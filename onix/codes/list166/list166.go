package list166

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	EmbargoDate              = "02"
	ExpectedAvailabilityDate = "08"
	LastDateForReturns       = "18"
	ReservationOrderDeadline = "25"
)

var (
	sortedCodes = []string{"02", "08", "18", "25"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		EmbargoDate:              {"Embargo date", "If there is an embargo on retail sales before a certain date, the date from which the embargo is lifted and retail sales are permitted.", "Frigivelsesdato", "Dersom det er restriksjoner på salg før en gitt dato, frigivelsesdato for når restriksjonene opphører og salg er tillatt"},
		ExpectedAvailabilityDate: {"Expected availability date", "The date on which physical stock is expected to be available to be shipped to retailers, or a digital product is expected to be released by the publisher or digital asset distributor to retailers or their retail platform providers.", "Forventet tilgjengelighetsdato", "Dato for når fysisk lagerbeholdning er klar for å sendes ut til forhandler, eller for når et digitalt produkt er forventet å være tilgjengelig for forhandler"},
		LastDateForReturns:       {"Last date for returns", "Last date when returns will be accepted, generally for a product which is being remaindered or put out of print.", "Siste dato for retur", "Siste dato for når retur aksepteres, vanligvis for produkter som har et restopplag eller som kommer til å bli utsolgt"},
		ReservationOrderDeadline: {"Reservation order deadline", "Latest date on which an order may be placed for guaranteed delivery prior to the publication date. May or may not be linked to a special reservation or pre-publication price.", "Tidsfrist for reservasjonsordre", "Siste dato for å legg ordre som garanterer levering før utgivelsesdato. Kan ha, men har ikke nødvendigvis, en spesiell rabattert pris"},
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
