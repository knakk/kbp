package list23

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	MainContentPageCount                  = "00"
	NumberOfWords                         = "02"
	FrontMatterPageCount                  = "03"
	BackMatterPageCount                   = "04"
	TotalNumberedPages                    = "05"
	ProductionPageCount                   = "06"
	AbsolutePageCount                     = "07"
	NumberOfPagesInPrintCounterpart       = "08"
	Duration                              = "09"
	NotionalNumberOfPagesInDigitalProduct = "10"
	ContentPageCount                      = "11"
	TotalUnnumberedInsertPageCount        = "12"
	DurationOfIntroductoryMatter          = "13"
	DurationOfMainContent                 = "14"
	Filesize                              = "22"
)

var (
	sortedCodes = []string{"00", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "22"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		MainContentPageCount:            {"Main content page count", "The highest-numbered page in a single numbered sequence of main content, usually the highest Arabic-numbered page in a book; or, for books without page numbers or (rarely) with multiple numbered sequences of main content, the total number of pages that carry the main content of the book. Note that this may include numbered but otherwise blank pages (eg pages inserted to ensure chapters start on a recto page) and may exclude unnumbered (but contentful) pages such as those in inserts/plate sections. It should exclude pages of back matter (eg any index) even when their numbering sequence continues from the main content. Either this or the Content Page count is the preferred page count for most books for the general reader. For books with substantial front and/or back matter, include also Front matter (03) and Back matter (04) page counts, or Total numbered pages (05). For books with inserts (plate sections), also include Total unnumbered insert page count whenever possible.", "Sidetall (bokas hovedinnhold)", "Sidetall i boka, vanligvis den høyest (arabisk-)nummererte sida i ei bok; eller, for bøker uten sidetall eller med flere nummereringer, det totale antall sider i boka. Merk at dette kan inkludere nummererte, men ellers blanke sider og kan ekskludere unummererte sider med innhold. Enten denne eller kode 11 er foretrukket for de fleste bøker beregnet for allmennmarkedet. "},
		NumberOfWords:                   {"Number of words", "Number of words of natural language text.", "Antall ord", "Antall ord i teksten"},
		FrontMatterPageCount:            {"Front matter page count", "The total number of numbered (usually Roman-numbered) pages that precede the main content of a book. This usually consists of various title and imprint pages, table of contents, an introduction, preface, foreword, etc.", "Sidetall for innledende stoff", "Antall sider (vanligvis romertall) for innledende sider før selve teksten i ei bok. Består vanligvis av ulike tittelsider, innholdsfortegnelse, introduksjon, forord osv."},
		BackMatterPageCount:             {"Back matter page count", "The total number of numbered (often Roman-numbered) pages that follow the main content of a book. This usually consists of an afterword, appendices, endnotes, index, etc. It excludes blank (or advertising) pages that are present only for convenience of printing and binding.", "Sidetall for avsluttende stoff", "Antall sider (vanligvis romertall) for sider som følger etter selve teksten i ei bok. Består vanligvis av etterord, appendikser, register osv."},
		TotalNumberedPages:              {"Total numbered pages", "The sum of all Roman- and Arabic-numbered pages. Note that this may include numbered but otherwise blank pages (eg pages inserted to ensure chapters start on a recto page) and may exclude unnumbered (but contentful) pages such as those in inserts/plate sections. It is the sum of the main content (00), front matter (03) and back matter (04) page counts.", "Totalt antall nummererte sider", "Summen av alle sider nummerert med romerske og arabiske tall. Merk at dette kan inkludere nummererte, men blankee sider og kan ekskludere unummererte sider med innhold. Det er summen av sidetall angitt i kodene 00, 03 og 04."},
		ProductionPageCount:             {"Production page count", "The total number of pages in a book, including unnumbered pages, front matter, back matter, etc. This includes any blank pages at the back that carry no content and are present only for convenience of printing and binding.", "Sidetall for produksjon", "Totalt antall sider i ei bok, inkludert unummererte sider, for- og etterord osv. Dette inkludere blanke sider som er med kun pga trykk og innbinding."},
		AbsolutePageCount:               {"Absolute page count", "The total number of pages of the book counting the cover as page 1. This page count type should be used only for digital publications delivered with fixed pagination.", "Fast sidetall (digitale publikasjoner)", "Totalt antall sider i boka, hvor man teller omslaget som side 1. Bør brukes kun for digitale publikasjoner som har fast sideinndeling."},
		NumberOfPagesInPrintCounterpart: {"Number of pages in print counterpart", "The total number of pages (equivalent to the Content page count) in the print counterpart of a digital product delivered without fixed pagination.", "Sidetall i tilsvarende trykt utgave (for digitale publikasjoner)", "Sidetall (tilsvarer kode 11) i trykt utgave av et digitalt produkt som ikke har fast sideinndeling."},
		Duration:                        {"Duration", "Total duration in time, expressed in the specified extent unit. This is the ‘running time’ equivalent of codes 05 or 11.", "Varighet/spilletid", "Brukes for å angi hvor lang tid avspillingen tar"},
		NotionalNumberOfPagesInDigitalProduct: {"Notional number of pages in digital product", "An estimate of the number of ‘pages’ in a digital product delivered without fixed pagination, and with no print counterpart, given as an indication of the size of the work. Equivalent to code 08, but for exclusively digital products.", "Anslått sidetall i digital publikasjon", "Et anslag over antall \"sider\" i et digitalt produkt som ikke har fast sideinndeling, og som ikke har ei trykt utgave. Angis som et anslag over verkets størrelse. Lik kode 08, men brukes kun for digitale produkter"},
		ContentPageCount:                      {"Content page count", "The sum of all Roman- and Arabic-numbered and contentful unnumbered pages. Sum of page counts with codes 00, 03, 04 and 12, and also the sum of 05 and 12.", "Sidetall", "Summen av alle nummererte (romersk og arabisk) og unummererte sider med innhold. Summen av sidetall angitt i kodene 00, 03, 04 og 12, og også summen av 05 og 12."},
		TotalUnnumberedInsertPageCount:        {"Total unnumbered insert page count", "The total number of unnumbered pages with content inserted within the main content of a book – for example inserts/plate sections that are not numbered.", "Totalt antall unummererte sider", "Totalt antall unummererte sider med innhold som er plassert i bokas hovedinnhold. F.eks. bolker med illustrasjoner som ikke er nummerert."},
		DurationOfIntroductoryMatter:          {"Duration of introductory matter", "Duration in time, expressed in the specified extent units, of introductory matter. This is the ‘running time’ equivalent of code 03, and comprises any significant amount of running time represented by announcements, titles, introduction or other material prefacing the main content.", "Spilletid for innledning", "Tilsvarer kode 03 for trykt materiale og omfatter spilletid for kunngjøringer, titler, introduksjon eller annet materiale som kommer før hovedinnholdet."},
		DurationOfMainContent:                 {"Duration of main content", "Duration in time, expressed in the specified extent units, of the main content. This is the ‘running time’ equivalent of code 00, and excludes time represented by announcements, titles, introduction or other prefatory material or ‘back matter’.", "Spilletid for hovedinnhold", "Tilsvarer kode 00 for trykt materiale, og inkluderer ikke spilletid for kunngjøringer, titler, introduksjon eller annet materiale som kommer før eller etter hovedinnholdet."},
		Filesize:                              {"Filesize", "The size of a digital file, expressed in the specified extent unit.", "Filstørrelse", "Filstørrelsen til et digitalt produkt"},
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
