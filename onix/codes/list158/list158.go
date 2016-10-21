package list158 // Resource content type
import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Resource content type
const (
	FrontCover                = "01"
	BackCover                 = "02"
	CoverPack                 = "03"
	ContributorPicture        = "04"
	SeriesImageArtwork        = "05"
	SeriesLogo                = "06"
	ProductImageArtwork       = "07"
	ProductLogo               = "08"
	PublisherLogo             = "09"
	ImprintLogo               = "10"
	ContributorInterview      = "11"
	ContributorPresentation   = "12"
	ContributorReading        = "13"
	ContributorEventSchedule  = "14"
	SampleContent             = "15"
	Widget                    = "16"
	Review                    = "17"
	OtherCommentaryDiscussion = "18"
	ReadingGroupGuide         = "19"
	TeachersGuide             = "20"
	FeatureArticle            = "21"
	Characterinterview        = "22"
	WallpaperScreensaver      = "23"
	PressRelease              = "24"
	TableOfContents           = "25"
	Trailer                   = "26"
	CoverThumbnail            = "27"
	FullContent               = "28"
	FullCover                 = "29"
	MasterBrandLogo           = "30"
	PublishersCatalogue       = "34"
	OnlineAdvertismentPanel   = "35"
	OnlineAdvertismentPage    = "36"
	PromotionalEventMaterial  = "37"
	DigitalReviewCopy         = "38"
	License                   = "99"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "34", "35", "36", "37", "38", "99"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		FrontCover:               {"Front cover", "", "Omslagets forside", ""},
		BackCover:                {"Back cover", "", "Omslagets bakside", ""},
		CoverPack:                {"Cover / pack", "Not limited to front or back.", "Omslag / pakke", "Ikke begrenset til forside eller bakside"},
		ContributorPicture:       {"Contributor picture", "Photograph or portrait of contributor(s).", "Bilde av bidragsyter", "Fotograf eller portrett av bidragsyter(e)"},
		SeriesImageArtwork:       {"Series image / artwork", "", "Seriebilde / kunstverk", ""},
		SeriesLogo:               {"Series logo", "", "Serielogo", ""},
		ProductImageArtwork:      {"Product image / artwork", "", "Produktbilde / kunstverk", ""},
		ProductLogo:              {"Product logo", "", "Produktlogo", ""},
		PublisherLogo:            {"Publisher logo", "", "Vareeiers logo", ""},
		ImprintLogo:              {"Imprint logo", "", "Forlagets logo", ""},
		ContributorInterview:     {"Contributor interview", "", "Intervju med bidragsyter", ""},
		ContributorPresentation:  {"Contributor presentation", "Contributor presentation and/or commentary.", "Presentasjon av bidragsyter", "Presentasjon av og/eller kommentar om bidragsyter"},
		ContributorReading:       {"Contributor reading", "", "Bidragsyteren leser", ""},
		ContributorEventSchedule: {"Contributor event schedule", "Link to a schedule in iCalendar format.", "Bidragsyters kalender", "Lenke til timeplan i iCalendar-format."},
		SampleContent:            {"Sample content", "For example: sample chapter text, page images, screenshots.", "Innholdseksempel", "F.eks: kapitteltekst, bilder av sider, skjermdumper"},
		Widget:                   {"Widget", "A ‘look inside’ feature presented as a small embeddable application.", "Widget", "En bla i boka-funksjon, levert som en innbyggbar applikasjon"},
		Review:                   {"Review", "Use the <TextContent> composite for review quotes carried in the ONIX record. Use the <CitedContent> composite for a third-party review which is referenced from the ONIX record. Use <SupportingResource> only for a review which is offered for reproduction as part of promotional material for the product.", "Anmeldelse", "Bruk <TextContent> for sitater fra anmeldelser. Bruk <CitedContent> for anmeldelser hentet fra en tredjepart som det vises til i Onix-posten. Bruk <SupportingResource> kun for en anmeldelse som sendes kun for å gjenbrukes som en del av produktets markedsføringsmateriell."},
		OtherCommentaryDiscussion: {"Other commentary / discussion", "", "Andre kommentarer / diskusjoner", ""},
		ReadingGroupGuide:         {"Reading group guide", "", "Guide for lesesirkler", ""},
		TeachersGuide:             {"Teacher’s guide", "", "Lærerveiledning", ""},
		FeatureArticle:            {"Feature article", "Feature article provided by publisher.", "Featureartikkel", "Featureartikkel levert av vareeier"},
		Characterinterview:        {"Character ‘interview’", "Fictional character ‘interview’.", "Intervju med oppdiktet / skjønnlitterær person", "Intervju' med en skjønnlitterær karakter"},
		WallpaperScreensaver:      {"Wallpaper / screensaver", "", "Skjermsparer / bakgrunnsbilde", ""},
		PressRelease:              {"Press release", "", "Pressemelding", ""},
		TableOfContents:           {"Table of contents", "A table of contents held on a webpage, not in the ONIX record.", "Innholdsfortegnelse", "Innholdsfortegnelse som er tilgjengelig på ei nettside, ikke i selve Onix-posten"},
		Trailer:                   {"Trailer", "A promotional video, similar to a movie trailer (sometimes referred to as a ‘book trailer’).", "Trailer", "Video for markedsføring, ligner på en filmtrailer (kalles noen ganger ‘boktrailer’)."},
		CoverThumbnail:            {"Cover thumbnail", "Intended ONLY for transitional use, where ONIX 2.1 records referencing existing thumbnail assets of unknown pixel size are being re-expressed in ONIX 3.0. Use code 01 for all new cover assets, and where the pixel size of older assets is known.", "Miniatyrbilde (av omslaget)", "Intended ONLY for transitional use, where ONIX 2.1 records referencing existing thumbnail assets of unknown pixel size are being re-expressed in ONIX 3.0. Use code 01 for all new cover assets, and where the pixel size of older assets is known."},
		FullContent:               {"Full content", "The full content of the product (or the product itself), supplied for example to support full-text search.", "Alt innhold", "Alt innhold i produktet eller selve produktet. Leveres f.eks for å støtte fulltekstsøk"},
		FullCover:                 {"Full cover", "Includes cover, back cover, spine and – where appropriate – any flaps.", "Komplett omslag", "Inkluderer omslagets forside, bakside, rygg, og -der det fins - innbretter"},
		MasterBrandLogo:           {"Master brand logo", "", "Merkevarelogo", ""},
		PublishersCatalogue:       {"Publisher's catalogue", "For example an PDF or other digital representation of a publisher's 'new titles' or range catalogue", "Utgivers katalog", "F.eks. en PDF eller et annet digitalt format som inneholder en oversikt over utgiverens 'nye titler' eller en katalog over utgivelsene"},
		OnlineAdvertismentPanel:   {"Online advertisment panel", "For example a banner ad for the product. Pixel dimensions should typically be included in <ResourceVersionFeature>", "Online bannerannonse", "For example a banner ad for the product. Pixel dimensions should typically be included in <ResourceVersionFeature>"},
		OnlineAdvertismentPage:    {"Online advertisment page", "German 'Bühnenbild'", "Online reklameside", "Tysk 'Bühnenbild'"},
		PromotionalEventMaterial:  {"Promotional event material", "Eg posters, logos, banners, advertising templates for use in connection with a promotional event", "Reklamekampanjemateriell", "F.eks. plakater, logoer, bannere, reklamemaler for bruk i forbindelse med reklamekampanje"},
		DigitalReviewCopy:         {"Digital review copy", "Availability of a digital review or proof copy, may be limited to authorised users or account holders", "Digitalt anmeldereksemplar", "Digitalt anmeldereksemplar, kan være begrensninger på tilgjengelighet, slik at det ikke er tilgjengelig for alle."},
		License:                   {"License", "Link to a license covering permitted usage of the product content. This is a workaround in ONIX 3.0, pending introduction of an license link element near to UsageConstraints.", "Lisens", "Lenke til en lisens som dekker tillatt bruk av produktets innhold. Erstattes av et element for å sende lenker til lisenser nær UsageConstraints."},
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
