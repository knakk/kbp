package list58

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	RRPExcludingTax                             = "01"
	RRPIncludingTax                             = "02"
	FixedRetailPriceExcludingTax                = "03"
	FixedRetailPriceIncludingTax                = "04"
	SuppliersNetPriceExcludingTax               = "05"
	SuppliersNetPriceExcludingTaxRentalGoods    = "06"
	SuppliersNetPriceIncludingTax               = "07"
	SuppliersAlternativeNetPriceExcludingTax    = "08"
	SuppliersAlternativeNetPriceIncludingTax    = "09"
	SpecialSaleRRPExcludingTax                  = "11"
	SpecialSaleRRPIncludingTax                  = "12"
	SpecialSaleFixedRetailPriceExcludingTax     = "13"
	SpecialSaleFixedRetailPriceIncludingTax     = "14"
	SuppliersNetPriceForSpecialSaleExcludingTax = "15"
	SuppliersNetPriceForSpecialSaleIncludingTax = "17"
	PrePublicationRRPExcludingTax               = "21"
	PrePublicationRRPIncludingTax               = "22"
	PrePublicationFixedRetailPriceExcludingTax  = "23"
	PrePublicationFixedRetailPriceIncludingTax  = "24"
	SuppliersPrePublicationNetPriceExcludingTax = "25"
	SuppliersPrePublicationNetPriceIncludingTax = "27"
	FreightPassThroughRRPExcludingTax           = "31"
	FreightPassThroughBillingPriceExcludingTax  = "32"
	PublishersRetailPriceExcludingTax           = "41"
	PublishersRetailPriceIncludingTax           = "42"
)

