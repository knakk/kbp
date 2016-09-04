package sip2

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

var testMF = NewMessageFactory(
	Field{Type: FieldRenewalPolicy, Value: "Y"}, // TODO constants
	Field{Type: FieldNoBlock, Value: "N"},       // TODO constants
	Field{Type: FieldInstitutionID, Value: "Acme co."},
	Field{Type: FieldTerminalPassword, Value: ""},
	Field{Type: FieldSecurityMarker, Value: "01"}, // TODO constants
	Field{Type: FieldFeeType, Value: "01"},        // TODO constants
	Field{Type: FieldMagneticMedia, Value: "N"},
	Field{Type: FieldDesentisize, Value: "N"},
)

type testHandler struct {
	lastMessage Message

	// circulation state:
	catalogue map[string]string // barcode -> title
	checkouts map[string]string // barcode ->  patron id
	patrons   map[string]string // patron id -> name
}

func (h *testHandler) Handle(req Message) (resp Message) {
	switch req.Type() {
	case MsgReqPatronStatus:
		resp = testMF.NewMessage(MsgRespPatronStatus)
	case MsgReqCheckout:
		barcode := req.Field(FieldItemIdentifier)
		resp = testMF.NewMessage(MsgRespCheckout).AddField(
			Field{Type: FieldPatronIdentifier, Value: req.Field(FieldPatronIdentifier)},
			Field{Type: FieldItemIdentifier, Value: barcode},
			// TODO make convenience function: Message.CopyFrom(msg, ..fields)
		)
		_, checkedout := h.checkouts[barcode]
		if checkedout {
			// allready checked out, decline
			resp.AddField(Field{Type: FieldOK, Value: "0"})
		} else {
			// OK!
			resp.AddField(Field{Type: FieldOK, Value: "1"})
			resp.AddField(Field{Type: FieldRenewalOK, Value: "N"})
			resp.AddField(Field{Type: FieldDueDate, Value: time.Now().AddDate(0, 1, 0).Format(DateLayout)})
		}
		resp.AddField(
			Field{Type: FieldTitleIdentifier, Value: h.catalogue[barcode]})
	case MsgReqCheckin:
		resp = testMF.NewMessage(MsgRespCheckin)
	case MsgReqBlockPatron:
		resp = testMF.NewMessage(MsgRespPatronStatus)
	case MsgReqStatus:
		resp = testMF.NewMessage(MsgRespStatus)
	case MsgReqResend:
		resp = h.lastMessage
	case MsgReqLogin:
		resp = testMF.NewMessage(MsgRespLogin)
	case MsgReqPatronInformation:
		resp = testMF.NewMessage(MsgRespPatronInformation)
	case MsgReqEndPatronSession:
		resp = testMF.NewMessage(MsgRespEndPatronSession)
	case MsgReqFeePaid:
		resp = testMF.NewMessage(MsgRespFeePaid)
	case MsgReqItemInformation:
		resp = testMF.NewMessage(MsgRespItemInformation).AddField(
			Field{Type: FieldCirulationStatus, Value: "03"}, // TODO status constants
			Field{Type: FieldTitleIdentifier, Value: h.catalogue[req.Field(FieldItemIdentifier)]},
		)
	case MsgReqItemStatusUpdate:
		resp = testMF.NewMessage(MsgRespItemStatusUpdate)
	case MsgReqPatronEnable:
		resp = testMF.NewMessage(MsgRespPatronEnable)
	case MsgReqHold:
		resp = testMF.NewMessage(MsgRespHold)
	case MsgReqRenew:
		resp = testMF.NewMessage(MsgRespRenew)
	case MsgReqRenewAll:
		resp = testMF.NewMessage(MsgRespRenewAll)
	default:
		resp = testMF.NewMessage(MsgReqResend)
	}

	h.lastMessage = resp
	return resp
}

func newTestHandler(books []string, patrons []string) *testHandler {
	h := testHandler{
		patrons:   make(map[string]string, len(patrons)),
		catalogue: make(map[string]string, len(books)),
		checkouts: make(map[string]string, len(books)),
	}
	for i, book := range books {
		h.catalogue[strconv.Itoa(i)] = book
	}
	for i, patron := range patrons {
		h.patrons[strconv.Itoa(i)] = patron
	}
	return &h
}

func sendMsg(t *testing.T, c net.Conn, msg Message) {
	if err := msg.Encode(c); err != nil {
		t.Fatalf("sendMsg: %v", err)
	}
}

func getMsg(t *testing.T, c net.Conn, typ msgType) Message {
	r := bufio.NewReader(c)
	b, err := r.ReadBytes('\r')
	if err != nil {
		t.Fatalf("getMsg: %v", err)
	}
	msg, err := Decode(b)
	if err != nil {
		t.Fatal(err)
	}
	if msg.typ != typ {
		t.Errorf("getMsg: %v; want %v", msg.typ, typ)
	}
	return msg
}

func wantMsg(t *testing.T, c net.Conn, typ msgType, fs ...Field) {
	msg := getMsg(t, c, typ)
	for _, f := range fs {
		if msg.Field(f.Type) != f.Value {
			t.Errorf("%v: got %q; want %q", f.Type, msg.Field(f.Type), f.Value)
		}
	}
}

func TestTransactionSession(t *testing.T) {

	// Setup

	books := []string{"Sult - Knut Hamsun", "Hvite Netter - Fjodor Dostovjevski", "Norsk-engelsk ordbok"}
	patrons := []string{"Ole Jensen", "Petter Åsen", "Mia Hansen"}
	h := newTestHandler(books, patrons)

	s, err := NewServer(h, 0)
	if err != nil {
		t.Fatal(err)
	}

	srvAddr := localAddr(s.ln.Addr().String())
	go s.Run()
	defer s.Close()

	client, err := net.Dial("tcp", srvAddr)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Interaction

	// 1) Item information

	sendMsg(t, client,
		testMF.NewMessage(MsgReqItemInformation).AddField(
			Field{Type: FieldItemIdentifier, Value: "1"},
		))

	wantMsg(t, client, MsgRespItemInformation,
		Field{Type: FieldTitleIdentifier, Value: books[1]},
	)

	// 2) Checkouts

	sendMsg(t, client,
		testMF.NewMessage(MsgReqCheckout).AddField(
			Field{Type: FieldPatronIdentifier, Value: "0"},
			Field{Type: FieldItemIdentifier, Value: "1"},
		))

	wantMsg(t, client, MsgRespCheckout,
		Field{Type: FieldOK, Value: "1"},
		Field{Type: FieldPatronIdentifier, Value: "0"},
		Field{Type: FieldItemIdentifier, Value: "1"},
		Field{Type: FieldTitleIdentifier, Value: books[1]},
	)

	// 3) Checkins

	/*

		resp := getMsg(t, client)
		wantFields(t, resp,
			Fields{Type: FieldOK, Value: "1"})
	*/

}

func localAddr(s string) string {
	return s[strings.LastIndex(s, ":"):]
}
