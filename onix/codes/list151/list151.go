package list151

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	BornIn             = "01"
	DiedIn             = "02"
	FormerlyResidedIn  = "03"
	CurrentlyResidesIn = "04"
	EducatedIn         = "05"
	WorkedIn           = "06"
	FlourishedIn       = "07"
	CitizenOf          = "08"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		BornIn:             {"Born in", "", "Født i", ""},
		DiedIn:             {"Died in", "", "Døde i", ""},
		FormerlyResidedIn:  {"Formerly resided in", "", "Bodde tidligere i", ""},
		CurrentlyResidesIn: {"Currently resides in", "", "Bor i", ""},
		EducatedIn:         {"Educated in", "", "Utdannet i", ""},
		WorkedIn:           {"Worked in", "", "Arbeidet i", ""},
		FlourishedIn:       {"Flourished in", "", "Blomstret i", ""},
		CitizenOf:          {"Citizen of", "Or nationality. For use with country codes only.", "Innbygger i", "Eller nasjonalitet. Skal kun brukes med landkoder"},
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
