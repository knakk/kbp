package list21

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	AbridgedEdition            = "ABR"
	ActingEdtion               = "ACT"
	AdaptedEdition             = "ADP"
	Alternate                  = "ALT"
	AnnotatedEdition           = "ANN"
	BilingualEdition           = "BLL"
	BrailleEdition             = "BRL"
	CombinedVolume             = "CMB"
	CriticalEdition            = "CRI"
	Coursepack                 = "CSP"
	DigitalOriginal            = "DGO"
	EnhancedEdition            = "ENH"
	EnlargedEdition            = "ENL"
	ExpurgatedEdition          = "EXP"
	FacsimileEdition           = "FAC"
	Festschrift                = "FST"
	IllustratedEdition         = "ILL"
	LargeTypeLargePrintEdition = "LTE"
	MicroprintEdition          = "MCP"
	MediaTieIn                 = "MDT"
	MultilingualEdition        = "MLL"
	NewEdition                 = "NED"
	EditionWithNumberedCopies  = "NUM"
	PreboundEdition            = "PRB"
	RevisedEdition             = "REV"
	SchoolEdition              = "SCH"
	SimplifiedLanguageEdition  = "SMP"
	SpecialEdition             = "SPE"
	StudentEdition             = "STU"
	TeachersEdition            = "TCH"
	UnabridgedEdition          = "UBR"
	UltraLargePrintEdition     = "ULP"
	UnexpurgatedEdition        = "UXP"
	VariorumEdition            = "VAR"
)

