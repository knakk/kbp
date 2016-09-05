package marc

import "bytes"

func mustDecode(input string) *Record {
	l := 64
	if len(input) < 64 {
		l = len(input)
	}

	format := DetectFormat([]byte(input[:l]))
	switch format {
	case MARC, LineMARC, MARCXML:
		break
	default:
		panic("mustDeocde: Unknown MARC format")
	}

	// Decode whole input stream
	dec := NewDecoder(bytes.NewBufferString(input), format)
	recs, err := dec.DecodeAll()
	if err != nil {
		panic("mustDecode: " + err.Error())
	}

	// We only want one record
	if len(recs) != 1 {
		panic("mustDecode: expected one record")
	}
	return recs[0]
}
