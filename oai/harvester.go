package oai

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// Harvester is a OAI harvester.
type Harvester struct {
	endpoint string
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
	}
}

func (h *Harvester) fetch() ([]Record, error) {
	url := h.endpoint
	if h.token != "" {
		url += "&resumptionToken=" + h.token
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

	// Store the resumptionToken.
	h.token = oaiResponse.ListRecords.ResumptionToken

	return oaiResponse.ListRecords.Records, nil
}
