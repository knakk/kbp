package marc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func mustDecode(input string) *Record {
	r, err := decode([]byte(input))
	if err != nil {
		panic(err)
	}
	return r
}

func decode(input []byte) (*Record, error) {
	l := 64
	if len(input) < 64 {
		l = len(input)
	}

	format := DetectFormat(input[:l])
	switch format {
	case MARC, LineMARC, MARCXML:
		break
	default:
		return nil, errors.New("decodeFile: Unknown MARC format")
	}

	// Decode whole input stream
	dec := NewDecoder(bytes.NewReader(input), format)
	recs, err := dec.DecodeAll()
	if err != nil {
		return nil, err
	}

	// We only want one record
	if len(recs) != 1 {
		return nil, errors.New("decodeFile: expected one record")
	}
	return recs[0], nil
}

func decodeFile(path string) (*Record, error) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return decode(input)
}

// Verify that MARC records can be decoded, encoded and redecoded without alterations.
func TestDecodeEncodeRoundtrip(t *testing.T) {
	files, err := filepath.Glob("testdata/*")
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		r, err := decodeFile(file)
		if err != nil {
			t.Errorf("decodeFile(%q) => %v", file, err)
		}
		for _, f := range []Format{MARCXML} { // TODO: []Format{MARCXML, LineMARC, MARC}
			var b bytes.Buffer
			if err := r.Marshall(&b, f); err != nil {
				t.Errorf("Marshall error: %v", err)
			}
			got := mustDecode(b.String())
			if !got.Eq(r) {
				t.Errorf("%s: decode/encode roundtrip failed when record serialized as %v", file, f)
			}
		}
	}
}
