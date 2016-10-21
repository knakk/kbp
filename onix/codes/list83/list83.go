package list83

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	AlbertoVaccari                               = "ALV"
	Amplified                                    = "AMP"
	AntonioMartini                               = "ANM"
	AmericanStandard                             = "ASV"
	ConferenzaEpiscopaleItaliana                 = "CEI"
	ConferenzaEpiscopaleItaliana2008             = "CEN"
	ContemporaryEnglish                          = "CEV"
	Concordata                                   = "CNC"
	Diodati                                      = "DDI"
	NuovaDiodati                                 = "DDN"
	DouayRheims                                  = "DOU"
	Einheitsübersetzung                          = "EIN"
	EnglishStandard                              = "ESV"
	Biblia1776                                   = "FBB"
	Raamattu19331938                             = "FRA"
	RaamattuKansalle                             = "FRK"
	Raamattu1992                                 = "FRM"
	GodsWord                                     = "GDW"
	Geneva                                       = "GEN"
	GoodNews                                     = "GNB"
	GalbiatiPennaRossanoUTET                     = "GPR"
	OriginalGreek                                = "GRK"
	GarofanoRinaldiMarietti                      = "GRM"
	OriginalHebrew                               = "HBR"
	HolmanChristianStandard                      = "HCS"
	InternationalChildrens                       = "ICB"
	TraduzioneInterconfessionaleInLinguaCorrente = "ILC"
	Jerusalem                                    = "JER"
	KingJames21stCentury                         = "KJT"
	KingJames                                    = "KJV"
	LivingBible                                  = "LVB"
	Luzzi                                        = "LZZ"
	MessageBible                                 = "MSG"
	NewAmerican                                  = "NAB"
	NewAmericanStandard                          = "NAS"
	NewAmericanStandardUpdated                   = "NAU"
	Bibelen1895                                  = "NBA"
	Bibelen1930                                  = "NBB"
	Bibelen1938                                  = "NBC"
	Bibelen197885                                = "NBD"
	Bibelen1978                                  = "NBE"
	Bibelen1985                                  = "NBF"
	Bibelen1988                                  = "NBG"
	Bibelen197885Rev2005                         = "NBH"
	Bibelen2011                                  = "NBI"
	NewCentury                                   = "NCV"
	NewEnglish                                   = "NEB"
	BibelenGudsOrd                               = "NGO"
	NewInternationalReaders                      = "NIR"
	NewInternational                             = "NIV"
	NewJerusalem                                 = "NJB"
	NewKingJames                                 = "NKJ"
	NewLiving                                    = "NLV"
	BibelenNynorsk                               = "NNK"
	NewRevisedStandard                           = "NRS"
	NuevaTraduccionVivienta                      = "NTV"
	NovissimaVersioneDellaBibbia                 = "NVB"
	NuevaBibliaAlDia                             = "NVD"
	NuevaVersionInternacional                    = "NVI"
	NewTestamentInModernEnglishPhillips          = "PHP"
	RevisedEnglish                               = "REB"
	RevisedVersion                               = "REV"
	RevisedStandard                              = "RSV"
	ReinaValera                                  = "RVL"
	Bibel2000                                    = "SBB"
	BibelenSamisk                                = "SMK"
	TodaysEnglish                                = "TEV"
	TodaysNewInternational                       = "TNI"
	Other                                        = "ZZZ"
)

