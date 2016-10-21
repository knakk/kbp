package list2 // Product composition

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Product composition
const (
	SingleItemRetailProduct                       = "00"
	MultipleItemRetailProduct                     = "10"
	MultipleItemCollectionRetailedAsSeparateParts = "11"
	TradeOnlyProduct                              = "20"
	MultipleItemTradePack                         = "30"
	MultipleItemPack                              = "31"
)

var (
	sortedCodes = []string{"00", "10", "11", "20", "30", "31"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		SingleItemRetailProduct:                       {"Single-item retail product", "", "Enkeltprodukt", ""},
		MultipleItemRetailProduct:                     {"Multiple-item retail product", "Multiple-item product retailed as a whole.", "Produkt bestående av flere enkeltprodukter", "Produktet selges som en enhet"},
		MultipleItemCollectionRetailedAsSeparateParts: {"Multiple-item collection, retailed as separate parts", "Used when an ONIX record is required for a collection-as-a-whole, even though it is not currently retailed as such.", "Produkt bestående av flere enkeltprodukter, solgt separat", "Brukes når en Onix-post kreves for produktet som helhet, selv om det ikke p.t. selges som en enhet, men som enkeltprodukter"},
		TradeOnlyProduct:                              {"Trade-only product", "Product not for retail, and not carrying retail items, eg empty dumpbin, empty counterpack, promotional material.", "Forhandlerprodukt", "Produktet er ikke ment for vanlig detaljhandel. F.eks. tomme gulv- eller diskdisplay, reklamemateriell"},
		MultipleItemTradePack:                         {"Multiple-item trade pack", "Carrying multiple copies for retailing as separate items, eg shrink-wrapped trade pack, filled dumpbin, filled counterpack.", "Forhandlerprodukt bestående av flere enkeltprodukter, selges separat", "Produktet inneholder flere eksemplarer av samme produkt, ment for salg som enkeltprodukter. F.eks. pakke med flere eksemplarer av samme bok, fylte gulv- eller diskdisplay"},
		MultipleItemPack:                              {"Multiple-item pack", "Carrying multiple copies, primarily for retailing as separate items. The pack may be split and retailed as separate items OR retailed as a single item. Use instead of Multiple item trade pack (code 30) only if the data provider specifically wishes to make explicit that the pack may optionally be retailed as a whole.", "Forhandlerprodukt bestående av flere enkeltprodukter", "Produktet inneholder flere eksemplarer av samme produkt, primært ment for salg som enkeltprodukter. Pakka kan splittes og selges som enkelprodukter ELLER hele pakka kan selges som et enkeltprodukt. Brukes istedenfor kode 30 bare dersom man spesifikt ønsker å uttrykke at det er mulig å selge pakka som en helhet."},
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
