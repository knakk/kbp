package list82

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	ApocryphaCatholicCanon                          = "AP"
	ApocryphaCanonUnspecified                       = "AQ"
	AdditionalApocryphalTextsGreekOrthodoxCanon     = "AX"
	AdditionalApocryphalTextsSlavonicOrthodoxCanon  = "AY"
	AdditionalApocryphalTexts                       = "AZ"
	GeneralCanonWithApocryphaCatholicCanon          = "GA"
	GeneralCanonWithApocryphalTextsCanonUnspecified = "GC"
	GeneralCanon                                    = "GE"
	Gospels                                         = "GS"
	NewTestamentWithPsalmsAndProverbs               = "NP"
	NewTestament                                    = "NT"
	OldTestament                                    = "OT"
	PaulsEpistles                                   = "PE"
	PsalmsAndProverbs                               = "PP"
	Psalms                                          = "PS"
	Pentateuch                                      = "PT"
	OtherPortions                                   = "ZZ"
)

var (
	sortedCodes = []string{"AP", "AQ", "AX", "AY", "AZ", "GA", "GC", "GE", "GS", "NP", "NT", "OT", "PE", "PP", "PS", "PT", "ZZ"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		ApocryphaCatholicCanon:                          {"Apocrypha (Catholic canon)", "The seven portions of the Apocrypha added to the Catholic canon at the Council of Trent in 1546: Tobit; Judith; Wisdom of Solomon; Sirach (Ecclesiasticus); Baruch, including the Letter of Jeremiah; I and II Maccabees; Extra portions of Esther and Daniel (Additions to Esther; the Prayer of Azariah; Song of the Three Jews; Susannah; Bel and the Dragon). These are not generally included in the Protestant canon.", "Apokryfer (katolsk kanon)", "Apokryfiske tekster som ble lagt til den katolske kanon av Konsilet i Trient i 1546: Tobits bok, Judits bok, Visdommens bok, Siraks bok, Baruks bok (inkl. Jeremias brev), første og andre Makkabeeerbok, tillegg til Esters og Daniels bok. Disse inkluderes vanligvis ikke i den protestantiske kanon"},
		ApocryphaCanonUnspecified:                       {"Apocrypha (canon unspecified)", "A collection of Apocryphal texts, canon not specified.", "Apokryfer (kanon er ikke angitt)", "Samling av apokryfiske tekster, kanon er ikke spesifisert"},
		AdditionalApocryphalTextsGreekOrthodoxCanon:     {"Additional Apocryphal texts: Greek Orthodox canon", "I Esdras; Prayer of Manasseh; Psalm 151; III Maccabees.", "Ytterligere apokryfiske tekster; gresk-ortodoks kanon", "Første Esra, Manasses bønn, Salme 151, Tredje Makkabeerbok"},
		AdditionalApocryphalTextsSlavonicOrthodoxCanon:  {"Additional Apocryphal texts: Slavonic Orthodox canon", "I and II Esdras; Prayer of Manasseh; Psalm 151; III and IV Maccabees.", "Ytterligere apokryfiske teksterM slavisk-ortodoks kanon", "Første og andre Esra; Manasses bønn; Salme 151; Tredje og fjerde Makkabeerbok"},
		AdditionalApocryphalTexts:                       {"Additional Apocryphal texts", "Additional Apocryphal texts included in some Bible versions: I and II Esdras; Prayer of Manasseh.", "Ytterligere apokryfiske tekster", "Ytterligere apokryfiske tekster, inkludert i noen Bibelversjoner: Første og andre Esra, Manasses bønn"},
		GeneralCanonWithApocryphaCatholicCanon:          {"General canon with Apocrypha (Catholic canon)", "The 66 books included in the Protestant, Catholic and Orthodox canons, together with the seven portions of the Apocrypha included in the Catholic canon.", "Generell kanon med apokryfer (katolsk kanon)", "De 66 bøkene som er inkludert i protestantisk, katolsk og ortodoks kanon, sammen med de sju apokryfiske tekstene som er inkludert i den katolske kanon."},
		GeneralCanonWithApocryphalTextsCanonUnspecified: {"General canon with Apocryphal texts (canon unspecified)", "The 66 books included in the Protestant, Catholic and Orthodox canons, together with Apocryphal texts, canon not specified.", "Generell kanon med apokryfer (kanon er ikke angitt)", "De 66 bøkene som er inkludert i protestantisk, katolsk og ortodoks kanon, sammen med apokryfiske tekster, kanon er ikke spesifisert."},
		GeneralCanon:                                    {"General canon", "The 66 books included in the Protestant, Catholic and Orthodox canons, 39 from the Old Testament and 27 from the New Testament. The sequence of books may differ in different canons.", "Generell kanon", "De 66 bøkene som er inkludert i protestantisk, katolsk og ortodoks kanon, 39 fra Det gamle testamentet og 27 fra Det nye testamentet. Rekkefølgen kan variere i de ulike kanonene."},
		Gospels:                                         {"Gospels", "The books of Matthew, Mark, Luke and John.", "Evangeliene", "Evangeliene etter Matteus, Markus, Lukas og Johannnes"},
		NewTestamentWithPsalmsAndProverbs: {"New Testament with Psalms and Proverbs", "Includes the 27 books of the New Testament plus Psalms and Proverbs from the Old Testament.", "Det nye testamentet med Salmene og Ordspråkene", "Inkluderer Det nye testamentets 27 bøker, Salmenes bok og Salomos ordspråk fra Det gamle testamentet."},
		NewTestament:                      {"New Testament", "The 27 books included in the Christian canon through the Easter Letter of Athanasius, Bishop of Alexandria and also by a general council of the Christian church held near the end of the 4th century CE.", "Det nye testamente", "De 27 bøkene som ble inkludert i den kristne kanon gjennom Påskebrevet fra Athanasius, biskop av Alexandria og også ved et generelt konsil for den kristne kirke, holdt nær enden av det 4. århundret e.Kr."},
		OldTestament:                      {"Old Testament", "Those 39 books which were included in the Jewish canon by the rabbinical academy established at Jamma in 90 CE. Also known as the Jewish or Hebrew scriptures.", "Det gamle testamente", "De 39 bøkene som ble inkludert i den jødiske kanon av akademiet for rabbinere, etablert i Jamma i år 90 e.Kr."},
		PaulsEpistles:                     {"Paul’s Epistles", "The books containing the letters of Paul to the various early Christian churches.", "Paulus' brev", "Paulus' brev til de forskjellige tidlige, kristne kirkene"},
		PsalmsAndProverbs:                 {"Psalms and Proverbs", "The book of Psalms and the book of Proverbs combined.", "Salmene og ordspråkene", "Salmenes bok og Salomos ordspråk"},
		Psalms:                            {"Psalms", "The book of Psalms.", "Salmenes bok", "Salmenes bok fra Det gamle testamentet"},
		Pentateuch:                        {"Pentateuch", "The first five books of the Bible: Genesis, Exodus, Numbers, Leviticus, Deuteronomy. Also applied to the Torah.", "Mosebøkene", "De fem første bøkene av  Bibelen. Tilsvarer  Torah."},
		OtherPortions:                     {"Other portions", "Selected books of either the OT or NT not otherwise noted.", "Andre deler", "Utvalgte deler av Bibelen som ikke nevnt andre steder"},
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
