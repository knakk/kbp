package list27 // Subject scheme identifier code

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Subject scheme identifier code
const (
	Dewey                                                 = "01"
	AbridgedDewey                                         = "02"
	LCClassification                                      = "03"
	LCSubjectHeading                                      = "04"
	NLMClassification                                     = "05"
	MeSHHeading                                           = "06"
	NALSubjectHeading                                     = "07"
	AAT                                                   = "08"
	UDC                                                   = "09"
	BISACSubjectHeading                                   = "10"
	BISACRegionCode                                       = "11"
	BICSubjectCategory                                    = "12"
	BICGeographicalQualifier                              = "13"
	BICLanguageQualifierLanguageAsSubject                 = "14"
	BICTimePeriodQualifier                                = "15"
	BICEducationalPurposeQualifier                        = "16"
	BICReadingLevelAndSpecialInterestQualifier            = "17"
	DDCSachgruppenDerDeutschenNationalbibliografie        = "18"
	LCFictionGenreHeading                                 = "19"
	Keywords                                              = "20"
	BICChildrensBookMarketingCategory                     = "21"
	BISACMerchandisingTheme                               = "22"
	PublishersOwnCategoryCode                             = "23"
	ProprietarySubjectScheme                              = "24"
	TablaDeMateriasISBN                                   = "25"
	WarengruppenSystematikDesDeutschenBuchhandels         = "26"
	SWD                                                   = "27"
	ThèmesElectre                                         = "28"
	CLIL                                                  = "29"
	DNBSachgruppen                                        = "30"
	NUGI                                                  = "31"
	NUR                                                   = "32"
	ECPAChristianBookCategory                             = "33"
	SISO                                                  = "34"
	KoreanDecimalClassificationKDC                        = "35"
	DDCDeutsch22                                          = "36"
	Bokgrupper                                            = "37"
	Varegrupper                                           = "38"
	Læreplaner                                            = "39"
	NipponDecimalClassification                           = "40"
	BSQ                                                   = "41"
	ANELEMaterias                                         = "42"
	Utdanningsprogram                                     = "43"
	Programfag                                            = "44"
	Undervisningsmateriell                                = "45"
	NorskDDK                                              = "46"
	Varugrupper                                           = "47"
	SAB                                                   = "48"
	Läromedel                                             = "49"
	Förhandsbeskrivning                                   = "50"
	SpanishISBNUDCSubset                                  = "51"
	ECISubjectCategories                                  = "52"
	SoggettoCCE                                           = "53"
	QualificatoreGeograficoCCE                            = "54"
	QualificatoreDiLinguaCCE                              = "55"
	QualificatoreDiPeriodoStoricoCCE                      = "56"
	QualificatoreDiLivelloScolasticoCCE                   = "57"
	QualificatoreDiEtàDiLetturaCCE                        = "58"
	VdSBildungsmedienFächer                               = "59"
	Fagkoder                                              = "60"
	JELClassification                                     = "61"
	CSH                                                   = "62"
	RVM                                                   = "63"
	YSA                                                   = "64"
	Allärs                                                = "65"
	YKL                                                   = "66"
	MUSA                                                  = "67"
	CILLA                                                 = "68"
	Kaunokki                                              = "69"
	Bella                                                 = "70"
	YSO                                                   = "71"
	PaikkatietoOntologia                                  = "72"
	SuomalainenKirjaAlanLuokitus                          = "73"
	Sears                                                 = "74"
	BICE4L                                                = "75"
	CSR                                                   = "76"
	SuomalainenOppiaineluokitus                           = "77"
	JapaneseBookTradeCCode                                = "78"
	JapaneseBookTradeGenreCode                            = "79"
	FiktiivisenAineistonLisäluokitus                      = "80"
	ArabicSubjectHeadingScheme                            = "81"
	ArabizedBICSubjectCategory                            = "82"
	ArabizedLCSubjectHeadings                             = "83"
	BibliothecaAlexandrinaSubjectHeadings                 = "84"
	PostalCode                                            = "85"
	GeoNamesID                                            = "86"
	NewBooksSubjectClassification                         = "87"
	ChineseLibraryClassification                          = "88"
	NTCPDSACClassification                                = "89"
	SeasonAndEventIndicator                               = "90"
	GND                                                   = "91"
	BICUKSLC                                              = "92"
	ThemaSubjectCategory                                  = "93"
	ThemaGeographicalQualifier                            = "94"
	ThemaLanguageQualifier                                = "95"
	ThemaTimePeriodQualifier                              = "96"
	ThemaEducationalPurposeQualifier                      = "97"
	ThemaInterestAgeSpecialInterestQualifier              = "98"
	ThemaStyleQualifier                                   = "99"
	Ämnesord                                              = "A2"
	StatystykaKsiążekPapierowychMówionychIElektronicznych = "A3"
	ThèmesBTLF                                            = "A4"
	Rameau                                                = "A5"
	NomenclatureDisciplineScolaire                        = "A6"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "A2", "A3", "A4", "A5", "A6"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Dewey:                                          {"Dewey", "Dewey Decimal Classification.", "Dewey", "Refererer til eksterne kodelister (Amerikansk DDC)"},
		AbridgedDewey:                                  {"Abridged Dewey", "", "Forkortet Dewey", "Refererer til eksterne kodelister (Abridged Dewey)"},
		LCClassification:                               {"LC classification", "US Library of Congress classification.", "LC class number", "Library of Congress klassifikasjon. Refererer til eksterne kodelister"},
		LCSubjectHeading:                               {"LC subject heading", "US Library of Congress subject heading.", "LC subject heading", "Library of Congress emneord. Refererer til eksterne kodelister"},
		NLMClassification:                              {"NLM classification", "US National Library of Medicine medical classification.", "NLM classification", "US National Library of Medicine medical classification."},
		MeSHHeading:                                    {"MeSH heading", "US National Library of Medicine Medical subject heading.", "MeSH heading", "US National Library of Medicine Medical subject heading."},
		NALSubjectHeading:                              {"NAL subject heading", "US National Agricultural Library subject heading.", "NAL subject heading", "US National Agricultural Library subject heading."},
		AAT:                                            {"AAT", "Getty Art and Architecture Thesaurus heading.", "AAT", "Getty Art and Architecture Thesaurus heading."},
		UDC:                                            {"UDC", "Universal Decimal Classification.", "UDK", "Universell desimalklassifikasjon. Refererer til eksterne kodelister"},
		BISACSubjectHeading:                            {"BISAC Subject Heading", "BISAC Subject Headings are used in the North American market to categorize books based on topical content. They serve as a guideline for shelving books in physical stores and browsing books in online stores. See ‘http://www.bisg.org/what-we-do-cat-20-classification-schemes.php’.", "BISAC category code", "For informasjon om BISAC emnekategorier se http://www.bisg.org.  Refererer til eksterne kodelister"},
		BISACRegionCode:                                {"BISAC region code", "A geographical qualifier used with a BISAC subject category.", "BISAC regionkoder", "En geografisk kvalifikator brukt sammen med BISAC-emnekategorier. Refererer til eksterne kodelister.\n"},
		BICSubjectCategory:                             {"BIC subject category", "For all BIC subject codes and qualifiers, see ‘http://www.bic.org.uk/7/BIC-Standard-Subject-Categories/’.", "BIC subsidiary subject", "For informasjon om BISAC emnekategorier se: http://www.bisg.org.  Refererer til eksterne kodelister"},
		BICGeographicalQualifier:                       {"BIC geographical qualifier", "", "BIC geographical qualifier", "Refererer til eksterne kodelister"},
		BICLanguageQualifierLanguageAsSubject:          {"BIC language qualifier (language as subject)", "", "BIC language qualifier (language as subject)", "Refererer til eksterne kodelister"},
		BICTimePeriodQualifier:                         {"BIC time period qualifier", "", "BIC time period qualifier", "Refererer til eksterne kodelister"},
		BICEducationalPurposeQualifier:                 {"BIC educational purpose qualifier", "", "BIC educational purpose qualifier", "Refererer til eksterne kodelister"},
		BICReadingLevelAndSpecialInterestQualifier:     {"BIC reading level and special interest qualifier", "", "BIC lesenivå og interessegrupper", "Refererer til eksterne kodelister"},
		DDCSachgruppenDerDeutschenNationalbibliografie: {"DDC-Sachgruppen der Deutschen Nationalbibliografie", "Used for German National Bibliography since 2004 (100 subjects). Is different from value 30. See http://www.d-nb.de/service/pdf/ddc_wv_aktuell.pdf (in German) or http://www.d-nb.de/eng/service/pdf/ddc_wv_aktuell_eng.pdf (English).", "DDC-Sachgruppen der Deutschen Nationalbibliografie", "Used for German National Bibliography since 2004 (100 subjects). Is different from value 30. See http://www.d-nb.de/service/pdf/ddc_wv_aktuell.pdf (in German) or http://www.d-nb.de/eng/service/pdf/ddc_wv_aktuell_eng.pdf (English)."},
		LCFictionGenreHeading:                          {"LC fiction genre heading", "", "LC fiction genre heading", "Library of Congress sjangerklassifisering. Refererer til eksterne kodelister"},
		Keywords:                                       {"Keywords", "Where multiple keywords or keyword phrases are sent in a single instance of the <SubjectHeadingText> element, it is recommended that they should be separated by semi-colons (this is consistent with Library of Congress preferred practice).", "Nøkkelord", "Frie emneord"},
		BICChildrensBookMarketingCategory:             {"BIC children’s book marketing category", "See ‘http://www.bic.org.uk/8/Children’s-Books-Marketing-Classifications/’.", "BIC children's book marketing category", "Se http://www.bic.org.uk/cbmc.html.  Refererer til eksterne kodelister"},
		BISACMerchandisingTheme:                       {"BISAC Merchandising Theme", "BISAC Merchandising Themes are used in addition to BISAC Subject Headings to denote an audience to which a work may be of particular appeal, a time of year or event for which a work may be especially appropriate, or to further describe fictional works that have been subject-coded by genre.", "BISAC book merchandising code", "For informasjon om BISAC emnekategorier se http://www.bisg.org.  Refererer til eksterne kodelister"},
		PublishersOwnCategoryCode:                     {"Publisher’s own category code", "", "Forleggerens egen kategorikode", "Refererer til eksterne kodelister"},
		ProprietarySubjectScheme:                      {"Proprietary subject scheme", "As specified in <SubjectSchemeName>.", "Proprietært emneskjema", "Refererer til eksterne kodelister"},
		TablaDeMateriasISBN:                           {"Tabla de materias ISBN", "Latin America.", "Tabla de materias ISBN", "Latin-Amerika. Refererer til eksterne kodelister."},
		WarengruppenSystematikDesDeutschenBuchhandels: {"Warengruppen-Systematik des deutschen Buchhandels", "See http://www.boersenverein.de/sixcms/media.php/976/WGSneuVersion2_0.pdf (in German).", "Warengruppen-Systematik des deutschen Buchhandels", "Refererer til eksterne kodelister"},
		SWD:            {"SWD", "Schlagwortnormdatei – Subject Headings Authority File in the German-speaking countries. See http://www.d-nb.de/standardisierung/normdateien/swd.htm (in German) and http://www.d-nb.de/eng/standardisierung/normdateien/swd.htm (English). DEPRECATED in favour of the GND.", "Schlagwort-Normdatei der Deutschen Bibliothek", "Tekst"},
		ThèmesElectre:  {"Thèmes Electre", "Subject classification used by Electre (France).", "Thèmes Electre", "Emneklassifisering brukt av Electre (France). Refererer til eksterne kodelister"},
		CLIL:           {"CLIL", "France. A four-digit number, see http://www.clil.org/information/documentation.html (in French). The first digit identifies the version of the scheme.", "CLIL", "Fransk. Refererer til eksterne kodelister"},
		DNBSachgruppen: {"DNB-Sachgruppen", "Deutsche Bibliothek subject groups. Used for German National Bibliography until 2003 (65 subjects). Is different from value 18. See http://www.d-nb.de/service/pdf/ddc_wv_alt_neu.pdf (in German).", "DNB-Sachgruppen", "Deutsche Bibliothek emnegrupper. Refererer til eksterne kodelister"},
		NUGI:           {"NUGI", "Nederlandse Uniforme Genre-Indeling (former Dutch book trade classification).", "NUGI", "Nederlandse Uniforme Genre-Indeling. Refererer til eksterne kodelister"},
		NUR:            {"NUR", "Nederlandstalige Uniforme Rubrieksindeling (Dutch book trade classification, from 2002, see http://reeks.boekwinkeltjes.nl/downloads/NUR-lijst.pdf (in Dutch).", "NUR", "Nederlandstalige Uniforme Rubrieksindeling. Refererer til eksterne kodelister"},
		ECPAChristianBookCategory: {"ECPA Christian Book Category", "ECPA Christian Product Category Book Codes, consisting of up to three x 3-letter blocks, for Super Category, Primary Category and Sub-Category. See ‘http://www.ecpa.org/ECPA/cbacategories.xls’.", "ECPA Christian Book Category", "ECPA Christian Product Category Book Codes, See http://www.ecpa.org/ECPA/cbacategories.xls. Refererer til eksterne kodelister"},
		SISO: {"SISO", "Schema Indeling Systematische Catalogus Openbare Bibliotheken (Dutch library classification).", "SISO", "Schema Indeling Systematische Catalogus Openbare Bibliotheken. Refererer til eksterne kodelister"},
		KoreanDecimalClassificationKDC: {"Korean Decimal Classification (KDC)", "A modified Dewey Decimal Classification used in the Republic of Korea.", "Korean Decimal Classification (KDC)", "En modifisert Dewey Decimal Classification brukt i sør-Korea.  Refererer til eksterne kodelister"},
		DDCDeutsch22:                   {"DDC Deutsch 22", "German Translation of Dewey Decimal Classification 22. Also known as DDC 22 ger. See ‘http://www.ddc-deutsch.de/produkte/uebersichten/’.", "DDC Deutsch", "Refererer til eksterne kodelister"},
		Bokgrupper:                     {"Bokgrupper", "Norwegian book trade product categories (Bokgrupper) administered by the Norwegian Publishers Association (http://www.forleggerforeningen.no/).", "Bokgrupper", "Norske bokgrupper. Administrert av Forleggerforeningen  (http://www.forleggerforeningen.no/)."},
		Varegrupper:                    {"Varegrupper", "Norwegian bookselling subject categories (Bokhandelens varegrupper) administered by the Norwegian Booksellers Association (http://bokhandlerforeningen.no/).", "Varegrupper", "Norske varegrupper. Administrert av Bokhandlerforeningen (http://bokhandlerforeningen.no)"},
		Læreplaner:                     {"Læreplaner", "Norwegian school curriculum version. Deprecated", "Læreplaner", "Norske læreplaner. UTGÅTT"},
		NipponDecimalClassification:    {"Nippon Decimal Classification", "Japanese subject classification scheme.", "Japansk desimalklassifikasjon", "Japansk emneklassifiseringssystem. Refererer til eksterne kodelister"},
		BSQ:                                 {"BSQ", "BookSelling Qualifier: Russian book trade classification.", "Russisk BSQ", "Russisk. Refererer til eksterne kodelister"},
		ANELEMaterias:                       {"ANELE Materias", "Spain: subject coding scheme of the Asociación Nacional de Editores de Libros y Material de Enseñanza.", "Spansk ANELE Materias", "Spania: emnekodingsskjema fra Asociación Nacional de Editores de Libros y Material de Enseñanza.  Refererer til eksterne kodelister"},
		Utdanningsprogram:                   {"Utdanningsprogram", "Codes for Norwegian ‘utdanningsprogram’ used in secondary education. See: http://www.udir.no/. (Formerly ‘Skolefag’.)", "Utdanningsprogram", "Koder for utdanningsprogram. Se: http://www.udir.no/"},
		Programfag:                          {"Programfag", "Codes for Norwegian ‘programfag’ used in secondary education. See http://www.udir.no/. (Formerly ‘Videregående’.)", "Programfag", "Koder for programfag. Se: http://www.udir.no/"},
		Undervisningsmateriell:              {"Undervisningsmateriell", "Norwegian list of categories for books and other material used in education (4707).", "Undervisningsmateriell", "Norsk liste over type undervisningsmateriell (4707)"},
		NorskDDK:                            {"Norsk DDK", "Norwegian version of Dewey Decimal Classification.", "DDK 5", "Norsk utgave av DDC 21"},
		Varugrupper:                         {"Varugrupper", "Swedish bookselling subject categories.", "Varegrupper i Sverige", "Svensk emneklassifisering for bokhandel. Ekstern kodeliste"},
		SAB:                                 {"SAB", "Swedish classification scheme.", "SAB", "Svensk klassifikasjonsskjema. Tekst"},
		Läromedel:                           {"Läromedel", "Swedish bookselling educational subject.", "Læremiddel i Sverige", "Svenske skoleemner. Koder"},
		Förhandsbeskrivning:                 {"Förhandsbeskrivning", "Swedish publishers preliminary subject classification.", "Forhåndsbeskrivelse i Sverige", "Svenske utgiveres forhåndsklassifikasjon. Koder"},
		SpanishISBNUDCSubset:                {"Spanish ISBN UDC subset", "Controlled subset of UDC codes used by the Spanish ISBN Agency.", "Spansk delmengde av ISBN UDC", "Kontrollert delmengde av UDC-koder, brukt av det spanske ISBN-kontoret. Koder"},
		ECISubjectCategories:                {"ECI subject categories", "Subject categories defined by El Corte Inglés and used widely in the Spanish book trade.", "ECI emnekategorier", "Emnekategorier definert av El Corte Inglés, mye brukt i spansk bokhandel.Koder"},
		SoggettoCCE:                         {"Soggetto CCE", "Classificazione commerciale editoriale (Italian book trade subject category based on BIC). CCE documentation available at ‘http://www.ie-online.it/CCE2_2.0.pdf’.", "Soggetto CCE", "Classificazione commerciale editoriale (Italienske emnekategorier, basert på BIC). Koder"},
		QualificatoreGeograficoCCE:          {"Qualificatore geografico CCE", "CCE Geographical qualifier.", "Qualificatore geografico CCE", "Refererer til eksterne kodelister"},
		QualificatoreDiLinguaCCE:            {"Qualificatore di lingua CCE", "CCE Language qualifier.", "Qualificatore di lingua CCE", "Refererer til eksterne kodelister"},
		QualificatoreDiPeriodoStoricoCCE:    {"Qualificatore di periodo storico CCE", "CCE Time Period qualifier.", "Qualificatore di periodo storico CCE", "Refererer til eksterne kodelister"},
		QualificatoreDiLivelloScolasticoCCE: {"Qualificatore di livello scolastico CCE", "CCE Educational Purpose qualifier.", "Qualificatore di livello scolastico CCE", "Refererer til eksterne kodelister"},
		QualificatoreDiEtàDiLetturaCCE:      {"Qualificatore di età di lettura CCE", "CCE Reading Level Qualifier.", "Qualificatore di età di lettura CCE", "Refererer til eksterne kodelister"},
		VdSBildungsmedienFächer:             {"VdS Bildungsmedien Fächer", "Subject code list of the German association of educational media publishers.", "VdS Bildungsmedien Fächer", "Emnekoder for tyske læremiddelutgivere"},
		Fagkoder:                            {"Fagkoder", "Norwegian primary and secondary school subject categories (fagkoder), see \"http://www.udir.no/\".", "Fagkoder", "Fagkoder for fag i grunn- og videregående skole. Se: http://www.udir.no/"},
		JELClassification:                   {"JEL classification", "Journal of Economic Literature classification scheme.", "JEL klassifikasjon", "Journal of Economic Literature klassifikasjonsskjema"},
		CSH:                                 {"CSH", "National Library of Canada subject heading (English).", "CSH", "National Library of Canada subject heading (English)."},
		RVM:                                 {"RVM", "Répertoire de vedettes-matière (Bibliothèque et Archives Canada et Bibliothèque de l’Université Laval) (French).", "RVM", "Répertoire de vedettes-matière (Bibliothèque et Archives Canada et Bibliothèque de l’Université Laval) (French)."},
		YSA:                                 {"YSA", "Yleinen suomalainen asiasanasto: Finnish General Thesaurus. See http://onki.fi/fi/browser/ (in Finnish).", "YSA", "Yleinen suomalainen asiasanasto: Finnish General Thesaurus. See http://onki.fi/fi/browser/ (in Finnish)."},
		Allärs:                              {"Allärs", "Allmän tesaurus på svenska: Swedish translation of the Finnish General Thesaurus. See http://onki.fi/fi/browser/ (in Finnish).", "Allärs", "Allmän tesaurus på svenska: Swedish translation of the Finnish General Thesaurus. See http://onki.fi/fi/browser/ (in Finnish)."},
		YKL:                                 {"YKL", "Yleisten kirjastojen luokitusjärjestelmä: Finnish Public Libraries Classification System. See http://ykl.kirjastot.fi/ (in Finnish).", "YKL", "Yleisten kirjastojen luokitusjärjestelmä: Finnish Public Libraries Classification System. See http://ykl.kirjastot.fi/ (in Finnish)."},
		MUSA:                                {"MUSA", "Musiikin asiasanasto: Finnish Music Thesaurus. See http://onki.fi/fi/browser/ (in Finnish).", "MUSA", "Musiikin asiasanasto: Finnish Music Thesaurus. See http://onki.fi/fi/browser/ (in Finnish)."},
		CILLA:                               {"CILLA", "Specialtesaurus för musik: Swedish translation of the Finnish Music Thesaurus. See http://onki.fi/fi/browser/ (in Finnish).", "CILLA", "Specialtesaurus för musik: Swedish translation of the Finnish Music Thesaurus. See http://onki.fi/fi/browser/ (in Finnish)."},
		Kaunokki:                            {"Kaunokki", "Fiktiivisen aineiston asiasanasto: Finnish thesaurus for fiction. See http://kaunokki.kirjastot.fi/ (in Finnish).", "Kaunokki", "Fiktiivisen aineiston asiasanasto: Finnish thesaurus for fiction. See http://kaunokki.kirjastot.fi/ (in Finnish)."},
		Bella:                               {"Bella", "Specialtesaurus för fiktivt material: Swedish translation of the Finnish thesaurus for fiction. See http://kaunokki.kirjastot.fi/sv-FI/ (in Finnish).", "Bella", "Specialtesaurus för fiktivt material: Swedish translation of the Finnish thesaurus for fiction. See http://kaunokki.kirjastot.fi/sv-FI/ (in Finnish)."},
		YSO:                                 {"YSO", "Yleinen suomalainen ontologia: Finnish General Upper Ontology. See http://onki.fi/fi/browser/ (In Finnish).", "YSO", "Yleinen suomalainen ontologia: Finnish General Upper Ontology. See http://onki.fi/fi/browser/ (In Finnish)."},
		PaikkatietoOntologia:                {"Paikkatieto ontologia", "Finnish Place Ontology. See http://onki.fi/fi/browser/ (in Finnish).", "Paikkatieto ontologia", "Finnish Place Ontology. See http://onki.fi/fi/browser/ (in Finnish)."},
		SuomalainenKirjaAlanLuokitus:        {"Suomalainen kirja-alan luokitus", "Finnish book trade categorisation.", "Suomalainen kirja-alan luokitus", "Finnish book trade categorisation."},
		Sears:  {"Sears", "Sears List of Subject Headings.", "Sears", "Sears List of Subject Headings."},
		BICE4L: {"BIC E4L", "BIC E4Libraries Category Headings, see ‘http://www.bic.org.uk/51/E4libraries-Subject-Category-Headings/’.", "BIC E4L", "BIC E4Libraries Category Headings, see ‘http://www.bic.org.uk/51/E4libraries-Subject-Category-Headings/’."},
		CSR:    {"CSR", "Code Sujet Rayon: subject categories used by bookstores in France.", "CSR", "Code Sujet Rayon: subject categories used by bookstores in France."},
		SuomalainenOppiaineluokitus:           {"Suomalainen oppiaineluokitus", "Finnish school subject categories.", "Suomalainen oppiaineluokitus", "Finnish school subject categories."},
		JapaneseBookTradeCCode:                {"Japanese book trade C-Code", "See ‘http://www.asahi-net.or.jp/~ax2s-kmtn/ref/ccode.html’ (in Japanese).", "Japanese book trade C-Code", "See ‘http://www.asahi-net.or.jp/~ax2s-kmtn/ref/ccode.html’ (in Japanese)."},
		JapaneseBookTradeGenreCode:            {"Japanese book trade Genre Code", "", "Japanese book trade Genre Code", ""},
		FiktiivisenAineistonLisäluokitus:      {"Fiktiivisen aineiston lisäluokitus", "Finnish fiction genre classification. See http://ykl.kirjastot.fi/fi-FI/lisaluokat/ (in Finnish).", "Fiktiivisen aineiston lisäluokitus", "Finnish fiction genre classification. See http://ykl.kirjastot.fi/fi-FI/lisaluokat/ (in Finnish)."},
		ArabicSubjectHeadingScheme:            {"Arabic Subject heading scheme", "", "Arabic Subject heading scheme", ""},
		ArabizedBICSubjectCategory:            {"Arabized BIC subject category", "Arabized version of BIC subject category scheme developed by ElKotob.com.", "Arabized BIC subject category", "Arabized version of BIC subject category scheme developed by ElKotob.com."},
		ArabizedLCSubjectHeadings:             {"Arabized LC subject headings", "Arabized version of Library of Congress scheme.", "Arabized LC subject headings", "Arabized version of Library of Congress scheme."},
		BibliothecaAlexandrinaSubjectHeadings: {"Bibliotheca Alexandrina Subject Headings", "Classification scheme used by Library of Alexandria.", "Bibliotheca Alexandrina Subject Headings", "Classification scheme used by Library of Alexandria."},
		PostalCode:                            {"Postal code", "Location defined by postal code. Format is two-letter country code (from List 91), space, postal code. Note some postal codes themselves contain spaces, eg “GB N7 9DP” or “US 10125”.", "Postal code", "Location defined by postal code. Format is two-letter country code (from List 91), space, postal code. Note some postal codes themselves contain spaces, eg “GB N7 9DP” or “US 10125”."},
		GeoNamesID:                            {"GeoNames ID", "ID number for geographical place, as defined at http://www.geonames.org (eg 2825297 is Stuttgart, Germany, see http://www.geonames.org/2825297).", "GeoNames ID", "ID number for geographical place, as defined at http://www.geonames.org (eg 2825297 is Stuttgart, Germany, see http://www.geonames.org/2825297)."},
		NewBooksSubjectClassification:         {"NewBooks Subject Classification", "Used for classification of academic and specialist publication in German-speaking countries. See http://www.newbooks-services.com/de/top/unternehmensportrait/klassifikation-und-mapping.html (German) and http://www.newbooks-services.com/en/top/about-newbooks/classification-mapping.html (English).", "NewBooks Subject Classification", "Used for classification of academic and specialist publication in German-speaking countries. See http://www.newbooks-services.com/de/top/unternehmensportrait/klassifikation-und-mapping.html (German) and http://www.newbooks-services.com/en/top/about-newbooks/classification-mapping.html (English)."},
		ChineseLibraryClassification:          {"Chinese Library Classification", "Subject classification maintained by the Editorial Board of Chinese Library Classification. See http://cct.nlc.gov.cn for access to details of the scheme.", "Chinese Library Classification", "Subject classification maintained by the Editorial Board of Chinese Library Classification. See http://cct.nlc.gov.cn for access to details of the scheme."},
		NTCPDSACClassification:                {"NTCPDSAC Classification", "Subject classification for Books, Audiovisual products and E-publications formulated by China National Technical Committee 505.", "NTCPDSAC Classification", "Subject classification for Books, Audiovisual products and E-publications formulated by China National Technical Committee 505."},
		SeasonAndEventIndicator:               {"Season and Event Indicator", "German code scheme indicating association with seasons, holidays, events (eg Autumn, Back to School, Easter).", "Season and Event Indicator", "German code scheme indicating association with seasons, holidays, events (eg Autumn, Back to School, Easter)."},
		GND:                                      {"GND", "Gemeinsame Normdatei – Joint Authority File in the German-speaking countries. See ‘http://www.dnb.de/EN/Standardisierung/Normdaten/GND/gnd_node.html (English)’. Combines the PND, SWD and GKD into a single authority file, and should be used in preference to the older codes.", "GND", "Gemeinsame Normdatei – Joint Authority File in the German-speaking countries. See ‘http://www.dnb.de/EN/Standardisierung/Normdaten/GND/gnd_node.html (English)’. Combines the PND, SWD and GKD into a single authority file, and should be used in preference to the older codes."},
		BICUKSLC:                                 {"BIC UKSLC", "UK Standard Library Categories, the successor to BIC’s E4L classification scheme.", "BIC UKSLC", "UK Standard Library Categories, the successor to BIC’s E4L classification scheme."},
		ThemaSubjectCategory:                     {"Thema subject category", "", "Thema subject category", ""},
		ThemaGeographicalQualifier:               {"Thema geographical qualifier", "", "Thema geographical qualifier", ""},
		ThemaLanguageQualifier:                   {"Thema language qualifier", "", "Thema language qualifier", ""},
		ThemaTimePeriodQualifier:                 {"Thema time period qualifier", "", "Thema time period qualifier", ""},
		ThemaEducationalPurposeQualifier:         {"Thema educational purpose qualifier", "", "Thema educational purpose qualifier", ""},
		ThemaInterestAgeSpecialInterestQualifier: {"Thema interest age / special interest qualifier", "", "Thema interest age / special interest qualifier", ""},
		ThemaStyleQualifier:                      {"Thema style qualifier", "", "Thema style qualifier", ""},
		Ämnesord:                                 {"Ämnesord", "Swedish subject categories maintained by Bokrondellen.", "Ämnesord", "Swedish subject categories maintained by Bokrondellen."},
		StatystykaKsiążekPapierowychMówionychIElektronicznych: {"Statystyka Książek Papierowych, Mówionych I Elektronicznych", "Polish Statistical Book and E-book Classification.", "Statystyka Książek Papierowych, Mówionych I Elektronicznych", "Polish Statistical Book and E-book Classification."},
		ThèmesBTLF: {"Thèmes BTLF", "Subject classification used by BTLF (Québec).", "Thèmes BTLF", "Subject classification used by BTLF (Québec)."},
		Rameau:     {"Rameau", "French library classification.", "Rameau", "French library classification."},
		NomenclatureDisciplineScolaire: {"Nomenclature discipline scolaire", "French educational subject classification, used for example on WizWiz.fr. See ‘http://www.kiosque-edu.com/html/onix/Nomenclature_disciplines.csv’.", "Nomenclature discipline scolaire", "French educational subject classification, used for example on WizWiz.fr. See ‘http://www.kiosque-edu.com/html/onix/Nomenclature_disciplines.csv’."},
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
