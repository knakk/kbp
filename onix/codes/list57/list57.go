package list57

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	FreeOfCharge       = "01"
	PriceToBeAnnounced = "02"
	NotSoldSeparately  = "03"
	ContactSupplier    = "04"
	NotSoldAsSet       = "05"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		FreeOfCharge:       {"Free of charge", "", "Gratis", ""},
		PriceToBeAnnounced: {"Price to be announced", "", "Pris ikke fastsatt", ""},
		NotSoldSeparately:  {"Not sold separately", "", "Ikke solgt separat", ""},
		ContactSupplier:    {"Contact supplier", "May be used for books that do not carry a recommended retail price, when an ONIX file is “broadcast” rather than sent one-to-one to a single trading partner; or for digital products offered on subscription or with pricing which is too complex to specify in ONIX.", "Kontakt forlag eller distributør", "Kan brukes for produkter som ikke har en anbefalt utsalgspris. F.eks. når en ONIX-post er ment for alle framfor å være sendt som en en-til-en melding til en gitt handelspartner; eller for digitale produkter som selges i abonnement, eller som har prising som er for kompleks til å angi i Onix."},
		NotSoldAsSet:       {"Not sold as set", "When a collection that is not sold as a set nevertheless has its own ONIX record.", "Selges ikke samlet", "Når en serie (collection) ikke selges som et sett, men likevel har sin egen Onix-post"},
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
