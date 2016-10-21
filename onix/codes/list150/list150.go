package list150 // Product form

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Product form
const (
	Undefined                              = "00"
	Audio                                  = "AA"
	AudioCassette                          = "AB"
	CDAudio                                = "AC"
	DAT                                    = "AD"
	AudioDisc                              = "AE"
	AudioTape                              = "AF"
	MiniDisc                               = "AG"
	CDExtra                                = "AH"
	DVDAudio                               = "AI"
	DownloadableAudioFile                  = "AJ"
	PreRecordedDigitalAudioPlayer          = "AK"
	PreRecordedSDCard                      = "AL"
	OtherAudioFormat                       = "AZ"
	Book                                   = "BA"
	Hardback                               = "BB"
	PaperbackSoftback                      = "BC"
	LooseLeaf                              = "BD"
	SpiralBound                            = "BE"
	Pamphlet                               = "BF"
	LeatherFineBinding                     = "BG"
	BoardBook                              = "BH"
	RagBook                                = "BI"
	BathBook                               = "BJ"
	NoveltyBook                            = "BK"
	SlideBound                             = "BL"
	BigBook                                = "BM"
	PartWorkFascículo                      = "BN"
	FoldOutBookOrChart                     = "BO"
	FoamBook                               = "BP"
	OtherBookFormat                        = "BZ"
	SheetMap                               = "CA"
	SheetMapFolded                         = "CB"
	SheetMapFlat                           = "CC"
	SheetMapRolled                         = "CD"
	Globe                                  = "CE"
	OtherCartographic                      = "CZ"
	DigitalOnPhysicalCarrier               = "DA"
	CDROM                                  = "DB"
	CDI                                    = "DC"
	GameCartridge                          = "DE"
	Diskette                               = "DF"
	DVDROM                                 = "DI"
	SecureDigitalSDMemoryCard              = "DJ"
	CompactFlashMemoryCard                 = "DK"
	MemoryStickMemoryCard                  = "DL"
	USBFlashDrive                          = "DM"
	DoubleSidedCDDVD                       = "DN"
	OtherDigitalCarrier                    = "DZ"
	DigitalDeliveredElectronically         = "EA"
	DigitalDownloadAndOnline               = "EB"
	DigitalOnline                          = "EC"
	DigitalDownload                        = "ED"
	FilmOrTransparency                     = "FA"
	Slides                                 = "FC"
	OHPTransparencies                      = "FD"
	Filmstrip                              = "FE"
	Film                                   = "FF"
	OtherFilmOrTransparencyFormat          = "FZ"
	DigitalProductLicense                  = "LA"
	DigitalProductLicenseKey               = "LB"
	DigitalProductLicenseCode              = "LC"
	Microform                              = "MA"
	Microfiche                             = "MB"
	Microfilm                              = "MC"
	OtherMicroform                         = "MZ"
	MiscellaneousPrint                     = "PA"
	AddressBook                            = "PB"
	Calendar                               = "PC"
	Cards                                  = "PD"
	Copymasters                            = "PE"
	Diary                                  = "PF"
	Frieze                                 = "PG"
	Kit                                    = "PH"
	SheetMusic                             = "PI"
	PostcardBookOrPack                     = "PJ"
	Poster                                 = "PK"
	RecordBook                             = "PL"
	WalletOrFolder                         = "PM"
	PicturesOrPhotographs                  = "PN"
	Wallchart                              = "PO"
	Stickers                               = "PP"
	PlateLámina                            = "PQ"
	NotebookBlankBook                      = "PR"
	Organizer                              = "PS"
	Bookmark                               = "PT"
	OtherPrintedItem                       = "PZ"
	MultipleItemRetailProduct              = "SA"
	MultipleItemRetailProductBoxed         = "SB"
	MultipleItemRetailProductSlipCased     = "SC"
	MultipleItemRetailProductShrinkwrapped = "SD"
	MultipleItemRetailProductLoose         = "SE"
	MultipleItemRetailProductPartSEnclosed = "SF"
	Video                                  = "VA"
	Videodisc                              = "VF"
	DVDVideo                               = "VI"
	VHSVideo                               = "VJ"
	BetamaxVideo                           = "VK"
	VCD                                    = "VL"
	SVCD                                   = "VM"
	HDDVD                                  = "VN"
	BluRay                                 = "VO"
	UMDVideo                               = "VP"
	CBHD                                   = "VQ"
	OtherVideoFormat                       = "VZ"
	TradeOnlyMaterial                      = "XA"
	DumpbinEmpty                           = "XB"
	DumpbinFilled                          = "XC"
	CounterpackEmpty                       = "XD"
	CounterpackFilled                      = "XE"
	PosterPromotional                      = "XF"
	ShelfStrip                             = "XG"
	WindowPiece                            = "XH"
	Streamer                               = "XI"
	Spinner                                = "XJ"
	LargeBookDisplay                       = "XK"
	ShrinkWrappedPack                      = "XL"
	OtherPointOfSale                       = "XZ"
	GeneralMerchandise                     = "ZA"
	Doll                                   = "ZB"
	SoftToy                                = "ZC"
	Toy                                    = "ZD"
	Game                                   = "ZE"
	TShirt                                 = "ZF"
	EBookReader                            = "ZG"
	TabletComputer                         = "ZH"
	AudiobookPlayer                        = "ZI"
	Jigsaw                                 = "ZJ"
	OtherApparel                           = "ZY"
	OtherMerchandise                       = "ZZ"
)

