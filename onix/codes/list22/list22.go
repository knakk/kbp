package list22 // Language role code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Language role code
const (
	LanguageOfText                           = "01"
	OriginalLanguageOfATranslatedText        = "02"
	LanguageOfAbstracts                      = "03"
	RightsLanguage                           = "04"
	RightsExcludedLanguage                   = "05"
	OriginalLanguageInAMultilingualEdition   = "06"
	TranslatedLanguageInAMultilingualEdition = "07"
	LanguageOfAudioTrack                     = "08"
	LanguageOfSubtitles                      = "09"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		LanguageOfText:                           {"Language of text", "", "Tekstens språk", ""},
		OriginalLanguageOfATranslatedText:        {"Original language of a translated text", "Where the text in the original language is NOT part of the current product.", "Originalspråk", "Where the text in the original language is NOT part of the current product."},
		LanguageOfAbstracts:                      {"Language of abstracts", "Where different from language of text: used mainly for serials.", "Språk i sammendrag", "Når det er annet språk på utdrag enn på hovedteksten: Mest aktuelt for tidsskrifter"},
		RightsLanguage:                           {"Rights language", "Language to which specified rights apply.", "Språk med spesifiserte rettighetsvilkår", "Språk som det er knyttet spesielle rettigheter til"},
		RightsExcludedLanguage:                   {"Rights-excluded language", "Language to which specified rights do not apply.", "Språk hvor spesifiserte rettigheter ikke gjelder", "Språk der spesifiserte rettigheter ikke gjelder"},
		OriginalLanguageInAMultilingualEdition:   {"Original language in a multilingual edition", "Where the text in the original language is part of a bilingual or multilingual edition.", "Originalspråket i en flerspråklig utgave", "Når teksten på orginalspråket er del av en to- eller flerspråklig utgave"},
		TranslatedLanguageInAMultilingualEdition: {"Translated language in a multilingual edition", "Where the text in a translated language is part of a bilingual or multilingual edition.", "Oversatt språk i en flerspråklig utgave", "Når den oversatte teksten er del av en to- eller flerspråklig utgave"},
		LanguageOfAudioTrack:                     {"Language of audio track", "For example, on a DVD.", "Språk på lydspor", "For eksempel på en DVD"},
		LanguageOfSubtitles:                      {"Language of subtitles", "For example, on a DVD.", "Språk på undertekst", "For eksempel på en DVD"},
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
