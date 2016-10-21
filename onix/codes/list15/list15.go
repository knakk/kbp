package list15 // Title type code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Title type code
const (
	Undefined                                                                          = "00"
	DistinctiveTitleBookCoverTitleSerialTitleOnItemSerialContentItemOrReviewedResource = "01"
	ISSNKeyTitleOfSerial                                                               = "02"
	TitleInOriginalLanguage                                                            = "03"
	TitleAcronymOrInitialism                                                           = "04"
	AbbreviatedTitle                                                                   = "05"
	TitleInOtherLanguage                                                               = "06"
	ThematicTitleOfJournalIssue                                                        = "07"
	FormerTitle                                                                        = "08"
	DistributorsTitle                                                                  = "10"
	AlternativeTitleOnCover                                                            = "11"
	AlternativeTitleOnBack                                                             = "12"
	ExpandedTitle                                                                      = "13"
	AlternativeTitle                                                                   = "14"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "10", "11", "12", "13", "14"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Undefined: {"Undefined", "", "Udefinert", ""},
		DistinctiveTitleBookCoverTitleSerialTitleOnItemSerialContentItemOrReviewedResource: {"Distinctive title (book); Cover title (serial); Title on item (serial content item or reviewed resource)", "The full text of the distinctive title of the item, without abbreviation or abridgement. For books, where the title alone is not distinctive, elements may be taken from a set or series title and part number etc to create a distinctive title. Where the item is an omnibus edition containing two or more works by the same author, and there is no separate combined title, a distinctive title may be constructed by concatenating the individual titles, with suitable punctuation, as in “Pride and prejudice / Sense and sensibility / Northanger Abbey”.", "Fullstendig tittel (bok eller serie)", "Hele teksten til tittelen, uten forkortelser eller lignende Når en tittel alene ikke er distinkt, kan elementer hentes fra serietittel og -nummer for å lage en distinkt tittel. Når en utgave er en sammenstilling som inneholder flere verk av samme forfatter og det ikke er en egen tittel på utgivelsen, kan en distinkt tittel lages ved å sammenstille de individuelle titlene med passende skilletegn, som i Pride and prejudice / Sense and sensibility / Northanger Abbey."},
		ISSNKeyTitleOfSerial:        {"ISSN key title of serial", "Serials only.", "ISSN-nøkkeltittel for tidskrift", "Kun for tidsskrifter/seriepublikasjoner"},
		TitleInOriginalLanguage:     {"Title in original language", "Where the subject of the ONIX record is a translated item.", "Originaltittel", "For oversatte bøker"},
		TitleAcronymOrInitialism:    {"Title acronym or initialism", "For serials: an acronym or initialism of Title Type 01, eg “JAMA”, “JACM”.", "Tittelakronym", "Kun for tidsskrifter: JACM = Journal of the Association for Computing Machinery"},
		AbbreviatedTitle:            {"Abbreviated title", "An abbreviated form of Title Type 01.", "Forkortet tittel", "En forkortet form av tittel type 01"},
		TitleInOtherLanguage:        {"Title in other language", "A translation of Title Type 01 into another language.", "Tittel på annet språk", "En oversettelse av tittel type 01 til et annet språk"},
		ThematicTitleOfJournalIssue: {"Thematic title of journal issue", "Serials only: when a journal issue is explicitly devoted to a specified topic.", "Tematisk tittel til en tidsskriftsutgave", "Kun for tidsskrifter: når en tidsskriftsutgave er eksplisitt om et spesifikt emne"},
		FormerTitle:                 {"Former title", "Books or serials: when an item was previously published under another title.", "Tidligere tittel", "Når en bok eller tidskriftsutgave tidligere har blitt publisert med en annen tittel"},
		DistributorsTitle:           {"Distributor’s title", "For books: the title carried in a book distributor’s title file: frequently incomplete, and may include elements not properly part of the title.", "Distributørens tittel", "For bøker: tittelen i distributørens system. Ofte ikke komplett og kan ha elementer som ikke er del av tittelen"},
		AlternativeTitleOnCover:     {"Alternative title on cover", "An alternative title that appears on the cover of a book.", "Alternativ tittel på omslaget", "En alternativ tittel på bokens omslag (forskjellig fra tittelsiden)"},
		AlternativeTitleOnBack:      {"Alternative title on back", "An alternative title that appears on the back of a book.", "Alternativ tittel på baksiden", "En alternativ tittel på baksiden (forskjellig fra tittelsiden)"},
		ExpandedTitle:               {"Expanded title", "An expanded form of the title, eg the title of a school text book with grade and type and other details added to make the title meaningful, where otherwise it would comprise only the curriculum subject. This title type is required for submissions to the Spanish ISBN Agency.", "Utvidet tittel", "En utvidet tittelform på en skolebok med f.eks. trinn og andre detaljer som gir mening hvor det ellers bare ville ha stått fag. Denne type tittel er påkrevet ved registrering ved det spanske ISBN-kontoret"},
		AlternativeTitle:            {"Alternative title", "An alternative title that the book is widely known by, whether it appears on the book or not", "Alternativ tittel", "En alternativ tittelen som brukes om boka, uavhengig av om tittelen forekommer på boka"},
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
