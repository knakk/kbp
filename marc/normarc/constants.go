package normarc

const (
	Uspesifisert = ' '

	// Postens status
	StatusRetted     = 'c'
	StatusSletted    = 'd'
	StatusNy         = 'n'
	StatusOppgradert = 'p'

	// Materialtyper
	MaterialtypeTekstlig         = 'a'
	MaterialtypeManuskript       = 'b'
	MaterialtypeMusikktrykk      = 'c'
	MaterialtypeMusikkmanuskript = 'd'
	MaterialtypeKartografisk     = 'e'
	MaterialtypeKartmanuskript   = 'f'
	MaterialtypeFilm             = 'g'
	MaterialtypeLydopptak        = 'i'
	MaterialtypeLydopptakMusikk  = 'j'
	MaterialtypeGrafisk          = 'k'
	MaterialtypeFiler            = 'm'
	MaterialtypeKombidokument    = 'o'
	Materialtype3DGjenstand      = 'r'

	// Bibliografisk kategori
	KategoriIkkePeriodiskAnalytt = 'a'
	KategoriSerieanalytt         = 'b'
	KategoriSamling              = 'c'
	KategoriMonografi            = 'm'
	KategoriPeriodiskAnalytt     = 'p'
	KategoriPeriodikum           = 's'

	// Beskrivelsesnivå
	NivåFullstendig = '0'
	Nivå1           = '1'
	Nivå2           = '2'
	NivåForeløbig   = '5'
)

// Posisjoner i kontrollfelt 008
const (
	PosRegDato                 = 0
	PosDateringskode           = 6
	PosUtgivelsesstatus        = 6
	PosÅrstall1                = 7
	PosÅrstall2                = 11
	PosUtgivelsesland          = 15
	PosPeriodisitet            = 18
	PostSpilletid              = 18
	PosIllustrasjoner          = 19
	PosRelieff                 = 18
	PosRegularitet             = 19
	PosISSNsenter              = 20
	PosFormat                  = 20
	PosPeriodikumType          = 21
	PosMålgruppe               = 22
	PosProjeksjon              = 22
	PosMaterialform            = 23
	PosMedfølgendeMateriale    = 23
	PosInnholdskoder           = 24
	PosLedsagendeMateriale     = 24
	PosMaterialtype            = 25
	PosDataType                = 26
	PosOffentligPublikasjon    = 28
	PosKonferanse              = 29
	PosFestskriftkode          = 30
	PosLittærTekstILydopptak   = 30
	PosRegister                = 31
	PosLitterærForm            = 33
	PosTittelensOriginalSkrift = 33
	PosSpesielleFormatDetaljer = 33
	PosTeknikk                 = 34
	PosBiografi                = 34
	PosSpråk                   = 35
	PosKatalogiseringsKilde    = 39
)

