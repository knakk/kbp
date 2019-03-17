package oai

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

// ProcessStatus represents the status of a record processing.
type ProcessStatus uint

// Available ProcessStatuses
const (
	StautsNew ProcessStatus = iota
	StatusUpdated
	StatusDeleted
	StatusFailed
)

// Harvester is a OAI harvester.
type Harvester struct {
	endpoint       string
	Token          string // resumptionToken
	From           time.Time
	Set            string
	Process        func(Record) (ProcessStatus, error)
	MetadataPrefix string
}

// NewHarvester returns a new Harvester targeting the given endpoint url.
func NewHarvester(url string) *Harvester {
	return &Harvester{
		endpoint: url,
	}
}

// HarvestStats keeps statistics about a harvest run.
type HarvestStats struct {
	NumAdded         int
	NumUpdated       int
	NumDeleted       int
	NumProcessFailed int
}

// Run will harvest from records from an OAI endpoint. If the configured processing
// function returns an error, it will return. Otherwise it will keep going until
// all records have been harvested, and then return harest statistics.
func (h *Harvester) Run() (HarvestStats, error) {
	stats := HarvestStats{}
	for {
		records, err := h.fetch()
		if err != nil {
			return stats, err
		}
		if len(records) == 0 {
			return stats, nil
		}
		for _, rec := range records {
			status, err := h.Process(rec)
			if err != nil {
				return stats, err
			}
			switch status {
			case StautsNew:
				stats.NumAdded++
			case StatusDeleted:
				stats.NumDeleted++
			case StatusUpdated:
				stats.NumUpdated++
			case StatusFailed:
				stats.NumProcessFailed++
			default:
				panic(fmt.Sprintf("unknown ProcessStatus %v", status))
			}
		}

		if h.Token == "" {
			// ResumptionToken is empty, which means we have harvested all records.
			return stats, nil
		}
	}
}

func (h *Harvester) fetch() ([]Record, error) {
	url := h.endpoint + "?verb=ListRecords"
	if h.Token != "" {
		url += "&resumptionToken=" + h.Token
	} else {
		url += "&metadataPrefix=" + h.MetadataPrefix
		if h.Set != "" {
			url += "&set=" + h.Set
		}
		if !h.From.IsZero() {
			url += "&from=" + h.From.Format("2006-01-02")
		}
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
