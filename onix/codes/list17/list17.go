package list17 // Contributor role code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Contributor role code
const (
	ByAuthor                         = "A01"
	With                             = "A02"
	ScreenplayBy                     = "A03"
	LibrettoBy                       = "A04"
	LyricsBy                         = "A05"
	ByComposer                       = "A06"
	ByArtist                         = "A07"
	ByPhotographer                   = "A08"
	CreatedBy                        = "A09"
	FromAnIdeaBy                     = "A10"
	DesignedBy                       = "A11"
	IllustratedBy                    = "A12"
	PhotographsBy                    = "A13"
	TextBy                           = "A14"
	PrefaceBy                        = "A15"
	PrologueBy                       = "A16"
	SummaryBy                        = "A17"
	SupplementBy                     = "A18"
	AfterwordBy                      = "A19"
	NotesBy                          = "A20"
	CommentariesBy                   = "A21"
	EpilogueBy                       = "A22"
	ForewordBy                       = "A23"
	IntroductionBy                   = "A24"
	FootnotesBy                      = "A25"
	MemoirBy                         = "A26"
	ExperimentsBy                    = "A27"
	IntroductionAndNotesBy           = "A29"
	SoftwareWrittenBy                = "A30"
	BookAndLyricsBy                  = "A31"
	ContributionsBy                  = "A32"
	AppendixBy                       = "A33"
	IndexBy                          = "A34"
	DrawingsBy                       = "A35"
	CoverDesignOrArtworkBy           = "A36"
	PreliminaryWorkBy                = "A37"
	OriginalAuthor                   = "A38"
	MapsBy                           = "A39"
	InkedOrColoredBy                 = "A40"
	PopUpsBy                         = "A41"
	ContinuedBy                      = "A42"
	Interviewer                      = "A43"
	Interviewee                      = "A44"
	OtherPrimaryCreator              = "A99"
	EditedBy                         = "B01"
	RevisedBy                        = "B02"
	RetoldBy                         = "B03"
	AbridgedBy                       = "B04"
	AdaptedBy                        = "B05"
	TranslatedBy                     = "B06"
	AsToldBy                         = "B07"
	TranslatedWithCommentaryBy       = "B08"
	SeriesEditedBy                   = "B09"
	EditedAndTranslatedBy            = "B10"
	EditorInChief                    = "B11"
	GuestEditor                      = "B12"
	VolumeEditor                     = "B13"
	EditorialBoardMember             = "B14"
	EditorialCoordinationBy          = "B15"
	ManagingEditor                   = "B16"
	FoundedBy                        = "B17"
	PreparedForPublicationBy         = "B18"
	AssociateEditor                  = "B19"
	ConsultantEditor                 = "B20"
	GeneralEditor                    = "B21"
	DramatizedBy                     = "B22"
	GeneralRapporteur                = "B23"
	LiteraryEditor                   = "B24"
	ArrangedByMusic                  = "B25"
	TechnicalEditor                  = "B26"
	ThesisAdvisorOrSupervisor        = "B27"
	ThesisExaminer                   = "B28"
	OtherAdaptationBy                = "B99"
	CompiledBy                       = "C01"
	SelectedBy                       = "C02"
	NonTextMaterialSelectedBy        = "C03"
	CuratedBy                        = "C04"
	OtherCompilationBy               = "C99"
	Producer                         = "D01"
	Director                         = "D02"
	Conductor                        = "D03"
	OtherDirectionBy                 = "D99"
	Actor                            = "E01"
	Dancer                           = "E02"
	Narrator                         = "E03"
	Commentator                      = "E04"
	VocalSoloist                     = "E05"
	InstrumentalSoloist              = "E06"
	ReadBy                           = "E07"
	PerformedByOrchestraBandEnsemble = "E08"
	Speaker                          = "E09"
	PerformedBy                      = "E99"
	FilmedPhotographedBy             = "F01"
	EditorFilmOrVideo                = "F02"
	OtherRecordingBy                 = "F99"
	AssistedBy                       = "Z01"
	Honored                          = "Z02"
	VariousRoles                     = "Z98"
	Other                            = "Z99"
)