// Landkoder
// Hentet fra http://nabo.nb.no/trip?_b=landekoder&_n=0&_q=500&ccl=&_O=kode
const (
	LandAndorra                           = "ad" // Fyrstedømmet Andorra
	LandDeForenteArabiskeEmirater         = "ae" // Uavhengig siden 1971
	LandAfghanistan                       = "af" // Den islamske republikken Afghanistan
	LandAntiguaOgBarbuda                  = "ag" // Uavhengig siden 1981
	LandAnguilla                          = "ai" // Indre selvstyre under Storbritannia
	LandAlbania                           = "al" // Republikken Albania
	LandArmenia                           = "am" // Republikken Armenia. Tidligere del av Sovjetunionen. Uavhengig siden 1991.
	LandDeNederlandskeAntiller            = "an" // Siden 1954 del av kongeriket Nederlandene
	LandAngola                            = "ao" // Republikken Angola. Uavhengig siden 1975
	LandAntarktis                         = "aq" // De foreldede kodene nq, ATN og 216 for Dronning Maud Land; bq og ATB for British Antarctic Territory; fq og ATF for Fransk sektor i Antarktis anbefales ikke brukt. (bq er nå tildelt et annet land)
	LandArgentina                         = "ar" // Republikken Argentina
	LandAmerikanskSamoa                   = "as"
	LandØsterrike                         = "at" // Republikken Østerrike/Austerrike
	LandAustralia                         = "au"
	LandAruba                             = "aw" // Siden 1954 del av Kongeriket Nederlandene
	LandÅland                             = "ax" // Ålandsøyene - Landskapet Åland. Indre selvstyre siden 1920. Alternativt kan kodene for Finland brukes.
	LandAserbajdsjan                      = "az" // Republikken Aserbajdsjan. Tidligere del av Sovjetunionen. Uavhengig siden 1991.
	LandBosniaHercegovina                 = "ba" // Tidligere del av Jugoslavia. Uavhengig siden 1.3.1992
	LandBarbados                          = "bb" // Uavhengig siden 1966
	LandBangladesh                        = "bd" // Folkerepublikken Bangladesh. Uavhengig siden 1971
	LandBelgia                            = "be" // Kongeriket Belgia
	LandBurkinaFaso                       = "BF" // 1898-1963: Øvre Volta. Uavhengig siden 1960.
	LandBulgaria                          = "bg" // Republikken Bulgaria
	LandBahrain                           = "bh" // Køngedømmet Bahrain. Uavhengig siden 1971
	LandBurundi                           = "bi" // Republikken Burundi. Uavhengig siden 1962
	LandBenin                             = "bj" // Republikken Benin. Uavhengig siden 1960. Før 1976 ved navn: Dahomey
	LandSaintBarthélemy                   = "bl" // Øy i Karibia som administreres av det franske oversjøiske departement Guadeloupe
	LandBermuda                           = "bm"
	LandBrunei                            = "bn" // Brunei Darussalam. Uavhengig siden 1984
	LandBolivia                           = "bo" // Den flernasjonale stat Bolivia
	LandBrasil                            = "br" // Forbundsrepublikken Brasil
	LandBahamas                           = "bs" // Bahamasambandet. Uavhengig siden 1973
	LandBhutan                            = "bt" // Kongeriket Bhutan
	LandBouvetøya                         = "bv" // Ubebodd norsk øy i Sør-Atlanteren
	LandBotswana                          = "bw" // Republikken Botswana. Uavhengig siden 1966
	LandHviterussland                     = "by" // Republikken Hviterussland. Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandBelize                            = "bz" // Uavhengig siden 1981
	LandCanada                            = "ca"
	LandKokosøyene                        = "cc" // Australsk besittelse siden 1955
	LandKongo                             = "cd" // Den demokratiske republikken Kongo. Tidl. Belgisk Kongo. Uavhengig fra 1960 som Kongo-Leopoldville. 1964-1971 og etter 1997: Den demokratiske republikken Kongo. 1971-1997: Zaïre
	LandDenSentralafrikanskeRepublikk     = "cf" // Uavhengig siden 1960
	LandKongoBrazzaville                  = "cg" // Republikken Kongo. Uavhengig siden 1960
	LandSveits                            = "ch" // Det sveitsiske edsforbund
	LandElfenbenskysten                   = "ci" // Republikken Elfenbenskysten. Uavhengig siden 1960
	LandCookøyene                         = "ck" // Selvstyrt øygruppe i fritt forbund med New Zealand
	LandChile                             = "cl" // Republikken Chile
	LandKamerun                           = "cm" // Republikken Kamerun. Uavhengig siden 1960
	LandKina                              = "cn" // Folkerepublikken Kina
	LandColombia                          = "co" // Republikken Colombia
	LandCostaRica                         = "cr" // Republikken Costa Rica
	LandCuba                              = "cu" // Republikken Cuba
	LandKappVerde                         = "cv" // Republikken Kapp Verde. Tidligere portugisisk besittelse. Uavhengig siden 1975
	LandChristmasøya                      = "cx" // Territoriet Christmasøya. Administreres av Australia
	LandKypros                            = "cy" // Republikken Kypros. Uavhengig siden 1960
	LandTsjekkia                          = "cz" // Den Tsjekkiske republikk. Selvstendig stat etter delingen av Tsjekkoslovakia i 1993
	LandTyskland                          = "de" // Forbundsrepublikken Tyskland. for Øst-Tyskland - Den Tyske Demokratiske Republikken (DDR), som ble innlemmet i Vest-Tyskland til Forbundsrepublikken Tyskland i 1990.
	LandDjibouti                          = "dj" // Republikken Djibouti. Før 1977 kjent som: French Afars and Issas. Uavhengig 1977. Den foreldede koden ai kan ikke brukes for tidligere år, da den nå er tildelt en annen stat.
	LandDanmark                           = "dk" // Kongeriket Danmark
	LandDominica                          = "dm" // Samveldet Dominica. Må ikke forveksles med Den dominikanske republikken
	LandDenDominikanskeRepublikk          = "do" // Må ikke forveksles med Samveldet Dominica
	LandAlgerie                           = "dz" // Den demokratiske folkerepublikken Algerie
	LandEcuador                           = "ec" // Republikken Ecuador
	LandEstland                           = "ee" // Republikken Estland. Tidligere del av Sovjetunionen. Uavhengig siden 1991.
	LandEgypt                             = "eg" // Den arabiske republikk Egypt
	LandVestSahara                        = "eh" // Tidligere Spansk Sahara. Nå land uten selvstendig status
	LandEritrea                           = "ER" // Staten Eritrea. Uavhengig siden 1993
	LandSpania                            = "es" // Kongeriket Spania. Inkluderer Kanariøyene
	LandEtiopia                           = "et" // Den føderale demokratiske republikk Etiopia
	LandFinland                           = "fi" // Republikken Finland
	LandFiji                              = "fj" // Republikken Fijiøyene. Uavhengig siden 1970
	LandFalklandsøyene                    = "fk" // Britisk selvstyrt oversjøisk territorium
	LandMikronesia                        = "fm" // Mikronesiaføderasjonen. Tidligere del av Pacific Islands Trust Territory. Uavhengig øynasjon siden 1986 består av over 600 øyer.
	LandFærøyene                          = "fo" // Indre selvstyre under Danmark
	LandFrankrike                         = "fr" // Republikken Frankrike. Kodene inkluderer Korsika. De foreldede kodene fx, FXX og 249 for Frankrike i Europa ("Metropolitan-Frankrike"), anbefales ikke brukt.
	LandGabon                             = "ga" // Republikken Gabon
	LandStorbritannia                     = "gb" // Det forente kongeriket Storbritannia og Nord-Irland
	LandGrenada                           = "gd" // Uavhengig siden 1974
	LandGeorgia                           = "ge" // Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandFranskGuyana                      = "gf" // Fransk oversjøisk departement (fylke)
	LandGuernsey                          = "gg" // Ikke del av United Kingdom, men en særskilt besittelse under Den Britiske Kronen. Omfatter også Kanaløyene Alderney, Brecqhou, Burhou, Herm, Jethou, Lihou, Sark med flere småøyer.
	LandGhana                             = "gh" // Republikken Ghana. Uavhengig siden 1957
	LandGibraltar                         = "gi"
	LandGrønland                          = "gl" // Dansk selvstyrt besittelse siden 1953
	LandGambia                            = "gm" // Republikken Gambia
	LandGuinea                            = "gn" // Republikken Guinea. Tidligere: Fransk Guinea. Uavhengig siden 1958
	LandGuadeloupe                        = "gp" // Oversjøisk fransk departement (fylke) i Karibia. (NB! Må ikke forveksles med Guadalupe - mange steder i den spansk-talende verden)
	LandEkvatorialGuinea                  = "gq" // Republikken Ekvatorial-Guinea. Uavhengig siden 1968
	LandHellas                            = "gr" // Republikken Hellas
	LandSørGeorgiaOgSørSandwichØyene      = "gs" // Britisk oversjøisk territorium i Sør-Atlanteren. Kun bebodd av forskere
	LandGuatemala                         = "gt" // Republikken Guatemala
	LandGuam                              = "gu" // Guamterritoriet. Besittelse i Stillehavet under USA siden 1944
	LandGuineaBissau                      = "gw" // Republikken Guinea-Bissau. Tidligere: Portugisisk Guinea. Uavhengig siden 1973
	LandGuyana                            = "gy" // Den kooperative republikken Guyana. Tidligere: Britisk Guyana. Uavhengig siden 1966
	LandHongKong                          = "hk" // Den spesielle administrative region Hong Kong. Tidligere Britisk besittelse. Overdratt til Kina 1997
	LandHeardOgMcDonaldøyene              = "hm" // Ubebodd australsk øygruppe
	LandHonduras                          = "hn" // Republikken Honduras
	LandKroatia                           = "hr" // Republikken Kroatia. Tidligere del av Jugoslavia. Uavhengig siden 1991
	LandHaiti                             = "ht" // Republikken Haiti
	LandUngarn                            = "hu" // Republikken Ungarn
	LandIndonesia                         = "id" // Republikken Indonesia. Uavhengig siden 1945
	LandIrland                            = "ie"
	LandIsrael                            = "il" // Staten Israel. Uavhengig siden 1948
	LandMan                               = "im" // Øya er ikke en del av United Kingdom, men en selvstyrt besittelse under Den Britiske Kronen.
	LandIndia                             = "in" // Republikken India. Tidligere Britisk besittelse. Uavhengig siden 1947.
	LandDetBritiskeTerritorietIIndiahavet = "io" // Britisk oversjøisk territorium i Indiahavet. Består av Chagos Archipelago - Chagosøyene, men bebos kun av militært personell.
	LandIrak                              = "iq" // Republikken Irak
	LandIran                              = "ir" // Den islamske republikken Iran
	LandIsland                            = "is" // Republikken Island
	LandItalia                            = "it" // Republikken Italia
	LandJersey                            = "je" // Ikke del av United Kingdom, men en særskilt besittelse under Den Britiske Kronen. Omfatter også de nesten ubebodde Kanaløyene Écréhous, Miniquiers og Pierres de Lecq og andre klipper og rev.
	LandJamaica                           = "jm" // Uavhengig siden 1962
	LandJordan                            = "jo" // Det hasjimittiske kongerike Jordan. Uavhengig siden 1946
	LandJapan                             = "jp"
	LandKenya                             = "ke" // Republikken Kenya
	LandKirgisistan                       = "kg" // Republikken Kirgisistan. Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandKambodsja                         = "kh" // Kongeriket Kambodsja. Uavhengig siden 1953. Tidligere: Kampuchea
	LandKiribati                          = "ki" // Republikken Kiribati. Før 1974 kjent som Gilbert Islands i øygruppen Gilbert and Ellis Islands. Uavhengig siden 1979 sammen med Canton- og Enderburyøyene (ct, CTE og 128) De foreldede kodene ge, GEL og 296 kan ikke brukes for tidligere år, da en av dem nå er tildelt en annen stat
	LandKomorene                          = "km" // Unionen Komorene. Består av øyene Grande Comore, Mohéli og Anjouan. Uavhengig siden 1975
	LandSaintKittsOgNevis                 = "kn" // Saint Kitts og Nevisføderasjonen. Også kjent som Sankt Christopher og Nevis
	LandNordKorea                         = "kp" // Den demokratiske folkerepublikken Korea. Uavhengig siden 1948
	LandSørKorea                          = "kr" // Republikken Korea. Uavhengig siden 1948
	LandKuwait                            = "kw" // Staten Kuwait. Uavhengig siden 1961
	LandCaymanøyene                       = "ky"
	LandKasakhstan                        = "kz" // Republikken Kazakhstan. Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandLaos                              = "la" // Den demokratiske folkerepublikken Laos. Uavhengig siden 1954
	LandLibanon                           = "lb" // Republikken Libanon
	LandSaintLucia                        = "lc" // Uavhengig siden 1979
	LandLiechtenstein                     = "li" // Fyrstedømmet Liechtenstein. Uavhengig siden 1866
	LandSriLanka                          = "lk" // Den demokratiske sosialistiske republikken Sri Lanka. Tidligere Ceylon. Uavhengig siden 1948
	LandLiberia                           = "lr" // Republikken Liberia
	LandLesotho                           = "ls" // Kongeriket Lesotho. Tidligere Basutoland. Uavhengig fra 1966
	LandLitauen                           = "lt" // Republikken Litauen. Tidligere del av Sovjetunionen. Uavhengig siden 1990
	LandLuxemburg                         = "lu" // Storhertugdømmet Luxemburg. Uavhengig siden 1815
	LandLatvia                            = "lv" // Republikken Latvia. Tidligere del av Sovjetunionen, også kjent som Lettland. Uavhengig siden 1991.
	LandLibya                             = "ly" // Den libyske arabiske jamahiriya- Den sosialistiske arabiske folkerepublikken Libya. (Jamahiriya: Massenes stat, dvs. styre ved (folke)massene)
	LandMarokko                           = "ma" // Kongeriket Marokko. Uavhengig siden 1956
	LandMonaco                            = "mc" // Fyrstedømmet Monaco
	LandMoldova                           = "md" // Republikken Moldova. Tidligere del av Sovjetunionen, også kjent som Moldavia. Uavhengig siden 1991
	LandMontenegro                        = "me" // Republikken Montenegro. Uavhengig fra 2006, men kodene kan også brukes fra 2004 mens landet var i union som Serbia og Montenegro
	LandSaintMartin                       = "mf" // Todelt øy i Karibia. Koden gjelder bare norddelen som administreres av det franske oversjøiske departement Guadeloupe. Sørdelen med navnet Sint Maarten administreres sammen med De Nederlandske Antiller (an).
	LandMadagaskar                        = "mg" // Republikken Madagaskar. Uavhengig siden 1960
	LandMarshalløyene                     = "mh" // Republikken Marshalløyene. Uavhengig siden 1960
	LandMakedonia                         = "MK" // Republikken Makedonia. Også kjent som Den tidligere jugoslaviske republikken Makedonia. Kortformen er ikke anerkjent av FN. Uavhengighetsprosess siden 1993.
	LandMali                              = "ml" // Republikken Mali. Uavhengig siden 1960
	LandMyanmar                           = "mm" // Unionen Myanmar. Før 1989: Burma med hovedstad Rangoon - Burmeser/burmesisk - Burmese.
	LandMongolia                          = "mn" // Tidligere: Folkerepublikken Mongolia
	LandMacao                             = "mo" // Spesiell administrativ region i Kina. Overført fra portugisisk til kinesisk administrasjon i 1999
	LandNordMarianene                     = "mp" // Samveldet Nord-Marianene. Samvelde i politisk union med USA.De største øyene: Saipan, Tinian og Rota
	LandMartinique                        = "mq" // Oversjøisk fransk departement (fylke)
	LandMauritania                        = "mr" // Den islamiske republikken Mauritania. Uavhengig siden 1960
	LandMontserrat                        = "ms" // Britisk oversjøisk territorium
	LandMalta                             = "mt" // Republikken Malta. Uavhengig siden 1964
	LandMauritius                         = "mu" // Republikken Mauritius. Uavhengig siden 1968
	LandMaldivene                         = "mv" // Repblikken Maldivene. Uavhengig siden 1965
	LandMalawi                            = "mw" // Republikken Malawi. Tidligere: Nyasaland. Uavhengig siden 1964
	LandMexico                            = "mx" // De forente stater Mexico
	LandMalaysia                          = "my"
	LandMosambik                          = "mz" // Republikken Mosambik
	LandNamibia                           = "na" // Republikken Namibia. Uavhengig siden 1990
	LandNyCaledonia                       = "nc" // Fransk oversjøisk territorium
	LandNiger                             = "ne" // Republikken Niger. Uavhengig siden 1960
	LandNorfolkøya                        = "nf" // Selvstyrt territorium i union med Australia
	LandNigeria                           = "ng" // Forbundsrepublikken Nigeria. Uavhengig siden 1960
	LandNicaragua                         = "ni" // Republikke Nicaragua
	LandNederland                         = "nl" // Kongeriket Nederland. Tidligere: Holland. Siden 1954 del av Kongeriket Nederlandene sammen med De nederlandske Antillene, Aruba (og Surinam til 1975)
	LandNorge                             = "no" // Kongeriket Norge/Noreg
	LandNepal                             = "np" // Den føderale republikken Nepal. Ved grunnlovesendring erklært å være republikk den 28.12.2007
	LandNauru                             = "nr" // Republikken Nauru. Uavhengig siden 1968
	LandNiue                              = "nu" // Selvstyre i fri union med New Zealand siden 1974
	LandNewZealand                        = "nz"
	LandOman                              = "om" // Sultanatet Oman
	LandPanama                            = "pa" // Republikken Panama. For årene 1903-1973 eksisterte Panama Kanalsone under overhøyhet av USA. Kodene pz, PCZ og 590 kan om ønskelig brukes for Kanalsonen i disse årene.
	LandPeru                              = "pe" // Republikken Peru
	LandFranskPolynesia                   = "pf" // Delvis uavhengig av Frankrike som oversjøisk territorium siden 1977
	LandPapuaNyGuinea                     = "pg" // Selvstyre under Australia fra 1973. Uavhengig fra 1975.
	LandFilippinene                       = "ph" // Republikken Filippinene
	LandPakistan                          = "pk" // Den islamske republikken Pakistan. Uavhengig siden 1947
	LandPolen                             = "pl" // Republikken Polen
	LandSaintPierreOgMiquelon             = "pm" // Fransk oversjøisk øygruppe rett sør for Newfoundland
	LandPitcairnøya                       = "pn" // egentlig: Pitcairnøyene. Britisk oversjøisk territorium. Bare en av fire øyer er bebodd.
	LandPuertoRico                        = "pr" // Samveldet Puerto Rico. Selvstyrt territorium under De Forente Stater
	LandPalestina                         = "ps" // De palestinske territoriene. Omfatter Vestbredden og Gazastripen
	LandPortugal                          = "pt" // Republikken Portugal. Inkluderer Madeira og Asorene
	LandPalau                             = "pw" // Republikken Palau. Tidligere del av Pacific Islands Trust Territory. Uavhengig siden 1994.
	LandParaguay                          = "py" // Republikken Paraguay
	LandQatar                             = "qa" // Staten Qatar. Uavhengig siden 1971
	LandRéunion                           = "re" // Oversjøisk fransk departement (fylke) siden 1974
	LandRomania                           = "ro"
	LandSerbia                            = "rs" // Republikken Serbia. Tidligere del av Jugoslavia, som opphørte som navn i 2003. - Uavhengig fra 2006
	LandRussland                          = "ru" // Den russiske føderasjon. Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandRwanda                            = "rw" // Republikken Rwanda. Uavhengig siden 1962
	LandSaudiArabia                       = "sa" // Kongeriket Saudi Arabia
	LandSalomonøyene                      = "sb" // Selvstyrt stat innenfor the British Commonwealth siden 1978
	LandSeychellene                       = "sc" // Republikken Seychellene
	LandSudan                             = "sd" // Republikken Sudan
	LandSverige                           = "se" // Kongeriket Sverige
	LandSingapore                         = "sg" // Republikken Singapore
	LandSanktHelena                       = "sh" // Øya inngår i British Overseas Territory in the South Atlantic Ocean, men landekoden omfatter også Ascension Island og Tristan da Cunha
	LandSlovenia                          = "si" // Republikken Slovenia. Tidligere del av Yugoslavia. Uavhengig siden 1991.
	LandSvalbardOgJanMayen                = "sj" // Del av Kongeriket Norge
	LandSlovakia                          = "sk" // Den slovakiske republikken. Tidligere del av Tsjekkoslovakia. Selvstendig stat fra 1993.
	LandSierraLeone                       = "sl" // Republikken Sierra Leone. Uavhengig siden 1961
	LandSanMarino                         = "sm" // Republikken San Marino
	LandSenegal                           = "sn" // Republikken Senegal. Uavhengig siden 1960
	LandSomalia                           = "so" // Republikken Somalia. Uavhengig siden 1960
	LandSurinam                           = "sr" // Republikken Surinam. Tidligere del av Kongeriket Nederlandene. Uavhengig siden 1975.
	LandSãoToméOgPríncipe                 = "st" // Den demokratiske republikken São Tomé og Príncipe. Uavhengig siden 1975
	LandElSalvador                        = "sv" // Republikken El Salvador
	LandSyria                             = "sy" // Den arabiske republikken Syria
	LandSwaziland                         = "sz" // Kongeriket Swaziland. Uavhengig siden 1968
	LandTurksOgCaicosøyene                = "tc" // Britisk oversjøisk territorium i Karibia
	LandTsjad                             = "td" // Republikken Tsjad. Uavhengig siden 1960
	LandFranskeSørterritorier             = "tf" // De franske sørterritorier. Omfatter de ubebodde øyene: Amsterdam, Bassas de India, Glorieusesøyene, Crozetøyene, Europaøya, Juan de Nova, Kerguelen, Saint-Paul, Tromelinøya og en del andre omstridte øyer. (for Fransk sektor i Antarktis: se aq)
	LandTogo                              = "tg" // Republikken Togo. Uavhengig siden 1960
	LandThailand                          = "th" // Kongeriket Thailand. Før 1949 med navnet Siam.
	LandTadsjikistan                      = "tj" // Republikken Tadsjikistan. Tidligere del av Sovjetunionen. Uavhengig siden 1991.
	LandTokelau                           = "tk" // Tidligere: Union Islands. Territorium under New Zealandsk styre.
	LandØstTimor                          = "tl" // Den demokratiske republikken Øst-Timor. Tidligere Portugisisk Timor. Del av Indonsia fra 1975. Uavhengig fra 2002
	LandTurkmenistan                      = "tm" // Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandTunisia                           = "tn" // Republikken Tunisia. Uavhengig siden 1956
	LandTonga                             = "to" // Kongeriket Tonga
	LandTyrkia                            = "tr" // Republikken Tyrkia
	LandTrinidadOgTobago                  = "tt" // Republikken Trinidad og Tobago. Uavhengig siden 1962
	LandTuvalu                            = "tv" // Før 1974 kjent som Ellice Islands i øygruppen Gilbert and Ellice Islands. Selvstendig stat innenfor det Britiske Samveldet siden 1978 Den foreldede koden ge kan ikke brukes for tidligere år, da den nå er tildelt en annen stat.
	LandTaiwan                            = "tw" // Den kinesiske Provinsen Taiwan. Tidligere: Formosa. Både Folkerepublikken Kina og Republikken Kina hevder suverenitet over Taiwan. Republikken har hevdet uavhengighet siden 1952.
	LandTanzania                          = "tz" // Forbundsrepublikken Tanzania. Tidligere: Zanzibar og Tanganyika. Fra 1964: Tanzania
	LandUkraina                           = "ua" // Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandUganda                            = "ug" // Republikken Uganda. Uavhengig siden 1962
	LandUSAsSmåFjerntliggendeØyer         = "um" // Omfatter (koder gjelder årene før 1986): Baker Island, Bajo Nuevo Bank, Howland Island, Jarvis Island, Johnston Atoll (jt), Kingman Reef, Navassa Island, Midwayøyene (mi), Palmyra Atoll, Serranilla Bank, Wake Island (wk)
	LandUSA                               = "us" // De Forente Stater/Sambandsstatene
	LandUruguay                           = "uy" // Republikken Uruguay
	LandUsbekistan                        = "uz" // Republikken Usbekistan. Tidligere del av Sovjetunionen. Uavhengig siden 1991
	LandVatikanstaten                     = "va" // Uavhengig fra 1929
	LandSaintVincentOgGrenadinene         = "vc"
	LandVenezuela                         = "ve" // Republikken Venezuela
	LandDeBritiskeJomfruøyene             = "vg" // De britiske Jomfruøyene
	LandDeAmerikanskeJomfruøyene          = "vi" // De amerikanske Jomfruøyene
	LandVietnam                           = "vn" // Den sosialistiske republikken Vietnam. Før 1976: Den demokratiske republikken Vietnam (VD) og Republikken Vitenam (VN)
	LandVanuatu                           = "vu" // Republikken Vanuatu. Tidligere: Ny Hebridene. Uavhengig siden 1980
	LandWallisOgFutunaøyene               = "wf" // Fransk oversjøisk fellesskap siden 2003
	LandSamoa                             = "ws" // Staten Samoa. Tidligere: Vest-Samoa. Uavhengig siden 1962
	LandJemen                             = "ye" // Republikken Jemen. Før 1990: Den demokratiske folkerepublikken Jemen, yd, (Sør-Jemen) og Den arabiske republikken Jemen, ye, YEM og 886 (Nord-Jemen)
	LandMayotte                           = "yt" // (Oversjøisk fransk departemental besittelse)
	LandSørAfrika                         = "za" // Republikken Sør-Afrika
	LandZambia                            = "zm" // Republikken Zambia. Tidligere del av Nord-Rhodesia. Uavhengig siden 1964
	LandZimbabwe                          = "zw" // Republic of Zimbabwe. Tidligere del av Sør-Rhodesia. Uavhengig siden 1980
)

