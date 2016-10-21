package list154

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Unrestricted              = "00"
	Restricted                = "01"
	Booktrade                 = "02"
	EndCustomers              = "03"
	Librarians                = "04"
	Teachers                  = "05"
	Students                  = "06"
	Press                     = "07"
	ShoppingComparisonService = "08"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Unrestricted: {"Unrestricted", "Any audience.", "Uten restriksjoner", "Ikke begrenset til et spesielt publikum"},
		Restricted:   {"Restricted", "Distribution by agreement between the parties to the ONIX exchange (this value is provided to cover applications where ONIX content includes material which is not for general distribution).", "Restriksjoner", "Distribueres etter avtale mellom de partene som utveksler Onix (denne koden sendes for å dekke applikasjoner hvor Onix-innholdet inkluderer materiale som ikke er for generell distribusjon."},
		Booktrade:    {"Booktrade", "Distributors, bookstores, publisher’s own staff etc.", "Bokbransjen", "Distributører, bokhandlere, forlagsansatte osv."},
		EndCustomers: {"End-customers", "", "Sluttkunder", ""},
		Librarians:   {"Librarians", "", "Bibliotekarer", ""},
		Teachers:     {"Teachers", "", "Lærere", ""},
		Students:     {"Students", "", "Studenter", ""},
		Press:        {"Press", "Press or other media.", "Presse", "Presse / media"},
		ShoppingComparisonService: {"Shopping comparison service", "Where a specially formatted description is required for this audience.", "Shopping comparison service", "Hvor det er krav om en spesielt formatert beskrivelse for dette pulikummet"},
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
