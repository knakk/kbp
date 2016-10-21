package list96

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	UAEDirham                   = "AED"
	Afghani                     = "AFA"
	Afghani                     = "AFN"
	Lek                         = "ALL"
	ArmenianDram                = "AMD"
	NetherlandsAntillianGuilder = "ANG"
	AngolanKwanza               = "AOA"
	ArgentinePeso               = "ARS"
	AustriaSchilling            = "ATS"
	AustralianDollar            = "AUD"
	ArubanFlorin                = "AWG"
	AzerbaijanianManat          = "AZN"
	ConvertibleMarks            = "BAM"
	BarbadosDollar              = "BBD"
	Taka                        = "BDT"
	BelgiumFranc                = "BEF"
	Lev                         = "BGL"
	Lev                         = "BGN"
	BahrainiDinar               = "BHD"
	BurundiFranc                = "BIF"
	BermudaDollar               = "BMD"
	BruneiDollar                = "BND"
	Boliviano                   = "BOB"
	BrazilianReal               = "BRL"
	BahamianDollar              = "BSD"
	Ngultrun                    = "BTN"
	Pula                        = "BWP"
	BelarussianRuble            = "BYR"
	BelizeDollar                = "BZD"
	CanadianDollar              = "CAD"
	FrancCongolais              = "CDF"
	SwissFranc                  = "CHF"
	ChileanPeso                 = "CLP"
	YuanRenminbi                = "CNY"
	ColombianPeso               = "COP"
	CostaRicanColon             = "CRC"
	SerbianDinar                = "CSD"
	CubanConvertiblePeso        = "CUC"
	CubanPeso                   = "CUP"
	CapeVerdeEscudo             = "CVE"
	CyprusPound                 = "CYP"
	CzechKoruna                 = "CZK"
	GermanyMark                 = "DEM"
	DjiboutiFranc               = "DJF"
	DanishKrone                 = "DKK"
	DominicanPeso               = "DOP"
	AlgerianDinar               = "DZD"
	Kroon                       = "EEK"
	EgyptianPound               = "EGP"
	Nakfa                       = "ERN"
	SpainPeseta                 = "ESP"
	EthiopianBirr               = "ETB"
	Euro                        = "EUR"
	FinlandMarkka               = "FIM"
	FijiDollar                  = "FJD"
	FalklandIslandsPound        = "FKP"
	FranceFranc                 = "FRF"
	PoundSterling               = "GBP"
	Lari                        = "GEL"
	Cedi                        = "GHC"
	Cedi                        = "GHS"
	GibraltarPound              = "GIP"
	Dalasi                      = "GMD"
	GuineaFranc                 = "GNF"
	GreeceDrachma               = "GRD"
	Quetzal                     = "GTQ"
	GuineaBissauPeso            = "GWP"
	GuyanaDollar                = "GYD"
	HongKongDollar              = "HKD"
	Lempira                     = "HNL"
	CroatianKuna                = "HRK"
	Gourde                      = "HTG"
	Forint                      = "HUF"
	Rupiah                      = "IDR"
	IrelandPunt                 = "IEP"
	IsraeliSheqel               = "ILS"
	IndianRupee                 = "INR"
	IraqiDinar                  = "IQD"
	IranianRial                 = "IRR"
	IcelandKrona                = "ISK"
	ItalyLira                   = "ITL"
	JamaicanDollar              = "JMD"
	JordanianDinar              = "JOD"
	Yen                         = "JPY"
	KenyanShilling              = "KES"
	Som                         = "KGS"
	Riel                        = "KHR"
	ComoroFranc                 = "KMF"
	NorthKoreanWon              = "KPW"
	Won                         = "KRW"
	KuwaitiDinar                = "KWD"
	CaymanIslandsDollar         = "KYD"
	Tenge                       = "KZT"
	Kip                         = "LAK"
	LebanesePound               = "LBP"
	SriLankaRupee               = "LKR"
	LiberianDollar              = "LRD"
	Loti                        = "LSL"
	LithuanianLitus             = "LTL"
	LuxembourgFranc             = "LUF"
	LatvianLats                 = "LVL"
	LibyanDinar                 = "LYD"
	MoroccanDirham              = "MAD"
	MoldovanLeu                 = "MDL"
	Ariary                      = "MGA"
	MalagasyFranc               = "MGF"
	Denar                       = "MKD"
	Kyat                        = "MMK"
	Tugrik                      = "MNT"
	Pataca                      = "MOP"
	Ouguiya                     = "MRO"
	MalteseLira                 = "MTL"
	MauritiusRupee              = "MUR"
	Rufiyaa                     = "MVR"
	Kwacha                      = "MWK"
	MexicanPeso                 = "MXN"
	MalaysianRinggit            = "MYR"
	Metical                     = "MZN"
	NamibiaDollar               = "NAD"
	Naira                       = "NGN"
	CordobaOro                  = "NIO"
	NetherlandsGuilder          = "NLG"
	NorwegianKrone              = "NOK"
	NepaleseRupee               = "NPR"
	NewZealandDollar            = "NZD"
	RialOmani                   = "OMR"
	Balboa                      = "PAB"
	NuevoSol                    = "PEN"
	Kina                        = "PGK"
	PhilippinePeso              = "PHP"
	PakistanRupee               = "PKR"
	Zloty                       = "PLN"
	PortugalEscudo              = "PTE"
	Guarani                     = "PYG"
	QatariRial                  = "QAR"
	OldLeu                      = "ROL"
	NewLeu                      = "RON"
	SerbianDinar                = "RSD"
	RussianRuble                = "RUB"
	RussianRuble                = "RUR"
	RwandaFranc                 = "RWF"
	SaudiRiyal                  = "SAR"
	SolomonIslandsDollar        = "SBD"
	SeychellesRupee             = "SCR"
	SudaneseDinar               = "SDD"
	SudanesePound               = "SDG"
	SwedishKrona                = "SEK"
	SingaporeDollar             = "SGD"
	SaintHelenaPound            = "SHP"
	Tolar                       = "SIT"
	SlovakKoruna                = "SKK"
	Leone                       = "SLL"
	SomaliShilling              = "SOS"
	SurinamDollar               = "SRD"
	SurinameGuilder             = "SRG"
	Dobra                       = "STD"
	ElSalvadorColon             = "SVC"
	SyrianPound                 = "SYP"
	Lilangeni                   = "SZL"
	Baht                        = "THB"
	Somoni                      = "TJS"
	Manat                       = "TMM"
	Manat                       = "TMT"
	TunisianDinar               = "TND"
	Paanga                      = "TOP"
	TimorEscudo                 = "TPE"
	TurkishLiraOld              = "TRL"
	TurkishLiraNew              = "TRY"
	TrinidadAndTobagoDollar     = "TTD"
	NewTaiwanDollar             = "TWD"
	TanzanianShilling           = "TZS"
	Hryvnia                     = "UAH"
	UgandaShilling              = "UGX"
	USDollar                    = "USD"
	PesoUruguayo                = "UYU"
	UzbekistanSum               = "UZS"
	Bolivar                     = "VEB"
	BolivarFuerte               = "VEF"
	Dong                        = "VND"
	Vatu                        = "VUV"
	Tala                        = "WST"
	CFAFrancBEAC                = "XAF"
	EastCaribbeanDollar         = "XCD"
	CFAFrancBCEAO               = "XOF"
	CFPFranc                    = "XPF"
	YemeniRial                  = "YER"
	YugoslavianDinar            = "YUM"
	Rand                        = "ZAR"
	Kwacha                      = "ZMK"
	ZambianKwacha               = "ZMW"
	ZimbabweDollar              = "ZWD"
	ZimbabweDollar              = "ZWL"
)

