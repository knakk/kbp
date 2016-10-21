package list73

import (
	"fmt"

	"github.com/knakk/kbp/onix/codes"
)

const (
	UnspecifiedSeeWebsiteDescription                          = "00"
	PublishersCorporateWebsite                                = "01"
	PublishersWebsiteForASpecifiedWork                        = "02"
	OnlineHostingServiceHomePage                              = "03"
	JournalHomePage                                           = "04"
	OnlineResourceavailableContentPage                        = "05"
	ContributorsOwnWebsite                                    = "06"
	PublishersWebsiteRelatingToSpecifiedContributor           = "07"
	OtherPublishersWebsiteRelatingToSpecifiedContributor      = "08"
	ThirdPartyWebsiteRelatingToSpecifiedContributor           = "09"
	ContributorsOwnWebsiteForSpecifiedWork                    = "10"
	OtherPublishersWebsiteRelatingToSpecifiedWork             = "11"
	ThirdPartyWebsiteRelatingToSpecifiedWork                  = "12"
	ContributorsOwnWebsiteForGroupOrSeriesOfWorks             = "13"
	PublishersWebsiteRelatingToGroupOrSeriesOfWorks           = "14"
	OtherPublishersWebsiteRelatingToGroupOrSeriesOfWorks      = "15"
	ThirdPartyWebsiteRelatingToGroupOrSeriesOfWorksEgAFanSite = "16"
	PublishersB2BWebsite                                      = "17"
	PublishersB2CWebsite                                      = "18"
	AuthorBlog                                                = "23"
	WebPageForAuthorPresentationCommentary                    = "24"
	WebPageForAuthorInterview                                 = "25"
	WebPageForAuthorReading                                   = "26"
	WebPageForCoverMaterial                                   = "27"
	WebPageForSampleContent                                   = "28"
	WebPageForFullContent                                     = "29"
	WebPageForOtherCommentaryDiscussion                       = "30"
	TransferURL                                               = "31"
	DOIWebsiteLink                                            = "32"
	SuppliersCorporateWebsite                                 = "33"
	SuppliersB2BWebsite                                       = "34"
	SuppliersB2CWebsite                                       = "35"
	SuppliersWebsiteForASpecifiedWork                         = "36"
	SuppliersB2BWebsiteForASpecifiedWork                      = "37"
	SuppliersB2CWebsiteForASpecifiedWork                      = "38"
	SuppliersWebsiteForAGroupOrSeriesOfWorks                  = "39"
	URLOfFullMetadataDescription                              = "40"
	SocialNetworkingURLForSpecificWorkOrProduct               = "41"
	AuthorsSocialNetworkingURL                                = "42"
	PublishersSocialNetworkingURL                             = "43"
	SocialNetworkingURLForSpecificArticleChapterOrContentItem = "44"
)

