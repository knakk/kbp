package list79

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	ColorOfCover                                           = "01"
	ColorOfPageEdge                                        = "02"
	TextFont                                               = "03"
	SpecialCoverMaterial                                   = "04"
	DVDRegion                                              = "05"
	OperatingSystemRequirements                            = "06"
	OtherSystemRequirements                                = "07"
	PointAndListenDeviceCompatibility                      = "08"
	EPublicationAccessibilityDetail                        = "09"
	EPublicationFormatVersion                              = "10"
	CPSIAChokingHazardWarning                              = "11"
	CPSIAChokingHazardWarning                              = "12"
	EUToySafetyHazardWarning                               = "13"
	IATADangerousGoodsWarning                              = "14"
	EPublicationFormatVersionCode                          = "15"
	NotFSCOrPEFCCertified                                  = "30"
	FSCCertifiedPure                                       = "31"
	FSCCertifiedMixedSources                               = "32"
	FSCCertifiedRecycled                                   = "33"
	PEFCCertified                                          = "34"
	PEFCRecycled                                           = "35"
	FSCOrPEFCCertifiedPreAndPostConsumerWastePCWPercentage = "36"
	ClaimedPreAndPostConsumerWastePCWPercentage            = "37"
	PaperProducedBygreenTechnology                         = "40"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "30", "31", "32", "33", "34", "35", "36", "37", "40"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		ColorOfCover:                      {"Color of cover", "For Product Form Feature values see code list 98.", "Innbindingsfarge", "For verdier, se kodeliste 98"},
		ColorOfPageEdge:                   {"Color of page edge", "For Product Form Feature values see code list 98.", "Papirkant farge", "For verdier, se kodeliste 98"},
		TextFont:                          {"Text font", "The principal font used for body text, when this is a significant aspect of product description, eg for some Bibles, and for large print product. The accompanying Product Form Feature Description is text specifying font size and, if desired, typeface.", "Typesnitt", "Hovedtekstfont og/eller størrelse brukt i hovedteksten når dette er vesentlig for produktbeskrivelsen, f.eks. for visse bibler. Bruk tekst ProductFormFeatureDescription til å spesifisere"},
		SpecialCoverMaterial:              {"Special cover material", "For Product Form Feature values see code list 99.", "Spesielt innbindingsmateriale", "For verdier, se kodeliste 99"},
		DVDRegion:                         {"DVD region", "For Product Form Feature values see code list 76.", "DVD regioner", "For verdier, se kodeliste 76"},
		OperatingSystemRequirements:       {"Operating system requirements", "A computer or handheld device operating system required to use a digital product, with version detail if applicable. The accompanying Product Form Feature Value is a code from List 176. Version detail, when applicable, is carried in Product Form Feature Description.", "Krav til operativsystem", "Krav til operativsystem for datamaskin eller håndholdt enhet for å kunne bruke et digitalt produkt, med informasjon om versjon der det er aktuelt. Product Form Feature Value inneholder koder fra liste 176. Versjonsdetaljer sendes i Product form feature description der det er aktuelt."},
		OtherSystemRequirements:           {"Other system requirements", "Other system requirements for a digital product, described by free text in Product Form Feature Description.", "Andre systemkrav", "Andre systemkrav for digitale produkter, sendt som fritekst i Product Form Feature Description."},
		PointAndListenDeviceCompatibility: {"‘Point and listen’ device compatibility", "Indicates compatibility with proprietary ‘point and listen’ devices such as Ting Pen (http://www.ting.eu) or the iSmart Touch and Read Pen. These devices scan invisible codes specially printed on the page to identify the book and position of the word, and the word is then read aloud by the device. The name of the compatible device (or range of devices) should be given in <ProductFormFeatureDescription>.", "Kompabilitet til 'Pek og hør'-enheter", "Indikerer komatibilitet med proprietære 'pek og hør'-enheter som Ting Pen (http://www.ting.eu) eller iSmart Touch og Read Pen. Disse enhetene skanner usynlige koder trykt på siden for å identifisere boka og ordets posisjon, og ordet blir så lest høyt av enheten. Navnet på den/de enheten(e) produktet er kompatibelt med bør angis i <ProductFormFeatureDescription>."},
		EPublicationAccessibilityDetail:   {"E-publication accessibility detail", "For <ProductFormFeatureValue> codes, see Codelist 196.", "E-publikasjoner, tilgjengelighetsdetaljer", "For <ProductFormFeatureValue>-koder, se liste 196."},
		EPublicationFormatVersion:         {"E-publication format version", "For versioned e-book file formats (or in some cases, devices). <ProductFormFeatureValue> should contain the version number as a period-separated list of numbers (eg ‘7’, ‘1.5’ or ‘3.10.7’). Use only with ONIX 3.0 – in ONIX 2.1, use <EpubTypeVersion> instead. For the most common file formats, code 15 and List 220 is strongly preferred.", "E-pubikasjoner, formatversjon", "For versjonerte e-bok-filformater (eller i noen tilfeller, enheter). <ProductFormFeatureValue> bør inneholde versjonsnummeret (f.eks. ‘7’, ‘1.5’ or ‘3.10.7’). Brukes kun med ONIX 3.0 – i ONIX 2.1 brukes <EpubTypeVersion> i steden. For de vanligste filformatene anbefales bruk av kode 15 og liste 220"},
		CPSIAChokingHazardWarning:         {"CPSIA choking hazard warning", "DEPRECATED – use code 12 and List 143.", "CPSIA choking hazard warning", "UTGÅTT – bruk kode 12 og liste 143."},
		CPSIAChokingHazardWarning:         {"CPSIA choking hazard warning", "Choking hazard warning required by US Consumer Product Safety Improvement Act (CPSIA) of 2008. Required, when applicable, for products sold in the US. The Product Form Feature Value is a code from List 143. Further explanation may be given in Product Form Feature Description.", "CPSIA choking hazard warning", "Choking hazard warning required by US Consumer Product Safety Improvement Act (CPSIA) of 2008. Required, when applicable, for products sold in the US. The Product Form Feature Value is a code from List 143. Further explanation may be given in Product Form Feature Description."},
		EUToySafetyHazardWarning:          {"EU Toy Safety Hazard warning", "Product carries hazard warning required by EU Toy Safety Directive. The Product Form Feature Value is a code from List 184, and (for some codes) the exact wording of the warning may be given in Product Form Feature Description.", "EU Toy Safety Hazard warning", "Product carries hazard warning required by EU Toy Safety Directive. The Product Form Feature Value is a code from List 184, and (for some codes) the exact wording of the warning may be given in Product Form Feature Description."},
		IATADangerousGoodsWarning:         {"IATA Dangerous Goods warning", "Product Form Feature Description must give further details of the warning.", "IATA Dangerous Goods warning", "Product Form Feature Description must give further details of the warning."},
		EPublicationFormatVersionCode:     {"E-publication format version code", "For common versioned e-book formats (or in some cases, devices) – for example EPUB 2.0.1 or EPUB 3.0. <ProductFormFeatureValue> is a code from list 220. Use in ONIX 3.0 only.", "E-publikasjoner, kode for formatversjon", "For e-bokformater som fins i flere versjoner (eller i noen tilfeller, enheter) – for eksempel EPUB 2.0.1 eller EPUB 3.0. <ProductFormFeatureValue> inneholder en kode fra liste 220. Use in ONIX 3.0 only."},
		NotFSCOrPEFCCertified:             {"Not FSC or PEFC certified", "Product does not carry FSC or PEFC logo. The Product Form Feature Value and Description elements are not used. The product may, however, still carry a claimed Pre- and Post-Consumer Waste (PCW) content (type code 37) in a separate repeat of the Product Form Feature composite.", "Not FSC or PEFC certified", "Product does not carry FSC or PEFC logo. The Product Form Feature Value and Description elements are not used. The product may, however, still carry a claimed Pre- and Post-Consumer Waste (PCW) content (type code 37) in a separate repeat of the Product Form Feature composite."},
		FSCCertifiedPure:                  {"FSC certified – pure", "Product carries FSC logo (Pure, 100%). <ProductFormFeatureValue> is the Certification number (ie either a Chain Of Custody (COC) number or a Trademark License number) printed on the book. Format: Chain of Custody number is two to five letters-COC-six digits (the digits should include leading zeros if necessary), eg “AB-COC-001234” or “ABCDE-COC-123456”; Trademark License number is C followed by six digits, eg “C005678” (this would normally be prefixed by ‘FSC®’ when displayed). By definition, a product certified Pure does not contain Pre- and Post-Consumer-Waste (PCW), so type code 31 can only occur on its own. Certification numbers may be checked at ‘http://info.fsc.org/’.", "FSC certified – pure", "Product carries FSC logo (Pure, 100%). <ProductFormFeatureValue> is the Certification number (ie either a Chain Of Custody (COC) number or a Trademark License number) printed on the book. Format: Chain of Custody number is two to five letters-COC-six digits (the digits should include leading zeros if necessary), eg “AB-COC-001234” or “ABCDE-COC-123456”; Trademark License number is C followed by six digits, eg “C005678” (this would normally be prefixed by ‘FSC®’ when displayed). By definition, a product certified Pure does not contain Pre- and Post-Consumer-Waste (PCW), so type code 31 can only occur on its own. Certification numbers may be checked at ‘http://info.fsc.org/’."},
		FSCCertifiedMixedSources:          {"FSC certified – mixed sources", "Product carries FSC logo (Mixed sources, Mix). <ProductFormFeatureValue> is the Certification number (ie either a Chain Of Custody (COC) number or a Trademark License number) printed on the book. Format: Chain of Custody number is two to five letters-COC-six digits (the digits should include leading zeros if necessary), eg “AB-COC-001234” or “ABCDE-COC-123456”; Trademark License number is C followed by six digits, eg “C005678” (this would normally be prefixed by ‘FSC®’ when displayed). May be accompanied by a Pre- and Post-Consumer-Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36. Certification numbers may be checked at http://info.fsc.org/", "FSC certified – mixed sources", "Product carries FSC logo (Mixed sources, Mix). <ProductFormFeatureValue> is the Certification number (ie either a Chain Of Custody (COC) number or a Trademark License number) printed on the book. Format: Chain of Custody number is two to five letters-COC-six digits (the digits should include leading zeros if necessary), eg “AB-COC-001234” or “ABCDE-COC-123456”; Trademark License number is C followed by six digits, eg “C005678” (this would normally be prefixed by ‘FSC®’ when displayed). May be accompanied by a Pre- and Post-Consumer-Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36. Certification numbers may be checked at http://info.fsc.org/"},
		FSCCertifiedRecycled:              {"FSC certified – recycled", "Product carries FSC logo (Recycled). <ProductFormFeatureValue> is the Certification number (ie either a Chain Of Custody (COC) number or a Trademark License number) printed on the book. Format: Chain of Custody number is two to five letters-COC-six digits (the digits should include leading zeroes if necessary), eg “AB-COC-001234” or “ABCDE-COC-123456”; Trademark License number is C followed by six digits, eg “C005678” (this would normally be prefixed by ‘FSC®’ when displayed). Should be accompanied by a Pre- and Post-Consumer-Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36. Certification numbers may be checked at‘ http://info.fsc.org/’.", "FSC certified – recycled", "Product carries FSC logo (Recycled). <ProductFormFeatureValue> is the Certification number (ie either a Chain Of Custody (COC) number or a Trademark License number) printed on the book. Format: Chain of Custody number is two to five letters-COC-six digits (the digits should include leading zeroes if necessary), eg “AB-COC-001234” or “ABCDE-COC-123456”; Trademark License number is C followed by six digits, eg “C005678” (this would normally be prefixed by ‘FSC®’ when displayed). Should be accompanied by a Pre- and Post-Consumer-Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36. Certification numbers may be checked at‘ http://info.fsc.org/’."},
		PEFCCertified:                     {"PEFC certified", "Product carries PEFC logo (certified). <ProductFormFeatureValue> is the Chain Of Custody (COC) number printed on the book. May be accompanied by a Post-Consumer Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36.", "PEFC certified", "Product carries PEFC logo (certified). <ProductFormFeatureValue> is the Chain Of Custody (COC) number printed on the book. May be accompanied by a Post-Consumer Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36."},
		PEFCRecycled:                      {"PEFC recycled", "Product carries PEFC logo (recycled). <ProductFormFeatureValue> is the Chain Of Custody (COC) number printed on the book. Should be accompanied by a Post-Consumer-Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36.", "PEFC recycled", "Product carries PEFC logo (recycled). <ProductFormFeatureValue> is the Chain Of Custody (COC) number printed on the book. Should be accompanied by a Post-Consumer-Waste (PCW) percentage value, to be reported in another instance of <ProductFormFeature> with type code 36."},
		FSCOrPEFCCertifiedPreAndPostConsumerWastePCWPercentage: {"FSC or PEFC certified Pre- and Post-Consumer Waste (PCW) percentage", "The percentage of recycled Pre- and Post-Consumer-Waste (PCW) used in a product where the composition is certified by FSC or PEFC. <ProductFormFeatureValue> is an integer. May occur together with type code 32, 33, 34 or 35.", "FSC or PEFC certified Pre- and Post-Consumer Waste (PCW) percentage", "The percentage of recycled Pre- and Post-Consumer-Waste (PCW) used in a product where the composition is certified by FSC or PEFC. <ProductFormFeatureValue> is an integer. May occur together with type code 32, 33, 34 or 35."},
		ClaimedPreAndPostConsumerWastePCWPercentage:            {"Claimed Pre- and Post-Consumer Waste (PCW) percentage", "The percentage of recycled Pre- and Post-Consumer Waste (PCW) claimed to be used in a product where the composition is not certified by FSC or PEFC. <Product FormFeatureValue> is an integer. <ProductFormFeatureDescription> may carry free text supporting the claim. Must be accompanied by type code 30.", "Claimed Pre- and Post-Consumer Waste (PCW) percentage", "The percentage of recycled Pre- and Post-Consumer Waste (PCW) claimed to be used in a product where the composition is not certified by FSC or PEFC. <Product FormFeatureValue> is an integer. <ProductFormFeatureDescription> may carry free text supporting the claim. Must be accompanied by type code 30."},
		PaperProducedBygreenTechnology:                         {"Paper produced by ‘green’ technology", "Product made from paper produced using environmentally-conscious technology. <ProductFormFeatureDescription> may carry free text with a more detailed statement.", "Papir produsert ved bruk av 'grønn' teknologi", "Produktet er lagd av papir som er produsert ved å bruke miljøvennlig teknologi. <ProductFormFeatureDescription> kan inneholde fritekst med mer detaljert informasjon"},
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
