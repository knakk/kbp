package list55

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	YYYYMMDD         = "00"
	YYYYMM           = "01"
	YYYYWW           = "02"
	YYYYQ            = "03"
	YYYYS            = "04"
	YYYY             = "05"
	YYYYMMDDYYYYMMDD = "06"
	YYYYMMYYYYMM     = "07"
	YYYYWWYYYYWW     = "08"
	YYYYQYYYYQ       = "09"
	YYYYSYYYYS       = "10"
	YYYYYYYY         = "11"
	TextString       = "12"
	YYYYMMDDThhmm    = "13"
	YYYYMMDDThhmmss  = "14"
	YYYYMMDDH        = "20"
	YYYYMMH          = "21"
	YYYYH            = "25"
	TextStringH      = "32"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "20", "21", "25", "32"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		YYYYMMDD:         {"YYYYMMDD", "Year month day (default).", "YYYYMMDD", "År måned dag (standardformat)"},
		YYYYMM:           {"YYYYMM", "Year and month.", "YYYYMM", "År og måned"},
		YYYYWW:           {"YYYYWW", "Year and week number.", "YYYYWW", "År og ukenummer"},
		YYYYQ:            {"YYYYQ", "Year and quarter (Q = 1, 2, 3, 4).", "YYYYQ", "År og kvartal (Q = 1, 2, 3, 4)"},
		YYYYS:            {"YYYYS", "Year and season (S = 1, 2, 3, 4, with 1 = “Spring”).", "YYYYS", "År og årstid (S = 1, 2, 3, 4, med 1 = vår)"},
		YYYY:             {"YYYY", "Year.", "YYYY", "Kun årstall"},
		YYYYMMDDYYYYMMDD: {"YYYYMMDDYYYYMMDD", "Spread of exact dates.", "YYYYMMDDYYYYMMDD", "Spenn mellom eksakte datoer"},
		YYYYMMYYYYMM:     {"YYYYMMYYYYMM", "Spread of months.", "YYYYMMYYYYMM", "Spenn mellom måneder"},
		YYYYWWYYYYWW:     {"YYYYWWYYYYWW", "Spread of week numbers.", "YYYYWWYYYYWW", "Spenn mellom ukenummer"},
		YYYYQYYYYQ:       {"YYYYQYYYYQ", "Spread of quarters.", "YYYYQYYYYQ ", "Spenn mellom kvartaler"},
		YYYYSYYYYS:       {"YYYYSYYYYS", "Spread of seasons.", "YYYYQYYYYQ ", "Spenn mellom sesonger"},
		YYYYYYYY:         {"YYYYYYYY", "Spread of years.", "YYYYYYYY", "Spenn mellom år"},
		TextString:       {"Text string", "For complex, approximate or uncertain dates.", "Tekststreng", "For komplekse eller usikre datoer"},
		YYYYMMDDThhmm:    {"YYYYMMDDThhmm", "Exact time. Use ONLY when exact times with hour/minute precision are relevant. By default, time is local. Alternatively, the time may be suffixed with an optional ‘Z’ for UTC times, or with ‘+’ or ‘-’ and an hhmm timezone offset from UTC. Times without a timezone are ‘rolling’ local times, times qualified with a timezone (using Z, + or -) specify a particular instant in time.", "YYYYMMDDThhmm", "Eksakt tid. Brukes BARE når det er nødvendig med presisjonsnivå på timer og minutter. Standard er at man angir lokal tid. Alternativt kan man legge til det valgfrie ’Z’ for UTC, eller bruke  ’+’ eller ’-’ og en tidsangivelse i hhmm offset fra UTC. Tid uten tidssone er ’rullende’ lokale tidsangiveiser, tidspunkter angitt med tidssone (ved å bruke Z, + eller -) angir et spesifikt tidspunkt."},
		YYYYMMDDThhmmss:  {"YYYYMMDDThhmmss", "Exact time. Use ONLY when exact times with second precision are relevant. By default, time is local. Alternatively, the time may be suffixed with an optional ‘Z’ for UTC times, or with ‘+’ or ‘-’ and an hhmm timezone offset from UTC. Times without a timezone are ‘rolling’ local times, times qualified with a timezone (using Z, + or -) specify a particular instant in time.", "YYYYMMDDThhmmss", "Eksakt tid. Brukes BARE når det er nødvendig med presisjonsnivå på sekunder. Standard er at man angir lokal tid. Alternativt kan man legge til det valgfrie ’Z’ for UTC, eller bruke  ’+’ eller ’-’ og en tidsangivelse i hhmm offset fra UTC. Tid uten tidssone er ’rullende’ lokale tidsangiveiser, tidspunkter angitt med tidssone (ved å bruke Z, + eller -) angir et spesifikt tidspunkt."},
		YYYYMMDDH:        {"YYYYMMDD (H)", "Year month day (Hijri calendar).", "YYYYMMDD (H)", "År måned dag (Islamsk kalender - Hirji)."},
		YYYYMMH:          {"YYYYMM (H)", "Year and month (Hijri calendar).", "YYYYMM (H)", "År og måned (Islamsk kalender - Hirji))."},
		YYYYH:            {"YYYY (H)", "Year (Hijri calendar).", "YYYY (H)", "År (Islamsk kalender - Hirji))."},
		TextStringH:      {"Text string (H)", "For complex, approximate or uncertain dates (Hijri calendar), text would usually be in Arabic script.", "Tekststreng (H)", "For komplekse eller usikre datoer (Islamsk kalender - Hirji), teksten er vanligvis angitt med arabiske tegn."},
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
