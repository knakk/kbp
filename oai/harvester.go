package oai

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

// Harvester is a OAI harvester.
type Harvester struct {
	endpoint       string
	Token          string // resumptionToken
	From           time.Time
	Set            string
	Process        func(Record) error
	MetadataPrefix string
}

// NewHarvester returns a new Harvester targeting the given endpoint url.
func NewHarvester(url string) *Harvester {
	return &Harvester{
		endpoint: url,
	}
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
			if err := h.Process(rec); err != nil {
				return err
			}
		}

		if h.Token == "" {
			// ResumptionToken is empty, which means we have harvested all records.
			return nil
		}
	}
}

func (h *Harvester) fetch() ([]Record, error) {
	url := h.endpoint + "?verb=ListRecords&metadataPrefix=" + h.MetadataPrefix
	if h.Token != "" {
		url += "&resumptionToken=" + h.Token
	}
	if h.Set != "" {
		url += "&set=" + h.Set
	}
	if !h.From.IsZero() {
		url += "&from=" + h.From.String()
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
	h.Token = oaiResponse.ListRecords.ResumptionToken

	return oaiResponse.ListRecords.Records, nil
}
