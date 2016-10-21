package list178

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	MP3          = "A103"
	WAV          = "A104"
	RealAudio    = "A105"
	WMA          = "A106"
	AAC          = "A107"
	AIFF         = "A111"
	RealVideo    = "D101"
	Quicktime    = "D102"
	AVI          = "D103"
	WMV          = "D104"
	MPEG4        = "D105"
	FLV          = "D106"
	SWF          = "D107"
	Format3GP    = "D108"
	WebM         = "D109"
	PDF          = "D401"
	GIF          = "D501"
	JPEG         = "D502"
	PNG          = "D503"
	TIFF         = "D504"
	EPUB         = "E101"
	HTML         = "E105"
	PDISO320001  = "E107"
	XHTML        = "E113"
	XPS          = "E115"
	AmazonKindle = "E116"
	CEB          = "E139"
	CEBX         = "E140"
)

var (
	sortedCodes = []string{"A103", "A104", "A105", "A106", "A107", "A111", "D101", "D102", "D103", "D104", "D105", "D106", "D107", "D108", "D109", "D401", "D501", "D502", "D503", "D504", "E101", "E105", "E107", "E113", "E115", "E116", "E139", "E140"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		MP3:          {"MP3", "MPEG 1/2 Audio Layer III file.", "MP3", "MPEG 1/2 Audio Layer III file."},
		WAV:          {"WAV", "Waveform Audio file.", "WAV", "Waveform Audio file."},
		RealAudio:    {"Real Audio", "Proprietary RealNetworks format.", "Real Audio", "Proprietary RealNetworks format."},
		WMA:          {"WMA", "Windows Media Audio format.", "WMA", "Windows media audio-format."},
		AAC:          {"AAC", "Advanced Audio Coding format.", "AAC", "Advanced Audio Coding format."},
		AIFF:         {"AIFF", "Audio Interchange File format.", "AIFF", "Audio Interchange File format."},
		RealVideo:    {"Real Video", "Proprietary RealNetworks format. Includes Real Video packaged within a .rm RealMedia container.", "Real Video-format", "Proprietary RealNetworks format. Includes Real Video packaged within a .rm RealMedia container."},
		Quicktime:    {"Quicktime", "Quicktime container format (.mov).", "Quicktime-format", "(.mov)."},
		AVI:          {"AVI", "Audio Video Interleave format.", "AVI-format", "Audio Video Interleave format."},
		WMV:          {"WMV", "Windows Media Video format.", "WMV", "Windows Media Video-format."},
		MPEG4:        {"MPEG-4", "MPEG-4 container format (.mp4, .m4a).", "MPEG-4", "MPEG-4 container format (.mp4, .m4a)."},
		FLV:          {"FLV", "Flash Video (.flv, .f4v).", "FLV", "Flash Video (.flv, .f4v)."},
		SWF:          {"SWF", "ShockWave (.swf).", "SWF", "ShockWave (.swf)."},
		Format3GP:    {"3GP", "3GPP container format (.3gp, .3g2).", "3GP", "3GPP container format (.3gp, .3g2)."},
		WebM:         {"WebM", "WebM container format (includes .mkv).", "WebM", "WebM container format (includes .mkv)."},
		PDF:          {"PDF", "Portable Document File format.", "PDF", "Portable Document File format."},
		GIF:          {"GIF", "Graphic Interchange File format.", "GIF", "Graphic Interchange File format."},
		JPEG:         {"JPEG", "Joint Photographic Experts Group format.", "JPEG", "Joint Photographic Experts Group format."},
		PNG:          {"PNG", "Portable Network Graphics format.", "PNG", "Portable Network Graphics format."},
		TIFF:         {"TIFF", "Tagged Image File format.", "TIFF", "Tagged Image File format."},
		EPUB:         {"EPUB", "The Open Publication Structure / OPS Container Format standard of the International Digital Publishing Forum (IDPF) [File extension .epub].", "EPUB", "The Open Publication Structure / OPS Container Format standard of the International Digital Publishing Forum (IDPF) [File extension .epub]."},
		HTML:         {"HTML", "HyperText Mark-up Language [File extension .html, .htm].", "HTML", "HyperText Mark-up Language [File extension .html, .htm]."},
		PDFISO320001: {"PDF", "Portable Document Format (ISO 32000-1:2008) [File extension .pdf].", "PDF", "Portable Document Format (ISO 32000-1:2008) [File extension .pdf]."},
		XHTML:        {"XHTML", "Extensible Hypertext Markup Language [File extension .xhtml, .xht, .xml, .html, .htm].", "XHTML", "Extensible Hypertext Markup Language [File extension .xhtml, .xht, .xml, .html, .htm]."},
		XPS:          {"XPS", "XML Paper Specification.", "XPS", "XML Paper Specification."},
		AmazonKindle: {"Amazon Kindle", "A format proprietary to Amazon for use with its Kindle reading devices or software readers [File extensions .azw, .mobi, .prc].", "Amazon Kindle", "A format proprietary to Amazon for use with its Kindle reading devices or software readers [File extensions .azw, .mobi, .prc]."},
		CEB:          {"CEB", "Founder Apabi’s proprietary basic e-book format.", "CEB", "Founder Apabi’s proprietary basic e-book format."},
		CEBX:         {"CEBX", "Founder Apabi’s proprietary XML e-book format.", "CEBX", "Founder Apabi’s proprietary XML e-book format."},
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
