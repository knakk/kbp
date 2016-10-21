package list93

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Unspecified                                  = "00"
	PublisherToRetailers                         = "01"
	PublishersExclusiveDistributorToRetailers    = "02"
	PublishersNonExclusiveDistributorToRetailers = "03"
	Wholesaler                                   = "04"
	SalesAgent                                   = "05"
	PublishersDistributorToRetailers             = "06"
	PODSupplier                                  = "07"
	Retailer                                     = "08"
	PublisherToEndCustomers                      = "09"
	ExclusiveDistributorToEndCustomers           = "10"
	NonExclusiveDistributorToEndCustomers        = "11"
	DistributorToEndCustomers                    = "12"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Unspecified:                                  {"Unspecified", "Default.", "Uspesifisert", "Default"},
		PublisherToRetailers:                         {"Publisher to retailers", "Publisher as supplier to retail trade outlets.", "Forlag", "Forlag som leverandør/distributør til forhandler"},
		PublishersExclusiveDistributorToRetailers:    {"Publisher’s exclusive distributor to retailers", "", "Forlagets enedistributør", "Forlagets eksklusive distributør til forhandler"},
		PublishersNonExclusiveDistributorToRetailers: {"Publisher’s non-exclusive distributor to retailers", "", "Distributør for forlaget", "Forlagets distributør til forhandler, ikke-eksklusiv"},
		Wholesaler:                            {"Wholesaler", "Wholesaler supplying retail trade outlets.", "Grossist", "Grossist som leverandør/distributør til forhandler"},
		SalesAgent:                            {"Sales agent", "DEPRECATED – use <MarketRepresentation> (ONIX 2.1) or <MarketPublishingDetail> (ONIX 3.0) to specify a sales agent.", "Salgsagent", "UTGÅTT: bruk <MarketRepresentation> (Onix 2.1) eller <MarketPublishingDetail> (Onix 3.0)"},
		PublishersDistributorToRetailers:      {"Publisher’s distributor to retailers", "In a specified supply territory. Use only where exclusive/non-exclusive status is not known. Prefer 02 or 03 as appropriate, where possible.", "Forlagets distributør", "I et spesifisert territorium. Brukes bare når ekslusiv/ikke-ekslusiv status er ukjent. Bruk 02 eller 03 når det er mulig."},
		PODSupplier:                           {"POD supplier", "Where a POD product is supplied to retailers and/or consumers direct from a POD source.", "POD-leverandør", "Når et print-on-demand-produkt leveres til forhandler og/eller konsument direkte fra POD kilde."},
		Retailer:                              {"Retailer", "", "Retailer", ""},
		PublisherToEndCustomers:               {"Publisher to end-customers", "Publisher as supplier direct to consumers and/or institutional customers.", "Forlag til sluttkunder", "Forlag som leverandør/distributør til sluttkunder og institusjoner"},
		ExclusiveDistributorToEndCustomers:    {"Exclusive distributor to end-customers", "Intermediary as exclusive distributor direct to consumers and/or institutional customers.", "Eksklusiv distributør til sluttkunder", "Mellomledd som fungerer som eksklusiv distributør til sluttkunder og/eller institusjoner"},
		NonExclusiveDistributorToEndCustomers: {"Non-exclusive distributor to end-customers", "Intermediary as non-exclusive distributor direct to consumers and/or institutional customers.", "Ikke-eksklusiv distributør til sluttkunder", "Mellomledd som fungerer som ikke-eksklusiv distributør til sluttkunder og/eller institusjoner"},
		DistributorToEndCustomers:             {"Distributor to end-customers", "Use only where exclusive/non-exclusive status is not known. Prefer 10 or 11 as appropriate, where possible.", "Distributør til sluttkunder", "Brukes bare når ekslusiv/ikke-ekslusiv status er ukjent. Bruk 10 eller 11 når det er mulig"},
	}
)

// Items returns alle the possible items in this list, with labels and description in the requested language.
func Items(lang codes.Language) (res []codes.Item) {
	for _, code := range sortedCodes {
		if lang == codes.Norwegian {
			res = append(res, codes.Item{Code: code, Label: all[code].labelNo, Notes: all[code].notesNo})
		} else {
			res = append(res, codes.Item{Code: code, Label: all[code].labelEn, Notes: all[code].notesEn})
		}
	}
	return res
}

// Item return the Item for the given code and language, or an error if it doesn't exist.
func Item(code string, lang codes.Language) (codes.Item, error) {
	item, ok := all[code]
	if !ok {
		return codes.Item{}, fmt.Errorf("no item with code: %q", code)
	}
	if lang == codes.Norwegian {
		return codes.Item{Code: code, Label: item.labelNo, Notes: item.notesNo}, nil
	}
	return codes.Item{Code: code, Label: item.labelEn, Notes: item.notesEn}, nil
}

// MustItem returns the Item for the given code. It will panic if it doesn't exist.
func MustItem(code string, lang codes.Language) codes.Item {
	item, ok := all[code]
	if !ok {
		panic(fmt.Errorf("no item with code: %q", code))
	}
	if lang == codes.Norwegian {
		return codes.Item{Code: code, Label: item.labelNo, Notes: item.notesNo}
	}
	return codes.Item{Code: code, Label: item.labelEn, Notes: item.notesEn}

}
