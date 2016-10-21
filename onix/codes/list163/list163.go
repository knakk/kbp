package list163 // Publishing date role

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Publishing date role
const (
	PublicationDate                                  = "01"
	EmbargoDate                                      = "02"
	PublicAnnouncementDate                           = "09"
	TradeAnnouncementDate                            = "10"
	DateOfFirstPublication                           = "11"
	LastReprintDate                                  = "12"
	OutOfPrintDeletionDate                           = "13"
	LastReissueDate                                  = "16"
	PublicationDateOfPrintCounterpart                = "19"
	DateOfFirstPublicationInOriginalLanguage         = "20"
	ForthcomingReissueDate                           = "21"
	ExpectedAvailabilityDateAfterTemporaryWithdrawal = "22"
	ReviewEmbargoDate                                = "23"
	PublishersReservationOrderDeadline               = "25"
	ForthcomingReprintDate                           = "26"
)

var (
	sortedCodes = []string{"01", "02", "09", "10", "11", "12", "13", "16", "19", "20", "21", "22", "23", "25", "26"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		PublicationDate:                                  {"Publication date", "Nominal date of publication.", "Utgivelsesdato", "Nominal date of publication."},
		EmbargoDate:                                      {"Embargo date", "If there is an embargo on retail sales in this market before a certain date, the date from which the embargo is lifted and retail sales are permitted.", "Frigivelsesdato", "For produkter med handelsrestriksjoner i dette markedet før en angitt dato; datoen for når handelsrestriksjonen oppheves og det er tillatt å forhandle produktet."},
		PublicAnnouncementDate:                           {"Public announcement date", "Date when a new product may be announced to the general public.", "Offentliggjøringsdato", "Dato for når informasjon om et produkt kan offentliggjøres for det allmenne publikum"},
		TradeAnnouncementDate:                            {"Trade announcement date", "Date when a new product may be announced for trade only.", "Offentliggjøringsdato for bransjen", "Dato for når informasjon om et produkt kan offentliggjøres kun for bransjen"},
		DateOfFirstPublication:                           {"Date of first publication", "Date when the work incorporated in a product was first published.", "Første utgivelsesdato", "Dato for når verket først ble utgitt"},
		LastReprintDate:                                  {"Last reprint date", "Date when a product was last reprinted.", "Siste opplagsdato", "Dato for når siste opplag av produktet ble trykket"},
		OutOfPrintDeletionDate:                           {"Out-of-print / deletion date", "Date when a product was (or will be) declared out-of-print or deleted.", "Utsolgt-/slettedato", "Dato for når produktet ble (eller kommer til å bli) utsolgt eller slettet"},
		LastReissueDate:                                  {"Last reissue date", "Date when a product was last reissued.", "Siste nyutgivelsesdato", "Dato for når et produkt sist ble utgitt på nytt"},
		PublicationDateOfPrintCounterpart:                {"Publication date of print counterpart", "Date of publication of a printed book which is the print counterpart to a digital edition.", "Utgivelsesdato for trykt versjon", "Utgivelsesdato for ei trykt bok som er det trykte motstykket til ei digital utgave"},
		DateOfFirstPublicationInOriginalLanguage:         {"Date of first publication in original language", "Year when the original language version of work incorporated in a product was first published (note, use only when different from code 11).", "Første utgivelsesdato på originalspråket", "Første utgivelsesår for verket på originalspråket (merk: brukes kun når forskjellig fra kode 11)."},
		ForthcomingReissueDate:                           {"Forthcoming reissue date", "Date when a product will be reissued.", "Forventet dato for nyutgivelse", "Dato for når et produkt vil bli utgitt på nytt"},
		ExpectedAvailabilityDateAfterTemporaryWithdrawal: {"Expected availability date after temporary withdrawal", "Date when a product that has been temporary withdrawn from sale or recalled for any reason is expected to become available again, eg after correction of quality or technical issues.", "Forventet tilgjengelighetsdato etter midlertidig ikke tilgjengelig", "Dato for når et produkt som har vært midlertidig ikke tilgjengelig er forventet å bli tilgjengelig igjen, f.eks. etter korreksjon av ting som har med kvalitet eller det tekniske å gjøre"},
		ReviewEmbargoDate:                                {"Review embargo date", "Date from which reviews of a product may be published eg in newspapers and magazines or online. Provided to the book trade for information only: newspapers and magazines are not expected to be recipients of ONIX metadata.", "Dato for anmeldelser", "Dato for når anmeldelser av et produkt kan publiseres, f.eks. i aviser og tidsskrifter og på nett. Kun for informasjon: aviser og tidsskrifter forventes ikke å motta Onix."},
		PublishersReservationOrderDeadline:               {"Publisher’s reservation order deadline", "Latest date on which an order may be placed with the publisher for guaranteed delivery prior to the publication date. May or may not be linked to a special reservation or pre-publication price.", "Utgivers frist for å legge forhåndsordre", "Siste dato for når ordre kan sendes til utgiver for at man skal være garantert leveranse før utgivelsesdatoen. Kan, men er ikke nødvendigvis, knyttet til en spesiell pris for produktet."},
		ForthcomingReprintDate:                           {"Forthcoming reprint date", "Date when a product will be reprinted.", "Forventet dato for nytt opplag", "Dato for når et nytt opplag til utgis"},
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
