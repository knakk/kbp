package list92

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Proprietary                         = "01"
	Proprietary                         = "02"
	BörsenvereinVerkehrsnummer          = "04"
	GermanISBNAgencyPublisherIdentifier = "05"
	GLN                                 = "06"
	SAN                                 = "07"
	DistributeurscodeBoekenbank         = "12"
	FondscodeBoekenbank                 = "13"
	VATIdentityNumber                   = "23"
)

var (
	sortedCodes = []string{"01", "02", "04", "05", "06", "07", "12", "13", "23"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Proprietary:                         {"Proprietary", "", "Proprietær", ""},
		Proprietary:                         {"Proprietary", "DEPRECATED – use 01.", "Proprietær", "UTGÅTT - bruk 01"},
		BörsenvereinVerkehrsnummer:          {"Börsenverein Verkehrsnummer", "", "Börsenverein Verkehrsnummer", ""},
		GermanISBNAgencyPublisherIdentifier: {"German ISBN Agency publisher identifier", "", "German ISBN Agency publisher identifier", ""},
		GLN: {"GLN", "GS1 global location number (formerly EAN location number).", "EAN-UCC GLN", "GS1 global location number (formerly EAN location number)."},
		SAN: {"SAN", "Book trade Standard Address Number – US, UK etc.", "SAN", "Book trade Standard Address Number - US, UK etc"},
		DistributeurscodeBoekenbank: {"Distributeurscode Boekenbank", "Flemish supplier code.", "Distributeurscode Boekenbank", "Flamsk distributørkode"},
		FondscodeBoekenbank:         {"Fondscode Boekenbank", "Flemish publisher code.", "Fondscode Boekenbank", "Flamsk forlagskode"},
		VATIdentityNumber:           {"VAT Identity Number", "Identifier for a business organization for VAT purposes, eg within the EU’s VIES system. See http://ec.europa.eu/taxation_customs/vies/faqvies.do for EU VAT ID formats, which vary from country to country. Generally these consist of a two-letter country code followed by the 8–12 digits of the national VAT ID. Some countries include one or two letters within their VAT ID. See http://en.wikipedia.org/wiki/VAT_identification_number for non-EU countries that maintain similar identifiers. Spaces, dashes etc should be omitted.", "VAT Identity Number", "Identifier for a business organization for VAT purposes, eg within the EU’s VIES system. See http://ec.europa.eu/taxation_customs/vies/faqvies.do for EU VAT ID formats, which vary from country to country. Generally these consist of a two-letter country code followed by the 8–12 digits of the national VAT ID. Some countries include one or two letters within their VAT ID. See http://en.wikipedia.org/wiki/VAT_identification_number for non-EU countries that maintain similar identifiers. Spaces, dashes etc should be omitted."},
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
