package list159

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Application = "01"
	Audio       = "02"
	Image       = "03"
	Text        = "04"
	Video       = "05"
	MultiMode   = "06"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Application: {"Application", "An executable together with data on which it operates.", "Applikasjon", "An executable together with data on which it operates."},
		Audio:       {"Audio", "A sound recording.", "Lyd", "Lydinnspilling"},
		Image:       {"Image", "A still image.", "Bilde", "Stillbilde"},
		Text:        {"Text", "Readable text, with or without associated images etc.", "Tekst", "Lesbar tekst, med eller uten bilder osv"},
		Video:       {"Video", "Moving images, with or without accompanying sound.", "Video", "Med eller uten tilhørende lyd"},
		MultiMode:   {"Multi-mode", "A website or other supporting resource delivering content in a variety of modes.", "Multi-mode", "Nettsted eller annen tilleggsressurs som har innhold på ulike former og formater"},
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
