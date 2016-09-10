package ncip

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"testing"
)

func mustEncodeRequest(req Request) string {
	output, err := xml.MarshalIndent(&NCIPRequestMessage{Request: req}, "", "  ")
	if err != nil {
		panic("mustEncodeRequest: " + err.Error())
	}

	return string(output)
}

func mustEncodeResponse(resp Response) string {
	output, err := xml.MarshalIndent(&NCIPResponseMessage{Response: resp}, "", "  ")
	if err != nil {
		panic("mustEncodeResponse: " + err.Error())
	}

	return string(output)
}

func readMsg(r *bufio.Reader) bytes.Buffer {
	var b bytes.Buffer
	for line, err := r.ReadBytes('\n'); err != io.EOF; line, err = r.ReadBytes('\n') {
		b.Write(line)
		if bytes.Equal(line, []byte("</NCIPMessage>\n")) {
			break
		}
	}
	return b
}

func TestRequestDecodeEncodeRoundtrip(t *testing.T) {
	f, err := os.Open("testdata/requests.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		b := readMsg(r)
		if b.Len() == 0 {
			break
		}
		want := b.String()
		want = want[:len(want)-1] // trim last newline
		req, err := DecodeRequest(&b)
		if err != nil {
			t.Fatal(err)
		}
		if got := mustEncodeRequest(req); want != got {
			t.Errorf("\n\ngot:\n%q\nwant:\n%q", got, want)
		}

	}

}

func TestResponseDecodeEncodeRoundtrip(t *testing.T) {
	f, err := os.Open("testdata/responses.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		b := readMsg(r)
		if b.Len() == 0 {
			break
		}
		want := b.String()
		want = want[:len(want)-1] // trim last newline
		resp, err := DecodeResponse(&b)
		if err != nil {
			t.Fatal(err)
		}
		if got := mustEncodeResponse(resp); want != got {
			t.Errorf("\n\ngot:\n%s\nwant:\n%s", got, want)
		}

	}

}
