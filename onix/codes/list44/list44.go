package list44

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	Proprietary                         = "01"
	Proprietary                         = "02"
	DNBPublisherIdentifier              = "03"
	BörsenvereinVerkehrsnummer          = "04"
	GermanISBNAgencyPublisherIdentifier = "05"
	GLN                                 = "06"
	SAN                                 = "07"
	CentraalBoekhuisRelatieID           = "10"
	FondscodeBoekenbank                 = "13"
	YTunnus                             = "15"
	ISNI                                = "16"
	PND                                 = "17"
	LCCN                                = "18"
	JapanesePublisherIdentifier         = "19"
	GKD                                 = "20"
	ORCID                               = "21"
	GAPPPublisherIdentifier             = "22"
	VATIdentityNumber                   = "23"
	JPDistributionIdentifier            = "24"
	GND                                 = "25"
	DUNS                                = "26"
	RinggoldID                          = "27"
	IdentifiantEditeurElectre           = "28"
	IdentifiantEditeurBTLF              = "29"
	IdentifiantMarqueElectre            = "30"
	IdentifiantMarqueBTLF               = "31"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "10", "13", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		Proprietary:                         {"Proprietary", "", "Proprietær", ""},
		Proprietary:                         {"Proprietary", "DEPRECATED – use 01.", "Proprietær", "UTGÅTT. Bruk 01"},
		DNBPublisherIdentifier:              {"DNB publisher identifier", "Deutsche Nationalbibliothek publisher identifier.", "Deutsche Bibliothek publisher identifier", "Deutsche Nationalbibliothek publisher identifier."},
		BörsenvereinVerkehrsnummer:          {"Börsenverein Verkehrsnummer", "", "Börsenverein Verkehrsnummer", ""},
		GermanISBNAgencyPublisherIdentifier: {"German ISBN Agency publisher identifier", "", "German ISBN Agency publisher identifier", ""},
		GLN: {"GLN", "GS1 global location number (formerly EAN location number).", "GLN (EAN-UCC)", "Global location number (tidligere: EAN location number)"},
		SAN: {"SAN", "Book trade Standard Address Number – US, UK etc.", "SAN", "Standard adressenummer i bokbransjen. Brukes i USA, UK etc."},
		CentraalBoekhuisRelatieID: {"Centraal Boekhuis Relatie ID", "Trading party identifier used in the Netherlands.", "Nederlandsk Centraal Boekhuis Relatie ID", "Trading party identifier used in the Netherlands"},
		FondscodeBoekenbank:       {"Fondscode Boekenbank", "Flemish publisher code.", "Fondscode Boekenbank", "Flamske utgiverkoder"},
		YTunnus:                   {"Y-tunnus", "Business Identity Code (Finland). See http://www.ytj.fi/ (in Finnish).", "Y-tunnus", "Business Identity Code (Finland). See http://www.ytj.fi/ (in Finnish)."},
		ISNI:                      {"ISNI", "International Standard Name Identifier. See ‘http://www.isni.org/’.", "ISNI", "International Standard Name Identifier. Se ‘http://www.isni.org/’."},
		PND:                       {"PND", "Personennamendatei – person name authority file used by Deutsche Nationalbibliothek and in other German-speaking countries. See http://www.d-nb.de/standardisierung/normdateien/pnd.htm (German) or http://www.d-nb.de/eng/standardisierung/normdateien/pnd.htm (English). DEPRECATED in favour of the GND.", "PND", "Personennamendatei – person name authority file used by Deutsche Nationalbibliothek and in other German-speaking countries. See http://www.d-nb.de/standardisierung/normdateien/pnd.htm (German) or http://www.d-nb.de/eng/standardisierung/normdateien/pnd.htm (English). DEPRECATED in favour of the GND."},
		LCCN:                      {"LCCN", "A control number assigned to a Library of Congress Name Authority record.", "LCCN", "A control number assigned to a Library of Congress Name Authority record."},
		JapanesePublisherIdentifier: {"Japanese Publisher identifier", "Publisher identifier administered by Japanese ISBN Agency.", "Japanese Publisher identifier", "Publisher identifier administered by Japanese ISBN Agency."},
		GKD:   {"GKD", "Gemeinsame Körperschaftsdatei – Corporate Body Authority File in the German-speaking countries. See http://www.d-nb.de/standardisierung/normdateien/gkd.htm (German) or http://www.d-nb.de/eng/standardisierung/normdateien/gkd.htm (English). DEPRECATED in favour of the GND.", "GKD", "Gemeinsame Körperschaftsdatei – Corporate Body Authority File in the German-speaking countries. See http://www.d-nb.de/standardisierung/normdateien/gkd.htm (German) or http://www.d-nb.de/eng/standardisierung/normdateien/gkd.htm (English). DEPRECATED in favour of the GND."},
		ORCID: {"ORCID", "Open Researcher and Contributor ID. See ‘http://www.orcid.org/’.", "ORCID", "Open Researcher and Contributor ID. See ‘http://www.orcid.org/’."},
		GAPPPublisherIdentifier:  {"GAPP Publisher Identifier", "Publisher identifier maintained by the Chinese ISBN Agency (GAPP).", "GAPP Publisher Identifier", "Publisher identifier maintained by the Chinese ISBN Agency (GAPP)."},
		VATIdentityNumber:        {"VAT Identity Number", "Identifier for a business organization for VAT purposes, eg within the EU’s VIES system. See http://ec.europa.eu/taxation_customs/vies/faqvies.do for EU VAT ID formats, which vary from country to country. Generally these consist of a two-letter country code followed by the 8–12 digits of the national VAT ID. Some countries include one or two letters within their VAT ID. See http://en.wikipedia.org/wiki/VAT_identification_number for non-EU countries that maintain similar identifiers. Spaces, dashes etc should be omitted.", "VAT Identity Number", "Identifier for a business organization for VAT purposes, eg within the EU’s VIES system. See http://ec.europa.eu/taxation_customs/vies/faqvies.do for EU VAT ID formats, which vary from country to country. Generally these consist of a two-letter country code followed by the 8–12 digits of the national VAT ID. Some countries include one or two letters within their VAT ID. See http://en.wikipedia.org/wiki/VAT_identification_number for non-EU countries that maintain similar identifiers. Spaces, dashes etc should be omitted."},
		JPDistributionIdentifier: {"JP Distribution Identifier", "4-digit business organization identifier controlled by the Japanese Publication Wholesalers Association.", "JP Distribution Identifier", "4-digit business organization identifier controlled by the Japanese Publication Wholesalers Association."},
		GND:                       {"GND", "Gemeinsame Normdatei – Joint Authority File in the German-speaking countries. See ‘http://www.dnb.de/EN/Standardisierung/Normdaten/GND/gnd_node.html (English)’. Combines the PND, SWD and GKD into a single authority file, and should be used in preference.", "GND", "Gemeinsame Normdatei – Joint Authority File in the German-speaking countries. See ‘http://www.dnb.de/EN/Standardisierung/Normdaten/GND/gnd_node.html (English)’. Combines the PND, SWD and GKD into a single authority file, and should be used in preference."},
		DUNS:                      {"DUNS", "Dunn and Bradstreet Universal Numbering System, see ‘www.dnb.co.uk/dandb-duns-number’.", "DUNS", "Dunn and Bradstreet Universal Numbering System, see ‘www.dnb.co.uk/dandb-duns-number’."},
		RinggoldID:                {"Ringgold ID", "Ringgold organizational identifier, see ‘http://www.ringgold.com/pages/identify.html’.", "Ringgold ID", "Ringgold organizational identifier, see ‘http://www.ringgold.com/pages/identify.html’."},
		IdentifiantEditeurElectre: {"Identifiant Editeur Electre", "French Electre publisher identifier.", "Identifiant Editeur Electre", "French Electre publisher identifier."},
		IdentifiantEditeurBTLF:    {"Identifiant Editeur BTLF", "Canadian BTLF publisher identifier.", "Identifiant Editeur BTLF", "Canadian BTLF publisher identifier."},
		IdentifiantMarqueElectre:  {"Identifiant Marque Electre", "French Electre imprint Identifier.", "Identifiant Marque Electre", "French Electre imprint Identifier."},
		IdentifiantMarqueBTLF:     {"Identifiant Marque BTLF", "Canadian BTLF imprint identifier.", "Identifiant Marque BTLF", "Canadian BTLF imprint identifier."},
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
