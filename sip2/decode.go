package sip2

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Decode decodes a SIP message from a byte slice.
//
// Decode does not check if SIP message has all the required fields, other
// than an initial check of length to see if the message is long enough to
// contain the required fields for the given message type. Neither does
// it validate that the contents of a field conforms to specification,
// but fixed-length fields are guaranteed to have the required length.
// To validate a Message call the Message.Validate() function.
//
// Fields which are not defined by the SIP protocol are ignored.
func Decode(msg []byte) (Message, error) {
	var m Message

	l := len(msg)

	// trim trailing carriage return if present
	if l > 0 && msg[l-1] == '\r' {
		l--
		msg = msg[:l]
	}

	if l < 2 {
		return m, errors.New("Decode: message too short")
	}

	m.typ = codeToMsg[string(msg[:2])]
	if m.typ == MsgUnknown {
		return m, fmt.Errorf("Decode: unknown message code: %q", string(msg[:2]))
	}

	if l < minMsgLength[m.typ] {
		return m, fmt.Errorf("Decode: message too short to contain required fields for %v: %d < %d", m.typ, len(msg), minMsgLength[m.typ])
	}

	m.fields = make(map[fieldType]string)
	m.repeateableFields = make(map[fieldType][]string)

	// Parse fixed-length fields:
	p := 2 // byte position in message
	for _, f := range msgDefinitions[m.typ].RequiredFixed {
		end := p + fixedFieldLengths[f] // end of token
		m.fields[f] = string(msg[p:end])
		p = end
	}

	// Parse variable-length fields:
outer:
	for {
		start := p + 2 // start of current field
		f := codeToField[string(msg[p:start])]
		p = start
		if f == FieldUnknown {
			// store unknown codes in message value
			start -= 2
		}

		for {
			r, w := utf8.DecodeRune(msg[p:])
			p += w
			if r == '|' {
				if repeatableField[f] {
					m.repeateableFields[f] = append(m.repeateableFields[f], string(msg[start:p-1]))
				} else {
					m.fields[f] = string(msg[start : p-1])
				}
				if p == l {
					break outer
				}
				continue outer
			} else if p == l {
				if repeatableField[f] {
					m.repeateableFields[f] = append(m.repeateableFields[f], string(msg[start:l]))
				} else {
					m.fields[f] = string(msg[start:l])
				}
				break outer
			}

		}
	}

	return m, nil
}
