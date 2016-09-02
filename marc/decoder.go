package marc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode/utf8"
)

const (
	linemarcRT  = 0x5E // ^ (record terminator)
	linemarcFS  = 0x2A // * (field separator)
	linemarcSFS = 0x24 // $ (subfield separator)
)

type Decoder struct {
	input  *bufio.Reader
	format Format
	lineN  int
}

func NewDecoder(r io.Reader, f Format) *Decoder {
	switch f {
	case MARC, MARCXML, LineMARC:
	default:
		panic("Cannot decode unknown MARC Format")
	}

	return &Decoder{
		input:  bufio.NewReader(r),
		format: f,
	}
}

// DecodeAll consumes the input stream and returns all decoded records.
// If there is an error, it will return, together with the succesfully
// parsed MARC records up til then.
func (d *Decoder) DecodeAll() ([]*Record, error) {
	res := make([]*Record, 0)
	for r, err := d.Decode(); err != io.EOF; r, err = d.Decode() {
		if err != nil {
			return res, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (d *Decoder) Decode() (*Record, error) {
	switch d.format {
	case LineMARC:
		return d.decodeLineMARC()
	case MARCXML:
		panic("TODO decode MARCXML")
	case MARC:
		panic("TODO decode ISOMARC")
	default:
		panic("Cannot decode unknown MARC Format")
	}
}

func (d *Decoder) decodeLineMARC() (r *Record, err error) {
	r = NewRecord()
decodeRecord:
	for {
		line, err := d.input.ReadBytes('\n')
		if err != nil && len(line) == 0 {
			return r, err
		}
		d.lineN++

		switch line[0] {
		case linemarcRT:
			break decodeRecord
		case '\n':
			continue decodeRecord // could be extra newline between records
		case linemarcFS:
			// OK
		default:
			err = fmt.Errorf("%d:0: expected '*', found %q", d.lineN, string(line))
			break
		}

		// strip newline
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}

		if len(line) < 4 || !isAllDigits(line[1:4]) {
			err = fmt.Errorf(
				"%d:0: expected a valid control field or data field number, found %q",
				d.lineN, string(line))
			break
		}

		p := 1 // position in line

		if bytes.HasPrefix(line[p:], []byte("00")) {
			if (line[p+2]) == '0' {
				// Parse leader
				copy(r.leader, line[p+3:])
			} else {
				// Parse Control field
				cField, err2 := parseControlField(line[p:])
				if err != nil {
					err = err2
					break
				}
				r.AddControlField(cField)
			}
		} else {
			// Parse Data field
			dField, err2 := parseDataField(line[p:])
			if err2 != nil {
				err = err2
				break
			}
			r.AddDataField(dField)
		}

	}
	return r, err
}

// DetectFormat tries to detect the MARC encoding of the given byte slice. It
// detects one of LineMARC/MARC/MARCXML, otherwise unknown.
func DetectFormat(data []byte) Format {
	// Find the first non-whitespace byte
	i := 0
	for ; i < len(data) && isWS(data[i]); i++ {
	}

	switch data[i] {
	case '<':
		return MARCXML
	case '*': // TODO, also '^' ?
		return LineMARC
	default:
		if data[i] >= '0' && data[i] <= '9' {
			return MARC
		}
		return unknownFormat
	}
}

// Mappings
var (
	str2controlfields = map[string]ControlTag{
		"001": Tag001, "003": Tag003, "005": Tag005, "006": Tag006, "007": Tag007, "008": Tag008, "009": Tag009,
	}

	controlfields2str = map[ControlTag]string{
		Tag001: "001", Tag003: "003", Tag005: "005", Tag006: "006", Tag007: "007", Tag008: "008", Tag009: "009",
	}

	str2datafields = map[string]DataTag{
		"010": Tag010, "020": Tag020, "041": Tag041, "111": Tag111, "245": Tag245, "260": Tag260,
		"300": Tag300, "440": Tag440, "508": Tag508, "546": Tag546, "700": Tag700,
		"911": Tag911,
	}

	datafield2str = map[DataTag]string{
		Tag010: "010", Tag020: "020", Tag041: "041", Tag111: "111", Tag245: "245", Tag260: "260",
		Tag300: "300", Tag440: "440", Tag508: "508", Tag546: "546", Tag700: "700",
		Tag911: "911",
	}
)

// Parsing helper functions:

func parseControlField(b []byte) (ControlField, error) {
	// We can asume that len(b) >= 3, and that b[0:3] is numeric,
	// and that b[3] != '0'.
	f := ControlField{
		Tag:   str2controlfields[string(b[:3])],
		value: b[3:],
	}
	return f, nil
}

func parseDataField(b []byte) (DataField, error) {
	// We can asume that len(b) >= 3, and that b[0:3] is numeric.
	f := DataField{
		Tag:       str2datafields[string(b[:3])],
		subfields: make(map[rune][]string),
	}
	if len(b) >= 5 {
		f.Indicator1 = rune(b[3])
		f.Indicator2 = rune(b[4])
	}

	p := 4 // position in b
subfields:
	for p+2 < len(b) {
		p += 2       // dollarsign + code
		code := b[p] // TODO use DecodeRune; Bibsys uses multibyte char like 'ø' as subfield code
		p++
		start := p
		for {
			ch, w := utf8.DecodeRune(b[p:])
			p += w

			if ch == linemarcSFS {
				f.Add(rune(code), string(b[start:p-1]))
				p -= 2
				continue subfields
			}

			if w == 0 { // eof
				if p > start {
					f.Add(rune(code), string(b[start:p]))
					break subfields
				}
			}
		}
	}

	return f, nil
}

func isWS(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAllDigits(b []byte) bool {
	for _, c := range b {
		if !isDigit(c) {
			return false
		}
	}
	return true
}