// Språkkoder
// Hentet fra http://nabo.nb.no/trip?_b=sprakkoder&_n=0&ccl=&_O=kode
const (
	SpråkAceh           = "ace" //  (Acehnesisk) (Aceh-provinsen i Indonesia. Ny kode 2008)
	SpråkAcholi         = "ach" //  (acoli) (Uganda. Ny kode 2014)
	SpråkAfrikaans      = "afr" //  (Sør-Afrika)
	SpråkAkkadisk       = "akk" //  (språk i det gamle Babylonia og Assyria)
	SpråkAlbansk        = "alb"
	SpråkAmharisk       = "amh" //  (Etiopias offisielle språk)
	SpråkAngelsaksisk   = "ang" //  (angelsaksisk engelsk) (ca. 450-1100)
	SpråkArabisk        = "ara"
	SpråkArameisk       = "arc" //  (Palestina, ca. 700 f.Kr.-ca. 700 e.Kr.)
	SpråkArmensk        = "arm"
	SpråkAvestisk       = "ave" //  (Ny kode 2012)
	SpråkAserbajdsjansk = "aze" //  (også kalt aseri, azeri)
	SpråkBambara        = "bam" //  (Vest-Afrika, bl.a. Mali)
	SpråkBaskisk        = "baq" //  (Nord-Spania)
	SpråkHviterussisk   = "bel"
	SpråkBemba          = "bem" //  (Zambia. Ny kode 2014)
	SpråkBengali        = "ben" //  (Bangladesh og Nordøst-India)
	SpråkBislama        = "bis" //  (Vanuatu, tidl. Ny-Hebridene)
	SpråkBantuspråk     = "bnt" // andre
	SpråkBosnisk        = "bos" //  (Ny kode 2004, erstatter delvis cro)
	SpråkBretonsk       = "bre" //  (Nordvest-Frankrike)
	SpråkBulgarsk       = "bul" //  (Kyrillisk skrift)
	SpråkBurmesisk      = "bur" //  (Myanmar, tidl. Burma)
	SpråkBilin          = "byn" //  (blin) (Eritrea. Ny kode 2004, erstatter delvis cus)
	// "cam" // Tidl. brukt for khmer (Utgår 2004, erstattes av khm)
	SpråkKatalansk    = "cat" //  (Nordøst-Spania, Andorra, Sør-Frankrike, Balearene)
	SpråkCebuano      = "ceb" //  (Filippinene)
	SpråkTsjetsjensk  = "che" //  (Ny kode 2013.)
	SpråkKinesisk     = "chi"
	SpråkKirkeslavisk = "chu" //  (Gammelkyrillisk skrift)
	SpråkSorani       = "ckb" //  (Sørkurdisk. Ny kode 2013. Hentet fra ISO 639-3.)
	SpråkKoptisk      = "cop" //  (Egyptisk språk, 2.-14. årh. e.Kr.)
	// "cro" // Tidl.  brukt for kroatisk (latinsk skrift) og bosnisk (Utgått 2004, da erstattet av scr (kroatisk) og bos (bosnisk). scr erstattes av hrv 2011)
	SpråkKusjittiske   = "cus" //  språk, andre (f.eks. iraqw) (Øst-Afrika, erstattes delvis av byn 2004)
	SpråkTsjekkisk     = "cze"
	SpråkDansk         = "dan"
	SpråkNederlandsk   = "dut" //  (inkl. Flamsk) Brukt for: Hollandsk
	SpråkDzongkha      = "dzo" //  (Bhutanesisk, bhotia i Bhutan. Ny kode 2005)
	SpråkEgyptisk      = "egy" //  (demotisk, hieratisk og hieroglyffer) (Se også koptisk)
	SpråkEngelsk       = "eng"
	SpråkMiddelEngelsk = "enm" //  (middelengelsk) (ca. 1100-1500)
	SpråkEsperanto     = "epo" //  (Ny kode 2004, erstatter esp)
	// "esk" // Tidl. brukt for eskimospråk (Utgår 2004, erstattes av iku, ipk, kal og ypk)
	// "esp" // Tidl. brukt for esperanto (Utgår 2004, erstattes av epo)
	SpråkEstisk = "est"
	// "eth" // Tidl. brukt for etiopisk (Utgår 2004, erstattes av gez)
	SpråkEwe     = "ewe" //  (Ghana. Ny kode 2014)
	SpråkFærøysk = "fao" //  (Ny kode 2004, erstatter far)
	// "far" // Tidl. brukt for færøysk (Utgår 2004, erstattes av fao)
	SpråkFante        = "fat" //  (fanti) (Ghana. Ny kode 2014)
	SpråkFijiansk     = "fij" //  (Fiji)
	SpråkFinsk        = "fin"
	SpråkFinskUgriske = "fiu" //  språk, andre (f.eks. karelsk) (Ny kode 2004, erstatter krj)
	SpråkKvensk       = "fkv" //  (Ny kode 2006. Hentet fra ISO 639-3.)
	SpråkFransk       = "fre"
	// "fri" // Tidl. brukt for frisisk (Utgår 2004, erstattes av fry)
	SpråkMellomFransk = "frm" //  (mellomfransk) (ca. 1400-1600)
	SpråkGammelFransk = "fro" //  (gammelfransk) (ca. 842-1400)
	SpråkFrisisk      = "fry" //  (Fryslân (nordnederlandsk provins) og Nordvest-Tyskland. Ny kode 2004 erstatter fri)
	SpråkFulfulde     = "ful" //  (også kalt fulani, fula) (Vest-Afrika)
	// "gae" // Tidl. brukt for gælisk (Utgår 2004, erstattes av gla)
	// "gag" // Tidl.  brukt for galegis/galisisk (Utgår 2004, erstattes av glg)
	SpråkGeorgisk      = "geo"
	SpråkTysk          = "ger"
	SpråkEtiopisk      = "gez" //  (geez) (Dødt språk, men er liturgisk språk for den etiopisk-koptiske kirke. Ny kode 2004, erstatter eth)
	SpråkSkotskGælisk  = "gla" //  (Ny kode 2004, erstatter gae)
	SpråkIrsk          = "gle" //  (Irsk-Gælisk) (Ny kode 2004, erstatter iri)
	SpråkGalegisk      = "glg" // el. Galisisk (Nordvest-Spania. Ny kode 2004, erstatter gag)
	SpråkMidelHøyTysk  = "gmh" //  (middelhøytysk) (ca 1050-1500. Ny kode 2004)
	SpråkGotisk        = "got"
	SpråkKlassiskGresk = "grc" //  klassisk og koine (fram til 1453)
	SpråkGresk         = "gre" //  (etter 1453)
	SpråkGujarati      = "guj" //  (Gujarat, India)
	SpråkHausa         = "hau" //  (Nigeria)
	SpråkHebraisk      = "heb"
	SpråkHiligaynon    = "hil" //  (Filippinene)
	SpråkHindi         = "hin" //  (Offisielt språk for hele India)
	SpråkHettittisk    = "hit" //  (hittittisk) (Ny kode 2009)
	SpråkKroatisk      = "hrv" //  (Erstatter scr 2011)
	SpråkUngarsk       = "hun"
	SpråkIbo           = "ibo" //  (også kalt igbo) (Nigeria)
	SpråkIslandsk      = "ice"
	SpråkInuittisk     = "iku" //  (inuktitut) (Canada. Ny kode 2004, erstatter delvis esk)
	SpråkIloko         = "ilo" //  (Filippinene)
	SpråkIndonesisk    = "ind"
	// "ing" // Tidl. brukt for "ingen tekst, tale eller sang" (Utgår 2004, erstattes av ^^^ (tre blanke tegn) hvis nødvendig.)
	SpråkInupiak = "ipk" //  (Alaska. Ny kode 2004, erstatter delvis esk)
	SpråkIranske = "ira" //  språk, andre (f.eks. dari)
	// "iri" // Tidl.  brukt for irsk (Utgår 2004, erstattes av gle)
	SpråkItaliensk   = "ita"
	SpråkJapansk     = "jpn"
	SpråkKalaalisut  = "kal" //  (grønlandsk) (Ny kode 2004, erstatter delvis esk)
	SpråkKhmer       = "khm" //  (Kambodsja. Ny kode 2004, erstatter cam)
	SpråkKhotanesisk = "kho" //  (Også kalt Khotan-sakisk eller Sakisk)
	SpråkKinyarwanda = "kin" //  (kinirwanda) (Rwanda. (Ny kode 2008))
	SpråkKurmanji    = "kmr" //  (Nordkurdisk. Ny kode 2013. Hentet fra ISO 639-3)
	SpråkKongo       = "kon" //  (også kalt kikongo)
	SpråkKoreansk    = "kor"
	// "krj" // Tidl.  brukt for karelsk (Utgår 2004, erstattes av fiu)
	SpråkKuanyama = "kua" //  (kwanyama) (Angola. Ny kode 2014)
	SpråkKurdisk  = "kur" //  (Latinsk, arabisk og kyrillisk skrift)
	// "lae" // Tidl. brukt for østsamisk (Utgår 2004, erstattes av smi)
	// "laf" // Tidl. brukt for nordsamisk (Utgår 2004, erstattes av sme)
	// "lai" // Tidl. brukt for inarisamisk (Utgår 2004, erstattes av smn)
	// "lak" // Tidl. brukt for skoltesamisk (Utgår 2004, erstattes av sms)
	// "lal" // Tidl. brukt for lulesamisk (Utgår 2004, erstattes av smj)
	SpråkLaotisk = "lao" //  (Laos og Thailand. Ny kode 2015)
	// "lap" // Tidl. brukt som samlekode for alle samiske språk (samisk) (Utgår 2004, erstattes av sma, sme, smi, smj, smn og sms)
	// "las" // Tidl. brukt for sørsamisk (Utgår 2004, erstattes av sma)
	SpråkLatin     = "lat"
	SpråkLatvisk   = "lav"
	SpråkLitauisk  = "lit"
	SpråkMakedonsk = "mac" //  (Den tidligere jugoslaviske republikken Makedonia)
	SpråkMaori     = "mao"
	SpråkMarathi   = "mar" //  (India)
	SpråkMalayisk  = "may" //  (Ny kode 2007)
	// "mis" // Tidl. brukt for "ubestemt" (Utgår 2004, erstattes av und)
	// "mla" // Tidl. brukt for madagassisk (Utgår 2004, erstattes av mlg)
	SpråkMadagassisk          = "mlg" //  (også kalt malgassisk, gassisk) (Madagaskar)
	SpråkMandinka             = "mnk" //  (Ny kode 2013. Hentet fra ISO 639-3.)
	SpråkMoldavisk            = "mol" //  (Moldavia, kyrillisk skrift)
	SpråkMongolsk             = "mon"
	SpråkFlerspråklig         = "mul"
	SpråkMaya                 = "myn" //  (Flere språk. Mexico, Guatemala, Belize )
	SpråkNahuatl              = "nah" //  (Mexico. Brukes for aztekisk og meksikansk. Ny kode 2012.)
	SpråkNdebeleSørAfrika     = "nbl" //  (Sør-Afrika (Transvaal). Ny kode 2008)
	SpråkNdebeleZimbabwe      = "nde" //  (Zimbabwe. Ny kode 2007)
	SpråkNedertysk            = "nds"
	SpråkNepali               = "nep" //  (Ny kode 2008)
	SpråkNigerKongospråk      = "nic" //  andre (f.eks. mbum) (Inntil 2004 kun brukt for Mbum)
	SpråkNorskNynorsk         = "nno" //  (nynorsk) (Brukes også for landsmål og dialekter)
	SpråkNorskBokmål          = "nob" //  (bokmål) (Ny kode 2004, erstatter nor i dets tidligere betydning. Brukes også for riksmål)
	SpråkNorskMellomnorsk     = "nom" //  (mellomnorsk) (ca. 1350-ca. 1550 (*))
	SpråkNorskGammelnorsk     = "non" //  (gammelnorsk), Islandsk (gammelislandsk), norrønt (I Norge til ca. 1350, på Island til ca. 1550)
	SpråkALTERNATIV           = "nor" //  kode for alt norsk (Ny betydning 2004, brukes dersom det ikke skilles mellom bokmål, nynorsk eller andre målformer. Tidligere brukt for bokmål, som nå kodes nob)
	SpråkProvençalskOksitansk = "oci" //  (langue d'oc, oksitansk) etter 1500 (Sør-Frankrike, ny kode 2004, erstatter delvis pro)
	SpråkOromo                = "orm" //  (Etiopisk folkegruppe. Ny kode 2004)
	SpråkPanjabi              = "pan" //  (Punjab, Nordvest-India)
	SpråkFarsi                = "per" //  (Persisk) (Iran)
	SpråkPali                 = "pli"
	SpråkPolsk                = "pol"
	SpråkPortugisisk          = "por" //  (Portugal og Brasil)
	SpråkPrakrit              = "pra" //  (prakritspråk er fellesbetegnelse på mellomindiske språk)
	SpråkProvençalsk          = "pro" //  til 1500 (Tidl. brukt for alt provençalsk, erstattes delvis av oci 2004)
	SpråkDari                 = "prs" //  (Afghanistan. Ny kode 2013. Hentet fra ISO 639-3.)
	SpråkPashto               = "pus" //  (pushto) (Afganistan, Pakistan)
	SpråkRetoromansk          = "roh" //  (Sørlige Sveits og tilgrensende områder i Nord-Italia)
	SpråkRomani               = "rom" //  (Felles kode for alle sigøynerspråk)
	SpråkRumensk              = "rum" //  (Romania)
	SpråkRundi                = "run" //  (kirundi) (Burundi. Ny kode 2014)
	SpråkRussisk              = "rus" //  (Kyrillisk skrift)
	SpråkJakutisk             = "sah" //  (sakha)
	SpråkSanskrit             = "san"
	// "sao" // Tidl. brukt for samoansk (Utgår 2004, erstattes av smo)
	SpråkSantali = "sat"
	// "scc" // Tidl. brukt for serbokroatisk (kyrillisk skrift), og serbisk (Erstattet ser 2004. scc erstattes av srp 2011)
	SpråkLavskotsk = "sco" //  skotsk-engelsk og lallans (Ny kode 2014)
	// "scr" // Tidl. brukt for serbokroatisk (latinsk skrift), og kroatisk (Erstattet delvis cro 2004. Erstattes av hrv 2011)
	// "ser" // Tidl. brukt for serbisk (Utgått 2004, da erstattet av scc, som erstattes av srp 2011)
	SpråkTegnspråk      = "sgn" //  (Her: alle tegnspråk for døve)
	SpråkSingalesisk    = "sin" //  (Sri Lanka. Ny kode 2004, erstatter snh)
	SpråkSinoTibetanske = "sit" //  språk, andre (f.eks. bodo (mech))
	SpråkSlovakisk      = "slo"
	SpråkSlovensk       = "slv"
	SpråkSørsamisk      = "sma" //  (Midt-Norge, Midt-Sverige. Ny kode 2004, erstatter las og delvis lap)
	SpråkNordsamisk     = "sme" //  (Troms og Finnmark. Ny kode 2004, erstatter laf og delvis lap)
	SpråkSamiske        = "smi" //  språk, andre (f.eks. kildinsamisk) (Ny kode 2004, erstatter lae og delvis lap)
	SpråkLulesamisk     = "smj" //  (Nord-Sverige. Ny kode 2004, erstatter lal og delvis lap)
	SpråkInarisamisk    = "smn" //  (enaresamisk) (Nord-Finland. Ny kode 2004, erstatter lai og delvis lap)
	SpråkSamoansk       = "smo" //  (Ny kode 2004, erstatter sao)
	SpråkSkoltesamisk   = "sms" //  (Finnmark og Murmansk. Ny kode 2004, erstatter lak og delvis lap)
	SpråkShona          = "sna" //  (Ny kode 2007)
	SpråkSindhi         = "snd" //  (Ny kode 2007)
	// "snh" // Tidl. brukt for singalesisk (Utgår 2004, erstattes av sin)
	SpråkSomali   = "som"
	SpråkSpansk   = "spa"
	SpråkSerbisk  = "srp" //  (Erstatter scc 2011)
	SpråkSumerisk = "sux"
	SpråkSwahili  = "swa" //  (Øst-Afrika)
	SpråkSvensk   = "swe"
	SpråkSyrisk   = "syr" //  (Kristent kirkespråk)
	// "tag" // Tidl. brukt for tagalog (Utgår 2004, erstattes av tgl)
	SpråkTamil        = "tam"
	SpråkTagalog      = "tgl" //  (Filippinenes offisielle språk. Ny kode 2004, erstatter tag)
	SpråkThai         = "tha"
	SpråkTibetansk    = "tib"
	SpråkTigrinja     = "tir" //  (Etiopia, Eritrea)
	SpråkTokelaisk    = "tkl" //  (Tokelau)
	SpråkTongansk     = "ton" //  (Tonga)
	SpråkTyrkisk      = "tur"
	SpråkTwi          = "twi" //  (Ghana. Ny kode 2013.)
	SpråkUigurisk     = "uig" //  (Xinjiang, autonom provins i Kina. Ny kode 2005)
	SpråkUkrainsk     = "ukr"
	SpråkUbestemt     = "und" //  (Ny kode 2004, erstatter mis og ubs (fra NORMARC))
	SpråkUrdu         = "urd"
	SpråkUsbekisk     = "uzb"
	SpråkVietnamesisk = "vie"
	SpråkWalisisk     = "wel"
	SpråkSorbiske     = "wen" //  språk, brukt for: vendisk (Sørvest-Polen og Sørøst-Tyskland)
	SpråkXhosa        = "xho" //  (Sør-Afrika. Ny kode 2014)
	SpråkJiddish      = "yid" //  (Tysk-hebraisk)
	SpråkYoruba       = "yor" //  (Ny kode 2013.)
	SpråkYupikspråk   = "ypk" //  (Alaska, Sibir. Ny kode 2004, erstatter delvis esk)
	SpråkZulu         = "zul" //  (Bantuspråk i det sørlige Afrika)
)
