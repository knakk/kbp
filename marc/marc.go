package marc

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"sort"
	"strconv"
)

// Format represents a MARC serialization format.
type Format int

// Supported MARC serialization formats:
const (
	unknownFormat Format = iota
	MARC
	LineMARC
	MARCXML
)

// String returns a string representation of a Format.
func (f Format) String() string {
	switch f {
	case MARC:
		return "Standard MARC (ISO2709)"
	case LineMARC:
		return "Line-MARC"
	case MARCXML:
		return "MarcXchange (ISO25577)"
	default:
		return "Unknown MARC format"
	}
}

// Record represents a MARC record.
type Record struct {
	leader  []byte
	cfields map[ControlTag][]byte
	dfields map[DataTag][]*DataField
}

var emptyLeader = []byte("                        ") // 24 spaces

// NewRecord returns a new MARC Record.
func NewRecord() *Record {
	leader := make([]byte, 24)
	copy(leader, emptyLeader)
	return &Record{
		leader:  leader,
		cfields: make(map[ControlTag][]byte),
		dfields: make(map[DataTag][]*DataField),
	}
}

// String renders the Record as a readable string in a format similar to LineMARC.
func (r *Record) String() string {
	var b bytes.Buffer
	b.WriteString("000  ")
	b.Write(r.leader)
	b.WriteRune('\n')
	for i := 1; i < 9; i++ {
		tag := ControlTag(i)
		if val, ok := r.cfields[tag]; ok {
			b.WriteString("00")
			b.WriteString(strconv.Itoa(i))
			b.WriteString("  ")
			b.Write(val)
			b.WriteRune('\n')
		}
	}
	for i := 10; i < 1000; i++ {
		tag := DataTag(i)
		if dfs, ok := r.dfields[tag]; ok {
			for _, df := range dfs {
				if i < 100 {
					b.WriteString("0") // zero-pad
				}
				b.WriteString(strconv.Itoa(i))
				b.WriteRune(df.Indicator1)
				b.WriteRune(df.Indicator2)
				// TODO sort codes alphabetically
				indented := true
				for code, vals := range df.subfields {
					for _, v := range vals {
						if !indented {
							b.WriteString("\n     ")
						}
						b.WriteRune('$')
						b.WriteRune(code)
						b.WriteRune(' ')
						b.WriteString(v)
						indented = false
					}
				}
				b.WriteRune('\n')
			}
		}
	}
	return b.String()
}

// Marshal (currently) serializes a marcxml representation of the MARC record.
// TODO different formats
func (r *Record) Marshal(w io.Writer, f Format) error {
	// TODO w = errWriter(w)
	w.Write([]byte("<record><leader>"))
	w.Write(r.leader)
	w.Write([]byte("</leader>"))
	for tag, val := range r.cfields {
		w.Write([]byte(`<controlfield tag="`))
		w.Write([]byte(tag.String()))
		w.Write([]byte(`">`))
		w.Write(val)
		w.Write([]byte("</controlfield>"))
	}
	for tag, dfs := range r.dfields {
		for _, df := range dfs {
			w.Write([]byte(`<datafield ind1="`))
			w.Write([]byte(string(df.Indicator1)))
			w.Write([]byte(`" ind2="`))
			w.Write([]byte(string(df.Indicator2)))
			w.Write([]byte(`" tag="`))
			w.Write([]byte(tag.String()))
			w.Write([]byte(`">`))
			for code, sfs := range df.subfields {
				for _, sf := range sfs {
					w.Write([]byte(`<subfield code="`))
					w.Write([]byte(string(code)))
					w.Write([]byte(`">`))
					xml.EscapeText(w, []byte(sf))
					w.Write([]byte("</subfield>"))
				}
			}
			w.Write([]byte("</datafield>"))
		}
	}
	w.Write([]byte("</record>"))

	return nil
}

type jsonRecord struct {
	Leader  string                  `json:"leader"`
	CFields map[string]string       `json:"cfields,omitempty"`
	DFields map[string][]jsonDField `json:"dfields,omitempty"`
}

type jsonDField struct {
	Ind1      string              `json:"ind1"`
	Ind2      string              `json:"ind2"`
	Subfields map[string][]string `json:"subfields"`
}

// MarshalJSON serializes a JSON representation of the MARC record.
func (r *Record) MarshalJSON() ([]byte, error) {
	var jr jsonRecord
	jr.Leader = string(r.leader)
	jr.CFields = make(map[string]string, len(r.cfields))
	for k, v := range r.cfields {
		jr.CFields[k.String()] = string(v)
	}
	jr.DFields = make(map[string][]jsonDField, len(r.dfields))
	for k, dfs := range r.dfields {
		for _, df := range dfs {
			d := jsonDField{
				Ind1:      string(df.Indicator1),
				Ind2:      string(df.Indicator2),
				Subfields: make(map[string][]string),
			}
			for sf, sfv := range df.subfields {
				d.Subfields[string(sf)] = append(d.Subfields[string(sf)], sfv...)
			}
			jr.DFields[k.String()] = append(jr.DFields[k.String()], d)

		}
	}
	return json.Marshal(jr)
}

// SetLeaderPos sets the Record leader position to the given value.
//
// It is not possible to set the values for "Record length" or
// "Base address of data"; they are computed automatically when
// serializing a Record.
func (r *Record) SetLeaderPos(pos leaderPos, val byte) *Record {
	r.leader[int(pos)] = val
	return r
}

// AddControlField adds the given control field to the record,
// overwriting any existing control field with the same tag.
func (r *Record) AddControlField(f *ControlField) *Record {
	r.cfields[f.Tag] = f.value
	return r
}