var (
	sortedCodes = []string{"00", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BZ", "CA", "CB", "CC", "CD", "CE", "CZ", "DA", "DB", "DC", "DE", "DF", "DI", "DJ", "DK", "DL", "DM", "DN", "DZ", "EA", "EB", "EC", "ED", "FA", "FC", "FD", "FE", "FF", "FZ", "LA", "LB", "LC", "MA", "MB", "MC", "MZ", "PA", "PB", "PC", "PD", "PE", "PF", "PG", "PH", "PI", "PJ", "PK", "PL", "PM", "PN", "PO", "PP", "PQ", "PR", "PS", "PT", "PZ", "SA", "SB", "SC", "SD", "SE", "SF", "VA", "VF", "VI", "VJ", "VK", "VL", "VM", "VN", "VO", "VP", "VQ", "VZ", "XA", "XB", "XC", "XD", "XE", "XF", "XG", "XH", "XI", "XJ", "XK", "XL", "XZ", "ZA", "ZB", "ZC", "ZD", "ZE", "ZF", "ZG", "ZH", "ZI", "ZJ", "ZY", "ZZ"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Undefined:                     {"Undefined", "", "Udefinert", ""},
		Audio:                         {"Audio", "Audio recording – detail unspecified.", "Lyd", "Lydinnspilling - detaljer er uspesifisert"},
		AudioCassette:                 {"Audio cassette", "Audio cassette (analogue).", "Lydkassett", "Analog lydkassett"},
		CDAudio:                       {"CD-Audio", "Audio compact disc, in any recording format: use for ‘red book’ (conventional audio CD) and SACD, and use coding in Product Form Detail to specify the format, if required.", "CD (lyd)", "Lyd-CD i alle innspillingsformat: brukes for vanlige lyd-CD-er og SACD, bruk i tillegg koder i Product Form Detail dersom det er nødvendig å spesifisere formatet"},
		DAT:                           {"DAT", "Digital audio tape cassette.", "DAT (lyd)", "Digital audio tape kassett"},
		AudioDisc:                     {"Audio disc", "Audio disc (excluding CD-Audio).", "Lyddisk", "Lyddisk som ikke er CD (AC)"},
		AudioTape:                     {"Audio tape", "Audio tape (analogue open reel tape).", "Lydbånd", "Analogt lydbånd"},
		MiniDisc:                      {"MiniDisc", "Sony MiniDisc format.", "MiniDisc", "Sony MiniDisc format."},
		CDExtra:                       {"CD-Extra", "Audio compact disc with part CD-ROM content, also termed CD-Plus or Enhanced-CD: use for ‘blue book’ and ‘yellow/red book’ two-session discs.", "Lyd- og data-CD", "Audio compact disk med delvis CD-ROM innhold, kalles også CD-extra, CD-plus eller beriket CD"},
		DVDAudio:                      {"DVD Audio", "", "DVD (lyd)", ""},
		DownloadableAudioFile:         {"Downloadable audio file", "Audio recording downloadable online.", "Nedlastbar lydfil", "Online nedlastbart lydopptak/lydbok"},
		PreRecordedDigitalAudioPlayer: {"Pre-recorded digital audio player", "For example, Playaway audiobook and player: use coding in Product Form Detail to specify the recording format, if required.", "MP3-spiller med innhold", "F.eks Playaway lydbok og spiller, eller i Norge: Digibok. Bruk koder i Product Form Detail for å spesifisere formatet der det er nødvendig"},
		PreRecordedSDCard:             {"Pre-recorded SD card", "For example, Audiofy audiobook chip.", "SD minnekort med innnhold", "For eksempel, Audiofy lydbokbrikke"},
		OtherAudioFormat:              {"Other audio format", "Other audio format not specified by AB to AK.", "Annet lydformat", "Annet lydformat, ikke spesifisert i AB-AK"},
		Book:                          {"Book", "Book – detail unspecified.", "Bok", "Bok - detaljer uspesifisert"},
		Hardback:                      {"Hardback", "Hardback or cased book.", "Innbundet", "Innbundet bok eller bok i kassett"},
		PaperbackSoftback:             {"Paperback / softback", "Paperback or other softback book.", "Heftet", "Heftet bok (eng paperback/softback)"},
		LooseLeaf:                     {"Loose-leaf", "Loose-leaf book.", "Løse ark", "Løse ark eller perm med ark"},
		SpiralBound:                   {"Spiral bound", "Spiral, comb or coil bound book.", "Spiralbundet", "Bok med spiralrygg"},
		Pamphlet:                      {"Pamphlet", "Pamphlet or brochure, stapled; German ‘geheftet’.", "Stiftet", "Stiftet hefte eller brosjyre, pamflett"},
		LeatherFineBinding:            {"Leather / fine binding", "", "Praktinnbinding", "Skinninnbinding eller eksklusiv innbinding"},
		BoardBook:                     {"Board book", "Child’s book with all pages printed on board.", "Pappbok", "Kartonert bok, som oftest for barn"},
		RagBook:                       {"Rag book", "Child’s book with all pages printed on textile.", "Tekstilbok", "Barnebok trykket på stoff (tekstil)"},
		BathBook:                      {"Bath book", "Child’s book printed on waterproof material.", "Badebok", "Barnebok trykket på vannbestandig materiale"},
		NoveltyBook:                   {"Novelty book", "A book whose novelty consists wholly or partly in a format which cannot be described by any other available code – a ‘conventional’ format code is always to be preferred; one or more Product Form Detail codes, eg from the B2nn group, should be used whenever possible to provide additional description.", "Eksperimentell innbinding", "Bøker med spesiell/eksperimentell innbinding. Brukes om innbindingen, ikke innholdet"},
		SlideBound:                    {"Slide bound", "Slide bound book.", "Fleksibind", "Plastbelagt mykbind / eng: slide bound"},
		BigBook:                       {"Big book", "Extra-large format for teaching etc; this format and terminology may be specifically UK; required as a top-level differentiator.", "Tavlebok", "Ekstra stort format for undervisning etc, dette formatet og terminologien er muligens spesifikk for UK"},
		PartWorkFascículo:             {"Part-work (fascículo)", "A part-work issued with its own ISBN and intended to be collected and bound into a complete book.", "Utdrag, del av større verk", "Del av en annen utgivelse som er tildelt eget ISBN og er ment som en komplett bok"},
		FoldOutBookOrChart:            {"Fold-out book or chart", "Concertina-folded book or chart, designed to fold to pocket or regular page size: use for German ‘Leporello’.", "Leporello (brettet)", "En concertina-brettet bok. Vanligvis en bildebok"},
		FoamBook:                      {"Foam book", "A children’s book whose cover and pages are made of foam.", "Skumgummibok", "Barnebok trykket på skumgummi"},
		OtherBookFormat:               {"Other book format", "Other book format or binding not specified by BB to BO.", "Annet bokformat", "Annet bokformat eller innbinding, ikke spesifisert i BB til BO"},
		SheetMap:                      {"Sheet map", "Sheet map – detail unspecified.", "Kart, uspesifisert", "Kart - detaljer uspesifisert"},
		SheetMapFolded:                {"Sheet map, folded", "", "Kart, falset", "Kart som er falset (brettet)"},
		SheetMapFlat:                  {"Sheet map, flat", "", "Kart, plano", "Kart som ikke er brettet, men flatt (plano)"},
		SheetMapRolled:                {"Sheet map, rolled", "See Code List 80 for ‘rolled in tube’.", "Kart, rullet", "Se liste 80 for kart som er  ‘rullet i et rør’."},
		Globe:                         {"Globe", "Globe or planisphere.", "Globus", "Globus eller 'planisphere'"},
		OtherCartographic:             {"Other cartographic", "Other cartographic format not specified by CB to CE.", "Annet kartografisk format", "Annet kartografisk format ikke spesifosert i CB til CE"},
		DigitalOnPhysicalCarrier:      {"Digital (on physical carrier)", "Digital content delivered on a physical carrier (detail unspecified).", "Digitalt format (på fysisk bærer)", "Digitalt innhold levert på en fysisk bærer (detaljer uspesifisert)"},
		CDROM:         {"CD-ROM", "", "CD-ROM", ""},
		CDI:           {"CD-I", "CD interactive: use for ‘green book’ discs.", "CD-I", "CD interaktiv"},
		GameCartridge: {"Game cartridge", "", "Game cartridge", "Spillkassett til f.eks. Game Boy"},
		Diskette:      {"Diskette", "AKA ‘floppy disc’.", "Diskett", "Diskett for datamaskin, tidligere kalt floppy disk"},
		DVDROM:        {"DVD-ROM", "", "DVD-ROM", ""},
		SecureDigitalSDMemoryCard:      {"Secure Digital (SD) Memory Card", "", "Secure Digital (SD) (minnekort)", ""},
		CompactFlashMemoryCard:         {"Compact Flash Memory Card", "", "Compact Flash (minnekort)", ""},
		MemoryStickMemoryCard:          {"Memory Stick Memory Card", "", "Memory Stick (minnekort)", ""},
		USBFlashDrive:                  {"USB Flash Drive", "", "USB Flash Drive", ""},
		DoubleSidedCDDVD:               {"Double-sided CD/DVD", "Double-sided disc, one side Audio CD/CD-ROM, other side DVD.", "Tosidig CD/DVD", "Tosidig disk, en side CD/CD-ROM, andre siden DVD"},
		OtherDigitalCarrier:            {"Other digital carrier", "Other carrier of digital content not specified by DB to DN.", "Annet digitalt format (på fysisk bærer)", "Annet digitalt format ikke spesifisert i DB til DN"},
		DigitalDeliveredElectronically: {"Digital (delivered electronically)", "Digital content delivered electronically (delivery method unspecified).", "Digitalt format (leveres elektronisk)", "Digitalt innhold som leveres elektronisk (leveringsmetode er ikke spesifisert)"},
		DigitalDownloadAndOnline:       {"Digital download and online", "Digital content available both by download and by online access.", "Digitalt format, nedlastbart og nettbasert (online)", "Digitalt innhold som er tilgjengelig både via nedlasting og ved online tilgang"},
		DigitalOnline:                  {"Digital online", "Digital content accessed online only.", "Digitalt format, nettbasert (online)", "Digitalt innhold som kun leveres online"},
		DigitalDownload:                {"Digital download", "Digital content delivered by download only.", "Digitalt format, nedlastbart (download)", "Digitalt innhold som kun leveres via nedlasting"},
		FilmOrTransparency:             {"Film or transparency", "Film or transparency – detail unspecified.", "Film eller lysark", "Film eller transparent for overhead - uspesifisert"},
		Slides:                         {"Slides", "Photographic transparencies mounted for projection.", "Lysbilder", "Fotografiske lysbilder montert for fremvisning"},
		OHPTransparencies:              {"OHP transparencies", "Transparencies for overhead projector.", "Lysark", "Transparenter/lysark til overhead prosjektor"},
		Filmstrip:                      {"Filmstrip", "", "Filmremse", "eng: filmstrip"},
		Film:                           {"Film", "Continuous movie film as opposed to filmstrip.", "Film", "Sammenhengende film i motsetning til kode FE"},
		OtherFilmOrTransparencyFormat:          {"Other film or transparency format", "Other film or transparency format not specified by FB to FF.", "Annet film- eller lysarkformat", "Annet film- eller lysarkformat ikke spesifisert i FB til FF"},
		DigitalProductLicense:                  {"Digital product license", "Digital product license (delivery method not encoded).", "Digitalt produkt, lisens", "Lisens for digitalt produkt (leveringsmåte er ikke angitt)"},
		DigitalProductLicenseKey:               {"Digital product license key", "Digital product license delivered through the retail supply chain as a physical “key”, typically a card or booklet containing a code enabling the purchaser to download the associated product.", "Digitalt produkt, lisensnøkkel", "Lisensnøkkel for digitalt produkt, levert som en fysisk 'nøkkel', som oftest et kort eller hefte som inneholder en kode som gjør at kunden kan laste ned det aktuelle produktet"},
		DigitalProductLicenseCode:              {"Digital product license code", "Digital product license delivered by email or other electronic distribution, typically providing a code enabling the purchaser to upgrade or extend the license supplied with the associated product.", "Digitalt produkt, lisenskode", "Lisenskode for digitalt produkt, levert via e-post eller annen elektronisk distribusjon. Koden gjør at kunden kan oppgradere eller utvide lisensen som hører til det aktuelle produktet."},
		Microform:                              {"Microform", "Microform – detail unspecified.", "Mikroform", "Mikroform - uspesifisert"},
		Microfiche:                             {"Microfiche", "", "Mikrofiche", "Mikrofilmkort"},
		Microfilm:                              {"Microfilm", "Roll microfilm.", "Mikrofilm", "Mikrofilm, fortløpende på rull"},
		OtherMicroform:                         {"Other microform", "Other microform not specified by MB or MC.", "Andre mikroformat", "Andre mikroformer, ikke spesifisert i MB eller MC"},
		MiscellaneousPrint:                     {"Miscellaneous print", "Miscellaneous printed material – detail unspecified.", "Diverse trykk", "Diverse trykket materiale, uspesifisert"},
		AddressBook:                            {"Address book", "May use product form detail codes P201 to P204 to specify binding.", "Adressebok", "Bok beregnet på å skrive inn adresser. Man kan bruke kodene P201-P204 i product form detail for å spesifisere innbinding"},
		Calendar:                               {"Calendar", "", "Kalender", ""},
		Cards:                                  {"Cards", "Cards, flash cards (eg for teaching reading).", "Undervisningskort", "Kort til diverse undervisningsformål"},
		Copymasters:                            {"Copymasters", "Copymasters, photocopiable sheets.", "Kopieringsoriginaler", "Orginal beregnet på kopiering. Til undervisningsformål"},
		Diary:                                  {"Diary", "May use product form detail codes P201 to P204 to specify binding.", "Dagbok", "Man kan bruke kodene P201-P204 i product form detail for å spesifisere innbinding"},
		Frieze:                                 {"Frieze", "Narrow strip-shaped printed sheet used mostly for education or children’s products (eg depicting alphabet, number line, procession of illustrated characters etc). Usually intended for horizontal display.", "Frise", "Eng: frieze. "},
		Kit:                                    {"Kit", "Parts for post-purchase assembly.", "Eske/samlesett", "Utstyrssett (eng: kit)"},
		SheetMusic:                             {"Sheet music", "", "Notetrykk", "Noter for musikk trykket på ark"},
		PostcardBookOrPack:                     {"Postcard book or pack", "", "Postkort", "Postkort i pakker"},
		Poster:                                 {"Poster", "Poster for retail sale – see also XF.", "Plakat", "Plakat for salg - se også XF"},
		RecordBook:                             {"Record book", "Record book (eg ‘birthday book’, ‘baby book’): binding unspecified; may use product form detail codes P201 to P204 to specify binding.", "Minnebok", "Minnebok (f.eks. 'Fødselsdagsbok', 'baby bok'). Man kan bruke kodene P201-P204 i product form detail for å spesifisere innbinding"},
		WalletOrFolder:                         {"Wallet or folder", "Wallet or folder (containing loose sheets etc): it is preferable to code the contents and treat ‘wallet’ as packaging (List 80), but if this is not possible the product as a whole may be coded as a ‘wallet’.", "Mappe", "Mappe som inneholder løsblad, f.eks undervisningsmateriell"},
		PicturesOrPhotographs:                  {"Pictures or photographs", "", "Bilder eller fotografier", ""},
		Wallchart:                              {"Wallchart", "", "Veggplansje", ""},
		Stickers:                               {"Stickers", "", "Klistremerker", ""},
		PlateLámina:                            {"Plate (lámina)", "A book-sized (as opposed to poster-sized) sheet, usually in colour or high quality print.", "Ark", "Ark i bokstørrelse i motsetning til poster, vanligvis i farger eller høykvalitets trykk (eng: plate, lámina)"},
		NotebookBlankBook:                      {"Notebook / blank book", "A book with all pages blank for the buyer’s own use; may use product form detail codes P201 to P204 to specify binding.", "Notisbok / skrivebok", "Bok som kun har blanke sider.  Man kan bruke kodene P201-P204 i product form detail for å spesifisere innbinding"},
		Organizer:                              {"Organizer", "May use product form detail codes P201 to P204 to specify binding.", "Syvende sans", "Man kan bruke kodene P201-P204 i product form detail for å spesifisere innbinding"},
		Bookmark:                               {"Bookmark", "", "Bokmerke", ""},
		OtherPrintedItem:                       {"Other printed item", "Other printed item not specified by PB to PQ.", "Andre trykte artikler", "Andre trykte artikler, ikke spesifisert i PB til PQ."},
		MultipleItemRetailProduct:              {"Multiple-item retail product", "Presentation unspecified: format of product items must be given in <ProductPart>.", "Produkt bestående av flere enkeltprodukter", "Uspesifisert. Enkeltproduktenes format angis i <ProductPart>."},
		MultipleItemRetailProductBoxed:         {"Multiple-item retail product, boxed", "Format of product items must be given in <ProductPart>.", "Produkt bestående av flere enkeltprodukter, i eske", "Enkeltproduktenes format angis i <ProductPart>."},
		MultipleItemRetailProductSlipCased:     {"Multiple-item retail product, slip-cased", "Format of product items must be given in <ProductPart>.", "Produkt bestående av flere enkeltprodukter, i kassett", "Enkeltproduktenes format angis i <ProductPart>."},
		MultipleItemRetailProductShrinkwrapped: {"Multiple-item retail product, shrinkwrapped", "Format of product items must be given in <ProductPart>. Use code XL for a shrink-wrapped pack for trade supply, where the retail items it contains are intended for sale individually.", "Produkt bestående av flere enkeltprodukter, plastpakket", "Enkeltproduktenes format angis i <ProductPart>. Bruk kode XL for plastpakkede pakker til forhandler, hvor enkeltproduktene er ment å selges enkeltvis"},
		MultipleItemRetailProductLoose:         {"Multiple-item retail product, loose", "Format of product items must be given in <ProductPart>.", "Produkt bestående av flere enkeltprodukter, ikke emballert", "Enkeltproduktenes format angis i <ProductPart>."},
		MultipleItemRetailProductPartSEnclosed: {"Multiple-item retail product, part(s) enclosed", "Multiple item product where subsidiary product part(s) is/are supplied as enclosures to the primary part, eg a book with a CD packaged in a sleeve glued within the back cover. Format of product items must be given in <ProductPart>.", "Produkt bestående av flere enkeltprodukter, del(er) er lagt ved", "Produkt bestående av enkeltprodukter hvor enkeltprodukt(er) er lagt ved som en del av produktets hoveddel, for eksempel ei bok med en CD som er lagt ved i ei lomme som er limt fast til bokas omslag. Enkeltproduktenes format angis i <ProductPart>."},
		Video:              {"Video", "Video – detail unspecified.", "Video", "Video – detaljer uspesifisert"},
		Videodisc:          {"Videodisc", "eg Laserdisc.", "Videodisk", "f.eks. laserdisk"},
		DVDVideo:           {"DVD video", "DVD video: specify TV standard in List 175.", "DVD video", "DVD video: spesifiser TV-standard i liste 175."},
		VHSVideo:           {"VHS video", "VHS videotape: specify TV standard in List 175.", "VHS video", "VHS videotape: spesifiser TV-standard i liste 175."},
		BetamaxVideo:       {"Betamax video", "Betamax videotape: specify TV standard in List 175.", "Betamax video", "Betamax videotape: spesifiser TV-standard i liste 175."},
		VCD:                {"VCD", "VideoCD.", "VCD", "VideoCD."},
		SVCD:               {"SVCD", "Super VideoCD.", "SVCD", "Super VideoCD."},
		HDDVD:              {"HD DVD", "High definition DVD disc, Toshiba HD DVD format.", "HD DVD", "Høyoppløsnings DVD-disk, HD DVD format"},
		BluRay:             {"Blu-ray", "High definition DVD disc, Sony Blu-ray format.", "Blu-ray", "Høyoppløsnings DVD-disk, Sony Blu ray-format"},
		UMDVideo:           {"UMD Video", "Sony Universal Media disc.", "UMD Video", "Sony Universal Media disc."},
		CBHD:               {"CBHD", "China Blue High-Definition, derivative of HD-DVD.", "CBHD", "China Blue High-Definition, derivative of HD-DVD."},
		OtherVideoFormat:   {"Other video format", "Other video format not specified by VB to VP.", "Andre videoformater", "Andre videoformater ikke spesifisert i VB til VP"},
		TradeOnlyMaterial:  {"Trade-only material", "Trade-only material (unspecified).", "Markedsføringsmateriell", "Uspesifisert maskedsførings-/reklamemateriell"},
		DumpbinEmpty:       {"Dumpbin – empty", "", "Tilbudskasse - tom", "Tilbudskasse uten innhold (eng: dumpbin)"},
		DumpbinFilled:      {"Dumpbin – filled", "Dumpbin with contents. ISBN (where applicable) and format of contained items must be given in Product Part.", "Tilbudskasse - fylt", "Tilbudskasse med innhold. ISBN og format for innholdselementene kan oppgis i Product part"},
		CounterpackEmpty:   {"Counterpack – empty", "", "Diskeske - tom", ""},
		CounterpackFilled:  {"Counterpack – filled", "Counterpack with contents. ISBN (where applicable) and format of contained items must be given in Product Part.", "Diskeske- fylt", "Diskeske med innhold. ISBN og format for innholdselementene kan oppgis i Product part"},
		PosterPromotional:  {"Poster, promotional", "Promotional poster for display, not for sale – see also PK.", "Reklameplakat", "Reklameplakat, ikke for salg. Se også PK"},
		ShelfStrip:         {"Shelf strip", "", "Hyllemarkør", ""},
		WindowPiece:        {"Window piece", "Promotional piece for shop window display.", "Vindusmateriell", "Reklamemateriell for utstillingsvindu"},
		Streamer:           {"Streamer", "", "Streamer", "Avlang reklameplakat"},
		Spinner:            {"Spinner", "", "Gulvstativ", "Display for utstilling av bøker, kart oa. På eng. begrenset til roterende stativ."},
		LargeBookDisplay:   {"Large book display", "Large scale facsimile of book for promotional display.", "Forstørret omslag", "Bokomslag i plakatform for reklameformål"},
		ShrinkWrappedPack:  {"Shrink-wrapped pack", "A quantity pack with its own product code, for trade supply only: the retail items it contains are intended for sale individually. ISBN (where applicable) and format of contained items must be given in Product Part. For products or product bundles supplied shrink-wrapped for retail sale, use code SD.", "Plastpakket", "En samlepakke med eget produktnummer. Kun for bruk i handelssammenheng. Innholdet er ment å kunne selges enkeltvis. ISBN og format for pakkas enkeltelementer kan oppgis i Product part. For produkter eller produktpakker som leveres for salg til sluttkunde, bruk kode SD."},
		OtherPointOfSale:   {"Other point of sale", "Other point of sale material not specified by XB to XL.", "Annet reklamemateriell", "Annet salgsmateriell som ikke er spesifisert i XB-XL"},
		GeneralMerchandise: {"General merchandise", "General merchandise – unspecified.", "Varer, uspesifisert", "Generelle varer, uspesifisert"},
		Doll:               {"Doll", "", "Dukke", ""},
		SoftToy:            {"Soft toy", "Soft or plush toy.", "Myk leke", "Kosebamser o.a.; myke leker"},
		Toy:                {"Toy", "", "Leke", ""},
		Game:               {"Game", "Board game, or other game (except computer game: see DE).", "Spill", "Spill eller brettspill (ikke dataspill, se DE)"},
		TShirt:             {"T-shirt", "", "T-skjorte", ""},
		EBookReader:        {"E-book reader", "Dedicated e-book reading device, typically with mono screen.", "Lesebrett", "Lesebrett beregnet kun på å lese bøker"},
		TabletComputer:     {"Tablet computer", "General purpose tablet computer, typically with color screen.", "Nettbrett", "General purpose tablet computer, typically with color screen."},
		AudiobookPlayer:    {"Audiobook player", "Dedicated audiobook player device, typically including book-related features like bookmarking", "Lydbokspiller", "Lydbokspiller beregnet kun på avspilling av lydbøker"},
		Jigsaw:             {"Jigsaw", "", "Puslespill", ""},
		OtherApparel:       {"Other apparel", "Other apparel items not specifieed by ZB-ZJ, including promotional or branded scarves, caps, aprons etc.", "Andre klær", "Andre typer klær og produkter som ikke er spesifisert av kodene ZB-ZJ, inkludert markedsføringsprodukter som skjerf, luer, forklær osv."},
		OtherMerchandise:   {"Other merchandise", "Other merchandise not specified by ZB to ZK.", "Andre varer", "Andre varer som ikke er spesifisert i ZB-ZK."},
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