var (
	sortedCodes = []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "11", "12", "13", "14", "15", "17", "21", "22", "23", "24", "25", "27", "31", "32", "41", "42"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		RRPExcludingTax:                             {"RRP excluding tax", "RRP excluding any sales tax or value-added tax.", "Foreslått utsalgsspris uten mva.", "RRP excluding any sales tax or value-added tax."},
		RRPIncludingTax:                             {"RRP including tax", "RRP including sales or value-added tax if applicable.", "Foreslått utsalgsspris inklusiv mva.", "RRP including sales or value-added tax if applicable."},
		FixedRetailPriceExcludingTax:                {"Fixed retail price excluding tax", "In countries where retail price maintenance applies by law to certain products: not used in USA.", "Fastpris uten mva.", "Brukes for å angi at boken er bundet av fastpris ihht. bransjeavtalen"},
		FixedRetailPriceIncludingTax:                {"Fixed retail price including tax", "In countries where retail price maintenance applies by law to certain products: not used in USA.", "Fastpris inklusiv mva.", "Brukes for å angi at boken er bundet av fastpris ihht. bransjeavtalen"},
		SuppliersNetPriceExcludingTax:               {"Supplier’s net price excluding tax", "Unit price charged by supplier to reseller excluding any sales tax or value-added tax: goods for retail sale.", "Distributørs enhetspris uten mva.: produkter for salg", "Unit price charged by supplier to reseller excluding any sales tax or value-added tax: goods for retail sale."},
		SuppliersNetPriceExcludingTaxRentalGoods:    {"Supplier’s net price excluding tax: rental goods", "Unit price charged by supplier to reseller / rental outlet, excluding any sales tax or value-added tax: goods for rental (used for video and DVD).", "Distributørs enhetspris uten mva.: produkter for utleie", "Brukes for video / DVD"},
		SuppliersNetPriceIncludingTax:               {"Supplier’s net price including tax", "Unit price charged by supplier to reseller including any sales tax or value-added tax if applicable: goods for retail sale.", "Distributørs enhetspris med mva.", "Produkter for salg"},
		SuppliersAlternativeNetPriceExcludingTax:    {"Supplier’s alternative net price excluding tax", "Unit price charged by supplier to a specified class of reseller excluding any sales tax or value-added tax: goods for retail sale (this value is for use only in countries, eg Finland, where trade practice requires two different net prices to be listed for different classes of resellers, and where national guidelines specify how the code should be used).", "Distributørs alternative enhetspris uten mva.", "Distributørs enhetspris for en spesifikk forhandler, ekskl. mva.: produkter for salg (denne koden brukes i land som Finland, hvor det er påkrevet å angi to enhetspriser for forskjellige typer av forhandlere, og hvor nasjonale retningslinjer angir hvordan koden skal brukes)"},
		SuppliersAlternativeNetPriceIncludingTax:    {"Supplier’s alternative net price including tax", "Unit price charged by supplier to a specified class of reseller including any sales tax or value-added tax: goods for retail sale (this value is for use only in countries, eg Finland, where trade practice requires two different net prices to be listed for different classes of resellers, and where national guidelines specify how the code should be used).", "Distributørs alternative enhetspris med mva.", "Distributørs enhetspris for en spesifikk forhandler, inkl. mva.: produkter for salg (denne koden brukes i land som Finland, hvor det er påkrevet å angi to enhetspriser for forskjellige typer av forhandlere, og hvor nasjonale retningslinjer angir hvordan koden skal brukes)"},
		SpecialSaleRRPExcludingTax:                  {"Special sale RRP excluding tax", "Special sale RRP excluding any sales tax or value-added tax. Note ‘special sales’ are sales where terms and conditions are different from normal trade sales, when for example products that are normally sold on a sale-or-return basis are sold on firm-sale terms, where a particular product is tailored for a specific retail outlet (often termed a ‘premium’\u00a0product), or where other specific conditions or qualiifications apply. Further details of the modified terms and conditions should be given in <PriceTypeDescription>.", "Foreslått spesialpris uten mva.", "Special sale RRP excluding any sales tax or value-added tax. Note ‘special sales’ are sales where terms and conditions are different from normal trade sales, when for example products that are normally sold on a sale-or-return basis are sold on firm-sale terms, where a particular product is tailored for a specific retail outlet (often termed a ‘premium’\u00a0product), or where other specific conditions or qualiifications apply. Further details of the modified terms and conditions should be given in <PriceTypeDescription>."},
		SpecialSaleRRPIncludingTax:                  {"Special sale RRP including tax", "Special sale RRP including sales or value-added tax if applicable.", "Foreslått spesialpris inklusiv mva.", "Special sale RRP including sales or value-added tax if applicable."},
		SpecialSaleFixedRetailPriceExcludingTax:     {"Special sale fixed retail price excluding tax", "In countries where retail price maintenance applies by law to certain products: not used in USA.", "Fast spesialpris uten mva.", "Brukes for angi en fastpris ihht. bokavtalen (-12.5%)"},
		SpecialSaleFixedRetailPriceIncludingTax:     {"Special sale fixed retail price including tax", "In countries where retail price maintenance applies by law to certain products: not used in USA.", "Fast spesialpris inklusiv mva.", "Brukes for angi en fastpris ihht. bokavtalen (-12.5%)"},
		SuppliersNetPriceForSpecialSaleExcludingTax: {"Supplier’s net price for special sale excluding tax", "Unit price charged by supplier to reseller for special sale excluding any sales tax or value-added tax.", "Distributørs spesialpris per enhet uten mva.", "Distributørs enhetspris til forhandler for spesielt salg, uten mva."},
		SuppliersNetPriceForSpecialSaleIncludingTax: {"Supplier’s net price for special sale including tax", "Unit price charged by supplier to reseller for special sale including any sales tax or value-added tax.", "Distributørs spesielle enhetspris inkl. mva.", "Distributørs enhetspris til forhandler for spesielt salg, inkl. mva."},
		PrePublicationRRPExcludingTax:               {"Pre-publication RRP excluding tax", "Pre-publication RRP excluding any sales tax or value-added tax.", "Foreløpig foreslått utsalgsspris uten mva.", "Pre-publication RRP excluding any sales tax or value-added tax."},
		PrePublicationRRPIncludingTax:               {"Pre-publication RRP including tax", "Pre-publication RRP including sales or value-added tax if applicable.", "Foreløpig foreslått utsalgsspris inklusiv mva.", "Pre-publication RRP including sales or value-added tax if applicable."},
		PrePublicationFixedRetailPriceExcludingTax:  {"Pre-publication fixed retail price excluding tax", "In countries where retail price maintenance applies by law to certain products: not used in USA.", "Foreløpig foreslått fastpris uten mva.", "Brukes før endelig fastpris er fastsatt"},
		PrePublicationFixedRetailPriceIncludingTax:  {"Pre-publication fixed retail price including tax", "In countries where retail price maintenance applies by law to certain products: not used in USA.", "Foreløpig foreslått fastpris inklusiv mva.", "Brukes før endelig fastpris er fastsatt"},
		SuppliersPrePublicationNetPriceExcludingTax: {"Supplier’s pre-publication net price excluding tax", "Unit price charged by supplier to reseller pre-publication excluding any sales tax or value-added tax.", "Distributørs foreløpige enhetspris uten mva.", "Distributørs enhetspris til forhandler før utgivelsen er et faktum, ekskl. mva."},
		SuppliersPrePublicationNetPriceIncludingTax: {"Supplier’s pre-publication net price including tax", "Unit price charged by supplier to reseller pre-publication including any sales tax or value-added tax.", "Distributørs foreløpige enhetspris inkl. mva.", "Distributørs enhetspris til forhandler før utgivelsen er et faktum, inkl. mva."},
		FreightPassThroughRRPExcludingTax:           {"Freight-pass-through RRP excluding tax", "In the US, books are sometimes supplied on ‘freight-pass-through’ terms, where a price that is different from the RRP is used as the basis for calculating the supplier’s charge to a reseller. To make it clear when such terms are being invoked, code 31 is used instead of code 01 to indicate the RRP. Code 32 is used for the ‘billing price’.", "Freight-pass-through foreslått utsalgsspris uten moms", "Brukes i USA"},
		FreightPassThroughBillingPriceExcludingTax:  {"Freight-pass-through billing price excluding tax", "When freight-pass-through terms apply, the price on which the supplier’s charge to a reseller is calculated, ie the price to which trade discount terms are applied. See also code 31.", "Freight-pass-through fakturert pris uten moms", "Brukes i USA"},
		PublishersRetailPriceExcludingTax:           {"Publishers retail price excluding tax", "For a product supplied on agency terms, the retail price set by the publisher, excluding any sales tax or value-added tax.", "Utgivers/vareeiers utsalgspris eks. mva.", "Utsalgsprisen angis av vareeier/utgiver, eks. mva."},
		PublishersRetailPriceIncludingTax:           {"Publishers retail price including tax", "For a product supplied on agency terms, the retail price set by the publisher, including sales or value-added tax if applicable.", "Utgivers/vareeiers utsalgspris inkl. mva.", "Utsalgsprisen angis av vareeier/utgiver, inkl. mva."},
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
