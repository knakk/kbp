package list91

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Andorra                                = "AD"
	UnitedArabEmirates                     = "AE"
	Afghanistan                            = "AF"
	AntiguaAndBarbuda                      = "AG"
	Anguilla                               = "AI"
	Albania                                = "AL"
	Armenia                                = "AM"
	NetherlandsAntilles                    = "AN"
	Angola                                 = "AO"
	Antarctica                             = "AQ"
	Argentina                              = "AR"
	AmericanSamoa                          = "AS"
	Austria                                = "AT"
	Australia                              = "AU"
	Aruba                                  = "AW"
	ÅlandIslands                           = "AX"
	Azerbaijan                             = "AZ"
	BosniaAndHerzegovina                   = "BA"
	Barbados                               = "BB"
	Bangladesh                             = "BD"
	Belgium                                = "BE"
	BurkinaFaso                            = "BF"
	Bulgaria                               = "BG"
	Bahrain                                = "BH"
	Burundi                                = "BI"
	Benin                                  = "BJ"
	SaintBarthélemy                        = "BL"
	Bermuda                                = "BM"
	BruneiDarussalam                       = "BN"
	BoliviaPlurinationalStateOf            = "BO"
	BonaireSintEustatiusAndSaba            = "BQ"
	Brazil                                 = "BR"
	Bahamas                                = "BS"
	Bhutan                                 = "BT"
	BouvetIsland                           = "BV"
	Botswana                               = "BW"
	Belarus                                = "BY"
	Belize                                 = "BZ"
	Canada                                 = "CA"
	CocosKeelingIslands                    = "CC"
	CongoDemocraticRepublicOfThe           = "CD"
	CentralAfricanRepublic                 = "CF"
	Congo                                  = "CG"
	Switzerland                            = "CH"
	CoteDIvoire                            = "CI"
	CookIslands                            = "CK"
	Chile                                  = "CL"
	Cameroon                               = "CM"
	China                                  = "CN"
	Colombia                               = "CO"
	CostaRica                              = "CR"
	SerbiaAndMontenegro                    = "CS"
	Cuba                                   = "CU"
	CapeVerde                              = "CV"
	Curaçao                                = "CW"
	ChristmasIsland                        = "CX"
	Cyprus                                 = "CY"
	CzechRepublic                          = "CZ"
	Germany                                = "DE"
	Djibouti                               = "DJ"
	Denmark                                = "DK"
	Dominica                               = "DM"
	DominicanRepublic                      = "DO"
	Algeria                                = "DZ"
	Ecuador                                = "EC"
	Estonia                                = "EE"
	Egypt                                  = "EG"
	WesternSahara                          = "EH"
	Eritrea                                = "ER"
	Spain                                  = "ES"
	Ethiopia                               = "ET"
	Finland                                = "FI"
	Fiji                                   = "FJ"
	FalklandIslandsMalvinas                = "FK"
	MicronesiaFederatedStatesOf            = "FM"
	FaroeIslands                           = "FO"
	France                                 = "FR"
	Gabon                                  = "GA"
	UnitedKingdom                          = "GB"
	Grenada                                = "GD"
	Georgia                                = "GE"
	FrenchGuiana                           = "GF"
	Guernsey                               = "GG"
	Ghana                                  = "GH"
	Gibraltar                              = "GI"
	Greenland                              = "GL"
	Gambia                                 = "GM"
	Guinea                                 = "GN"
	Guadeloupe                             = "GP"
	EquatorialGuinea                       = "GQ"
	Greece                                 = "GR"
	SouthGeorgiaAndTheSouthSandwichIslands = "GS"
	Guatemala                              = "GT"
	Guam                                   = "GU"
	GuineaBissau                           = "GW"
	Guyana                                 = "GY"
	HongKong                               = "HK"
	HeardIslandAndMcDonaldIslands          = "HM"
	Honduras                               = "HN"
	Croatia                                = "HR"
	Haiti                                  = "HT"
	Hungary                                = "HU"
	Indonesia                              = "ID"
	Ireland                                = "IE"
	Israel                                 = "IL"
	IsleOfMan                              = "IM"
	India                                  = "IN"
	BritishIndianOceanTerritory            = "IO"
	Iraq                                   = "IQ"
	IranIslamicRepublicOf                  = "IR"
	Iceland                                = "IS"
	Italy                                  = "IT"
	Jersey                                 = "JE"
	Jamaica                                = "JM"
	Jordan                                 = "JO"
	Japan                                  = "JP"
	Kenya                                  = "KE"
	Kyrgyzstan                             = "KG"
	Cambodia                               = "KH"
	Kiribati                               = "KI"
	Comoros                                = "KM"
	SaintKittsAndNevis                     = "KN"
	KoreaDemocraticPeoplesRepublicOf       = "KP"
	KoreaRepublicOf                        = "KR"
	Kuwait                                 = "KW"
	CaymanIslands                          = "KY"
	Kazakhstan                             = "KZ"
	LaoPeoplesDemocraticRepublic           = "LA"
	Lebanon                                = "LB"
	SaintLucia                             = "LC"
	Liechtenstein                          = "LI"
	SriLanka                               = "LK"
	Liberia                                = "LR"
	Lesotho                                = "LS"
	Lithuania                              = "LT"
	Luxembourg                             = "LU"
	Latvia                                 = "LV"
	Libya                                  = "LY"
	Morocco                                = "MA"
	Monaco                                 = "MC"
	MoldovaRepubicOf                       = "MD"
	Montenegro                             = "ME"
	SaintMartinFrenchPart                  = "MF"
	Madagascar                             = "MG"
	MarshallIslands                        = "MH"
	MacedoniaTheFormerYugoslavRepublicOf   = "MK"
	Mali                                   = "ML"
	Myanmar                                = "MM"
	Mongolia                               = "MN"
	Macao                                  = "MO"
	NorthernMarianaIslands                 = "MP"
	Martinique                             = "MQ"
	Mauritania                             = "MR"
	Montserrat                             = "MS"
	Malta                                  = "MT"
	Mauritius                              = "MU"
	Maldives                               = "MV"
	Malawi                                 = "MW"
	Mexico                                 = "MX"
	Malaysia                               = "MY"
	Mozambique                             = "MZ"
	Namibia                                = "NA"
	NewCaledonia                           = "NC"
	Niger                                  = "NE"
	NorfolkIsland                          = "NF"
	Nigeria                                = "NG"
	Nicaragua                              = "NI"
	Netherlands                            = "NL"
	Norway                                 = "NO"
	Nepal                                  = "NP"
	Nauru                                  = "NR"
	Niue                                   = "NU"
	NewZealand                             = "NZ"
	Oman                                   = "OM"
	Panama                                 = "PA"
	Peru                                   = "PE"
	FrenchPolynesia                        = "PF"
	PapuaNewGuinea                         = "PG"
	Philippines                            = "PH"
	Pakistan                               = "PK"
	Poland                                 = "PL"
	SaintPierreAndMiquelon                 = "PM"
	Pitcairn                               = "PN"
	PuertoRico                             = "PR"
	PalestineStateOf                       = "PS"
	Portugal                               = "PT"
	Palau                                  = "PW"
	Paraguay                               = "PY"
	Qatar                                  = "QA"
	Réunion                                = "RE"
	Romania                                = "RO"
	Serbia                                 = "RS"
	RussianFederation                      = "RU"
	Rwanda                                 = "RW"
	SaudiArabia                            = "SA"
	SolomonIslands                         = "SB"
	Seychelles                             = "SC"
	Sudan                                  = "SD"
	Sweden                                 = "SE"
	Singapore                              = "SG"
	SaintHelenaAscensionAndTristanDaCunha  = "SH"
	Slovenia                               = "SI"
	SvalbardAndJanMayen                    = "SJ"
	Slovakia                               = "SK"
	SierraLeone                            = "SL"
	SanMarino                              = "SM"
	Senegal                                = "SN"
	Somalia                                = "SO"
	Suriname                               = "SR"
	SouthSudan                             = "SS"
	SaoTomeAndPrincipe                     = "ST"
	ElSalvador                             = "SV"
	SintMaartenDutchPart                   = "SX"
	SyrianArabRepublic                     = "SY"
	Swaziland                              = "SZ"
	TurksAndCaicosIslands                  = "TC"
	Chad                                   = "TD"
	FrenchSouthernTerritories              = "TF"
	Togo                                   = "TG"
	Thailand                               = "TH"
	Tajikistan                             = "TJ"
	Tokelau                                = "TK"
	TimorLeste                             = "TL"
	Turkmenistan                           = "TM"
	Tunisia                                = "TN"
	Tonga                                  = "TO"
	Turkey                                 = "TR"
	TrinidadAndTobago                      = "TT"
	Tuvalu                                 = "TV"
	TaiwanProvinceOfChina                  = "TW"
	TanzaniaUnitedRepublicOf               = "TZ"
	Ukraine                                = "UA"
	Uganda                                 = "UG"
	UnitedStatesMinorOutlyingIslands       = "UM"
	UnitedStates                           = "US"
	Uruguay                                = "UY"
	Uzbekistan                             = "UZ"
	HolySeeVaticanCityState                = "VA"
	SaintVincentAndTheGrenadines           = "VC"
	VenezuelaBolivarianRepublicOf          = "VE"
	VirginIslandsBritish                   = "VG"
	VirginIslandsUS                        = "VI"
	VietNam                                = "VN"
	Vanuatu                                = "VU"
	WallisAndFutuna                        = "WF"
	Samoa                                  = "WS"
	Yemen                                  = "YE"
	Mayotte                                = "YT"
	Yugoslavia                             = "YU"
	SouthAfrica                            = "ZA"
	Zambia                                 = "ZM"
	Zimbabwe                               = "ZW"
)

