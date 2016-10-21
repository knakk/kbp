package list1 // Notification or update type code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Notification or update type code
const (
	EarlyNotification                         = "01"
	AdvanceNotificationConfirmed              = "02"
	NotificationConfirmedOnPublication        = "03"
	UpdatePartial                             = "04"
	Delete                                    = "05"
	NoticeOfSale                              = "08"
	NoticeOfAcquisition                       = "09"
	UpdateSupplyDetailOnly                    = "12"
	UpdateMarketRepresentationOnly            = "13"
	UpdateSupplyDetailAndMarketRepresentation = "14"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "08", "09", "12", "13", "14"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		EarlyNotification:                         {"Early notification", "Use for a complete record issued earlier than approximately six months before publication.", "Tidlig melding om utgivelse", "Bruk for en komplett ONIX-post sendt tidligere enn ca. seks måneder før utgivelse"},
		AdvanceNotificationConfirmed:              {"Advance notification (confirmed)", "Use for a complete record issued to confirm advance information approximately six months before publication; or for a complete record issued after that date and before information has been confirmed from the book-in-hand.", "Forhåndsmelding om utgivelse (bekreftet)", "Bruk for en komplett ONIX-post sendt for å bekrefte informasjonen ca. seks måneder før utgivelse"},
		NotificationConfirmedOnPublication:        {"Notification confirmed on publication", "Use for a complete record issued to confirm advance information at or just before actual publication date; or for a complete record issued at any later date.", "Melding med bekreftet lagerstatus", "Brukes for en komplett ONIX-post som er sendt på grunnlag av bok-i-handa (den fysiske boka) på eller like før utgivelsesdato"},
		UpdatePartial:                             {"Update (partial)", "In ONIX 3.0 only, use when sending a ‘block update’ record. In previous ONIX releases, ONIX updating has generally been by complete record replacement using code 03, and code 04 is not used.", "Oppdatering (blocks)", "Kun i Onix 3.0, brukes når man sender en \"block update\". I tidligere versjoner av Onix har oppdatering av Onix stort sett blitt gjort ved å sende komplette poster og ved å bruke kode 03, kode 04 blir ikke brukt"},
		Delete:                                    {"Delete", "Use when sending an instruction to delete a record which was previously issued. Note that a Delete instruction should NOT be used when a product is cancelled, put out of print, or otherwise withdrawn from sale: this should be handled as a change of Publishing status, leaving the receiver to decide whether to retain or delete the record. A Delete instruction is only used when there is a particular reason to withdraw a record completely, eg because it was issued in error.", "Sletting", "Brukes når man sender beskjed om å slette en post som tidligere er oversendt. Merk at kode for sletting IKKE skal brukes når et produkt ikke vil utkomme, blir utsolgt eller av andre grunnes ikke er i salg lenger: disse tilfellene skal behandles som endringer av Publishing status, og det er opp til mottakeren om posten skal beholdes eller slettes. Beskjed om sletting brukes bare når det en angitt grunn til at posten skal slettes. F.eks. at informasjon om produktet ble sendt ved en feil."},
		NoticeOfSale:                              {"Notice of sale", "Notice of sale of a product, from one publisher to another: sent by the publisher disposing of the product.", "Melding om salg av produkt", "Melding om salg/overdragelse av et produkt fra en vareeier til en annen. Sendes fra den vareeier som selger"},
		NoticeOfAcquisition:                       {"Notice of acquisition", "Notice of acquisition of a product, by one publisher from another: sent by the acquiring publisher.", "Melding om kjøp av produkt", "Melding om salg/overdragelse av en tittel fra en vareeier til en annen. Sendes fra den vareeier som kjøper"},
		UpdateSupplyDetailOnly:                    {"Update – SupplyDetail only", "ONIX Books 2.1 supply update – <SupplyDetail> only (not used in ONIX 3.0).", "Oppdatering - kun SupplyDetail", "ONIX Books 2.1 supply update - kun SupplyDetail (ikke brukt i Onix 3.0)"},
		UpdateMarketRepresentationOnly:            {"Update – MarketRepresentation only", "ONIX Books 2.1 supply update – <MarketRepresentation> only (not used in ONIX 3.0).", "Oppdatering - kun MarketRepresentation", "ONIX Books 2.1 supply update - kun MarketRepresentation (ikke brukt i Onix 3.0)"},
		UpdateSupplyDetailAndMarketRepresentation: {"Update – SupplyDetail and MarketRepresentation", "ONIX Books 2.1 supply update – both <SupplyDetail> and <MarketRepresentation> (not used in ONIX 3.0).", "Oppdatering - både SupplyDetail og MarketRepresentation", "ONIX Books 2.1 supply update - både SupplyDetail og MarketRepresentation (ikke brukt i Onix 3.0)"},
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
