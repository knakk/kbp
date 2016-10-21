package list165 // Supplier own code type
import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Supplier own code type
const (
	SuppliersSalesClassification              = "01"
	SuppliersBonusEligibility                 = "02"
	PublishersSalesClassification             = "03"
	SuppliersPricingRestrictionClassification = "04"
	SuppliersSalesExpecation                  = "05"
	PublishersSalesExpectation                = "06"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		SuppliersSalesClassification:              {"Supplier’s sales classification", "A rating applied by a supplier (typically a wholesaler) to indicate its assessment of the expected or actual sales performance of a product.", "Distributørens salgsklassifikasjon", "Rating angitt av distributør eller grossist for å gi indikasjon om forventet salg av et produkt"},
		SuppliersBonusEligibility:                 {"Supplier’s bonus eligibility", "A supplier’s coding of the eligibility of a product for a bonus scheme on overall sales.", "Supplier’s bonus eligibility", "A supplier’s coding of the eligibility of a product for a bonus scheme on overall sales."},
		PublishersSalesClassification:             {"Publisher’s sales classification", "A rating applied by the publisher to indicate a sales category (eg backlist/frontlist, core stock etc). Use only when the publisher is not the ‘supplier’.", "Utgivers salgsklassifikasjon", "Rating angitt av utgiver for å gi indikasjon om produktets salgskategori (f.eks. backlist/frontlist, basisutvalg osv). Brukes kun når utgiver ikke er den samme som 'distributør'"},
		SuppliersPricingRestrictionClassification: {"Supplier’s pricing restriction classification", "A classification applied by a supplier to a product sold on Agency terms, to indicate that retail price restrictions are applicable.", "Distributørens klassifikasjon for prisrestriksjoner", "A classification applied by a supplier to a product sold on Agency terms, to indicate that retail price restrictions are applicable."},
		SuppliersSalesExpecation:                  {"Supplier's sales expecation", "Code is the ISBN of another book that had sales (both in terms of copy numbers and customer profile) comparable to that the distributor or supplier estimates for the product. <SupplierCodeValue> must be an ISBN-13 or GTIN-13", "Distributørens salgsforventning", "Koden er ISBN-et til en annen bok som har hatt salg som kan sammenlignes med distributørens anslag for produktet. <SupplierCodeValue> må være ISBN-13 eller GTIN-13."},
		PublishersSalesExpectation:                {"Publisher's sales expectation", "Code is the ISBN of another book that had sales (both in terms of copy numbers and customer profile) comparable to that the publisher estimates for the product. <SupplierCodeValue> must be an ISBN-13 or GTIN-13", "Utgivers salgsforventning", "Koden er ISBN-et til en annen bok som har hatt salg som kan sammenlignes med utgiverens anslag for produktet. <SupplierCodeValue> må være ISBN-13 eller GTIN-13."},
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
