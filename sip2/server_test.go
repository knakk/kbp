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
	// circulation state:
	catalogue map[string]string // barcode -> title
	checkouts map[string]string // barcode ->  patron id
	patrons   map[string]string // patron id -> name
}

func (h *testHandler) Handle(req Message) (resp Message) {
	switch req.Type() {
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
			resp.AddField(Field{Type: FieldRenewalOK, Value: "N"})
			resp.AddField(Field{Type: FieldDueDate, Value: time.Now().Format(DateLayout)})
		} else {
			// OK!
			resp.AddField(Field{Type: FieldOK, Value: "1"})
			resp.AddField(Field{Type: FieldRenewalOK, Value: "Y"})
			resp.AddField(Field{Type: FieldDueDate, Value: time.Now().AddDate(0, 1, 0).Format(DateLayout)})
			h.checkouts[barcode] = req.Field(FieldPatronIdentifier)
		}
		resp.AddField(
			Field{Type: FieldTitleIdentifier, Value: h.catalogue[barcode]})
	case MsgReqCheckin:
		// delete any checkouts
		delete(h.checkouts, req.Field(FieldItemIdentifier))
		resp = testMF.NewMessage(MsgRespCheckin).AddField(
			Field{Type: FieldOK, Value: "1"},
			Field{Type: FieldResentisize, Value: "Y"},
			Field{Type: FieldAlert, Value: "N"},
			Field{Type: FieldItemIdentifier, Value: req.Field(FieldItemIdentifier)},
			Field{Type: FieldPermanentLocation, Value: "here"})
	case MsgReqItemInformation:
		resp = testMF.NewMessage(MsgRespItemInformation).AddField(
			Field{Type: FieldTitleIdentifier, Value: h.catalogue[req.Field(FieldItemIdentifier)]},
		)
		_, checkedout := h.checkouts[req.Field(FieldItemIdentifier)]
		if checkedout {
			resp.AddField(Field{Type: FieldCirulationStatus, Value: "04"})
		} else {
			resp.AddField(Field{Type: FieldCirulationStatus, Value: "03"})
		}
	default:
		resp = testMF.NewMessage(MsgReqResend)
	}

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

	// 1) Item information (available)

	sendMsg(t, client,
		testMF.NewMessage(MsgReqItemInformation).AddField(
			Field{Type: FieldItemIdentifier, Value: "1"},
		))

	wantMsg(t, client, MsgRespItemInformation,
		Field{Type: FieldTitleIdentifier, Value: books[1]},
		Field{Type: FieldCirulationStatus, Value: "03"}, // = avialable, TODO constant
	)

	// 2) Checkout

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

	// 3) Checkout (allready checked out)

	sendMsg(t, client,
		testMF.NewMessage(MsgReqCheckout).AddField(
			Field{Type: FieldPatronIdentifier, Value: "0"},
			Field{Type: FieldItemIdentifier, Value: "1"},
		))

	wantMsg(t, client, MsgRespCheckout,
		Field{Type: FieldOK, Value: "0"},
		Field{Type: FieldPatronIdentifier, Value: "0"},
		Field{Type: FieldItemIdentifier, Value: "1"},
		Field{Type: FieldTitleIdentifier, Value: books[1]},
	)

	// 4) Item information (checked out)

	sendMsg(t, client,
		testMF.NewMessage(MsgReqItemInformation).AddField(
			Field{Type: FieldItemIdentifier, Value: "1"},
		))

	wantMsg(t, client, MsgRespItemInformation,
		Field{Type: FieldTitleIdentifier, Value: books[1]},
		Field{Type: FieldCirulationStatus, Value: "04"}, // = charged, TODO constant
	)

	// 5) Checkin

	sendMsg(t, client,
		testMF.NewMessage(MsgReqCheckin).AddField(
			Field{Type: FieldItemIdentifier, Value: "1"},
			Field{Type: FieldCurrentLocation, Value: "here"},
			Field{Type: FieldReturnDate, Value: time.Now().Format(DateLayout)},
		))

	wantMsg(t, client, MsgRespCheckin,
		Field{Type: FieldOK, Value: "1"},
		Field{Type: FieldItemIdentifier, Value: "1"},
	)

	// 6) Item information (available again)

	sendMsg(t, client,
		testMF.NewMessage(MsgReqItemInformation).AddField(
			Field{Type: FieldItemIdentifier, Value: "1"},
		))

	wantMsg(t, client, MsgRespItemInformation,
		Field{Type: FieldTitleIdentifier, Value: books[1]},
		Field{Type: FieldCirulationStatus, Value: "03"}, // = avialable, TODO constant
	)
}

func localAddr(s string) string {
	return s[strings.LastIndex(s, ":"):]
}
