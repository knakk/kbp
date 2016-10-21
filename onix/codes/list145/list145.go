package list145

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Preview            = "01"
	Print              = "02"
	CopyPaste          = "03"
	Share              = "04"
	TextToSpeech       = "05"
	Lend               = "06"
	TimeLimitedLicense = "07"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Preview:            {"Preview", "Preview before purchase.", "Forhåndsvisning", "Forhåndsvisning før kjøp"},
		Print:              {"Print", "Print paper copy of extract.", "Utskrift", "Utskrift av (utdrag av) innholdet"},
		CopyPaste:          {"Copy / paste", "Make digital copy of extract.", "Kopier/lim inn", "Lag digital kopi av (utdrag av) innholdet"},
		Share:              {"Share", "Share product across multiple concurrent devices.", "Del", "Del produktet samtidig på ulike eneheter"},
		TextToSpeech:       {"Text to speech", "‘Read aloud’ with text to speech functionality.", "Tekst til tale", "‘Opplesning’ med tekst til tale-funksjonalitet"},
		Lend:               {"Lend", "Lendable to other device owner or account holder, eg ‘Lend-to-a-friend’, library lending. The ‘primary’ copy becomes unusable while the secondary copy is ‘on loan’ unless a number of concurrent borrowers is also specified).", "Lån", "Utlån til eiere av andre enheter eller kontoer er mulig, f.eks. ‘Lån-til-en-venn', bibliotekutlån. Det opprinnelige eksemplaret kan ikke brukes når en kopi av dette er lånt bort, bortsett fra de tilfellene hvor et antall samtidige lånere er spesifisert"},
		TimeLimitedLicense: {"Time-limited license", "E-publication license is time limited. Use with 02 from List 146 and a number of days in <EpubUsageLimit>.", "Tidsbegrenset lisens", "Tidsbegrenset lisens til en epublikasjon. Brukes med 02 fra liste 146 og antall dager i   <EpubUsageLimit>"},
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
