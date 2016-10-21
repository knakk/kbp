package list30

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	USSchoolGradeRange                = "11"
	UKSchoolGrade                     = "12"
	ReadingSpeedWordsPerMinute        = "15"
	InterestAgeMonths                 = "16"
	InterestAgeYears                  = "17"
	ReadingAgeYears                   = "18"
	SpanishSchoolGrade                = "19"
	Skoletrinn                        = "20"
	Nivå                              = "21"
	ItalianSchoolGrade                = "22"
	Schulform                         = "23"
	Bundesland                        = "24"
	Ausbildungsberuf                  = "25"
	CanadianSchoolGradeRange          = "26"
	FinnishSchoolGradeRange           = "27"
	FinnishUpperSecondarySchoolCourse = "28"
	ChineseSchoolGradeCode            = "29"
	NomenclatureNiveaux               = "30"
)

var (
	sortedCodes = []string{"11", "12", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		USSchoolGradeRange:                {"US school grade range", "Values for <AudienceRangeValue> are specified in List 77.", "Amerikanske skoletrinn", "Verdier for <AudienceRangeValue>  er spesifisert i liste 77"},
		UKSchoolGrade:                     {"UK school grade", "Values are defined by BIC for England and Wales, Scotland and N Ireland.", "Skoletrinn i Storbritannia", "Verdier definert av BIC for England, Wales, Skottland og nord-Irland"},
		ReadingSpeedWordsPerMinute:        {"Reading speed, words per minute", "Values in <AudienceRangeValue> must be integers.", "Lesehastighet - ord per minutt", "Verdier i &lt;AudienceRangeValue&gt;; må være heltall"},
		InterestAgeMonths:                 {"Interest age, months", "For use up to 30 months only: values in <AudienceRangeValue> must be integers.", "Målgruppe, angitt i måneder", "Brukes kun opp til 30 måneder: verdier i <AudienceRangeValue> må være heltall"},
		InterestAgeYears:                  {"Interest age, years", "Values in <AudienceRangeValue> must be integers.", "Interessealder, år", "Verdier i  <AudienceRangeValue> må være heltall"},
		ReadingAgeYears:                   {"Reading age, years", "Values in <AudienceRangeValue> must be integers.", "Lesealder, år", "Verdier i  <AudienceRangeValue> må være heltall"},
		SpanishSchoolGrade:                {"Spanish school grade", "Spain: combined grade and region code, maintained by the Ministerio de Educación.", "Skoletrinn i Spania", "Spania: Kombinerte trinn- og regionskoder. Vedlikeholdes av Ministerio de Educación"},
		Skoletrinn:                        {"Skoletrinn", "Norwegian educational level for primary and secondary education", "Skoletrinn i Norge", "Norske skoletrinn"},
		Nivå:                              {"Nivå", "Swedish educational qualifier (code).", "Skoletrinn i Sverige", "Svensk kodeliste for skoletrinn"},
		ItalianSchoolGrade:                {"Italian school grade", "", "Skoletrinn i Italia", "Refererer til eksterne kodelister"},
		Schulform:                         {"Schulform", "DEPRECATED – assigned in error: see List 29.", "Schulform", "UTGÅTT. Se liste 29"},
		Bundesland:                        {"Bundesland", "DEPRECATED – assigned in error: see List 29.", "Bundesland", "UTGÅTT. Se liste 29"},
		Ausbildungsberuf:                  {"Ausbildungsberuf", "DEPRECATED – assigned in error: see List 29.", "Ausbildungsberuf", "UTGÅTT. Se liste 29"},
		CanadianSchoolGradeRange:          {"Canadian school grade range", "Values for <AudienceRangeValue> are specified in List 77.", "Canadian school grade range", "Values for <AudienceRangeValue> are specified in List 77."},
		FinnishSchoolGradeRange:           {"Finnish school grade range", "", "Finnish school grade range", ""},
		FinnishUpperSecondarySchoolCourse: {"Finnish Upper secondary school course", "Lukion kurssi.", "Finnish Upper secondary school course", "Lukion kurssi."},
		ChineseSchoolGradeCode:            {"Chinese School Grade code", "Values are P, K, 1–17 (including college-level audiences).", "Chinese School Grade code", "Values are P, K, 1–17 (including college-level audiences)."},
		NomenclatureNiveaux:               {"Nomenclature niveaux", "French educational level classification, used for example on WizWiz.fr. See http://www.kiosque-edu.com/html/onix/Nomenclature_niveaux.csv’.", "Nomenclature niveaux", "French educational level classification, used for example on WizWiz.fr. See http://www.kiosque-edu.com/html/onix/Nomenclature_niveaux.csv’."},
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