var (
	sortedCodes = []string{"ALV", "AMP", "ANM", "ASV", "CEI", "CEN", "CEV", "CNC", "DDI", "DDN", "DOU", "EIN", "ESV", "FBB", "FRA", "FRK", "FRM", "GDW", "GEN", "GNB", "GPR", "GRK", "GRM", "HBR", "HCS", "ICB", "ILC", "JER", "KJT", "KJV", "LVB", "LZZ", "MSG", "NAB", "NAS", "NAU", "NBA", "NBB", "NBC", "NBD", "NBE", "NBF", "NBG", "NBH", "NBI", "NCV", "NEB", "NGO", "NIR", "NIV", "NJB", "NKJ", "NLV", "NNK", "NRS", "NTV", "NVB", "NVD", "NVI", "PHP", "REB", "REV", "RSV", "RVL", "SBB", "SMK", "TEV", "TNI", "ZZZ"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		AlbertoVaccari:                               {"Alberto Vaccari", "Alberto Vaccari – Pontificio Istituto Biblico.", "Alberto Vaccari", "Alberto Vaccari - Pontificio Istituto Biblico"},
		Amplified:                                    {"Amplified", "A translation based on the American Standard Version and showing multiple options for the translation of ancient text. Published in full in 1965. Sponsored by the Lockman Foundation.", "Amplified", "A translation based on the American Standard Version and showing multiple options for the translation of ancient text. Published in full in 1965.  Sponsored by the Lockman Foundation."},
		AntonioMartini:                               {"Antonio Martini", "Most popular Catholic Bible translation in Italian prior to the CEI translation in 1971.", "Antonio Martini", "Most popular Catholic Bible translation in Italian prior to the CEI translation in 1971."},
		AmericanStandard:                             {"American Standard", "A 1901 translation using verbal equivalence techniques with the purpose of Americanizing the REV.", "American Standard", "A 1901 translation using verbal equivalence techniques with the purpose of Americanizing the King James version."},
		ConferenzaEpiscopaleItaliana:                 {"Conferenza Episcopale Italiana", "Italian Episcopal Conference 1971 translation suitable for Italian Catholic liturgy. (Includes minor 1974 revision).", "Conferenza Episcopale Italiana", "Italian Episcopal Conference 1971 translation suitable for Italian Catholic liturgy. (Includes minor 1974 revision)."},
		ConferenzaEpiscopaleItaliana2008:             {"Conferenza Episcopale Italiana 2008", "New translation of the C.E.I. first published in 2008 – the version most widely used by the Italian Catholic Church.", "Conferenza Episcopale Italiana 2008", "New translation of the C.E.I. first published in 2008 – the version most widely used by the Italian Catholic Church."},
		ContemporaryEnglish:                          {"Contemporary English", "A translation completed in 1995 and sponsored by the American Bible Society under the leadership of Barclay Newman.", "Contemporary English", "A translation completed in 1995 and sponsored by the American Bible Society under the leadership of Barclay Newman."},
		Concordata:                                   {"Concordata", "1968 Interfaith version promoted by the Italian Bible Society. Has a Catholic ‘imprimateur’, but its ecumenical approach has Jewish, Protestant and Christian Orthodox approval.", "Concordata", "1968 Interfaith version promoted by the Italian Bible Society. Has a Catholic ‘imprimateur’, but its ecumenical approach has Jewish, Protestant and Christian Orthodox approval."},
		Diodati:                                      {"Diodati", "Version based on original documents, edited by Giovanni Diodati in 1607, revised by Diodati in 1641 and again in 1894. It is the reference version for many Italian Protestants.", "Diodati", "Version based on original documents, edited by Giovanni Diodati in 1607, revised by Diodati in 1641 and again in 1894. It is the reference version for many Italian Protestants."},
		NuovaDiodati:                                 {"Nuova Diodati", "Revision of the Diodati Bible dating to the 1990s, aiming at highest fidelity to original ancient Greek (New Testament) and Hebrew (Old Testament) texts.", "Nuova Diodati", "Revision of the Diodati Bible dating to the 1990s, aiming at highest fidelity to original ancient Greek (New Testament) and Hebrew (Old Testament) texts."},
		DouayRheims:                                  {"Douay-Rheims", "An early (1580-1609) English translation from the Latin Vulgate designed for Catholics and performed by George Martin.", "Douay-Rheims", "An early (1580-1609) English translation from the Latin Vulgate designed for Catholics and performed by George Martin."},
		Einheitsübersetzung:                          {"Einheitsübersetzung", "A German translation of the Bible for use in Roman Catholic churches.", "Einheitsübersetzung", "Tysk oversettelse av bibelen til bruk i romersk-katolske kirker"},
		EnglishStandard:                              {"English Standard", "An update of the Revised Standard Version that makes ‘modest’ use of gender-free terminology.", "English Standard", "An update of the Revised Standard Version  that makes 'modest' use of gender-free terminology."},
		Biblia1776:                                   {"Biblia (1776)", "Finnish Bible translation.", "Biblia (1776)", "Finnish Bible translation."},
		Raamattu19331938:                             {"Raamattu (1933/1938)", "Finnish Bible translation.", "Raamattu (1933/1938)", "Finnish Bible translation."},
		RaamattuKansalle:                             {"Raamattu kansalle", "Finnish Bible translation.", "Raamattu kansalle", "Finnish Bible translation."},
		Raamattu1992:                                 {"Raamattu (1992)", "Finnish Bible translation.", "Raamattu (1992)", "Finnish Bible translation."},
		GodsWord:                                     {"God’s Word", "A 1995 translation by the World Bible Publishing Company using the English language in a manner to communicate to the late 20th century American.", "God's Word", "A 1995 translation by the World Bible Publishing Company using the English language in a manner to communicate to the late 20th century American."},
		Geneva:                                       {"Geneva", "An early (1560) English version of the Bible translated by William Whittingham with strong Protestant leanings.", "Geneva", "An early (1560) English version of the Bible translated by William Whittingham with strong Protestant leanings."},
		GoodNews:                                     {"Good News", "A translation sponsored by the American Bible Society. The New Testament was first published (as “Today’s English Version” TEV) in 1966. The Old Testament was completed in 1976, and the whole was published as the “Good News Bible”.", "Good News", "A translation sponsored by the American Bible Society.  The New Testament was first published (as Today's English Version' TEV) in 1966.  The Old Testament was completed in 1976 and the whole was published as the Good News Bible"},
		GalbiatiPennaRossanoUTET:                     {"Galbiati, Penna, Rossano – UTET", "Version edited by E. Galbiati, A. Penna and P. Rossano, and published by UTET. This version, based on original texts, is rich in notes and has been used as the basis for CEI translation.", "Galbiati, Penna, Rossano", "E. Galbiati A. Penna P. Rossano - UTET"},
		OriginalGreek:                                {"Original Greek", "New Testament text in an original Greek version.", "Original Greek", "New Testament text in an original Greek version"},
		GarofanoRinaldiMarietti:                      {"Garofano, Rinaldi – Marietti", "Richly annotated 1963 Version edited by S. Garofano and S. Rinaldi, and published by Marietti.", "Garofano, Rinaldi - Marietti", "S. Garofano S. Rinaldi - Marietti"},
		OriginalHebrew:                               {"Original Hebrew", "Old Testament text in an original Hebrew version.", "Original Hebrew", "Old Testament text in an original Hebrew version"},
		HolmanChristianStandard:                      {"Holman Christian Standard", "Published by Broadman and Holman this translation rejects all forms of gender-neutral wording and is written with strong influences from the Southern Baptist perspective of biblical scholarship.", "Holman Christian Standard", "Published by Broadman and Holman this translation rejects all forms of gender-neutral wording and is written with strong influences from the Southern Baptist perspective of biblical scholarship."},
		InternationalChildrens:                       {"International Children’s", "A translation completed in 1986 targeting readability at the US third grade level.", "International Children's", "A translation completed in 1986 targeting readability at the US third grade level."},
		TraduzioneInterconfessionaleInLinguaCorrente: {"Traduzione Interconfessionale in Lingua Corrente", "Interconfessional translation resulting from 1985 effort by Catholic and Protestant scholars, aimed at delivering an easy-to-understand message.", "Traduzione Interconfessionale in Lingua Corrente", "Interconfessional translation resulting from 1985 effort by Catholic and Protestant scholars, aimed at delivering an easy-to-understand message."},
		Jerusalem:                           {"Jerusalem", "A translation designed for English speaking Catholics based on the original languages. It is based on French as well as ancient texts and was first published in 1966.", "Jerusalem", "A translation designed for English speaking Catholics based on the original languages.  It is based on French as well as ancient texts and was first published in 1966."},
		KingJames21stCentury:                {"21st Century King James", "A verbal translation led by William Prindele. Published in 1994, it was designed to modernize the language of the King James Version based on Webster’s New International Dictionary, 2nd edition, unabridged.", "21st Century King James", "A verbal translation led by William Prindele.  Published in 1994, it was designed to modernize the language of the King James Version based on Webster's New International Dictionary, 2nd edition, unabridged."},
		KingJames:                           {"King James", "A translation commissioned by King James I of England and first published in 1611.", "King James", "A translation commissioned by King James I of England and first published in 1611."},
		LivingBible:                         {"Living Bible", "A paraphrase translation led by Kenneth N Taylor and first published in 1972.", "Living Bible", "A paraphrase translation led by Kenneth N Taylor and first published in 1972."},
		Luzzi:                               {"Luzzi", "1924 translation by Giovanni Luzzi, Professor at the Waldensian Faculty of Theology in Rome, who revised the 17th Century Diodati version.", "Luzzi", "1924 translation by Giovanni Luzzi, Professor at the Waldensian Faculty of Theology in Rome, who revised the 17th Century Diodati version."},
		MessageBible:                        {"Message Bible", "A paraphrase translation of the New Testament by Eugene Peterson first published in 1993.", "Message Bible", "A paraphrase translation of the New Testament by Eugene Peterson first published in 1993."},
		NewAmerican:                         {"New American", "A translation aimed at Catholic readers first published in its entirety in 1970. A revised New Testament was issued in 1986 as the 2nd Edition. The 3rd Edtion was published in 1991 with a revision to Psalms. The 4th Edition (also known as the New American Bible Revised Edition) was published in 2011, incorporating revisions to the Old Testament.", "New American", "A translation aimed at Catholic readers first published in its entirely in 1970.  A revised New Testament was issued in 1986."},
		NewAmericanStandard:                 {"New American Standard", "A translation commissioned by the Lockman Foundation. The New Testament was published in 1960 followed by the entire Bible in 1971.", "New American Standard", "A translation commissioned by the Lockman Foundation.  The New Testament was published in 1960 followed by the entire Bible in 1971."},
		NewAmericanStandardUpdated:          {"New American Standard, Updated", "A 1995 translation using more modern language than the NASB.", "New American Standard, Updated", "A 1995 translation using more modern language than the NASB.  "},
		Bibelen1895:                         {"Bibelen 1895", "Norwegian Bible translation.", "Bibelen 1895", "Norsk bibeloversettelse"},
		Bibelen1930:                         {"Bibelen 1930", "Norwegian Bible translation.", "Bibelen 1930", "Norsk bibeloversettelse"},
		Bibelen1938:                         {"Bibelen 1938", "Norwegian Bible translation.", "Bibelen 1938", "Norsk bibeloversettelse"},
		Bibelen197885:                       {"Bibelen 1978-85", "Norwegian Bible translation.", "Bibelen 1978-85", "Norsk bibeloversettelse"},
		Bibelen1978:                         {"Bibelen 1978", "Norwegian Bible translation.", "Bibelen 1978", "Norsk bibeloversettelse"},
		Bibelen1985:                         {"Bibelen 1985", "Norwegian Bible translation.", "Bibelen 1985", "Norsk bibeloversettelse"},
		Bibelen1988:                         {"Bibelen 1988", "Norwegian Bible translation.", "Bibelen 1988", "Norsk bibeloversettelse"},
		Bibelen197885Rev2005:                {"Bibelen 1978-85/rev. 2005", "Norwegian Bible translation.", "Bibelen 1978-85/rev. 2005", "Norsk bibeloversettelse"},
		Bibelen2011:                         {"Bibelen 2011", "Norwegian Bible translation.", "Bibelen 2011", "Norsk bibeloversettelse"},
		NewCentury:                          {"New Century", "A translation inspired by the International Children’s version. First published by World Publishing in 1991.", "New Century", "A translation inspired by the International Children's version.  First published by World Publishing in 1991."},
		NewEnglish:                          {"New English", "A translation first issued in 1961 (New Testament) and 1970 (complete Bible) as a result of a proposal at the 1946 General Assembly of the Church of Scotland.", "New English", "A translation first issued in 1970 as a result of a proposal at the 1946 General Assembly of the Church of Scotland."},
		BibelenGudsOrd:                      {"Bibelen Guds ord", "Norwegian Bible translation.", "Bibelen Guds ord", "Norsk bibeloversettelse. Bygger på Nye King James"},
		NewInternationalReaders:             {"New International Reader’s", "A 1996 translation designed for people with limited literacy in English and based on the NIV.", "New International Reader's", "A 1996 translation designed for people with limited literacy in English and based on the NIV."},
		NewInternational:                    {"New International", "A translation underwritten by Biblica (formerly the International Bible Society, and previously the New York Bible Society). The New Testament was published in 1973 followed by the entire Bible in 1978. The NIV text was revised in 1984 and again in 2011.", "New International", "A translation underwritten by the International Bible Society (formerly New York Bible Society).  The New Testament was published in 1973 followed by the entire Bible in 1978."},
		NewJerusalem:                        {"New Jerusalem", "A revision of the Jerusalem Bible. First published in 1986.", "New Jerusalem", "A revision of the Jerusalem Bible.  First published in 1986."},
		NewKingJames:                        {"New King James", "A version issued by Thomas Nelson Publishers in 1982-83 designed to update the language of the King James Version while maintaining the phrasing and rhythm and using the same sources as its predecessor.", "New King James", "A version issued by Thomas Nelson Publishers in 1982-83 designed to update the language of the King James Version while maintaining the phrasing and rhythm and using the same sources as its predecessor."},
		NewLiving:                           {"New Living", "A translation sponsored by Tyndale House and first released in 1996. It is considered a revision and updating of the Living Bible.", "New Living", "A translation sponsored by Tyndale House and first released in 1996.  It is considered a revision and updating of the Living Bible."},
		BibelenNynorsk:                      {"Bibelen, nynorsk", "Norwegian ‘nynorsk’ Bible translation.", "Bibelen, nynorsk", "Norsk bibeloversettelse til nynorsk"},
		NewRevisedStandard:                  {"New Revised Standard", "A revision of the Revised Standard based on ancient texts but updating language to American usage of the 1980s.", "New Revised Standard", "A revision of the Revised Standard based on ancient texts but updating language to American usage of the 1980s."},
		NuevaTraduccionVivienta:             {"Nueva Traduccion Vivienta", "A Spanish translation from the original Greek and Hebrew, sponsored by Tyndale House.", "Nueva Traduccion Vivienta", "A Spanish translation from the original Greek and Hebrew, sponsored by Tyndale House."},
		NovissimaVersioneDellaBibbia:        {"Novissima Versione della Bibbia", "Nuovissima version – a Catholic-oriented translation in modern Italian, edited by a group including Carlo Martini, Gianfranco Ravasi and Ugo Vanni and first published (in 48 volumes, 1967-1980) by Edizioni San Paolo.", "Novissima Versione della Bibbia", "Nuovissima version – a Catholic-oriented translation in modern Italian, edited by a group including Carlo Martini, Gianfranco Ravasi and Ugo Vanni and first published (in 48 volumes, 1967-1980) by Edizioni San Paolo."},
		NuevaBibliaAlDia:                    {"Nueva Biblia al Dia", "A Spanish translation from the original Greek and Hebrew, sponsored by the International Bible Society/Sociedad Bíblica Internacional.", "Nueva Biblia al Dia", "A Spanish translation from the original Greek and Hebrew, sponsored by the International Bible Society/Sociedad Bíblica Internacional."},
		NuevaVersionInternacional:           {"Nueva Version Internacional", "A Spanish translation underwritten by the International Bible Society.", "Nueva Version Internacional", "A Spanish translation underwritten by the International Bible Society."},
		NewTestamentInModernEnglishPhillips: {"New Testament in Modern English (Phillips)", "An idiomatic translation by J B Phillips, first completed in 1966.", "New Testament in Modern English (Phillips)", "An idiomatic translation by J B Phillips, first completed in 1966"},
		RevisedEnglish:                      {"Revised English", "A 1989 revision of the NEB. A significant effort was made to reduce the British flavor present in the NEB.", "Revised English", "A 1989 revision of the NEB.  A significant effort was made to reduce the British flavor present in the NEB."},
		RevisedVersion:                      {"Revised Version", "The first major revision of the King James Version, the Revised Version incorporates insights from early manuscripts discovered between 1611 and 1870, and corrects readings in the KJV which nineteenth-century scholarship deemed mistaken. The New Testament was published in 1881, the Old Testament in 1885, and the Apocrypha in 1895.", "Revised Version", "The first major revision of the King James Version, the Revised Version incorporates insights from early manuscripts discovered between 1611 and 1870, and corrects readings in the KJV which nineteenth-century scholarship deemed mistaken.  The New Testament was published in 1881, the Old Testament in 1885, and the Apocrypha in 1895."},
		RevisedStandard:                     {"Revised Standard", "A translation authorized by the National Council of Churches of Christ in the USA. The New Testament was published in 1946 followed by a complete Protestant canon in 1951.", "Revised Standard", "A translation authorized by the National Council of Churches of Christ in the USA.  The New Testament was published in 1946 followed by a complete Protestant canon in 1951."},
		ReinaValera:                         {"Reina Valera", "A Spanish translation based on the original texts.", "Reina Valera", "A Spanish translation based on the original texts."},
		Bibel2000:                           {"Bibel 2000", "Swedish Bible translation.", "Bibel 2000", "Svensk bibeloversettelse"},
		BibelenSamisk:                       {"Bibelen, samisk", "Norwegian ‘samisk’ Bible translation.", "Bibelen, samisk", "Norsk bibeloversettelse til samisk"},
		TodaysEnglish:                       {"Today’s English", "A translation of the New Testament sponsored by the American Bible Society and first published in 1966. It was incorporated into the “Good News Bible” GNB in 1976.", "Today's English", "A translation of the New Testament sponsored by the American Bible Society and first published in 1966. It was incorporated into the Good News Bible' GNB in 1976."},
		TodaysNewInternational:              {"Today’s New International", "An updating of the New International Version. The New Testament was published in 2002, and the entire Bible in 2005. Superseded by the 2011 NIV update.", "Today's New International", "An updating of the New International Version.  The New Testament was published in 2002, and the entire Bible is scheduled for 2005."},
		Other: {"Other", "Other translations not otherwise noted.", "Other", "Other translations not otherwise noted."},
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
