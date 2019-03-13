package oai

import "time"

// Error represents error returned from a OAI endpoint. This only includes
// protocol level errors, and not errors originating from the transport layer (io).
type Error struct {
	code string
}

// Defined OAI protocol errors.
var (
	ErrBadArgument             = Error{code: "badArgument"}
	ErrBadResumptionToken      = Error{code: "badResumptionToken"}
	ErrBadVerb                 = Error{code: "badVerb"}
	ErrCannotDisseminateFormat = Error{code: "cannotDisseminateFormat"}
	ErrNoMetadataFormats       = Error{code: "noMetadataFormats"}
	ErrNoSetHierarchy          = Error{code: "noSetHierarchy"}
)

// Error return a string representation of an Error.
func (e Error) Error() string {
	return "oai: " + e.code
}

func errorFromCode(code string) Error {
	return Error{code: code}
}

// Record represents a OAI record.
type Record struct {
	Header struct {
		Status     string    `xml:"status,attr"`
		Identifier string    `xml:"identifier"`
		Datestamp  time.Time `xml:"datestamp"`
	} `xml:"header"`
	Metadata []byte `xml:",innerxml"`
}

type response struct {
	Error struct {
		Code    string `xml:"code,attr"`
		Message string `xml:",chardata"`
	} `xml:"error"`
	ListRecords struct {
		Records         []Record `xml:"record"`
		ResumptionToken string   `xml:"resumptionToken"`
	} `xml:"ListRecords,omitempty"`
}
