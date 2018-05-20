package oai

import "time"

type Record struct {
	Header struct {
		Identifier string    `xml:"identifier"`
		Datestamp  time.Time `xml:"datestamp"`
	} `xml:"header"`
	Metadata []byte `xml:",innerxml"`
}

type response struct {
	ListRecords struct {
		Records         []Record `xml:"record"`
		ResumptionToken string   `xml:"resumptionToken"`
	} `xml:"ListRecords,omitempty"`
}
