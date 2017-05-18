// package sip2 is a library for working with the SIP2 protocol used for
// communication between self-service circulation automats and library vendor systems.

package sip2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"
)

// DateLayout is the date time format that the SIP-protocol expects.
const DateLayout = "20060102    150405"

// A Message defines a request sent by the SC to the ACS, or the response
// sent by the ACS to the SC. The Message Type defines which kind it is.
type Message struct {
	typ               msgType
	fields            map[fieldType]string
	repeateableFields map[fieldType][]string
}

// Field is a component of a SIP message, consiting of a type and a string value.
type Field struct {
	Type  fieldType
	Value string
}

// NewMessage returns a new Message with the give message type.
func NewMessage(f msgType) Message {
	return Message{
		typ:               f,
		fields:            make(map[fieldType]string),
		repeateableFields: make(map[fieldType][]string),
	}
}

func (m Message) Type() msgType {
	return m.typ
}

// AddField adds a field to the Message. If the field is not repeatable, it
// will overwrite any existing value for the field.
func (m Message) AddField(fs ...Field) Message {
	for _, f := range fs {
		if repeatableField[f.Type] {
			m.repeateableFields[f.Type] = append(m.repeateableFields[f.Type], f.Value)
		} else {
			m.fields[f.Type] = f.Value
		}
	}
	return m
}

// String encodes the SIP message to a string.
func (m Message) String() string {
	var b bytes.Buffer
	m.Encode(&b)
	return b.String()
}

// Field returns the value of a field. An undefined field
// returns an empty string. There is no way of knowing if
// the string is defined or not. Use FieldOK for that.
// If the field is repeatable and several are defined,
// only one of them will be returned
func (m Message) Field(f fieldType) string {
	if repeatableField[f] && len(m.repeateableFields[f]) > 0 {
		return m.repeateableFields[f][0]
	}
	return m.fields[f]
}

// FieldFound returns the value of a field, along with a boolean stating if
// the field is defined or not. If the field is repeatable and several are
// defined, only one of them will be returned
func (m Message) FieldOK(f fieldType) (string, bool) {
	v, ok := m.fields[f]
	if !ok && repeatableField[f] {
		if len(m.repeateableFields[f]) > 0 {
			v = m.repeateableFields[f][0]
			ok = true
		}
	}
	return v, ok
}

// Fields returns the values defined for a reatable field, or nil if none are defined.
func (m Message) Fields(f fieldType) []string {
	res, ok := m.repeateableFields[f]
	if !ok {
		return nil
	}
	return res
}

// Encode encodes a SIP messgae to a stream.
// It will fail if there are missing required fields, or if
// writing to the stream fails.
// Unknown fields and defined fields which are not required or
// optional for the give type are not encoded.
func (m Message) Encode(w io.Writer) error {
	bw := bufio.NewWriter(w)
	if _, err := bw.WriteString(msgToCode[m.typ]); err != nil {
		return err
	}

	for _, f := range msgDefinitions[m.typ].RequiredFixed {
		if !m.hasField(f) {
			return fmt.Errorf("Encode: %v missing required fixed-length field: %v", m.typ, f)
		}

		if _, err := bw.WriteString(m.fields[f]); err != nil {
			return err
		}
	}

	for _, f := range msgDefinitions[m.typ].RequiredVar {
		if !m.hasField(f) { // TODO leave out?
			return fmt.Errorf("Encode: %v missing required variable-length field: %v", m.typ, f)
		}

		if _, err := bw.WriteString(fieldToCode[f]); err != nil {
			return err
		}

		if _, err := bw.WriteString(m.fields[f]); err != nil {
			return err
		}

		if _, err := bw.WriteRune('|'); err != nil {
			return err
		}
	}

	for _, f := range msgDefinitions[m.typ].OptionalVar {
		if !m.hasField(f) {
			continue
		}

		if _, err := bw.WriteString(fieldToCode[f]); err != nil {
			return err
		}

		if _, err := bw.WriteString(m.fields[f]); err != nil {
			return err
		}

		if _, err := bw.WriteRune('|'); err != nil {
			return err
		}
	}

	for f, fs := range m.repeateableFields {
		if !isOptional(m.typ, f) {
			continue
		}

		for _, s := range fs {
			if _, err := bw.WriteString(fieldToCode[f]); err != nil {
				return err
			}

			if _, err := bw.WriteString(s); err != nil {
				return err
			}

			if _, err := bw.WriteRune('|'); err != nil {
				return err
			}
		}

	}

	if _, err := bw.WriteRune('\r'); err != nil {
		return err
	}

	return bw.Flush()
}

func (m Message) hasField(f fieldType) bool {
	_, ok := m.fields[f]
	return ok
}

