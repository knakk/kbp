package list176

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Android       = "01"
	BlackBerryOS  = "02"
	IOS           = "03"
	Linux         = "04"
	MacOS         = "05"
	MacOSX        = "06"
	PalmOS        = "07"
	WebOS         = "08"
	Symbian       = "09"
	Windows       = "10"
	WindowsCE     = "11"
	WindowsMobile = "12"
	MacOS         = "13"
	WindowsPhone7 = "14"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Android:       {"Android", "An Open Source mobile device operating system originally developed by Google and supported by the Open Handset Alliance.", "Android", "An Open Source mobile device operating system originally developed by Google and supported by the Open Handset Alliance."},
		BlackBerryOS:  {"BlackBerry OS", "A proprietary operating system supplied by Research In Motion for its BlackBerry handheld devices.", "BlackBerry OS", "A proprietary operating system supplied by Research In Motion for its BlackBerry handheld devices."},
		IOS:           {"iOS", "A proprietary operating system based on Mac OS X supplied by Apple for its iPhone, iPad and iPod Touch handheld devices.", "iOS", "A proprietary operating system based on Mac OS X supplied by Apple for its iPhone, iPad and iPod Touch handheld devices."},
		Linux:         {"Linux", "An operating system based on the Linux kernel.", "Linux", "An operating system based on the Linux kernel."},
		MacOS:         {"Mac OS", "[A proprietary operating system supplied by Apple on Macintosh computers up to 2002] DEPRECATED – use code 13 for all Mac OS versions.", "Mac OS", "[A proprietary operating system supplied by Apple on Macintosh computers up to 2002] DEPRECATED – use code 13 for all Mac OS versions."},
		MacOSX:        {"Mac OS X", "[A proprietary operating system supplied by Apple on Macintosh computers from 2001/2002] DEPRECATED – use code 13 for all Mac OS versions.", "Mac OS X", "[A proprietary operating system supplied by Apple on Macintosh computers from 2001/2002] DEPRECATED – use code 13 for all Mac OS versions."},
		PalmOS:        {"Palm OS", "A proprietary operating system (AKA Garnet OS) originally developed for handheld devices.", "Palm OS", "A proprietary operating system (AKA Garnet OS) originally developed for handheld devices."},
		WebOS:         {"webOS", "A proprietry Linux-based operating system for handheld devices, originally developed by Palm (now owned by HP).", "webOS", "A proprietry Linux-based operating system for handheld devices, originally developed by Palm (now owned by HP)."},
		Symbian:       {"Symbian", "An operating system for hand-held devices, originally developed as a proprietary system, but planned to become wholly Open Source by 2010.", "Symbian", "An operating system for hand-held devices, originally developed as a proprietary system, but planned to become wholly Open Source by 2010."},
		Windows:       {"Windows", "A proprietary operating system supplied by Microsoft.", "Windows", "A proprietary operating system supplied by Microsoft."},
		WindowsCE:     {"Windows CE", "A proprietary operating system (AKA Windows Embedded Compact, WinCE) supplied by Microsoft for small-scale devices.", "Windows CE", "A proprietary operating system (AKA Windows Embedded Compact, WinCE) supplied by Microsoft for small-scale devices."},
		WindowsMobile: {"Windows Mobile", "A proprietary operating system supplied by Microsoft for mobile devices.", "Windows Mobile", "A proprietary operating system supplied by Microsoft for mobile devices."},
		MacOS:         {"Mac OS", "A proprietary operating system supplied by Apple on Macintosh computers.", "Mac OS", "A proprietary operating system supplied by Apple on Macintosh computers."},
		WindowsPhone7: {"Windows Phone 7", "A proprietary operating system supplied by Microsoft for mobile devices, successor to Windows Mobile.", "Windows Phone 7", "A proprietary operating system supplied by Microsoft for mobile devices, successor to Windows Mobile."},
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
