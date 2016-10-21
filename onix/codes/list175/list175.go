package list175 // Product form detail
import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Available codes for Product form detail:
const (
	CDStandardAudioFormat                       = "A101"
	SACDSuperAudioFormat                        = "A102"
	MP3Format                                   = "A103"
	WAVFormat                                   = "A104"
	RealAudioFormat                             = "A105"
	WMA                                         = "A106"
	AAC                                         = "A107"
	OggVorbis                                   = "A108"
	Audible                                     = "A109"
	FLAC                                        = "A110"
	AIFF                                        = "A111"
	ALAC                                        = "A112"
	DAISY2FullAudioWithTitleOnlyNoNavigation    = "A201"
	DAISY2FullAudioWithNavigationNoText         = "A202"
	DAISY2FullAudioWithNavigationAndPartialText = "A203"
	DAISY2FullAudioWithNavigationAndFullText    = "A204"
	DAISY2FullTextWithNavigationAndPartialAudio = "A205"
	DAISY2FullTextWithNavigationAndNoAudio      = "A206"
	DAISY3FullAudioWithTitleOnlyNoNavigation    = "A207"
	DAISY3FullAudioWithNavigationNoText         = "A208"
	DAISY3FullAudioWithNavigationAndPartialText = "A209"
	DAISY3FullAudioWithNavigationAndFullText    = "A210"
	DAISY3FullTextWithNavigationAndPartialAudio = "A211"
	DAISY3FullTextWithNavigationAndNoAudio      = "A212"
	StandaloneAudio                             = "A301"
	ReadalongAudio                              = "A302"
	PlayalongAudio                              = "A303"
	SpeakalongAudio                             = "A304"
	MassMarketRackPaperback                     = "B101"
	TradePaperbackUS                            = "B102"
	DigestFormatPaperback                       = "B103"
	AFormatPaperback                            = "B104"
	BFormatPaperback                            = "B105"
	TradePaperbackUK                            = "B106"
	TallRackPaperbackUS                         = "B107"
	A5SizeTankobon                              = "B108"
	JISB5SizeTankobon                           = "B109"
	JISB6SizeTankobon                           = "B110"
	A6SizeBunko                                 = "B111"
	B40DoriShinsho                              = "B112"
	PocketSweden                                = "B113"
	StorpocketSweden                            = "B114"
	KartonnageSweden                            = "B115"
	FlexbandSweden                              = "B116"
	Mook                                        = "B117"
	Dwarsligger                                 = "B118"
	Form46Size                                  = "B119"
	Form46HenkeiSize                            = "B120"
	A4                                          = "B121"
	A4HenkeiSize                                = "B122"
	A5HenkeiSize                                = "B123"
	B5HenkeiSize                                = "B124"
	B6HenkeiSize                                = "B125"
	ABSize                                      = "B126"
	JISB7Size                                   = "B127"
	KikuSize                                    = "B128"
	KikuHenkeiSize                              = "B129"
	JISB4Size                                   = "B130"
	PaperbackDE                                 = "B131"
	ColoringJoinTheDotBook                      = "B201"
	LiftTheFlapBook                             = "B202"
	MiniatureBook                               = "B204"
	MovingPictureFlickerBook                    = "B205"
	PopUpBook                                   = "B206"
	ScentedsmellyBook                           = "B207"
	SoundStorynoisyBook                         = "B208"
	StickerBook                                 = "B209"
	TouchAndFeelBook                            = "B210"
	DieCutBook                                  = "B212"
	BookAsToy                                   = "B213"
	SoftToTouchBook                             = "B214"
	FuzzyFeltBook                               = "B215"
	PictureBook                                 = "B221"
	CarouselBook                                = "B222"
	LooseLeafSheetsAndBinder                    = "B301"
	LooseLeafBinderOnly                         = "B302"
	LooseLeafSheetsOnly                         = "B303"
	Sewn                                        = "B304"
	UnsewnAdhesiveBound                         = "B305"
	LibraryBinding                              = "B306"
	ReinforcedBinding                           = "B307"
	HalfBound                                   = "B308"
	QuarterBound                                = "B309"
	SaddleSewn                                  = "B310"
	CombBound                                   = "B311"
	WireO                                       = "B312"
	ConcealedWire                               = "B313"
	ClothOverBoards                             = "B401"
	PaperOverBoards                             = "B402"
	LeatherReal                                 = "B403"
	LeatherImitation                            = "B404"
	LeatherBonded                               = "B405"
	Vellum                                      = "B406"
	Cloth                                       = "B409"
	ImitationCloth                              = "B410"
	Velvet                                      = "B411"
	FlexiblePlasticVinylCover                   = "B412"
	PlasticCovered                              = "B413"
	VinylCovered                                = "B414"
	LaminatedCover                              = "B415"
	WithDustJacket                              = "B501"
	WithPrintedDustJacket                       = "B502"
	WithTranslucentDustCover                    = "B503"
	WithFlaps                                   = "B504"
	WithThumbIndex                              = "B505"
	WithRibbonMarkerS                           = "B506"
	WithZipFastener                             = "B507"
	WithButtonSnapFastener                      = "B508"
	WithLeatherEdgeLining                       = "B509"
	RoughFront                                  = "B510"
	Foldout                                     = "B511"
	TurnAroundBook                              = "B601"
	UnflippedMangaFormat                        = "B602"
	Syllabification                             = "B610"
	UKUncontractedBraille                       = "B701"
	UKContractedBraille                         = "B702"
	USBraille                                   = "B703"
	USUncontractedBraille                       = "B704"
	USContractedBraille                         = "B705"
	UnifiedEnglishBraille                       = "B706"
	Moon                                        = "B707"
	RealVideoFormat                             = "D101"
	QuicktimeFormat                             = "D102"
	AVIFormat                                   = "D103"
	WindowsMediaVideoFormat                     = "D104"
	MPEG4                                       = "D105"
	MSDOS                                       = "D201"
	Windows                                     = "D202"
	Macintosh                                   = "D203"
	UNIXLINUX                                   = "D204"
	OtherOperatingSystemS                       = "D205"
	PalmOS                                      = "D206"
	WindowsMobile                               = "D207"
	MicrosoftXBox                               = "D301"
	NintendoGameboyColor                        = "D302"
	NintendoGameboyAdvanced                     = "D303"
	NintendoGameboy                             = "D304"
	NintendoGamecube                            = "D305"
	Nintendo64                                  = "D306"
	SegaDreamcast                               = "D307"
	SegaGenesisMegadrive                        = "D308"
	SegaSaturn                                  = "D309"
	SonyPlayStation1                            = "D310"
	SonyPlayStation2                            = "D311"
	NintendoDualScreen                          = "D312"
	SonyPlayStation3                            = "D313"
	Xbox360                                     = "D314"
	NintendoWii                                 = "D315"
	SonyPlayStationPortablePSP                  = "D316"
	Other                                       = "E100"
	EPUB                                        = "E101"
	OEB                                         = "E102"
	DOC                                         = "E103"
	DOCX                                        = "E104"
	HTML                                        = "E105"
	ODF                                         = "E106"
	PDF                                         = "E107"
	PDFA                                        = "E108"
	RTF                                         = "E109"
	SGML                                        = "E110"
	TCR                                         = "E111"
	TXT                                         = "E112"
	XHTML                                       = "E113"
	ZTXT                                        = "E114"
	XPS                                         = "E115"
	AmazonKindle                                = "E116"
	BBeB                                        = "E117"
	DXReader                                    = "E118"
	EBL                                         = "E119"
	Ebrary                                      = "E120"
	EReader                                     = "E121"
	Exebook                                     = "E122"
	FranklinEBookman                            = "E123"
	GemstarRocketbook                           = "E124"
	ISilo                                       = "E125"
	MicrosoftReader                             = "E126"
	Mobipocket                                  = "E127"
	MyiLibrary                                  = "E128"
	NetLibrary                                  = "E129"
	Plucker                                     = "E130"
	VitalBook                                   = "E131"
	Vook                                        = "E132"
	GoogleEdition                               = "E133"
	BookappForIOS                               = "E134"
	BookappForAndroid                           = "E135"
	BookappForOtherOperatingSystem              = "E136"
	CEB                                         = "E139"
	CEBX                                        = "E140"
	IBook                                       = "E141"
	EPIB                                        = "E142"
	SCORM                                       = "E143"
	EBP                                         = "E144"
	Reflowable                                  = "E200"
	FixedFormat                                 = "E201"
	ReadableOffline                             = "E202"
	RequiresNetworkConnection                   = "E203"
	ContentRemoved                              = "E204"
	Landscape                                   = "E210"
	Portrait                                    = "E211"
	Aspect54                                    = "E221"
	Aspect43                                    = "E222"
	Aspect32                                    = "E223"
	Aspect1610                                  = "E224"
	Aspect169                                   = "E225"
	Laminated                                   = "L101"
	DeskCalendar                                = "P101"
	MiniCalendar                                = "P102"
	EngagementCalendar                          = "P103"
	DayByDayCalendar                            = "P104"
	PosterCalendar                              = "P105"
	WallCalendar                                = "P106"
	PerpetualCalendar                           = "P107"
	AdventCalendar                              = "P108"
	BookmarkCalendar                            = "P109"
	StudentCalendar                             = "P110"
	ProjectCalendar                             = "P111"
	AlmanacCalendar                             = "P112"
	OtherCalendar                               = "P113"
	OtherCalendarOrOrganiserProduct             = "P114"
	HardbackStationery                          = "P201"
	PaperbackSoftbackStationery                 = "P202"
	SpiralBoundStationery                       = "P203"
	LeatherFineBindingStationery                = "P204"
	PAL                                         = "V201"
	NTSC                                        = "V202"
	SECAM                                       = "V203"
)

