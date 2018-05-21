package oai

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

// Harvester is a OAI harvester.
type Harvester struct {
	endpoint string
	from     time.Time
	token    string // resumptionToken
	process  func(Record) error
}

// NewHarvester returns a new Harvester targeting the given endpoint url and
// using the given function to process the records.
func NewHarvester(url string, process func(r Record) error) *Harvester {
	return &Harvester{
		endpoint: url + "?verb=ListRecords&metadataPrefix=marcxchange",
		process:  process,
	}
}

// WithToken returns a Harvester to be started with the given resumptionToken.
func (h *Harvester) WithToken(s string) *Harvester {
	h.token = s
	return h
}

// From returns a Harvester to be started harvesting from the given timestamp.
func (h *Harvester) From(from time.Time) *Harvester {
	h.from = from
	return h
}

// Run will harvest from records from an OAI endpoint. If the configured processing
// function returns an error, it will return. Otherwise it will keep going until
// all records have been harvested, and then return nil.
func (h *Harvester) Run() error {
	for {
		records, err := h.fetch()
		if err != nil {
			return err
		}
		if len(records) == 0 {
			return nil
		}
		for _, rec := range records {
			if err := h.process(rec); err != nil {
				return err
			}
		}

		if h.token == "" {
			// ResumptionToken is empty, which means we have harvested all records.
			return nil
		}
	}
}

func (h *Harvester) fetch() ([]Record, error) {
	url := h.endpoint
	if h.token != "" {
		url += "&resumptionToken=" + h.token
	}
	if !h.from.IsZero() {
		url += "&from=" + h.from.String()
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/xml")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch oai reponse: %v", resp.Status)
	}
	var oaiResponse response
	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(&oaiResponse); err != nil {
		return nil, err
	}

	// Return an error if there is any OAI error in the response
	if errCode := oaiResponse.Error.Code; errCode != "" {
		return nil, errorFromCode(errCode)
	}

	// Store the resumptionToken (even if empty string).
	h.token = oaiResponse.ListRecords.ResumptionToken

	return oaiResponse.ListRecords.Records, nil
}
