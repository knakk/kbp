package list65

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Cancelled                                             = "01"
	NotYetAvailable                                       = "10"
	AwaitingStock                                         = "11"
	NotYetAvailableWillBePOD                              = "12"
	Available                                             = "20"
	InStock                                               = "21"
	ToOrder                                               = "22"
	POD                                                   = "23"
	TemporarilyUnavailable                                = "30"
	OutOfStock                                            = "31"
	Reprinting                                            = "32"
	AwaitingReissue                                       = "33"
	TemporarilyWithdrawnFromSale                          = "34"
	NotAvailableReasonUnspecified                         = "40"
	NotAvailableReplacedByNewProduct                      = "41"
	NotAvailableOtherFormatAvailable                      = "42"
	NoLongerSuppliedByUs                                  = "43"
	ApplyDirect                                           = "44"
	NotSoldSeparately                                     = "45"
	WithdrawnFromSale                                     = "46"
	Remaindered                                           = "47"
	NotAvailableReplacedByPOD                             = "48"
	Recalled                                              = "49"
	NotSoldAsSet                                          = "50"
	NotAvailablePublisherIndicatesOP                      = "51"
	NotAvailablePublisherNoLongerSellsProductInThisMarket = "52"
	NoRecentUpdateReceived                                = "97"
	NoLongerReceivingUpdates                              = "98"
	ContactSupplier                                       = "99"
)