var (
	sortedCodes = []string{"A101", "A102", "A103", "A104", "A105", "A106", "A107", "A108", "A109", "A110", "A111", "A112", "A201", "A202", "A203", "A204", "A205", "A206", "A207", "A208", "A209", "A210", "A211", "A212", "A301", "A302", "A303", "A304", "B101", "B102", "B103", "B104", "B105", "B106", "B107", "B108", "B109", "B110", "B111", "B112", "B113", "B114", "B115", "B116", "B117", "B118", "B119", "B120", "B121", "B122", "B123", "B124", "B125", "B126", "B127", "B128", "B129", "B130", "B131", "B201", "B202", "B204", "B205", "B206", "B207", "B208", "B209", "B210", "B212", "B213", "B214", "B215", "B221", "B222", "B301", "B302", "B303", "B304", "B305", "B306", "B307", "B308", "B309", "B310", "B311", "B312", "B313", "B401", "B402", "B403", "B404", "B405", "B406", "B409", "B410", "B411", "B412", "B413", "B414", "B415", "B501", "B502", "B503", "B504", "B505", "B506", "B507", "B508", "B509", "B510", "B511", "B601", "B602", "B610", "B701", "B702", "B703", "B704", "B705", "B706", "B707", "D101", "D102", "D103", "D104", "D105", "D201", "D202", "D203", "D204", "D205", "D206", "D207", "D301", "D302", "D303", "D304", "D305", "D306", "D307", "D308", "D309", "D310", "D311", "D312", "D313", "D314", "D315", "D316", "E100", "E101", "E102", "E103", "E104", "E105", "E106", "E107", "E108", "E109", "E110", "E111", "E112", "E113", "E114", "E115", "E116", "E117", "E118", "E119", "E120", "E121", "E122", "E123", "E124", "E125", "E126", "E127", "E128", "E129", "E130", "E131", "E132", "E133", "E134", "E135", "E136", "E139", "E140", "E141", "E142", "E143", "E144", "E200", "E201", "E202", "E203", "E204", "E210", "E211", "E221", "E222", "E223", "E224", "E225", "L101", "P101", "P102", "P103", "P104", "P105", "P106", "P107", "P108", "P109", "P110", "P111", "P112", "P113", "P114", "P201", "P202", "P203", "P204", "V201", "V202", "V203"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		CDStandardAudioFormat: {"CD standard audio format", "CD ‘red book’ format.", "CD standard lydformat", "CD ‘red book’ format."},
		SACDSuperAudioFormat:  {"SACD super audio format", "", "SACD super audio format", ""},
		MP3Format:             {"MP3 format", "MPEG-1/2 Audio Layer III file.", "MP3 format", "MPEG-1/2 Audio Layer III file."},
		WAVFormat:             {"WAV format", "", "WAV format", ""},
		RealAudioFormat:       {"Real Audio format", "", "Real Audio format", ""},
		WMA:                   {"WMA", "Windows Media Audio format.", "WMA", "Windows Media Audio format."},
		AAC:                   {"AAC", "Advanced Audio Coding format.", "AAC", "Advanced Audio Coding format."},
		OggVorbis:             {"Ogg/Vorbis", "Vorbis audio format in the Ogg container.", "Ogg/Vorbis", "Vorbis lydformat i Ogg container"},
		Audible:               {"Audible", "Audio format proprietary to Audible.com.", "Audible", "Proprietært lydformat for Audible.com."},
		FLAC:                  {"FLAC", "Free lossless audio codec.", "FLAC", "Free lossless audio codec."},
		AIFF:                  {"AIFF", "Audio Interchangeable File Format.", "AIFF", "Audio Interchangeable File Format."},
		ALAC:                  {"ALAC", "Apple Lossless Audio Codec.", "ALAC", "Apple Lossless Audio Codec."},
		DAISY2FullAudioWithTitleOnlyNoNavigation:    {"DAISY 2: full audio with title only (no navigation)", "Deprecated, as does not meet DAISY 2 standard. Use conventional audiobook codes instead.", "DAISY 2: full audio with title only (no navigation)", "UTGÅTT"},
		DAISY2FullAudioWithNavigationNoText:         {"DAISY 2: full audio with navigation (no text)", "", "DAISY 2: full audio with navigation (no text)", ""},
		DAISY2FullAudioWithNavigationAndPartialText: {"DAISY 2: full audio with navigation and partial text", "", "DAISY 2: full audio with navigation and partial text", ""},
		DAISY2FullAudioWithNavigationAndFullText:    {"DAISY 2: full audio with navigation and full text", "", "DAISY 2: full audio with navigation and full text", ""},
		DAISY2FullTextWithNavigationAndPartialAudio: {"DAISY 2: full text with navigation and partial audio", "Reading systems may provide full audio via text-to-speech.", "DAISY 2: full text with navigation and partial audio", "Lesesystemer kan tilby fullstendig lyd via tekst-til-tale"},
		DAISY2FullTextWithNavigationAndNoAudio:      {"DAISY 2: full text with navigation and no audio", "Reading systems may provide full audio via text-to-speech.", "DAISY 2: full text with navigation and no audio", "Lesesystemer kan tilby fullstendig lyd via tekst-til-tale"},
		DAISY3FullAudioWithTitleOnlyNoNavigation:    {"DAISY 3: full audio with title only (no navigation)", "Deprecated, as does not meet DAISY 3 standard. Use conventional audiobook codes instead.", "DAISY 3: full audio with title only (no navigation)", "UTGÅTT"},
		DAISY3FullAudioWithNavigationNoText:         {"DAISY 3: full audio with navigation (no text)", "", "DAISY 3: full audio with navigation (no text)", ""},
		DAISY3FullAudioWithNavigationAndPartialText: {"DAISY 3: full audio with navigation and partial text", "", "DAISY 3: full audio with navigation and partial text", ""},
		DAISY3FullAudioWithNavigationAndFullText:    {"DAISY 3: full audio with navigation and full text", "", "DAISY 3: full audio with navigation and full text", ""},
		DAISY3FullTextWithNavigationAndPartialAudio: {"DAISY 3: full text with navigation and partial audio", "Reading systems may provide full audio via text-to-speech.", "DAISY 3: full text with navigation and partial audio", "Lesesystemer kan tilby fullstendig lyd via tekst-til-tale"},
		DAISY3FullTextWithNavigationAndNoAudio:      {"DAISY 3: full text with navigation and no audio", "Reading systems may provide full audio via text-to-speech.", "DAISY 3: full text with navigation and no audio", "Lesesystemer kan tilby fullstendig lyd via tekst-til-tale"},
		StandaloneAudio:                             {"Standalone audio", "", "Lydfil, frittstående", ""},
		ReadalongAudio:                              {"Readalong audio", "Audio intended exclusively for use alongside a printed copy of the book. Most often a children’s product. Normally contains instructions such as “turn the page now” and other references to the printed item, and is usually sold packaged together with a printed copy.", "Lydfil, opplesning", "Lyd ment for bruk sammen med trykt eksemplar av produktet. Som oftest produkter for barn. Inneholder vanligvis instruksjoner som \"bla om\" og andre henvisninger til det trykte produktet, og blir vanligvis solgt sammen med det trykte produktet"},
		PlayalongAudio:                              {"Playalong audio", "Audio intended for musical accompaniment, eg ‘Music minus one’, etc, often used for music learning. Includes singalong backing audio for musical learning or for Karaoke-style entertainment.", "Lydfil, musikk", "Lyd med musikalsk akkompagnement, ofte bruk for musikkopplæring. Inkluderer lydspor uten vokal for musikkopplæring eller underholdning à la karaoke"},
		SpeakalongAudio:                             {"Speakalong audio", "Audio intended for language learning, which includes speech plus gaps intended to be filled by the listener.", "Lydfil, tale", "Lydfil ment brukt for språkopplæring, og som inkluderer tale med pauser som skal fylles av lytteren"},
		MassMarketRackPaperback:                     {"Mass market (rack) paperback", "In North America, a category of paperback characterized partly by page size (typically 4¼ x 7 1/8 inches) and partly by target market and terms of trade. Use with Product Form code BC.", "Mass market (rack) paperback", "I Nord-Amerika, en kategori \"paperback\" karakterisert delvis på størrelse (4 1/4 x 7 1/8 inches) og delvis på markedet den er beregnet for, samt handelsbetingelsene, Brukes med \"Product Form\" kode BC"},
		TradePaperbackUS:                            {"Trade paperback (US)", "In North America, a category of paperback characterized partly by page size and partly by target market and terms of trade. AKA ‘quality paperback’, and including textbooks. Most paperback books sold in North America except ‘mass-market’ (B101) and ‘tall rack’ (B107) are correctly described with this code. Use with Product Form code BC.", "Trade paperback (US)", "I Nord-Amerika, en kategori \"paperback\" karakterisert delvis på størrelse og delvis på markedet den er beregnet for, samt handelsbetingelsene, Brukes om \"quality paperback\". Brukes med \"Product Form\" kode BC"},
		DigestFormatPaperback:                       {"Digest format paperback", "In North America, a category of paperback characterized by page size and generally used for children’s books; use with Product Form code BC. Note: was wrongly shown as B102 (duplicate entry) in Issue 3.", "Digest format paperback", "I Nord-Amerika, en kategori \"paperback\" karakterisert på størrelse og brukes generelt for barnebøker. Brukes med \"Product Form\" kode BC"},
		AFormatPaperback:                            {"A-format paperback", "In UK, a category of paperback characterized by page size (normally 178 x 111 mm approx); use with Product Form code BC.", "A-format paperback", "UK, en kategori \"paperback\" karakterisert på størrelse (normalt ca. 178 x 111 mm), Brukes med \"Product Form\" kode BC"},
		BFormatPaperback:                            {"B-format paperback", "In UK, a category of paperback characterized by page size (normally 198 x 129 mm approx); use with Product Form code BC.", "B-format paperback", "UK, en kategori \"paperback\" karakterisert på størrelse (normalt ca. 198 x 129 mm), Brukes med \"Product Form\" kode BC"},
		TradePaperbackUK:                            {"Trade paperback (UK)", "In UK, a category of paperback characterized partly by size (usually in traditional hardback dimensions), and often used for paperback originals; use with Product Form code BC (replaces ‘C-format’ from former List 8).", "Trade paperback (UK)", "UK, en kategori \"paperback\" karakterisert delvis på størrelse (vanligvis i tradisjonell \"hardback\" dimensjoner) og ofte brukt på \"paperback\" originaler. Bruk med \"Product Form\" kode BC. (erstatter 'C-format' fra utgått liste 8)"},
		TallRackPaperbackUS:                         {"Tall rack paperback (US)", "In North America, a category of paperback characterised partly by page size and partly by target market and terms of trade; use with Product Form code BC.", "Tall rack paperback (US)", "I Nord-Amerika, en kategori \"paperback\" karakterisert delvis på størrelse og delvis på markedet den er beregnet for, samt handelsbetingelsen. Brukes med \"Product Form\" kode BC"},
		A5SizeTankobon:                              {"A5 size Tankobon", "210x148mm.", "A5 size Tankobon", "Japansk innbundet format, 210x148mm."},
		JISB5SizeTankobon:                           {"JIS B5 size Tankobon", "Japanese B-series size, 257x182mm.", "JIS B5 size Tankobon", "Japansk innbundet format, 257x182mm."},
		JISB6SizeTankobon:                           {"JIS B6 size Tankobon", "Japanese B-series size, 182x128mm.", "JIS B6 size Tankobon", "Japansk innbundet format, 182x128mm."},
		A6SizeBunko:                                 {"A6 size Bunko", "148x105mm.", "A6 size Bunko", "Japansk heftet format, 148x105mm."},
		B40DoriShinsho:                              {"B40-dori Shinsho", "Japanese format, 182x103mm or 173x105mm.", "B40-dori Shinsho", "Japansk heftet format, 182x103mm eller 173x105mm."},
		PocketSweden:                                {"Pocket (Sweden)", "A Swedish paperback format, use with Product Form Code BC.", "Pocket (Sweden)", "Et svensk heftet format, bruk med Product Form Code BC"},
		StorpocketSweden:                            {"Storpocket (Sweden)", "A Swedish paperback format, use with Product Form Code BC.", "Storpocket (Sweden)", "Et svensk heftet format, bruk med Product Form Code BC"},
		KartonnageSweden:                            {"Kartonnage (Sweden)", "A Swedish hardback format, use with Product Form Code BB.", "Kartonnage (Sweden)", "Et svensk innbundet format, bruk med Code BB"},
		FlexbandSweden:                              {"Flexband (Sweden)", "A Swedish softback format, use with Product Form Code BC.", "Flexband (Sweden)", "Et svensk heftet format, bruk med Product Form Code BC"},
		Mook:                                        {"Mook", "In Japan, a softback book in the format of a magazine but sold like a book.", "Mook", "Japansk heftet format som ligner på magasin-/tidsskriftsformat, men selges som bok"},
		Dwarsligger:                                 {"Dwarsligger", "Also called ‘Flipback’. A softback book in a specially compact proprietary format with pages printed in landscape on very thin paper and bound along the long (top) edge – see www.dwarsligger.com.", "Dwarsligger", "Kalles også ‘Flipback’. Et heftet format i et spesielt kompakt, proprietært format med sider som er liggende, hvor papiret er svært tynt, og sidene er bundet sammen langs toppen, se www.dwarsligger.com."},
		Form46Size:                                  {"46 size", "Japanese format: 188x127mm.", "46 size", "Japansk format: 188x127mm."},
		Form46HenkeiSize:                            {"46-Henkei size", "Japanese format.", "46-Henkei size", "Japansk format."},
		A4:                                          {"A4", "297x210mm.", "A4", "297x210mm."},
		A4HenkeiSize:                                {"A4-Henkei size", "Japanese format.", "A4-Henkei size", "Japansk format."},
		A5HenkeiSize:                                {"A5-Henkei size", "Japanese format.", "A5-Henkei size", "Japansk format."},
		B5HenkeiSize:                                {"B5-Henkei size", "Japanese format.", "B5-Henkei size", "Japansk format."},
		B6HenkeiSize:                                {"B6-Henkei size", "Japanese format.", "B6-Henkei size", "Japansk format."},
		ABSize:                                      {"AB size", "257x210mm.", "AB size", "257x210mm."},
		JISB7Size:                                   {"JIS B7 size", "Japanese B-series size, 128x91mm.", "JIS B7 size", "Japansk format, 128x91mm."},
		KikuSize:                                    {"Kiku size", "Japanese format: 218x152mm or 227x152mm.", "Kiku size", "Japansk format, 218x152mm eller 227x152mm."},
		KikuHenkeiSize:                              {"Kiku-Henkei size", "Japanese format.", "Kiku-Henkei size", "Japansk format"},
		JISB4Size:                                   {"JIS B4 size", "Japanese B-series size, 257x364mm.", "JIS B4 size", "Japansk format, 364x257mm."},
		PaperbackDE:                                 {"Paperback (DE)", "German paperback format, greater than 205mm high, with flaps. Use with Product form code BC.", "Paperback (DE)", "Tysk heftet format. Høyere enn 205mm, med innbretter. Brukes med Product form-kode BC"},
		ColoringJoinTheDotBook:                      {"Coloring / join-the-dot book", "", "Fargebok", ""},
		LiftTheFlapBook:                             {"Lift-the-flap book", "", "Utbrettsbok/klaffebok", ""},
		MiniatureBook:                               {"Miniature book", "Note: was wrongly shown as B203 (duplicate entry) in Issue 3.", "Miniatyrbok", "Note: was wrongly shown as B203 (duplicate entry) in Issue 3."},
		MovingPictureFlickerBook:                    {"Moving picture / flicker book", "", "Flimrebok", "Eng: Moving picture / flicker book"},
		PopUpBook:                                   {"Pop-up book", "", "Sprett-opp-bok", ""},
		ScentedsmellyBook:                           {"Scented / ‘smelly’ book", "", "Bok med lukteffekter\n", ""},
		SoundStorynoisyBook:                         {"Sound story / ‘noisy’ book", "", "Bok med lydeffekter\n", ""},
		StickerBook:                                 {"Sticker book", "", "Klistremerkebok", ""},
		TouchAndFeelBook:                            {"Touch-and-feel book", "A book whose pages have a variety of textured inserts designed to stimulate tactile exploration: see also B214 and B215.", "Ta og føle på-bok", "Inneholder flere forskjellige materialer for å stimulere taktil utforskning, se også B214 and B215."},
		DieCutBook:                                  {"Die-cut book", "A book which is cut into a distinctive non-rectilinear shape and/or in which holes or shapes have been cut internally. (‘Die-cut’ is used here as a convenient shorthand, and does not imply strict limitation to a particular production process.).", "Die-cut book", "A book which is cut into a distinctive non-rectilinear shape and/or in which holes or shapes have been cut internally. (‘Die-cut’ is used here as a convenient shorthand, and does not imply strict limitation to a particular production process.)."},
		BookAsToy:                                   {"Book-as-toy", "A book which is also a toy, or which incorporates a toy as an integral part. (Do not, however, use B213 for a multiple-item product which includes a book and a toy as separate items.).", "Bok-som-leke", "Bok som også er en leke, eller som har en leke som vesentlig del av produktet. (Bruk ikke B213 for produkter hvor bok og leke er separate produkter solgt sammen)."},
		SoftToTouchBook:                             {"Soft-to-touch book", "A book whose cover has a soft textured finish, typically over board.", "Bokomslag-som-er-mykt-å-ta-på", "Bok hvor omslaget har en myk overflate med tekstur, ofte kartonert"},
		FuzzyFeltBook:                               {"Fuzzy-felt book", "A book with detachable felt pieces and textured pages on which they can be arranged.", "Filt-bok", "Bok med avtagbare deler i filt og sider med tekstur hvor figurene kan festes"},
		PictureBook:                                 {"Picture book", "Children’s picture book: use with applicable Product Form code.", "Billedbok", "Bildebok for barn. Bruk med passende produktformkode"},
		CarouselBook:                                {"‘Carousel’ Book", "(aka ‘Star’ book). Tax treatment of products may differ from that of products with similar codes such as Book as toy or Pop-up book).", "Karusell'-bok", "Kalles også 'stjernebok'"},
		LooseLeafSheetsAndBinder:                    {"Loose leaf – sheets and binder", "Use with Product Form code BD.", "Løsark - ark og perm", "Brukes med Product form-kode BD."},
		LooseLeafBinderOnly:                         {"Loose leaf – binder only", "Use with Product Form code BD.", "Løsark - kun perm", "Brukes med Product form-kode BD."},
		LooseLeafSheetsOnly:                         {"Loose leaf – sheets only", "Use with Product Form code BD.", "Løsark - kun ark", "Brukes med Product form-kode BD."},
		Sewn:                                        {"Sewn", "AKA stitched; for ‘saddle-sewn’, see code B310.", "Sydd", "eng stitched; se også kode B310."},
		UnsewnAdhesiveBound:                         {"Unsewn / adhesive bound", "Including ‘perfect bound’, ‘glued’.", "Limt", "Including ‘perfect bound’, ‘glued’."},
		LibraryBinding:                              {"Library binding", "Strengthened binding intended for libraries.", "Biblioteksinnbinding", "Forsterket innbinding beregnet på bibliotek"},
		ReinforcedBinding:                           {"Reinforced binding", "Strengthened binding, not specifically intended for libraries.", "Forsterket innbinding", "Forsterket innbinding ikke spesielt beregnet på biblioteker"},
		HalfBound:                                   {"Half bound", "Must be accompanied by a code specifiying a material, eg ‘half-bound real leather’.", "Halvbind", "Må ledsages av kode som angir materialet. F.eks. \"halvbind - ekte skinn\""},
		QuarterBound:                                {"Quarter bound", "Must be accompanied by a code specifiying a material, eg ‘quarter bound real leather’.", "Kvartbind", "Må ledsages av kode som angir materialet. F.eks. \"'kvartbind - ekte skinn\""},
		SaddleSewn:                                  {"Saddle-sewn", "AKA ‘saddle-stitched’ or ‘wire-stitched’.", "Stiftet", "Stiftet"},
		CombBound:                                   {"Comb bound", "Round or oval plastic forms in a clamp-like configuration: use with Product Form code BE.", "Plastspiral", "Runde eller ovale plastikklemmer (eng: comb bound). Brukes sammen med Product form-kode BE"},
		WireO:                                       {"Wire-O", "Twin loop metal or plastic spine: use with Product Form code BE.", "Wire-O", "Doble gjennomgående metall- eller plastikkspiraler. Brukes sammen med Product form-kode BE"},
		ConcealedWire:                               {"Concealed wire", "Cased over Wire-O binding: use with Product Form code BE.", "Skjult spiral", "Wire-O som er skjult. Brukes sammen med Product form-kode BE"},
		ClothOverBoards:                             {"Cloth over boards", "AKA fabric, linen over boards.", "Tekstilkledd innbinding", "eng: fabric, linen over boards."},
		PaperOverBoards:                             {"Paper over boards", "", "Papirkledd innbinding", ""},
		LeatherReal:                                 {"Leather, real", "", "Ekte skinn", ""},
		LeatherImitation:                            {"Leather, imitation", "", "Kunstskinn", ""},
		LeatherBonded:                               {"Leather, bonded", "", "Forsterket skinn", ""},
		Vellum:                                      {"Vellum", "", "Pergamentbind", ""},
		Cloth:                                       {"Cloth", "Cloth, not necessarily over boards – cf B401.", "Tekstil", "Tekstil, men ikke nødvendigvis over kartong. Se også kode B401"},
		ImitationCloth:                              {"Imitation cloth", "Spanish ‘simil-tela’.", "Imitert tekstil", "Spansk: simil-tela"},
		Velvet:                                      {"Velvet", "", "Fløyel", ""},
		FlexiblePlasticVinylCover: {"Flexible plastic/vinyl cover", "AKA “flexibound”: use with Product Form code BC.", "Fleksibelt plast- eller vinylomslag", "Også kalt 'flexibound\": brukes med produktformkode BC"},
		PlasticCovered:            {"Plastic-covered", "", "Plastet", "eng: Plastic-covered"},
		VinylCovered:              {"Vinyl-covered", "", "Vinyl-dekket", "eng: Vinyl-covered"},
		LaminatedCover:            {"Laminated cover", "Book, laminating material unspecified: use L101 for ‘whole product laminated’, eg a laminated sheet map or wallchart.", "Laminert omslag", "Bok hvor lamineringsmaterialet er uspesifisert: bruk L101 for produkter hvor hele produktet er laminert, for eksempel et laminert veggkart"},
		WithDustJacket:            {"With dust jacket", "Type unspecified.", "Har smussomslag", "Type uspesifisert"},
		WithPrintedDustJacket:     {"With printed dust jacket", "Used to distinguish from B503.", "Har trykket smussomslag", "Brukes for å skille fra B503"},
		WithTranslucentDustCover:  {"With translucent dust cover", "With translucent paper or plastic protective cover.", "Har gjennomsiktig smussomslag", "Med gjennomsiktig papir eller beskyttelsesomslag av plastikk"},
		WithFlaps:                 {"With flaps", "For paperback with flaps.", "Med innbretter", "For heftede bøker med innbretter"},
		WithThumbIndex:            {"With thumb index", "", "Har utstanset register", "Har register utstanset i snittet på boken"},
		WithRibbonMarkerS:         {"With ribbon marker(s)", "If the number of markers is significant, it can be stated as free text in <ProductFormDescription>.", "Med lesebånd", "Hvis antall lesebånd er viktig så kan det spesifiseres i <ProductFormDescription>"},
		WithZipFastener:           {"With zip fastener", "", "Med glidelås", ""},
		WithButtonSnapFastener:    {"With button snap fastener", "", "Med trykknapp", ""},
		WithLeatherEdgeLining:     {"With leather edge lining", "AKA yapp edge?.", "Med skinnrygg", "AKA yapp edge?."},
		RoughFront:                {"Rough front", "With edge trimming such that the front edge is ragged, not neatly and squarely trimmed: AKA deckle edge, feather edge, uncut edge, rough cut.", "Grovt frontsnitt", "Med sidekanter som er frynsete/fillete, ikke pent kuttet"},
		Foldout:                   {"Foldout", "With one or more gatefold or foldout sections bound in.", "Med utbrett", "Med en eller flere utbrettsseksjoner"},
		TurnAroundBook:            {"Turn-around book", "A book in which half the content is printed upside-down, to be read the other way round. Also known as a ‘flip-book’ (usually an omnibus of two works).", "Opp-ned-bok", "Ei bok hvor halvparten av innholdet er trykket opp-ned"},
		UnflippedMangaFormat:      {"Unflipped manga format", "Manga with pages and panels in the sequence of the original Japanese, but with Western text.", "Unflipped manga format", "Manga med sider og paneler i samme sekvens som den japanske originalen, men med vestlig tekst"},
		Syllabification:           {"Syllabification", "Text shows syllable breaks", "Syllabification", "Text shows syllable breaks"},
		UKUncontractedBraille:     {"UK Uncontracted Braille", "Single letters only. Was formerly identified as UK Braille Grade 1.", "UK Uncontracted Braille", "Single letters only. Was formerly identified as UK Braille Grade 1."},
		UKContractedBraille:       {"UK Contracted Braille", "With some letter combinations. Was formerly identified as UK Braille Grade 2.", "UK Contracted Braille", "With some letter combinations. Was formerly identified as UK Braille Grade 2."},
		USBraille:                 {"US Braille", "DEPRECATED- use B704/B705 as appropriate instead.", "US Braille", "UTGÅTTT. Bruk B704/B705 isteden"},
		USUncontractedBraille:     {"US Uncontracted Braille", "", "US Uncontracted Braille", ""},
		USContractedBraille:       {"US Contracted Braille", "", "US Contracted Braille", ""},
		UnifiedEnglishBraille:     {"Unified English Braille", "", "Unified English Braille", ""},
		Moon:                    {"Moon", "Moon embossed alphabet, used by some print-impaired readers who have difficulties with Braille.", "Moon", "Moon-alfabet, brukt av personer som har vanskeligheter med Braille"},
		RealVideoFormat:         {"Real Video format", "Proprietary RealNetworks format. Includes Real Video packaged within a .rm RealMedia container.", "Real video-format", "Proprietary RealNetworks format. Includes Real Video packaged within a .rm RealMedia container."},
		QuicktimeFormat:         {"Quicktime format", "", "Quicktime-format", ""},
		AVIFormat:               {"AVI format", "", "AVI format", ""},
		WindowsMediaVideoFormat: {"Windows Media Video format", "", "Windows Media Video-format", ""},
		MPEG4:                      {"MPEG-4", "", "MPEG-4", ""},
		MSDOS:                      {"MS-DOS", "Use with an applicable Product Form code D*; note that more detail of operating system requirements can be given in a Product Form Feature composite.", "MS-DOS", "Brukes med D*-kode fra Product Form; ytterligere detaljer om krav til operativsystem kan sendes i Product Form Feature"},
		Windows:                    {"Windows", "Use with an applicable Product Form code D*; see note on D201.", "Windows", "Brukes med D*-kode fra Product Form; se note for D201."},
		Macintosh:                  {"Macintosh", "Use with an applicable Product Form code D*; see note on D201.", "Macintosh", "Brukes med D*-kode fra Product Form; se note for D201."},
		UNIXLINUX:                  {"UNIX / LINUX", "Use with an applicable Product Form code D*; see note on D201.", "UNIX / LINUX", "Brukes med D*-kode fra Product Form; se note for D201."},
		OtherOperatingSystemS:      {"Other operating system(s)", "Use with an applicable Product Form code D*; see note on D201.", "Andre operativsystem(er)", "Brukes med D*-kode fra Product Form; se note for D201."},
		PalmOS:                     {"Palm OS", "Use with an applicable Product Form code D*; see note on D201.", "Palm OS", "Brukes med D*-kode fra Product Form; se note for D201."},
		WindowsMobile:              {"Windows Mobile", "Use with an applicable Product Form code D*; see note on D201.", "Windows Mobile", "Brukes med D*-kode fra Product Form; se note for D201."},
		MicrosoftXBox:              {"Microsoft XBox", "Use with Product Form code DE or DB as applicable.", "Microsoft XBox", "Brukes sammen med Product Form-kode DE eller DB"},
		NintendoGameboyColor:       {"Nintendo Gameboy Color", "Use with Product Form code DE or DB as applicable.", "Nintendo Gameboy Color", "Brukes sammen med Product Form-kode DE eller DB"},
		NintendoGameboyAdvanced:    {"Nintendo Gameboy Advanced", "Use with Product Form code DE or DB as applicable.", "Nintendo Gameboy Advanced", "Brukes sammen med Product Form-kode DE eller DB"},
		NintendoGameboy:            {"Nintendo Gameboy", "Use with Product Form code DE or DB as applicable.", "Nintendo Gameboy", "Brukes sammen med Product Form-kode DE eller DB"},
		NintendoGamecube:           {"Nintendo Gamecube", "Use with Product Form code DE or DB as applicable.", "Nintendo Gamecube", "Brukes sammen med Product Form-kode DE eller DB"},
		Nintendo64:                 {"Nintendo 64", "Use with Product Form code DE or DB as applicable.", "Nintendo 64", "Brukes sammen med Product Form-kode DE eller DB"},
		SegaDreamcast:              {"Sega Dreamcast", "Use with Product Form code DE or DB as applicable.", "Sega Dreamcast", "Brukes sammen med Product Form-kode DE eller DB"},
		SegaGenesisMegadrive:       {"Sega Genesis/Megadrive", "Use with Product Form code DE or DB as applicable.", "Sega Genesis/Megadrive", "Brukes sammen med Product Form-kode DE eller DB"},
		SegaSaturn:                 {"Sega Saturn", "Use with Product Form code DE or DB as applicable.", "Sega Saturn", "Brukes sammen med Product Form-kode DE eller DB"},
		SonyPlayStation1:           {"Sony PlayStation 1", "Use with Product Form code DE or DB as applicable.", "Sony PlayStation 1", "Brukes sammen med Product Form-kode DE eller DB"},
		SonyPlayStation2:           {"Sony PlayStation 2", "Use with Product Form code DE or DB as applicable.", "Sony PlayStation 2", "Brukes sammen med Product Form-kode DE eller DB"},
		NintendoDualScreen:         {"Nintendo Dual Screen", "", "Nintendo Dual Screen", ""},
		SonyPlayStation3:           {"Sony PlayStation 3", "", "Sony PlayStation 3", ""},
		Xbox360:                    {"Xbox 360", "", "Xbox 360", ""},
		NintendoWii:                {"Nintendo Wii", "", "Nintendo Wii", ""},
		SonyPlayStationPortablePSP: {"Sony PlayStation Portable (PSP)", "", "Sony PlayStation Portable (PSP)", ""},
		Other:                          {"Other", "No code allocated for this e-publication format yet.", "Andre", "Ingen kode er tildelt dette epublikasjonsformatet ennå"},
		EPUB:                           {"EPUB", "The Open Publication Structure / OPS Container Format standard of the International Digital Publishing Forum (IDPF) [File extension .epub].", "EPUB", "The Open Publication Structure / OPS Container Format standard of the International Digital Publishing Forum (IDPF) [File extension .epub]."},
		OEB:                            {"OEB", "The Open EBook format of the IDPF, a predecessor of the full EPUB format, still (2008) supported as part of the latter [File extension .opf]. Includes EPUB format up to and including version 2 – but prefer code E101 for EPUB 2, and always use code E101 for EPUB 3.", "OEB", "The Open EBook format of the IDPF, a predecessor of the full EPUB format, still (2008) supported as part of the latter [File extension .opf]. Includes EPUB format up to and including version 2 – but prefer code E101 for EPUB 2, and always use code E101 for EPUB 3."},
		DOC:                            {"DOC", "Microsoft Word binary document format [File extension .doc].", "DOC", "Microsoft Word binary document format [File extension .doc]."},
		DOCX:                           {"DOCX", "Office Open XML / Microsoft Word XML document format (ISO/IEC 29500:2008) [File extension .docx].", "DOCX", "Office Open XML / Microsoft Word XML document format (ISO/IEC 29500:2008) [File extension .docx]."},
		HTML:                           {"HTML", "HyperText Mark-up Language [File extension .html, .htm].", "HTML", "HyperText Mark-up Language [File extension .html, .htm]."},
		ODF:                            {"ODF", "Open Document Format [File extension .odt].", "ODF", "Open Document Format [File extension .odt]."},
		PDF:                            {"PDF", "Portable Document Format (ISO 32000-1:2008) [File extension .pdf].", "PDF", "Portable Document Format (ISO 32000-1:2008) [File extension .pdf]."},
		PDFA:                           {"PDF/A", "PDF archiving format defined by ISO 19005-1:2005 [File extension .pdf].", "PDF/A", "PDF archiving format defined by ISO 19005-1:2005 [File extension .pdf]."},
		RTF:                            {"RTF", "Rich Text Format [File extension .rtf].", "RTF", "Rich Text Format [File extension .rtf]."},
		SGML:                           {"SGML", "Standard Generalized Mark-up Language.", "SGML", "Standard Generalized Mark-up Language."},
		TCR:                            {"TCR", "A compressed text format mainly used on Psion handheld devices [File extension .tcr].", "TCR", "A compressed text format mainly used on Psion handheld devices [File extension .tcr]."},
		TXT:                            {"TXT", "Text file format [File extension .txt].", "TXT", "Text file format [File extension .txt]."},
		XHTML:                          {"XHTML", "Extensible Hypertext Markup Language [File extension .xhtml, .xht, .xml, .html, .htm].", "XHTML", "Extensible Hypertext Markup Language [File extension .xhtml, .xht, .xml, .html, .htm]."},
		ZTXT:                           {"zTXT", "A compressed text format mainly used on Palm handheld devices [File extension .pdb – see also E121, E125, E130].", "zTXT", "A compressed text format mainly used on Palm handheld devices [File extension .pdb – see also E121, E125, E130]."},
		XPS:                            {"XPS", "XML Paper Specification format [File extension .xps].", "XPS", "XML Paper Specification format [File extension .xps]."},
		AmazonKindle:                   {"Amazon Kindle", "A format proprietary to Amazon for use with its Kindle reading devices or software readers [File extensions .azw, .mobi, .prc].", "Amazon Kindle", "A format proprietary to Amazon for use with its Kindle reading devices or software readers [File extensions .azw, .mobi, .prc]."},
		BBeB:                           {"BBeB", "A Sony proprietary format for use with the Sony Reader and LIBRIé reading devices [File extension .lrf].", "BBeB", "A Sony proprietary format for use with the Sony Reader and LIBRIé reading devices [File extension .lrf]."},
		DXReader:                       {"DXReader", "A proprietary format for use with DXReader software.", "DXReader", "A proprietary format for use with DXReader software."},
		EBL:                            {"EBL", "A format proprietary to the Ebook Library service.", "EBL", "A format proprietary to the Ebook Library service."},
		Ebrary:                         {"Ebrary", "A format proprietary to the Ebrary service.", "Ebrary", "A format proprietary to the Ebrary service."},
		EReader:                        {"eReader", "A proprietary format for use with eReader (AKA ‘Palm Reader’) software on various hardware platforms [File extension .pdb – see also E114, E125, E130].", "eReader", "A proprietary format for use with eReader (AKA ‘Palm Reader’) software on various hardware platforms [File extension .pdb – see also E114, E125, E130]."},
		Exebook:                        {"Exebook", "A proprietary format with its own reading system for Windows platforms [File extension .exe].", "Exebook", "A proprietary format with its own reading system for Windows platforms [File extension .exe]."},
		FranklinEBookman:               {"Franklin eBookman", "A proprietary format for use with the Franklin eBookman reader.", "Franklin eBookman", "A proprietary format for use with the Franklin eBookman reader."},
		GemstarRocketbook:              {"Gemstar Rocketbook", "A proprietary format for use with the Gemstar Rocketbook reader [File extension .rb].", "Gemstar Rocketbook", "A proprietary format for use with the Gemstar Rocketbook reader [File extension .rb]."},
		ISilo:                          {"iSilo", "A proprietary format for use with iSilo software on various hardware platforms [File extension .pdb – see also E114, E121, E130].", "iSilo", "A proprietary format for use with iSilo software on various hardware platforms [File extension .pdb – see also E114, E121, E130]."},
		MicrosoftReader:                {"Microsoft Reader", "A proprietary format for use with Microsoft Reader software on Windows and Pocket PC platforms [File extension .lit].", "Microsoft Reader", "A proprietary format for use with Microsoft Reader software on Windows and Pocket PC platforms [File extension .lit]."},
		Mobipocket:                     {"Mobipocket", "A proprietary format for use with Mobipocket software on various hardware platforms [File extensions .mobi, .prc]. Includes Amazon Kindle formats up to and including version 7 – but prefer code E116 for version 7, and always use E116 for KF8.", "Mobipocket", "A proprietary format for use with Mobipocket software on various hardware platforms [File extensions .mobi, .prc]. Includes Amazon Kindle formats up to and including version 7 – but prefer code E116 for version 7, and always use E116 for KF8."},
		MyiLibrary:                     {"MyiLibrary", "A format proprietary to the MyiLibrary service.", "MyiLibrary", "A format proprietary to the MyiLibrary service."},
		NetLibrary:                     {"NetLibrary", "A format proprietary to the NetLibrary service.", "NetLibrary", "A format proprietary to the NetLibrary service."},
		Plucker:                        {"Plucker", "A proprietary format for use with Plucker reader software on Palm and other handheld devices [File extension .pdb – see also E114, E121, E125].", "Plucker", "A proprietary format for use with Plucker reader software on Palm and other handheld devices [File extension .pdb – see also E114, E121, E125]."},
		VitalBook:                      {"VitalBook", "A format proprietary to the VitalSource service.", "VitalBook", "A format proprietary to the VitalSource service."},
		Vook:                           {"Vook", "A proprietary digital product combining text and video content and available to be used online or as a downloadable application for a mobile device – see www.vook.com.", "Vook", "A proprietary digital product combining text and video content and available to be used online or as a downloadable application for a mobile device – see www.vook.com."},
		GoogleEdition:                  {"Google Edition", "An epublication made available by Google in association with a publisher; readable online on a browser-enabled device and offline on designated ebook readers.", "Google Edition", "An epublication made available by Google in association with a publisher; readable online on a browser-enabled device and offline on designated ebook readers."},
		BookappForIOS:                  {"Book ‘app’ for iOS", "Epublication packaged as application for iOS (eg Apple iPhone, iPad etc), containing both executable code and content. Use <ProductContentType> to describe content, and <ProductFormFeatureType> to list detailed technical requirements.", "Bok-app for iOS", "Epublication packaged as application for iOS (eg Apple iPhone, iPad etc), containing both executable code and content. Use <ProductContentType> to describe content, and <ProductFormFeatureType> to list detailed technical requirements."},
		BookappForAndroid:              {"Book ‘app’ for Android", "Epublication packaged as application for Android (eg Android phone or tablet), containing both executable code and content. Use <ProductContentType> to describe content, and <ProductFormFeatureType> to list detailed technical requirements.", "Bok-app for Android", "Epublication packaged as application for Android (eg Android phone or tablet), containing both executable code and content. Use <ProductContentType> to describe content, and <ProductFormFeatureType> to list detailed technical requirements."},
		BookappForOtherOperatingSystem: {"Book ‘app’ for other operating system", "Epublication packaged as application, containing both executable code and content. Use where other ‘app’ codes are not applicable. Technical requirements such as target operating system and/or device should be provided eg in <ProductFormFeatureType>. Content type (text or text plus various ‘enhancements’) may be described with <ProductContentType>.", "Bok-app for andre operativsystemer", "Epublication packaged as application, containing both executable code and content. Use where other ‘app’ codes are not applicable. Technical requirements such as target operating system and/or device should be provided eg in <ProductFormFeatureType>. Content type (text or text plus various ‘enhancements’) may be described with <ProductContentType>."},
		CEB:                             {"CEB", "Founder Apabi’s proprietary basic e-book format.", "CEB", "Founder Apabi’s proprietary basic e-book format."},
		CEBX:                            {"CEBX", "Founder Apabi’s proprietary XML e-book format.", "CEBX", "Founder Apabi’s proprietary XML e-book format."},
		IBook:                           {"iBook", "Apple’s iBook format (a proprietary extension of EPUB), can only be read on Apple iOS devices.", "iBook", "Apple’s iBook format (a proprietary extension of EPUB), can only be read on Apple iOS devices."},
		EPIB:                            {"ePIB", "Proprietary format used by Barnes and Noble, readable on NOOK devices and Nook reader software.", "ePIB", "Proprietary format used by Barnes and Noble, readable on NOOK devices and Nook reader software."},
		SCORM:                           {"SCORM", "Sharable Content Object Reference Model, standard content and packaging format for e-learning objects.", "SCORM", "Sharable Content Object Reference Model, standard content and packaging format for e-learning objects."},
		EBP:                             {"EBP", "E-book Plus (proprietary Norwegian e-book format).", "EBP", "E-bok pluss (proprietært norsk e-bokformat)"},
		Reflowable:                      {"Reflowable", "Use when a particular e-publication type (specified using codes E100 and upwards) has both fixed format and reflowable variants.", "Flytende sidevisning", "Brukes når en epublikasjon (spesifisert av kodene E100 og oppover) har både flytende sidevisning og fast sidevisning"},
		FixedFormat:                     {"Fixed format", "Use when a particular e-publication type (specified using codes E100 and upwards) has both fixed format and reflowable variants.", "Fast sidevisning", "Brukes når en epublikasjon (spesifisert av kodene E100 og oppover) har både har en variant med flytende sidevisning og en med fast sidevisning"},
		ReadableOffline:                 {"Readable offline", "All e-publication resources are included within the e-publication package.", "Kan leses offline", "Alt innhold er inkludert i epublikasjonen"},
		RequiresNetworkConnection:       {"Requires network connection", "E-publication requires a network connection to access some resources (eg an enhanced e-book where video clips are not stored within the e-publication package itself, but are delivered via an internet connection).", "Krever nettilgang", "Epublikasjonen krever nettilgang for å kunne bruke noe av innholdet (f.eks. beriket e-bok hvor videoklipp ikke er lagret i selve epublikasjonen, men er levert via en internettilgang)"},
		ContentRemoved:                  {"Content removed", "Resources (eg images) present in other editions have been removed from this product, eg due to rights issues.", "Innhold er fjernet", "Ressurser (f.eks. bilder) som fins i andre utgaver er fjernet fra dette produktet, for eksempel pga. rettighetsspørsmål"},
		Landscape:                       {"Landscape", "Use for fixed-format e-books optimised for landscape display. Also include an indication of the optimal screen aspect ratio.", "Liggende", "Brukes for e-bøker med fast sidevisning som er optimalisert for liggende visning. Man kan også sende informasjon om optimal skjermstørrelse"},
		Portrait:                        {"Portrait", "Use for fixed-format e-books optimised for portrait display. Also include an indication of the optimal screen aspect ratio.", "Stående", "Brukes for e-bøker med fast sidevisning som er optimalisert for stående visning. Man kan også sende informasjon om optimal skjermstørrelse"},
		Aspect54:                        {"5:4", "Use for fixed-format e-books optimised for displays with a 5:4 aspect ratio (eg 1280x1024 pixels etc, assuming square pixels). Note that aspect ratio codes are NOT specific to actual screen dimensions or pixel counts, but to the ratios between two dimensions or two pixel counts.", "5:4", "Brukes for e-bøker med fast sidevisning som er optimalisert for å vises på skjermer med sideforhold 5:4 (f.eks. 1280x1024 piksler, dersom man går ut fra at det snakk om kvadratiske piksler). Merk at kodene for sideforhold IKKE er knyttet til faktiske skjermstørrelser eller konkrete antall piksler, men til forholdet mellom de to sidene eller to pikselantall."},
		Aspect43:                        {"4:3", "Use for fixed-format e-books optimised for displays with a 4:3 aspect ratio (eg 800x600, 1024x768, 2048x1536 pixels etc).", "4:3", "Brukes for e-bøker med fast sidevisning som er optimalisert for å vises på skjermer med sideforhold 4:3 (f.eks. 800x600, 1024x768, 2048x1536 piksler osv.)"},
		Aspect32:                        {"3:2", "Use for fixed-format e-books optimised for displays with a 3:2 aspect ratio (eg 960x640, 3072x2048 pixels etc).", "3:2", "Brukes for e-bøker med fast sidevisning som er optimalisert for å vises på skjermer med sideforhold 3:2 (f.eks. 960x640, 3072x2048 piksler osv.)"},
		Aspect1610:                      {"16:10", "Use for fixed-format e-books optimised for displays with a 16:10 aspect ratio (eg 1440x900, 2560x1600 pixels etc).", "16:10", "Brukes for e-bøker med fast sidevisning som er optimalisert for å vises på skjermer med sideforhold 16:10 (f.eks. 1440x900, 2560x1600 piksler osv.)"},
		Aspect169:                       {"16:9", "Use for fixed-format e-books optimised for displays with a 16:9 aspect ratio (eg 1024x576, 1920x1080, 2048x1152 pixels etc).", "16:9", "Brukes for e-bøker med fast sidevisning som er optimalisert for å vises på skjermer med sideforhold 16:9 (f.eks. 1024x576, 1920x1080, 2048x1152 piksler osv.)"},
		Laminated:                       {"Laminated", "Whole product laminated (eg laminated map, fold-out chart, wallchart, etc): use B415 for book with laminated cover.", "Laminert", "Hele produktet er laminert (f.eks laminerte kart, utbrettsplansjer, veggplansjer osv.): bruk B415 for bøker med laminert omslag"},
		DeskCalendar:                    {"Desk calendar", "Use with Product Form code PC.", "Bordkalender", "Brukes med produktform PC"},
		MiniCalendar:                    {"Mini calendar", "Use with Product Form code PC.", "Minikalender", "Brukes med produktform PC"},
		EngagementCalendar:              {"Engagement calendar", "Use with Product Form code PC.", "Avtalekalender", "Brukes med produktform PC (eng: Engagement calendar)"},
		DayByDayCalendar:                {"Day by day calendar", "Use with Product Form code PC.", "Dag-for-dag kalender", "Brukes med produktform PC"},
		PosterCalendar:                  {"Poster calendar", "Use with Product Form code PC.", "Plakatkalender", "Brukes med produktform PC"},
		WallCalendar:                    {"Wall calendar", "Use with Product Form code PC.", "Veggkalender", "Brukes med produktform PC"},
		PerpetualCalendar:               {"Perpetual calendar", "Use with Product Form code PC.", "Evighetskalender", "Brukes med produktform PC"},
		AdventCalendar:                  {"Advent calendar", "Use with Product Form code PC.", "Adventskalender", "Brukes med produktform PC"},
		BookmarkCalendar:                {"Bookmark calendar", "Use with Product Form code PC.", "Bokmerkekalender", "Brukes med produktform PC"},
		StudentCalendar:                 {"Student calendar", "Use with Product Form code PC.", "Studentkalender", "Brukes med produktform PC"},
		ProjectCalendar:                 {"Project calendar", "Use with Product Form code PC.", "Prosjektkalender", "Brukes med produktform PC"},
		AlmanacCalendar:                 {"Almanac calendar", "Use with Product Form code PC.", "Almanakk", "Brukes med produktform PC"},
		OtherCalendar:                   {"Other calendar", "A calendar that is not one of the types specified elsewhere: use with Product Form code PC.", "Andre kalendere", "En kalender som ikke faller inn under noen av de typene som er spesifisert andre steder i kodelista. Brukes med produktform PC"},
		OtherCalendarOrOrganiserProduct: {"Other calendar or organiser product", "A product that is associated with or ancillary to a calendar or organiser, eg a deskstand for a calendar, or an insert for an organiser: use with Product Form code PC or PS.", "Andre kalender- eller dagboksprodukter", "Et produkt som er tilleggsutstyr til en kalender eller en filofax, for eksempel et kalenderstativ eller ekstra innhold til en filofax. Brukes med produktform PC eller PS."},
		HardbackStationery:              {"Hardback (stationery)", "Stationery item in hardback book format.", "Kontormateriell (innbundet)", "Kontormateriell i innbundet bokformat"},
		PaperbackSoftbackStationery:     {"Paperback / softback (stationery)", "Stationery item in paperback/softback book format.", "Kontormateriell (heftet)", "kontormateriell i heftet bokformat"},
		SpiralBoundStationery:           {"Spiral bound (stationery)", "Stationery item in spiral-bound book format.", "kontormateriell (spiral)", "Kontormateriell med spiralrygg (bokformat)"},
		LeatherFineBindingStationery:    {"Leather / fine binding (stationery)", "Stationery item in leather-bound book format, or other fine binding.", "Kontormateriell (lærinnbinding/luksusinnbinding)", "kontormateriell med lærinnbinding eller annen luksusinnbinding (bokformat)"},
		PAL:   {"PAL", "TV standard for video or DVD.", "PAL", "TV standard for video og DVD"},
		NTSC:  {"NTSC", "TV standard for video or DVD.", "NTSC", "TV standard for video og DVD"},
		SECAM: {"SECAM", "TV standard for video or DVD.", "SECAM", "TV standard for video og DVD"},
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
