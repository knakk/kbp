package list24

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Words                      = "02"
	Pages                      = "03"
	HoursIntegerAndDecimals    = "04"
	MinutesIntegerAndDecimals  = "05"
	SecondsIntegerOnly         = "06"
	Tracks                     = "11"
	HoursHHH                   = "14"
	HoursAndMinutesHHHMM       = "15"
	HoursMinutesSecondsHHHMMSS = "16"
	Bytes                      = "17"
	Kbytes                     = "18"
	Mbytes                     = "19"
)

var (
	sortedCodes = []string{"02", "03", "04", "05", "06", "11", "14", "15", "16", "17", "18", "19"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Words: {"Words", "Words of natural language text.", "Ord", "Ord i selve hovedteksten"},
		Pages: {"Pages", "", "Sider", ""},
		HoursIntegerAndDecimals:    {"Hours (integer and decimals)", "", "Timer (heltall og desimaler)", "Spilletid i timer"},
		MinutesIntegerAndDecimals:  {"Minutes (integer and decimals)", "", "Minutter (heltall og desimaler)", "Spilletid i minutter"},
		SecondsIntegerOnly:         {"Seconds (integer only)", "", "Sekunder (kun heltall)", "Spilletid i sekunder"},
		Tracks:                     {"Tracks", "Of an audiobook on CD (or a similarly divided selection of audio files). Conventionally, each track is 3–6 minutes of running time, and track counts are misleading and inappropriate if the average track duration is significantly more or less than this. Note that track breaks are not necessarily aligned with structural breaks in the text (eg chapter breaks).", "Spor", "For lydbok på CD (eller lydfiler som er delt inn på en lignende mate)."},
		HoursHHH:                   {"Hours HHH", "Fill with leading zeroes if any elements are missing.", "Timer HHH", "Angis med tre siffer"},
		HoursAndMinutesHHHMM:       {"Hours and minutes HHHMM", "Fill with leading zeroes if any elements are missing.", "Timer og minutter HHHMM", "Angis med fem siffer"},
		HoursMinutesSecondsHHHMMSS: {"Hours minutes seconds HHHMMSS", "Fill with leading zeroes if any elements are missing.", "Timer, minutter og sekunder HHHMMSS", "Angis med syv siffer"},
		Bytes:  {"Bytes", "", "Bytes", ""},
		Kbytes: {"Kbytes", "", "Kilobytes", ""},
		Mbytes: {"Mbytes", "", "Megabytes", ""},
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