// Validate validates a SIP message. It returns a slice of strings
// listing the violations of the SIP protocol for the given message type.
func (m Message) Validate() []string {
	errs := []string{}
	if m.typ == MsgUnknown {
		errs = append(errs, "unknown message type")
		return errs
	}

	for _, f := range msgDefinitions[m.typ].RequiredFixed {
		if !m.hasField(f) {
			errs = append(errs, fmt.Sprintf("missing required fixed-length field: %v", f))
			continue
		}
		if rxp := fieldValdiation[f]; rxp != nil {
			if v, ok := m.fields[f]; ok && !rxp.MatchString(v) {
				errs = append(errs, fmt.Sprintf("fixed-length field %v with value %q does not match %v",
					f, v, rxp))
			}
		}
	}

	for _, f := range msgDefinitions[m.typ].RequiredVar {
		if !m.hasField(f) {
			errs = append(errs, fmt.Sprintf("missing required variable-length field: %v", f))
			continue
		}
		if rxp := fieldValdiation[f]; rxp != nil {
			if v, ok := m.fields[f]; ok && !rxp.MatchString(v) {
				errs = append(errs, fmt.Sprintf("variable-length field %v with value %q does not match %v",
					f, v, rxp))
			}
		}
	}

	// TODO repeatable fields are never required?
	/*for _, f := range m.RepeateableFields {

	}*/

	return errs
}

// MessageFactory can generate Message with default values set.
// Some fields will always have the same value in a lot of
// configurations, so this is provided as a convenience in those cases.
type MessageFactory struct {
	defaults map[fieldType]string
}

// NewMessageFactory returns a MessageFacyory with the given default
// fields, which will be used creating messages with NewMessage.
func NewMessageFactory(defaults ...Field) MessageFactory {
	mf := MessageFactory{
		defaults: make(map[fieldType]string, len(defaults)),
	}

	for _, f := range defaults {
		mf.defaults[f.Type] = f.Value
	}

	return mf
}

// NewMessage returns a new Message, with the default values defined in the MessageFactory, but
// only when applicable to the message type.
//
// In addition, the fields FieldDateTieSync, FieldTransactionDate and FieldNbDueDate are set with
// the current timestamp when required by given the message type.
func (mf MessageFactory) NewMessage(t msgType) Message {
	now := time.Now().Format(DateLayout)

	msg := NewMessage(t)
	switch msg.typ {
	case MsgReqStatus:
		msg.AddField(Field{Type: FieldDateTimeSync, Value: time.Now().Format(DateLayout)})
	case MsgReqResend, MsgReqLogin, MsgRespStatus, MsgRespLogin:
		// No timestamp
	case MsgReqCheckin:
		msg.AddField(Field{Type: FieldReturnDate, Value: time.Now().Format(DateLayout)})
		fallthrough
	case MsgReqCheckout, MsgReqRenew:
		msg.AddField(Field{Type: FieldNbDueDate, Value: now})
		fallthrough
	default:
		// Matches all except those types in the first case statement
		msg.AddField(Field{Type: FieldTransactionDate, Value: now})
	}

	for _, f := range msgDefinitions[msg.typ].RequiredFixed {
		if def, ok := mf.defaults[f]; ok {
			msg.AddField(Field{Type: f, Value: def})
		}
	}

	for _, f := range msgDefinitions[msg.typ].RequiredVar {
		if def, ok := mf.defaults[f]; ok {
			msg.AddField(Field{Type: f, Value: def})
		}
	}

	for _, f := range msgDefinitions[msg.typ].OptionalVar {
		if def, ok := mf.defaults[f]; ok {
			msg.AddField(Field{Type: f, Value: def})
		}
	}

	return msg
}

// msgType represents the request or response message type.
type msgType int

// All possible message types
const (
	MsgUnknown msgType = iota
	// Requests:
	MsgReqPatronStatus      // 23
	MsgReqCheckout          // 11
	MsgReqCheckin           // 09
	MsgReqBlockPatron       // 01
	MsgReqStatus            // 99
	MsgReqResend            // 97
	MsgReqLogin             // 93
	MsgReqPatronInformation // 63
	MsgReqEndPatronSession  // 35
	MsgReqFeePaid           // 37
	MsgReqItemInformation   // 17
	MsgReqItemStatusUpdate  // 19
	MsgReqPatronEnable      // 25
	MsgReqHold              // 15
	MsgReqRenew             // 29
	MsgReqRenewAll          // 65
	// Responses:
	MsgRespPatronStatus      // 24
	MsgRespCheckout          // 12
	MsgRespCheckin           // 10
	MsgRespStatus            // 98
	MsgRespLogin             // 94
	MsgRespPatronInformation // 64
	MsgRespEndPatronSession  // 36
	MsgRespFeePaid           // 38
	MsgRespItemInformation   // 18
	MsgRespItemStatusUpdate  // 20
	MsgRespPatronEnable      // 26
	MsgRespHold              // 16
	MsgRespRenew             // 30
	MsgRespRenewAll          // 66
)

// fieldType represents a request/response message field ID.
type fieldType int

