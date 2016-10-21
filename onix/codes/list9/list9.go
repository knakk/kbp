package list9 // Product classification type code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Product classification type code
const (
	WCOHarmonizedSystem                         = "01"
	UNSPSC                                      = "02"
	HMRC                                        = "03"
	WarenverzeichnisFürDieAußenhandelsstatistik = "04"
	TARIC                                       = "05"
	Fondsgroep                                  = "06"
	SendersProductCategory                      = "07"
	GAPPProductClass                            = "08"
	CPA                                         = "09"
	NCM                                         = "10"
	ElectreGenre                                = "50"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "50"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		WCOHarmonizedSystem: {"WCO Harmonized System", "World Customs Organization Harmonized Commodity Coding and Description System.", "WCO harmonisert system", "World Customs Organization Harmonized Commodity Coding & Description System"},
		UNSPSC:              {"UNSPSC", "UN Standard Product and Service Classification.", "UNSPSC", "UN Standard Product & Service Classification"},
		HMRC:                {"HMRC", "UK Revenue and Customs classifications, based on the Harmonized System.", "HMCE", "UK Customs & Excise classifications, based on the Harmonized System"},
		WarenverzeichnisFürDieAußenhandelsstatistik: {"Warenverzeichnis für die Außenhandelsstatistik", "German export trade classification, based on the Harmonised System.", "Tysk varefortegnelse for eksportstatistikk", "German export trade classification, based on the Harmonised System"},
		TARIC:                  {"TARIC", "EU TARIC codes, an extended version of the Harmonized System.", "TARIC", "EU TARIC codes, an extended version of the Harmonized System"},
		Fondsgroep:             {"Fondsgroep", "Centraal Boekhuis free classification field for publishers.", "Fondsgroep", "Centraal Boekhuis free classification field for publishers."},
		SendersProductCategory: {"Sender’s product category", "A product category (not a subject classification) assigned by the sender.", "Sender’s product category", "A product category (not a subject classification) assigned by the sender."},
		GAPPProductClass:       {"GAPP Product Class", "Product classification maintained by the Chinese General Administration of Press and Publication (http://www.gapp.gov.cn).", "GAPP Product Class", "Product classification maintained by the Chinese General Administration of Press and Publication (http://www.gapp.gov.cn)."},
		CPA:                    {"CPA", "Statistical Classification of Products by Activity in the European Economic Community, see http://ec.europa.eu/eurostat/ramon/nomenclatures/index.cfm?TargetUrl=LST_NOM_DTL&StrNom=CPA_2008. Up to six digits, with one or two periods. For example, printed children’s books are “58.11.13”, but the periods are normally ommited in ONIX.", "CPA", "Statistical Classification of Products by Activity in the European Economic Community, see http://ec.europa.eu/eurostat/ramon/nomenclatures/index.cfm?TargetUrl=LST_NOM_DTL&StrNom=CPA_2008. Up to six digits, with one or two periods. For example, printed children’s books are “58.11.13”, but the periods are normally ommited in ONIX."},
		NCM:                    {"NCM", "Mercosur/Mercosul Common Nomenclature, based on the Harmonised System.", "NCM", "Mercosur/Mercosul Common Nomenclature, based on the Harmonised System."},
		ElectreGenre:           {"Electre genre", "Typologie de marché géré par Electre (Market segment code maintained by Electre).", "Electre genre", "Typologie de marché géré par Electre (Market segment code maintained by Electre)."},
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
