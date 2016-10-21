package list29

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	ONIXAudienceCodes                          = "01"
	Proprietary                                = "02"
	MPAARating                                 = "03"
	BBFCRating                                 = "04"
	FSKRating                                  = "05"
	BTLFAudienceCode                           = "06"
	ElectreAudienceCode                        = "07"
	ANELETipo                                  = "08"
	AVI                                        = "09"
	USKRating                                  = "10"
	AWS                                        = "11"
	Schulform                                  = "12"
	Bundesland                                 = "13"
	Ausbildungsberuf                           = "14"
	SuomalainenKouluasteluokitus               = "15"
	CBGAgeGuidance                             = "16"
	NielsenBookAudienceCode                    = "17"
	AVIRevised                                 = "18"
	LexileMeasure                              = "19"
	FryReadabilityScore                        = "20"
	JapaneseChildrensAudienceCode              = "21"
	ONIXAdultAudienceRating                    = "22"
	CommonEuropeanFrameworkForLanguageLearning = "23"
	KoreanPublicationEthicsCommissionRating    = "24"
	IoEBookBand                                = "25"
	FSKLehrInfoprogramm                        = "26"
	IntendedAudienceLanguage                   = "27"
	PEGIRating                                 = "28"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		ONIXAudienceCodes:                          {"ONIX audience codes", "Using a code from List 28.", "ONIX målgruppekoder", "Bruker kodeliste 28"},
		Proprietary:                                {"Proprietary", "", "Proprietær", ""},
		MPAARating:                                 {"MPAA rating", "Motion Picture Association of America rating applied to movies.", "MPAA rating", "Motion Picture Association of America rating for filmer"},
		BBFCRating:                                 {"BBFC rating", "British Board of Film Classification rating applied to movies.", "BBFC rating", "British Board of Film Classification rating for filmer"},
		FSKRating:                                  {"FSK rating", "German FSK (Freiwillige Selbstkontrolle der Filmwirtschaft) rating applied to movies.", "FSK rating", "Tysk FSK (Freiwillige Selbstkontrolle der Filmwirtschaft) brukt for filmer"},
		BTLFAudienceCode:                           {"BTLF audience code", "French Canadian audience code list, used by BTLF for Memento.", "Fransk-kanadisk BTLF målgruppekoder", "Fransk-kanadiske målgruppekoder, brukt av BTLF for Memento"},
		ElectreAudienceCode:                        {"Electre audience code", "Audience code used by Electre (France).", "Fransk Electre målgruppekoder", "Franske målgruppekoder, brukt av Electre"},
		ANELETipo:                                  {"ANELE Tipo", "Spain: educational audience and material type code of the Asociación Nacional de Editores de Libros y Material de Enseñanza.", "Spansk ANELE Tipo", "Spania: koder for utdanningstrinn og materialtyper fra Asociación Nacional de Editores de Libros y Material de Enseñanza"},
		AVI:                                        {"AVI", "Code list used to specify reading levels for children’s books, used in Flanders, and formerly in the Netherlands – see also code 18.", "AVI", "Kodeliste brukt for å angi nivå på barnebøker, brukt i Flandern og Nederland"},
		USKRating:                                  {"USK rating", "German USK (Unterhaltungssoftware Selbstkontrolle) rating applied to video or computer games.", "USK rating", "German USK (Unterhaltungssoftware Selbstkontrolle). For video eller dataspill"},
		AWS:                                        {"AWS", "Audience code used in Flanders.", "AWS", "Audience code used in Flanders."},
		Schulform:                                  {"Schulform", "Type of school: codelist maintained by VdS Bildungsmedien eV, the German association of educational media publishers.", "Schulform", "Type of school: codelist maintained by VdS Bildungsmedien eV, the German association of educational media publishers."},
		Bundesland:                                 {"Bundesland", "School region: codelist maintained by VdS Bildungsmedien eV, the German association of educational media publishers, indicating where products are licensed to be used in schools.", "Bundesland", "School region: codelist maintained by VdS Bildungsmedien eV, the German association of educational media publishers, indicating where products are licensed to be used in schools."},
		Ausbildungsberuf:                           {"Ausbildungsberuf", "Occupation: codelist for vocational training materials, maintained by VdS Bildungsmedien eV, the German association of educational media publishers.", "Ausbildungsberuf", "Occupation: codelist for vocational training materials, maintained by VdS Bildungsmedien eV, the German association of educational media publishers."},
		SuomalainenKouluasteluokitus:               {"Suomalainen kouluasteluokitus", "Finnish school or college level.", "Suomalainen kouluasteluokitus", "Finnish school or college level."},
		CBGAgeGuidance:                             {"CBG age guidance", "UK Publishers Association, Children’s Book Group, coded indication of intended reader age, carried on book covers.", "CBG age guidance", "UK Publishers Association, Children’s Book Group, coded indication of intended reader age, carried on book covers."},
		NielsenBookAudienceCode:                    {"Nielsen Book audience code", "Audience code used in Nielsen Book Services.", "Nielsen Book audience code", "Audience code used in Nielsen Book Services."},
		AVIRevised:                                 {"AVI (revised)", "Code list used to specify reading levels for children’s books, used in the Netherlands – see also code 09.", "AVI (revised)", "Code list used to specify reading levels for children’s books, used in the Netherlands – see also code 09."},
		LexileMeasure:                              {"Lexile measure", "Lexile measure (the Lexile measure in <AudienceCodeValue> may optionally be prefixed by the Lexile code). Examples might be ‘880L’, ‘AD0L’ or ‘HL600L’. Deprecated – use <Complexity> instead.", "Lexile measure", "Lexile measure (the Lexile measure in <AudienceCodeValue> may optionally be prefixed by the Lexile code). Examples might be ‘880L’, ‘AD0L’ or ‘HL600L’. Deprecated – use <Complexity> instead."},
		FryReadabilityScore:                        {"Fry Readability score", "Fry readability metric based on number of sentences and syllables per 100 words. Expressed as a number from 1 to 15 in <AudienceCodeValue>. Deprecated – use <Complexity> instead.", "Fry Readability score", "Fry readability metric based on number of sentences and syllables per 100 words. Expressed as a number from 1 to 15 in <AudienceCodeValue>. Deprecated – use <Complexity> instead."},
		JapaneseChildrensAudienceCode:              {"Japanese Children’s audience code", "Children’s audience code (対象読者), two-digit encoding of intended target readership from 0–2 years up to High School level.", "Japanese Children’s audience code", "Children’s audience code (対象読者), two-digit encoding of intended target readership from 0–2 years up to High School level."},
		ONIXAdultAudienceRating:                    {"ONIX Adult audience rating", "Publisher’s rating indicating suitability for an particular adult audience, using a code from List 203.", "ONIX Adult audience rating", "Publisher’s rating indicating suitability for an particular adult audience, using a code from List 203."},
		CommonEuropeanFrameworkForLanguageLearning: {"Common European Framework for Language Learning", "Codes A1 to C2 indicating standardised level of language learning or teaching material, from beginner to advanced, used in EU.", "Common European Framework for Language Learning", "Codes A1 to C2 indicating standardised level of language learning or teaching material, from beginner to advanced, used in EU."},
		KoreanPublicationEthicsCommissionRating:    {"Korean Publication Ethics Commission rating", "Rating used in Korea to control selling of books and e-books to minors. Current values are 0 (suitable for all) and 19 (only for sale to ages 19+). See ‘http://www.kpec.or.kr/english/’.", "Korean Publication Ethics Commission rating", "Rating used in Korea to control selling of books and e-books to minors. Current values are 0 (suitable for all) and 19 (only for sale to ages 19+). See ‘http://www.kpec.or.kr/english/’."},
		IoEBookBand:                                {"IoE Book Band", "UK Institute of Education Book Bands for Guided Reading scheme (see http://www.ioe.ac.uk/research/4664.html). <AudienceCodeValue> is a colour, eg ‘Pink A’ or ‘Copper’. Deprecated – use <Complexity> instead.", "IoE Book Band", "UK Institute of Education Book Bands for Guided Reading scheme (see http://www.ioe.ac.uk/research/4664.html). <AudienceCodeValue> is a colour, eg ‘Pink A’ or ‘Copper’. Deprecated – use <Complexity> instead."},
		FSKLehrInfoprogramm:                        {"FSK Lehr-/Infoprogramm", "Used for German videos/DVDs with educational or informative content; value for <AudienceCodeValue> must be either ‘Infoprogramm gemäß § 14 JuSchG’ or ‘Lehrprogramm gemäß § 14 JuSchG’.", "FSK Lehr-/Infoprogramm", "Used for German videos/DVDs with educational or informative content; value for <AudienceCodeValue> must be either ‘Infoprogramm gemäß § 14 JuSchG’ or ‘Lehrprogramm gemäß § 14 JuSchG’."},
		IntendedAudienceLanguage:                   {"Intended audience language", "Where this is different from the language of the text of the book recorded in <Language>. <AudienceCodeValue> should be a value from list 74", "Intended audience language", "Where this is different from the language of the text of the book recorded in <Language>. <AudienceCodeValue> should be a value from list 74"},
		PEGIRating:                                 {"PEGI rating", "Pan European Game Information rating used primarily for video games", "PEGI rating", "Pan European Game Information rating used primarily for video games"},
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
