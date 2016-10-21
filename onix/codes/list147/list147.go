package list147

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Copies                  = "01"
	Characters              = "02"
	Words                   = "03"
	Pages                   = "04"
	Percentage              = "05"
	Devices                 = "06"
	ConcurrentUsers         = "07"
	PercentagePerTimePeriod = "08"
	Days                    = "09"
	Times                   = "10"
	AllowedUsageStartPage   = "11"
	AllowedUsageEndPage     = "12"
	Weeks                   = "13"
	Months                  = "14"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Copies:                  {"Copies", "Maximum number of copies that may be made of a permitted extract.", "Kopier", "Maks antall kopier som kan lages av et tillatt utdrag"},
		Characters:              {"Characters", "Maximum number of characters in a permitted extract for a specified usage.", "Tegn", "Maks antall tegn som er tillatt brukt i et utdrag for spesifisert bruk"},
		Words:                   {"Words", "Maximum number of words in a permitted extract for a specified usage.", "Ord", "Maks antall ord som er tillatt brukt i et utdrag for spesifisert bruk"},
		Pages:                   {"Pages", "Maximum number of pages in a permitted extract for a specified usage.", "Sider", "Maks antall sider som er tillatt brukt i et utdrag for spesifisert bruk"},
		Percentage:              {"Percentage", "Maximum percentage of total content in a permitted extract for a specified usage.", "Prosent", "Maks prosentandel av det totale innholdet som er tillatt brukt i et utdrag for spesifisert bruk."},
		Devices:                 {"Devices", "Maximum number of devices in ‘share group’.", "Enheter", "Maks antall enheter i ei gruppe"},
		ConcurrentUsers:         {"Concurrent users", "Maximum number of concurrent users. NB where the number of concurrent users is specifically not limited, set the number of concurrent users to zero.", "Samtidige brukere", "Maks antall samtidige brukere. NB: dersom antall samtidige brukere ikke er spesifikt begrenset, sett antall samtidige brukere til null"},
		PercentagePerTimePeriod: {"Percentage per time period", "Maximum percentage of total content which may be used in a specified usage per time period; the time period being specified as another EpubUsageQuantity.", "Prosentandel av tidsperiode", "Maks prosentandel av det totale innholdet som er tillatt brukt et antall ganger innen en tidsperiode: tidsperioden spesifiseres i EpubUsageQuantity."},
		Days:  {"Days", "Maximum time period in days.", "Dager", "Maks lengde på tidsperiode, angitt i dager"},
		Times: {"Times", "Maximum number of times a specified usage may occur.", "Ganger", "Maks antall ganger en spesifisert bruk kan forekomme"},
		AllowedUsageStartPage: {"Allowed usage start page", "Page number where allowed usage begins. <Quantity> should contain an absolute page number, counting the cover as page 1. (This type of page numbering should not be used where the e-publication has no fixed pagination). Use with (max number of) Pages, Percentage of content, or End page to specify pages allowed in Preview.", "Startside for tillatt bruk", "Sidetall som angir hvor tillatt bruk starter. <Quantity>  bør inneholde det totale sideantallet, hvor omslaget regnes som side 1 (Denne typen sidenummerering bør ikke brukes for epublikasjoner som ikke har fast sideangivelse). Brukes med (maks antall) sider, prosentandel av innholdet, eller sidetall for siste side for å spesifisere hvilke sider som er tillatt for forhåndsvisning"},
		AllowedUsageEndPage:   {"Allowed usage end page", "Page number at which allowed usage ends. <Quantity> should contain an absolute page number, counting the cover as page 1. (This type of page numbering should not be used where the e-publication has no fixed pagination). Use with Start page to specify pages allowed in a preview.", "Sluttside for tillatt bruk", "Sidetall som angir når tillatt bruk slutter.  <Quantity>  bør inneholde det totale sideantallet, hvor omslaget regnes som side 1 (Denne typen sidenummerering bør ikke brukes for epublikasjoner som ikke har fast sideangivelse). Brukes med startside for å spesifisere hvilke sider som er tillatt brukt for forhåndsvisning."},
		Weeks:                 {"Weeks", "Maximum time period in weeks.", "Uker", "Maks lengde på tidsperiode, angitt i uker"},
		Months:                {"Months", "Maximum time period in months.", "Måneder", "Maks lengde på tidsperiode, angitt i måneder"},
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
