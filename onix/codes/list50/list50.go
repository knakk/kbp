package list50

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Centimeters = "cm"
	Grams       = "gr"
	InchesUS    = "in"
	Kilograms   = "kg"
	PoundsUS    = "lb"
	Millimeters = "mm"
	OuncesUS    = "oz"
	Pixels      = "px"
)

var (
	sortedCodes = []string{"cm", "gr", "in", "kg", "lb", "mm", "oz", "px"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Centimeters: {"Centimeters", "Millimeters are the preferred metric unit of length.", "Centimeter", "Millimeter er den foretrukne måleenheten"},
		Grams:       {"Grams", "", "Gram", ""},
		InchesUS:    {"Inches (US)", "", "Tommer (US)", ""},
		Kilograms:   {"Kilograms", "Grams are the preferred metric unit of weight.", "Kilograms", "Gram er den foretrukne måleenheten"},
		PoundsUS:    {"Pounds (US)", "", "Pund (US)", ""},
		Millimeters: {"Millimeters", "", "Millimeter", ""},
		OuncesUS:    {"Ounces (US)", "", "Unser (US)", ""},
		Pixels:      {"Pixels", "", "Pixels", ""},
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
