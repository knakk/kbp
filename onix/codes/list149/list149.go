package list149

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Product         = "01"
	CollectionLevel = "02"
	Subcollection   = "03"
	ContentItem     = "04"
	MasterBrand     = "05"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Product:         {"Product", "The title element refers to an individual product.", "Produkt", "Tittelelementet viser til et enkeltprodukt"},
		CollectionLevel: {"Collection level", "The title element refers to the top level of a bibliographic collection.", "Samlingsnivå (øverste nivå)", "Tittelelementet viser til øverste nivå av ei bibliografisk samling (for eksempel serietittel)"},
		Subcollection:   {"Subcollection", "The title element refers to an intermediate level of a bibliographic collection that comprises two or more ‘sub-collections’.", "Samlingsnivå (underordnet nivå)", "Tittelelementet viser til et underordnet nivå i ei bibliografisk samling som består av to eller flere ‘sub-collections’."},
		ContentItem:     {"Content item", "The title element refers to a content item within a product, eg a work included in a combined or ‘omnibus’ edition, or a chapter in a book.", "Innholdselement", "Tittelelementet viser til et innholdselement i et produkt, f.eks. et verk som er inkludert i ei sammensatt utgave (omnibus edition), eller et kapittel i ei bok"},
		MasterBrand:     {"Master brand", "The title element names a master brand where the use of the brand spans multiple collections and product forms, and possibly multiple imprints and publishers. Used only for branded media properties such as children’s character properties.", "Merkevarenavn", "Tittelelementet viser til et merkevarenavn hvor bruken av dette navnet strekker seg over flere samlinger og produktformater, og mulighens ulike forlag og utgivere. Brukes kun for produkter som er merket med merkevarenavnet, slik som hovedpersoner i barneprodukter"},
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
