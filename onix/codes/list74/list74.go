package list74 // Language code – ISO 639-2/B

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

// Possible values for Language code – ISO 639-2/B
const (
	Afar                                                               = "aar"
	Abkhaz                                                             = "abk"
	Achinese                                                           = "ace"
	Acoli                                                              = "ach"
	Adangme                                                            = "ada"
	Adygei                                                             = "ady"
	AfroAsiaticLanguages                                               = "afa"
	Afrihili                                                           = "afh"
	Afrikaans                                                          = "afr"
	Ainu                                                               = "ain"
	Akan                                                               = "aka"
	Akkadian                                                           = "akk"
	Albanian                                                           = "alb"
	Aleut                                                              = "ale"
	AlgonquianLanguages                                                = "alg"
	SouthernAltai                                                      = "alt"
	Amharic                                                            = "amh"
	EnglishOldCa4501100                                                = "ang"
	Angika                                                             = "anp"
	ApacheLanguages                                                    = "apa"
	Arabic                                                             = "ara"
	OfficialAramaicImperialAramaic700300BCE                            = "arc"
	Aragonese                                                          = "arg"
	Armenian                                                           = "arm"
	MapudungunMapuche                                                  = "arn"
	Arapaho                                                            = "arp"
	ArtificialLanguages                                                = "art"
	Arawak                                                             = "arw"
	Assamese                                                           = "asm"
	AsturianBableLeoneseAsturleonese                                   = "ast"
	AthapascanLanguages                                                = "ath"
	AustralianLanguages                                                = "aus"
	Avaric                                                             = "ava"
	Avestan                                                            = "ave"
	Awadhi                                                             = "awa"
	Aymara                                                             = "aym"
	Azerbaijani                                                        = "aze"
	BandaLanguages                                                     = "bad"
	BamilekeLanguages                                                  = "bai"
	Bashkir                                                            = "bak"
	Baluchi                                                            = "bal"
	Bambara                                                            = "bam"
	Balinese                                                           = "ban"
	Basque                                                             = "baq"
	Basa                                                               = "bas"
	BalticLanguages                                                    = "bat"
	BejaBedawiyet                                                      = "bej"
	Belarusian                                                         = "bel"
	Bemba                                                              = "bem"
	Bengali                                                            = "ben"
	BerberLanguages                                                    = "ber"
	Bhojpuri                                                           = "bho"
	BihariLanguages                                                    = "bih"
	Bikol                                                              = "bik"
	BiniEdo                                                            = "bin"
	Bislama                                                            = "bis"
	Siksika                                                            = "bla"
	BantuLanguages                                                     = "bnt"
	Bosnian                                                            = "bos"
	Braj                                                               = "bra"
	Breton                                                             = "bre"
	BatakLanguages                                                     = "btk"
	Buriat                                                             = "bua"
	Buginese                                                           = "bug"
	Bulgarian                                                          = "bul"
	Burmese                                                            = "bur"
	BlinBilin                                                          = "byn"
	Caddo                                                              = "cad"
	CentralAmericanIndianLanguages                                     = "cai"
	GalibiCarib                                                        = "car"
	Catalan                                                            = "cat"
	CaucasianLanguages                                                 = "cau"
	Cebuano                                                            = "ceb"
	CelticLanguages                                                    = "cel"
	Chamorro                                                           = "cha"
	Chibcha                                                            = "chb"
	Chechen                                                            = "che"
	Chagatai                                                           = "chg"
	Chinese                                                            = "chi"
	ChuukeseTruk                                                       = "chk"
	Mari                                                               = "chm"
	ChinookJargon                                                      = "chn"
	Choctaw                                                            = "cho"
	ChipewyanDeneSuline                                                = "chp"
	Cherokee                                                           = "chr"
	ChurchSlavicOldSlavonicChurchSlavonicOldBulgarianOldChurchSlavonic = "chu"
	Chuvash                                                            = "chv"
	Cheyenne                                                           = "chy"
	ChamicLanguages                                                    = "cmc"
	Coptic                                                             = "cop"
	Cornish                                                            = "cor"
	Corsican                                                           = "cos"
	CreolesAndPidginsEnglishBased                                      = "cpe"
	CreolesAndPidginsFrenchBased                                       = "cpf"
	CreolesAndPidginsPortugueseBased                                   = "cpp"
	Cree                                                               = "cre"
	CrimeanTurkishCrimeanTatar                                         = "crh"
	CreolesAndPidgins                                                  = "crp"
	Kashubian                                                          = "csb"
	CushiticLanguages                                                  = "cus"
	Czech                                                              = "cze"
	Dakota                                                             = "dak"
	Danish                                                             = "dan"
	Dargwa                                                             = "dar"
	LandDayakLanguages                                                 = "day"
	Delaware                                                           = "del"
	SlaveAthapascan                                                    = "den"
	Dogrib                                                             = "dgr"
	Dinka                                                              = "din"
	DivehiDhivehiMaldivian                                             = "div"
	Dogri                                                              = "doi"
	DravidianLanguages                                                 = "dra"
	LowerSorbian                                                       = "dsb"
	Duala                                                              = "dua"
	DutchMiddleCa10501350                                              = "dum"
	DutchFlemish                                                       = "dut"
	Dyula                                                              = "dyu"
	Dzongkha                                                           = "dzo"
	Efik                                                               = "efi"
	EgyptianAncient                                                    = "egy"
	Ekajuk                                                             = "eka"
	Elamite                                                            = "elx"
	English                                                            = "eng"
	EnglishMiddle11001500                                              = "enm"
	Esperanto                                                          = "epo"
	Estonian                                                           = "est"
	Ewe                                                                = "ewe"
	Ewondo                                                             = "ewo"
	Fang                                                               = "fan"
	Faroese                                                            = "fao"
	Fanti                                                              = "fat"
	Fijian                                                             = "fij"
	FilipinoPilipino                                                   = "fil"
	Finnish                                                            = "fin"
	FinnoUgrianLanguages                                               = "fiu"
	Kvensk                                                             = "fkv"
	Fon                                                                = "fon"
	French                                                             = "fre"
	FrenchMiddleCa14001600                                             = "frm"
	FrenchOldCa8421400                                                 = "fro"
	NorthernFrisian                                                    = "frr"
	EasternFrisian                                                     = "frs"
	WesternFrisian                                                     = "fry"
	Fulah                                                              = "ful"
	Friulian                                                           = "fur"
	Gã                                                                 = "gaa"
	Gayo                                                               = "gay"
	Gbaya                                                              = "gba"
	GermanicLanguages                                                  = "gem"
	Georgian                                                           = "geo"
	German                                                             = "ger"
	EthiopicGeez                                                       = "gez"
	Gilbertese                                                         = "gil"
	ScottishGaelic                                                     = "gla"
	Irish                                                              = "gle"
	Galician                                                           = "glg"
	Manx                                                               = "glv"
	GermanMiddleHighCa10501500                                         = "gmh"
	GermanOldHighCa7501050                                             = "goh"
	Gondi                                                              = "gon"
	Gorontalo                                                          = "gor"
	Gothic                                                             = "got"
	Grebo                                                              = "grb"
	GreekAncientTo1453                                                 = "grc"
	GreekModern1453                                                    = "gre"
	Guarani                                                            = "grn"
	SwissGermanAlemannic                                               = "gsw"
	Gujarati                                                           = "guj"
	Gwichin                                                            = "gwi"
	Haida                                                              = "hai"
	HaitianFrenchCreole                                                = "hat"
	Hausa                                                              = "hau"
	Hawaiian                                                           = "haw"
	Hebrew                                                             = "heb"
	Herero                                                             = "her"
	Hiligaynon                                                         = "hil"
	HimachaliLanguagesWesternPahariLanguages                           = "him"
	Hindi                                                              = "hin"
	Hittite                                                            = "hit"
	HmongMong                                                          = "hmn"
	HiriMotu                                                           = "hmo"
	Croatian                                                           = "hrv"
	UpperSorbian                                                       = "hsb"
	Hungarian                                                          = "hun"
	Hupa                                                               = "hup"
	Iban                                                               = "iba"
	Igbo                                                               = "ibo"
	Icelandic                                                          = "ice"
	Ido                                                                = "ido"
	SichuanYiNuosu                                                     = "iii"
	IjoLanguages                                                       = "ijo"
	Inuktitut                                                          = "iku"
	InterlingueOccidental                                              = "ile"
	Iloko                                                              = "ilo"
	InterlinguaInternationalAuxiliaryLanguageAssociation               = "ina"
	IndicLanguages                                                     = "inc"
	Indonesian                                                         = "ind"
	IndoEuropeanLanguages                                              = "ine"
	Ingush                                                             = "inh"
	Inupiaq                                                            = "ipk"
	IranianLanguages                                                   = "ira"
	IroquoianLanguages                                                 = "iro"
	Italian                                                            = "ita"
	Javanese                                                           = "jav"
	Lojban                                                             = "jbo"
	Japanese                                                           = "jpn"
	JudeoPersian                                                       = "jpr"
	JudeoArabic                                                        = "jrb"
	KaraKalpak                                                         = "kaa"
	Kabyle                                                             = "kab"
	KachinJingpho                                                      = "kac"
	KalâtdlisutGreenlandic                                             = "kal"
	Kamba                                                              = "kam"
	Kannada                                                            = "kan"
	KarenLanguages                                                     = "kar"
	Kashmiri                                                           = "kas"
	Kanuri                                                             = "kau"
	Kawi                                                               = "kaw"
	Kazakh                                                             = "kaz"
	KabardianCircassian                                                = "kbd"
	Khasi                                                              = "kha"
	KhoisanLanguages                                                   = "khi"
	CentralKhmer                                                       = "khm"
	KhotaneseSakan                                                     = "kho"
	KikuyuGikuyu                                                       = "kik"
	Kinyarwanda                                                        = "kin"
	KirghizKyrgyz                                                      = "kir"
	Kimbundu                                                           = "kmb"
	Konkani                                                            = "kok"
	Komi                                                               = "kom"
	Kongo                                                              = "kon"
	Korean                                                             = "kor"
	KusaieanCarolineIslands                                            = "kos"
	Kpelle                                                             = "kpe"
	KarachayBalkar                                                     = "krc"
	Karelian                                                           = "krl"
	KruLanguages                                                       = "kro"
	Kurukh                                                             = "kru"
	Kuanyama                                                           = "kua"
	Kumyk                                                              = "kum"
	Kurdish                                                            = "kur"
	Kutenai                                                            = "kut"
	Ladino                                                             = "lad"
	Lahnda                                                             = "lah"
	Lamba                                                              = "lam"
	Lao                                                                = "lao"
	Latin                                                              = "lat"
	Latvian                                                            = "lav"
	Lezgian                                                            = "lez"
	Limburgish                                                         = "lim"
	Lingala                                                            = "lin"
	Lithuanian                                                         = "lit"
	MongoNkundu                                                        = "lol"
	Lozi                                                               = "loz"
	LuxembourgishLetzeburgesch                                         = "ltz"
	LubaLulua                                                          = "lua"
	LubaKatanga                                                        = "lub"
	Ganda                                                              = "lug"
	Luiseño                                                            = "lui"
	Lunda                                                              = "lun"
	LuoKenyaAndTanzania                                                = "luo"
	Lushai                                                             = "lus"
	Macedonian                                                         = "mac"
	Madurese                                                           = "mad"
	Magahi                                                             = "mag"
	Marshallese                                                        = "mah"
	Maithili                                                           = "mai"
	Makasar                                                            = "mak"
	Malayalam                                                          = "mal"
	Mandingo                                                           = "man"
	Maori                                                              = "mao"
	AustronesianLanguages                                              = "map"
	Marathi                                                            = "mar"
	Masai                                                              = "mas"
	Malay                                                              = "may"
	Moksha                                                             = "mdf"
	Mandar                                                             = "mdr"
	Mende                                                              = "men"
	IrishMiddleCa11001550                                              = "mga"
	MikmaqMicmac                                                       = "mic"
	Minangkabau                                                        = "min"
	UncodedLanguages                                                   = "mis"
	MonKhmerLanguages                                                  = "mkh"
	Malagasy                                                           = "mlg"
	Maltese                                                            = "mlt"
	Manchu                                                             = "mnc"
	Manipuri                                                           = "mni"
	ManoboLanguages                                                    = "mno"
	Mohawk                                                             = "moh"
	MoldavianMoldovan                                                  = "mol"
	Mongolian                                                          = "mon"
	MooréMossi                                                         = "mos"
	MultipleLanguages                                                  = "mul"
	MundaLanguages                                                     = "mun"
	Creek                                                              = "mus"
	Mirandese                                                          = "mwl"
	Marwari                                                            = "mwr"
	MayanLanguages                                                     = "myn"
	Erzya                                                              = "myv"
	NahuatlLanguages                                                   = "nah"
	NorthAmericanIndianLanguages                                       = "nai"
	Neapolitan                                                         = "nap"
	Nauruan                                                            = "nau"
	Navajo                                                             = "nav"
	NdebeleSouth                                                       = "nbl"
	NdebeleNorth                                                       = "nde"
	Ndonga                                                             = "ndo"
	LowGermanLowSaxon                                                  = "nds"
	Nepali                                                             = "nep"
	NewariNepalBhasa                                                   = "new"
	Nias                                                               = "nia"
	NigerKordofanianLanguages                                          = "nic"
	Niuean                                                             = "niu"
	NorwegianNynorsk                                                   = "nno"
	NorwegianBokmål                                                    = "nob"
	Nogai                                                              = "nog"
	OldNorse                                                           = "non"
	Norwegian                                                          = "nor"
	NKo                                                                = "nqo"
	PediSepediNorthernSotho                                            = "nso"
	NubianLanguages                                                    = "nub"
	ClassicalNewariOldNewariClassicalNepalBhasa                        = "nwc"
	ChichewaChewaNyanja                                                = "nya"
	Nyamwezi                                                           = "nym"
	Nyankole                                                           = "nyn"
	Nyoro                                                              = "nyo"
	Nzima                                                              = "nzi"
	OccitanPost1500                                                    = "oci"
	Ojibwa                                                             = "oji"
	Oriya                                                              = "ori"
	Oromo                                                              = "orm"
	Osage                                                              = "osa"
	OssetianOssetic                                                    = "oss"
	TurkishOttoman                                                     = "ota"
	OtomianLanguages                                                   = "oto"
	PapuanLanguages                                                    = "paa"
	Pangasinan                                                         = "pag"
	Pahlavi                                                            = "pal"
	PampangaKapampangan                                                = "pam"
	Panjabi                                                            = "pan"
	Papiamento                                                         = "pap"
	Palauan                                                            = "pau"
	OldPersianCa600400BC                                               = "peo"
	Persian                                                            = "per"
	PhilippineLanguages                                                = "phi"
	Phoenician                                                         = "phn"
	Pali                                                               = "pli"
	Polish                                                             = "pol"
	Ponapeian                                                          = "pon"
	Portuguese                                                         = "por"
	PrakritLanguages                                                   = "pra"
	ProvençalOldTo1500OccitanOldTo1500                                 = "pro"
	PushtoPashto                                                       = "pus"
	Aranés                                                             = "qar"
	Valencian                                                          = "qav"
	Quechua                                                            = "que"
	Rajasthani                                                         = "raj"
	Rapanui                                                            = "rap"
	RarotonganCookIslandsMaori                                         = "rar"
	RomanceLanguages                                                   = "roa"
	Romansh                                                            = "roh"
	Romany                                                             = "rom"
	Romanian                                                           = "rum"
	Rundi                                                              = "run"
	AromanianArumanianMacedoRomanian                                   = "rup"
	Russian                                                            = "rus"
	Sandawe                                                            = "sad"
	Sango                                                              = "sag"
	Yakut                                                              = "sah"
	SouthAmericanIndianLanguages                                       = "sai"
	SalishanLanguages                                                  = "sal"
	SamaritanAramaic                                                   = "sam"
	Sanskrit                                                           = "san"
	Sasak                                                              = "sas"
	Santali                                                            = "sat"
	SerbianDeprecated                                                  = "scc"
	Sicilian                                                           = "scn"
	ScotsLallans                                                       = "sco"
	CroatianDeprecated                                                 = "scr"
	Selkup                                                             = "sel"
	SemiticLanguages                                                   = "sem"
	IrishOldTo1100                                                     = "sga"
	SignLanguages                                                      = "sgn"
	Shan                                                               = "shn"
	Sidamo                                                             = "sid"
	SinhalaSinhalese                                                   = "sin"
	SiouanLanguages                                                    = "sio"
	SinoTibetanLanguages                                               = "sit"
	SlavicLanguages                                                    = "sla"
	Slovak                                                             = "slo"
	Slovenian                                                          = "slv"
	SouthernSami                                                       = "sma"
	NorthernSami                                                       = "sme"
	SamiLanguages                                                      = "smi"
	LuleSami                                                           = "smj"
	InariSami                                                          = "smn"
	Samoan                                                             = "smo"
	SkoltSami                                                          = "sms"
	Shona                                                              = "sna"
	Sindhi                                                             = "snd"
	Soninke                                                            = "snk"
	Sogdian                                                            = "sog"
	Somali                                                             = "som"
	SonghaiLanguages                                                   = "son"
	SothoSesotho                                                       = "sot"
	Spanish                                                            = "spa"
	Sardinian                                                          = "srd"
	SrananTongo                                                        = "srn"
	Serbian                                                            = "srp"
	Serer                                                              = "srr"
	NiloSaharanLanguages                                               = "ssa"
	SwaziSwati                                                         = "ssw"
	Sukuma                                                             = "suk"
	Susu                                                               = "sus"
	Sumerian                                                           = "sux"
	Swahili                                                            = "swa"
	Swedish                                                            = "swe"
	ClassicalSyriac                                                    = "syc"
	Syriac                                                             = "syr"
	Tahitian                                                           = "tah"
	TaiLanguages                                                       = "tai"
	Tamil                                                              = "tam"
	Tatar                                                              = "tat"
	Telugu                                                             = "tel"
	TemneTime                                                          = "tem"
	Terena                                                             = "ter"
	Tetum                                                              = "tet"
	Tajik                                                              = "tgk"
	Tagalog                                                            = "tgl"
	Thai                                                               = "tha"
	Tibetan                                                            = "tib"
	Tigré                                                              = "tig"
	Tigrinya                                                           = "tir"
	Tiv                                                                = "tiv"
	Tokelauan                                                          = "tkl"
	KlingonTlhInganHol                                                 = "tlh"
	Tlingit                                                            = "tli"
	Tamashek                                                           = "tmh"
	TongaNyasa                                                         = "tog"
	Tongan                                                             = "ton"
	TokPisin                                                           = "tpi"
	Tsimshian                                                          = "tsi"
	Tswana                                                             = "tsn"
	Tsonga                                                             = "tso"
	Turkmen                                                            = "tuk"
	Tumbuka                                                            = "tum"
	TupiLanguages                                                      = "tup"
	Turkish                                                            = "tur"
	AltaicLanguages                                                    = "tut"
	Tuvaluan                                                           = "tvl"
	Twi                                                                = "twi"
	Tuvinian                                                           = "tyv"
	Udmurt                                                             = "udm"
	Ugaritic                                                           = "uga"
	UighurUyghur                                                       = "uig"
	Ukrainian                                                          = "ukr"
	Umbundu                                                            = "umb"
	UndeterminedLanguage                                               = "und"
	Urdu                                                               = "urd"
	Uzbek                                                              = "uzb"
	Vai                                                                = "vai"
	Venda                                                              = "ven"
	Vietnamese                                                         = "vie"
	Volapük                                                            = "vol"
	Votic                                                              = "vot"
	WakashanLanguages                                                  = "wak"
	WolaittaWolaytta                                                   = "wal"
	Waray                                                              = "war"
	Washo                                                              = "was"
	Welsh                                                              = "wel"
	SorbianLanguages                                                   = "wen"
	Walloon                                                            = "wln"
	Wolof                                                              = "wol"
	Kalmyk                                                             = "xal"
	Xhosa                                                              = "xho"
	Yao                                                                = "yao"
	Yapese                                                             = "yap"
	Yiddish                                                            = "yid"
	Yoruba                                                             = "yor"
	YupikLanguages                                                     = "ypk"
	Zapotec                                                            = "zap"
	BlissymbolsBlissymbolicsBliss                                      = "zbl"
	Zenaga                                                             = "zen"
	ZhuangChuang                                                       = "zha"
	ZandeLanguages                                                     = "znd"
	Zulu                                                               = "zul"
	Zuni                                                               = "zun"
	NoLinguisticContent                                                = "zxx"
	ZazaDimiliDimliKirdkiKirmanjkiZazaki                               = "zza"
)

