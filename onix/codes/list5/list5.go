package list5 // Product identifier type code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Product identifier type code
const (
	Proprietary        = "01"
	ISBN10             = "02"
	GTIN13             = "03"
	UPC                = "04"
	ISMN10             = "05"
	DOI                = "06"
	LCCN               = "13"
	GTIN14             = "14"
	ISBN13             = "15"
	LegalDepositNumber = "17"
	URN                = "22"
	OCLCNumber         = "23"
	CoPublishersISBN13 = "24"
	ISMN13             = "25"
	ISBNA              = "26"
	JPECode            = "27"
	OLCCNumber         = "28"
	JPMagazineID       = "29"
	BNFControlNumber   = "31"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "13", "14", "15", "17", "22", "23", "24", "25", "26", "27", "28", "29", "31"}

	all = map[string]struct {
		labelEn, notesEn, labelNo, notesNo string
	}{
		Proprietary:        {"Proprietary", "For example, a publisher’s or wholesaler’s product number.", "Proprietær", "F.eks., et forlags eller grossists produktnummer"},
		ISBN10:             {"ISBN-10", "International Standard Book Number, pre-2007, unhyphenated (10 characters) – now DEPRECATED in ONIX for Books, except where providing historical information for compatibility with legacy systems. It should only be used in relation to products published before 2007 – when ISBN-13 superseded it – and should never be used as the ONLY identifier (it should always be accompanied by the correct GTIN-13 / ISBN-13).", "ISBN-10", "International Standard Book Number, pre-2007, uten bindestreker (10 siffer) – nå UTGÅTT i Onix for Books, bortsett fra hvor det er nøvendig for å sende historisk informasjon for å sikre kompatibilitet med andre systemer. Det bør bare brukes for produkter utgitt før 2007, og bør aldri brukes som eneste identifikator (man bør alltid også sende det korrekte GTIN-13/ISBN-13)"},
		GTIN13:             {"GTIN-13", "GS1 Global Trade Item Number, formerly known as EAN article number (13 digits).", "GTIN-13", "GS1 Global trade item number, tidligere EAN (13 siffer)"},
		UPC:                {"UPC", "UPC product number (12 digits).", "UPC", "UPC produktnummer (12 siffer)"},
		ISMN10:             {"ISMN-10", "International Standard Music Number (M plus nine digits). Pre-2008 – now DEPRECATED in ONIX for Books, except where providing historical information for compatibility with legacy systems. It should only be used in relation to products published before 2008 – when ISMN-13 superseded it – and should never be used as the ONLY identifier (it should always be accompanied by the correct ISMN-13).", "ISMN", "International Standard Music Number (10 siffer). Pre-2008 – nå UTGÅTT i Onix for books, bortsett fra hvor det er nøvendig for å sende historisk informasjon for å sikre kompatibilitet med andre systemer. Det bør bare brukes for produkter utgitt før 2008, og bør aldri brukes som eneste identifikator (man bør alltid også sende det korrekte ISMN-13)"},
		DOI:                {"DOI", "Digital Object Identifier (variable length and character set).", "DOI", "Digital Object Identifier (variabel lengde og tegnsett)"},
		LCCN:               {"LCCN", "Library of Congress Control Number (12 characters, alphanumeric).", "LCCN", "Library of Congress Control Number (12 tegn, alfanumerisk)"},
		GTIN14:             {"GTIN-14", "GS1 Global Trade Item Number (14 digits).", "GTIN-14", "GS1 Global Trade Item Number (14 siffer)"},
		ISBN13:             {"ISBN-13", "International Standard Book Number, from 2007, unhyphenated (13 digits starting 978 or 9791–9799).", "ISBN-13", "International Standard Book Number, fra 2007, Uten bindestreker (13-siffer, starter med 978 eller 9791-9799)"},
		LegalDepositNumber: {"Legal deposit number", "The number assigned to a publication as part of a national legal deposit process.", "Legal deposit number", "The number assigned to a publication as part of a national legal deposit process"},
		URN:                {"URN", "Uniform Resource Name: note that in trade applications an ISBN must be sent as a GTIN-13 and, where required, as an ISBN-13 – it should not be sent as a URN.", "URN", "Uniform Resource Name"},
		OCLCNumber:         {"OCLC number", "A unique number assigned to a bibliographic item by OCLC.", "OCLC number", "En unik identifikator fra OCLC"},
		CoPublishersISBN13: {"Co-publisher’s ISBN-13", "An ISBN-13 assigned by a co-publisher. The ‘main’ ISBN sent with ID type code 03 and/or 15 should always be the ISBN that is used for ordering from the supplier identified in Supply Detail. However, ISBN rules allow a co-published title to carry more than one ISBN. The co-publisher should be identified in an instance of the <Publisher> composite, with the applicable <PublishingRole> code.", "Medutgivers ISBN-13", "Et ISBN-13 som tilhører en medutgiver av produktet. Hoved-ISBN-et som sendes med ID type code 03 og/eller 15 bør alltid være det ISBN-et som brukes for bestilling fra distributøren idenfifisert i Supply detail. Imidlertid tillater ISBN-regelverket et produkt som utgis av to utgivere å ha to ISBN. Medutgiveren bør identifiseres i <Publisher>, med den riktige <PublishingRole>-koden."},
		ISMN13:             {"ISMN-13", "International Standard Music Number, from 2008 (13-digit number starting 9790).", "ISMN-13", "International Standard Music Number, fra 2008 (13-siffer,  starter med 9790)."},
		ISBNA:              {"ISBN-A", "Actionable ISBN, in fact a special DOI incorporating the ISBN-13 within the DOI syntax. Begins ‘10.978.’ or ‘10.979.’ and includes a / character between the registrant element (publisher prefix) and publication element of the ISBN, eg 10.978.000/1234567. Note the ISBN-A should always be accompanied by the ISBN itself, using codes 03 and/or 15.", "ISBN-A", "Actionable ISBN, in fact a special DOI incorporating the ISBN-13 within the DOI syntax. Begins ‘10.978.’ or ‘10.979.’ and includes a / character between the registrant element (publisher prefix) and publication element of the ISBN, eg 10.978.000/1234567. Note the ISBN-A should always be accompanied by the ISBN itself, using codes 03 and/or 15."},
		JPECode:            {"JP e-code", "E-publication identifier controlled by JPOIID’s Committee for Research and Management of Electronic Publishing Codes.", "JP e-code", "E-publication identifier controlled by JPOIID’s Committee for Research and Management of Electronic Publishing Codes."},
		OLCCNumber:         {"OLCC number", "Unique number assigned by the Chinese Online Library Cataloging Center (see http://olcc.nlc.gov.cn).", "OLCC nummer", "Unique number assigned by the Chinese Online Library Cataloging Center (see http://olcc.nlc.gov.cn)."},
		JPMagazineID:       {"JP Magazine ID", "Japanese magazine identifier, similar in scope to ISSN but identifying a specific issue of a serial publication. Five digits to identify the periodical, plus a hyphen and two digits to identify the issue.", "JP Magazine ID", "Japanese magazine identifier, similar in scope to ISSN but identifying a specific issue of a serial publication. Five digits to identify the periodical, plus a hyphen and two digits to identify the issue."},
		BNFControlNumber:   {"BNF control number", "Numéro de la notice BNF", "BNF control number", "Numéro de la notice BNF"},
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
	} else {
		return codes.Item{Code: code, Label: item.labelEn, Notes: item.notesEn}, nil
	}
}

// MustItem returns the Item for the given code. It will panic if it doesn't exist.
func MustItem(code string, lang codes.Language) codes.Item {
	item, ok := all[code]
	if !ok {
		panic(fmt.Errorf("no item with code: %q", code))
	}
	if lang == codes.Norwegian {
		return codes.Item{Code: code, Label: item.labelNo, Notes: item.notesNo}
	} else {
		return codes.Item{Code: code, Label: item.labelEn, Notes: item.notesEn}
	}
}
