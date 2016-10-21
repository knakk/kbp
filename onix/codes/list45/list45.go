package list45

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Publisher                          = "01"
	CoPublisher                        = "02"
	Sponsor                            = "03"
	PublisherOfOriginalLanguageVersion = "04"
	HostDistributorOfElectronicContent = "05"
	PublishedForOnBehalfOf             = "06"
	PublishedInAssociationWith         = "07"
	PublishedOnBehalfOf                = "08"
	NewOrAcquiringPublisher            = "09"
	PublishingGroup                    = "10"
	PublisherOfFacsimileOriginal       = "11"
	RepackagerOfPreboundEdition        = "12"
	FormerPublisher                    = "13"
	PublicationFunder                  = "14"
	ResearchFunder                     = "15"
	FundingBody                        = "16"
	Printer                            = "17"
	Binder                             = "18"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Publisher:   {"Publisher", "", "Vareeier/utgiver", ""},
		CoPublisher: {"Co-publisher", "Use where two or more publishers co-publish the exact same product, either under a single ISBN (in which case both publishers are co-publishers), or under different ISBNs (in which case the publisher of THIS ISBN is the publisher and the publishers of OTHER ISBNs are co-publishers. Note this is different from publication of ‘co-editions’.", "Sampubliserert med", "Use where two or more publishers co-publish the exact same product, either under a single ISBN (in which case both publishers are co-publishers), or under different ISBNs (in which case the publisher of THIS ISBN is the publisher and the publishers of OTHER ISBNs are co-publishers. Note this is different from publication of ‘co-editions’."},
		Sponsor:     {"Sponsor", "", "Sponsor", ""},
		PublisherOfOriginalLanguageVersion: {"Publisher of original-language version", "Of a translated work.", "Forlag/utgiver for versjonen på orginalspråket", "For oversatt arbeid"},
		HostDistributorOfElectronicContent: {"Host/distributor of electronic content", "", "Vert/distributør av elektronisk innhold", ""},
		PublishedForOnBehalfOf:             {"Published for/on behalf of", "", "Utgitt for/på vegne av", ""},
		PublishedInAssociationWith:         {"Published in association with", "Use also for “Published in cooperation with”.", "Utgitt i samarbeid med", "Bruk også for publisert sammen med "},
		PublishedOnBehalfOf:                {"Published on behalf of", "DEPRECATED: use code 06.", "Utgitt på vegne av", "UTGÅR: bruk kode 06"},
		NewOrAcquiringPublisher:            {"New or acquiring publisher", "When ownership of a product or title is transferred from one publisher to another.", "Ny eller overtagende utgiver", "Når eierskap til et produkt eller tittel er overført fra en utgiver til en annen"},
		PublishingGroup:                    {"Publishing group", "The group to which a publisher (publishing role 01) belongs: use only if a publisher has been identified with role code 01.", "Utgivergruppe", "Den gruppa en vareeier/utgiver (kode 01) hører til. Brukes kun dersom en vareeier/utgiver har blitt identifisert med kode 01."},
		PublisherOfFacsimileOriginal:       {"Publisher of facsimile original", "The publisher of the edition of which a product is a facsimile.", "Utgiver av original (for faksimile)", "Utgiveren av den utgava som produktet er en faksimile av."},
		RepackagerOfPreboundEdition:        {"Repackager of prebound edition", "The repackager of a prebound edition that has been assigned its own identifier. (In the US, a ‘prebound edition’ is a book that was previously bound, normally as a paperback, and has been rebound with a library-quality hardcover binding by a supplier other than the original publisher.) Required when the <EditionType> is coded PRB. The original publisher should be named as the ‘publisher’.", "Repackager of prebound edition", "USA: The repackager of a prebound edition that has been assigned its own identifier. (In the US, a ‘prebound edition’ is a book that was previously bound, normally as a paperback, and has been rebound with a library-quality hardcover binding by a supplier other than the original publisher.) Required when the <EditionType> is coded PRB. The original publisher should be named as the ‘publisher’."},
		FormerPublisher:                    {"Former publisher", "When ownership of a product or title is transferred from one publisher to another (complement of code 09).", "Tidligere utgiver", "Når eierskapet til et produkt overføres fra en vareeier til en annen (brukes sammen med kode 09)"},
		PublicationFunder:                  {"Publication funder", "Body funding publication fees, if different from the body funding the underlying research. For use with open access publications.", "Ansvarlig for finansiering av publikasjon", "Den som finansierer publikasjonen, dersom denne er forskjellig fra den som finansierer den underliggende forskninga. Brukes med open access-publikasjoner."},
		ResearchFunder:                     {"Research funder", "Body funding the research on which publication is based, if different from the body funding the publication. For use with open access publications.", "Ansvarlig for finansiering av forskning", "Den som finansierer forskningen publikasjonen baserer seg på, dersom denne er forskjellig fra den som finansierer utgivelsen av publikasjonen. Brukes med open access-publikasjoner."},
		FundingBody:                        {"Funding body", "Body funding research and publication. For use with open access publications.", "Ansvarlig for finansiering", "Den som finansierer forskning og publisering. Brukes med open access-publikasjoner."},
		Printer:                            {"Printer", "Organisation responsible for printing a printed product. Supplied primarily to meet legal deposit requirements, and may apply only to the first impression. The organisation may also be responsible for binding, when a separate binder is not specified.", "Trykkeri", "Organisasjon som er ansvarlig for trykking av et fysisk produkt. Angis primært for å oppfylle krav til til pliktavlevering, og det kan hende det gjelder kun det første opplaget. Organisasjonen kan også være ansvarlig for innbinding, når en separat organisasjon for dette ikke er angitt."},
		Binder:                             {"Binder", "Organisation responsible for binding a printed product (where distinct from the printer). Supplied primarily to meet legal deposit requirements, and may apply only to the first impression.", "Bokbinder", "Organisasjon som er ansvarlig for innbinding av et trykt product (der organisasjonen er en annen enn trykkeriet). Angis primært for å oppfylle krav til pliktavlevering, og det kan hende det gjelder kun det første opplaget."},
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
