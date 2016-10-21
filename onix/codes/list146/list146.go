package list146

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	PermittedUnlimited      = "01"
	PermittedSubjectToLimit = "02"
	Prohibited              = "03"
)

var (
	sortedCodes = []string{"01", "02", "03"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		PermittedUnlimited:      {"Permitted unlimited", "", "Ubegrenset", ""},
		PermittedSubjectToLimit: {"Permitted subject to limit", "Limit should be specified in <EpubUsageLimit>.", "Tillatt, med begrensninger", "Begrensninger bør spesifiseres i <EpubUsageLimit>."},
		Prohibited:              {"Prohibited", "", "Forbudt", ""},
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
