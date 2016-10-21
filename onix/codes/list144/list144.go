package list144

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	None                = "00"
	DRM                 = "01"
	DigitalWatermarking = "02"
	AdobeDRM            = "03"
	AppleDRM            = "04"
	OMADRM              = "05"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		None:                {"None", "Has no technical protection.", "Ingen", "Har ingen teknisk beskyttelse"},
		DRM:                 {"DRM", "Has DRM protection.", "DRM", "Har DRM (kopibeskyttelse)"},
		DigitalWatermarking: {"Digital watermarking", "Has digital watermarking.", "Vannmerking", "Har digital vannmerking"},
		AdobeDRM:            {"Adobe DRM", "Has DRM protection applied by the Adobe CS4 Content Server Package or by the Adobe ADEPT hosted service.", "Adobe DRM", "Har DRM-beskyttelse (Adobe CS4 Content Server Package eller Adobe ADEPT hosted service)"},
		AppleDRM:            {"Apple DRM", "FairPlay’ DRM protection applied via Apple proprietary online store.", "Apple DRM", "FairPlay’ DRM (Apple proprietary online store)"},
		OMADRM:              {"OMA DRM", "Has OMA v2 DRM protection applied, as used to protect some mobile phone content.", "OMA DRM", "Har OMA v2 DRM-beskyttelse. For å beskytte innhold på mobiltelefoner"},
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