// All possible field types:
const (
	FieldUnknown fieldType = iota

	// Fixed length fields:
	FieldAlert                 // 1 char: Y or N
	FieldAvialable             // 1 char: Y or N
	FieldCardRetained          // 1 char: Y or N
	FieldChargedItemsCount     // 4 char: 0000 to 9999
	FieldFineItemsCount        // 4 char: 0000 to 9999
	FieldHoldItemsCount        // 4 char: 0000 to 9999
	FieldOverdueItemsCount     // 4 char: 0000 to 9999
	FieldRecallItemsCount      // 4 char: 0000 to 9999
	FieldUnavailableHoldsCount // 4 char: 0000 to 9999
	FieldCheckinOK             // 1 char: Y or N
	FieldCheckoutOK            // 1 char: Y or N
	FieldCirulationStatus      // 2 char: 00 to 99 (1-13 are defined by the protocol)
	FieldDateTimeSync          // 18 char: YYYYMMDDZZZZHHMMSS
	FieldDesentisize           // 1 char: Y or N
	FieldEndSession            // 1 char: Y or N
	FieldHoldMode              // 1 char: + or - or *
	FieldItemPropertiesOK      // 1 char: '1' = OK, other = not NOK
	FieldLanguage              // 3 char: 000 - 999 (some language codes defined by the protocol)
	FieldMagneticMedia         // 1 char: Y or N or U
	FieldMaxPrintWidth         // 3 char: 000 to 999
	FieldNbDueDate             // 1 char: Y or N
	FieldNoBlock               // 1 char: Y or N
	FieldOffLineOK             // 1 char: Y or N
	FieldOK                    // 1 char: '1' = OK, other = not OK
	FieldOnLineStatus          // 1 char: Y or N
	FieldPatronStatus          // 14 char: each char is Y or N
	FieldPaymentAccepted       // 1 char: Y or N
	FieldPaymentType           // 2 char: 00=cash, 01=VISA, 02=credit card
	FieldProtocolVersion       // 4 char: x.xx
	FieldPWDAlgorithm          // 1 char: '0' = not encrypted
	FieldRenewalOK             // 1 char: Y or N
	FieldRenewedCount          // 4 char: 0000 to 9999
	FieldResentisize           // 1 char: Y or N
	FieldRetriesAllowed        // 3 char: ?
	FieldReturnDate            // 18 char: YYYYMMDDZZZZHHMMSS
	FieldRenewalPolicy         // 1 char: Y or N
	FieldSecurityMarker        // 2 char: 00=other, 01=none, 02,03=3M specific
	FieldStatusCode            // 1 char: 0=OK, 1=printer out of paper, 2=about to shut down
	FieldStatusUpdateOK        // 1 char: Y or N
	FieldSummary               // 10 char:
	FieldThirdPartyAllowd      // 1 char: Y or N
	FieldTimeoutPeriod         // 3 char: 000 - 999
	FieldTransactionDate       // 18 char: YYYYMMDDZZZZHHMMSS
	FieldUIDAlgorithm          // 1 char: 0=not encrypted
	FieldUnrenewedCount        // 4 char: 0000 to 9999

	// Variable length fields:
	FieldPatronIdentifier      // AA
	FieldItemIdentifier        // AB
	FieldTerminalPassword      // AC
	FieldPatronPassword        // AD
	FieldPersonalName          // AE
	FieldScreenMessage         // AF
	FieldPrintLine             // AG
	FieldDueDate               // AH
	FieldTitleIdentifier       // AJ
	FieldBlockedCardMsg        // AL
	FieldLibraryName           // AM
	FieldTerminalLocation      // AN
	FieldInstitutionID         // AO
	FieldCurrentLocation       // AP
	FieldPermanentLocation     // AQ
	FieldHoldItems             // AS
	FieldOverdueItems          // AT
	FieldChargedItems          // AU
	FieldFineItems             // AV
	FieldHomeAddress           // BD
	FieldEmailAddress          // BE
	FieldHomePhoneNumber       // BF
	FieldOwner                 // BG
	FieldCurrenyType           // BH
	FieldCancel                // BI
	FieldTransactionID         // BK
	FieldValidPatron           // BL
	FieldRenewedItems          // BM
	FieldUnrenewedItems        // BN
	FieldFeeAcknowledged       // BO
	FieldStartItem             // BP
	FieldEndItem               // BQ
	FieldQueuePosition         // BR
	FieldPickupLocation        // BS
	FieldFeeType               // BT
	FieldRecallItems           // BU
	FieldFeeAmount             // BV
	FieldExpirationDate        // BW
	FieldSupportedMessages     // BX
	FieldHoldType              // BY
	FieldHoldItemsLimit        // BZ
	FieldOverdueItemsLimit     // CA
	FieldChargedItemsLimit     // CB
	FieldFeeLimit              // CC
	FieldUnavailableHoldsItems // CD
	FieldHoldQueueLength       // CF
	FieldFeeIdentifier         // CG
	FieldItemProperties        // CH
	FieldSecurityInhibit       // CI
	FieldRecallDate            // CJ
	FieldMediaType             // CK
	FieldSortBin               // CL
	FieldHoldPickupDate        // CM
	FieldLoginUserID           // CN
	FieldLoginPassword         // CO
	FieldLocationCode          // CP
	FieldValidPatronPassword   // CQ
	FieldSequenceNumber        // AY
	FieldChecksum              // AZ
)
