package list218

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	HumanReadable        = "01"
	ProfessionalReadable = "02"
	ONIXPL               = "10"
)

var (
	sortedCodes = []string{"01", "02", "10"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		HumanReadable:        {"Human readable", "Document (eg Word file, PDF or web page) Intended for the lay reader.", "For allmenn bruk", "Document (eg Word file, PDF or web page) Intended for the lay reader."},
		ProfessionalReadable: {"Professional readable", "Document (eg Word file, PDF or web page) Intended for the legal specialist reader.", "For juridiske eksperter", "Document (eg Word file, PDF or web page) Intended for the legal specialist reader."},
		ONIXPL:               {"ONIX-PL", "", "ONIX-PL", ""},
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
