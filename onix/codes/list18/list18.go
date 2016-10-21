package list18

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Unspecified                     = "00"
	Pseudonym                       = "01"
	AuthorityControlledName         = "02"
	EarlierName                     = "03"
	RealName                        = "04"
	TransliteratedFormOfPrimaryName = "05"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Unspecified:                     {"Unspecified", "", "Uspesifisert", ""},
		Pseudonym:                       {"Pseudonym", "May be used to give a well-known pseudonym, where the primary name is a ‘real’ name.", "Pseudonym", "May be used to give a well-known pseudonym, where the primary name is a ‘real’ name."},
		AuthorityControlledName:         {"Authority-controlled name", "", "Autoritetskontrollert navn", ""},
		EarlierName:                     {"Earlier name", "", "Tidligere navn", ""},
		RealName:                        {"‘Real’ name", "May be used to identify a well-known real name, where the primary name is a pseudonym.", "Virkelig navn", "Kan brukes til å identifisere et kjent, virkelig navn, hvor det primære navnet er et pseudonym"},
		TransliteratedFormOfPrimaryName: {"Transliterated form of primary name", "Use only within <AlternativeName>, when the primary name type is unspecified.", "Transkripsjon fra primært navn", "Brukes kun i <AlternativeName>, når primær navnetype er uspesifisert"},
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