var (
	sortedCodes = []string{"01", "10", "11", "12", "20", "21", "22", "23", "30", "31", "32", "33", "34", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "97", "98", "99"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Cancelled:                {"Cancelled", "Cancelled: product was announced, and subsequently abandoned.", "Vil ikke utkomme", "Produktet ble annonsert, men har siden blitt trukket."},
		NotYetAvailable:          {"Not yet available", "Not yet available (requires <ExpectedShipDate>, except in exceptional circumstances where no date is known).", "Ikke tilgjengelig ennå", "Ikke tilgjengelig ennå. Krever <ExpectedShipDate>, unntatt i spesielle tilfeller hvor datoen er ukjent"},
		AwaitingStock:            {"Awaiting stock", "Not yet available, but will be a stock item when available (requires <ExpectedShipDate>, except in exceptional circumstances where no date is known). Used particularly for imports which have been published in the country of origin but have not yet arrived in the importing country.", "Ikke ankommet lager", "Ikke tilgjengelig ennå, men lagerføres ved tilgjengelighet. Krever <ExpectedShipDate>, unntatt i spesielle tilfeller hvor datoen er ukjent. Brukes særlig for import når produktet er publisert i opprinnelseslandet, men har ennå ikke ankommet til det importerende landet."},
		NotYetAvailableWillBePOD: {"Not yet available, will be POD", "Not yet available, to be published as print-on-demand only. May apply either to a POD successor to an existing conventional edition, when the successor will be published under a different ISBN (normally because different trade terms apply); or to a title that is being published as a POD original.", "Ikke tilgjengelig ennå, vil komme som POD", "Ikke tilgjengelig ennå. Vil kun bli tilgjengelig som Print-On-Demand. Kan brukes både om en POD-utgivelse som er en videreføring av en trykket utgave, når denne utgis under et annet ISBN (normalt fordi det er ulike handelsbetingelser). Eller det kan brukes om en tittel som blir utgitt som POD i orginalutgave."},
		Available:                {"Available", "Available from us (form of availability unspecified).", "Tilgjengelig", "Tilgjengelig fra oss. Type tilgjengelig er ikke spesifisert."},
		InStock:                  {"In stock", "Available from us as a stock item.", "Tilgjengelig på lager", "Tilgjengelig fra vårt lager"},
		ToOrder:                  {"To order", "Available from us as a non-stock item, by special order.", "Skaffes på bestilling", "Tilgjengelig fra oss, men lagerføres ikke og bestilles på ordre"},
		POD:                      {"POD", "Available from us by print-on-demand.", "POD - Print on demand", "Tilgjengelig fra oss som print on demand"},
		TemporarilyUnavailable:                                {"Temporarily unavailable", "Temporarily unavailable: temporarily unavailable from us (reason unspecified) (requires expected date, either as <ExpectedShipDate> (ONIX 2.1) or as <SupplyDate> with <SupplyDateRole> coded ‘08’ (ONIX 3.0), except in exceptional circumstances where no date is known).", "Midlertidig ikke tilgjengelig", "Midlertidig utilgjengelig fra oss. Årsak uspesifisert. Krever dato i enten <ExpectedShipDate> (Onix 2.1) eller i <SupplyDate> med kode ’08’ i <SupplyDateRole> (Onix 3.0), unntatt i spesielle tilfeller hvor datoen er ukjent"},
		OutOfStock:                                            {"Out of stock", "Stock item, temporarily out of stock (requires expected date, either as <ExpectedShipDate> (ONIX 2.1) or as <SupplyDate> with <SupplyDateRole> coded ‘08’ (ONIX 3.0), except in exceptional circumstances where no date is known).", "Midlertidig tomt på lager", "Lagerføres, midlertidig tomt på lager. Krever dato i enten <ExpectedShipDate> (Onix 2.1) eller i <SupplyDate> med kode ’08’ i <SupplyDateRole> (Onix 3.0), unntatt i spesielle tilfeller hvor datoen er ukjent"},
		Reprinting:                                            {"Reprinting", "Temporarily unavailable, reprinting (requires expected date, either as <ExpectedShipDate> (ONIX 2.1) or as <SupplyDate> with <SupplyDateRole> coded ‘08’ (ONIX 3.0), except in exceptional circumstances where no date is known).", "Under opptrykk", "Midlertidig utilgjengelig, under opptrykk. Krever dato, enten i <ExpectedShipDate> (Onix 2.1), eller som <SupplyDate> med kode ’08’ i <SupplyDateRole>, unntatt i spesielle tilfeller hvor datoen er ukjent."},
		AwaitingReissue:                                       {"Awaiting reissue", "Temporarily unavailable, awaiting reissue (requires the <Reissue> composite, and expected date, either as <ExpectedShipDate> (ONIX 2.1) or as <SupplyDate> with <SupplyDateRole> coded ‘08’ (ONIX 3.0), except in exceptional circumstances where no date is known).", "Nytt opplag ventes", "Midlertidig utilgjengelig, venter på ny utgave. krever <Reissue> composite, og dato enten i <ExpectedShipDate> (Onix 2.1) eller som <SupplyDate> med kode ’08’ i <SupplyDateRole> (Onix 3.0), unntatt i spesielle tilfeller hvor datoen er ukjent"},
		TemporarilyWithdrawnFromSale:                          {"Temporarily withdrawn from sale", "May be for quality or technical reasons. Requires expected availability date, either as <ExpectedShipDate> (ONIX 2.1) or as <SupplyDate> with <SupplyDateRole> coded ‘08’ (ONIX 3.0), except in exceptional circumstances where no date is known.", "Midlertidig trukket tilbake fra salg", "Midlertidig trukket tilbake, som oftest pga. forhold rundt kvalitet eller teknisk årsaker. Krever dato, enten i <ExpectedShipDate> (Onix 2.1) eller i <SupplyDate> med kode ’08’ i <SupplyDateRole>, unntatt i spesielle tilfeller hvor datoen er ukjent."},
		NotAvailableReasonUnspecified:                         {"Not available (reason unspecified)", "Not available from us (for any reason).", "Ikke tilgjengelig fra oss", "Ikke tilgjengelig fra oss. Årsak uspesifisert. "},
		NotAvailableReplacedByNewProduct:                      {"Not available, replaced by new product", "This product is unavailable, but a successor product or edition is or will be available from us (identify successor in <RelatedProduct>).", "Erstattet av nytt produkt", "Produktet er ikke tilgjengelig, men et alternativt produkt eller utgave vil bli tilgjengelig fra oss. Angi alternativt produkt i <RelatedProduct>"},
		NotAvailableOtherFormatAvailable:                      {"Not available, other format available", "This product is unavailable, but the same content is or will be available from us in an alternative format (identify other format product in <RelatedProduct>).", "Annet format tilgjengelig", "Produktet er ikke tilgjengelig, men samme innholdet er eller vil bli tilgjengelig fra oss i et alternativt format. Angi alternativt produkt/format i <RelatedProduct>"},
		NoLongerSuppliedByUs:                                  {"No longer supplied by us", "Identify new supplier in <NewSupplier> if possible.", "Ikke lenger distribuert av oss", "Angi ny distributør/leverandør i <NewSupplier> om mulig"},
		ApplyDirect:                                           {"Apply direct", "Not available to trade, apply direct to publisher.", "Må bestilles direkte fra vareeier", "Kontakt vareeier/utgiver direkte"},
		NotSoldSeparately:                                     {"Not sold separately", "Must be bought as part of a set (identify set in <RelatedProduct>).", "Selges ikke enkeltvis", "Kan kun kjøpes som del av et verk. Angi verket i <RelatedProduct>"},
		WithdrawnFromSale:                                     {"Withdrawn from sale", "May be for legal reasons or to avoid giving offence.", "Trukket tilbake fra salg", "Kan være av juridiske årsaker"},
		Remaindered:                                           {"Remaindered", "Remaindered.", "Reservert nedsettelse", "Restopplaget er reservert til mammut eller lignende salgskampanjer. Ikke i ordinært salg lenger"},
		NotAvailableReplacedByPOD:                             {"Not available, replaced by POD", "Out of print, but a print-on-demand edition is or will be available under a different ISBN. Use only when the POD successor has a different ISBN, normally because different trade terms apply.", "Utsolgt, erstattes av POD", "Utsolgt, men \"print-on-demand\" utgave er eller vil bli tilgjengelig under et annet ISBN. Brukes vare når POD utgaven har et annet ISBN fordi handelsbetingelsene da vil være forskjellige"},
		Recalled:                                              {"Recalled", "Recalled for reasons of consumer safety.", "Tilbakekalt", "Recalled for reasons of consumer safety."},
		NotSoldAsSet:                                          {"Not sold as set", "When a collection that is not sold as a set nevertheless has its own ONIX record.", "Selges ikke samlet", "Brukes når en 'collecttion' ikke selges samleg, men likevel har en egen Onix-post."},
		NotAvailablePublisherIndicatesOP:                      {"Not available, publisher indicates OP", "This product is unavailable, no successor product or alternative format is available or planned. Use this code only when the publisher has indicated the product is out of print.", "Ikke tilgjengelig, utgiver/vareeier angir at produktet er utsolgt", "Produktet er ikke tilgjengelig. Det er ikke planer om å utgi et erstatningsprodukt eller produktet i et annet format. Bruk koden kun når utgiver/vareeier har angitt at produktet er utsolgt."},
		NotAvailablePublisherNoLongerSellsProductInThisMarket: {"Not available, publisher no longer sells product in this market", "This product is unavailable in this market, no successor product or alternative format is available or planned. Use this code when a publisher has indicated the product is permanently unavailable (in this market) while remaining available elsewhere.", "Ikke tilgjengelig. Utgiver/vareeier selger ikke lenger produktet i dette markedet", "Produktet er ikke tilgjengelig i dette markedet. Det er ikke planer om å utgi et erstatningsprodukt eller produktet i et annet format. Bruk koden kun når utgiver/vareeier har angitt at produktet ikke er tilgjengelig lenger (i dette markedet), men kan være tilgjengelig andre steder."},
		NoRecentUpdateReceived:                                {"No recent update received", "Sender has not received any recent update for this product from the publisher/supplier (for use when the sender is a data aggregator): the definition of “recent” must be specified by the aggregator, or by agreement between parties to an exchange.", "Ukjent, ikke mottatt oppdateringer i det siste", "Avsender har ikke mottatt oppdateringer (i det siste) fra utgiver/distributør om dette produktet (brukes når avsender er et bibliografisk firma): definisjonen av ’i det siste’ må defineres av avsender, eller i en avtale mellom de partene som utveksler metadata."},
		NoLongerReceivingUpdates:                              {"No longer receiving updates", "Sender is no longer receiving any updates from the publisher/supplier of this product (for use when the sender is a data aggregator).", "Ukjent, oppdateringer mottas ikke", "Avsender mottar ikke lenger oppdateringer fra utgiver/distributor for dette produktet (brukes når avsender er et bibliografisk firma)."},
		ContactSupplier:                                       {"Contact supplier", "Availability not known to sender.", "Kontakt kundetjenesten", "Kontakt kundetjenesten for mer informasjon"},
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
