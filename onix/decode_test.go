package onix

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestDecodeEncodeRoundtrip(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		want, err := ioutil.ReadFile(filepath.Join("testdata", file.Name()))
		if err != nil {
			t.Fatal(err)
		}
		var msg Message

		if err := xml.Unmarshal(want, &msg); err != nil {
			t.Fatal(err)
		}

		got, err := xml.MarshalIndent(msg, "  ", "  ")
		if err != nil {
			t.Fatal(err)
		}

		if err := ignoreWhiteSpaceEqual(want, got); err != nil {
			t.Fatalf("%s: %v", file.Name(), err)
		}
	}

}

// ignoreWhiteSpaceEqual returns an error with the diff of the first which is
// different in b from a, ignoring whitespace at the beginning and end of string.
func ignoreWhiteSpaceEqual(a, b []byte) error {
	alines := bytes.Split(a, []byte("\n"))
	blines := bytes.Split(b, []byte("\n"))
	for i, line := range alines {
		aline := bytes.TrimSpace(line)
		if len(blines) <= i {
			return fmt.Errorf("number of lines in b < %d", i+1)
		}
		bline := bytes.TrimSpace(blines[i])
		if !bytes.Equal(aline, bline) {
			return fmt.Errorf("%d: %q != %q", i+1, string(aline), string(bline))
		}
	}
	return nil
}
