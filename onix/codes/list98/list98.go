package list98

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Black                 = "BLK"
	Blue                  = "BLU"
	Brown                 = "BRN"
	BurgundyMaroon        = "BUR"
	Cream                 = "CRE"
	FourColor             = "FCO"
	FourColorAndSpotColor = "FCS"
	Gold                  = "GLD"
	Green                 = "GRN"
	Grey                  = "GRY"
	Multicolor            = "MUL"
	NavyDarkBlue          = "NAV"
	Orange                = "ORG"
	Pink                  = "PNK"
	Purple                = "PUR"
	Red                   = "RED"
	SkyPaleBlue           = "SKY"
	Silver                = "SLV"
	TanLightBrown         = "TAN"
	TealTurquoiseGreen    = "TEA"
	White                 = "WHI"
	Yellow                = "YEL"
	Other                 = "ZZZ"
)

var (
	sortedCodes = []string{"BLK", "BLU", "BRN", "BUR", "CRE", "FCO", "FCS", "GLD", "GRN", "GRY", "MUL", "NAV", "ORG", "PNK", "PUR", "RED", "SKY", "SLV", "TAN", "TEA", "WHI", "YEL", "ZZZ"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Black:                 {"Black", "", "Svart", ""},
		Blue:                  {"Blue", "", "Blå", ""},
		Brown:                 {"Brown", "", "Brun", ""},
		BurgundyMaroon:        {"Burgundy/maroon", "", "Burgunder", ""},
		Cream:                 {"Cream", "", "Krem", ""},
		FourColor:             {"Four-color", "", "Four-color", ""},
		FourColorAndSpotColor: {"Four-color and spot-color", "", "Four-color and spot-color", ""},
		Gold:               {"Gold", "", "Gull", ""},
		Green:              {"Green", "", "Grønn", ""},
		Grey:               {"Grey", "", "Grå", ""},
		Multicolor:         {"Multicolor", "Use <ProductFormFeatureDescription> to add brief details if required.", "Flerfarget", "Bruk <ProductFormFeatureDescription> for å legge til detaljer der det er nødvendig"},
		NavyDarkBlue:       {"Navy/Dark blue", "", "Mørk blå", ""},
		Orange:             {"Orange", "", "Oransje", ""},
		Pink:               {"Pink", "", "Rosa", ""},
		Purple:             {"Purple", "", "Fiolett", ""},
		Red:                {"Red", "", "Rød", ""},
		SkyPaleBlue:        {"Sky/Pale blue", "", "Lys blå", ""},
		Silver:             {"Silver", "", "Sølv", ""},
		TanLightBrown:      {"Tan/Light brown", "", "Lys brun", ""},
		TealTurquoiseGreen: {"Teal/Turquoise green", "", "Turkis", ""},
		White:              {"White", "", "Hvit", ""},
		Yellow:             {"Yellow", "", "Gul", ""},
		Other:              {"Other", "", "Andre", ""},
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
