package list162

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	FileFormat                             = "01"
	ImageHeightInPixels                    = "02"
	ImageWidthInPixels                     = "03"
	Filename                               = "04"
	ApproximateDownloadFileSizeInMegabytes = "05"
	MD5HashValue                           = "06"
	ExactDownloadFileSizeInBytes           = "07"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		FileFormat:          {"File format", "Resource Version Feature Value carries a code from List 178.", "Filformat", "Resource Version Feature Value inneholder kode fra liste 178."},
		ImageHeightInPixels: {"Image height in pixels", "Resource Version Feature Value carries an integer.", "Bilde: høyde i piksler", "Resource Version Feature Value inneholder et heltall"},
		ImageWidthInPixels:  {"Image width in pixels", "Resource Version Feature Value carries an integer.", "Bilde: bredde i piksler", "Resource Version Feature Value inneholder et heltall"},
		Filename:            {"Filename", "Resource Version Feature Value carries the filename of the supporting resource.", "Filnavn", "Resource Version Feature Value inneholder ressursens filnavn"},
		ApproximateDownloadFileSizeInMegabytes: {"Approximate download file size in megabytes", "Resource Version Feature Value carries a decimal number only, suggested no more than 2 significant digits (eg 1.7, not 1.7462).", "Nedlastbar fil: omtrentlig størrelse i megabyte", "Resource Version Feature Value inneholder et desimaltall, helst ikke med mer en to sifre (f.eks. 1.7, ikke 1.7462)."},
		MD5HashValue:                           {"MD5 hash value", "MD5 hash value of the resource file. <ResourceVersionFeatureValue> should contain the hash value (as 32 hexadecimal digits). Can be used as a cryptographic check on the integrity of a resource after it has been retrieved.", "MD5 hash value", "MD5 hash value of the resource file. <ResourceVersionFeatureValue> should contain the hash value (as 32 hexadecimal digits). Can be used as a cryptographic check on the integrity of a resource after it has been retrieved."},
		ExactDownloadFileSizeInBytes:           {"Exact download file size in bytes", "Resource Version Feature Value carries a integer number only (eg 1831023).", "Nedlastbar fil: eksakt størrelse i bytes", "Resource Version Feature Value inneholder et heltall (f.eks. 1831023)."},
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
