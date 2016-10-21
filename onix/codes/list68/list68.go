package list68

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Unspecified                  = "00"
	Cancelled                    = "01"
	Forthcoming                  = "02"
	PostponedIndefinitely        = "03"
	Active                       = "04"
	NoLongerOurProduct           = "05"
	OutOfStockIndefinitely       = "06"
	OutOfPrint                   = "07"
	Inactive                     = "08"
	Unknown                      = "09"
	Remaindered                  = "10"
	WithdrawnFromSale            = "11"
	NotAvailableInThisMarket     = "12"
	ActiveButNotSoldSeparately   = "13"
	ActiveWithMarketRestrictions = "14"
	Recalled                     = "15"
	TemporarilyWithdrawnFromSale = "16"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Unspecified:                  {"Unspecified", "Status is not specified (as distinct from unknown): the default if the <MarketPublishingStatus> element is not sent.", "Ikke spesifisert", "Status er ikke spesifisert (til forskjell fra ukjent). Dette brukes dersom <MarketPublishingStatus> elementet ikke er sendt fra vareeier."},
		Cancelled:                    {"Cancelled", "The product was announced for publication in this market, and subsequently abandoned.", "Vil ikke utkomme", "Kanselert: produktet ble annonsert, men har siden blitt trukket"},
		Forthcoming:                  {"Forthcoming", "Not yet published in this market, should be accompanied by expected local publication date..", "Forventes utgitt", "Ikke utgitt ennå. Må etterfølges av en forventet lokal utgivelses dato"},
		PostponedIndefinitely:        {"Postponed indefinitely", "The product was announced for publication in this market, and subsequently postponed with no expected local publication date.", "Utsatt inntil videre", "Produktet ble annonsert, men har siden blitt utsatt uten at forventet dato er angitt."},
		Active:                       {"Active", "The product was published in this market, and is still active in the sense that the publisher will accept orders for it, though it may or may not be immediately available, for which see <SupplyDetail>.", "I salg", "Produktet ble annonsert og er fortsatt aktivt i den betydning at forlaget forsatt vil akseptere ordre, men det er ikke sikkert når det blir tilgjengelig. Se også <SupplyDetail>"},
		NoLongerOurProduct:           {"No longer our product", "Responsibility for the product in this market has been transferred elsewhere.", "Ikke lenger vårt produkt", "Eierskap til produktet har blitt overført til noen andre"},
		OutOfStockIndefinitely:       {"Out of stock indefinitely", "The product was active, but is now inactive in the sense that (a) no further stock is expected to be made available in this market, though stock may still be available elsewhere in the supply chain, and (b) there are no current plans to bring it back into stock.", "Utsolgt, til vurdering", "Ikke tilgjengelig i dette markedet, men det kan finnes andre steder. Det er ikke besluttet om produktet vil bli tatt inn på lager igjen eller ikke"},
		OutOfPrint:                   {"Out of print", "The product was active, but is now permanently inactive in the sense that (a) no further stock is expected to be made available in this market, though stock may still be available elsewhere in the supply chain, and (b) the product will not be made available again under the same ISBN.", "Utsolgt", "Ikke forventet å bli tilgjengelig igjen i dette markedet, men det kan finnes andre steder. Vil ikke bli tatt inn på lager igjen."},
		Inactive:                     {"Inactive", "The product was active, but is now permanently or indefinitely inactive in the sense that no further stock is expected to be made available in this market, though stock may still be available elsewhere in the supply chain. Code 08 covers both of codes 06 and 07, and may be used where the distinction between those values is either unnecessary or meaningless.", "Ikke i ordinært salg", "Produktet var aktivt, men er nå tatt ut av ordinært salg i den betydning at produktet er ikke forventet å få lagerbeholdning i dette markedet, selv om det kan være tilgjengelige eksemplarer andre steder i verdikjeden. Kode 08 dekker både kode 06 og 07, og kan brukes der skillet mellom de to kodene ikke er nødvendig, eller ikke gir mening."},
		Unknown:                      {"Unknown", "The sender of the ONIX record does not know the current publishing status in this market.", "Ukjent", "Avsenderen av ONIX posten vet ikke utgivelsesstatus i dette markedet"},
		Remaindered:                  {"Remaindered", "The product is no longer available in this market from the local publisher, under the current ISBN, at the current price. It may be available to be traded through another channel, usually at a reduced price.", "Reservert nedsettelse", "Produktet er ikke lenger tilgjengelig i dette markedet fra lokal utgiver under dette ISBN til angitt pris. Det kan være tilgjengelig i andre kanaler. Vanligvis til en redusert pris."},
		WithdrawnFromSale:            {"Withdrawn from sale", "Withdrawn from sale in this market, typically for legal reasons.", "Trukket tilbake fra salg", "Trukket tilbake fra dette markedet. Ofte av juridiske årsaker"},
		NotAvailableInThisMarket:     {"Not available in this market", "Either no rights are held for the product in this market, or for other reasons the publisher has decided not to make it available in this market.", "Ikke tilgjengelig i dette markedet", "Enten fordi ingen har rettighetene i dett markedet eller fordi utgiver av andre årsaker har bestemt å ikke gjøre det tilgjengelig."},
		ActiveButNotSoldSeparately:   {"Active, but not sold separately", "The product is published in this market and active but, as a publishing decision, it is not sold separately – only in an assembly or as part of a package.", "Selges ikke enkeltvis", "Kan kun kjøpes som del av et verk eller pakke"},
		ActiveWithMarketRestrictions: {"Active, with market restrictions", "The product is published in this market and active, but is not available to all customer types, typically because the market is split between exclusive sales agents for different market segments. In ONIX 2.1, should be accompanied by a free-text statement in <MarketRestrictionDetail> describing the nature of the restriction. In ONIX 3.0, the <SalesRestriction> composite in Group P.24 should be used.", "I salg, med restiksjoner", "Produktet er utgitt, men ikke tilgjengelig for alle kundetyper. Typisk fordi markedet er delt mellom salgsagenter for ulike segmenter. Skal følges av tekst i <MarketRestrictionDetail> som beskriver disse restriksjonene"},
		Recalled:                     {"Recalled", "Recalled in this market for reasons of consumer safety.", "Tilbakekalt", "Recalled in this market for reasons of consumer safety."},
		TemporarilyWithdrawnFromSale: {"Temporarily withdrawn from sale", "Temporarily withdrawn from sale in this market, typically for quality or technical reasons. In ONIX 3.0, must be accompanied by expected availability date coded ‘22’ within the <MarketPublishingDate> composite, except in exceptional circumstances where no date is known.", "Midlertidig trukket tilbake fra salg", "Midlertidig trukket tilbake, som oftest pga. forhold rundt kvalitet eller teknisk årsaker. I Onix 3.0 må koden følges av en dato med kode ’22’ i <MarketPublishingDate>, bortsett fra i eksepsjonelle tilfeller, der dato ikke er kjent."},
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
