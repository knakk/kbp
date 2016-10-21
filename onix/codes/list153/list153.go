package list153

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	SenderDefinedText                       = "01"
	ShortDescriptionAnnotation              = "02"
	Description                             = "03"
	TableOfContents                         = "04"
	FlapCoverCopy                           = "05"
	ReviewQuote                             = "06"
	ReviewQuotePreviousEdition              = "07"
	ReviewQuotePreviousWork                 = "08"
	Endorsement                             = "09"
	PromotionalHeadline                     = "10"
	Feature                                 = "11"
	BiographicalNote                        = "12"
	PublishersNotice                        = "13"
	Excerpt                                 = "14"
	Index                                   = "15"
	ShortDescriptionAnnotationForCollection = "16"
	DescriptionForCollection                = "17"
	NewFeature                              = "18"
	VersionHistory                          = "19"
	OpenAccessStatement                     = "20"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		SenderDefinedText:          {"Sender-defined text", "To be used only in circumstances where the parties to an exchange have agreed to include text which (a) is not for general distribution, and (b) cannot be coded elsewhere. If more than one type of text is sent, it must be identified by tagging within the text itself.", "Senderdefinert tekst", "Brukes for tekst som ikke faller inn under noen av de andre kodene. Brukes bl.a. for bibliografisk note"},
		ShortDescriptionAnnotation: {"Short description/annotation", "Limited to a maximum of 350 characters.", "Kort beskrivelse", "Kort beskrivelse av produktet"},
		Description:                {"Description", "Length unrestricted.", "Forlagsomtale", "Forlagets egen tekst om produktet. Denne koden brukes til forlagets omtale, uavhengig av omtalens lengde"},
		TableOfContents:            {"Table of contents", "Used for a table of contents sent as a single text field, which may or may not carry structure expressed as XHTML.", "Innholdsfortegnelse", "Brukes som innholdsfortegnelse som sendes som et enkelt tekstfelt, som kan, men ikke nødvendigvis innholder strukturen uttrykt som XHTML."},
		FlapCoverCopy:              {"Flap / cover copy", "Descriptive blurb taken from the back cover and/or flaps.", "Beskrivende blurb", "Beskrivende blurb hentet fra omslagets bakside og/eller innbretter"},
		ReviewQuote:                {"Review quote", "A quote taken from a review of the product or of the work in question where there is no need to take account of different editions.", "Sitat fra anmeldelse", "Sitat hentet fra en anmeldelse av produktet, hvor det ikke er nødvendig å skille på ulike utgaver"},
		ReviewQuotePreviousEdition: {"Review quote: previous edition", "A quote taken from a review of a previous edition of the work.", "Sitat fra anmeldelse av tidligere utgave", "Sitat hentet fra en anmeldelse av en tidligere utgave av verket"},
		ReviewQuotePreviousWork:    {"Review quote: previous work", "A quote taken from a review of a previous work by the same author(s) or in the same series.", "Sitat fra anmeldelse av tidligere verk", "Sitat hentet fra anmeldelse av et tidligere verk av den samme forfatteren/de samme forfatterne eller som er i samme serie"},
		Endorsement:                {"Endorsement", "A quote usually provided by a celebrity to promote a new book, not from a review.", "Blurb", "Sitat, vanligvis av en kjendis for å markedsføre ei ny bok, ikke fra en anmeldelse"},
		PromotionalHeadline:        {"Promotional headline", "A promotional phrase which is intended to headline a description of the product.", "Markedsførende overskrift", "Overskrift. Setning ment brukt til markedsføring og som overskrift til en beskrivelse av produktet"},
		Feature:                    {"Feature", "Text describing a feature of a product to which the publisher wishes to draw attention for promotional purposes. Each separate feature should be described by a separate repeat, so that formatting can be applied at the discretion of the receiver of the ONIX record, or multiple features can be described using appropriate XHTML markup.", "Egenskap/særpreg", "Tekst som beskriver et særpreg eller en karakteristisk egenskap ved produktet, som forlaget ønsker å framheve i markedsføringsøyemed. "},
		BiographicalNote:           {"Biographical note", "A note referring to all contributors to a product – NOT linked to a single contributor.", "Biografisk tekst", "Tekst som omhandler alle bidragsytere - IKKE knyttet til en enkelt bidragsyter"},
		PublishersNotice:           {"Publisher’s notice", "A statement included by a publisher in fulfillment of contractual obligations, such as a disclaimer, sponsor statement, or legal notice of any sort. Note that the inclusion of such a notice cannot and does not imply that a user of the ONIX record is obliged to reproduce it.", "Informasjon fra forlaget", "Informasjon fra forlaget for å oppfylle kontraktuelle forpliktelser, som f.eks. en fraskrivningsklausul eller annen juridisk informasjon. Merk at selv om denne informasjonen sendes i Onix-posten, betyr ikke det at mottakeren er forpliktet til å reprodusere den."},
		Excerpt:                    {"Excerpt", "A short excerpt from the work.", "Utdrag", "Et kort utdrag fra verket"},
		Index:                      {"Index", "Used for an index sent as a single text field, which may be structured using XHTML.", "Register", "Brukes når registeret sendes som et enkelt tekstfelt. Kan være strukturert ved bruk av XHTML."},
		ShortDescriptionAnnotationForCollection: {"Short description/annotation for collection", "(of which the product is a part.) Limited to a maximum of 350 characters.", "Kort beskrivelse av samling/serie (collection)", "Av serien som produktet er en del av. Maks 350 tegn"},
		DescriptionForCollection:                {"Description for collection", "(of which the product is a part.) Length unrestricted.", "Beskrivelse av samling/serie (collection)", "Av serien som produktet er en del av."},
		NewFeature:                              {"New feature", "As code 11 but used for a new feature of this edition or version.", "Ny egenskap/nytt særpreg", "Som kode 11, men blir brukt for en ny egenskap som er spesiell for denne utgaven eller versjonen"},
		VersionHistory:                          {"Version history", "", "Versjonshistorie", ""},
		OpenAccessStatement:                     {"Open access statement", "Short summary statement of open access status and any related conditions (eg “Open access – no commercial use”), primarily for marketing purposes. Should always be accompanied by a link to the complete license (see code 99 in List 158).", "Open access statement", "Short summary statement of open access status and any related conditions (eg “Open access – no commercial use”), primarily for marketing purposes. Should always be accompanied by a link to the complete license (see code 99 in List 158)."},
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