var (
	sortedCodes = []string{"AED", "AFA", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "ATS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BEF", "BGL", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN", "BWP", "BYR", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CSD", "CUC", "CUP", "CVE", "CYP", "CZK", "DEM", "DJF", "DKK", "DOP", "DZD", "EEK", "EGP", "ERN", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GRD", "GTQ", "GWP", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "IEP", "ILS", "INR", "IQD", "IRR", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LUF", "LVL", "LYD", "MAD", "MDL", "MGA", "MGF", "MKD", "MMK", "MNT", "MOP", "MRO", "MTL", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PTE", "PYG", "QAR", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "STD", "SVC", "SYP", "SZL", "THB", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VEB", "VEF", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "YUM", "ZAR", "ZMK", "ZMW", "ZWD", "ZWL"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		UAEDirham:                   {"UAE Dirham", "United Arab Emirates.", "Emiratarabisk dirham", "United Arab Emirates"},
		Afghani:                     {"Afghani", "DEPRECATED, replaced by AFN.", "Afghani", "UTGÅTT, erstattet av AFN"},
		Afghani:                     {"Afghani", "Afghanistan.", "Afghani", "Afghanistan"},
		Lek:                         {"Lek", "Albania.", "Albanske lek", "Albania"},
		ArmenianDram:                {"Armenian Dram", "Armenia.", "Armenske dram", "Armenia"},
		NetherlandsAntillianGuilder: {"Netherlands Antillian Guilder", "Curaçao, Sint Maarten.", "Antilliansk gylden", "Nederlandske antiller"},
		AngolanKwanza:               {"Angolan Kwanza", "Angola.", "Angolansk kwanza", "Angola"},
		ArgentinePeso:               {"Argentine Peso", "Argentina.", "Argentinsk peso", "Argentina"},
		AustriaSchilling:            {"Austria, Schilling", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Østerikske schilling", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		AustralianDollar:            {"Australian Dollar", "Australia, Christmas Island, Cocos (Keeling) Islands, Heard Island and McDonald Islands, Kiribati, Nauru, Norfolk Island, Tuvalu.", "Australske dollar", "Australia, Christmas Island, Cocos (Keeling) Islands, Heard Island and McDonald Islands, Kiribati, Nauru, Norfolk Island, Tuvalu"},
		ArubanFlorin:                {"Aruban Florin", "Aruba.", "Arubansk florin", "Aruba"},
		AzerbaijanianManat:          {"Azerbaijanian Manat", "Azerbaijan.", "Aserbajdsjansk manat", "Aserbajdsjan"},
		ConvertibleMarks:            {"Convertible Marks", "Bosnia and Herzegovina.", "Konvertibel mark", "Bosnia og Herzegovina"},
		BarbadosDollar:              {"Barbados Dollar", "Barbados.", "Barbadisk dollar", "Barbados"},
		Taka:                        {"Taka", "Bangladesh.", "Bangladeshisk taka", "Bangladesh"},
		BelgiumFranc:                {"Belgium, Franc", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Belgiske franc", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		Lev:                         {"Lev", "DEPRECATED, replaced by BGN.", "Bulgarske lev", "UTGÅTT, erstattet av BGN"},
		Lev:                         {"Lev", "Bulgaria.", "Bulgarske lev", "Bulgaria"},
		BahrainiDinar:               {"Bahraini Dinar", "Bahrain.", "Bahrainsk dinar", "Bahrain"},
		BurundiFranc:                {"Burundi Franc", "Burundi.", "Burundisk franc", "Burundi"},
		BermudaDollar:               {"Bermuda Dollar", "Bermuda.", "Bermudisk dollar", "Bermuda"},
		BruneiDollar:                {"Brunei Dollar", "Brunei Darussalam.", "Bruneisk dollar", "Brunei"},
		Boliviano:                   {"Boliviano", "Bolivia.", "Boliviano", "Bolivia"},
		BrazilianReal:               {"Brazilian Real", "Brazil.", "Brasiliansk real", "Brasil"},
		BahamianDollar:              {"Bahamian Dollar", "Bahamas.", "Bahamansk dollar", "Bahamas"},
		Ngultrun:                    {"Ngultrun", "Bhutan.", "Ngultrum", "Bhutan"},
		Pula:                        {"Pula", "Botswana.", "Pula", "Botswana"},
		BelarussianRuble:            {"Belarussian Ruble", "Belarus.", "Hviterussiske rubel", "Hviterussland"},
		BelizeDollar:                {"Belize Dollar", "Belize.", "Belizisk dollar", "Belize"},
		CanadianDollar:              {"Canadian Dollar", "Canada.", "Kanadisk dollar", "Canada"},
		FrancCongolais:              {"Franc Congolais", "Congo (Democratic Republic of the).", "Kongolesiske Franc", "Kongo, Den demokratiske rebublikken"},
		SwissFranc:                  {"Swiss Franc", "Switzerland, Liechtenstein.", "Sveitserfranc", "Sveits, Liechtenstein"},
		ChileanPeso:                 {"Chilean Peso", "Chile.", "Chilensk peso", "Chile"},
		YuanRenminbi:                {"Yuan Renminbi", "China.", "Yuan Renminbi", "Kina"},
		ColombianPeso:               {"Colombian Peso", "Colombia.", "Colombiansk peso", "Colombia"},
		CostaRicanColon:             {"Costa Rican Colon", "Costa Rica.", "Costaricansk colón", "Costa Rica"},
		SerbianDinar:                {"Serbian Dinar", "Deprecated, replaced by RSD.", "Serbisk dinar", "UTGÅTT. Erstattet av RSD"},
		CubanConvertiblePeso:        {"Cuban Convertible Peso", "Cuba (alternative currency).", "Cubansk konvertibel peso", "Cuba (alternativ valuta)"},
		CubanPeso:                   {"Cuban Peso", "Cuba.", "Cubansk peso", "Cuba"},
		CapeVerdeEscudo:             {"Cape Verde Escudo", "Cape Verde.", "Kappverdisk escudo", "Cape Verde"},
		CyprusPound:                 {"Cyprus Pound", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Kypriotisk pund", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		CzechKoruna:                 {"Czech Koruna", "Czech Republic.", "Tsjekkisk koruna", "Tsjekkia"},
		GermanyMark:                 {"Germany, Mark", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Tyske mark", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		DjiboutiFranc:               {"Djibouti Franc", "Djibouti.", "Djiboutisk franc", "Djibouti"},
		DanishKrone:                 {"Danish Krone", "Denmark, Faroe Islands, Greenland.", "Danske kroner", "Danmark, Grønnland og Færøyene"},
		DominicanPeso:               {"Dominican Peso", "Dominican Republic.", "Dominikansk peso", "Den dominikanske republikk "},
		AlgerianDinar:               {"Algerian Dinar", "Algeria.", "Algerisk dinar", "Algerie"},
		Kroon:                       {"Kroon", "Estonia – now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Kroon", "Estland"},
		EgyptianPound:               {"Egyptian Pound", "Egypt.", "Egyptisk pund\n", "Egypt"},
		Nakfa:                       {"Nakfa", "Eritrea.", "Eritreisk nakfa", "Eritrea"},
		SpainPeseta:                 {"Spain, Peseta", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Spansk peseta", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		EthiopianBirr:               {"Ethiopian Birr", "Ethiopia.", "Etiopiske birr", "Etiopia"},
		Euro:                        {"Euro", "Andorra, Austria, Belgium, Cyprus, Estonia, Finland, France, Fr Guiana, Fr S Territories, Germany, Greece, Guadeloupe, Holy See (Vatican City), Ireland, Italy, Luxembourg, Martinique, Malta, Mayotte, Monaco, Montenegro, Netherlands, Portugal, Réunion, St Pierre and Miquelon, San Marino, Spain.", "Euro", "Andorra, Austria, Belgium, Cyprus, Estonia, Finland, France, Fr Guiana, Fr S Territories, Germany, Greece, Guadeloupe, Holy See (Vatican City), Ireland, Italy, Luxembourg, Martinique, Malta, Mayotte, Monaco, Montenegro, Netherlands, Portugal, Réunion, St Pierre and Miquelon, San Marino, Spain."},
		FinlandMarkka:               {"Finland, Markka", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Finske mark", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		FijiDollar:                  {"Fiji Dollar", "Fiji.", "Fijiansk dollar", "Fiji"},
		FalklandIslandsPound:        {"Falkland Islands Pound", "Falkland Islands (Malvinas).", "Falklandspund", "Falklandsøyene (Las Malvinas)"},
		FranceFranc:                 {"France, Franc", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Franske franc", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		PoundSterling:               {"Pound Sterling", "United Kingdom, Isle of Man, Channel Islands, South Georgia, South Sandwich Islands, British Indian Ocean Territory.", "Britiske pund", "United Kingdom, Isle of Man, Channel Islands, South Georgia, South Sandwich Islands, British Indian Ocean Territory."},
		Lari:                        {"Lari", "Georgia.", "Lari", "Georgia"},
		Cedi:                        {"Cedi", "Deprecated, replaced by GHS.", "Ghanesisk cedi", "UTGÅTT. Erstattet av GHS"},
		Cedi:                        {"Cedi", "Ghana.", "Cedi", "Ghana"},
		GibraltarPound:              {"Gibraltar Pound", "Gibraltar.", "Gibraltarsk pund", "Gibraltar"},
		Dalasi:                      {"Dalasi", "Gambia.", "Gambisk dalasi", "Gambia"},
		GuineaFranc:                 {"Guinea Franc", "Guinea.", "Guineansk franc", "Guinea"},
		GreeceDrachma:               {"Greece, Drachma", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Greske drakmer", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		Quetzal:                     {"Quetzal", "Guatemala.", "Quetzal", "Guatemala"},
		GuineaBissauPeso:            {"Guinea-Bissau Peso", "Now replaced by the CFA Franc BCEAO XOF use only for historical prices that pre-date use of the CFA Franc.", "Guinea-Bissau Peso", "Guinea-Bissau"},
		GuyanaDollar:                {"Guyana Dollar", "Guyana.", "Guyansk dollar", "Guyana"},
		HongKongDollar:              {"Hong Kong Dollar", "Hong Kong, Macao.", "Hong Kong dollar", "Hong Kong, Macao"},
		Lempira:                     {"Lempira", "Honduras.", "Lempira", "Honduras"},
		CroatianKuna:                {"Croatian Kuna", "Croatia.", "Kroatisk kuna", "Croatia"},
		Gourde:                      {"Gourde", "Haiti.", "Haitisk gourde", "Haiti"},
		Forint:                      {"Forint", "Hungary.", "Forint", "Ungarn"},
		Rupiah:                      {"Rupiah", "Indonesia.", "Rupiah", "Indonesia"},
		IrelandPunt:                 {"Ireland, Punt", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Irske pund", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		IsraeliSheqel:               {"Israeli Sheqel", "Israel.", "Israelsk shekel", "Israel"},
		IndianRupee:                 {"Indian Rupee", "India.", "Indisk rupee", "India"},
		IraqiDinar:                  {"Iraqi Dinar", "Iraq.", "Irakisk dinar", "Iraq"},
		IranianRial:                 {"Iranian Rial", "Iran (Islamic Republic of).", "Iransk rial", "Iran"},
		IcelandKrona:                {"Iceland Krona", "Iceland.", "Islandsk krona", "Island"},
		ItalyLira:                   {"Italy, Lira", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Italienske lire", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		JamaicanDollar:              {"Jamaican Dollar", "Jamaica.", "Jamaicansk dollar", "Jamaica"},
		JordanianDinar:              {"Jordanian Dinar", "Jordan.", "Jordansk dinar", "Jordan"},
		Yen:                         {"Yen", "Japan.", "Yen", "Japan"},
		KenyanShilling:              {"Kenyan Shilling", "Kenya.", "Kenyansk shilling", "Kenya"},
		Som:                         {"Som", "Kyrgyzstan.", "Som", "Kirgisistan"},
		Riel:                        {"Riel", "Cambodia.", "Riel", "Kambodsja"},
		ComoroFranc:                 {"Comoro Franc", "Comoros.", "Komorisk franc", "Komorene"},
		NorthKoreanWon:              {"North Korean Won", "Korea (Democratic People’s Republic of).", "Nordkoreansk won", "Nord-Korea"},
		Won:                         {"Won", "Korea (Republic of).", "Sørkoreansk won", "Sør-Korea"},
		KuwaitiDinar:                {"Kuwaiti Dinar", "Kuwait.", "Kuwaitisk dinar", "Kuwait"},
		CaymanIslandsDollar:         {"Cayman Islands Dollar", "Cayman Islands.", "Caymansk dollar", "Caymanøyene"},
		Tenge:                       {"Tenge", "Kazakstan.", "Tenge", "Kasakhstan"},
		Kip:                         {"Kip", "Lao People’s Democratic Republic.", "Kip", "Laos"},
		LebanesePound:               {"Lebanese Pound", "Lebanon.", "Libanesisk pund", "Libanon"},
		SriLankaRupee:               {"Sri Lanka Rupee", "Sri Lanka.", "Srilankisk rupi", "Sri Lanka"},
		LiberianDollar:              {"Liberian Dollar", "Liberia.", "Liberiansk dollar", "Liberia"},
		Loti:                        {"Loti", "Lesotho.", "Loti", "Lesotho"},
		LithuanianLitus:             {"Lithuanian Litus", "Lithuania.", "Litauisk litas", "Litauen"},
		LuxembourgFranc:             {"Luxembourg, Franc", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Luxemburg franc", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		LatvianLats:                 {"Latvian Lats", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Latvisk lat", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		LibyanDinar:                 {"Libyan Dinar", "Libyan Arab Jamahiriya.", "Libysk dinar", "Libya"},
		MoroccanDirham:              {"Moroccan Dirham", "Morocco, Western Sahara.", "Marokkansk dirham", "Marokko, Vest Sahara"},
		MoldovanLeu:                 {"Moldovan Leu", "Moldova, Republic of.", "Moldovisk leu", "Moldova"},
		Ariary:                      {"Ariary", "Madagascar.", "Ariary", "Madagaskar"},
		MalagasyFranc:               {"Malagasy Franc", "Now replaced by the Ariary (MGA).", "Gassisk franc", "UTGÅTT. Erstattet av MGA"},
		Denar:                       {"Denar", "Macedonia (former Yugoslav Republic of).", "Makedonsk denar", "Makedonia (FYR)"},
		Kyat:                        {"Kyat", "Myanmar.", "Kyat", "Myanmar"},
		Tugrik:                      {"Tugrik", "Mongolia.", "Tugrik", "Mongolia"},
		Pataca:                      {"Pataca", "Macau.", "Pataca", "Macau"},
		Ouguiya:                     {"Ouguiya", "Mauritania.", "Ouguiya", "Mauritania"},
		MalteseLira:                 {"Maltese Lira", "Malta – now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Maltesisk lira", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		MauritiusRupee:              {"Mauritius Rupee", "Mauritius.", "Mauritisk rupi", "Mauritius"},
		Rufiyaa:                     {"Rufiyaa", "Maldives.", "Maldivisk rufiyaa", "Maldivene"},
		Kwacha:                      {"Kwacha", "Malawi.", "Malawisk kwacha", "Malawi"},
		MexicanPeso:                 {"Mexican Peso", "Mexico.", "Mexikansk peso", "Mexico       "},
		MalaysianRinggit:            {"Malaysian Ringgit", "Malaysia.", "Ringgit", "Malaysia"},
		Metical:                     {"Metical", "Mozambique.", "Metical", "Mosambik"},
		NamibiaDollar:               {"Namibia Dollar", "Namibia.", "Namibisk dollar", "Namibia"},
		Naira:                       {"Naira", "Nigeria.", "Naira", "Nigeria"},
		CordobaOro:                  {"Cordoba Oro", "Nicaragua.", "Córdoba", "Nicaragua"},
		NetherlandsGuilder:          {"Netherlands, Guilder", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Nederlandsk gylden", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		NorwegianKrone:              {"Norwegian Krone", "Norway, Bouvet Island, Svalbard and Jan Mayen.", "Norsk krone", "Norge, Svalbard og Jan Mayen"},
		NepaleseRupee:               {"Nepalese Rupee", "Nepal.", "Nepalsk rupi", "Nepal"},
		NewZealandDollar:            {"New Zealand Dollar", "New Zealand, Cook Islands, Niue, Pitcairn, Tokelau.", "Newzealandsk dollar", "New Zealand, Cook Islands, Niue, Pitcairn, Tokelau"},
		RialOmani:                   {"Rial Omani", "Oman.", "Omansk rial", "Oman"},
		Balboa:                      {"Balboa", "Panama.", "Balboa", "Panama"},
		NuevoSol:                    {"Nuevo Sol", "Peru.", "Peruansk nuevo sol", "Peru"},
		Kina:                        {"Kina", "Papua New Guinea.", "Kina", "Papua Ny-Guinea"},
		PhilippinePeso:              {"Philippine Peso", "Philippines.", "Filippinsk peso", "Filippinene"},
		PakistanRupee:               {"Pakistan Rupee", "Pakistan.", "Pakistansk rupi", "Pakistan"},
		Zloty:                       {"Zloty", "Poland.", "Polsk zloty", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		PortugalEscudo:              {"Portugal, Escudo", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Portugisiske escudo", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		Guarani:                     {"Guarani", "Paraguay.", "Guarani", "Paraguay"},
		QatariRial:                  {"Qatari Rial", "Qatar.", "Qatarsk riyal", "Qatar"},
		OldLeu:                      {"Old Leu", "Deprecated, replaced by RON.", "Rumensk leu", "UTGÅTT. Erstattet av RON"},
		NewLeu:                      {"New Leu", "Romania.", "Ny rumensk Leu", "Romania"},
		SerbianDinar:                {"Serbian Dinar", "Serbia.", "Serbisk dinar", "Serbia"},
		RussianRuble:                {"Russian Ruble", "Russian Federation.", "Russisk rubel", "Russland"},
		RussianRuble:                {"Russian Ruble", "DEPRECATED, replaced by RUB.", "Russisk rubel", "UTGÅTT, erstattet av RUB"},
		RwandaFranc:                 {"Rwanda Franc", "Rwanda.", "Rwandisk franc", "Rwanda"},
		SaudiRiyal:                  {"Saudi Riyal", "Saudi Arabia.", "Saudiarabisk riyal", "Saudi Arabia"},
		SolomonIslandsDollar:        {"Solomon Islands Dollar", "Solomon Islands.", "Salomonsk dollar", "Solomon øyene"},
		SeychellesRupee:             {"Seychelles Rupee", "Seychelles.", "Seychellisk rupi", "Seychelles"},
		SudaneseDinar:               {"Sudanese Dinar", "Now replaced by the Sudanese Pound (SDG).", "Sudanesisk dinar", "Erstattet av Sudanesisk pund (SDG)"},
		SudanesePound:               {"Sudanese Pound", "Sudan.", "Sudanesisk pund", "Sudan"},
		SwedishKrona:                {"Swedish Krona", "Sweden.", "Svensk krona", "Sverige"},
		SingaporeDollar:             {"Singapore Dollar", "Singapore.", "Singaporsk dollar", "Singapore"},
		SaintHelenaPound:            {"Saint Helena Pound", "Saint Helena.", "Sankthelensk\u00a0pund\u00a0", "St. Helena"},
		Tolar:                       {"Tolar", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Tolar", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		SlovakKoruna:                {"Slovak Koruna", "Now replaced by the Euro (EUR): use only for historical prices that pre-date the introduction of the Euro.", "Slovakisk koruna", "Erstattet av Euro (EUR). Brukes bare for historiske priser fra før Euro ble innført"},
		Leone:                       {"Leone", "Sierra Leone.", "Leone", "Sierra Leone"},
		SomaliShilling:              {"Somali Shilling", "Somalia.", "Somalisk shilling\n", "Somalia"},
		SurinamDollar:               {"Surinam Dollar", "Suriname.", "Surinamsk dollar", "Suriname"},
		SurinameGuilder:             {"Suriname Guilder", "DEPRECATED, replaced by SRD.", "Surinamsk gylden", "UTGÅTT, erstattet av SRD"},
		Dobra:                       {"Dobra", "São Tome and Principe.", "Dobra", "São Tome og Principe"},
		ElSalvadorColon:             {"El Salvador Colon", "El Salvador.", "Salvadoransk colon", "El Salvador"},
		SyrianPound:                 {"Syrian Pound", "Syrian Arab Republic.", "Syrisk pund", "Syria"},
		Lilangeni:                   {"Lilangeni", "Swaziland.", "Lilangeni", "Swaziland"},
		Baht:                        {"Baht", "Thailand.", "Baht", "Thailand"},
		Somoni:                      {"Somoni", "Tajikistan.", "Tadsjikisk somoni", "Tadsjikistan"},
		Manat:                       {"Manat", "Deprecated, replaced by TMT.", "Turkmensk manat", "UTGÅTT. Erstattet av TMT"},
		Manat:                       {"Manat", "Turkmenistan.", "Manat", "Turkmenistan"},
		TunisianDinar:               {"Tunisian Dinar", "Tunisia.", "Tunisisk dinar", "Tunisia"},
		Paanga:                      {"Pa’anga", "Tonga.", "Pa'anga", "Tonga"},
		TimorEscudo:                 {"Timor Escudo", "NO LONGER VALID, Timor-Leste now uses the US Dollar.", "Timor Escudo", "UTGÅTT OG IKKE LENGER GYLDIG. Øst-Timor bruker US dollar"},
		TurkishLiraOld:              {"Turkish Lira (old)", "Deprecated, replaced by TRY.", "Tyrkisk lira (gammel)", "UTGÅTT. Erstattet av TRY"},
		TurkishLiraNew:              {"Turkish Lira (new)", "Turkey, from 1 January 2005.", "Tyrkisk lira (ny)", "Tyrkia,fra 1 Januar 2005"},
		TrinidadAndTobagoDollar:     {"Trinidad and Tobago Dollar", "Trinidad and Tobago.", "Trinidad and Tobago-dollar", "Trinidad and Tobago"},
		NewTaiwanDollar:             {"New Taiwan Dollar", "Taiwan (Province of China).", "Taiwansk dollar", "Taiwan "},
		TanzanianShilling:           {"Tanzanian Shilling", "Tanzania (United Republic of).", "Tanzaniansk shilling", "Tanzania"},
		Hryvnia:                     {"Hryvnia", "Ukraine.", "Hryvnia", "Ukraina"},
		UgandaShilling:              {"Uganda Shilling", "Uganda.", "Uganda shilling", "Uganda"},
		USDollar:                    {"US Dollar", "United States, American Samoa, British Indian Ocean Territory, Ecuador, Guam, Marshall Is, Micronesia (Federated States of), Northern Mariana Is, Palau, Puerto Rico, Timor-Leste, Turks and Caicos Is, US Minor Outlying Is, Virgin Is (British), Virgin Is (US).", "Amerikanske dollar", "United States, American Samoa, British Indian Ocean Territory, Ecuador, Guam, Marshall Is, Micronesia (Federated States of), Northern Mariana Is, Palau, Puerto Rico, Timor-Leste, Turks and Caicos Is, US Minor Outlying Is, Virgin Is (British), Virgin Is (US)."},
		PesoUruguayo:                {"Peso Uruguayo", "Uruguay.", "Uruguayisk peso\n", "Uruguay"},
		UzbekistanSum:               {"Uzbekistan Sum", "Uzbekistan.", "Usbekisk sum", "Usbekistan"},
		Bolivar:                     {"Bolivar", "Deprecated, replaced by VEF.", "Bolivar", "UTGÅTT. Erstattet av VEF"},
		BolivarFuerte:               {"Bolivar fuerte", "Venezuela.", "Bolivar fuerte", "Venezuela"},
		Dong:                        {"Dong", "Viet Nam.", "Vietamesisk dong", "Vietnam"},
		Vatu:                        {"Vatu", "Vanuatu.", "Vatu", "Vanuatu"},
		Tala:                        {"Tala", "Samoa.", "Tala", "Samoa"},
		CFAFrancBEAC:                {"CFA Franc BEAC", "Cameroon, Central African Republic, Chad, Congo, Equatorial Guinea, Gabon.", "CFA franc BEAC", "Kamerun, Sentralafrikanske republikk, Tsjad, Kongo, Equatorial Guinea, Gabon"},
		EastCaribbeanDollar:         {"East Caribbean Dollar", "Anguilla, Antigua and Barbuda, Dominica, Grenada, Montserrat, Saint Kitts and Nevis, Saint Lucia, Saint Vincent and the Grenadines.", "Østkaribisk dollar", "Anguilla, Antigua and Barbuda, Dominica, Grenada, Montserrat, Saint Kitts and Nevis, Saint Lucia, Saint Vincent and the Grenadines."},
		CFAFrancBCEAO:               {"CFA Franc BCEAO", "Benin, Burkina Faso, Côte D’Ivoire, Guinea-Bissau, Mali, Niger, Senegal, Togo.", "CFA franc BCEAO", "Benin, Burkina Faso, Elfenbenskysten, Guinea-Bissau, Mali, Niger, Senegal, Togo."},
		CFPFranc:                    {"CFP Franc", "French Polynesia, New Caledonia, Wallis and Futuna.", "CFP franc", "Fransk Polynesia, Ny Kaledonia, Wallis and Futuna "},
		YemeniRial:                  {"Yemeni Rial", "Yemen.", "Jemenittisk rial", "Jemen"},
		YugoslavianDinar:            {"Yugoslavian Dinar", "DEPRECATED, replaced by CSD.", "Jugoslavisk dinar", "UTGÅTT, erstattet av CSD"},
		Rand:                        {"Rand", "South Africa.", "Rand", "Sør-Afrika"},
		Kwacha:                      {"Kwacha", "Deprecated, replaced with ZMW.", "Kwacha", "UTGÅTT. Erstattet av ZMW"},
		ZambianKwacha:               {"Zambian Kwacha", "Zambia.", "Zambisk kwacha", "Zambia"},
		ZimbabweDollar:              {"Zimbabwe Dollar", "Deprecated, replaced with ZWL.", "Zimbabwe dollar", "UTGÅTT. Erstattet av ZWL"},
		ZimbabweDollar:              {"Zimbabwe Dollar", "Zimbabwe.", "Zimbabwe dollar", "Zimbabwe"},
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