var (
	sortedCodes = []string{"aar", "abk", "ace", "ach", "ada", "ady", "afa", "afh", "afr", "ain", "aka", "akk", "alb", "ale", "alg", "alt", "amh", "ang", "anp", "apa", "ara", "arc", "arg", "arm", "arn", "arp", "art", "arw", "asm", "ast", "ath", "aus", "ava", "ave", "awa", "aym", "aze", "bad", "bai", "bak", "bal", "bam", "ban", "baq", "bas", "bat", "bej", "bel", "bem", "ben", "ber", "bho", "bih", "bik", "bin", "bis", "bla", "bnt", "bos", "bra", "bre", "btk", "bua", "bug", "bul", "bur", "byn", "cad", "cai", "car", "cat", "cau", "ceb", "cel", "cha", "chb", "che", "chg", "chi", "chk", "chm", "chn", "cho", "chp", "chr", "chu", "chv", "chy", "cmc", "cop", "cor", "cos", "cpe", "cpf", "cpp", "cre", "crh", "crp", "csb", "cus", "cze", "dak", "dan", "dar", "day", "del", "den", "dgr", "din", "div", "doi", "dra", "dsb", "dua", "dum", "dut", "dyu", "dzo", "efi", "egy", "eka", "elx", "eng", "enm", "epo", "est", "ewe", "ewo", "fan", "fao", "fat", "fij", "fil", "fin", "fiu", "fkv", "fon", "fre", "frm", "fro", "frr", "frs", "fry", "ful", "fur", "gaa", "gay", "gba", "gem", "geo", "ger", "gez", "gil", "gla", "gle", "glg", "glv", "gmh", "goh", "gon", "gor", "got", "grb", "grc", "gre", "grn", "gsw", "guj", "gwi", "hai", "hat", "hau", "haw", "heb", "her", "hil", "him", "hin", "hit", "hmn", "hmo", "hrv", "hsb", "hun", "hup", "iba", "ibo", "ice", "ido", "iii", "ijo", "iku", "ile", "ilo", "ina", "inc", "ind", "ine", "inh", "ipk", "ira", "iro", "ita", "jav", "jbo", "jpn", "jpr", "jrb", "kaa", "kab", "kac", "kal", "kam", "kan", "kar", "kas", "kau", "kaw", "kaz", "kbd", "kha", "khi", "khm", "kho", "kik", "kin", "kir", "kmb", "kok", "kom", "kon", "kor", "kos", "kpe", "krc", "krl", "kro", "kru", "kua", "kum", "kur", "kut", "lad", "lah", "lam", "lao", "lat", "lav", "lez", "lim", "lin", "lit", "lol", "loz", "ltz", "lua", "lub", "lug", "lui", "lun", "luo", "lus", "mac", "mad", "mag", "mah", "mai", "mak", "mal", "man", "mao", "map", "mar", "mas", "may", "mdf", "mdr", "men", "mga", "mic", "min", "mis", "mkh", "mlg", "mlt", "mnc", "mni", "mno", "moh", "mol", "mon", "mos", "mul", "mun", "mus", "mwl", "mwr", "myn", "myv", "nah", "nai", "nap", "nau", "nav", "nbl", "nde", "ndo", "nds", "nep", "new", "nia", "nic", "niu", "nno", "nob", "nog", "non", "nor", "nqo", "nso", "nub", "nwc", "nya", "nym", "nyn", "nyo", "nzi", "oci", "oji", "ori", "orm", "osa", "oss", "ota", "oto", "paa", "pag", "pal", "pam", "pan", "pap", "pau", "peo", "per", "phi", "phn", "pli", "pol", "pon", "por", "pra", "pro", "pus", "qar", "qav", "que", "raj", "rap", "rar", "roa", "roh", "rom", "rum", "run", "rup", "rus", "sad", "sag", "sah", "sai", "sal", "sam", "san", "sas", "sat", "scc", "scn", "sco", "scr", "sel", "sem", "sga", "sgn", "shn", "sid", "sin", "sio", "sit", "sla", "slo", "slv", "sma", "sme", "smi", "smj", "smn", "smo", "sms", "sna", "snd", "snk", "sog", "som", "son", "sot", "spa", "srd", "srn", "srp", "srr", "ssa", "ssw", "suk", "sus", "sux", "swa", "swe", "syc", "syr", "tah", "tai", "tam", "tat", "tel", "tem", "ter", "tet", "tgk", "tgl", "tha", "tib", "tig", "tir", "tiv", "tkl", "tlh", "tli", "tmh", "tog", "ton", "tpi", "tsi", "tsn", "tso", "tuk", "tum", "tup", "tur", "tut", "tvl", "twi", "tyv", "udm", "uga", "uig", "ukr", "umb", "und", "urd", "uzb", "vai", "ven", "vie", "vol", "vot", "wak", "wal", "war", "was", "wel", "wen", "wln", "wol", "xal", "xho", "yao", "yap", "yid", "yor", "ypk", "zap", "zbl", "zen", "zha", "znd", "zul", "zun", "zxx", "zza"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Afar:                 {"Afar", "", "Afar", ""},
		Abkhaz:               {"Abkhaz", "", "Abkhasisk", ""},
		Achinese:             {"Achinese", "", "Achenesisk", ""},
		Acoli:                {"Acoli", "", "Acholi", ""},
		Adangme:              {"Adangme", "", "Dangme", ""},
		Adygei:               {"Adygei", "", "Adygei", ""},
		AfroAsiaticLanguages: {"Afro-Asiatic languages", "Collective name.", "Afroasiatiske språk", "Fellesbetegnelse"},
		Afrihili:             {"Afrihili", "Artificial language.", "Afrihili", "Konstruert språk"},
		Afrikaans:            {"Afrikaans", "", "Afrikaans", ""},
		Ainu:                 {"Ainu", "", "Ainu", ""},
		Akan:                 {"Akan", "", "Akan", ""},
		Akkadian:             {"Akkadian", "", "Akkadisk", ""},
		Albanian:             {"Albanian", "", "Albansk", ""},
		Aleut:                {"Aleut", "", "Aleut", ""},
		AlgonquianLanguages:  {"Algonquian languages", "Collective name.", "Algonkinske språk", "Fellesbetegnelse"},
		SouthernAltai:        {"Southern Altai", "", "Sør-altaisk", ""},
		Amharic:              {"Amharic", "", "Amharisk", ""},
		EnglishOldCa4501100:  {"English, Old (ca. 450-1100)", "", "Angelsaksisk (ca. 450-1100)", ""},
		Angika:               {"Angika", "", "Angika", ""},
		ApacheLanguages:      {"Apache languages", "Collective name.", "Apachespråk", "Fellesbetegnelse"},
		Arabic:               {"Arabic", "", "Arabisk", ""},
		OfficialAramaicImperialAramaic700300BCE: {"Official Aramaic; Imperial Aramaic (700-300 BCE)", "", "Arameisk  (700-300 f.Kr.)", ""},
		Aragonese:           {"Aragonese", "", "Aragonesisk", ""},
		Armenian:            {"Armenian", "", "Armensk", ""},
		MapudungunMapuche:   {"Mapudungun; Mapuche", "", "Mapudungun; Mapuche", ""},
		Arapaho:             {"Arapaho", "", "Arapaho", ""},
		ArtificialLanguages: {"Artificial languages", "Collective name.", "Konstruert språk", "Fellesbetegnelse"},
		Arawak:              {"Arawak", "", "Arawak", ""},
		Assamese:            {"Assamese", "", "Assamesisk", ""},
		AsturianBableLeoneseAsturleonese: {"Asturian; Bable; Leonese; Asturleonese", "", "Asturiansk; Leonesisk; Asturleonesisk", ""},
		AthapascanLanguages:              {"Athapascan languages", "Collective name.", "Athapask-språk", "Fellesbetegnelse"},
		AustralianLanguages:              {"Australian languages", "Collective name.", "Australske språk", "Fellesbetegnelse"},
		Avaric:                           {"Avaric", "", "Avarisk", ""},
		Avestan:                          {"Avestan", "", "Avestisk", ""},
		Awadhi:                           {"Awadhi", "", "Awadhi", ""},
		Aymara:                           {"Aymara", "", "Aymara", ""},
		Azerbaijani:                      {"Azerbaijani", "", "Aserbajdsjansk", ""},
		BandaLanguages:                   {"Banda languages", "Collective name.", "Banda", "Collective name."},
		BamilekeLanguages:                {"Bamileke languages", "Collective name.", "Bamileke-språk", "Collective name."},
		Bashkir:                          {"Bashkir", "", "Basjkirsk", ""},
		Baluchi:                          {"Baluchi", "", "Baluchi", ""},
		Bambara:                          {"Bambara", "", "Bambara", ""},
		Balinese:                         {"Balinese", "", "Balinese", ""},
		Basque:                           {"Basque", "", "Baskisk", ""},
		Basa:                             {"Basa", "", "Basa", ""},
		BalticLanguages:                  {"Baltic languages", "Collective name.", "Baltiske språk", "Fellesbetegnelse"},
		BejaBedawiyet:                    {"Beja; Bedawiyet", "", "Beja; Bedawi", ""},
		Belarusian:                       {"Belarusian", "", "Hviterussisk", ""},
		Bemba:                            {"Bemba", "", "Bemba", ""},
		Bengali:                          {"Bengali", "", "Bengali", ""},
		BerberLanguages:                  {"Berber languages", "Collective name.", "Berberspråk", "Fellesbetegnelse"},
		Bhojpuri:                         {"Bhojpuri", "", "Bhojpuri", ""},
		BihariLanguages:                  {"Bihari languages", "Collective name.", "Bihari", "Fellesbetegnelse"},
		Bikol:                            {"Bikol", "", "Bikol", ""},
		BiniEdo:                          {"Bini; Edo", "", "Bini ; Edo", ""},
		Bislama:                          {"Bislama", "", "Bislama", ""},
		Siksika:                          {"Siksika", "", "Siksika", ""},
		BantuLanguages:                   {"Bantu languages", "Collective name.", "Bantuspråk", "Fellesbetegnelse"},
		Bosnian:                          {"Bosnian", "", "Bosnisk", ""},
		Braj:                             {"Braj", "", "Braj bhasha", ""},
		Breton:                           {"Breton", "", "Bretonsk", ""},
		BatakLanguages:                   {"Batak languages", "Collective name.", "Batak-språk", "Fellesbetegnelse"},
		Buriat:                           {"Buriat", "", "Burjatisk", ""},
		Buginese:                         {"Buginese", "", "Buginesisk", ""},
		Bulgarian:                        {"Bulgarian", "", "Bulgarsk", ""},
		Burmese:                          {"Burmese", "", "Burmesisk", ""},
		BlinBilin:                        {"Blin; Bilin", "", "Blin ; Bilen", ""},
		Caddo:                            {"Caddo", "", "Caddo", ""},
		CentralAmericanIndianLanguages: {"Central American Indian languages", "Collective name.", "Sentralamerikanske indianske språk", "Fellesbetegnelse"},
		GalibiCarib:                    {"Galibi Carib", "", "Karib", ""},
		Catalan:                        {"Catalan", "", "Katalansk", ""},
		CaucasianLanguages:             {"Caucasian languages", "Collective name.", "Kaukasiske språk", "Fellesbetegnelse"},
		Cebuano:                        {"Cebuano", "", "Cebuano", ""},
		CelticLanguages:                {"Celtic languages", "Collective name.", "Keltiske språk", "Fellesbetegnelse"},
		Chamorro:                       {"Chamorro", "", "Chamorro", ""},
		Chibcha:                        {"Chibcha", "", "Chibcha", ""},
		Chechen:                        {"Chechen", "", "Tsjetsjensk", ""},
		Chagatai:                       {"Chagatai", "", "Tsjagataisk", ""},
		Chinese:                        {"Chinese", "", "Kinesisk", ""},
		ChuukeseTruk:                   {"Chuukese (Truk)", "", "Chuukese (Truk)", ""},
		Mari:                           {"Mari", "", "Mari", ""},
		ChinookJargon:                  {"Chinook jargon", "", "Chinook-sjargong", ""},
		Choctaw:                        {"Choctaw", "", "Choctaw", ""},
		ChipewyanDeneSuline:            {"Chipewyan; Dene Suline", "", "Chipewyan", ""},
		Cherokee:                       {"Cherokee", "", "Cherokee", ""},
		ChurchSlavicOldSlavonicChurchSlavonicOldBulgarianOldChurchSlavonic: {"Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic", "", "Kirkeslavisk ; Gammelkirkeslavisk ; Gammelslavisk ; Gammelbulgarsk", ""},
		Chuvash:                          {"Chuvash", "", "Tsjuvasjisk", ""},
		Cheyenne:                         {"Cheyenne", "", "Cheyenne", ""},
		ChamicLanguages:                  {"Chamic languages", "Collective name.", "Chamic-språk", "Fellesbetegnelse"},
		Coptic:                           {"Coptic", "", "Koptisk", ""},
		Cornish:                          {"Cornish", "", "Kornisk", ""},
		Corsican:                         {"Corsican", "", "Korsikansk", ""},
		CreolesAndPidginsEnglishBased:    {"Creoles and pidgins, English-based", "Collective name.", "Kreol- og pidginspråk. Basert på engelsk", "Fellesbetegnelse"},
		CreolesAndPidginsFrenchBased:     {"Creoles and pidgins, French-based", "Collective name.", "Kreol- og pidginspråk. Basert på fransk", "Fellesbetegnelse"},
		CreolesAndPidginsPortugueseBased: {"Creoles and pidgins, Portuguese-based", "Collective name.", "Kreol- og pidginspråk. Basert på portugisisk", "Fellesbetegnelse"},
		Cree: {"Cree", "", "Cree", ""},
		CrimeanTurkishCrimeanTatar: {"Crimean Turkish; Crimean Tatar", "", "Krimtatarisk", ""},
		CreolesAndPidgins:          {"Creoles and pidgins", "Collective name.", "Kreol- og pidginspråk", "Fellesbetegnelse"},
		Kashubian:                  {"Kashubian", "", "Kasjubisk", ""},
		CushiticLanguages:          {"Cushitic languages", "Collective name.", "Kusjittiske språk", "Fellesbetegnelse"},
		Czech:                      {"Czech", "", "Tsjekkisk", ""},
		Dakota:                     {"Dakota", "", "Dakota", ""},
		Danish:                     {"Danish", "", "Dansk", ""},
		Dargwa:                     {"Dargwa", "", "Dargwa", ""},
		LandDayakLanguages:         {"Land Dayak languages", "Collective name.", "Dajak", "Collective name."},
		Delaware:                   {"Delaware", "", "Delaware", ""},
		SlaveAthapascan:            {"Slave (Athapascan)", "", "Slave (Athabaskisk)", ""},
		Dogrib:                     {"Dogrib", "", "Dogrib", ""},
		Dinka:                      {"Dinka", "", "Dinka", ""},
		DivehiDhivehiMaldivian: {"Divehi; Dhivehi; Maldivian", "", "Divehi", ""},
		Dogri:              {"Dogri", "", "Dogri", ""},
		DravidianLanguages: {"Dravidian languages", "Collective name.", "Dravidiske språk", "Fellesbetegnelse"},
		LowerSorbian:       {"Lower Sorbian", "", "Nedersorbisk", ""},
		Duala:              {"Duala", "", "Duala", ""},
		DutchMiddleCa10501350:  {"Dutch, Middle (ca. 1050-1350)", "", "Mellomnederlandsk (ca. 1050-1350)", ""},
		DutchFlemish:           {"Dutch; Flemish", "", "Nederlandsk ; Flamsk", ""},
		Dyula:                  {"Dyula", "", "Dioula", ""},
		Dzongkha:               {"Dzongkha", "", "Dzongkha", ""},
		Efik:                   {"Efik", "", "Efik", ""},
		EgyptianAncient:        {"Egyptian (Ancient)", "", "Egyptisk (oldtiden)", ""},
		Ekajuk:                 {"Ekajuk", "", "Ekajuk", ""},
		Elamite:                {"Elamite", "", "Elamittisk språk", ""},
		English:                {"English", "", "Engelsk", "English"},
		EnglishMiddle11001500:  {"English, Middle (1100-1500)", "", "Mellomengelsk (ca. 1100-1500)", ""},
		Esperanto:              {"Esperanto", "Artificial language.", "Esperanto", "Konstruert språk"},
		Estonian:               {"Estonian", "", "Estisk", ""},
		Ewe:                    {"Ewe", "", "Ewe", ""},
		Ewondo:                 {"Ewondo", "", "Ewondo", ""},
		Fang:                   {"Fang", "", "Fang", ""},
		Faroese:                {"Faroese", "", "Færøysk", ""},
		Fanti:                  {"Fanti", "", "Fanti", ""},
		Fijian:                 {"Fijian", "", "Fijiansk", ""},
		FilipinoPilipino:       {"Filipino; Pilipino", "", "Filippinsk", ""},
		Finnish:                {"Finnish", "", "Finsk", ""},
		FinnoUgrianLanguages:   {"Finno-Ugrian languages", "Collective name.", "Finsk-ugriske språk", "Fellesbetegnelse"},
		Kvensk:                 {"Kvensk", "ONIX local code, equivalent to fkv in ISO 639-3.", "Kvensk", "Tilsvarer kode fkv i ISO 639-3"},
		Fon:                    {"Fon", "", "Fon", ""},
		French:                 {"French", "", "Fransk", ""},
		FrenchMiddleCa14001600: {"French, Middle (ca. 1400-1600)", "", "Mellomfransk (ca. 1400-1600)", ""},
		FrenchOldCa8421400:     {"French, Old (ca. 842-1400)", "", "Gammelfransk (ca. 842-1400)", ""},
		NorthernFrisian:        {"Northern Frisian", "", "Nordfrisisk", ""},
		EasternFrisian:         {"Eastern Frisian", "", "Østfrisisk", ""},
		WesternFrisian:         {"Western Frisian", "", "Vestfrisisk", ""},
		Fulah:                  {"Fulah", "", "Fulfulde (fulani, fula)", ""},
		Friulian:               {"Friulian", "", "Friulisk", ""},
		Gã:                     {"Gã", "", "Gã", ""},
		Gayo:                   {"Gayo", "", "Gayo", ""},
		Gbaya:                  {"Gbaya", "", "Gbaya", ""},
		GermanicLanguages:      {"Germanic languages", "Collective name.", "Germanske språk", "Fellesbetegnelse"},
		Georgian:               {"Georgian", "", "Georgisk", ""},
		German:                 {"German", "", "Tysk", ""},
		EthiopicGeez:           {"Ethiopic (Ge’ez)", "", "Etiopisk (geez)", ""},
		Gilbertese:             {"Gilbertese", "", "Kiribatisk (gilbertesisk)", ""},
		ScottishGaelic:         {"Scottish Gaelic", "", "Skotsk-gælisk", ""},
		Irish:                  {"Irish", "", "Irsk", ""},
		Galician:               {"Galician", "", "Galisisk", ""},
		Manx:                   {"Manx", "", "Mansk", ""},
		GermanMiddleHighCa10501500: {"German, Middle High (ca. 1050-1500)", "", "Middelhøytysk (ca. 1050-1500)", ""},
		GermanOldHighCa7501050:     {"German, Old High (ca. 750-1050)", "", "Gammelhøytysk (ca. 750-1050)", ""},
		Gondi:                {"Gondi", "", "Gondi", ""},
		Gorontalo:            {"Gorontalo", "", "Gorontalo", ""},
		Gothic:               {"Gothic", "", "Gotisk", ""},
		Grebo:                {"Grebo", "", "Grebo", ""},
		GreekAncientTo1453:   {"Greek, Ancient (to 1453)", "", "Gresk, klassisk (fram til 1453)", ""},
		GreekModern1453:      {"Greek, Modern (1453-)", "", "Gresk, moderne (etter 1453)", ""},
		Guarani:              {"Guarani", "", "Guarani", ""},
		SwissGermanAlemannic: {"Swiss German; Alemannic", "", "Sveitsertysk; Alemannisk", ""},
		Gujarati:             {"Gujarati", "", "Gujarati", ""},
		Gwichin:              {"Gwich’in", "", "Gwich'in", ""},
		Haida:                {"Haida", "", "Haida", ""},
		HaitianFrenchCreole:  {"Haitian French Creole", "", "Haitisk kreol (franskbasert)", ""},
		Hausa:                {"Hausa", "", "Hausa", ""},
		Hawaiian:             {"Hawaiian", "", "Hawaiisk", ""},
		Hebrew:               {"Hebrew", "", "Hebraisk", ""},
		Herero:               {"Herero", "", "Herero", ""},
		Hiligaynon:           {"Hiligaynon", "", "Hiligaynon", ""},
		HimachaliLanguagesWesternPahariLanguages: {"Himachali languages; Western Pahari languages", "Collective name.", "Himachali; Pahari", "Fellesbetegnelse"},
		Hindi:                 {"Hindi", "", "Hindi", ""},
		Hittite:               {"Hittite", "", "Hettittisk", ""},
		HmongMong:             {"Hmong; Mong", "", "Hmong", ""},
		HiriMotu:              {"Hiri Motu", "", "Hiri motu", ""},
		Croatian:              {"Croatian", "", "Kroatisk", ""},
		UpperSorbian:          {"Upper Sorbian", "", "Høysorbisk", ""},
		Hungarian:             {"Hungarian", "", "Ungarsk", ""},
		Hupa:                  {"Hupa", "", "Hupa", ""},
		Iban:                  {"Iban", "", "Iban", ""},
		Igbo:                  {"Igbo", "", "Ibo", ""},
		Icelandic:             {"Icelandic", "", "Islandsk", ""},
		Ido:                   {"Ido", "Artificial language.", "Ido", "Konstruert språk"},
		SichuanYiNuosu:        {"Sichuan Yi; Nuosu", "", "Sichuan yi; Nuosu", ""},
		IjoLanguages:          {"Ijo languages", "Collective name.", "Ijo-språk", "Fellesbetegnelse"},
		Inuktitut:             {"Inuktitut", "", "Inuktitut", ""},
		InterlingueOccidental: {"Interlingue; Occidental", "Artificial language.", "Interlingue; Occidental", "Konstruert språk"},
		Iloko: {"Iloko", "", "Iloko", ""},
		InterlinguaInternationalAuxiliaryLanguageAssociation: {"Interlingua (International Auxiliary Language Association)", "Artificial language.", "Interlingua (International Auxiliary Language Association)", "Konstruert språk"},
		IndicLanguages:                                       {"Indic languages", "Collective name.", "Indoariske språk", "Fellesbetegnelse"},
		Indonesian:                                           {"Indonesian", "", "Indonesisk", ""},
		IndoEuropeanLanguages:                                {"Indo-European languages", "Collective name.", "Indo-europesiske språk", "Fellesbetegnelse"},
		Ingush:                                               {"Ingush", "", "Ingusjisk", ""},
		Inupiaq:                                              {"Inupiaq", "", "Inupiak", ""},
		IranianLanguages:                                     {"Iranian languages", "Collective name.", "Iranske språk", "Fellesbetegnelse"},
		IroquoianLanguages:                                   {"Iroquoian languages", "Collective name.", "Irokesisk", "Fellesbetegnelse"},
		Italian:                                              {"Italian", "", "Italiensk", ""},
		Javanese:                                             {"Javanese", "", "Javanesisk", ""},
		Lojban:                                               {"Lojban", "", "Lojban", ""},
		Japanese:                                             {"Japanese", "", "Japansk", ""},
		JudeoPersian:                                         {"Judeo-Persian", "", "Jødepersisk", ""},
		JudeoArabic:                                          {"Judeo-Arabic", "", "Jødearabisk", ""},
		KaraKalpak:                                           {"Kara-Kalpak", "", "Karakalpakisk", ""},
		Kabyle:                                               {"Kabyle", "", "Kabylsk", ""},
		KachinJingpho:                                        {"Kachin; Jingpho", "", "Kachin", ""},
		KalâtdlisutGreenlandic:                               {"Kalâtdlisut; Greenlandic", "", "Grønlandsk", ""},
		Kamba:                   {"Kamba", "", "Kamba", ""},
		Kannada:                 {"Kannada", "", "Kannada", ""},
		KarenLanguages:          {"Karen languages", "Collective name.", "Karenspråk", "Fellesbetegnelse"},
		Kashmiri:                {"Kashmiri", "", "Kasjmiri", ""},
		Kanuri:                  {"Kanuri", "", "Kanuri", ""},
		Kawi:                    {"Kawi", "", "Kawi", ""},
		Kazakh:                  {"Kazakh", "", "Kasakhisk", ""},
		KabardianCircassian:     {"Kabardian (Circassian)", "", "Kabardisk", ""},
		Khasi:                   {"Khasi", "", "Khasi", ""},
		KhoisanLanguages:        {"Khoisan languages", "Collective name.", "Khoisanspråk", "Fellesbetegnelse"},
		CentralKhmer:            {"Central Khmer", "", "Khmer", ""},
		KhotaneseSakan:          {"Khotanese; Sakan", "", "Khotanesisk", ""},
		KikuyuGikuyu:            {"Kikuyu; Gikuyu", "", "Kikuyu", ""},
		Kinyarwanda:             {"Kinyarwanda", "", "Kinjarwanda", ""},
		KirghizKyrgyz:           {"Kirghiz; Kyrgyz", "", "Kirgisisk", ""},
		Kimbundu:                {"Kimbundu", "", "Kimbundu", ""},
		Konkani:                 {"Konkani", "", "Konkani", ""},
		Komi:                    {"Komi", "", "Komi", ""},
		Kongo:                   {"Kongo", "", "Kikongo; Kongo", ""},
		Korean:                  {"Korean", "", "Koreansk", ""},
		KusaieanCarolineIslands: {"Kusaiean (Caroline Islands)", "", "Kosraeansk\u00a0", ""},
		Kpelle:                  {"Kpelle", "", "Kpelle", ""},
		KarachayBalkar:          {"Karachay-Balkar", "", "Karachay-Balkar", ""},
		Karelian:                {"Karelian", "", "Karelsk", ""},
		KruLanguages:            {"Kru languages", "Collective name.", "Kruspråk", "Fellesbetegnelse"},
		Kurukh:                  {"Kurukh", "", "Kurukh", ""},
		Kuanyama:                {"Kuanyama", "", "Kuanyama", ""},
		Kumyk:                   {"Kumyk", "", "Kumyk", ""},
		Kurdish:                 {"Kurdish", "", "Kurdisk", ""},
		Kutenai:                 {"Kutenai", "", "Kutenai", ""},
		Ladino:                  {"Ladino", "", "Ladinsk", ""},
		Lahnda:                  {"Lahnda", "", "Lahnda", ""},
		Lamba:                   {"Lamba", "", "Lamba", ""},
		Lao:                     {"Lao", "", "Lao; Laotisk", ""},
		Latin:                   {"Latin", "", "Latin", ""},
		Latvian:                 {"Latvian", "", "Latvisk", ""},
		Lezgian:                 {"Lezgian", "", "Lezghian", ""},
		Limburgish:              {"Limburgish", "", "Limburgisk", ""},
		Lingala:                 {"Lingala", "", "Lingala", ""},
		Lithuanian:              {"Lithuanian", "", "Litauisk", ""},
		MongoNkundu:             {"Mongo-Nkundu", "", "Mongo", ""},
		Lozi:                    {"Lozi", "", "Lozi", ""},
		LuxembourgishLetzeburgesch: {"Luxembourgish; Letzeburgesch", "", "Luxemburgsk", ""},
		LubaLulua:                  {"Luba-Lulua", "", "Luba-Lulua", ""},
		LubaKatanga:                {"Luba-Katanga", "", "Luba-Katanga", ""},
		Ganda:                      {"Ganda", "", "Ganda", ""},
		Luiseño:                    {"Luiseño", "", "Luiseño", ""},
		Lunda:                      {"Lunda", "", "Lunda", ""},
		LuoKenyaAndTanzania:        {"Luo (Kenya and Tanzania)", "", "Luo", ""},
		Lushai:                     {"Lushai", "", "Lushai", ""},
		Macedonian:                 {"Macedonian", "", "Makedonsk", ""},
		Madurese:                   {"Madurese", "", "Maduresisk", ""},
		Magahi:                     {"Magahi", "", "Magahi", ""},
		Marshallese:                {"Marshallese", "", "Marshallesisk", ""},
		Maithili:                   {"Maithili", "", "Maithili", ""},
		Makasar:                    {"Makasar", "", "Makasar", ""},
		Malayalam:                  {"Malayalam", "", "Malayalam", ""},
		Mandingo:                   {"Mandingo", "", "Mandingo", ""},
		Maori:                      {"Maori", "", "Maori", ""},
		AustronesianLanguages: {"Austronesian languages", "Collective name.", "Austronesiske språk", "Fellesbetegnelse"},
		Marathi:               {"Marathi", "", "Marathi", ""},
		Masai:                 {"Masai", "", "Masai", ""},
		Malay:                 {"Malay", "", "Malayisk", ""},
		Moksha:                {"Moksha", "", "Moksha", ""},
		Mandar:                {"Mandar", "", "Mandar", ""},
		Mende:                 {"Mende", "", "Mende", ""},
		IrishMiddleCa11001550:        {"Irish, Middle (ca. 1100-1550)", "", "Mellomirsk (ca. 1100-1550)", ""},
		MikmaqMicmac:                 {"Mi’kmaq; Micmac", "", "Micmac", ""},
		Minangkabau:                  {"Minangkabau", "", "Minangkabau", ""},
		UncodedLanguages:             {"Uncoded languages", "", "Ukodete språk", ""},
		MonKhmerLanguages:            {"Mon-Khmer languages", "Collective name.", "Mon-Khmer-språk", "Fellesbetegnelse"},
		Malagasy:                     {"Malagasy", "", "Madagassisk", ""},
		Maltese:                      {"Maltese", "", "Maltesisk", ""},
		Manchu:                       {"Manchu", "", "Mandsju", ""},
		Manipuri:                     {"Manipuri", "", "Manipuri", ""},
		ManoboLanguages:              {"Manobo languages", "Collective name.", "Manobospråk", "Fellesbetegnelse"},
		Mohawk:                       {"Mohawk", "", "Mohawk", ""},
		MoldavianMoldovan:            {"Moldavian; Moldovan", "DEPRECATED – use rum.", "Moldovsk", "UTGÅTT. Bruk rum"},
		Mongolian:                    {"Mongolian", "", "Mongolsk", ""},
		MooréMossi:                   {"Mooré; Mossi", "", "Mossi", ""},
		MultipleLanguages:            {"Multiple languages", "", "Flerspråklig", ""},
		MundaLanguages:               {"Munda languages", "Collective name.", "Mundaspråk", "Fellesbetegnelse"},
		Creek:                        {"Creek", "", "Creek", ""},
		Mirandese:                    {"Mirandese", "", "Mirandesisk", ""},
		Marwari:                      {"Marwari", "", "Marwari", ""},
		MayanLanguages:               {"Mayan languages", "Collective name.", "Mayaspråk", "Fellesbetegnelse"},
		Erzya:                        {"Erzya", "", "Erzya", ""},
		NahuatlLanguages:             {"Nahuatl languages", "Collective name.", "Nahuatlspråk", "Fellesbetegnelse"},
		NorthAmericanIndianLanguages: {"North American Indian languages", "Collective name.", "Nordamerikanske, urfolks språk", "Fellesbetegnelse"},
		Neapolitan:                   {"Neapolitan", "", "Napolitansk", ""},
		Nauruan:                      {"Nauruan", "", "Nauru", ""},
		Navajo:                       {"Navajo", "", "Navajo", ""},
		NdebeleSouth:                 {"Ndebele, South", "", "Sørndebele", ""},
		NdebeleNorth:                 {"Ndebele, North", "", "Nordndebele", ""},
		Ndonga:                       {"Ndonga", "", "Ndonga", ""},
		LowGermanLowSaxon:            {"Low German; Low Saxon", "", "Nedertysk", ""},
		Nepali:                       {"Nepali", "", "Nepalsk", ""},
		NewariNepalBhasa:             {"Newari; Nepal Bhasa", "", "Newari", ""},
		Nias:                         {"Nias", "", "Nias", ""},
		NigerKordofanianLanguages: {"Niger-Kordofanian languages", "Collective name.", "Niger-Kongospråk", "Fellesbetegnelse"},
		Niuean:           {"Niuean", "", "Niuisk", ""},
		NorwegianNynorsk: {"Norwegian Nynorsk", "", "Norsk, nynorsk", "Brukes også for landsmål og dialekter"},
		NorwegianBokmål:  {"Norwegian Bokmål", "", "Norsk, bokmål", "Brukes også for riksmål"},
		Nogai:            {"Nogai", "", "Nogai", ""},
		OldNorse:         {"Old Norse", "", "Gammelnorsk", ""},
		Norwegian:        {"Norwegian", "", "Norsk", "Brukes dersom det ikke skilles mellom bokmål og nynorsk"},
		NKo:              {"N’Ko", "", "N-kå", ""},
		PediSepediNorthernSotho:                     {"Pedi; Sepedi; Northern Sotho", "", "Nordsotho", ""},
		NubianLanguages:                             {"Nubian languages", "Collective name.", "Nubiske språk", "Fellesbetegnelse"},
		ClassicalNewariOldNewariClassicalNepalBhasa: {"Classical Newari; Old Newari; Classical Nepal Bhasa", "", "Klassisk newari", ""},
		ChichewaChewaNyanja:                         {"Chichewa; Chewa; Nyanja", "", "Nyanja", ""},
		Nyamwezi:                                    {"Nyamwezi", "", "Nyamwezi", ""},
		Nyankole:                                    {"Nyankole", "", "Nyankole", ""},
		Nyoro:                                       {"Nyoro", "", "Nyoro", ""},
		Nzima:                                       {"Nzima", "", "Nzima", ""},
		OccitanPost1500:                             {"Occitan (post 1500)", "", "Oksitansk (etter 1500)", ""},
		Ojibwa:                                      {"Ojibwa", "", "Ojibwa", ""},
		Oriya:                                       {"Oriya", "", "Oriya", ""},
		Oromo:                                       {"Oromo", "", "Oromo", ""},
		Osage:                                       {"Osage", "", "Osage", ""},
		OssetianOssetic:                             {"Ossetian; Ossetic", "", "Ossetisk", ""},
		TurkishOttoman:                              {"Turkish, Ottoman", "", "Ottomansk tyrkisk", ""},
		OtomianLanguages:                            {"Otomian languages", "Collective name.", "Ottomanske språk", "Fellesbetegnelse"},
		PapuanLanguages:                             {"Papuan languages", "Collective name.", "Papuanske språk", "Fellesbetegnelse"},
		Pangasinan:                                  {"Pangasinan", "", "Pangasinan", ""},
		Pahlavi:                                     {"Pahlavi", "", "Pahlavi", ""},
		PampangaKapampangan:                         {"Pampanga; Kapampangan", "", "Pampanga", ""},
		Panjabi:                                     {"Panjabi", "", "Panjabi", ""},
		Papiamento:                                  {"Papiamento", "", "Papiamento", ""},
		Palauan:                                     {"Palauan", "", "Palauisk", ""},
		OldPersianCa600400BC:                        {"Old Persian (ca. 600-400 B.C.)", "", "Gammelpersisk (ca. 600-400 f.Kr.)", ""},
		Persian:                                     {"Persian", "", "Persisk; Farsi", ""},
		PhilippineLanguages:                         {"Philippine languages", "Collective name.", "Filippinske språk", "Fellesbetegnelse"},
		Phoenician:                                  {"Phoenician", "", "Fønikisk", ""},
		Pali:                                        {"Pali", "", "Pali", ""},
		Polish:                                      {"Polish", "", "Polsk", ""},
		Ponapeian:                                   {"Ponapeian", "", "Ponapisk", ""},
		Portuguese:                                  {"Portuguese", "", "Portugisisk", ""},
		PrakritLanguages:                            {"Prakrit languages", "Collective name.", "Prakritspråk", "Fellesbetegnelse"},
		ProvençalOldTo1500OccitanOldTo1500:          {"Provençal, Old (to 1500); Occitan, Old (to 1500)", "", "Gammelprovençalsk (til 1500); Gammeloksitansk (til 1500)", ""},
		PushtoPashto:                                {"Pushto; Pashto", "", "Pashto; Pushto", ""},
		Aranés:                                      {"Aranés", "ONIX local code, distinct dialect of Occitan (not distinguished from oci by ISO 639-3).", "Aranesisk", "Lokal ONIX-kode for en bestemt dialekt av oksitansk (ikke forskjellig fra oci av ISO-639-3)"},
		Valencian:                                   {"Valencian", "ONIX local code, distinct dialect of Catalan (not distinguished from cat by ISO 639-3).", "Valensiansk", "Lokal ONIX-kode for en bestemt dialekt av katalansk (ikke forskjellig fra cat i ISO-639-3)"},
		Quechua:                                     {"Quechua", "", "Quechua", ""},
		Rajasthani:                                  {"Rajasthani", "", "Rajasthani", ""},
		Rapanui:                                     {"Rapanui", "", "Rapanui", ""},
		RarotonganCookIslandsMaori:                  {"Rarotongan; Cook Islands Maori", "", "Rarotongansk\u00a0", ""},
		RomanceLanguages:                            {"Romance languages", "Collective name.", "Romanske språk", "Fellesbetegnelse"},
		Romansh:                                     {"Romansh", "", "Retoromansk", ""},
		Romany:                                      {"Romany", "", "Romani", ""},
		Romanian:                                    {"Romanian", "", "Rumensk", ""},
		Rundi:                                       {"Rundi", "", "Rundi", ""},
		AromanianArumanianMacedoRomanian: {"Aromanian; Arumanian; Macedo-Romanian", "", "Armensk; Arumensk", ""},
		Russian: {"Russian", "", "Russisk", ""},
		Sandawe: {"Sandawe", "", "Sandawe", ""},
		Sango:   {"Sango", "", "Sango", ""},
		Yakut:   {"Yakut", "", "Jakutsk", ""},
		SouthAmericanIndianLanguages: {"South American Indian languages", "Collective name.", "Søramerikanske, indianske språk", "Fellesbetegnelse"},
		SalishanLanguages:            {"Salishan languages", "Collective name.", "Salishanspråk", "Fellesbetegnelse"},
		SamaritanAramaic:             {"Samaritan Aramaic", "", "Samaritansk arameisk", ""},
		Sanskrit:                     {"Sanskrit", "", "Sanskrit", ""},
		Sasak:                        {"Sasak", "", "Sasak", ""},
		Santali:                      {"Santali", "", "Santali", ""},
		SerbianDeprecated:            {"Serbian", "DEPRECATED – use srp.", "Serbisk", "UTGÅTT. Bruk srp"},
		Sicilian:                     {"Sicilian", "", "Siciliansk", ""},
		ScotsLallans:                 {"Scots (lallans)", "", "Skotsk (lallans)", ""},
		CroatianDeprecated:           {"Croatian", "DEPRECATED – use hrv.", "Kroatisk", "UTGÅTT. Bruk hrv"},
		Selkup:                       {"Selkup", "", "Selkupisk", ""},
		SemiticLanguages:             {"Semitic languages", "Collective name.", "Semittiske språk", "Fellesbetegnelse"},
		IrishOldTo1100:               {"Irish, Old (to 1100)", "", "Gammelirsk (til 1100)", ""},
		SignLanguages:                {"Sign languages", "Collective name.", "Tegnspråk", "Fellesbetegnelse"},
		Shan:                         {"Shan", "", "Shan", ""},
		Sidamo:                       {"Sidamo", "", "Sidamo", ""},
		SinhalaSinhalese:             {"Sinhala; Sinhalese", "", "Singalesisk", ""},
		SiouanLanguages:              {"Siouan languages", "Collective name.", "Siouxspråk", "Fellesbetegnelse"},
		SinoTibetanLanguages:         {"Sino-Tibetan languages", "Collective name.", "Sino-tibetanske språk", "Fellesbetegnelse"},
		SlavicLanguages:              {"Slavic languages", "Collective name.", "Slaviske språk", "Fellesbetegnelse"},
		Slovak:                       {"Slovak", "", "Slovakisk", ""},
		Slovenian:                    {"Slovenian", "", "Slovensk", ""},
		SouthernSami:                 {"Southern Sami", "", "Sørsamisk", ""},
		NorthernSami:                 {"Northern Sami", "", "Nordsamisk", ""},
		SamiLanguages:                {"Sami languages", "Collective name.", "Samiske språk", "Fellesbetegnelse"},
		LuleSami:                     {"Lule Sami", "", "Lulesamisk", ""},
		InariSami:                    {"Inari Sami", "", "Enaresamisk", ""},
		Samoan:                       {"Samoan", "", "Samoansk", ""},
		SkoltSami:                    {"Skolt Sami", "", "Skoltesamisk", ""},
		Shona:                        {"Shona", "", "Shona", ""},
		Sindhi:                       {"Sindhi", "", "Sindhi", ""},
		Soninke:                      {"Soninke", "", "Soninke", ""},
		Sogdian:                      {"Sogdian", "", "Sogdisk", ""},
		Somali:                       {"Somali", "", "Somali", ""},
		SonghaiLanguages:             {"Songhai languages", "Collective name.", "Songhaispråk", "Collective name."},
		SothoSesotho:                 {"Sotho; Sesotho", "", "Sotho; Sesotho", ""},
		Spanish:                      {"Spanish", "", "Spansk", ""},
		Sardinian:                    {"Sardinian", "", "Sardinsk", ""},
		SrananTongo:                  {"Sranan Tongo", "", "Sranan Tongo", ""},
		Serbian:                      {"Serbian", "", "Serbisk", ""},
		Serer:                        {"Serer", "", "Serer", ""},
		NiloSaharanLanguages:         {"Nilo-Saharan languages", "Collective name.", "Nilosahariske språk\u00a0", "Fellesbetegnelse"},
		SwaziSwati:                   {"Swazi; Swati", "", "Swati", ""},
		Sukuma:                       {"Sukuma", "", "Sukuma", ""},
		Susu:                         {"Susu", "", "Susu", ""},
		Sumerian:                     {"Sumerian", "", "Sumerisk", ""},
		Swahili:                      {"Swahili", "", "Swahili", ""},
		Swedish:                      {"Swedish", "", "Svensk", ""},
		ClassicalSyriac:              {"Classical Syriac", "", "Gammelsyrisk", ""},
		Syriac:                       {"Syriac", "", "Syrisk", ""},
		Tahitian:                     {"Tahitian", "", "Tahitisk", ""},
		TaiLanguages:                 {"Tai languages", "Collective name.", "Tai-språk", "Fellesbetegnelse"},
		Tamil:                        {"Tamil", "", "Tamil", ""},
		Tatar:                        {"Tatar", "", "Tatarisk", ""},
		Telugu:                       {"Telugu", "", "Telugu", ""},
		TemneTime:                    {"Temne; Time", "", "Temne", ""},
		Terena:                       {"Terena", "", "Terena", ""},
		Tetum:                        {"Tetum", "", "Tetum", ""},
		Tajik:                        {"Tajik", "", "Tadsjikisk", ""},
		Tagalog:                      {"Tagalog", "", "Tagalog", ""},
		Thai:                         {"Thai", "", "Thai", ""},
		Tibetan:                      {"Tibetan", "", "Tibetansk", ""},
		Tigré:                        {"Tigré", "", "Tigré", ""},
		Tigrinya:                     {"Tigrinya", "", "Tigrinja", ""},
		Tiv:                          {"Tiv", "", "Tiv", ""},
		Tokelauan:                    {"Tokelauan", "", "Tokelauisk", ""},
		KlingonTlhInganHol:           {"Klingon; tlhIngan-Hol", "Artificial language.", "Klingon", "Konstruert språk"},
		Tlingit:                      {"Tlingit", "", "Tlingit", ""},
		Tamashek:                     {"Tamashek", "", "Tamasjek", ""},
		TongaNyasa:                   {"Tonga (Nyasa)", "", "Tonga (Nyasa)", ""},
		Tongan:                       {"Tongan", "", "Tongansk", ""},
		TokPisin:                     {"Tok Pisin", "", "Tok Pisin", ""},
		Tsimshian:                    {"Tsimshian", "", "Tsimshian", ""},
		Tswana:                       {"Tswana", "AKA Setswana.", "Setswana", "AKA Setswana."},
		Tsonga:                       {"Tsonga", "", "Tsonga", ""},
		Turkmen:                      {"Turkmen", "", "Turkmensk", ""},
		Tumbuka:                      {"Tumbuka", "", "Tumbuka", ""},
		TupiLanguages:                {"Tupi languages", "Collective name.", "Tupianske språk", "Fellesbetegnelse"},
		Turkish:                      {"Turkish", "", "Tyrkisk", ""},
		AltaicLanguages:              {"Altaic languages", "", "Altaiske språk", ""},
		Tuvaluan:                     {"Tuvaluan", "", "Tuvalu", ""},
		Twi:                          {"Twi", "", "Twi", ""},
		Tuvinian:                     {"Tuvinian", "", "Tuvinsk", ""},
		Udmurt:                       {"Udmurt", "", "Udmurtisk", ""},
		Ugaritic:                     {"Ugaritic", "", "Ugarittisk", ""},
		UighurUyghur:                 {"Uighur; Uyghur", "", "Uigurisk", ""},
		Ukrainian:                    {"Ukrainian", "", "Ukrainsk", ""},
		Umbundu:                      {"Umbundu", "", "Umbundu", ""},
		UndeterminedLanguage:         {"Undetermined language", "", "Ubestemt", ""},
		Urdu:                         {"Urdu", "", "Urdu", ""},
		Uzbek:                        {"Uzbek", "", "Usbekisk", ""},
		Vai:                          {"Vai", "", "Vai", ""},
		Venda:                        {"Venda", "", "Venda", ""},
		Vietnamese:                   {"Vietnamese", "", "Vietnamesisk", ""},
		Volapük:                      {"Volapük", "Artificial language.", "Volapük", "Konstruert språk"},
		Votic:                        {"Votic", "", "Votisk", ""},
		WakashanLanguages:            {"Wakashan languages", "Collective name.", "Wakashan-språk", "Fellesbetegnelse"},
		WolaittaWolaytta:             {"Wolaitta; Wolaytta", "", "Walamo", ""},
		Waray:                        {"Waray", "", "Waray", ""},
		Washo:                        {"Washo", "", "Washo", ""},
		Welsh:                        {"Welsh", "", "Walisisk", ""},
		SorbianLanguages:             {"Sorbian languages", "Collective name.", "Sorbiske språk", "Fellesbetegnelse"},
		Walloon:                      {"Walloon", "", "Vallonsk", ""},
		Wolof:                        {"Wolof", "", "Wolof", ""},
		Kalmyk:                       {"Kalmyk", "", "Kalmykisk", ""},
		Xhosa:                        {"Xhosa", "", "Xhosa", ""},
		Yao:                          {"Yao", "", "Yao", ""},
		Yapese:                       {"Yapese", "", "Yapesisk", ""},
		Yiddish:                      {"Yiddish", "", "Jiddisk (Jiddisch)", ""},
		Yoruba:                       {"Yoruba", "", "Joruba", ""},
		YupikLanguages:               {"Yupik languages", "Collective name.", "Yupik-språk", "Fellesbetegnelse"},
		Zapotec:                      {"Zapotec", "", "Zapotec", ""},
		BlissymbolsBlissymbolicsBliss: {"Blissymbols; Blissymbolics; Bliss", "Artificial language.", "Bliss symbolspråk", "Konstruert språk"},
		Zenaga:                               {"Zenaga", "", "Zenaga", ""},
		ZhuangChuang:                         {"Zhuang; Chuang", "", "Zhuang", ""},
		ZandeLanguages:                       {"Zande languages", "Collective name.", "Zande-språk", "Fellesbetegnelse"},
		Zulu:                                 {"Zulu", "", "Zulu", ""},
		Zuni:                                 {"Zuni", "", "Zuni", ""},
		NoLinguisticContent:                  {"No linguistic content", "", "Har ikke språklig innhold", ""},
		ZazaDimiliDimliKirdkiKirmanjkiZazaki: {"Zaza; Dimili; Dimli; Kirdki; Kirmanjki; Zazaki", "", "Zaza; Dimli; Kirmancki; Zazaisk", ""},
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
