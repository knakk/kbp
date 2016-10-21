package list148 // Collection type

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Collection type
const (
	UnspecifiedDefault  = "00"
	PublisherCollection = "10"
	AscribedCollection  = "20"
)

var (
	sortedCodes = []string{"00", "10", "20"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		UnspecifiedDefault:  {"Unspecified (default)", "Collection type is not determined.", "Uspesifisert (standard)", "Collection type er ikke angitt"},
		PublisherCollection: {"Publisher collection", "The collection is a bibliographic collection (eg a series) defined and identified by a publisher, either on the product itself or in product information supplied by the publisher.", "Forlagsserie / Forlagsdefinert samling", "Eng: publisher collection.  Samlinga (eller serien) er definert og identifisert av forlaget, enten ved at den er trykt/angitt på selve produktet eller i annen produktinformasjon som forlaget leverer"},
		AscribedCollection:  {"Ascribed collection", "The collection has been defined and identified by a party in the metadata supply chain other than the publisher, typically an aggregator.", "Annen tilknytning / Tillagt samling", "Eng: ascribed collection. Samlinga har blitt definert og identifisert av en annen part i verdikjeden enn forlaget, typisk av en metadataprodusent"},
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