var (
	sortedCodes = []string{"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AN", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CS", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "YU", "ZA", "ZM", "ZW"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Andorra:                     {"Andorra", "", "Andorra", ""},
		UnitedArabEmirates:          {"United Arab Emirates", "", "De Forente arabiske emirater", ""},
		Afghanistan:                 {"Afghanistan", "", "Afghanistan", ""},
		AntiguaAndBarbuda:           {"Antigua and Barbuda", "", "Antigua og Barbuda", ""},
		Anguilla:                    {"Anguilla", "", "Anguilla", ""},
		Albania:                     {"Albania", "", "Albania", ""},
		Armenia:                     {"Armenia", "", "Armenia", ""},
		NetherlandsAntilles:         {"Netherlands Antilles", "Deprecated – use BQ, CW or SX as appropriate.", "Antillene", "Deprecated – use BQ, CW or SX as appropriate."},
		Angola:                      {"Angola", "", "Angola", ""},
		Antarctica:                  {"Antarctica", "", "Antarktis", ""},
		Argentina:                   {"Argentina", "", "Argentina", ""},
		AmericanSamoa:               {"American Samoa", "", "Amerikansk Samoa", ""},
		Austria:                     {"Austria", "", "Østerrike", ""},
		Australia:                   {"Australia", "", "Australia", ""},
		Aruba:                       {"Aruba", "", "Aruba", ""},
		ÅlandIslands:                {"Åland Islands", "", "Åland", ""},
		Azerbaijan:                  {"Azerbaijan", "", "Aserbajdsjan", ""},
		BosniaAndHerzegovina:        {"Bosnia and Herzegovina", "", "Bosnia og Hercegovina", ""},
		Barbados:                    {"Barbados", "", "Barbados", ""},
		Bangladesh:                  {"Bangladesh", "", "Bangladesh", ""},
		Belgium:                     {"Belgium", "", "Belgia", ""},
		BurkinaFaso:                 {"Burkina Faso", "", "Burkina Faso", ""},
		Bulgaria:                    {"Bulgaria", "", "Bulgaria", ""},
		Bahrain:                     {"Bahrain", "", "Bahrain", ""},
		Burundi:                     {"Burundi", "", "Burundi", ""},
		Benin:                       {"Benin", "", "Benin", ""},
		SaintBarthélemy:             {"Saint Barthélemy", "", "Saint-Barthélemy", "Også kjent som\u00a0Saint Barts"},
		Bermuda:                     {"Bermuda", "", "Bermuda", ""},
		BruneiDarussalam:            {"Brunei Darussalam", "", "Brunei", ""},
		BoliviaPlurinationalStateOf: {"Bolivia, Plurinational State of", "", "Bolivia", ""},
		BonaireSintEustatiusAndSaba: {"Bonaire, Sint Eustatius and Saba", "", "Bonaire, Sint Eustatius og Saba", "Også kjent som Karibisk Nederland"},
		Brazil:                       {"Brazil", "", "Brasil", ""},
		Bahamas:                      {"Bahamas", "", "Bahamas", ""},
		Bhutan:                       {"Bhutan", "", "Bhutan", ""},
		BouvetIsland:                 {"Bouvet Island", "", "Bouvetøya", ""},
		Botswana:                     {"Botswana", "", "Botswana", ""},
		Belarus:                      {"Belarus", "", "Hviterussland", ""},
		Belize:                       {"Belize", "", "Belize", ""},
		Canada:                       {"Canada", "", "Canada", ""},
		CocosKeelingIslands:          {"Cocos (Keeling) Islands", "", "Kokosøyene", ""},
		CongoDemocraticRepublicOfThe: {"Congo, Democratic Republic of the", "", "Kongo, Den demokratiske rebublikken", ""},
		CentralAfricanRepublic:       {"Central African Republic", "", "Den sentralafrikanske republikk", ""},
		Congo:               {"Congo", "", "Kongo", ""},
		Switzerland:         {"Switzerland", "", "Sveits", ""},
		CoteDIvoire:         {"Cote d’Ivoire", "", "Elfenbenskysten", ""},
		CookIslands:         {"Cook Islands", "", "Cookøyene", ""},
		Chile:               {"Chile", "", "Chile", ""},
		Cameroon:            {"Cameroon", "", "Kamerun", ""},
		China:               {"China", "", "Kina", ""},
		Colombia:            {"Colombia", "", "Colombia", ""},
		CostaRica:           {"Costa Rica", "", "Costa Rica", ""},
		SerbiaAndMontenegro: {"Serbia and Montenegro", "DEPRECATED, replaced by ME – Montenegro and RS – Serbia.", "Serbia og Montenegro", "UTGÅTT - bruk RS for Serbia og ME for Montenegro"},
		Cuba:                {"Cuba", "", "Cuba", ""},
		CapeVerde:           {"Cape Verde", "", "Kapp Verde", ""},
		Curaçao:             {"Curaçao", "", "Curaçao", ""},
		ChristmasIsland:     {"Christmas Island", "", "Christmasøya", ""},
		Cyprus:              {"Cyprus", "", "Kypros", ""},
		CzechRepublic:       {"Czech Republic", "", "Tsjekkia", ""},
		Germany:             {"Germany", "", "Tyskland", ""},
		Djibouti:            {"Djibouti", "", "Djibouti", ""},
		Denmark:             {"Denmark", "", "Danmark", ""},
		Dominica:            {"Dominica", "", "Dominica", ""},
		DominicanRepublic:   {"Dominican Republic", "", "Den dominikanske republikk", ""},
		Algeria:             {"Algeria", "", "Algerie", ""},
		Ecuador:             {"Ecuador", "", "Ecuador", ""},
		Estonia:             {"Estonia", "", "Estland", ""},
		Egypt:               {"Egypt", "", "Egypt", ""},
		WesternSahara:       {"Western Sahara", "", "Vest-Sahara", ""},
		Eritrea:             {"Eritrea", "", "Eritrea", ""},
		Spain:               {"Spain", "", "Spania", ""},
		Ethiopia:            {"Ethiopia", "", "Etiopia", ""},
		Finland:             {"Finland", "", "Finland", ""},
		Fiji:                {"Fiji", "", "Fiji", ""},
		FalklandIslandsMalvinas:     {"Falkland Islands (Malvinas)", "", "Falklandsøyene", ""},
		MicronesiaFederatedStatesOf: {"Micronesia, Federated States of", "", "Mikronesiaføderasjonen", ""},
		FaroeIslands:                {"Faroe Islands", "", "Færøyene", ""},
		France:                      {"France", "", "Frankrike", ""},
		Gabon:                       {"Gabon", "", "Gabon", ""},
		UnitedKingdom:               {"United Kingdom", "", "Storbritannia", ""},
		Grenada:                     {"Grenada", "", "Grenada", ""},
		Georgia:                     {"Georgia", "", "Georgia", ""},
		FrenchGuiana:                {"French Guiana", "", "Fransk Guyana", ""},
		Guernsey:                    {"Guernsey", "", "Guernsey", ""},
		Ghana:                       {"Ghana", "", "Ghana", ""},
		Gibraltar:                   {"Gibraltar", "", "Gibraltar", ""},
		Greenland:                   {"Greenland", "", "Grønland", ""},
		Gambia:                      {"Gambia", "", "Gambia", ""},
		Guinea:                      {"Guinea", "", "Guinea", ""},
		Guadeloupe:                  {"Guadeloupe", "", "Guadeloupe", ""},
		EquatorialGuinea:            {"Equatorial Guinea", "", "Ekvatorial-Guinea", ""},
		Greece:                      {"Greece", "", "Hellas", ""},
		SouthGeorgiaAndTheSouthSandwichIslands: {"South Georgia and the South Sandwich Islands", "", "Sør-Georgia og Sør-Sandwichøyene", ""},
		Guatemala:                     {"Guatemala", "", "Guatemala", ""},
		Guam:                          {"Guam", "", "Guam", ""},
		GuineaBissau:                  {"Guinea-Bissau", "", "Guinea-Bissau", ""},
		Guyana:                        {"Guyana", "", "Guyana", ""},
		HongKong:                      {"Hong Kong", "", "Hongkong", ""},
		HeardIslandAndMcDonaldIslands: {"Heard Island and McDonald Islands", "", "Heard- og MacDonaldøyene", ""},
		Honduras:                      {"Honduras", "", "Honduras", ""},
		Croatia:                       {"Croatia", "", "Kroatia", ""},
		Haiti:                         {"Haiti", "", "Haiti", ""},
		Hungary:                       {"Hungary", "", "Ungarn", ""},
		Indonesia:                     {"Indonesia", "", "Indonesia", ""},
		Ireland:                       {"Ireland", "", "Irland", ""},
		Israel:                        {"Israel", "", "Israel", ""},
		IsleOfMan:                     {"Isle of Man", "", "Isle of Man", ""},
		India:                         {"India", "", "India", ""},
		BritishIndianOceanTerritory: {"British Indian Ocean Territory", "", "Chagosøyene", ""},
		Iraq: {"Iraq", "", "Irak", ""},
		IranIslamicRepublicOf:            {"Iran, Islamic Republic of", "", "Iran", ""},
		Iceland:                          {"Iceland", "", "Island", ""},
		Italy:                            {"Italy", "", "Italia", ""},
		Jersey:                           {"Jersey", "", "Jersey", ""},
		Jamaica:                          {"Jamaica", "", "Jamaica", ""},
		Jordan:                           {"Jordan", "", "Jordan", ""},
		Japan:                            {"Japan", "", "Japan", ""},
		Kenya:                            {"Kenya", "", "Kenya", ""},
		Kyrgyzstan:                       {"Kyrgyzstan", "", "Kirgisistan", ""},
		Cambodia:                         {"Cambodia", "", "Kambodsja", ""},
		Kiribati:                         {"Kiribati", "", "Kiribati", ""},
		Comoros:                          {"Comoros", "", "Komorene", ""},
		SaintKittsAndNevis:               {"Saint Kitts and Nevis", "", "Sankt Kitts-Nevis", ""},
		KoreaDemocraticPeoplesRepublicOf: {"Korea, Democratic People’s Republic of", "", "Nord-Korea", "Den demokratiske folkerepublikken Korea"},
		KoreaRepublicOf:                  {"Korea, Republic of", "", "Sør-Korea", "Republikken Korea"},
		Kuwait:                           {"Kuwait", "", "Kuwait", ""},
		CaymanIslands:                    {"Cayman Islands", "", "Caymanøyene", ""},
		Kazakhstan:                       {"Kazakhstan", "", "Kasakhstan", ""},
		LaoPeoplesDemocraticRepublic:     {"Lao People’s Democratic Republic", "", "Laos", ""},
		Lebanon:                              {"Lebanon", "", "Libanon", ""},
		SaintLucia:                           {"Saint Lucia", "", "Saint Lucia", ""},
		Liechtenstein:                        {"Liechtenstein", "", "Liechtenstein", ""},
		SriLanka:                             {"Sri Lanka", "", "Sri Lanka", ""},
		Liberia:                              {"Liberia", "", "Liberia", ""},
		Lesotho:                              {"Lesotho", "", "Lesotho", ""},
		Lithuania:                            {"Lithuania", "", "Litauen", ""},
		Luxembourg:                           {"Luxembourg", "", "Luxembourg", ""},
		Latvia:                               {"Latvia", "", "Latvia", ""},
		Libya:                                {"Libya", "", "Libya", ""},
		Morocco:                              {"Morocco", "", "Marokko", ""},
		Monaco:                               {"Monaco", "", "Monaco", ""},
		MoldovaRepubicOf:                     {"Moldova, Repubic of", "", "Moldova", ""},
		Montenegro:                           {"Montenegro", "", "Montenegro", ""},
		SaintMartinFrenchPart:                {"Saint Martin (French part)", "", "Saint Martin (den franske delen)", ""},
		Madagascar:                           {"Madagascar", "", "Madagaskar", ""},
		MarshallIslands:                      {"Marshall Islands", "", "Marshalløyene", ""},
		MacedoniaTheFormerYugoslavRepublicOf: {"Macedonia, the former Yugoslav Republic of", "", "Makedonia (FYR)", ""},
		Mali:     {"Mali", "", "Mali", ""},
		Myanmar:  {"Myanmar", "", "Myanmar", ""},
		Mongolia: {"Mongolia", "", "Mongolia", ""},
		Macao:    {"Macao", "", "Macau", ""},
		NorthernMarianaIslands: {"Northern Mariana Islands", "", "Marianene", ""},
		Martinique:             {"Martinique", "", "Martinique", ""},
		Mauritania:             {"Mauritania", "", "Mauritania", ""},
		Montserrat:             {"Montserrat", "", "Montserrat", ""},
		Malta:                  {"Malta", "", "Malta", ""},
		Mauritius:              {"Mauritius", "", "Mauritius", ""},
		Maldives:               {"Maldives", "", "Maldivene", ""},
		Malawi:                 {"Malawi", "", "Malawi", ""},
		Mexico:                 {"Mexico", "", "Mexico", ""},
		Malaysia:               {"Malaysia", "", "Malaysia", ""},
		Mozambique:             {"Mozambique", "", "Mosambik", ""},
		Namibia:                {"Namibia", "", "Namibia", ""},
		NewCaledonia:           {"New Caledonia", "", "Ny-Caledonia", ""},
		Niger:                  {"Niger", "", "Niger", ""},
		NorfolkIsland:          {"Norfolk Island", "", "Norfolkøya", ""},
		Nigeria:                {"Nigeria", "", "Nigeria", ""},
		Nicaragua:              {"Nicaragua", "", "Nicaragua", ""},
		Netherlands:            {"Netherlands", "", "Nederland", ""},
		Norway:                 {"Norway", "", "Norge", ""},
		Nepal:                  {"Nepal", "", "Nepal", ""},
		Nauru:                  {"Nauru", "", "Nauru", ""},
		Niue:                   {"Niue", "", "Niue", ""},
		NewZealand:             {"New Zealand", "", "New Zealand", ""},
		Oman:                   {"Oman", "", "Oman", ""},
		Panama:                 {"Panama", "", "Panama", ""},
		Peru:                   {"Peru", "", "Peru", ""},
		FrenchPolynesia:        {"French Polynesia", "", "Fransk Polynesia", ""},
		PapuaNewGuinea:         {"Papua New Guinea", "", "Papua Ny-Guinea", ""},
		Philippines:            {"Philippines", "", "Filippinene", ""},
		Pakistan:               {"Pakistan", "", "Pakistan", ""},
		Poland:                 {"Poland", "", "Polen", ""},
		SaintPierreAndMiquelon: {"Saint Pierre and Miquelon", "", "Sankt Pierre og Miquelon", ""},
		Pitcairn:               {"Pitcairn", "", "Pitcairnøya", ""},
		PuertoRico:             {"Puerto Rico", "", "Puerto Rico", ""},
		PalestineStateOf:       {"Palestine, State of", "", "Palestina", ""},
		Portugal:               {"Portugal", "", "Portugal", ""},
		Palau:                  {"Palau", "", "Palau", ""},
		Paraguay:               {"Paraguay", "", "Paraguay", ""},
		Qatar:                  {"Qatar", "", "Qatar", ""},
		Réunion:                {"Réunion", "", "Réunion", ""},
		Romania:                {"Romania", "", "Romania", ""},
		Serbia:                 {"Serbia", "", "Serbia", ""},
		RussianFederation:      {"Russian Federation", "", "Russland", ""},
		Rwanda:                 {"Rwanda", "", "Rwanda", ""},
		SaudiArabia:            {"Saudi Arabia", "", "Saudi-Arabia", ""},
		SolomonIslands:         {"Solomon Islands", "", "Salomonøyene", ""},
		Seychelles:             {"Seychelles", "", "Seychellene", ""},
		Sudan:                  {"Sudan", "", "Sudan", ""},
		Sweden:                 {"Sweden", "", "Sverige", ""},
		Singapore:              {"Singapore", "", "Singapore", ""},
		SaintHelenaAscensionAndTristanDaCunha: {"Saint Helena, Ascension and Tristan da Cunha", "", "St. Helena, Ascension og Tristan da Cunha", ""},
		Slovenia:              {"Slovenia", "", "Slovenia", ""},
		SvalbardAndJanMayen:   {"Svalbard and Jan Mayen", "", "Svalbard og Jan Mayen", ""},
		Slovakia:              {"Slovakia", "", "Slovakia", ""},
		SierraLeone:           {"Sierra Leone", "", "Sierra Leone", ""},
		SanMarino:             {"San Marino", "", "San Marino", ""},
		Senegal:               {"Senegal", "", "Senegal", ""},
		Somalia:               {"Somalia", "", "Somalia", ""},
		Suriname:              {"Suriname", "", "Surinam", ""},
		SouthSudan:            {"South Sudan", "", "Sør-Sudan", ""},
		SaoTomeAndPrincipe:    {"Sao Tome and Principe", "", "São Tomé og Príncipe", ""},
		ElSalvador:            {"El Salvador", "", "El Salvador", ""},
		SintMaartenDutchPart:  {"Sint Maarten (Dutch part)", "", "Sint Maarten (den nederlandske delen av Saint Martin)", ""},
		SyrianArabRepublic:    {"Syrian Arab Republic", "", "Syria", ""},
		Swaziland:             {"Swaziland", "", "Swaziland", ""},
		TurksAndCaicosIslands: {"Turks and Caicos Islands", "", "Turks- og Caicosøyene", ""},
		Chad: {"Chad", "", "Tsjad", ""},
		FrenchSouthernTerritories: {"French Southern Territories", "", "De franske sørterritorier", ""},
		Togo:                     {"Togo", "", "Togo", ""},
		Thailand:                 {"Thailand", "", "Thailand", ""},
		Tajikistan:               {"Tajikistan", "", "Tadsjikistan", ""},
		Tokelau:                  {"Tokelau", "", "Tokelau", ""},
		TimorLeste:               {"Timor-Leste", "", "Øst-Timor", ""},
		Turkmenistan:             {"Turkmenistan", "", "Turkmenistan", ""},
		Tunisia:                  {"Tunisia", "", "Tunisia", ""},
		Tonga:                    {"Tonga", "", "Tonga", ""},
		Turkey:                   {"Turkey", "", "Tyrkia", ""},
		TrinidadAndTobago:        {"Trinidad and Tobago", "", "Trinidad og Tobago", ""},
		Tuvalu:                   {"Tuvalu", "", "Tuvalu", ""},
		TaiwanProvinceOfChina:    {"Taiwan, Province of China", "", "Taiwan", ""},
		TanzaniaUnitedRepublicOf: {"Tanzania, United Republic of", "", "Tanzania", ""},
		Ukraine:                  {"Ukraine", "", "Ukraina", ""},
		Uganda:                   {"Uganda", "", "Uganda", ""},
		UnitedStatesMinorOutlyingIslands: {"United States Minor Outlying Islands", "", "USAs ytre småøyer", ""},
		UnitedStates:                     {"United States", "", "USA", ""},
		Uruguay:                          {"Uruguay", "", "Uruguay", ""},
		Uzbekistan:                       {"Uzbekistan", "", "Usbekistan", ""},
		HolySeeVaticanCityState:          {"Holy See (Vatican City State)", "", "Vatikanstaten", ""},
		SaintVincentAndTheGrenadines:     {"Saint Vincent and the Grenadines", "", "Saint Vincent og Grenadinene", ""},
		VenezuelaBolivarianRepublicOf:    {"Venezuela, Bolivarian Republic of", "", "Venezuela", ""},
		VirginIslandsBritish:             {"Virgin Islands, British", "", "Jomfruøyene (GB)\n", ""},
		VirginIslandsUS:                  {"Virgin Islands, US", "", "Jomfruøyene (US)", ""},
		VietNam:                          {"Viet Nam", "", "Vietnam", ""},
		Vanuatu:                          {"Vanuatu", "", "Vanuatu", ""},
		WallisAndFutuna:                  {"Wallis and Futuna", "", "Wallis- og Futunaøyene", ""},
		Samoa:                            {"Samoa", "", "Samoa", ""},
		Yemen:                            {"Yemen", "", "Jemen", ""},
		Mayotte:                          {"Mayotte", "", "Mayotte", ""},
		Yugoslavia:                       {"Yugoslavia", "DEPRECATED, replaced by ME – Montenegro and RS – Serbia.", "Jugoslavia", "UTGÅTT - bruk RS for Serbia og ME for Montenegro"},
		SouthAfrica:                      {"South Africa", "", "Sør-Afrika", ""},
		Zambia:                           {"Zambia", "", "Zambia", ""},
		Zimbabwe:                         {"Zimbabwe", "", "Zimbabwe", ""},
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