// AddDataField adds the given data field to the record,
// appending to any existing data fields with the same tag.
func (r *Record) AddDataField(f *DataField) *Record {
	r.dfields[f.Tag] = append(r.dfields[f.Tag], f)
	return r
}

// ControlField returns the ControlField for the given tag, along
// with a boolean indicating whether the Record has that ControlField or not.
// If not, the ControlField will be empty.
func (r *Record) ControlField(tag ControlTag) (cf *ControlField, ok bool) {
	cf = NewControlField(tag)
	var b []byte
	if b, ok = r.cfields[tag]; ok {
		cf.value = b
	}
	return cf, ok
}

// DataField returns the DataField for the given tag, along
// with a boolean indicating whether the Record has that DataField or not.
// If not, the DataField will be empty.
// Only one DataField will be returned; if the field is repeated, use
// the DataFields() to retrieve all.
func (r *Record) DataField(tag DataTag) (*DataField, bool) {
	dfs, ok := r.dfields[tag]
	if !ok {
		return nil, ok
	}

	return dfs[0], true
}

// DataFields returns all the DataFields for the given tag.
func (r *Record) DataFields(tag DataTag) []*DataField {
	return r.dfields[tag]
}

// NewDataField return a new DataField.
func NewDataField(tag DataTag) *DataField {
	return &DataField{
		Tag:        tag,
		Indicator1: ' ',
		Indicator2: ' ',
		subfields:  make(map[rune][]string),
	}
}

// NewDataFieldWithIndicators returns a new DataField with specified Indicators.
func NewDataFieldWithIndicators(tag DataTag, ind1, ind2 rune) *DataField {
	return &DataField{
		Tag:        tag,
		Indicator1: ind1,
		Indicator2: ind2,
		subfields:  make(map[rune][]string),
	}
}

// DataField represents a DataField in a MARC record.
type DataField struct {
	Tag        DataTag
	Indicator1 rune
	Indicator2 rune
	subfields  map[rune][]string
}

// Add will add the given code,value pair to DataField's subfields.
func (df *DataField) Add(code rune, value string) *DataField {
	df.subfields[code] = append(df.subfields[code], value)
	return df
}

// Subfields returns the values set for a DataField's given subfield.
func (df *DataField) Subfields(code rune) []string {
	return df.subfields[code]
}

// Subfield returns a value for a DataField's given subfield, or an
// empty string if there aren't any.
func (df *DataField) Subfield(code rune) string {
	if v, ok := df.subfields[code]; ok {
		if len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

// ControlField represents a MARC control field. Control fields have
// no indicators or subfield codes, unlike Data fields.
type ControlField struct {
	Tag   ControlTag
	value []byte
}

// NewControlField returns a new ControlField.
func NewControlField(tag ControlTag) *ControlField {
	return &ControlField{Tag: tag}
}

// Set sets the ControlField value to the given string.
func (f *ControlField) Set(s string) *ControlField {
	f.value = []byte(s)
	return f
}

// SetPos inserts the given string in he ControlField value at
// the specified (byte) position.
func (f *ControlField) SetPos(pos int, s string) *ControlField {
	if l := pos + len(s); len(f.value) < l {
		b := make([]byte, l)
		for i := range b {
			b[i] = ' '
		}
		copy(b, f.value)
		f.value = b
	}
	copy(f.value[pos:], s)
	return f
}

// GetPos returns n characters starting at the given position in the ControlField.
func (f ControlField) GetPos(pos int, n int) string {
	if len(f.value) < pos+n {
		return ""
	}
	return string(f.value[pos : pos+n])
}

func (f ControlField) String() string {
	return string(f.value)
}

// Eq checks if two MARC records are equal.
func (r *Record) Eq(other *Record) bool {
	if !bytes.Equal(r.leader, other.leader) {
		return false
	}
	for tag, a := range r.cfields {
		b, ok := other.cfields[tag]
		if !ok {
			return false
		}
		if !bytes.Equal(a, b) {
			return false
		}
	}

	if len(r.dfields) != len(other.dfields) {
		return false
	}

	for tag, dfs := range r.dfields {
		if len(other.dfields[tag]) != len(dfs) {
			return false
		}

		sort.Sort(byTagAndSubfields(r.dfields[tag]))
		sort.Sort(byTagAndSubfields(other.dfields[tag]))

		for i, df := range dfs {
			if df.Indicator1 != other.dfields[tag][i].Indicator1 {
				return false
			}
			if df.Indicator2 != other.dfields[tag][i].Indicator2 {
				return false
			}
			for code, vals := range df.subfields {
				if len(vals) != len(other.dfields[tag][i].subfields[code]) {
					return false
				}
				for _, v := range vals {
					if !strInSlice(v, other.dfields[tag][i].subfields[code]) {
						return false
					}
				}
			}
		}
	}

	return true
}

func strInSlice(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// byTagAndSubfields is a wrapper type to sort DataFields.
type byTagAndSubfields []*DataField

func (b byTagAndSubfields) Len() int      { return len(b) }
func (b byTagAndSubfields) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b byTagAndSubfields) Less(i, j int) bool {
	// TODO revisit this function
	if b[i].Tag < b[j].Tag {
		return true
	}
	if b[i].Indicator1 < b[j].Indicator1 {
		return true
	}
	if b[i].Indicator2 < b[j].Indicator2 {
		return true
	}
	if len(b[i].subfields) < len(b[j].subfields) {
		return true
	}
	for code, vals := range b[i].subfields {
		if len(vals) < len(b[j].subfields[code]) {
			return true
		}
		// TODO sort vals
		if len(vals) < len(b[j].subfields[code]) {
			return true
		}
		for k, v := range vals {
			if k >= len(b[j].subfields[code]) {
				return true
			}
			if v < b[j].subfields[code][k] {
				return true
			}
		}
	}
	return false
}