var (
	sortedCodes = []string{"ABR", "ACT", "ADP", "ALT", "ANN", "BLL", "BRL", "CMB", "CRI", "CSP", "DGO", "ENH", "ENL", "EXP", "FAC", "FST", "ILL", "LTE", "MCP", "MDT", "MLL", "NED", "NUM", "PRB", "REV", "SCH", "SMP", "SPE", "STU", "TCH", "UBR", "ULP", "UXP", "VAR"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		AbridgedEdition:            {"Abridged edition", "Content has been shortened: use for abridged, shortened, concise, condensed.", "Forkortet utgave", "Innholdet har blitt forkortet. "},
		ActingEdtion:               {"Acting edtion", "Version of a play or script intended for use of those directly involved in a production, usually including full stage directions in addition to the text of the script.", "Sceneutgave", "Versjon av et skuespill eller manuskript som er ment for bruk av de som er direkte involvert i en produksjon. Inkluderer vanligvis sceneanvisninger i tillegg til selve teksten."},
		AdaptedEdition:             {"Adapted edition", "Content has been adapted to serve a different purpose or audience, or from one medium to another: use for dramatization, novelization etc. Use <EditionStatement> to describe the exact nature of the adaptation.", "Tilpasset utgave", "Innholdet har blitt tilpasset for et annet formål, målgruppe, form eller fra et medium til et annet. Bruk for dramatisering etc. Bruk <EditionStatement> til å beskrive type tilpasning som er gjort."},
		Alternate:                  {"Alternate", "Do not use. This code is now DEPRECATED, but is retained in the list for reasons of backwards compatibility.", "Alternativ", "UTGÅTT. Ikke bruk denne koden. Den er med for å holde kompatiblitet med tidligere utgaver av ONIX kodelister"},
		AnnotatedEdition:           {"Annotated edition", "Content is augmented by the addition of notes.", "Annotert utgave", "Innholdet er utvidet med noter til teksten"},
		BilingualEdition:           {"Bilingual edition", "Both languages should be specified in the ‘Language’ group. Use MLL for an edition in more than two languages.", "Tospråklig utgave", "Begge språk skal spesifiseres i <Language> elementet.  Bruk MLL for en utgave med flere enn to språk"},
		BrailleEdition:             {"Braille edition", "Braille edition.", "Braille-utgave", "Braille-utgaver skal også ha tilhørende produktformkode."},
		CombinedVolume:             {"Combined volume", "An edition in which two or more works also published separately are combined in a single volume; AKA ‘omnibus’ edition.", "Sammensatt utgave", "Ei utgave hvor to titler som også er utgitt separate, er utgitt i et enkelt bind. Engelsk: omnibus"},
		CriticalEdition:            {"Critical edition", "Content includes critical commentary on the text.", "Kommentert utgave", "Inneholder kritiske kommentarer til teksten"},
		Coursepack:                 {"Coursepack", "Content was compiled for a specified educational course.", "Kurspakke", "Innholdet er sammensatt for et spesielt undervisningskurs"},
		DigitalOriginal:            {"Digital original", "A digital product which has no print counterpart and is not expected to have a print counterpart.", "Digital original", "Et digitalt produkt som ikke har og som heller ikke forventes å få en trykt versjon."},
		EnhancedEdition:            {"Enhanced edition", "Use for e-publications that have been enhanced with additional text, speech, other audio, video, interactive or other content.", "Beriket utgave", "Brukes for epublikasjoner som har blitt beriket med ekstra tekstmateriale, tale, annen lyd, video, interaktivitet eller annet innhold."},
		EnlargedEdition:            {"Enlarged edition", "Content has been enlarged or expanded from that of a previous edition.", "Utvidet utgave", "Innholdet er utvidet fra tildligere utgaver"},
		ExpurgatedEdition:          {"Expurgated edition", "‘Offensive’ content has been removed.", "Renset (expurgated) utgave", "Støtende innhold er fjernet. Sensurert utgave"},
		FacsimileEdition:           {"Facsimile edition", "Exact reproduction of the content and format of a previous edition.", "Faksimileutgave", "Eksakt reproduksjon av innhold og format fra tidligere utgave"},
		Festschrift:                {"Festschrift", "A collection of writings published in honor of a person, an institution or a society.", "Festskrift", "Ei samling tekster utgitt for å hedre en person, en institusjon eller en organisasjon."},
		IllustratedEdition:         {"Illustrated edition", "Content includes extensive illustrations which are not part of other editions.", "Illustrert utgave", "Innholdet inkluderer ekstra illustrasjoner som ikke er med i andre utgaver"},
		LargeTypeLargePrintEdition: {"Large type / large print edition", "Large print edition, print sizes 14 to 19 pt – see also ULP.", "Storskriftutgave", "Trykket med ekstra store bokstaver for svaksynte. Skal også følge med tilsvarende produktform kode fra liste 7"},
		MicroprintEdition:          {"Microprint edition", "A printed edition in a type size too small to be read without a magnifying glass.", "Mikroskriftutgave", "Trykket med bokstaver som er for små til å lese uten forstørrelsesglass. Skal også følge med tilsvarende produktform kode fra liste 7"},
		MediaTieIn:                 {"Media tie-in", "An edition published to coincide with the release of a film, TV program, or electronic game based on the same work. Use <EditionStatement> to describe the exact nature of the tie-in.", "Media tie-in", "En utgave publisert for å følge utgivelse av en film, TV-program eller et elektronisk spill basert på samme arbeid. Bruk <EditionStatement> til å beskrive type kobling til annet medium"},
		MultilingualEdition:        {"Multilingual edition", "All languages should be specified in the ‘Language’ group. Use BLL for a bilingual edition.", "Flerspråklig utgave", "Alle språk skal legges i <Language> elementet. Bruk BLL for en tospråklig utgave"},
		NewEdition:                 {"New edition", "Where no other information is given, or no other coded type is applicable.", "Ny utgave", "Brukes når ingen annen informasjon er gitt eller når ingen andre koder for utgavetype passer"},
		EditionWithNumberedCopies:  {"Edition with numbered copies", "A limited edition in which each copy is individually numbered.", "Nummerert utgave", "Begrenset opplag hvor hvert eksemplar er nummerert"},
		PreboundEdition:            {"Prebound edition", "In the US, a book that was previously bound, normally as a paperback, and has been rebound with a library-quality hardcover binding by a supplier other than the original publisher. See also the <Publisher> and <RelatedProduct> composites for other aspects of the treatment of prebound editions in ONIX.", "Prebound edition", "USA: ei bok som opprinnelig hadde ei innbinding, vanligvis heftet, og som har blitt bundet inn på nytt i forsterket biblioteksinnbinding. Se også <Publisher> og <RelatedProduct> for andre aspekter rundt dette"},
		RevisedEdition:             {"Revised edition", "Content has been revised from that of a previous edition.", "Revidert utgave", "Innholdet er revidert fra forrige utgave"},
		SchoolEdition:              {"School edition", "An edition intended specifically for use in schools.", "Skoleutgave", "En utgave spesielt beregnet for bruk i skolen."},
		SimplifiedLanguageEdition:  {"Simplified language edition", "An edition that uses simplified language (Finnish ‘Selkokirja’).", "Utgave med forenklet språk", "Utgave som bruker forenklet språk (finsk: selkokirja)"},
		SpecialEdition:             {"Special edition", "Use for anniversary, collectors’, de luxe, gift, limited, numbered, autographed edition. Use <EditionStatement> to describe the exact nature of the special edition.", "Spesialutgave", "Bruk for jubileum-, samler-, gave-, luksus, begrenset, nummerert, signert utgave. Bruk <EditionStatement> for å beskrive type spesialutgave"},
		StudentEdition:             {"Student edition", "Where a text is available in both student and teacher’s editions.", "Studentutgave", "Brukes når teksten er tilgjengelig i både en student- og en lærerutgave"},
		TeachersEdition:            {"Teacher’s edition", "Where a text is available in both student and teacher’s editions; use also for instructor’s or leader’s editions.", "Lærerutgave", "Brukes når teksten er tilgjengelig i både en student- og en lærerutgave"},
		UnabridgedEdition:          {"Unabridged edition", "Where a title has also been published in an abridged edition; also for audiobooks, regardless of whether an abridged audio version also exists.", "Uforkortet", "Brukes når tittelen også har vært publisert i en forkortet utgave"},
		UltraLargePrintEdition:     {"Ultra large print edition", "For print sizes 20pt and above, and with typefaces designed for the visually impaired – see also LTE.", "Utgave med ekstra stor skrift", "Til bruk for skriftstørrelse på 20 punkter og større og med skrifttyper velegnet for svaksynte - se også LTE"},
		UnexpurgatedEdition:        {"Unexpurgated edition", "Content previously considered ‘offensive’ has been restored.", "Urenset (unexpurgated) utgave", "Innhold som tidligere ble ansett som støtende er gjenopprettet. Usensurert utgave"},
		VariorumEdition:            {"Variorum edition", "Content includes notes by various commentators, and/or includes and compares several variant texts of the same work.", "Variorumutgave", "Variantutgave. Innholdet inkluderer noter av forskjellige kommentatorer og /eller inkluderer og sammenligner flere tekstvarianter av samme verket."},
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