var (
	sortedCodes = []string{"A01", "A02", "A03", "A04", "A05", "A06", "A07", "A08", "A09", "A10", "A11", "A12", "A13", "A14", "A15", "A16", "A17", "A18", "A19", "A20", "A21", "A22", "A23", "A24", "A25", "A26", "A27", "A29", "A30", "A31", "A32", "A33", "A34", "A35", "A36", "A37", "A38", "A39", "A40", "A41", "A42", "A43", "A44", "A99", "B01", "B02", "B03", "B04", "B05", "B06", "B07", "B08", "B09", "B10", "B11", "B12", "B13", "B14", "B15", "B16", "B17", "B18", "B19", "B20", "B21", "B22", "B23", "B24", "B25", "B26", "B27", "B28", "B99", "C01", "C02", "C03", "C04", "C99", "D01", "D02", "D03", "D99", "E01", "E02", "E03", "E04", "E05", "E06", "E07", "E08", "E09", "E99", "F01", "F02", "F99", "Z01", "Z02", "Z98", "Z99"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		ByAuthor:                   {"By (author)", "Author of a textual work.", "Av (forfatter)", "Forfatter av av teksten"},
		With:                       {"With", "With or as told to: ‘ghost’ author of a literary work.", "Med", "Sammen med eller fortalt til forfatter av et litterært verk"},
		ScreenplayBy:               {"Screenplay by", "Writer of screenplay or script (film or video).", "Manus av", "Forfatter av manus eller skript (film eller video)"},
		LibrettoBy:                 {"Libretto by", "Writer of libretto (opera): see also A31.", "Libretto av", "Forfatter av libretto (opera): se også A31"},
		LyricsBy:                   {"Lyrics by", "Author of lyrics (song): see also A31.", "Sangtekst av", "Forfatter av sangtekster: se også A31"},
		ByComposer:                 {"By (composer)", "Composer of music.", "Av (komponist)", "Komponist"},
		ByArtist:                   {"By (artist)", "Visual artist when named as the primary creator of, eg, a book of reproductions of artworks.", "Av (kunstner)", "Visuell artist når navngitt som primær produsent av f.eks. en bok med reproduksjoner av kunst."},
		ByPhotographer:             {"By (photographer)", "Photographer when named as the primary creator of, eg, a book of photographs.", "Av (fotograf)", "Fotograf når navngitt som primær produsent av f.eks. en bok med fotografier"},
		CreatedBy:                  {"Created by", "", "Skapt av", ""},
		FromAnIdeaBy:               {"From an idea by", "", "Etter ide av", ""},
		DesignedBy:                 {"Designed by", "", "Designet av", ""},
		IllustratedBy:              {"Illustrated by", "Artist when named as the creator of artwork which illustrates a text, or of the artwork of a graphic novel or comic book.", "Illustrert av", "Illustratør når den som har laget illustrasjonene til en tekst er navngitt"},
		PhotographsBy:              {"Photographs by", "Photographer when named as the creator of photographs which illustrate a text.", "Fotografier av", "Fotograf når den som har tatt fotografiene er navngitt i en bok med fotografier som illustrerer en tekst"},
		TextBy:                     {"Text by", "Author of text which accompanies art reproductions or photographs, or which is part of a graphic novel or comic book.", "Tekst av", "Forfatter av tekst som følger en reproduksjon av kunst eller fotografier"},
		PrefaceBy:                  {"Preface by", "Author of preface.", "Innledning av", "Forfatter av innledningen"},
		PrologueBy:                 {"Prologue by", "Author of prologue.", "Prolog av", "Forfatter av prologen"},
		SummaryBy:                  {"Summary by", "Author of summary.", "Sammendrag av", "Forfatter av sammendrag"},
		SupplementBy:               {"Supplement by", "Author of supplement.", "Tillegg av", "Forfatter av tillegg"},
		AfterwordBy:                {"Afterword by", "Author of afterword.", "Etterord av", "Forfatter av etterord"},
		NotesBy:                    {"Notes by", "Author of notes or annotations: see also A29.", "Noter ved", "Forfatter av noter eller merknader: se også A29"},
		CommentariesBy:             {"Commentaries by", "Author of commentaries on the main text.", "Kommentarer av", "Forfatter av kommentarer i hovedteksten"},
		EpilogueBy:                 {"Epilogue by", "Author of epilogue.", "Epilog av", "Forfatter av epilog"},
		ForewordBy:                 {"Foreword by", "Author of foreword.", "Forord av", "Forfatter av forord"},
		IntroductionBy:             {"Introduction by", "Author of introduction: see also A29.", "Introduksjon av", "Forfatter av introduksjon: se også A29"},
		FootnotesBy:                {"Footnotes by", "Author/compiler of footnotes.", "Fotnoter av", "Forfatter/redaktør av fotnoter"},
		MemoirBy:                   {"Memoir by", "Author of memoir accompanying main text.", "Biografi av", "Forfatter av biografi/memoarer som hører med hovedteksten"},
		ExperimentsBy:              {"Experiments by", "Person who carried out experiments reported in the text.", "Eksperimenter av", "Personen som utførte eksperimentene som er beskrevet i teksten"},
		IntroductionAndNotesBy:     {"Introduction and notes by", "Author of introduction and notes: see also A20 and A24.", "Introduksjon og noter av", "Forfatter av introduksjon og noter. Se også A20 og A24"},
		SoftwareWrittenBy:          {"Software written by", "Writer of computer programs ancillary to the text.", "Programvare utviklet av", "Den som har skrevet programkoden som er gjengitt sammen med eller i tillegg til teksten"},
		BookAndLyricsBy:            {"Book and lyrics by", "Author of the textual content of a musical drama: see also A04 and A05.", "Bok og sangtekst av", "Forfatter av teksten i et musikkdrama: se også A04 og A05"},
		ContributionsBy:            {"Contributions by", "Author of additional contributions to the text.", "Bidrag av", "Forfatter av ekstra bidrag til teksten"},
		AppendixBy:                 {"Appendix by", "Author of appendix.", "Appendiks av", "Forfatter av appendiks"},
		IndexBy:                    {"Index by", "Compiler of index.", "Indeks av", "Den som har laget indeks"},
		DrawingsBy:                 {"Drawings by", "", "Tegninger av", ""},
		CoverDesignOrArtworkBy:     {"Cover design or artwork by", "Use also for the cover artist of a graphic novel or comic book if named separately.", "Omslag av", "Use also for the cover artist of a graphic novel or comic book if named separately."},
		PreliminaryWorkBy:          {"Preliminary work by", "Responsible for preliminary work on which the work is based.", "Forarbeid av", "Ansvarlig for det innledende arbeid som denne teksten er basert på"},
		OriginalAuthor:             {"Original author", "Author of the first edition (usually of a standard work) who is not an author of the current edition.", "Opprinnelig forfatter", "Forfatter av først utgave, men som ikke er medforfatter av denne utgaven"},
		MapsBy:                     {"Maps by", "Maps drawn or otherwise contributed by.", "Kart av", "Kart tegnet eller frambrakt av"},
		InkedOrColoredBy:           {"Inked or colored by", "When separate persons are named as having respectively drawn and colored artwork, eg for a graphic novel or comic book, use A12 for ‘drawn by’ and A40 for ‘colored by’.", "Fargelagt av", "Nå forskjellig personer har respektive tegnet og fargelagt"},
		PopUpsBy:                   {"Pop-ups by", "Designer of pop-ups in a pop-up book, who may be different from the illustrator.", "Sprett-opp-effekter ved", "Designer av sprett-opp-effekter i ei bok, kan være en annen enn illustratøren"},
		ContinuedBy:                {"Continued by", "Use where a standard work is being continued by somebody other than the original author.", "Videreført av", "Brukes når et arbeid har blitt videreført av en annen enn opprinnelig forfatter"},
		Interviewer:                {"Interviewer", "", "Intervjuer", ""},
		Interviewee:                {"Interviewee", "", "Intervjuobjekt", ""},
		OtherPrimaryCreator:        {"Other primary creator", "Other type of primary creator not specified above.", "Annen hovedbidragsyter", "Annen type primært arbeid som ikke er spesifisert ovenfor"},
		EditedBy:                   {"Edited by", "", "Redigert av", ""},
		RevisedBy:                  {"Revised by", "", "Revidert av", ""},
		RetoldBy:                   {"Retold by", "", "Gjenfortalt av", ""},
		AbridgedBy:                 {"Abridged by", "", "Forkortet av", ""},
		AdaptedBy:                  {"Adapted by", "", "Tilpasset av", ""},
		TranslatedBy:               {"Translated by", "", "Oversatt av", ""},
		AsToldBy:                   {"As told by", "", "Som fortalt av", ""},
		TranslatedWithCommentaryBy: {"Translated with commentary by", "This code applies where a translator has provided a commentary on issues relating to the translation. If the translator has also provided a commentary on the work itself, codes B06 and A21 should be used.", "Oversatt med kommentarer av", "Denne koden brukes når oversetteren har bidratt med kommentarer på ting som vedrører oversettelsen. Hvis oversetteren har bidratt med kommentarer til selve teksten, skal B06 og A21 brukes"},
		SeriesEditedBy:             {"Series edited by", "Name of a series editor when the product belongs to a series.", "Serien redigert av", "Navn på seriens redaktør når et produkt inngår i en serie"},
		EditedAndTranslatedBy:      {"Edited and translated by", "", "Redigert og oversatt av", ""},
		EditorInChief:              {"Editor-in-chief", "", "Hovedredaktør", ""},
		GuestEditor:                {"Guest editor", "", "Gjesteredaktør", ""},
		VolumeEditor:               {"Volume editor", "", "Bindredaktør", ""},
		EditorialBoardMember:       {"Editorial board member", "", "Redaksjonsmedlem", ""},
		EditorialCoordinationBy:    {"Editorial coordination by", "", "Redaksjonskoordinator", ""},
		ManagingEditor:             {"Managing editor", "", "Sjefredaktør", ""},
		FoundedBy:                  {"Founded by", "Usually the founder editor of a serial publication: Begruendet von.", "Grunnlagt av", "Vanligvis den som grunnla en seriepublikasjon. "},
		PreparedForPublicationBy:   {"Prepared for publication by", "", "Tilrettelagt for publisering av", ""},
		AssociateEditor:            {"Associate editor", "", "Assisterende redaktør", ""},
		ConsultantEditor:           {"Consultant editor", "Use also for ‘advisory editor’.", "Konsulent/redaktør", "Brukes også for rådgivende redaktør"},
		GeneralEditor:              {"General editor", "", "Generell redaktør", ""},
		DramatizedBy:               {"Dramatized by", "", "Dramatisert av", ""},
		GeneralRapporteur:          {"General rapporteur", "In Europe, an expert editor who takes responsibility for the legal content of a collaborative law volume.", "Juridisk redaktør", "I Europa, en redaktør med ekspertise som har tatt ansvaret for det juridiske innholdet av en lovbok med flere forfattere"},
		LiteraryEditor:             {"Literary editor", "An editor who is responsible for establishing the text used in an edition of a literary work, where this is recognised as a distinctive role (in Spain, ‘editor literario’).", "Litterær redaktør", "En redaktør som er ansvarlig for teksten som er brukt i et litterært verk når dette er å betrakte som en separat rolle. (spansk: 'editor literario')"},
		ArrangedByMusic:            {"Arranged by (music)", "", "Arrangert av (musikk)", ""},
		TechnicalEditor:            {"Technical editor", "", "Teknisk redaktør", ""},
		ThesisAdvisorOrSupervisor:  {"Thesis advisor or supervisor", "", "Avhandling, veileder", ""},
		ThesisExaminer:             {"Thesis examiner", "", "Avhandling, opponent", ""},
		OtherAdaptationBy:          {"Other adaptation by", "Other type of adaptation or editing not specified above.", "Andre tilpasninger av", "Andre type tilpasninger eller redigering som ikke er spesifisert i andre koder"},
		CompiledBy:                 {"Compiled by", "", "Samlet av", ""},
		SelectedBy:                 {"Selected by", "", "Utvalgt av", ""},
		NonTextMaterialSelectedBy:  {"Non-text material selected by", "eg for a collection of selected photos etc.", "Ikke-tekstlig materiale utvalgt av", "F.eks. for ei samling av utvalgte fotografier osv."},
		CuratedBy:                  {"Curated by", "eg for an exhibition", "Kurator", "F.eks. for ei utstilling"},
		OtherCompilationBy:         {"Other compilation by", "Other type of compilation not specified above.", "Andre utvalg av", "Ander typer sammenstillinger ikke spesifisert i andre koder"},
		Producer:                   {"Producer", "", "Produsent", ""},
		Director:                   {"Director", "", "Regissør", ""},
		Conductor:                  {"Conductor", "Conductor of a musical performance.", "Dirigent", "Dirigent for en musikalsk framføring"},
		OtherDirectionBy:           {"Other direction by", "Other type of direction not specified above.", "Annen ledelse av", "Annen type ledelse ikke spesifisert i andre koder"},
		Actor:                      {"Actor", "", "Skuespiller", ""},
		Dancer:                     {"Dancer", "", "Danser", ""},
		Narrator:                   {"Narrator", "", "Forteller", ""},
		Commentator:                {"Commentator", "", "Kommentator", ""},
		VocalSoloist:               {"Vocal soloist", "Singer etc.", "Solist (vokal)", "Sanger etc."},
		InstrumentalSoloist:        {"Instrumental soloist", "", "Solist (instrumental)", ""},
		ReadBy:                     {"Read by", "Reader of recorded text, as in an audiobook.", "Lest av", "Den som leser teksten. F.eks. i en lydbok"},
		PerformedByOrchestraBandEnsemble: {"Performed by (orchestra, band, ensemble)", "Name of a musical group in a performing role.", "Framført av (orkester, band, ensemble)", "Navn på musikkgruppe som framfører "},
		Speaker:              {"Speaker", "Of a speech, lecture etc.", "Taler", "Navn på person som holder en tale, forelesning osv."},
		PerformedBy:          {"Performed by", "Other type of performer not specified above: use for a recorded performance which does not fit a category above, eg a performance by a stand-up comedian.", "Framført av", "Annen type utøver som ikke er spesifisert i andre koder: Brukes for en framførelse som ikke passer i annen kategori. F.eks. en framføring av en stand--up-komiker"},
		FilmedPhotographedBy: {"Filmed/photographed by", "Cinematographer, etc.", "Filmet/fotografert av", "Cinematographer, etc."},
		EditorFilmOrVideo:    {"Editor (film or video)", "", "Redaktør (film eller video)", ""},
		OtherRecordingBy:     {"Other recording by", "Other type of recording not specified above.", "Andre opptak av", "Andre typer opptak som ikke er spesifisert i andre koder"},
		AssistedBy:           {"Assisted by", "May be associated with any contributor role, and placement should therefore be controlled by contributor sequence numbering.", "Assistert av", "Kan bli sammenstilt med andre bidragsroller. Rekkefølgen bør derfor kontrolleres av \"contributor sequence numbering\""},
		Honored:              {"Honored/dedicated to", "", "Dedikert til", ""},
		VariousRoles:         {"(Various roles)", "For use ONLY with ‘et al’ or ‘Various’ within <UnnamedPersons>, where the roles of the multiple contributors vary.", "(Ulike roller)", "Skal brukes KUN med “et al” eller “Forskjellige personer” i <UnnamedPersons>, hvor rollene til de ulike bidragsyterne er forskjellige."},
		Other:                {"Other", "Other creative responsibility not falling within A to F above.", "Annet", "Andre kreative oppgaver som ikke faller innenfor de øvrige kodene"},
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
