package list81

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Audiobook                                  = "01"
	PerformanceSpokenWord                      = "02"
	MusicRecording                             = "03"
	OtherAudio                                 = "04"
	GamePuzzle                                 = "05"
	Video                                      = "06"
	StillImagesGraphics                        = "07"
	Software                                   = "08"
	Data                                       = "09"
	TextEyeReadable                            = "10"
	MusicalNotation                            = "11"
	MapsAndOrOtherCartographicContent          = "12"
	OtherSpeechContent                         = "13"
	ExtensiveLinksToExternalContent            = "14"
	ExtensiveLinksBetweenInternalContent       = "15"
	AdditionalEyeReadableTextNotPartOfMainWork = "16"
	PromotionalTextForOtherBookProduct         = "17"
	Photographs                                = "18"
	FiguresDiagramsChartsGraphs                = "19"
	AdditionalImagesGraphicsNotPartOfMainWork  = "20"
	PartialPerformanceSpokenWord               = "21"
	AdditionalAudioContentNotPartOfMainWork    = "22"
	PromotionalAudioForOtherBookProduct        = "23"
	AnimatedInteractiveIllustrations           = "24"
	NarrativeAnimation                         = "25"
	VideoRecordingOfAReading                   = "26"
	PerformanceVisual                          = "27"
	OtherVideo                                 = "28"
	PartialPerformanceVideo                    = "29"
	AdditionalVideoContentNotPartOfMainWork    = "30"
	PromotionalVideoForOtherBookProduct        = "31"
	Contest                                    = "32"
	DataSetPlusSoftware                        = "33"
	BlankPages                                 = "34"
	AdvertisingContent                         = "35"
	AdvertisingCoupons                         = "36"
	AdvertisingFirstParty                      = "37"
	AdvertisingThirdPartyDisplay               = "38"
	AdvertisingThirdPartyTextual               = "39"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Audiobook:                                  {"Audiobook", "Audio recording of a reading of a book or other text.", "Lydbok", "Lydopptak av opplesning av ei bok eller en annen tekst"},
		PerformanceSpokenWord:                      {"Performance – spoken word", "Audio recording of a drama or other spoken word performance.", "Framføring - tale", "Lydopptak av en framføring med muntlig tale"},
		MusicRecording:                             {"Music recording", "Audio recording of a music performance, including musical drama and opera.", "Musikkinnspilling", "Lydopptak av en musikkframføring, inkludert musikalsk drama og opera"},
		OtherAudio:                                 {"Other audio", "Audio recording of other sound, eg birdsong.", "Annen lyd", "Lydopptak av annen lyd, f.eks. fuglesang"},
		GamePuzzle:                                 {"Game / Puzzle", "No multi-user functionality. Formerly just ‘Game’.", "Spill / Hjernetrim", "Ingen flerbrukerfunksjonalitet. Tidligere kun 'Spill'"},
		Video:                                      {"Video", "Includes Film, video, animation etc. Use only when no more detailed specification is provided. Formerly ‘Moving images’.", "Video", "Film, video, animasjon etc.  Brukes kun når det ikke fins en mer detaljert spesifikasjon. Tidligere ‘Moving images’."},
		StillImagesGraphics:                        {"Still images / graphics", "Use only when no more detailed specification is provided.", "Stillbilder / grafikk", "Brukes kun når det ikke fins en mer detaljert spesifikasjon. "},
		Software:                                   {"Software", "Largely ‘content free’.", "Programvare", "Largely ‘content free’."},
		Data:                                       {"Data", "Data files.", "Datafiler", "Data files."},
		TextEyeReadable:                            {"Text (eye-readable)", "Readable text of the main work: this value is required, together with applicable <ProductForm> and <ProductFormDetail> values, to designate an e-book or other digital product whose primary content is eye-readable text.", "Tekst (lesbar)", "Lesbar tekst. Koden må sendes sammen med <ProductForm> og <ProductFormDetail> for beskrive ei e-bok eller et annet digital produkt som har vanlig lesbar tekst som sitt hovedinnhold."},
		MusicalNotation:                            {"Musical notation", "", "Musikknotasjon", ""},
		MapsAndOrOtherCartographicContent:          {"Maps and/or other cartographic content", "", "Kart og/eller annet kartografisk materiale", ""},
		OtherSpeechContent:                         {"Other speech content", "eg an interview, not a ‘reading’ or ‘performance’).", "Annet lydinnhold, tale", "f.eks. intervju, ikke opplesning eller framføring"},
		ExtensiveLinksToExternalContent:            {"Extensive links to external content", "E-publication is enhanced with a significant number of actionable (clickable) web links.", "Lenker til eksternt innhold", "Epublikasjonen er beriket med et vesentlig antall lenker til innhold på nettet."},
		ExtensiveLinksBetweenInternalContent:       {"Extensive links between internal content", "E-publication is enhanced with a significant number of actionable cross-references, hyperlinked notes and annotations, or with other links between largely textual elements (eg quiz/test questions, ‘choose your own ending’ etc).", "Lenker mellom innhold i produktet", "Epublikasjonen er beriket med et vesentlig antall kryssreferanser, lenker og annotasjoner, eller med andre lenker mellom (i hovedsak) tekstelementer (som oftest spørsmål og tester, 'velg din egen slutt' etc)."},
		AdditionalEyeReadableTextNotPartOfMainWork: {"Additional eye-readable text not part of main work", "E-publication is enhanced with additional textual content such as interview, feature article, essay, bibliography, quiz/test, other background material or text that is not included in a primary or ‘unenhanced’ version.", "Ekstra lesbar tekst som ikke er en del av verkets hoveddel", "Epublikasjonen er beriket med ytterligere tekstinnhold, for eksempel intervjuer, artikler, essay, bibliografi, quiz, tester,  annet bakgrunnsmateriale eller tekst som ikke er en del av den versjonen som ikke er beriket"},
		PromotionalTextForOtherBookProduct:         {"Promotional text for other book product", "eg Teaser chapter.", "Markedsføringstekst for annet bokprodukt", "F.eks. et kapittel fra annen bok"},
		Photographs:                                {"Photographs", "Whether in a plate section / insert, or not.", "Fotografier", "Enten de er inkludert som en egen del av boka eller ikke"},
		FiguresDiagramsChartsGraphs:                {"Figures, diagrams, charts, graphs", "Including other ‘mechanical’ (ie non-photographic) illustrations.", "Figurer, diagrammer, grafer", "Inkluderer andre 'mekaniske' illutrasjoner (dvs. illustrasjoner som ikke er fotografier)"},
		AdditionalImagesGraphicsNotPartOfMainWork:  {"Additional images / graphics not part of main work", "E-publication is enhanced with additional images or graphical content such as supplementary photographs that are not included in a primary or ‘unenhanced’ version.", "Ytterligere bildemateriale som ikke er del av hovedverket", "Epublikasjon som er beriket med illustrasjonsinnhold som for eksempel ekstra fotografier som ikke er med i den versjonen av produktet som ikke er beriket"},
		PartialPerformanceSpokenWord:               {"Partial performance – spoken word", "Audio recording of a reading, performance or dramatization of part of the work.", "Framføring av deler av verket, tale", "Lydopptak av opplesning, framføring eller dramatisering av en del av verket"},
		AdditionalAudioContentNotPartOfMainWork:    {"Additional audio content not part of main work", "Product is enhanced with audio recording of full or partial reading, performance, dramatization, interview, background documentary or other audio content not included in the primary or ‘unenhanced’ version.", "Ekstra lydinnhold som ikke er en del av hovedverket", "Produktet er beriket med lydopptak som kan være opplesning, framføring, dramatisering, intervju, bakgrunnsdokumentar eller annet lydinnhold som ikke er en del av den versjonen av verket som ikke er beriket."},
		PromotionalAudioForOtherBookProduct:        {"Promotional audio for other book product", "eg Reading of teaser chapter.", "Lydopptak, markedsføringsmateriale for annet bokprodukt", "F.eks. opplesning av et kapittel fra annen bok"},
		AnimatedInteractiveIllustrations:           {"Animated / interactive illustrations", "eg animated diagrams, charts, graphs or other illustrations.", "Animasjon / interaktive illustrasjoner", "F.eks. animerte diagrammer, grafer eller illustrasjoner"},
		NarrativeAnimation:                         {"Narrative animation", "eg cartoon, animatic or CGI animation.", "Fortellende animasjon", "F.eks. tegneserie, dataanimasjon"},
		VideoRecordingOfAReading:                   {"Video recording of a reading", "", "Videoopptak av opplesning", ""},
		PerformanceVisual:                          {"Performance – visual", "Video recording of a drama or other performance, including musical performance.", "Framføring, visuell", "Videoopptak av et skuespill eller annen framføring, inkluderer musikkframføring"},
		OtherVideo:                                 {"Other video", "Other video content eg interview, not a reading or performance.", "Annen video", "Annet videoinnhold, som oftest intervju, ikke opplesning eller framføring"},
		PartialPerformanceVideo:                    {"Partial performance – video", "Video recording of a reading, performance or dramatization of part of the work.", "Framføring av deler av verket, video", "Videoopptak av opplesning, framføring eller dramatisering av en del av verket"},
		AdditionalVideoContentNotPartOfMainWork:    {"Additional video content not part of main work", "E-publication is enhanced with video recording of full or partial reading, performance, dramatization, interview, background documentary or other content not included in the primary or ‘unenhanced’ version.", "Ekstra videoinnhold som ikke er en del av verkets hovedinnhold", "E-publication is enhanced with video recording of full or partial reading, performance, dramatization, interview, background documentary or other content not included in the primary or ‘unenhanced’ version."},
		PromotionalVideoForOtherBookProduct:        {"Promotional video for other book product", "eg Book trailer.", "Markedsføringsvideo for annet bokprodukt", "eng: Book trailer."},
		Contest:                      {"Contest", "Includes some degree of multi-user functionality.", "Konkurranse", "Inkluderer en viss grad av flerbrukerfunksjonalitet"},
		DataSetPlusSoftware:          {"Data set plus software", "", "Datasett pluss software", ""},
		BlankPages:                   {"Blank pages", "Intended to be filled in by the reader.", "Blanke sider", "Leseren skal fylle ut sidene selv"},
		AdvertisingContent:           {"Advertising content", "Use only where type of advertising content is not stated.", "Reklameinnhold", "Brukes kun når typen reklameinnhold ikke er spesifisert"},
		AdvertisingCoupons:           {"Advertising – coupons", "Eg to obtain discounts on other products.", "Reklame - kuponger", "Kuponger for å oppnå rabatter på andre produkter"},
		AdvertisingFirstParty:        {"Advertising – first party", "‘Back ads’ – promotional pages for other books (that do not include sample content, cf codes 17, 23).", "Reklame –  andre bøker", "Sider som markedsfører andre bøker (men som ikke inneholder smakebiter på innholdet, jf. kodene 17, 23)."},
		AdvertisingThirdPartyDisplay: {"Advertising – third party display", "", "Reklame – andre produkter (visuell)", ""},
		AdvertisingThirdPartyTextual: {"Advertising – third party textual", "", "Reklame - andre produkter (tekst)", ""},
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
