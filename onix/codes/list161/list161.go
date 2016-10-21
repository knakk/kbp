package list161

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	LinkableResource      = "01"
	DownloadableFile      = "02"
	EmbeddableApplication = "03"
)

var (
	sortedCodes = []string{"01", "02", "03"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		LinkableResource:      {"Linkable resource", "A resource that may be accessed by a hyperlink. The current host (eg the ONIX sender, who may be the publisher) will provide ongoing hosting services for the resource for the active life of the product (or at least until the Until Date specified in <ContentDate>). The ONIX recipient may embed the URL in a consumer facing-website (eg as the src attribute in an <img> link), and need not host an independent copy of the resource.", "Lenkbar ressurs", "Ressurs som er tilgjengelig via en hyperlenke. Den som sender Onix-posten sørger for at ressursen er tilgjengelig så lenge produktet er aktivt (eller minimum til Til og med-dato (Until Date) angitt i <ContentDate>. Mottageren av Onix-posten kan inkludere URL-en på et nettsted for sine kunder (som src-attributt til en <img>-lenke), og trenger ikke lagre en egen kopi av ressursen."},
		DownloadableFile:      {"Downloadable file", "A file that may be downloaded on demand for third-party use. The ONIX sender will host a copy of the resource until the specified Until Date, but only for the ONIX recipient’s direct use. The ONIX recipient should download a copy of the resource, and must host an independent copy of the resource if it is used on a consumer-facing website. Special attention should be paid to the ‘Last Updated’ <ContentDate> to ensure the independent copy of the resource is kept up to date.", "Nedlastbar fil", "Fil som kan lastes ned for bruk av en tredjepart. Den som sender Onix-posten vil lagre en kopi av ressursen til den angitte Til og med-dato (Until date), men kun for Onix-mottagerens direkte bruk. Onix-mottakgeren bør laste ned en kopi av ressursen, og må lagre en egen kopi for bruk på egen nettside. Man bør se spesielt på ’Sist oppdatert’ <ContentDate> for å sikre at den nedlastede kopien er oppdatert."},
		EmbeddableApplication: {"Embeddable application", "An application which is supplied in a form which can be embedded into a third-party webpage. As type 02, except the resource contains active content such as JavaScript, Flash, etc.", "Applikasjon som kan inkluderes på nettsted", "En applikasjon som leveres i et format som gjør det mulig å inkludere den i en tredjeparts nettside. Som type 02, bortsett fra at ressursen inneholder aktivt innhold som JavaScript, Flash osv."},
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
