package list71

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	UnspecifiedSeeText                = "00"
	RetailerExclusiveOwnBrand         = "01"
	OfficeSuppliesEdition             = "02"
	InternalPublisherUseOnlyDoNotList = "03"
	RetailerExclusive                 = "04"
	RetailerOwnBrand                  = "05"
	LibraryEdition                    = "06"
	SchoolsOnlyEdition                = "07"
	Indiziert                         = "08"
	NotForSaleToLibraries             = "09"
	NewsOutletEdition                 = "10"
	RetailerException                 = "11"
	NotForSaleToSubscriptionServices  = "12"
	SubscriptionServicesOnly          = "13"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		UnspecifiedSeeText:                {"Unspecified – see text", "Restriction must be described in <SalesRestrictionDetail> (ONIX 2.1) or <SalesRestrictionNote> (ONIX 3.0).", "Uspesifisert - se tekst", "Restriksjoner må beskrives i <SalesRestrictionDetail> (Onix 2.1) eller <SalesRestrictionNote> (Onix 3.0)"},
		RetailerExclusiveOwnBrand:         {"Retailer exclusive / own brand", "For sale only through designated retailer. Retailer must be identified or named in an instance of the <SalesOutlet> composite. Use only when it is not possible to assign the more explicit code 04 or 05.", "Bokhandlerutgave / eget merke", "Kun for salg gjennom bestemt forhandler. Forhandler må bli navngitt i <SalesOutletName>"},
		OfficeSuppliesEdition:             {"Office supplies edition", "For editions sold only though office supplies wholesalers. Retailer(s) and/or distributor(s) may be identified or named in an instance of the <SalesOutlet> composite.", "Kontorrekvisitautgave", "For utgaver som kun selges gjennom forhandlere av kontorrekvisita. Forhandler(e) og/eller distributør(er) må navngis i <SalesOutlet>"},
		InternalPublisherUseOnlyDoNotList: {"Internal publisher use only: do not list", "For an ISBN that is assigned for a publisher’s internal purposes.", "Kun for internt forlagsbruk: skal ikke vises", "For et ISBN som er tildelt kun for et forlags interne bruk"},
		RetailerExclusive:                 {"Retailer exclusive", "For sale only through designated retailer, though not under retailer’s own brand/imprint. Retailer must be identified or named in an instance of the <SalesOutlet> composite.", "Bokhandlerutgave", "Kun for salg gjennom bestemte bokhandlere/utsalg, men ikke med forhandlerens eget merkenavn eller imprint.Forhandler må navngis <SalesOutlet>."},
		RetailerOwnBrand:                  {"Retailer own brand", "For sale only through designated retailer under retailer’s own brand/imprint. Retailer must be identified or named in an instance of the <SalesOutlet> composite.", "Bokhandlers eget merke", "Kun for salg gjennom bestemte bokhandlere/utsalg under forhandlerens eget merkenavn eller imprint. Forhandler må navngis  <SalesOutlet>."},
		LibraryEdition:                    {"Library edition", "For sale to libraries only; not for sale through retail trade.", "Bibliotekutgave\n", "Kun for salg til biblioteker. Ikke for salg gjennom forhandler eller andre salgsledd."},
		SchoolsOnlyEdition:                {"Schools only edition", "For sale directly to schools only; not for sale through retail trade.", "Skoleutgave", "Kun for salg direkte til skoler"},
		Indiziert:                         {"Indiziert", "Indexed for the German market – in Deutschland indiziert.", "Indiziert", "Indeksert for det tyske markedet– tysk: indiziert."},
		NotForSaleToLibraries:             {"Not for sale to libraries", "Expected to apply in particular to digital products for consumer sale where the publisher does not permit the product to be supplied to libraries who provide an ebook loan service.", "Ikke for salg til bibliotek", "Forventes å bli brukt for digitale produkter hvor utgiver ikke tillater at produktet selges til bibliotek som har utlånsløsninger for e-bøker"},
		NewsOutletEdition:                 {"News outlet edition", "For editions sold only through newsstands/newsagents.", "Kioskutgave", "For utgaver som selges i kiosker"},
		RetailerException:                 {"Retailer exception", "Not for sale through designated retailer. Retailer must be identified or named in an instance of the <SalesOutlet> composite.", "Forhandlerunntak", "Ikke for salg gjennom angitt forhandler. Forhandler må identifiseres eller navngis i <SalesOutlet>"},
		NotForSaleToSubscriptionServices:  {"Not for sale to subscription services", "Not for sale to organisations or services offering consumers subscription access to a library of books", "Ikke for abonnementstjenester", "Ikke for salg via tjenester som tilbyr abonnementstilgang til et utvalg bøker"},
		SubscriptionServicesOnly:          {"Subscription services only", "Restricted to organisations or services offering consumers subscription access to a library of books", "Kun for abonnementstjenester", "Kun tilgjengelig for tjenester som tilbyr abonnementstjenester som gir kundene tilgang til et utvalg av bøker"},
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
