package list48

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Height                        = "01"
	Width                         = "02"
	Thickness                     = "03"
	PageTrimHeight                = "04"
	PageTrimWidth                 = "05"
	UnitWeight                    = "08"
	DiameterSphere                = "09"
	UnfoldedUnrolledSheetHeight   = "10"
	UnfoldedUnrolledSheetWidth    = "11"
	DiameterTubeOrCylinder        = "12"
	RolledSheetPackageSideMeasure = "13"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "08", "09", "10", "11", "12", "13"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Height:                        {"Height", "For a book, the spine height when standing on a shelf. For a folded map, the height when folded. In general, the height of a product in the form in which it is presented or packaged for retail sale.", "Høyde", "For en bok, høyden på ryggen når den står i en hylle"},
		Width:                         {"Width", "For a book, the horizontal dimension of the cover when standing upright. For a folded map, the width when folded. In general, the width of a product in the form in which it is presented or packaged for retail sale.", "Bredde", "For en bok, den horisontale lengden på omslaget når den står oppreist"},
		Thickness:                     {"Thickness", "For a book, the thickness of the spine. For a folded map, the thickness when folded. In general, the thickness or depth of a product in the form in which it is presented or packaged for retail sale.", "Tykkelse", "For en bok, tykkelsen på bokryggen"},
		PageTrimHeight:                {"Page trim height", "Not recommended for general use.", "Skjærekant høyde", "Anbefales ikke brukt for annet enn spesielle formål"},
		PageTrimWidth:                 {"Page trim width", "Not recommended for general use.", "Skjærekant bredde", "Anbefales ikke brukt for annet enn spesielle formål"},
		UnitWeight:                    {"Unit weight", "", "Enhetsvekt", ""},
		DiameterSphere:                {"Diameter (sphere)", "Of a globe, for example.", "Diameter", "Av en globus for eksempel"},
		UnfoldedUnrolledSheetHeight:   {"Unfolded/unrolled sheet height", "The height of a folded or rolled sheet map, poster etc when unfolded.", "Utbrettet plakathøyde", "Høyden på brettet eller sammenrullet kart, plakat e.l. når det er brettet ut"},
		UnfoldedUnrolledSheetWidth:    {"Unfolded/unrolled sheet width", "The width of a folded or rolled sheet map, poster etc when unfolded.", "Utbrettet plakatbredde", "Bredden på brettet eller sammenfullet kart, plakat eller lignende når det er brettet ut"},
		DiameterTubeOrCylinder:        {"Diameter (tube or cylinder)", "The diameter of the cross-section of a tube or cylinder, usually carrying a rolled sheet product. Use 01 “height” for the height or length of the tube.", "Diameter (rør eller sylinder)", "Diamater på et rør eller en sylinder som vanligvis inneholder et sammenrullet produkt. Bruk 01 'høyde' for høyden eller lengden på røret."},
		RolledSheetPackageSideMeasure: {"Rolled sheet package side measure", "The length of a side of the cross-section of a long triangular or square package, usually carrying a rolled sheet product. Use 01 “height” for the height or length of the package.", "Rolled sheet package side measure", "The length of a side of the cross-section of a long triangular or square package, usually carrying a rolled sheet product. Use 01 'height' for the height or length of the package."},
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