var (
	sortedCodes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44"}

	all = map[string]struct{ labelEn, notesEn, labelNo, notesNo string }{
		UnspecifiedSeeWebsiteDescription:                          {"Unspecified, see website description", "", "Uspesifisert", ""},
		PublishersCorporateWebsite:                                {"Publisher’s corporate website", "See also codes 17 and 18.", "Forlagets hjemmeside", "Se også kode 17 og 18"},
		PublishersWebsiteForASpecifiedWork:                        {"Publisher’s website for a specified work", "A publisher’s informative and/or promotional webpage relating to a specified work (book, journal, online resource or other publication type).", "Forlagets nettsted for et bestemt produkt", "Utgivers informative eller markedsførende websider relatert til et spesifikt arbeid. (bok, tidsskrift, online ressurser eller andre publikasjonsformer)"},
		OnlineHostingServiceHomePage:                              {"Online hosting service home page", "A webpage giving access to an online content hosting service as a whole.", "Nettsted for webhotell", "En webside som gir tilgang til en webhotell for online innhold"},
		JournalHomePage:                                           {"Journal home page", "A webpage giving general information about a serial, in print or electronic format or both.", "Nettsted for tidsskrift", "En webside som gir generell informasjon om et tidsskrift (trykket, online eller begge)"},
		OnlineResourceavailableContentPage:                        {"Online resource ‘available content’ page", "A webpage giving direct access to the content that is available online for a specified resource version. Generally used for content available online under subscription terms.", "Nettsted for tilgjengelig innhold i tidsskrift", "En webside som gir direkte tilgang til innhold som er tilgjengelig online for et gitt tidsskift "},
		ContributorsOwnWebsite:                                    {"Contributor’s own website", "A webpage maintained by an author or other contributor about her/his publications and personal background.", "Forfatters/bidragsyters eget nettsted", "En webside vedlikeholdt av forfatteren eller annen bidragsyter om han/hennes publikasjoner og personlige bakgrunn "},
		PublishersWebsiteRelatingToSpecifiedContributor:           {"Publisher’s website relating to specified contributor", "A publisher’s webpage devoted to a specific author or other contributor.", "Forlagets nettsted for en forfatter/bidragsyter", "En utgivers webside som omhandler en gitt forfatter eller annen bidragsyter"},
		OtherPublishersWebsiteRelatingToSpecifiedContributor:      {"Other publisher’s website relating to specified contributor", "A webpage devoted to a specific author or other contributor, and maintained by a publisher other than the publisher of the item described in the ONIX record.", "Annets forlags nettsted for en bidragsyter", "En webside om en gitt forfatter eller bidragsyter. Vedlikeholdes av annen utgiver enn den som er beskrevet i ONIX-posten"},
		ThirdPartyWebsiteRelatingToSpecifiedContributor:           {"Third-party website relating to specified contributor", "A webpage devoted to a specific author or other contributor, and maintained by a third party (eg a fan site).", "Annet nettsted for en bidragsyter", "En webside om en gitt forfatter eller bidragsyter. Vedlikeholdes av tredjepart (f.eks. en fanklubb)"},
		ContributorsOwnWebsiteForSpecifiedWork:                    {"Contributor’s own website for specified work", "A webpage maintained by an author or other contributor and specific to an individual work.", "Forfatters/bidragsyters eget nettsted for et bestemt produkt", "En webside vedlikeholdt av forfatteren eller annen bidragsyter spesielt for et bestemt arbeid"},
		OtherPublishersWebsiteRelatingToSpecifiedWork:             {"Other publisher’s website relating to specified work", "A webpage devoted to an individual work, and maintained by a publisher other than the publisher of the item described in the ONIX record.", "Annet forlags nettsted for et bestemt produkt", "En webside dedikert til et bestemt arbeid. Vedlikeholdes av annen utgiver enn den som er beskrevet i ONIX-posten"},
		ThirdPartyWebsiteRelatingToSpecifiedWork:                  {"Third-party website relating to specified work", "A webpage devoted to an individual work, and maintained by a third party (eg a fan site).", "Annet nettsted for et bestemt produkt", "En webside dedikert til et bestemt arbeid. Vedlikeholdes av tredjepart (f.eks. en fanklubb)"},
		ContributorsOwnWebsiteForGroupOrSeriesOfWorks:             {"Contributor’s own website for group or series of works", "A webpage maintained by an author or other contributor and specific to a group or series of works.", "Forfatters/bidragsyters eget nettsted for en gruppe eller en serie produkter", "En webside vedlikeholdt av forfatteren eller annen bidragsyter om en gruppe eller serie med arbeider"},
		PublishersWebsiteRelatingToGroupOrSeriesOfWorks:           {"Publisher’s website relating to group or series of works", "A publisher’s webpage devoted to a group or series of works.", "Forlagets eget nettsted for en gruppe eller en serie produkter", "En utgivers websider om en gruppe eller serie med arbeider"},
		OtherPublishersWebsiteRelatingToGroupOrSeriesOfWorks:      {"Other publisher’s website relating to group or series of works", "A webpage devoted to a group or series of works, and maintained by a publisher other than the publisher of the item described in the ONIX record.", "Annet forlags nettsted for en gruppe eller en serie produkter", "En webside om en gruppe eller serie med arbeider. Vedlikeholdes av annen utgiver enn den som er beskrevet i ONIX-posten"},
		ThirdPartyWebsiteRelatingToGroupOrSeriesOfWorksEgAFanSite: {"Third-party website relating to group or series of works (eg a fan site)", "A webpage devoted to a group or series of works, and maintained by a third party (eg a fan site).", "Annet nettsted for en gruppe eller en serie produkter", "En webside om en gruppe eller serie med arbeider.Vedlikeholdes av tredjepart (f.eks. en fanklubb)"},
		PublishersB2BWebsite:                                      {"Publisher’s B2B website", "Use instead of code 01 to specify a publisher’s website for trade users.", "Forlagets nettsted for bedriftsmarkedet", "Bruk istedet for kode 01 for å spesifisere en utgivers webside for handelspartnere"},
		PublishersB2CWebsite:                                      {"Publisher’s B2C website", "Use instead of code 01 to specify a publisher’s website for end customers (consumers).", "Forlagets nettsted for konsument", "Bruk istedet for kode 01 for å spesifisere en utgivers webside for sluttkunder"},
		AuthorBlog:                                                {"Author blog", "For example, a Blogger or Tumblr URL, a Wordpress website or other blog URL.", "Forfatterblogg", "For example, a Blogger or Tumblr URL, a Wordpress website or other blog URL."},
		WebPageForAuthorPresentationCommentary:                    {"Web page for author presentation / commentary", "", "Nettside for presentasjon av forfatteren", ""},
		WebPageForAuthorInterview:                                 {"Web page for author interview", "", "Nettside for forfatterintervju", ""},
		WebPageForAuthorReading:                                   {"Web page for author reading", "", "Nettside for forfatteropplesning", ""},
		WebPageForCoverMaterial:                                   {"Web page for cover material", "", "Nettside for omslagsmateriale", ""},
		WebPageForSampleContent:                                   {"Web page for sample content", "", "Nettside for eksempel på innhold", ""},
		WebPageForFullContent:                                     {"Web page for full content", "Use this value in the <Website> composite in <SupplyDetail> when sending a link to a webpage at which a digital product is available for download and/or online access.", "Nettside for komplett innhold", "Bruk denne koden i <Website> i <SupplyDetail> når en lenke til nettstedet for et digitalt produkt som er tilgjengelig for nedlasting og/eller online sendes."},
		WebPageForOtherCommentaryDiscussion:                       {"Web page for other commentary / discussion", "", "Nettside for andre kommentarer / diskusjoner", ""},
		TransferURL:                                               {"Transfer-URL", "URL needed by the German National Library for direct access, harvesting and storage of an electronic resource.", "Transfer-URL", "Kun for bruk i Tyskland"},
		DOIWebsiteLink:                                            {"DOI Website Link", "Link needed by German Books in Print (VLB) for DOI registration and ONIX DOI conversion.", "DOI Website Link", "Kun for bruk i Tyskland"},
		SuppliersCorporateWebsite:                                 {"Supplier’s corporate website", "A corporate website operated by a distributor or other supplier (not the publisher).", "Distributørens nettsted", "Nettsted som vedlikeholdes av en distributør (ikke forlaget), ment for handelspartnere"},
		SuppliersB2BWebsite:                                       {"Supplier’s B2B website", "A website operated by a distributor or other supplier (not the publisher) and aimed at trade customers.", "Distributørens nettsted for bedriftsmarkedet", "Nettsted vedlikeholdt av en distributør (ikke forlaget), ment for handelspartnere"},
		SuppliersB2CWebsite:                                       {"Supplier’s B2C website", "A website operated by a distributor or other supplier (not the publisher) and aimed at consumers.", "Distributørens nettsted for konsument", "Nettsted vedlikeholdt av en distributør (ikke forlaget), for sluttkunder"},
		SuppliersWebsiteForASpecifiedWork:                         {"Supplier’s website for a specified work", "A distributor or supplier’s webpage describing a specified work.", "Distributørens nettsted for et bestemt produkt", "Distributørens nettsted som beskriver et bestemt produkt"},
		SuppliersB2BWebsiteForASpecifiedWork:                      {"Supplier’s B2B website for a specified work", "A distributor or supplier’s webpage describing a specified work, and aimed at trade customers.", "Distributørens nettsted for bedriftsmarkedet for et bestemt produkt", "Distributørens nettsted for et bestemt produkt, ment for bedriftsmarkedet"},
		SuppliersB2CWebsiteForASpecifiedWork:                      {"Supplier’s B2C website for a specified work", "A distributor or supplier’s webpage describing a specified work, and aimed at consumers.", "Distributørens nettsted for konsument for et bestemt produkt", "Distributørens nettsted for et bestemt produkt, ment for sluttkunder"},
		SuppliersWebsiteForAGroupOrSeriesOfWorks:                  {"Supplier’s website for a group or series of works", "A distributor or supplier’s webpage describing a group or series of works.", "Distributørens nettsted for en gruppe eller en serie produkter", "Distributørens nettsted for en gruppe eller en serie produkter"},
		URLOfFullMetadataDescription:                              {"URL of full metadata description", "For example an ONIX or MARC record for the product, available online.", "URL til full metadatapost", "F.eks Onix- eller Marc-poster for produktet, tilgjengelig online"},
		SocialNetworkingURLForSpecificWorkOrProduct:               {"Social networking URL for specific work or product", "For example, a Facebook, Google+ or Twitter URL for the product or work.", "URL til sosiale medier for et bestemt produkt", "F.eks. Facebook, Google+ eller Twitter URL for produktet"},
		AuthorsSocialNetworkingURL:                                {"Author’s social networking URL", "For example, a Facebook, Google+ or Twitter page.", "URL til forfatterens sider på sosiale medier", "F.eks. en Facebook-, Google+ eller Twitter-side"},
		PublishersSocialNetworkingURL:                             {"Publisher’s social networking URL", "For example, a Facebook, Google+ or Twitter page.", "URL til utgivers/vareeiers sider på sosiale medier", "F.eks. en Facebook-, Google+ eller Twitter-side"},
		SocialNetworkingURLForSpecificArticleChapterOrContentItem: {"Social networking URL for specific article, chapter or content item", "For example, a Facebook, Google+ or Twitter page. Use only in the context of a specific content item (eg within <ContentItem>).", "URL til sider på sosiale medier for en bestemt artikkel, et kapittel eller et annet innholdselement", "F.eks. en Facebook-, Google+ eller Twitter-side"},
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
