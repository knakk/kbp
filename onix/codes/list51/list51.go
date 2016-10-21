package list51

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Unspecified                              = "00"
	Includes                                 = "01"
	IsPartOf                                 = "02"
	Replaces                                 = "03"
	ReplacedBy                               = "05"
	AlternativeFormat                        = "06"
	HasAncillaryProduct                      = "07"
	IsAncillaryTo                            = "08"
	IsRemainderedAs                          = "09"
	IsRemainderOf                            = "10"
	IsOtherLanguageVersionOf                 = "11"
	PublishersSuggestedAlternative           = "12"
	EpublicationBasedOnPrintProduct          = "13"
	EpublicationIsDistributedAs              = "14"
	EpublicationIsARenderingOf               = "15"
	PODReplacementFor                        = "16"
	ReplacedByPOD                            = "17"
	IsSpecialEditionOf                       = "18"
	HasSpecialEdition                        = "19"
	IsPreboundEditionOf                      = "20"
	IsOriginalOfPreboundEdition              = "21"
	ProductBySameAuthor                      = "22"
	SimilarProduct                           = "23"
	IsFacsimileOf                            = "24"
	IsOriginalOfFacsimile                    = "25"
	IsLicenseFor                             = "26"
	ElectronicVersionAvailableAs             = "27"
	EnhancedVersionAvailableAs               = "28"
	BasicVersionAvailableAs                  = "29"
	ProductInSameCollection                  = "30"
	HasAlternativeInADifferentMarketSector   = "31"
	HasEquivalentIntendedForADifferentMarket = "32"
	HasAlternativeIntendedForDifferentMarket = "33"
	Cites                                    = "34"
	IsCitedBy                                = "35"
	SalesExpectation                         = "36"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Unspecified:                              {"Unspecified", "<Product> is related to <RelatedProduct> in a way that cannot be specified by another code value.", "Uspesifisert", "<Product> er relatert til <RelatedProduct> på en måte som ikke kan beskrives vha. andre koder i denne listen"},
		Includes:                                 {"Includes", "<Product> includes <RelatedProduct>.", "Inkluderer", "<Product> inkluderer <RelatedProduct>"},
		IsPartOf:                                 {"Is part of", "<Product> is part of <RelatedProduct>: use for ‘also available as part of’.", "Er del av", "<Product> er en del av <RelatedProduct>: brukes for ”også tilgjengelig som del av”."},
		Replaces:                                 {"Replaces", "<Product> replaces, or is new edition of, <RelatedProduct>.", "Erstatter", "<Product> erstatter eller er ny utgave av <RelatedProduct>."},
		ReplacedBy:                               {"Replaced by", "<Product> is replaced by, or has new edition, <RelatedProduct> (reciprocal of code 03).", "Erstattet av", "<Product> er erstattet av, eller har ny utgave, <RelatedProduct> (motsvarer kode 03)"},
		AlternativeFormat:                        {"Alternative format", "<Product> is available in an alternative format as <RelatedProduct> – indicates an alternative format of the same content which is or may be available.", "Alternativt format", "<Product> er tilgengelig i et alternativt format i <RelatedProduct>. Indikasjon om at alternativt format med samme innhold er eller kan gjøres tilgjengelig"},
		HasAncillaryProduct:                      {"Has ancillary product", "<Product> has an ancillary or supplementary product <RelatedProduct>.", "Har tilleggsprodukt", "<Product> har et del- eller tilleggsprodukt <RelatedProduct>"},
		IsAncillaryTo:                            {"Is ancillary to", "<Product> is ancillary or supplementary to <RelatedProduct>.", "Er tilleggsprodukt til", "<Product> er del- eller tilleggsprodukt til <RelatedProduct>"},
		IsRemainderedAs:                          {"Is remaindered as", "<Product> is remaindered as <RelatedProduct>, when a remainder merchant assigns its own identifier to the product.", "Restopplaget identifisert som", "<Product> har et restopplag som er tildelt en egen identifikator hos en forhandler/distributør <RelatedProduct>"},
		IsRemainderOf:                            {"Is remainder of", "<Product> was originally sold as <RelatedProduct>, indicating the publisher’s original identifier for a title which is offered as a remainder under a different identifier (reciprocal of code 09).", "Er restopplag av", "<Product> ble opprinnelig solgt som <RelatedProduct>. Angir utgiverens opprinnelige identifikator for en tittel som nå selges som et restopplag under en annen identifikator (motsvarer kode 09)"},
		IsOtherLanguageVersionOf:                 {"Is other-language version of", "<Product> is an other-language version of <RelatedProduct>.", "Er versjon på annet språk av\n", "<Product> er en versjon av <RelatedProduct> på et annet språk"},
		PublishersSuggestedAlternative:           {"Publisher’s suggested alternative", "<Product> has a publisher’s suggested alternative <RelatedProduct>, which does not, however, carry the same content (cf 05 and 06).", "Utgivers/vareeiers foreslåtte alternativ", "<Product> har et alternativ foreslått av utgiver/vareeier <RelatedProduct>, men som ikke har helt samme innhold (jf. 05 og 06)."},
		EpublicationBasedOnPrintProduct:          {"Epublication based on (print product)", "<Product> is an epublication based on printed product <RelatedProduct>.", "E-bok basert på (trykt publikasjon)", "<Product> er en epublikasjon som er basert på et trykt produkt <RelatedProduct>"},
		EpublicationIsDistributedAs:              {"Epublication is distributed as", "<Product> is an epublication ‘rendered’ as <RelatedProduct>: use in ONIX 2.1 only when the <Product> record describes a package of electronic content which is available in multiple ‘renderings’ (coded 000 in <EpubTypeCode>): NOT USED in ONIX 3.0.", "E-bok er distribuert som", "BRUKES IKKE I ONIX 3.0. Bruk når ONIX-posten beskriver en pakke med elektronisk innhold som er tilgjengelig i flere formater."},
		EpublicationIsARenderingOf:               {"Epublication is a rendering of", "<Product> is a ‘rendering’ of an epublication <RelatedProduct>: use in ONIX 2.1 only when the <Product> record describes a specific rendering of an epublication content package, to identify the package: NOT USED in ONIX 3.0.", "E-bok er en gjengivelse av", "BRUKES IKKE I ONIX 3.0. Bruk når ONIX-posten beskriver et bestemt format i en epublikasjonspakke"},
		PODReplacementFor:                        {"POD replacement for", "<Product> is a POD replacement for <RelatedProduct>. <RelatedProduct> is an out-of-print product replaced by a print-on-demand version under a new ISBN.", "POD erstatning for", "<Product> er en POD-erstatning for <RelatedProduct>. <RelatedProduct> er et trykt produkt som er utsolgt, og som er erstattet med et print on demand-produkt med nytt ISBN."},
		ReplacedByPOD:                            {"Replaced by POD", "<Product> is replaced by POD <RelatedProduct>. <RelatedProduct> is a print-on-demand replacement, under a new ISBN, for an out-of-print <Product> (reciprocal of code 16).", "Erstattet av POD", "<Product> er erstattet av POD <RelatedProduct>. <RelatedProduct> er en print on demand-erstatning, med et nytt ISBN, for et utsolgt <Product> (motsvarer kode 16)"},
		IsSpecialEditionOf:                       {"Is special edition of", "<Product> is a special edition of <RelatedProduct>. Used for a special edition (German: Sonderausgabe) with different cover, binding etc – more than ‘alternative format’ – which may be available in limited quantity and for a limited time.", "Er spesialutgave av", "<Product> er spesialutgave av <RelatedProduct>. Brukes for ei spesialutgave (tysk: Sonderausgabe) med ulikt omslag, innbinding osv., mer enn ’alternativt format’, som kan være tilgjengelig i et lite opplag i et avgrenset tidsrom."},
		HasSpecialEdition:                        {"Has special edition", "<Product> has a special edition <RelatedProduct> (reciprocal of code 18).", "Har spesialutgave", "<Product> har spesialutgave <RelatedProduct> (motsvarer kode 18)"},
		IsPreboundEditionOf:                      {"Is prebound edition of", "<Product> is a prebound edition of <RelatedProduct> (in the US, a prebound edition is ‘a book that was previously bound and has been rebound with a library quality hardcover binding. In almost all commercial cases, the book in question began as a paperback.’).", "Is prebound edition of", "<Product> is a prebound edition of <RelatedProduct> (in the US, a prebound edition is ‘a book that was previously bound and has been rebound with a library quality hardcover binding. In almost all commercial cases, the book in question began as a paperback.’)."},
		IsOriginalOfPreboundEdition:              {"Is original of prebound edition", "<Product> is the regular edition of which <RelatedProduct> is a prebound edition.", "Is original of prebound edition", "<Product> is the regular edition of which <RelatedProduct> is a prebound edition."},
		ProductBySameAuthor:                      {"Product by same author", "<Product> and <RelatedProduct> have a common author.", "Produkt av samme forfatter", "<Product> og <RelatedProduct> har en felles forfatter."},
		SimilarProduct:                           {"Similar product", "<RelatedProduct> is another product that is suggested as similar to <Product> (‘if you liked <Product>, you may also like <RelatedProduct>’).", "Lignende produkt", "<RelatedProduct> er et annet produkt som er foreslått som lignende til <Product> (’hvis du likte <Product>, kan det hende du også liker <RelatedProduct>)."},
		IsFacsimileOf:                            {"Is facsimile of", "<Product> is a facsimile edition of <RelatedProduct>.", "Er faksimile av", "<Product> er faksimile av <RelatedProduct>."},
		IsOriginalOfFacsimile:                    {"Is original of facsimile", "<Product> is the original edition from which a facsimile edition <RelatedProduct> is taken (reciprocal of code 25).", "Er originalen til faksimile", "<Product> er den originale utgaven som faksimilen baserer seg på <RelatedProduct> (motsvarer kode 25)."},
		IsLicenseFor:                             {"Is license for", "<Product> is a license for a digital <RelatedProduct>, traded or supplied separately.", "Er lisens til", "<Product> er lisens for et digitalt <RelatedProduct>, solgt eller levert separat."},
		ElectronicVersionAvailableAs:             {"Electronic version available as", "<RelatedProduct> is an electronic version of print <Product> (reciprocal of code 13).", "Digital versjon tilgjengelig som", "<RelatedProduct> er en digital versjon av trykt <Product> (motsvarer kode 13)"},
		EnhancedVersionAvailableAs:               {"Enhanced version available as", "<RelatedProduct> is an ‘enhanced’ version of <Product>, with additional content. Typically used to link an enhanced e-book to its original ‘unenhanced’ equivalent, but not specifically limited to linking e-books – for example, may be used to link illustrated and non-illustrated print books. <Product> and <RelatedProduct> should share the same <ProductForm>.", "Beriket versjon tilgjengelig som", "<RelatedProduct> er en ’beriket’ versjon av <Product>, med ekstra innhold. Brukes som oftest for å lenke en beriket e-bok til sitt ’originale’ motstykke, men er ikke avgrenset til kun å gjelde e-bøker. For eksempel kan koden også brukes til å lenke utgaver av trykte bøker med og uten illustrasjoner. <Product> og <RelatedProduct> bør ha samme <ProductForm>"},
		BasicVersionAvailableAs:                  {"Basic version available as", "<RelatedProduct> is a basic version of <Product> (reciprocal of code 28). <Product> and <RelatedProduct> should share the same <ProductForm>.", "Uberiket versjon tilgjengelig som", "<RelatedProduct> er den originale ('uberikede') versjonen av <Product> (motsvarer kode 28). <Product> og <RelatedProduct> bør ha samme <ProductForm>"},
		ProductInSameCollection:                  {"Product in same collection", "<RelatedProduct> and <Product> are part of the same collection (eg two products in same series or set).", "Produkt i samme 'serie' (collection)", "<RelatedProduct> og <Product> hører til samme serie (’collection’). "},
		HasAlternativeInADifferentMarketSector:   {"Has alternative in a different market sector", "<RelatedProduct> is an alternative product in another sector (of the same geographical market). Indicates an alternative that carries the same content, but available to a different set of customers, as one or both products are retailer-, channel- or market sector-specific.", "Har alternativ i en annen sektor av markedet", "<RelatedProduct> er et alternativt produkt i en annen sektor (i det same geografiske markedet). Indikterer et alternativ som har samme innhold, men som er tilgjengelig for en annen kundegruppe, siden et av eller begge produktene er forhandler- eller sektorspesifikt."},
		HasEquivalentIntendedForADifferentMarket: {"Has equivalent intended for a different market", "<RelatedProduct> is an equivalent product, often intended for another (geographical) market. Indicates an alternative that carries essentially the same content, though slightly adapted for local circumstances (as opposed to a translation – use code 11).", "Har tilsvarende produkt ment for et annet marked", "<RelatedProduct> er et tilsvarende produkt, ofte ment for et annet (geografisk) marked. Indikerer et alternativ som i hovedsak har samme innhold, men som kan være delvis tilpasset lokale forhold (i motsetning til en oversettelse – bruk kode 11)"},
		HasAlternativeIntendedForDifferentMarket: {"Has alternative intended for different market", "<RelatedProduct> is an alternative product, often intended for another (geographical) market. Indicates the content of the alternative is identical in all respects.", "Har et alternativ ment for et annet marked", "<RelatedProduct> er et alternativt produkt, ofte ment for et annet (geografisk) marked. Indikerer at innholdet er identisk med <Product>."},
		Cites:            {"Cites", "<Product> cites <RelatedProduct>.", "Siterer", "<Product> siterer <RelatedProduct>."},
		IsCitedBy:        {"Is cited by", "<Product> is the object of a citation in <RelatedProduct>.", "Er sitert av", "<Product> blir sitert i<RelatedProduct>."},
		SalesExpectation: {"Sales expectation", "Use to give the ISBN of another book that had sales (both in terms of copy numbers and customer  profile) comparable to that the publisher or distributor estimates for the product. Use in ONIX 2.1 ONLY", "Sales expectation", "Use to give the ISBN of another book that had sales (both in terms of copy numbers and customer  profile) comparable to that the publisher or distributor estimates for the product. Use in ONIX 2.1 ONLY"},
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
