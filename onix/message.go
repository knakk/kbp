// Package onix provides data structures and means for working with ONIX
// for books, including encoding and decoding of records.
//
// The package is based on ONIX 3.03, as specified in the XSD schema
// available here: http://www.editeur.org/93/Release-3.0-Downloads/#Specification
package onix

import "encoding/xml"

// Addressee represents agroup of data elements which together specify the
// addressee of an ONIX for Books message.  Optional, and repeatable if
// there are several addressees.
type Addressee struct {
	ContactName         *ContactName
	EmailAddress        *EmailAddress
	AddresseeIdentifier []AddresseeIdentifier
	AddresseeName       *AddresseeName
	GeneralAttributes
}

// AddresseeIdentifier represents a group of data elements which together
// define an identifier of the addressee. The composite is optional, and
// repeatable if more than one identifier of different types is sent; but
// either an <AddresseeName> or an <AddresseeIdentifier> must be included.
type AddresseeIdentifier struct {
	AddresseeIDType AddresseeIDType
	IDTypeName      *IDTypeName
	IDValue         IDValue
	GeneralAttributes
}

// AddresseeIDType represents an ONIX code identifying a scheme from which
// an identifier in the <IDValue> element is taken. Mandatory in each
// occurrence of the <AddresseeIdentifier> composite, and non-repeating.
type AddresseeIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// AddresseeName represents the name of the addressee organization, which
// should always be stated in a standard form agreed with the addressee.
// Optional and non-repeating; but either a <AddresseeName> element or a
// <AddresseeIdentifier> composite must be included.
type AddresseeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// Affiliation represents an organization to which a contributor to the
// product was affiliated at the time of its creation, and – if the
// <ProfessionalPosition> element is also present – where s/he held that
// position. Optional and non-repeating.
type Affiliation struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// AgentIdentifier represents a group of data elements together defining
// the identifier of an agent or local publisher in accordance with a
// specified scheme. Optional, but each occurrence of the
// <PublisherRepresentative> composite must carry either at least one agent
// identifier, or an <AgentName>. Repeatable only if two or more identifiers
// are sent using different schemes.
type AgentIdentifier struct {
	AgentIDType AgentIDType
	IDTypeName  *IDTypeName
	IDValue     IDValue
	GeneralAttributes
}

// AgentIDType represents an ONIX code identifying the scheme from which the
// identifier in the <IDValue> element is taken. Mandatory in each occurrence
// of the <AgentIdentifier> composite, and non-repeating.
type AgentIDType struct {
	Value string `xml:",innerxml"` // List92
	GeneralAttributes
}

// AgentName represents the name of an agent or local publisher. Optional and
// non-repeating; required if no agent identifier is sent in an occurrence of
// the <PublisherRepresentative> composite.
type AgentName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// AgentRole represents n ONIX code identifying the role of an agent in
// relation to the product in the specified market, eg Exclusive sales
// agent, Local publisher, etc. Mandatory in each occurrence of the
// <PublisherRepresentative>
type AgentRole struct {
	Value string `xml:",innerxml"` // List69
	GeneralAttributes
}

// AlternativeName represents a repeatable group of data elements which
// together represent an alternative name of a contributor, and specify its
// type. The <AlternativeName> composite is optional. It may be used to send
// a pseudonym as well as a real name, where both names are on the product,
// eg to handle such cases as ‘Ian Rankin writing as Jack Harvey’; or to
// send an authority-controlled form of a name; or to identify the real name
// of the contributor where the book is written under a pseudonym (and the
// real identity need not be kept private). Note that in all cases, the primary
// name is that used on the product, and the alternative name merely provides
// a dditional information. Each instance of the composite must contain the
// <NameType> element with either:
//
// * one or more of the forms of representation of a person name, with or without an occurrence of the <NameIdentifier> composite; or
//
// * one or both of the forms of representation of a corporate name, with  or without an occurrence of the <NameIdentifier> composite; or
//
// * an occurrence of the <NameIdentifier> composite without any accompanying name element(s).
type AlternativeName struct {
	NameType           NameType
	NameIdentifier     []NameIdentifier
	PersonName         PersonName
	PersonNameInverted *PersonNameInverted
	TitlesBeforeNames  *TitlesBeforeNames
	NamesBeforeKey     *NamesBeforeKey
	PrefixToKey        *PrefixToKey
	KeyNames           KeyNames
	NamesAfterKey      *NamesAfterKey
	SuffixToKey        *SuffixToKey
	LettersAfterNames  *LettersAfterNames
	TitlesAfterNames   *TitlesAfterNames
	Gender             *Gender
	GeneralAttributes
}

// AncillaryContent represents a eepeatable group of data elements which
// together specify the number of illustrations or other content items
// of a stated type which the product carries. Use of the <AncillaryContent>
// composite is optional.
type AncillaryContent struct {
	AncillaryContentType        AncillaryContentType
	AncillaryContentDescription []AncillaryContentDescription
	Number                      *Number
	GeneralAttributes
}

// AncillaryContentType represents an ONIX code which identifies the
// type of illustration or other content to which an occurrence of the
// composite refers. Mandatory in each occurrence of the <AncillaryContent>
// composite, and non-repeating.
type AncillaryContentType struct {
	Value string `xml:",innerxml"` // List25
	GeneralAttributes
}

// AncillaryContentDescription represents a text describing the type of
// illustration or other content to which an occurrence of the composite
// refers, when a code is insufficient. Optional, and repeatable if parallel
// descriptive text is provided in multiple languages. Required when
// <AncillaryContentType> carries the value 00. The language attribute
// is optional for a single instance of <AncillaryContentDescription>,
// but must be included in each instance if <AncillaryContentDescription>
// is repeated.
type AncillaryContentDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// Audience represents a repeatable group of data elements which together
// describe an audience to which the product is directed.
type Audience struct {
	AudienceCodeType     AudienceCodeType
	AudienceCodeTypeName *AudienceCodeTypeName
	AudienceCodeValue    AudienceCodeValue
	GeneralAttributes
}

// AudienceCode [DEPRECATED] represents an ONIX code, derived from BISAC
// and BIC lists, which identifies the broad audience or readership for
// which a product is intended. Optional, and repeatable if the product
// is intended for two or more groups. Deprecated, in favor of providing
// the same information within the <Audience> composite using code 01
type AudienceCode struct {
	Value string `xml:",innerxml"` // List28
	GeneralAttributes
}

// AudienceCodeType represents an ONIX code which identifies the scheme
// from which the code in <AudienceCodeValue> is taken. Mandatory in
// each occurrence of the <Audience> composite, and non-repeating.
type AudienceCodeType struct {
	Value string `xml:",innerxml"` // List29
	GeneralAttributes
}

// AudienceCodeTypeName represents a name which identifies a proprietary
// audience code when the code in <AudienceCodeType> indicates a proprietary
// scheme, eg a vendor’s own code. Optional and non-repeating.
type AudienceCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// AudienceCodeValue represents a code value taken from the scheme
// specified in <AudienceCodeType>. Mandatory in each occurrence of
// the <Audience> composite, and non-repeating.
type AudienceCodeValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// AudienceDescription represents a free text describing the audience
// for which a product is intended. Optional, and repeatable if parallel
// descriptive text is provided in multiple languages. The language
// attribute is optional for a single instance of <AudienceDescription>,
// but must be included in each instance if <AudienceDescription> is repeated.
type AudienceDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// AudienceRange represents an optional and repeatable group of data
// elements which together describe an audience or readership range
// for which a product is intended. The composite can carry a single
// value from, to, or exact, or a pair of values with an explicit
// from and to.
type AudienceRange struct {
	AudienceRangeQualifier AudienceRangeQualifier
	AudienceRangePrecision AudienceRangePrecision
	AudienceRangeValue     AudienceRangeValue
	GeneralAttributes
}

// AudienceRangePrecision represents an ONIX code specifying the
// ‘precision’ of the value in the <AudienceRangeValue> element which
// follows (from, to, exact). Mandatory in each occurrence of the
// <AudienceRange> composite, and non-repeating.
type AudienceRangePrecision struct {
	Value string `xml:",innerxml"` // List31
	GeneralAttributes
}

// AudienceRangeQualifier represents an ONIX code specifying the
// attribute (age, school grade etc) which is measured by the value
// in the <AudienceRangeValue> element. Mandatory in each occurrence
// of the <AudienceRange> composite, and non-repeating.
type AudienceRangeQualifier struct {
	Value string `xml:",innerxml"` // List30
	GeneralAttributes
}

// AudienceRangeValue represents a value indicating an exact position
// within a range, or the upper or lower end of a range.
type AudienceRangeValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// Barcode represents a group of data elements which together specify
// a barcode type and its position on a product. Optional: expected
// to be used only in North America. Repeatable if more than one type
// of barcode is carried on a single product. The absence of this
// composite does not mean that a product is not bar-coded.
type Barcode struct {
	BarcodeType       BarcodeType
	PositionOnProduct *PositionOnProduct
	GeneralAttributes
}

// BarcodeType represents an ONIX code indicating whether, and in what form,
// a barcode is carried on a product. Mandatory in any instance of the
// <Barcode> composite, and non-repeating.
type BarcodeType struct {
	Value string `xml:",innerxml"` // List141
	GeneralAttributes
}

// BatchBonus represents a repeatable group of data elements which together
// specify a batch bonus, ie a quantity of free copies which are supplied
// with a certain order quantity. The <BatchBonus> composite is optional.
type BatchBonus struct {
	BatchQuantity BatchQuantity
	FreeQuantity  FreeQuantity
	GeneralAttributes
}

// BatchQuantity represents the number of copies which must be ordered to
// obtain the free copies specified in <FreeQuantity>. Mandatory in each
// occurrence of the <BatchBonus> composite, and non-repeating.
type BatchQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// Bible represents a group of data elements which together describe
// features of an edition of the Bible or of a selected Biblical text.
// Mandatory in each occurrence of the <ReligiousText> composite that
// does not include a <ReligiousTextIdentifier> element, and non-repeating.
type Bible struct {
	BibleContents          []BibleContents
	BibleVersion           []BibleVersion
	StudyBibleType         *StudyBibleType
	BiblePurpose           []BiblePurpose
	BibleTextOrganization  *BibleTextOrganization
	BibleReferenceLocation *BibleReferenceLocation
	BibleTextFeature       []BibleTextFeature
	GeneralAttributes
}

// BibleContents represents an ONIX code indicating the content of an edition
// of the Bible or selected Biblical text, for example ‘New Testament’,
// ‘Apocrypha’, ‘Pentateuch’. Mandatory in each occurrence of the <Bible>
// composite, and repeatable so that a list such as ‘Old Testament and
// Apocrypha’ can be expressed.
type BibleContents struct {
	Value string `xml:",innerxml"` // List82
	GeneralAttributes
}

// BiblePurpose represents an ONIX code indicating the purpose for which a
// Bible or selected Biblical text is intended, for example ‘Family’,
//‘Lectern/pulpit’. Optional and repeatable.
type BiblePurpose struct {
	Value string `xml:",innerxml"` // List85
	GeneralAttributes
}

// BibleReferenceLocation represents an ONIX code indicating where references
// are located as part of the content of a Bible or selected Biblical text,
// for example ‘Center column’. Optional and non-repeating.
type BibleReferenceLocation struct {
	Value string `xml:",innerxml"` // List87
	GeneralAttributes
}

// BibleTextFeature represents an ONIX code specifying a feature of a Bible
// text not covered elsewhere, eg red letter. Optional and repeatable.
type BibleTextFeature struct {
	Value string `xml:",innerxml"` // List97
	GeneralAttributes
}

// BibleTextOrganization represents an ONIX code indicating the way in
// which the content of a Bible or selected Biblical text is organized, for
// example ‘Chronological’, ‘Chain reference’. Optional and non-repeating.
type BibleTextOrganization struct {
	Value string `xml:",innerxml"` // List86
	GeneralAttributes
}

// BibleVersion represents an ONIX code indicating the version of a Bible or
// selected Biblical text, for example ‘King James’, ‘Jerusalem’,  ‘New
// American Standard’, ‘Reina Valera’. Mandatory in each occurrence of the
// <Bible> composite, and repeatable if a work includes text in two or more versions.
type BibleVersion struct {
	Value string `xml:",innerxml"` // List83
	GeneralAttributes
}

// BiographicalNote represents a biographical note about a contributor to the
// product. (See the <TextContent> composite in Group P.14 for a biographical
// note covering all contributors to a product in a single text.) Optional,
// and repeatable to provide parallel biographical notes in multiple languages.
// The language attribute is optional for a single instance of <BiographicalNote>,
// but must be included in each instance if <BiographicalNote> is repeated.
// May occur with a person name or with a corporate name. A biographical note
// in ONIX should always contain the name of the person or body concerned, and
// it should always be presented as a piece of continuous text consisting of
// full sentences. Some recipients of ONIX data feeds will not accept text which
// has embedded URLs. A contributor website link can be sent using the <Website> composite.
type BiographicalNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// BookClubAdoption represents a free text describing the adoption of the
// product as a book club selection. Optional, and repeatable if parallel
// text is provided in multiple languages. The language attribute is optional
// for a single instance of <BookClubAdoption>, but must be included in each
// instance if <BookClubAdoption> is repeated.
type BookClubAdoption struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// CBO (Committed backorder quantity) represents the quantity of stock on
// order which is already committed to meet backorders. Optional and non-repeating.
type CBO struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

// CitationNote represents a free text note giving any additional information
// about cited content, for example a detailed volume, issue and page reference
// to content cited from a periodical. Optional, and repeatable when parallel
// text is provided in multiple languages. The language attribute is optional
// for a single instance of <CitationNote>, but must be included in each
// instance if <CitationNote> is repeated.
type CitationNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// CitedContent represents a group of data elements which together describe
// a piece of cited content. Optional and repeatable.
type CitedContent struct {
	CitedContentType CitedContentType
	ContentAudience  []ContentAudience
	Territory        *Territory
	SourceType       *SourceType
	SourceTitle      []SourceTitle
	CitationNote     []CitationNote
	ResourceLink     []ResourceLink
	ContentDate      []ContentDate
	ReviewRating     *ReviewRating
	ListName         []ListName
	PositionOnList   *PositionOnList
	GeneralAttributes
}

// CitedContentType represents an ONIX code indicating the type of content
// which is being cited. Mandatory in each occurrence of the <CitedContent>
// composite, and non-repeating.
type CitedContentType struct {
	Value string `xml:",innerxml"` // List156
	GeneralAttributes
}

// CityOfPublication represents the name of a city or town associated with
// the imprint or publisher. Optional, and repeatable if parallel names
// for a single location appear on the title page in multiple languages,
// or if the imprint carries two or more cities of publication.
//
// A place of publication is normally given in the form in which it appears
// on the title page. If the place name appears in more than one language,
// <CityOfPublication> may be repeated. The language attribute is optional
// with a single instance of <CityOfPublication>, but must be included in
// each instance if <CityOfPublication> is repeated.
type CityOfPublication struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// CollateralDetail represents covers data element Groups P.14 to P.17, all of
// which are primarily concerned with information and/or resources which in
// one way or another support the marketing of the product. The block as a
// whole is non-repeating. It is not mandatory within the <Product> record,
// nor are any of the individual sections mandatory within an occurrence of
// the block.
type CollateralDetail struct {
	TextContent        []TextContent
	CitedContent       []CitedContent
	SupportingResource []SupportingResource
	Prize              []Prize
	GeneralAttributes
}

// Collection represents a repeatable group of data elements which carry
// attributes of a collection of which the product is part. The composite
// is optional.
type Collection struct {
	CollectionType       CollectionType
	SourceName           *SourceName
	CollectionIdentifier []CollectionIdentifier
	CollectionSequence   []CollectionSequence
	TitleDetail          []TitleDetail
	GeneralAttributes
}

// CollectionIdentifier represents a repeatable group of data elements
// which together define an identifier of a bibliographic collection. The
// composite is optional, and may only repeat if two or more identifiers
// of different types are sent. It is not permissible to have two identifiers
// of the same type.
type CollectionIdentifier struct {
	CollectionIDType CollectionIDType
	IDTypeName       *IDTypeName
	IDValue          IDValue
	GeneralAttributes
}

// CollectionIDType represents an ONIX code identifying a scheme from which
// an identifier in the <IDValue> element is taken. Mandatory in each
// occurrence of the <CollectionIdentifier> composite, and non-repeating.
type CollectionIDType struct {
	Value string `xml:",innerxml"` // List13
	GeneralAttributes
}

// CollectionSequence represents an optional and repeatable group of data
// elements which indicates some ordinal position of a product within a
// collection. Different ordinal positions may be specified using separate
// repeats of the composite – for example, a product may be published first
// while also being third in narrative order within a collection.
type CollectionSequence struct {
	CollectionSequenceType     CollectionSequenceType
	CollectionSequenceTypeName *CollectionSequenceTypeName
	CollectionSequenceNumber   CollectionSequenceNumber
	GeneralAttributes
}

// CollectionSequenceNumber represents a number which specifies the ordinal
// position of the product in a collection. The ordinal position may be a
// simple number (1st, 2nd, 3rd etc) or may be multi-level if the collection
// has a multi-level structure (ie there are both collection and sub-collection
// title elements). Mandatory and non-repeating within the <CollectionSequence>
// composite.
type CollectionSequenceNumber struct {
	Value string `xml:",innerxml"` // dt.MultiLevelNumber
	GeneralAttributes
}

// CollectionSequenceType represents an ONIX code identifying the type of
// ordering used for the product’s sequence number within the collection.
// Mandatory and non-repeating within the <CollectionSequence> composite.
type CollectionSequenceType struct {
	Value string `xml:",innerxml"` // List197
	GeneralAttributes
}

// CollectionSequenceTypeName represents a name which describes a proprietary
// order used for the product’s sequence number within the collection. Must
// be included when, and only when, the code in the <CollectionSequenceType>
// field indicates a proprietary scheme. Optional and non-repeating.
type CollectionSequenceTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// CollectionType represents an ONIX code indicating the type of a collection:
// publisher collection, ascribed collection, or unspecified. Mandatory in each
// occurrence of the <Collection> composite, and non-repeating.
type CollectionType struct {
	Value string `xml:",innerxml"` // List148
	GeneralAttributes
}

// ComparisonProductPrice represents an optional and repeatable group of data
// elements that together define a price for a directly comparable product, to
// facilitate supply of price data to retailers who do not receive a full ONIX
// record for that comparable product. This is primarily intended for use
// within a <Product> record for a digital product, to provide a price for a
// comparable physical product.
//
// Those using this composite should be wary of the volatile nature of product
// prices: special note should be taken of the risk of stale data being stored
// in data recipients’ systems when prices for the comparison product are
// updated, as those updates to the comparison product may occur outside the
// context of the main product being described in the <Product> record. Because
// of this, ONIX suppliers are cautioned of the risk of contradictory data in
// separate data feeds. This composite should not be supplied unless
// specifically requested by a particular recipient.
//
// The inclusion of a comparison price in itself implies nothing about the
// availability or status of the comparable product. However, there may be
// legal requirements in particular territories relating to the use of
// comparison prices in promotion that users of this data must comply with.
//
// Note that the comparison product price composite does not include all the
// features of the <Price> composite: for example, <PriceQualifier> is not
// included. Thus data providers should ensure that any conditions attached
// to the comparison product price are such that it is directly comparable
// to the price of the main product being described.
type ComparisonProductPrice struct {
	ProductIdentifier []ProductIdentifier
	PriceType         *PriceType
	PriceAmount       PriceAmount
	CurrencyCode      *CurrencyCode
	GeneralAttributes
}

// Complexity represents an optional and repeatable group of data elements
// which together describe the level of complexity of a text.
type Complexity struct {
	ComplexitySchemeIdentifier ComplexitySchemeIdentifier
	ComplexityCode             ComplexityCode
	GeneralAttributes
}

// ComplexityCode represents a code specifying the level of complexity of
// a text. Mandatory in each occurrence of the <Complexity> composite, and
// non-repeating.
type ComplexityCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ComplexitySchemeIdentifier represents an ONIX code specifying the scheme
// from which the value in <ComplexityCode> is taken. Mandatory in each
// occurrence of the <Complexity> composite, and non-repeating.
type ComplexitySchemeIdentifier struct {
	Value string `xml:",innerxml"` // List32
	GeneralAttributes
}

// ComponentNumber represents the number (if any) which is given to the
// content item in the product, in the form (eg Arabic or Roman) in which
// it is given in the product. Optional and non-repeating.
type ComponentNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ComponentTypeName represents the generic name (if any) which is given in
// the product to the type of section which the content item represents, eg
// Chapter, Part, Track. Optional and non-repeating; but either this field
// or a title (in the <TitleDetail> composite), or both, must be present in
// each occurrence of the <ContentItem>.
type ComponentTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// Conference [DEPRECATED] represents a group of data elements which together
// describe a conference to which the product is related. Optional, and
// repeatable if the product contains material from two or more conferences.
// The whole of the <Conference> composite is deprecated, in favor of the
// <Event> composite which has an equivalent structure.
type Conference struct {
	ConferenceRole    *ConferenceRole
	ConferenceName    ConferenceName
	ConferenceAcronym *ConferenceAcronym
	ConferenceNumber  *ConferenceNumber
	ConferenceTheme   *ConferenceTheme
	ConferenceDate    *ConferenceDate
	ConferencePlace   *ConferencePlace
	ConferenceSponsor []ConferenceSponsor
	Website           []Website
	GeneralAttributes
}

// ConferenceAcronym [DEPRECATED] represents an acronym used as a short form
// of the name of a conference or conference series given in the
// <ConferenceName> element. Optional and non-repeating.
type ConferenceAcronym struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ConferenceDate [DEPRECATED] represents the date of a conference to which
// the product is related. Optional and non-repeating.
type ConferenceDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// ConferenceName [DEPRECATED] represents the name of a conference or
// conference series to which the product is related. This element is
// mandatory in each occurrence of the <Conference> composite, and non-
// repeating.
type ConferenceName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ConferenceNumber [DEPRECATED] represents the number of a conference to
// which the product is related, within a conference series. Optional and
// non-repeating.
type ConferenceNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// ConferencePlace [DEPRECATED] represents the place of a conference to which
// the product is related. Optional and non-repeating.
type ConferencePlace struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ConferenceRole [DEPRECATED] represents an ONIX code which indicates the
// relationship between the product and a conference to which it is related,
// eg Proceedings of / Selected papers from / Developed from. Optional and
// non-repeating.
type ConferenceRole struct {
	Value string `xml:",innerxml"` // List20
	GeneralAttributes
}

// ConferenceSponsor [DEPRECATED] represents an optional and repeatable group
// of data elements which together identify a sponsor of a conference. Either
// an identifier, or one or other of <PersonName> or <CorporateName>, or both
// an identifier and a name, must be present in each occurrence of the composite.
type ConferenceSponsor struct {
	ConferenceSponsorIdentifier []ConferenceSponsorIdentifier
	PersonName                  *PersonName
	CorporateName               *CorporateName
	GeneralAttributes
}

// ConferenceSponsorIDType [DEPRECATED] represents an ONIX code which
// identifies the scheme from which the value in the <IDValue> element is
// taken. Mandatory in each occurrence of the <ConferenceSponsorIdentifier>
// composite, and non-repeating.
type ConferenceSponsorIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// ConferenceSponsorIdentifier [DEPRECATED] represents an optional and
// repeatable group of data elements which together carry a coded identifier
// for a sponsor of a conference.
type ConferenceSponsorIdentifier struct {
	ConferenceSponsorIDType ConferenceSponsorIDType
	IDTypeName              *IDTypeName
	IDValue                 IDValue
	GeneralAttributes
}

// ConferenceTheme [DEPRECATED] represents the thematic title of an individual
// conference in a series that has a series name in the <ConferenceName>
// element. Optional and non-repeating.
type ConferenceTheme struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// ContactName represents a free text giving the name, department,
// phone number, etc for a contact person in the product contact organization
// who is responsible for the product. Optional and non-repeating.
type ContactName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ContentAudience represents an ONIX code which identifies the audience for
// which the text in the <Text> element is intended. Mandatory in each
// occurrence of the <TextContent> composite, and repeatable.
type ContentAudience struct {
	Value string `xml:",innerxml"` // List154
	GeneralAttributes
}

// ContentDate represents an optional and repeatable group of data elements
// which together specify a date associated with the text carried in an
// occurrence of the <TextContent> composite, eg date when quoted text was published.
type ContentDate struct {
	ContentDateRole ContentDateRole
	DateFormat      *DateFormat
	Date            Date
	GeneralAttributes
}

// ContentDateRole represents an ONIX code indicating the significance of the
// date in relation to the text content. Mandatory in each occurrence of the
// <ContentDate> composite, and non-repeating.
type ContentDateRole struct {
	Value string `xml:",innerxml"` // List155
	GeneralAttributes
}

// ContentDetail represents the single data element Group P.18. The block as a
// whole is non-repeating. It is not mandatory within the <Product> record,
// and is used only when there is a requirement to describe individual chapters
// or parts within a product in a fully structured way. The more usual ONIX
// practice is to send a table of contents as text, possibly in XHTML, in Group P.14.
type ContentDetail struct {
	ContentItem []ContentItem
	GeneralAttributes
}

// ContentItem represents a repeatable group of data elements which together
// describe a content item within a product. Mandatory in any occurrence of
// the <ContentDetail> composite.
type ContentItem struct {
	LevelSequenceNumber *LevelSequenceNumber
	TextItem            TextItem
	Contributor         []Contributor
	Subject             []Subject
	NameAsSubject       []NameAsSubject
	TextContent         []TextContent
	CitedContent        []CitedContent
	SupportingResource  []SupportingResource
	RelatedWork         []RelatedWork
	RelatedProduct      []RelatedProduct
	ComponentTypeName   ComponentTypeName
	ComponentNumber     *ComponentNumber
	TitleDetail         []TitleDetail
	GeneralAttributes
}

// Contributor represents a group of data elements which together describe a
// personal or corporate contributor to a content item. Optional and repeatable.
type Contributor struct {
	SequenceNumber          *SequenceNumber
	ContributorRole         []ContributorRole
	FromLanguage            []FromLanguage
	ToLanguage              []ToLanguage
	NameType                *NameType
	AlternativeName         []AlternativeName
	ContributorDate         []ContributorDate
	ProfessionalAffiliation []ProfessionalAffiliation
	Prize                   []Prize
	Website                 []Website
	ContributorDescription  []ContributorDescription
	ContributorPlace        []ContributorPlace
	NameIdentifier          []NameIdentifier
	PersonName              *PersonName
	PersonNameInverted      *PersonNameInverted
	TitlesBeforeNames       *TitlesBeforeNames
	NamesBeforeKey          *NamesBeforeKey
	PrefixToKey             *PrefixToKey
	KeyNames                *KeyNames
	BiographicalNote        []BiographicalNote
	NamesAfterKey           *NamesAfterKey
	SuffixToKey             *SuffixToKey
	LettersAfterNames       *LettersAfterNames
	TitlesAfterNames        *TitlesAfterNames
	Gender                  *Gender
	UnnamedPersons          *UnnamedPersons
	GeneralAttributes
}

// ContributorDate represents a repeatable group of data elements which
// together specify a date associated with the person or organization
// identified in an occurrence of the <Contributor> composite, eg birth or
// death. Optional.
type ContributorDate struct {
	ContributorDateRole ContributorDateRole
	DateFormat          *DateFormat
	Date                Date
	GeneralAttributes
}

// ContributorDateRole represents an ONIX code indicating the significance of
// the date in relation to the contributor name. Mandatory in each occurrence
// of the <ContributorDate> composite, and non-repeating.
type ContributorDateRole struct {
	Value string `xml:",innerxml"` // List177
	GeneralAttributes
}

// ContributorDescription represents a brief text describing a contributor to
// the product, at the publisher’s discretion. Optional, and repeatable to
// provide parallel descriptions in multiple languages. The language attribute
// is optional for a single instance of <ContributorDescription>, but must be
// included in each instance if <ContributorDescription> is repeated. It may
// be used with either a person or corporate name, to draw attention to any
// aspect of a contributor’s background which supports the promotion of the book.
type ContributorDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// ContributorPlace represents an optional and repeatable group of data
// elements which together identify a geographical location with which a
// contributor is associated, used to support ‘local interest’ promotions.
type ContributorPlace struct {
	ContributorPlaceRelator ContributorPlaceRelator
	LocationName            []LocationName
	CountryCode             CountryCode
	RegionCode              *RegionCode
	GeneralAttributes
}

// ContributorPlaceRelator represents an ONIX code identifying the relationship
// between a contributor and a geographical location. Mandatory in each
// occurrence of <ContributorPlace> and non-repeating.
type ContributorPlaceRelator struct {
	Value string `xml:",innerxml"` // List151
	GeneralAttributes
}

// ContributorRole represents an ONIX code indicating the role played by a
// person or corporate body in the creation of the product. Mandatory in each
// occurrence of a <Contributor> composite, and may be repeated if the same
// person or corporate body has more than one role in relation to the product.
type ContributorRole struct {
	Value string `xml:",innerxml"` // List17
	GeneralAttributes
}

// ContributorStatement represents a free text showing how the authorship
// should be described in an online display, when a standard concatenation of
// individual contributor elements would not give a satisfactory presentation.
// Optional, and repeatable if parallel text is provided in multiple languages.
// The language attribute is optional for a single instance of
// <ContributorStatement>, but must be included in each instance if
//  <ContributorStatement> is repeated. When the <ContributorStatement> field
// is sent, the receiver should use it to replace all name detail sent in the
// <Contributor> composite for display purposes only. It does not replace the
// <BiographicalNote> element. The individual name detail must also be sent in
// the <Contributor> composite for indexing and retrieval.
type ContributorStatement struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// CopiesSold represents a free text detailing the number of copies already
// sold, eg for a new paperback, the copies sold in hardback. Optional, and
// repeatable if parallel text is provided in multiple languages. The language
// attribute is optional for a single instance of <CopiesSold>, but must be
// included in each instance if <CopiesSold> is repeated.
type CopiesSold struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// CopyrightOwner represents a repeatable group of data elements which together
// name a copyright owner. Optional, so that a copyright statement can be
// limited to <CopyrightYear>. Each occurrence of the <CopyrightOwner>
// composite must carry a single name (personal or corporate), or an
// identifier, or both.
type CopyrightOwner struct {
	CopyrightOwnerIdentifier []CopyrightOwnerIdentifier
	PersonName               *PersonName
	CorporateName            *CorporateName
	GeneralAttributes
}

// CopyrightOwnerIDType represents an ONIX code which identifies the scheme
// from which the value in the <IDValue> element is taken. Mandatory in each
// occurrence of the <CopyrightOwnerIdentifier> composite, and non-repeating.
type CopyrightOwnerIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// CopyrightOwnerIdentifier represents a group of data elements which together
// represent a coded identification of a copyright owner. Optional, and
// repeatable if sending more than one identifier of different types. May be
// sent either instead of or as well as a name.
type CopyrightOwnerIdentifier struct {
	CopyrightOwnerIDType CopyrightOwnerIDType
	IDTypeName           *IDTypeName
	IDValue              IDValue
	GeneralAttributes
}

// CopyrightStatement represents an optional and repeatable group of data
// elements which together represent a copyright or neighbouring right
// statement for the product. Either the copyright year alone, or a structured
// rights statement listing year(s) and rights holder(s), may be sent as an
// instance of the composite.
type CopyrightStatement struct {
	CopyrightType  *CopyrightType
	CopyrightYear  []CopyrightYear
	CopyrightOwner []CopyrightOwner
	GeneralAttributes
}

// CopyrightType represents an optional ONIX code indicating the type of right
// covered by the statement, typically a copyright or neighbouring right. If
// omitted, the default is that the statement represents a copyright.
type CopyrightType struct {
	Value string `xml:",innerxml"` // List219
	GeneralAttributes
}

// CopyrightYear represents the copyright year as it appears in a copyright
// statement on the product. Mandatory in each occurrence of the
// <CopyrightStatement> composite, and repeatable if several years are listed.
type CopyrightYear struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// CorporateName represents the name of a corporate body, used here for a
// corporate copyright owner. Optional and non- repeating. Each occurrence of
// the <CopyrightOwner> composite may carry a single name (personal or
// corporate), or an identifier, or both a name and an identifier.
type CorporateName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// CorporateNameInverted represents the name of a corporate body which
// contributed to the creation of the product, presented in inverted order,
// with the element used for alphabetical sorting placed first. Optional and
// non-repeating: see Group P.7 introductory text for valid options.
type CorporateNameInverted struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// CountriesIncluded represents one or more ISO standard codes identifying
// countries included in the territory. Successive codes must be separated by
// spaces. Optional and non-repeating, but either <CountriesIncluded> or
// <RegionsIncluded> is mandatory in each occurrence of the <Territory> composite.
type CountriesIncluded struct {
	Value string `xml:",innerxml"` // CountryCodeList
	GeneralAttributes
}

// CountriesExcluded represents one or more ISO standard codes identifying
// countries excluded from the territory. Successive codes must be separated by
// spaces. Optional and non-repeating, and can only occur if the
// <RegionsIncluded> element is also present and includes a supra-national
// region code (such as ‘World’).
type CountriesExcluded struct {
	Value string `xml:",innerxml"` // CountryCodeList
	GeneralAttributes
}

// CountryCode represents a code identifying a country with which a contributor
// is particularly associated. Optional and non-repeatable. There must be an
// occurrence of either the <CountryCode> or the <RegionCode> elements in each
// occurrence of <ContributorPlace>.
type CountryCode struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

// CountryOfManufacture represents an ISO code identifying the country of
// manufacture of a single-item product, or of a multiple-item product when
// all items are manufactured in the same country. This information is needed
// in some countries to meet regulatory requirements. Optional and non-repeating.
type CountryOfManufacture struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

// CountryOfPublication represents a code identifying the country where the
// product is published. Optional and non-repeating.
type CountryOfPublication struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

// CurrencyCode represents an ISO standard code identifying the currency in
// which all monetary amounts in an occurrence of the <Price> composite are
// stated. Optional and non-repeating; and required if the currency is not the
// default currency for the message (which may be set in <DefaultCurrencyCode>).
// All ONIX messages must include an explicit statement of the currency used
// for any prices. To avoid any possible ambiguity, it is strongly recommended
// that the currency should be repeated here for each individual price.
type CurrencyCode struct {
	Value string `xml:",innerxml"` // List96
	GeneralAttributes
}

// CurrencyZone [DEPRECATED] represents an ONIX code identifying a currency
// zone in which the price stated in an occurrence of the <Price> composite
// is applicable. Optional and non-repeating. Deprecated – use Country or
// Region codes instead.
type CurrencyZone struct {
	Value string `xml:",innerxml"` // List172
	GeneralAttributes
}

// Date represents the date specified in the <PriceDateRole> field. Mandatory
// in each occurrence of the <PriceDate> composite, and non-repeating. <Date>
// may carry a dateformat attribute: if the attribute is missing, then
// <DateFormat> indicates the format of the date; if both dateformat attribute
// and <DateFormat> element are missing, the default format is YYYYMMDD.
type Date struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// DateFormat [DEPRECATED] represents an ONIX code indicating the format in
// which the date is given in <Date>. Optional in each occurrence of the
// <PriceDate> composite, and non-repeating. Deprecated – where possible,
// use dateformat attribute instead.
type DateFormat struct {
	Value string `xml:",innerxml"` // List55
	GeneralAttributes
}

// DefaultCurrencyCode represents an ISO standard code indicating the currency
// which is assumed for prices listed in the message, unless explicitly stated
// otherwise in a <Price> composite in a product record. Optional and non-
// repeating. All ONIX messages must include an explicit statement of the
// currency used for any prices. To avoid any possible ambiguity, it is
// strongly recommended that the currency should be repeated in the <Price>
// composite for each individual price.
type DefaultCurrencyCode struct {
	Value string `xml:",innerxml"` // List96
	GeneralAttributes
}

// DefaultLanguageOfText represents an ISO standard code indicating the default
// language which is assumed for the text of products listed in the message,
// unless explicitly stated otherwise by sending a ‘language of text’ element
// in the product record. This default will be assumed for all product records
// which do not specify a language in Group P.10. Optional and non-repeating.
type DefaultLanguageOfText struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

// DefaultPriceType represents an ONIX code indicating the default price type
// which is assumed for prices listed in the message, unless explicitly stated
// otherwise in a <Price> composite in the product record. Optional and non-repeating.
type DefaultPriceType struct {
	Value string `xml:",innerxml"` // List58
	GeneralAttributes
}

// DeletionText represents a free text which indicates the reason why an ONIX
// record is being deleted. Optional and repeatable, and may occur only when
// the <NotificationType> element carries the code value 05. The language
// attribute is optional for a single instance of <DeletionText>, but must be
// included in each instance if <DeletionText> is repeated. Note that it refers
// to the reason why the record is being deleted, not the reason why a product
// has been ‘deleted’ (in industries which use this terminology when a product
// is withdrawn).
type DeletionText struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// DescriptiveDetail represents  data element Groups P.3 to P.13, all of which
// are essentially part of the factual description of the form and content of
// a product. The block as a whole is non-repeating. It is mandatory in any
// <Product> record unless the <NotificationType> in Group P.1 indicates that
// the record is an update notice which carries only those blocks in which
// changes have occurred.
type DescriptiveDetail struct {
	ProductComposition      ProductComposition
	ProductForm             ProductForm
	ProductFormDetail       []ProductFormDetail
	ProductFormFeature      []ProductFormFeature
	ProductPackaging        *ProductPackaging
	ProductFormDescription  []ProductFormDescription
	TradeCategory           *TradeCategory
	PrimaryContentType      *PrimaryContentType
	ProductContentType      []ProductContentType
	Measure                 []Measure
	CountryOfManufacture    *CountryOfManufacture
	EpubTechnicalProtection []EpubTechnicalProtection
	EpubUsageConstraint     []EpubUsageConstraint
	EpubLicense             *EpubLicense
	MapScale                []MapScale
	ProductClassification   []ProductClassification
	ProductPart             []ProductPart
	Collection              []Collection
	NoCollection            *NoCollection
	TitleDetail             []TitleDetail
	NoContributor           *NoContributor
	Contributor             []Contributor
	ContributorStatement    []ContributorStatement
	Event                   []Event
	Conference              []Conference
	ReligiousText           *ReligiousText
	EditionType             []EditionType
	EditionStatement        []EditionStatement
	NoEdition               *NoEdition
	Language                []Language
	Extent                  []Extent
	Illustrated             *Illustrated
	NumberOfIllustrations   *NumberOfIllustrations
	IllustrationsNote       []IllustrationsNote
	AncillaryContent        []AncillaryContent
	Subject                 []Subject
	NameAsSubject           []NameAsSubject
	AudienceCode            []AudienceCode
	Audience                []Audience
	AudienceRange           []AudienceRange
	AudienceDescription     []AudienceDescription
	Complexity              []Complexity
	GeneralAttributes
}

// Discount represents a repeatable group of data elements which together
// define a discount either as a percentage or as an absolute amount.
// Optional. Used only when an ONIX message is sent within the context of
// a specific trading relationship.
type Discount struct {
	DiscountType    *DiscountType
	DiscountPercent DiscountPercent
	DiscountAmount  *DiscountAmount
	GeneralAttributes
}

// DiscountAmount represents a discount expressed as an absolute amount per
// copy. Optional and non-repeating; but either <DiscountPercent> or
// <DiscountAmount> (or both) must be present in each occurrence of the
// <Discount> composite. Note that when both are present, they represent two
// different expressions of the same discount – the discounts are not cumulative.
type DiscountAmount struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

// DiscountCode represents a repeatable group of data elements which together
// define a discount code from a specified scheme, and allowing different
// discount code schemes to be supported without defining additional data
// elements. Optional. A discount code is generally used when the exact
// percentage discount (or commission, in an agency business model) that a
// code represents may vary from reseller to reseller (or from agent to agent),
// or if terms must be kept confidential. If the discount (or commission) is
// the same for all resellers (or agents) and need not be kept confidential,
// use <Discount> and <DiscountPercent> instead.
type DiscountCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// DiscountCodeType represents an ONIX code identifying the scheme from which
// the value in the <DiscountCode> element is taken. Mandatory in each
// occurrence of the <DiscountCoded> composite, and non-repeating.
type DiscountCodeType struct {
	Value string `xml:",innerxml"` // List100
	GeneralAttributes
}

// DiscountCodeTypeName represents a name which identifies a proprietary
// discount code. Must be used when, and only when the code in the
// <DiscountCodeType> element indicates a proprietary scheme, eg a wholesaler’s
// own code. Optional and non-repeating.
type DiscountCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// DiscountCoded represents a repeatable group of data elements which together
// define a discount code from a specified scheme, and allowing different
// discount code schemes to be supported without defining additional data
// elements. Optional. A discount code is generally used when the exact
// percentage discount (or commission, in an agency business model) that a code
// represents may vary from reseller to reseller (or from agent to agent), or
// if terms must be kept confidential. If the discount (or commission) is the
// same for all resellers (or agents) and need not be kept confidential, use
// <Discount> and <DiscountPercent> instead.
type DiscountCoded struct {
	DiscountCodeType     DiscountCodeType
	DiscountCodeTypeName *DiscountCodeTypeName
	DiscountCode         DiscountCode
	GeneralAttributes
}

// DiscountType represents an ONIX code identifying a discount type or reason.
// Optional, and non-repeating. When omitted, the default is a simple or rising
// discount (the discount is applied to all units in a qualifying order).
type DiscountType struct {
	Value string `xml:",innerxml"` // List170
	GeneralAttributes
}

// EditionNumber represents the number of a numbered edition. Optional and non-
// repeating. Normally sent only for the second and subsequent editions of a
// work, but by agreement between parties to an ONIX exchange a first edition
// may be explicitly numbered.
type EditionNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// EditionStatement represents a short free-text description of a version or
// edition. Optional, and repeatable if parallel text is provided in multiple
// languages. The language attribute is optional for a single instance of
// <EditionStatement>, but must be included in each instance if
// <EditionStatement> is repeated. When used, an <EditionStatement> must be
// complete in itself, ie it should not be treated as merely supplementary to
// an <EditionType> or an <EditionNumber>, nor as a replacement for them
// Appropriate edition type and number must also be sent, for indexing and
// retrieval. An <EditionStatement> should be strictly limited to describing
// features of the content of the edition, and should not include aspects
// such as rights or market restrictions which are properly covered elsewhere
// in the ONIX record.
type EditionStatement struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// EditionType represents an ONIX code, indicating the type of a version or
// edition. Optional, and repeatable if the product has characteristics of two
// or more types (eg ‘revised’ and ‘annotated’).
type EditionType struct {
	Value string `xml:",innerxml"` // List21
	GeneralAttributes
}

// EditionVersionNumber represents the number of a numbered revision within an
// edition number. To be used only where a publisher uses such two-level
// numbering to indicate revisions which do not constitute a new edition under
// a new ISBN or other distinctive product identifier. Optional and non-
// repeating. If this field is used, an <EditionNumber> must also be present.
type EditionVersionNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// EmailAddress represents a text field giving the e-mail address for a contact
// person in the product contact organization who is responsible for the
// product. Optional and non-repeating.
type EmailAddress struct {
	Value string `xml:",innerxml"` // dt.EmailString
	GeneralAttributes
}

// EndDate represents the date until which a sales restriction is effective.
// Optional and non-repeating.
type EndDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// EpubLicense represents an optional and non-repeatable composite carrying
// the name or title of the license governing use of the product, and a link to
// the license terms in eye-readable or machine-readable form.
type EpubLicense struct {
	EpubLicenseName       []EpubLicenseName
	EpubLicenseExpression []EpubLicenseExpression
	GeneralAttributes
}

// EpubLicenseExpression represents an optional and repeatable composite that
// carries details of a link to an expression of the license terms, which may
// be in human-readable or machine-readable form. Repeatable when there is more
// than one expression of the license.
type EpubLicenseExpression struct {
	EpubLicenseExpressionType     EpubLicenseExpressionType
	EpubLicenseExpressionTypeName *EpubLicenseExpressionTypeName
	EpubLicenseExpressionLink     EpubLicenseExpressionLink
	GeneralAttributes
}

// EpubLicenseExpressionLink represents the URI for the license expression.
// Mandatory in each instance of the <EpubLicenseExpression> composite, and
// non-repeating.
type EpubLicenseExpressionLink struct {
	Value string `xml:",innerxml"` // dt.NonEmptyURI
	GeneralAttributes
}

// EpubLicenseExpressionType represents an ONIX code identifying the nature or
// format of the license expression specified in the <EpubLicenseExpressionLink>
// element. Mandatory within the <EpubLicenseExpression> composite, and non-
// repeating.
type EpubLicenseExpressionType struct {
	Value string `xml:",innerxml"` // List218
	GeneralAttributes
}

// EpubLicenseExpressionTypeName represents A short free-text name for a
// license expression type, when the code in <EpubLicenseExpressionType>
// provides insufficient detail – for example when a machine-readable license
// is expressed using a particular proprietary encoding scheme. Optional and
// non-repeating, and must be included when (and only when) the
// <EpubLicenseExpressionType> element indicates the expression is encoded in a
// proprietary way.
type EpubLicenseExpressionTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// EpubLicenseName represents the name or title of the license. Mandatory in
// any <EpubLicense> composite, and repeatable to provide the license name in
// multiple languages. The language attribute is optional for a single instance
// of <EpubLicenseName>, but must be included in each instance if
// <EpubLicenseName> is repeated.
type EpubLicenseName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// EpubTechnicalProtection represents an ONIX code specifying whether a digital
// product has DRM or other technical protection features. Optional and
// repeatable, if a product has two or more kinds of protection.
type EpubTechnicalProtection struct {
	Value string `xml:",innerxml"` // List144
	GeneralAttributes
}

// EpubUsageConstraint represents an optional and repeatable group of data
// elements which together describe a usage constraint on a digital product
// (or the absence of such a constraint), whether enforced by DRM technical
// protection, inherent in the platform used, or specified by license agreement.
type EpubUsageConstraint struct {
	EpubUsageType   EpubUsageType
	EpubUsageStatus EpubUsageStatus
	EpubUsageLimit  []EpubUsageLimit
	GeneralAttributes
}

// EpubUsageLimit represents an optional and repeatable group of data elements
// which together specify a quantitative limit on a particular type of usage of
// a digital product.
type EpubUsageLimit struct {
	Quantity      Quantity
	EpubUsageUnit EpubUsageUnit
	GeneralAttributes
}

// EpubUsageStatus represents an ONIX code specifying the status of a usage of
// a digital product, eg permitted without limit, permitted with limit,
// prohibited. Mandatory in each occurrence of the <EpubUsageConstraint>
// composite, and non-repeating.
type EpubUsageStatus struct {
	Value string `xml:",innerxml"` // List146
	GeneralAttributes
}

// EpubUsageType represents an ONIX code specifying a usage of a digital
// product. Mandatory in each occurrence of the <EpubUsageConstraint>
// composite, and non-repeating.
type EpubUsageType struct {
	Value string `xml:",innerxml"` // List145
	GeneralAttributes
}

// EpubUsageUnit represents an ONIX code for a unit in which a permitted usage
// quantity is stated. Mandatory in each occurrence of the <EpubUsageLimit>
// composite, and non-repeating.
type EpubUsageUnit struct {
	Value string `xml:",innerxml"` // List147
	GeneralAttributes
}

// Event represents a group of data elements which together describe an event
// to which the product is related. Optional, and repeatable if the product
// contains material from or is related to two or more events.
type Event struct {
	EventRole    EventRole
	EventName    []EventName
	EventAcronym []EventAcronym
	EventNumber  *EventNumber
	EventTheme   []EventTheme
	EventDate    *EventDate
	EventPlace   []EventPlace
	EventSponsor []EventSponsor
	Website      []Website
	GeneralAttributes
}

// EventAcronym represents an acronym used as a short form of the name of an
// event or series of events given in the <EventName> element. Optional, and
// repeatable to provide parallel acronyms for a single event in multiple
// languages. The language attribute is optional for a single instance of
// <EventAcronym>, but must be included in each instance if <EventAcronym> is
// repeated.
type EventAcronym struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// EventDate represents the date of an event to which the product is related.
// Optional and non-repeating.
type EventDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// EventName represents the name of an event or series of events to which the
// product is related. This element is mandatory in each occurrence of the
// <Event> composite, and repeatable to provide parallel names for a single
// event in multiple languages (eg ‘United Nations Climate Change Conference’
// and « Conférences des Nations unies sur les changements climatiques »).
// The language attribute is optional for a single instance of <EventName>, but
// must be included in each instance if <EventName> is repeated.
type EventName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// EventNumber represents the number of an event to which the product is
// related, within a series of events. Optional and non-repeating.
type EventNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// EventPlace represents the place of an event to which the product is
// related. Optional, and repeatable to provide parallel placenames for a
// single location in multiple languages. The language attribute is optional
// for a single instance of <EventPlace>, but must be included in each instance
// if <EventPlaceTheme> is repeated.
type EventPlace struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// EventRole represents an ONIX code which indicates the relationship between
// the product and an event to which it is related, eg Proceedings of
// conference / Selected papers from conference / Programme for sporting event /
// Guide for art exhibition. Mandatory and non-repeating.
type EventRole struct {
	Value string `xml:",innerxml"` // List20
	GeneralAttributes
}

// EventSponsor represents an optional and repeatable group of data elements
// which together identify a sponsor of a event. Either an identifier, or one
// or other of <PersonName> or <CorporateName>, or both an identifier and a
// name, must be present in each occurrence of the composite.
type EventSponsor struct {
	EventSponsorIdentifier []EventSponsorIdentifier
	PersonName             *PersonName
	CorporateName          *CorporateName
	GeneralAttributes
}

// EventSponsorIDType represents an ONIX code which identifies the scheme from
// which the value in the <IDValue> element is taken. Mandatory in each
// occurrence of the <EventSponsorIdentifier> composite, and non-repeating.
type EventSponsorIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// EventSponsorIdentifier represents an optional and repeatable group of data
// elements which together carry a coded identifier for a sponsor of an event.
type EventSponsorIdentifier struct {
	EventSponsorIDType EventSponsorIDType
	IDTypeName         *IDTypeName
	IDValue            IDValue
	GeneralAttributes
}

// EventTheme represents the thematic title of an individual event in a series
// that has an event series name in the <EventName> element. Optional, and
// repeatable to provide parallel thematic titles for a single event in
// multiple languages. The language attribute is optional for a single
// instance of <EventTheme>, but must be included in each instance if
// <EventTheme> is repeated.
type EventTheme struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ExpectedDate represents the date on which a stock shipment is expected.
// Mandatory in each occurrence of the <OnOrderDetail> composite, and non-
// repeating.
type ExpectedDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// Extent represents a repeatable group of data elements which together
// describe an extent pertaining to the product. Optional, but in practice
// required for most products, eg to give the number of pages in a printed
// book or paginated ebook, or to give the running time of an audiobook.
type Extent struct {
	ExtentType       ExtentType
	ExtentValue      ExtentValue
	ExtentUnit       ExtentUnit
	ExtentValueRoman *ExtentValueRoman
	GeneralAttributes
}

// ExtentType represents an ONIX code which identifies the type of extent
// carried in the composite, eg running time for an audio or video product.
// Mandatory in each occurrence of the <Extent> composite, and non-repeating.
// From Issue 9 of the code lists, an extended set of values for <ExtentType>
// has been defined to allow more accurate description of pagination.
type ExtentType struct {
	Value string `xml:",innerxml"` // List23
	GeneralAttributes
}

// ExtentUnit represents an ONIX code indicating the unit used for the
// <ExtentValue> and the format in which the value is presented. Mandatory in
// each occurrence of the <Extent> composite, and non-repeating.
type ExtentUnit struct {
	Value string `xml:",innerxml"` // List24
	GeneralAttributes
}

// ExtentValue represents the numeric value of the extent specified in
// <ExtentType>. Optional, and non-repeating. However, either <ExtentValue>
// or <ExtentValueRoman> must be present in each occurrence of the <Extent>
// composite; and it is very strongly recommended that <ExtentValue> should
// always be included, even when the original product uses Roman numerals.
type ExtentValue struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

// ExtentValueRoman represents the value of the extent expressed in Roman
// numerals. Optional, and non-repeating. Used only for page runs which are
// numbered in Roman.
type ExtentValueRoman struct {
	Value string `xml:",innerxml"` // dt.RomanNumeralString
	GeneralAttributes
}

// FaxNumber [DEPRECATED] represents a fax number of an agent or local
// publisher. Optional and repeatable. Deprecated in this context, in favor of
// providing contact details in the <ProductContact> composite.
type FaxNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// FeatureNote represents a short text note describing a fetature, to be used
// if the <ResourceFeatureType> requires free text rather than a code value,
// or if the code in <FeatureValue> does not adequately describe the feature,
// a short text note may be added. Optional, and repeatable when parallel notes
// are provided in multiple languages. The language attribute is optional for a
// single instance of <FeatureNote>, but must be included in each instance if
// <FeatureNote> is repeated.
type FeatureNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// FeatureValue represents a controlled value that describes a resource feature.
// Presence or absence of this element depends on the <ResourceFeatureType>,
// since some features may not require an accompanying value, while others may
// require free text in <FeatureNote>; and others may have both code and free
// text. Non-repeating.
type FeatureValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// FirstPageNumber represents the number of the first page of a sequence of
// contiguous pages. Mandatory in each occurrence of the <PageRun> composite,
// and non-repeating. Note that here and in the <LastPageNumber> element a page
// ‘number’ may be Arabic, Roman, or an alphanumeric string (eg L123).
type FirstPageNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// FreeQuantity represents the number of free copies which will be supplied
// with an order for the batch quantity specified in the <BatchQuantity> field.
// Mandatory in each occurrence of the <BatchBonus> composite, and non-
// repeating.
type FreeQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// FromLanguage represents source language of translation. Used only when the
// <ContributorRole> code value is B06, B08 or B10 indicating a translator, to
// specify the source language from which the translation was made. This
// element makes it possible to specify a translator’s exact responsibility
// when a work involves translation from two or more languages. Optional, and
// repeatable in the event that a single person has been responsible for
// translation from two or more languages.
type FromLanguage struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

// Funding represents an optional group of data elements which together
// identify a grant or award provided by the entity specified as a funder in
// an occurence of the <Publisher> composite, to subsidise research or
// publication. Repeatable when the funder provides multiple grants or
// awards. Used only when <PublishingRole> indicates the role of a funder.
type Funding struct {
	FundingIdentifier []FundingIdentifier
	GeneralAttributes
}

// FundingIdentifier represents a repeatable group of data elements which
// together identify a particular grant or award. At least one
// <FundingIdentifier> composite must occur in each instance of the <Funding>
// composite. Repeatable when the grant or award has multiple identifiers.
type FundingIdentifier struct {
	FundingIDType FundingIDType
	IDTypeName    *IDTypeName
	IDValue       IDValue
	GeneralAttributes
}

// FundingIDType represents an ONIX code identifying the scheme from which the
// identifier in the <IDValue> element is taken. Mandatory in each occurrence
// of the <FundingIdentifier> composite, and non-repeating.
type FundingIDType struct {
	Value string `xml:",innerxml"` // List228
	GeneralAttributes
}

// Gender represents an optional ONIX code specifying the gender of a personal
// contributor. Not repeatable. Note that this indicates the gender of the
// contributor’s public identity (which may be pseudonymous) based on
// designations used in ISO 5218, rather than the gender identity, biological
// sex or sexuality of a natural person.
type Gender struct {
	Value string `xml:",innerxml"` // List229
	GeneralAttributes
}

// Header represents a group of data elements which together constitute a
// message header. Mandatory in any ONIX for Books message, and non-repeating.
// In ONIX 3.0, a number of redundant elements have been deleted, and the
// Sender and Addressee structures and the name and format of the
// <SentDateTime> element have been made consistent with other current ONIX
// formats.
type Header struct {
	Sender                Sender
	Addressee             []Addressee
	MessageNumber         *MessageNumber
	MessageRepeat         *MessageRepeat
	SentDateTime          SentDateTime
	MessageNote           []MessageNote
	DefaultLanguageOfText *DefaultLanguageOfText
	DefaultPriceType      *DefaultPriceType
	DefaultCurrencyCode   *DefaultCurrencyCode
	GeneralAttributes
}

// IDTypeName represents a name which identifies a proprietary identifier
// scheme (ie a scheme which is not a standard and for which there is no
// individual ID type code). Must be used when, and only when, the code in
// the <SenderIDType> element indicates a proprietary scheme. Optional and
// non-repeating.
type IDTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// IDValue represents an identifier of the type specified in the
// <SenderIDType> element. Mandatory in each occurrence of the
// <SenderIdentifier> composite, and non-repeating.
type IDValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// Illustrated represents an ONIX code indicating whether a book or other
// textual (usually printed) product has illustrations. The more informative
// free text field <IllustrationsNote> and/or the <AncillaryContent> composite
// are strongly preferred. This element has been added specifically to cater
// for a situation where a sender of product information maintains only a
// yes/no flag, and it should not otherwise be used. Optional and non-
// repeating.
type Illustrated struct {
	Value string `xml:",innerxml"` // List152
	GeneralAttributes
}

// IllustrationsNote represents an illustration and other contents note. For
// books or other text media only, this data element carries text stating the
// number and type of illustrations. The text may also include other content
// items, eg maps, bibliography, tables, index etc. Optional, and repeatable if
// parallel notes are provided in multiple languages. The language attribute is
// optional for a single instance of <IllustrationsNote>, but must be included
// in each instance if <IllustrationsNote> is repeated.
type IllustrationsNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// Imprint represents a repeatable group of data elements which together
// identify an imprint or brand under which the product is marketed. The
// composite must carry either a name identifier or a name or both.
type Imprint struct {
	ImprintIdentifier []ImprintIdentifier
	ImprintName       *ImprintName
	GeneralAttributes
}

// ImprintIdentifier represents a group of data elements which together
// define the identifier of an imprint name. Optional and repeatable, but
// mandatory if the <Imprint> composite does not carry an <ImprintName>.
type ImprintIdentifier struct {
	ImprintIDType ImprintIDType
	IDTypeName    *IDTypeName
	IDValue       IDValue
	GeneralAttributes
}

// ImprintIDType represents an ONIX code which identifies the scheme from
// which the value in the <IDValue> element is taken. Mandatory in each
// occurrence of the <ImprintIdentifier> composite.
type ImprintIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// ImprintName represents the name of an imprint or brand under which the
// product is issued, as it appears on the product. Mandatory if there is no
// imprint identifier in an occurrence of the <Imprint> composite, and optional
// if an imprint identifier is included. Non-repeating.
type ImprintName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// InitialPrintRun represents a free text detailing the number of copies which
// will be printed and any related aspects of the initial publishing effort.
// Optional, and repeatable if parallel text is provided in multiple languages.
// The language attribute is optional for a single instance of
// <InitialPrintRun>, but must be include in each instance if <InitialPrintRun>
// is repeated..
type InitialPrintRun struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// KeyNames represents the name of a person who contributed to the creation of
// the product: key name(s), ie the name elements normally used to open an
// entry in an alphabetical list, eg ‘Smith’ or ‘Garcia Marquez’ or ‘Madonna’
// or ‘Francis de Sales’ (in Saint Francis de Sales). Non-repeating. Required
// if name part elements P.7.11 to P.7.18 are used.
type KeyNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// Language represents an optional and repeatable group of data elements which
// together represent a language, and specify its role and, where required,
// whether it is a country variant.
type Language struct {
	LanguageRole LanguageRole
	LanguageCode LanguageCode
	CountryCode  *CountryCode
	ScriptCode   *ScriptCode
	GeneralAttributes
}

// LanguageCode represents an ISO code indicating a language. Mandatory in each
// occurrence of the <Language> composite, and non-repeating.
type LanguageCode struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

// LanguageRole represents an ONIX code indicating the ‘role’ of a language in
// the context of the ONIX record. Mandatory in each occurrence of the
// <Language> composite, and non-repeating.
type LanguageRole struct {
	Value string `xml:",innerxml"` // List22
	GeneralAttributes
}

// LastPageNumber represents the number of the last page of a sequence of
// contiguous pages (ignoring any blank verso which is left after the last
// text page). This element is omitted if an item begins and ends on the same
// page; otherwise it should occur once and only once in each occurrence of the
// <PageRun> composite.
type LastPageNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// LatestReprintNumber represents the number of the most recent reprint (or
// current ‘impression number’) of the product. Optional and non-repeating.
// This element is used only in certain countries where there is a legal
// requirement to record reprints.
type LatestReprintNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// LettersAfterNames represents the qualifications and honors following a
// person’s names, eg ‘CBE FRS’. Optional and non-repeating.
type LettersAfterNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// LevelSequenceNumber represents a number which specifies the position of a
// content item in a multi-level hierarchy of such items. Numbering starts at
// the top level in the hierarchy, which may represent (eg) chapters in a
// printed book, and the first item at the top level is numbered 1. Numbers
// should be assigned solely with a view to the logic of the ONIX description
// and not in relation to any other characteristics of the items being numbered
// (such as their typographical layout in a printed table of contents).
// <LevelSequenceNumber> is not a required field, but it is strongly
// recommended for structured tables of contents. If used, it must occur once
// and only once in each occurrence of the <ContentItem> composite.
type LevelSequenceNumber struct {
	Value string `xml:",innerxml"` // dt.MultiLevelNumber
	GeneralAttributes
}

// ListName represents the name of a bestseller list, when the <CitedContent>
// composite is used to refer to a position which a product has reached on such
// a list. Optional, and repeatable to provide a parallel list name in multiple
// languages. The language attribute is optional for a single instance of
// <ListName>, but must be included in each instance if <ListName> is repeated.
type ListName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// LocationIdentifier represents an optional and repeatable group of data
// elements which together define the identifier of a stock location in
// accordance with a specified scheme, and allowing different types of location
// identifier to be supported without defining additional data elements.
type LocationIdentifier struct {
	LocationIDType LocationIDType
	IDTypeName     *IDTypeName
	IDValue        IDValue
	GeneralAttributes
}

// LocationIDType represents an ONIX code identifying the scheme from which the
// identifier in the <IDValue> element is taken. Mandatory in each occurrence
// of the <LocationIdentifier> composite, and non-repeating.
type LocationIDType struct {
	Value string `xml:",innerxml"` // List92
	GeneralAttributes
}

// LocationName represents the name of a stock location. Optional and
// repeatable to provide parallel names for a single location in multiple
// languages (eg Baile Átha Cliath and Dublin, or Bruxelles and Brussel). The
// language attribute is optional for a single instance of <LocationName>, but
// must be included in each instance if <LocationName> is repeated..
type LocationName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// MainSubject represents an empty element that identifies an instance of the
// <Subject> composite as representing the main subject category for the
// product. The main category may be expressed in more than one subject scheme,
// ie there may be two or more instances of the <Subject> composite, using
// different schemes, each carrying the <MainSubject/> flag. Optional and non-
// repeating in each occurrence of the <Subject> composite.
type MainSubject struct {
	Dummy string `xml:",innerxml"` // An empty element, but needed for encoding/xml to marshal it
	GeneralAttributes
}

// MapScale represents the scale of a map, expressed as a ratio 1:nnnnn; only
// the number nnnnn is carried in the data element, without spaces or
// punctuation. Optional, and repeatable if a product comprises maps with two
// or more different scales.
type MapScale struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// Market represents a repeatable group of data elements which together give
// details of a geographical territory and any non-geographical sales
// restrictions that apply within it. Optional in terms of the schema
// definitions, but required in most ONIX applications.
type Market struct {
	Territory        Territory
	SalesRestriction []SalesRestriction
	GeneralAttributes
}

// MarketDate represents a repeatable group of data elements which together
// specify a date associated with the publishing status of the product in a
// specified market, eg ‘local publication date’. Optional, but a date of
// publication must be specified either here as a ‘local pubdate’ or in P.20.
// Other dates relating to the publication of the product in the specific
//  market may be sent in further repeats of the composite.
type MarketDate struct {
	MarketDateRole MarketDateRole
	DateFormat     *DateFormat
	Date           Date
	GeneralAttributes
}

// MarketDateRole represents an ONIX code indicating the significance of the
// date. Mandatory in each occurrence of the <MarketDate> composite, and non-
// repeating.
type MarketDateRole struct {
	Value string `xml:",innerxml"` // List163
	GeneralAttributes
}

// MarketPublishingDetail represents a group of data elements which together
// give details of the publishing status of a product within a specified
// market. Optional and non-repeating within an occurrence of the
// <ProductSupply> block.
type MarketPublishingDetail struct {
	PublisherRepresentative    []PublisherRepresentative
	ProductContact             []ProductContact
	MarketPublishingStatus     MarketPublishingStatus
	MarketPublishingStatusNote []MarketPublishingStatusNote
	MarketDate                 []MarketDate
	PromotionCampaign          []PromotionCampaign
	PromotionContact           *PromotionContact
	InitialPrintRun            []InitialPrintRun
	ReprintDetail              []ReprintDetail
	CopiesSold                 []CopiesSold
	BookClubAdoption           []BookClubAdoption
	GeneralAttributes
}

// MarketPublishingStatus represents an ONIX code which identifies the status
// of a published product in a specified market. Mandatory in each occurrence
// of the <MarketPublishingDetail> composite, and non-repeating.
type MarketPublishingStatus struct {
	Value string `xml:",innerxml"` // List68
	GeneralAttributes
}

// MarketPublishingStatusNote represents a free text that describes the status
// of a product in a specified market, when the code in <MarketPublishingStatus>
// is insufficient. Optional, but when used, must be accompanied by the
// <MarketPublishingStatus> element. Repeatable if parallel text is provided in
// multiple languages. The language attribute is optional for a single instance
// of <MarketPublishingStatusNote>, but must be included in each instance if
// <MarketPublishingStatusNote> is repeated.
type MarketPublishingStatusNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// Measure represents an optional and repeatable group of data elements which
// together identify a measurement and the units in which it is expressed; used
// to specify the overall dimensions of a physical product including its
// packaging (if any).
type Measure struct {
	MeasureType     MeasureType
	Measurement     Measurement
	MeasureUnitCode MeasureUnitCode
	GeneralAttributes
}

// Measurement represents the number which represents the dimension specified
// in <MeasureType> in the measure units specified in<MeasureUnitCode>.
// Mandatory in each occurrence of the <Measure> composite, and non-repeating.
type Measurement struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

// MeasureType represents an ONIX code indicating the dimension which is
// specified by an occurrence of the measure composite. Mandatory in each
// occurrence of the <Measure> composite, and non-repeating.
type MeasureType struct {
	Value string `xml:",innerxml"` // List48
	GeneralAttributes
}

// MeasureUnitCode represents an ONIX code indicating the measure unit in which
// dimensions are given. Mandatory in each occurrence of the <Measure>
// composite, and non-repeating. This element must follow the dimension to
// which the measure unit applies. See example below.
type MeasureUnitCode struct {
	Value string `xml:",innerxml"` // List50
	GeneralAttributes
}

// MessageNote represents a free text giving additional information about the
// message. Optional, and repeatable in order to provide a note in multiple
// languages. The language attribute is optional for a single instance of
// <MessageNote>, but must be included in each instance if <MessageNote> is
// repeated.
type MessageNote struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// MessageNumber represents a monotonic sequence number of the messages in a
// series sent between trading partners, to enable the receiver to check
// against gaps and duplicates. Optional and non-repeating.
type MessageNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// MessageRepeat represents a number which distinguishes any repeat
// transmissions of a message. If this element is used, the original is
// numbered 1 and repeats are numbered 2, 3 etc. Optional and non-repeating.
type MessageRepeat struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// MinimumOrderQuantity represents the minimum number of copies which must be
// ordered to obtain the price carried in an occurrence of the <Price>
// composite. Optional and non-repeating. If the field is present, the price is
// a quantity price (and only whole multiples of the qualifying quantity may be
// ordered at that price). If the field is omitted, the price applies to a
// single unit.
//
// Note the similarity between <MinimumOrderQuantity> and
// <OrderQuantityMinimum> in P.26.41a: only <MinimumOrderQuantity> has an
// effect on the specification of <Price>. Use of <MinimumOrderQuantity> is
// close in effect to a multi-item trade pack (see <ProductComposition>).
type MinimumOrderQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// NameAsSubject represents an optional and repeatable group of data
// elements which together represent the name of a person or organization
// that is part of the subject of a product. Each instance of the composite
// must contain either:
//
// * one or more of the forms of representation of a person name, with or without an occurrence of the <NameIdentifier> composite; or
//
// * one or more of the forms of representation of a corporate name, with or without an occurrence of the <NameIdentifier> composite; or
//
// * an occurrence of the <NameIdentifier> composite without any accompanying name element(s). The name of a person (not of a corporation) may optionally be followed by details of that person’s professional affiliation.
type NameAsSubject struct {
	NameType                *NameType
	AlternativeName         []AlternativeName
	SubjectDate             []SubjectDate
	ProfessionalAffiliation []ProfessionalAffiliation
	NameIdentifier          []NameIdentifier
	PersonName              PersonName
	PersonNameInverted      *PersonNameInverted
	TitlesBeforeNames       *TitlesBeforeNames
	NamesBeforeKey          *NamesBeforeKey
	PrefixToKey             *PrefixToKey
	KeyNames                KeyNames
	NamesAfterKey           *NamesAfterKey
	SuffixToKey             *SuffixToKey
	LettersAfterNames       *LettersAfterNames
	TitlesAfterNames        *TitlesAfterNames
	Gender                  *Gender
	GeneralAttributes
}

// NameIdentifier represents a repeatable group of data elements which together
// specify a name identifier, used here to carry an identifier for a person or
// organization name given in an occurrence of the <Contributor> composite.
// Optional: see Group P.7 introductory text for valid options.
type NameIdentifier struct {
	NameIDType NameIDType
	IDTypeName *IDTypeName
	IDValue    IDValue
	GeneralAttributes
}

// NameIDType represents an ONIX code which identifies the scheme from which
// the value in the <IDValue> element is taken. Mandatory in each occurrence
// of the <NameIdentifier> composite, and non-repeating.
type NameIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// NamesAfterKey represents the name suffix, or name(s) following a person’s key
// name(s), eg ‘Ibrahim’ (in Anwar Ibrahim). Optional and non-repeating.
type NamesAfterKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// NamesBeforeKey represents the name(s) and/or initial(s) preceding a person’s
// key name(s), eg James J. Optional and non-repeating.
type NamesBeforeKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// NameType represents an ONIX code indicating the type of the name sent in an
// occurrence of the <AlternativeName> composite. Mandatory in each occurrence
// of the composite, and non-repeating.
type NameType struct {
	Value string `xml:",innerxml"` // List18
	GeneralAttributes
}

// NewSupplier represents a group of data elements which together specify a new
// supply source to which orders are referred. Use only when the code in
// <ProductAvailability> indicates ‘no longer available from us, refer to new
// supplier’. Only one occurrence of the composite is permitted in this context.
type NewSupplier struct {
	TelephoneNumber    []TelephoneNumber
	FaxNumber          []FaxNumber
	EmailAddress       []EmailAddress
	SupplierIdentifier []SupplierIdentifier
	SupplierName       *SupplierName
	GeneralAttributes
}

// NoCollection represents an empty element that provides a positive indication
// that a product does not belong to a collection (or ‘series’). This element
// is intended to be used in an ONIX accreditation scheme to confirm that
// collection information is being consistently supplied in publisher ONIX
// feeds. Optional and non-repeating. Must only be sent in a record that has no
// instances of the <Collection> composite and has no collection level title
// elements in Group P.6.
type NoCollection struct {
	Dummy string `xml:",innerxml"` // An empty element, but needed for encoding/xml to marshal it
	GeneralAttributes
}

// NoContributor represents an empty element that provides a positive
// indication that a product has no stated authorship. Intended to be used in
// an ONIX accreditation scheme to confirm that author information is being
// consistently supplied in publisher ONIX feeds. Optional and non-repeating.
// Must only be sent in a record that has no other elements from Group P.7.
type NoContributor struct {
	Dummy string `xml:",innerxml"` // An empty element, but needed for encoding/xml to marshal it
	GeneralAttributes
}

// NoEdition represents an empty element that provides a positive indication
// that a product does not carry any edition information. Intended to be used
// an ONIX accreditation scheme to confirm that edition information is being
// consistently supplied in publisher ONIX feeds. Optional and non-repeating.
// Must only be sent in a record that has no instances of any of the four
// preceding Edition elements.
type NoEdition struct {
	Dummy string `xml:",innerxml"` // An empty element, but needed for encoding/xml to marshal it
	GeneralAttributes
}

// NoPrefix represents an empty element that provides a positive indication
// that a title element does not include any prefix that is ignored for
// sorting purposes. Optional and non-repeating, and must only be used when
// <TitleWithoutPrefix> is used and no <TitlePrefix> element is present.
type NoPrefix struct {
	Dummy string `xml:",innerxml"` // An empty element, but needed for encoding/xml to marshal it
	GeneralAttributes
}

// NoProduct represents an empty element that provides a positive indication
// that a message does not carry any Product records. Intended to be used only
// in empty ‘delta’ update messages to provide confirmation that there have
// been no updates since the previous message. Optional and non-repeating, but
// must be used in an ONIX message that contains no <Product> composites.
type NoProduct struct {
	Dummy string `xml:",innerxml"` // An empty element, but needed for encoding/xml to marshal it
	GeneralAttributes
}

// NotificationType represents an ONIX code which indicates the type of
// notification or update which you are sending. Mandatory and non-repeating.
type NotificationType struct {
	Value string `xml:",innerxml"` // List1
	GeneralAttributes
}

// Number represents the number of illustrations or other content items of the
// type specified in <AncillaryContentType>. Optional and non-repeating.
type Number struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

// NumberOfCopies represents the number of copies in a Product. When product
// parts are listed as a specified number of copies of a single item, usually
// identified by a <ProductIdentifier>, <NumberOfCopies> must be used to
// specify the quantity, even if the number is ‘1’. It must be used when a
// multiple-item product or pack contains (a) a quantity of a single item; or
// (b) one of each of several different items (as in a multi-volume set); or
// (c) one or more of each of several different items (as in a dumpbin carrying
// copies of two different books, or a classroom pack containing a teacher’s
// text and twenty student texts). Consequently the element is mandatory, and
// non-repeating, in an occurrence of the <ProductPart> composite if
// <NumberOfItemsOfThisForm> is not present. It is normally accompanied by a
// <ProductIdentifier>; but in exceptional circumstances, if the sender’s
// system is unable to provide an identifier at this level, it may be sent with
// product form coding and without an ID.
type NumberOfCopies struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// NumberOfIllustrations represents the total number of illustrations in a book
// or other printed product. The more informative free text field
// <IllustrationsNote> and/or the <AncillaryContent> composite are strongly
// preferred, but where a sender of product information maintains only a simple
// numeric field, the <NumberOfIllustrations> element may be used. Optional and
// non-repeating.
type NumberOfIllustrations struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

// NumberOfItemsOfThisForm represents the number of items of a specified form
// in a product. When product parts are listed as a specified number of
// different items in a specified form, without identifying the individual
// items, <NumberOfItemsOfThisForm> must be used to carry the quantity, even
// if the number is ‘1’. Consequently the element is mandatory and non-
// repeating in an occurrence of the <ProductPart> composite if
// <NumberOfCopies> is not present; and it must not be used if
// <ProductIdentifier> is present.
type NumberOfItemsOfThisForm struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// NumberOfPages represents the page extent of a text item within a paginated
// product. Optional and non-repeating, but normally expected when the text
// item is being referenced as part of a structured table of contents.
type NumberOfPages struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type Message struct {
	XMLName   xml.Name `xml:"ONIXMessage"`
	Release   string   `xml:"release,attr"`
	Header    Header
	NoProduct *NoProduct
	Product   []Product
	GeneralAttributes
	ReleaseAttribute
}

// OnHand represents the quantity of stock on hand and available to fulfill new
// orders. Either <StockQuantityCoded> or <OnHand> is mandatory in each
// occurrence of the <Stock> composite, even if the quantity on hand is zero.
// Non-repeating.
type OnHand struct {
	Value string `xml:",innerxml"` // dt.Integer
	GeneralAttributes
}

// OnOrder represents the quantity of stock on order. Optional and non-
// repeating.
type OnOrder struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

// OnOrderDetail represents a repeatable group of data elements which together
// specify details of a stock shipment currently awaited, normally from
// overseas. Optional, and repeatable if more than a single shipment is
// outstanding. Note that quantities in the <OnOrderDetail> composite must be
// included in any total quantity on order given in P.26.37 <OnOrder>, and
// detail need not be given for all outstanding shipments (ie the P.26.37
// <OnOrder> must be greater than or equal to the total of the <OnOrder>
// elements in repeats of the composite).
type OnOrderDetail struct {
	OnOrder      OnOrder
	Proximity    *Proximity
	ExpectedDate ExpectedDate
	GeneralAttributes
}

// OrderQuantityMinimum represents the minimum quantity of the product that may
// be ordered in a single order placed with the supplier. Optional. If omitted,
// any number may be ordered. If supplied without a succeeding Minimum initial
// order quantity, the Minimum order quantity applies to all orders for the
// product. If followed by a Minimum initial order quantity, the Minimum order
// quantity applies to the second and subsequent orders for the product.
// TODO check double defs in spec: P.26.41a & P.26.41b
type OrderQuantityMinimum struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// OrderQuantityMultiple represents the order quantity multiple that must be
// used in any order for the product placed with the supplier. Optional, but
// where supplied, must be preceded by at least one <OrderQuantityMinimum>
// element. For example with a minimum order quantity of 6 and a multiple of 4,
// orders for 6, 10 or 14 copies are acceptable, but orders for fewer than 6,
// or for 7, 8, 9 or 11 copies are not. If omitted, the minimum or any larger
// quantity may be ordered.
type OrderQuantityMultiple struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// OrderTime represents the expected average number of days from receipt of
// order to despatch (for items ‘manufactured on demand’ or ‘only to order’).
// Optional and non-repeating.
type OrderTime struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

// PackQuantity represents the quantity in each carton or binder’s pack in
// stock currently held by the supplier. Optional and non- repeating.
//
// Note that orders do not have to be aligned with multiples of the pack size,
// but such orders may be more convenient to handle.
type PackQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// PageRun represents a repeatable group of data elements which together define
// a run of contiguous pages on which a text item appears. The composite is
// optional, but may be repeated where the text item covers two or more separate
// page runs.
type PageRun struct {
	FirstPageNumber FirstPageNumber
	LastPageNumber  *LastPageNumber
	GeneralAttributes
}

// PartNumber represents a product's part number. When a title element includes
// a part designation within a larger whole (eg Part I, or Volume 3), this
// field should be used to carry the number and its ‘caption’ as text. Optional
// and non-repeating.
type PartNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
	TextscriptAttribute
}

// Percent represents the percentage of the unit value of the product that is
// assignable to a designated product classification. Optional and non-
// repeating. Used when a mixed product (eg book and CD) belongs partly to two
// or more product classifications. If omitted, the product classification
// code applies to 100% of the product.
type Percent struct {
	Value string `xml:",innerxml"` // dt.PercentDecimal
	GeneralAttributes
}

// DiscountPercent represents a discount percentage applicable to the price
// carried in an occurrence of the <Price> composite. Optional and non-
// repeating; but either <DiscountPercent> or <DiscountAmount> or both must be
// present in each occurrence of the <Discount> composite.
type DiscountPercent struct {
	Value string `xml:",innerxml"` // dt.PercentDecimal
	GeneralAttributes
}

// PersonName represents the name of a person who contributed to the creation
// of the product, unstructured, and presented in normal order. Optional and
// non-repeating: see Group P.7 introductory text for valid options.
type PersonName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// PersonNameInverted represents the name of a person who contributed to the
// creation of the product, presented with the element used for alphabetical
// sorting placed first (‘inverted order’). Optional and non-repeating: see
// Group P.7 introductory text for valid options.
type PersonNameInverted struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// PositionOnList represents the position that a product has reached on a
// bestseller list specified in <ListName>. Optional and non-repeating. The
// <ListName> element must also be present if <PositionOnList> is included.
type PositionOnList struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// PositionOnProduct represents an ONIX code indicating a position on a
// product; in this case, the position in which a price appears. Optional,
// but required if the <PrintedOnProduct> element indicates that the price
// appears on the product, even if the position is ‘unknown’. Non-repeating.
type PositionOnProduct struct {
	Value string `xml:",innerxml"` // List142
	GeneralAttributes
}

// PrefixToKey represents a prefix which precedes the key name(s) but which is
// not to be treated as part of the key name, eg ‘van’ in Ludwig van Beethoven.
// This element may also be used for titles that appear after given names and
// before key names, eg ‘Lord’ in Alfred, Lord Tennyson. Optional and non-
//repeating.
type PrefixToKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// Price represents an optional and repeatable group of data elements which
// together specify a unit price. Each <SupplyDetail> composite must include
// either one or more prices, or a single <UnpricedItemType> element (see P.26.42).
// Where multiple prices are specified for a product, each price option should
// specify a distinct combination of its terms of trade and the group of end
// customers it is applicable to, any relevant price conditions, periods of
// time, currency and territory etc, so that the data recipient can clearly
// select the correct pricing option. If, under a particular combination, the
// product is free of charge or its price is not yet set, an <UnpricedItemType>
// element (P.26.70a) must be used in place of a <PriceAmount>. Each pricing
// option may optionally be given an identifier for use in subsequent revenue
// reporting or for other internal purposes.
type Price struct {
	PriceIdentifier        []PriceIdentifier
	PriceType              *PriceType
	PriceQualifier         *PriceQualifier
	PriceConstraint        []PriceConstraint
	PriceTypeDescription   []PriceTypeDescription
	PricePer               *PricePer
	PriceCondition         []PriceCondition
	MinimumOrderQuantity   *MinimumOrderQuantity
	BatchBonus             []BatchBonus
	DiscountCoded          []DiscountCoded
	Discount               []Discount
	PriceStatus            *PriceStatus
	PriceAmount            PriceAmount
	CurrencyZone           *CurrencyZone
	ComparisonProductPrice []ComparisonProductPrice
	PriceDate              []PriceDate
	Tax                    []Tax
	CurrencyCode           *CurrencyCode
	Territory              *Territory
	PrintedOnProduct       *PrintedOnProduct
	PositionOnProduct      *PositionOnProduct
	PriceCoded             *PriceCoded
	UnpricedItemType       *UnpricedItemType
	GeneralAttributes
}

// PriceAmount represents the amount of a price. Optional and non-repeating,
// but each occurrence of the <Price> composite must include either a
// <PriceAmount> or a <PriceCoded> composite, with optional tax details, or an
// <UnpricedItemType> element. Note that free-of-charge products must use
// <UnpricedItemType> rather than a zero price.
type PriceAmount struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

// PriceCode represents a price code from the scheme specified in the
// <PriceCodeType> element. Mandatory in each occurrence of the <PriceCoded>
// composite, and non-repeating.
type PriceCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// PriceCoded represents an optional group of data elements to carry a price
// that is expressed as one of a discrete set of price points, tiers or bands,
// rather than actual currency amounts. Each occurrence of the <Price>
// composite must include either a <PriceAmount> or a <PriceCoded> composite,
// with optional tax details, or an <UnpricedItemType> element.
type PriceCoded struct {
	PriceCodeType     PriceCodeType
	PriceCodeTypeName *PriceCodeTypeName
	PriceCode         PriceCode
	GeneralAttributes
}

// PriceCodeType represents an ONIX code identifying the scheme from which the
// value in the <PriceCode> element is taken. Mandatory in an occurrence of the
// <PriceCoded> composite, and non-repeating.
type PriceCodeType struct {
	Value string `xml:",innerxml"` // List179
	GeneralAttributes
}

// PriceCodeTypeName represents a name which identifies a proprietary price
// code type. Must be used when, and only when, the code in the
// <PriceCodeType> element indicates a proprietary scheme, eg a retailer’s
// price banding scheme. Optional and non-repeating.
type PriceCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// PriceCondition represents a repeatable group of data elements which together
// specify a condition relating to a price.
type PriceCondition struct {
	PriceConditionType     PriceConditionType
	PriceConditionQuantity []PriceConditionQuantity
	ProductIdentifier      []ProductIdentifier
	GeneralAttributes
}

// PriceConditionQuantity represents an optional and repeatable group of data
// elements which together specify a price condition quantity, for example
// a minimum number of copies, or a period of time for which updates are
// supplied or must be purchased.
type PriceConditionQuantity struct {
	PriceConditionQuantityType PriceConditionQuantityType
	Quantity                   Quantity
	QuantityUnit               QuantityUnit
	GeneralAttributes
}

// PriceConditionQuantityType represents an ONIX code identifying a type of
// price condition quantity. Mandatory in each occurrence of the
// <PriceConditionQuantity> composite, and non-repeating.
type PriceConditionQuantityType struct {
	Value string `xml:",innerxml"` // List168
	GeneralAttributes
}

// PriceConditionType represents an ONIX code identifying a type of price
// condition. Mandatory in each occurrence of the <PriceCondition> composite,
// and non-repeating.
type PriceConditionType struct {
	Value string `xml:",innerxml"` // List167
	GeneralAttributes
}

// PriceDate represents an optional and repeatable group of data elements which
// together specify a date associated with a price.
type PriceDate struct {
	PriceDateRole PriceDateRole
	DateFormat    *DateFormat
	Date          Date
	GeneralAttributes
}

// PriceDateRole represents an ONIX code indicating the significance of the
// date. Mandatory in each occurrence of the <PriceDate> composite, and non-
// repeating.
type PriceDateRole struct {
	Value string `xml:",innerxml"` // List173
	GeneralAttributes
}

// PriceIdentifier represents an optional and repeatable group of elements
// that provide an identifier for a particular price. For products that may be
// available at potentially many different prices, to different groups of
// purchasers or under different terms and conditions, this identifier may then
// be used in subsequent revenue reporting to specify which price the product
// was traded at.
//
// Note that the price identifier will always be proprietary and must be
// unique across multiple pricing options for one product, but need not be
// unique across all products, nor need it be the same across all products
// offered at the same price point or under the same terms.
type PriceIdentifier struct {
	PriceIDType PriceIDType
	IDTypeName  *IDTypeName
	IDValue     IDValue
	GeneralAttributes
}

// PriceIDType represents an ONIX code identifying the scheme from which the
// identifier in the <IDValue> element is taken. Mandatory in each occurrence
// of the <PriceIdentifier> composite, and non-repeating. There is no
// particular public ‘standard’ for price identifiers, so at present only
// proprietary identifiers may be specified.
type PriceIDType struct {
	Value string `xml:",innerxml"` // List217
	GeneralAttributes
}

// PricePer represents an ONIX code indicating the unit of product which is the
// basis for the price carried in an occurrence of the <Price> composite.
// Optional and non-repeating. Where the price applies to a copy of the whole
// product, this field is normally omitted.
type PricePer struct {
	Value string `xml:",innerxml"` // List60
	GeneralAttributes
}

// PriceQualifier represents an ONIX code which further specifies the type of
// price, eg member price, reduced price when purchased as part of a set.
// Optional, and repeatable when more than once qualifier applies.
type PriceQualifier struct {
	Value string `xml:",innerxml"` // List59
	GeneralAttributes
}

// PriceStatus represents an ONIX code which specifies the status of a price.
// Optional and non-repeating. If the field is omitted, the default
// ‘unspecified’ will apply.
type PriceStatus struct {
	Value string `xml:",innerxml"` // List61
	GeneralAttributes
}

// PriceType represents an ONIX code indicating the type of the comparison
// price in the <PriceAmount> element within the <ComparisonProductPrice>
// composite. Optional if a <DefaultPriceType> has been specified in the
// message header, otherwise mandatory. Non-repeating.
type PriceType struct {
	Value string `xml:",innerxml"` // List58
	GeneralAttributes
}

// PriceTypeDescription represents a free text which further describes the
// price type, qualifier, constraints and other parameters of the price.
// Optional, and repeatable if parallel descriptions are provided in multiple
// languages. The language attribute is optional for a single instance of
// <PriceTypeDescription>, but must be included in each instance if
// <PriceTypeDescription> is repeated in multiple languages. In the Netherlands
// and elsewhere, when the <PriceQualifier> code identifies a ‘voucher price’,
// the <PriceTypeDescription> should give the ‘EAN action number’ that
// identifies the offer.
type PriceTypeDescription struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// PriceConstraint represents an optional and repeatable group of data elements
// which together describe a contractual term or constraint (or the absence of
// such a constraint) that forms part of the commercial offer for a product.
// The Price constraint composite has the same structure as
// <EpubUsageConstraint>. Use <EpubUsageConstraint> for constraints that limit
// the user experience of the product, whether or not they are enforced by
// technical protection measures (DRM). Use <PriceConstraint> where a single
// product is available under multiple terms and conditions (ie multiple
// commercial offers for the same product).
type PriceConstraint struct {
	PriceConstraintType   PriceConstraintType
	PriceConstraintStatus PriceConstraintStatus
	PriceConstraintLimit  []PriceConstraintLimit
	GeneralAttributes
}

// PriceConstraintLimit represents an optional and repeatable group of data
// elements which together specify a quantitative limit on a particular type of
// contractual term or constraint.
type PriceConstraintLimit struct {
	Quantity            Quantity
	PriceConstraintUnit PriceConstraintUnit
	GeneralAttributes
}

// PriceConstraintStatus represents An ONIX code specifying the status of a
// contractual term or constraint, eg permitted without limit, permitted with
// limit, prohibited. Mandatory in each occurrence of the <PriceConstraint>
// composite, and non-repeating.
type PriceConstraintStatus struct {
	Value string `xml:",innerxml"` // List146
	GeneralAttributes
}

// PriceConstraintType represents an ONIX code specifying a type of commercial
// term or constraint forming part of the commercial offer for a digital
// product. Mandatory in each occurrence of the <PriceConstraint> composite,
// and non-repeating.
type PriceConstraintType struct {
	Value string `xml:",innerxml"` // List230
	GeneralAttributes
}

// PriceConstraintUnit represents an ONIX code for a unit in which a maximum
// permitted quantity or limit is stated. Mandatory in each occurrence of the
// <PriceConstraintLimit> composite, and non-repeating.
type PriceConstraintUnit struct {
	Value string `xml:",innerxml"` // List147
	GeneralAttributes
}

// PrimaryContentType represents an ONIX code which indicates the primary or
// only content type included in a product. The element is intended to be used
// in particular for digital products, when the sender wishes to make it clear
// that one of a number of content types (eg text, audio, video) is the primary
// type for the product. Other content types may be specified in the
// <ProductContentType>. Optional and non-repeating.
type PrimaryContentType struct {
	Value string `xml:",innerxml"` // List81
	GeneralAttributes
}

// PrimaryPart represents an empty element that allows a sender to identify a
// product part as the ‘primary’ part of a multiple-item product. For example,
// in a ‘book and toy’ or ‘book and DVD’ product, the book may be regarded as
// the primary part. Optional and non-repeating.
type PrimaryPart struct {
	GeneralAttributes
}

// PrintedOnProduct represents an ONIX code indicating whether the price in a
// <Price> composite is printed on the product. Optional and non-repeating.
// Omission of this element must not be interpreted as indicating that the
// price is not printed on the product.
type PrintedOnProduct struct {
	Value string `xml:",innerxml"` // List174
	GeneralAttributes
}

// Prize represents an optional and repeatable group of data elements which
// together describe a prize or award won by the contributor for a body of
// work (rather than for this or other specific works or products).
type Prize struct {
	PrizeName      []PrizeName
	PrizeYear      *PrizeYear
	PrizeCountry   *PrizeCountry
	PrizeCode      *PrizeCode
	PrizeStatement []PrizeStatement
	PrizeJury      []PrizeJury
	GeneralAttributes
}

// PrizeCode represents An ONIX code indicating the achievement of the product
// in relation to a prize or award, eg winner, runner-up, shortlisted. Optional
// and non-repeating.
type PrizeCode struct {
	Value string `xml:",innerxml"` // List41
	GeneralAttributes
}

// PrizeCountry represents An ISO standard code identifying the country in
// which a prize or award is given. Optional and non-repeating.
type PrizeCountry struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

// PrizeJury represents a free text listing members of the jury that awarded
// the prize. Optional, and repeatable if the text is provided in more than one
// language. The language attribute is optional for a single instance of
// <PrizeJury>, but must be included in each instance if <PrizeJury> is repeated.
type PrizeJury struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// PrizeName represents the name of a prize or award which the product or work
// has received. Mandatory in each occurrence of the <Prize> composite, and
// repeatable to provide a parallel award name in multiple languages. The
// language attribute is optional for a single instance of <PrizeName>, but
// must be included in each instance if <PrizeName> is repeated.
type PrizeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// PrizeStatement represents a short free-text description of the prize or
// award, intended primarily for display. Optional, and repeatable if the text
// is provided in more than one language. The language attribute is optional
// for a single instance of <PrizeStatement>, but must be included in each
// instance if <PrizeStatement> is repeated. <PrizeStatement> is intended for
// display purposes only. When used, a <PrizeStatement> must be complete in
// itself, ie it should not be treated as merely supplementary to other
// elements within the <Prize> composite. Nor should <PrizeStatement> be
// supplied instead of those other elements – at minimum, the <PrizeCode>
// element, and whenever appropriate the <PrizeYear> element should be supplied.
type PrizeStatement struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// PrizeYear represents the year in which a prize or award was given. Optional
// and non-repeating.
type PrizeYear struct {
	Value string `xml:",innerxml"` // dt.Year
	GeneralAttributes
}

// Product represents an ONIX product record. A product is described by a group
// of data elements beginning with an XML label <Product> and ending with an
// XML label </Product>. The entire group of data elements which is enclosed
// between these two labels constitutes an ONIX product record. The product
// record is the fundamental unit within an ONIX Product Information message.
// In almost every case, each product record describes an individually tradable
// item; and in all circumstances, each tradable item identified by a
// recognized product identifier should be described by one, and only one,
// ONIX product record.
//
// In ONIX 3.0, a <Product> record has a mandatory ‘preamble’ comprising data
// element Groups P.1 and P.2, and carrying data that identifies the record and
// the product to which it refers. This is followed by up to six ‘blocks’, each
// optional, some of which are repeatable.
type Product struct {
	RecordReference        RecordReference
	NotificationType       NotificationType
	DeletionText           []DeletionText
	RecordSourceType       *RecordSourceType
	RecordSourceIdentifier []RecordSourceIdentifier
	RecordSourceName       *RecordSourceName
	ProductIdentifier      []ProductIdentifier
	Barcode                []Barcode
	DescriptiveDetail      *DescriptiveDetail
	CollateralDetail       *CollateralDetail
	ContentDetail          *ContentDetail
	PublishingDetail       *PublishingDetail
	RelatedMaterial        *RelatedMaterial
	ProductSupply          []ProductSupply
	GeneralAttributes
}

// ProductAvailability represents an ONIX code indicating the availability of a
// product from a supplier. Mandatory in each occurrence of the <SupplyDetail>
// composite, and non-repeating.
type ProductAvailability struct {
	Value string `xml:",innerxml"` // List65
	GeneralAttributes
}

// ProductClassification represents a repeatable group of data elements which
// together define a product classification (not to be confused with a subject
// classification). The intended use is to enable national or international
// trade classifications (also known as commodity codes) to be carried in an
// ONIX record. Optional.
type ProductClassification struct {
	ProductClassificationType ProductClassificationType
	ProductClassificationCode ProductClassificationCode
	Percent                   *Percent
	GeneralAttributes
}

// ProductClassificationCode represents a classification code from the scheme
// specified in <ProductClassificationType>. Mandatory in each occurrence of
// the <ProductClassification> composite, and non-repeating.
type ProductClassificationCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ProductClassificationType represents an ONIX code identifying the scheme from
// which the identifier in <ProductClassificationCode> is taken. Mandatory in
// each occurrence of the <ProductClassification> composite, and non-repeating.
type ProductClassificationType struct {
	Value string `xml:",innerxml"` // List9
	GeneralAttributes
}

// ProductComposition represents an ONIX code which indicates whether a product
// consists of a single item or multiple items. Mandatory in an occurrence of
// <DescriptiveDetail>, and non-repeating.
type ProductComposition struct {
	Value string `xml:",innerxml"` // List2
	GeneralAttributes
}

// ProductContact represents a group of data elements which together specify an
// organization (which may or may not be the publisher) responsible for dealing
// with enquiries related to the product.
type ProductContact struct {
	ProductContactRole       ProductContactRole
	ContactName              *ContactName
	EmailAddress             *EmailAddress
	ProductContactIdentifier []ProductContactIdentifier
	ProductContactName       *ProductContactName
	GeneralAttributes
}

// ProductContactIdentifier represents a group of data elements which together
// define an identifier of the product contact. The composite is optional, and
// repeatable if more than one identifier of different types is sent; but
// either a <ProductContactName> or a <ProductContactIdentifier> must be
// included.
type ProductContactIdentifier struct {
	ProductContactIDType ProductContactIDType
	IDTypeName           *IDTypeName
	IDValue              IDValue
	GeneralAttributes
}

// ProductContactIDType represents an ONIX code identifying a scheme from which
// an identifier in the <IDValue> element is taken. Mandatory in each
// occurrence of the <ProductContactIdentifier> composite, and non-repeating.
type ProductContactIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// ProductContactName represents the name of the product contact organization,
// which should always be stated in a standard form. Optional and non-
// repeating; but either a <ProductContactName> element or a
// <ProductContactIdentifier> composite must be included.
type ProductContactName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ProductContactRole represents an ONIX code which identifies the role played
// by the product contact in relation to the product – for example answering
// enquiries related to sales or to promotion.
type ProductContactRole struct {
	Value string `xml:",innerxml"` // List198
	GeneralAttributes
}

// ProductContentType represents an ONIX code which indicates a content type
// included in a product. The element is intended to be used in particular for
// digital products, to specify content types other than the primary type, or
// to list content types when none is singled out as the primary type.
// Optional and repeatable.
type ProductContentType struct {
	Value string `xml:",innerxml"` // List81
	GeneralAttributes
}

// ProductForm represents an ONIX code which indicates the primary form of a
// product part. Mandatory in each occurrence of <ProductPart>, and non-
// repeating.
type ProductForm struct {
	Value string `xml:",innerxml"` // List150
	GeneralAttributes
}

// ProductFormDescription represents a text describing the product form. If
// product form codes do not adequately describe the contained item, a short
// text description may be added. Optional and repeatable. The language
// attribute is optional for a single instance of <ProductFormDescription>,
// but must be included in each instance if <ProductFormDescription> is
// repeated.
type ProductFormDescription struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ProductFormDetail represents an ONIX code which provides added detail of
// the medium and/or format of a related product. Optional and repeatable.
type ProductFormDetail struct {
	Value string `xml:",innerxml"` // List175
	GeneralAttributes
}

// ProductFormFeature represents a repeatable group of data elements which
// together describe an aspect of product form that is too specific to be
// covered in the <ProductForm> and <ProductFormDetail> elements. Optional.
type ProductFormFeature struct {
	ProductFormFeatureType        ProductFormFeatureType
	ProductFormFeatureValue       *ProductFormFeatureValue
	ProductFormFeatureDescription []ProductFormFeatureDescription
	GeneralAttributes
}

// ProductFormFeatureDescription represents a product form feature description
// text. If the <ProductFormFeatureType> requires free text rather than a code
// value, or if the code in <ProductFormFeatureValue> does not adequately
// describe the feature, a short text description may be added. Optional, and
// repeatable to provide parallel descriptive text in multiple languages. The
// language attribute is optional for a single instance of
// <ProductFormFeatureDescription>, but must be included in each instance if
// <ProductFormFeatureDescription> is repeated to provide parallel text in
// multiple languages.
type ProductFormFeatureDescription struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ProductFormFeatureType represents an ONIX code which specifies the feature
// described by an instance of the <ProductFormFeature> composite, eg binding
// color. Mandatory in each occurrence of the composite, and non-repeating.
type ProductFormFeatureType struct {
	Value string `xml:",innerxml"` // List79
	GeneralAttributes
}

// ProductFormFeatureValue represents a controlled value that describes a
// product form feature. Presence or absence of this element depends on the
// <ProductFormFeatureType>, since some product form features (eg thumb index)
// do not require an accompanying value, while others (eg text font) require
// free text in <ProductFormFeatureDescription>; and others may have both code
// and free text. Non-repeating.
type ProductFormFeatureValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ProductIdentifier represents a repeatable group of data elements which
// together define an identifier of a product in accordance with a specified
// scheme, used here to carry the product identifier of a product part.
// Optional, but required when an occurrence of <ProductPart> specifies an
// individual item with its own identifier.
type ProductIdentifier struct {
	ProductIDType ProductIDType
	IDTypeName    *IDTypeName
	IDValue       IDValue
	GeneralAttributes
}

// ProductIDType represents an ONIX code identifying the scheme from which
// the identifier in the <IDValue> element is taken. Mandatory in each
// occurrence of the <ProductIdentifier> composite, and non-repeating.
type ProductIDType struct {
	Value string `xml:",innerxml"` // List5
	GeneralAttributes
}

// ProductPackaging represents an ONIX code which indicates the type of
// packaging used for the product part. Optional, and not repeatable.
type ProductPackaging struct {
	Value string `xml:",innerxml"` // List80
	GeneralAttributes
}

// ProductPart represents a repeatable group of data elements which together
// describe an item which is part of or contained within a multiple-item
// product or a trade pack. The composite must be used with all multiple-item
// products to specify (for example) the item(s) and item quantities included
// in a combined book plus audiobook product. a multi-volume set, a filled
// dumpbin, or a classroom pack. In other cases, where parts are not
// individually identified, it is used to state the product form(s) and the
// quantity or quantities of each form contained within the product.
//
// Each instance of the <ProductPart> composite must carry a <ProductForm> code
// and a quantity, even if the quantity is ‘1’. If the composite refers to a
// number of copies of a single item, the quantity must be sent as
// <NumberOfCopies>, normally accompanied by a <ProductIdentifier>. If the
// composite refers to a number of different items of the same form, without
// identifying them individually, the quantity must be sent as
// <NumberOfItemsOfThisForm>.
type ProductPart struct {
	PrimaryPart             *PrimaryPart
	ProductIdentifier       []ProductIdentifier
	ProductForm             ProductForm
	ProductFormDetail       []ProductFormDetail
	ProductFormFeature      []ProductFormFeature
	ProductPackaging        *ProductPackaging
	ProductFormDescription  []ProductFormDescription
	ProductContentType      []ProductContentType
	CountryOfManufacture    *CountryOfManufacture
	NumberOfItemsOfThisForm NumberOfItemsOfThisForm
	NumberOfCopies          *NumberOfCopies
	GeneralAttributes
}

// ProductRelationCode represents an ONIX code which identifies the nature of
// the relationship between two products, eg ‘replaced-by’. Mandatory in each
// occurrence of the <RelatedProduct> composite, and repeatable where the
// related product has multiple types of relationship to the product described.
type ProductRelationCode struct {
	Value string `xml:",innerxml"` // List51
	GeneralAttributes
}

// ProductSupply represents the product supply block covers data element Groups
// P.24 to P.26, specifying a market, the publishing status of the product in
// that market, and the supply arrangements for the product in that market. The
// block is repeatable to describe multiple markets. At least one occurrence is
// expected in a <Product> record unless the <NotificationType> in Group P.1
// indicates that the record is an update notice which carries only those blocks
// in which changes have occurred.
type ProductSupply struct {
	Market                 []Market
	MarketPublishingDetail *MarketPublishingDetail
	SupplyDetail           []SupplyDetail
	GeneralAttributes
}

// ProfessionalAffiliation represents an optional and repeatable group of data
// elements which together identify a contributor’s professional position and/or
// affiliation, allowing multiple positions and affiliations to be specified.
type ProfessionalAffiliation struct {
	ProfessionalPosition []ProfessionalPosition
	Affiliation          *Affiliation
	GeneralAttributes
}

// ProfessionalPosition represents a professional position held by a
// contributor to the product at the time of its creation. Optional, and
// repeatable to provide  parallel text in multiple languages. The language
// attribute is optional for a single instance of <ProfessionalPosition>, but
// must be included in each instance if <ProfessionalPosition> is repeated.
type ProfessionalPosition struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// PromotionCampaign represents a free text describing the promotion and
// adverting campaign for the product. Optional, and repeatable if parallel
// text is provided in multiple languages. The language attribute is optional
// for a single instance of <PromotionCampaign>, but must be included in each
// instance if <PromotionCampaign> is repeated.
type PromotionCampaign struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// PromotionContact represents a free text giving the name, department, phone
// number, e-mail address etc for a promotional contact person for the product.
// Optional and non-repeating. Deprecated, in favor of supplying this
// information via the <ProductContact> composite.
type PromotionContact struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// Proximity represents an ONIX code which specifies the precision of the stock
// quantity on hand. Optional, non-repeating and may only be used if an <OnHand>
// value is specified.
type Proximity struct {
	Value string `xml:",innerxml"` // List215
	GeneralAttributes
}

// Publisher represents a repeatable group of data elements which together
// identify an entity which is associated with the publishing of a product. The
// composite allows additional publishing roles to be introduced without adding
// new fields. Each occurrence of the composite must carry a publishing role
// code and either a name identifier code or a name or both.
type Publisher struct {
	PublishingRole      PublishingRole
	PublisherName       *PublisherName
	Funding             []Funding
	Website             []Website
	PublisherIdentifier []PublisherIdentifier
	GeneralAttributes
}

// PublisherIdentifier represents a group of data elements which together
// define the identifier of a publisher name. Optional and repeatable, but
// mandatory if the <Publisher> composite does not carry a <PublisherName>.
type PublisherIdentifier struct {
	PublisherIDType PublisherIDType
	IDTypeName      *IDTypeName
	IDValue         IDValue
	GeneralAttributes
}

// PublisherIDType represents an ONIX code which identifies the scheme from
// which the value in the <IDValue> element is taken. Mandatory in each
// occurrence of the <PublisherIdentifier> composite.
type PublisherIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// PublisherName represents the name of an entity associated with the
// publishing of a product. Mandatory if there is no publisher identifier in an
// occurrence of the <Publisher> composite, and optional if a publisher
// identifier is included. Non-repeating.
type PublisherName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// PublisherRepresentative represents a repeatable group of data elements which
// together identify a publisher representative in a specified market. Optional,
// and repeated only if the publisher has two or more representatives.
type PublisherRepresentative struct {
	AgentRole       AgentRole
	TelephoneNumber []TelephoneNumber
	FaxNumber       []FaxNumber
	EmailAddress    []EmailAddress
	Website         []Website
	AgentIdentifier []AgentIdentifier
	AgentName       *AgentName
	GeneralAttributes
}

// PublishingDate represents a repeatable group of data elements which together
// specify a date associated with the publishing of the product. Optional, but
// a date of publication must be specified either here (as a ‘global’ pubdate)
// or in <MarketPublishingDetail> (P.25). Other dates related to the publishing
// of a product can be sent in further repeats.
type PublishingDate struct {
	PublishingDateRole PublishingDateRole
	DateFormat         *DateFormat
	Date               Date
	GeneralAttributes
}

// PublishingDateRole represents an ONIX code indicating the significance of
// the date, eg pubdate, announcement date, latest reprint date. Mandatory in
// each occurrence of the <PublishingDate> composite, and non-repeating.
type PublishingDateRole struct {
	Value string `xml:",innerxml"` // List163
	GeneralAttributes
}

// PublishingDetail represents a block of data element Groups P.19 to P.21,
// carrying information on the publisher(s), ‘global’ publishing status, and
// rights attaching to a product. The block as a whole is non-repeating. It is
// mandatory in any <Product> record unless the <NotificationType> in Group P.1
// indicates that the record is an update notice which carries only those
// blocks in which changes have occurred.
type PublishingDetail struct {
	Imprint              []Imprint
	Publisher            []Publisher
	CityOfPublication    []CityOfPublication
	CountryOfPublication *CountryOfPublication
	PublishingStatus     []PublishingStatus
	PublishingDate       []PublishingDate
	SalesRights          []SalesRights
	ROWSalesRightsType   *ROWSalesRightsType
	ProductContact       []ProductContact
	LatestReprintNumber  *LatestReprintNumber
	CopyrightStatement   []CopyrightStatement
	SalesRestriction     []SalesRestriction
	GeneralAttributes
}

// PublishingRole represents an ONIX code which identifies a role played by an
// entity in the publishing of a product. Mandatory in each occurrence of the
// <Publisher> composite, and non-repeating.
type PublishingRole struct {
	Value string `xml:",innerxml"` // List45
	GeneralAttributes
}

// PublishingStatus represents an ONIX code which identifies the status of a
// published product. Optional and non-repeating, but required if publishing
// status is not identified at market level in <MarketPublishingDetail> (P.25).
// Where the element is sent by a sender who is not the publisher, based on
// information that has been previously supplied by the publisher, it is
// strongly recommended that it should carry a datestamp attribute to indicate
// its likely reliability. See Section 1 for further details of the datestamp
// attribute.
type PublishingStatus struct {
	Value string `xml:",innerxml"` // List64
	GeneralAttributes
}

// PublishingStatusNote represents a free text that describes the status of a
// published product, when the code in <PublishingStatus> is insufficient.
// Optional, but when used, must be accompanied by the <PublishingStatus>
// element. Repeatable if parallel notes are provided in multiple languages.
// The language attribute is optional for a single instance of
// <PublishingStatusNote>, but must be included in each instance if
// <PublishingStatusNote> is repeated.
type PublishingStatusNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// Quantity is a numeric value representing the maximum permitted quantity or
// limit of a particular type of constraint. Mandatory in each occurrence of
// the <PriceConstraintLimit> composite, and non-repeating.
type Quantity struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

// QuantityUnit represents an ONIX code value specifying the unit in which a
// price condition quantity is stated. Mandatory in each occurrence of the
// <PriceConditionQuantity> composite, and non-repeating.
type QuantityUnit struct {
	Value string `xml:",innerxml"` // List169
	GeneralAttributes
}

// Rate represents the stock depletion rate (as a number of copies, rounded to
// the nearest integer), measured according to the metric in <VelocityMetric>.
type Rate struct {
	Value string `xml:",innerxml"` // dt.Integer
	GeneralAttributes
}

// ReviewRating represents an optional group of data elements which together
// specify a ‘star rating’ awarded as part of a review of the publication,
// used when <TextType> indicates the text is a review. Not repeatable.
type ReviewRating struct {
	Rating      Rating
	RatingLimit *RatingLimit
	RatingUnits []RatingUnits
	GeneralAttributes
}

// Rating represents the ‘star rating’ awarded as part of a review of the
// publication. Mandatory within an occurrence of the <ReviewRating> composite,
// and non-repeating.
type Rating struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

// RatingLimit represents the maximum possible rating that may be awarded as
// part of a review of the publication. Optional, but where used, it must be
// greater than or equal to the specified <Rating>.
type RatingLimit struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

// RatingUnits represents the ‘units’ used by a rating, eg stars, tomatoes etc.
// Optional, and repeatable to provide the units in multiple languages. The
// language attribute is optional for a single instance of <RatingUnits>, but
// must be included in each instance if <RatingUnits> is repeated.
type RatingUnits struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// RecordReference represents a record reference which will uniquely identify
// the Information record which you send out about that product, and which will
// remain as its permanent identifier every time you send an update. It doesn’t
// matter what reference you choose, provided that it is unique and permanent.
// This record reference doesn’t identify the product – even though you may
// choose to use the ISBN or another product identifier as a part or the whole
// of your record reference – it identifies your information record about the
// product, so that the person to whom you are sending an update can match it
// with what you have previously sent. A good way of generating references
// which are not part of a recognized product identification scheme but which
// can be guaranteed to be unique is to prefix a product identifier or a
// meaningless row ID from your internal database with a reversed Internet
// domain name which is registered to your organization (reversal prevents the
// record reference appearing to be a resolvable URL). Alternatively, use a
// UUID.
//
// This field is mandatory and non-repeating.
type RecordReference struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// RecordSourceIdentifier represents a repeatable group of data elements which
// together define an identifier of the organization which is the source of the
// ONIX record. Optional.
type RecordSourceIdentifier struct {
	RecordSourceIDType RecordSourceIDType
	IDTypeName         *IDTypeName
	IDValue            IDValue
	GeneralAttributes
}

// RecordSourceIDType represents an ONIX code identifying the scheme from which
// the identifier in the <IDValue> element is taken. Mandatory in each
// occurrence of the <RecordSourceIdentifier> composite, and non-repeating.
type RecordSourceIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// RecordSourceName represents the name of the party which issued the record,
// as free text. Optional and non-repeating, independently of the occurrence
// of any other field.
type RecordSourceName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// RecordSourceType represents an ONIX code which indicates the type of source
// which has issued the ONIX record. Optional and non-repeating, independently
// of the occurrence of any other field.
type RecordSourceType struct {
	Value string `xml:",innerxml"` // List3
	GeneralAttributes
}

// RegionCode represents an ONIX code identifying a region with which a
// contributor is particularly associated. Optional and non-repeatable. There
// must be an occurrence of either the <CountryCode> or the <RegionCode>
// elements in each occurrence of <ContributorPlace>. A region is an area which
// is not a country, but which is precisely defined in geographical terms, eg
// Northern Ireland, Australian Capital Territory. Note that US States have
// region codes, while US overseas territories have distinct ISO Country Codes.
type RegionCode struct {
	Value string `xml:",innerxml"` // List49
	GeneralAttributes
}

// RegionsIncluded represents one or more ONIX codes identifying regions
// included in the territory. A region is an area which is not a country, but
// which is precisely defined in geographical terms, eg World, Northern
// Ireland, Australian Capital Territory. Successive codes must be separated
// by spaces. Optional and non-repeating, but either <CountriesIncluded> or
// <RegionsIncluded> is mandatory in each occurrence of the <Territory>
// composite.
type RegionsIncluded struct {
	Value string `xml:",innerxml"` // RegionCodeList
	GeneralAttributes
}

// RegionsExcluded represents one or more ONIX codes identifying regions
// excluded from the territory. Successive codes must be separated by spaces.
// Optional and non-repeating, and can only occur if the <CountriesIncluded>
// element is also present (and specifies countries of which the excluded
// regions are a part), or if <RegionsIncluded> is present and includes a
// supra-national region code (such as ‘World’).
type RegionsExcluded struct {
	Value string `xml:",innerxml"` // RegionCodeList
	GeneralAttributes
}

// Reissue [DEPRECATED] represents An optional and non-repeating group of data
// elements which together specify that a product is to be reissued within the
// market to which the <SupplyDetail> composite applies. The entire <Reissue>
// composite is deprecated. Suppliers should supply information about planned
// reissues in other parts of the Product record – the date of a planned
// reissue in <PublishingDate> and/or <MarketDate>, and new collateral material
// alongside old collateral in Block 2 where they may be associated with
// appropriate end and start dates using <ContentDate>. The <Reissue> composite
// was (prior to deprecation) used only when the publisher intended to re-launch
//the product under the same ISBN. There are two possible cases:
//
// * When the product is unavailable during the period immediately before reissue. In this case, <ProductAvailability> should carry the value 33 for ‘unavailable, awaiting reissue’, and the ONIX       record can be updated to describe the reissued product as soon as details can be made available;
//
// * When the product is still available during the period up to the reissue date. In this case, the ONIX record should continue to describe the existing product and the <ProductAvailability> value should continue to record the product as ‘available’ (eg code 21) right up to the reissue date. At that date, the record should be updated to describe the reissued product, with the <ProductAvailability> value usually remaining unchanged.
type Reissue struct {
	ReissueDate        ReissueDate
	ReissueDescription *ReissueDescription
	Price              []Price
	SupportingResource []SupportingResource
	GeneralAttributes
}

// ReissueDate [DEPRECATED] represents the date on which the product will be
// reissued, or (after reissue) the date when it was last reissued. Mandatory
// in each occurrence of the <Reissue> composite, and non-repeating.
// Deprecated.
type ReissueDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// ReissueDescription [DEPRECATED] represents a text explaining the nature of
// the reissue. Optional and non-repeating. Deprecated.
type ReissueDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// RelatedMaterial represents a group of data element Groups P.22 and P.23,
// providing links to related works and related products. The block as a whole
// is optional and non-repeating.
type RelatedMaterial struct {
	RelatedWork    []RelatedWork
	RelatedProduct []RelatedProduct
	GeneralAttributes
}

// RelatedProduct represents an optional and repeatable group of data elements
// which together describe a product which has a specified relationship to the
// product described in the ONIX record.
type RelatedProduct struct {
	ProductRelationCode []ProductRelationCode
	ProductIdentifier   []ProductIdentifier
	GeneralAttributes
}

// RelatedWork represents a group of data elements which together describe a
// work which has a specified relationship to a content item. Optional and
// repeatable.
type RelatedWork struct {
	WorkRelationCode WorkRelationCode
	WorkIdentifier   []WorkIdentifier
	GeneralAttributes
}

// ReligiousText represents an optional, non-repeating, group of data elements
// which together describe features of the content of an edition of a religious
// text, and intended to meet the special needs of religious publishers and
// booksellers. The <ReligiousText> composite may carry either a <Bible>
// composite or a <ReligiousTextIdentifier> element accompanied by multiple
// repeats of the <ReligiousTextFeature> composite. This approach is adopted
// to enable other devotional texts to be included if need arises without
// requiring a new ONIX release.
type ReligiousText struct {
	ReligiousTextIdentifier ReligiousTextIdentifier
	ReligiousTextFeature    []ReligiousTextFeature
	Bible                   *Bible
	GeneralAttributes
}

// ReligiousTextFeature represents a repeatable group of data elements which
// together specify and describe a feature of a religious text. Mandatory if
// and only if <ReligiousTextIdentifier> is present.
type ReligiousTextFeature struct {
	ReligiousTextFeatureType        ReligiousTextFeatureType
	ReligiousTextFeatureCode        ReligiousTextFeatureCode
	ReligiousTextFeatureDescription []ReligiousTextFeatureDescription
	GeneralAttributes
}

// ReligiousTextFeatureCode represents an ONIX code specifying a feature
// described in the associated <ReligiousTextFeatureCode> element. Mandatory
// in each occurrence of the <ReligiousTextFeature> composite, and non-
// repeating
type ReligiousTextFeatureCode struct {
	Value string `xml:",innerxml"` // List90
	GeneralAttributes
}

// ReligiousTextFeatureDescription represents  a free text describing a feature
// that is not adequately defined by code values alone. Optional, and
// repeatable if parallel text is provided in multiple languages. The language
// attribute is optional for a single instance of
// <ReligiousTextFeatureDescription>, but must be included in each instance if
// <ReligiousTextFeatureDescription> is repeated.
type ReligiousTextFeatureDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// ReligiousTextFeatureType represents an ONIX code specifying a feature
// described in the associated <ReligiousTextFeatureCode> element. Mandatory
// in each occurrence of the <ReligiousTextFeature> composite, and non-
// repeating.
type ReligiousTextFeatureType struct {
	Value string `xml:",innerxml"` // List89
	GeneralAttributes
}

// ReligiousTextIdentifier represents an ONIX code indicating a religious text
// other than the Bible. Mandatory in each occurrence of the <ReligiousText>
// composite that does not include a <Bible> composite, and non-repeating.
type ReligiousTextIdentifier struct {
	Value string `xml:",innerxml"` // List88
	GeneralAttributes
}

// ReprintDetail represents a free text used to give details of the reprint
// history as part of the promotion of a book. Optional, and repeatable if
// parallel text is provided in multiple languages. The language attribute is
// optional for a single instance of <ReprintDetail>, but must be included in
// each instance if <ReprintDetail> is repeated. (For compatibility purposes,
// <ReprintDetail> is also repeatable – without the language attribute, or with
// the same language attribute – to give information about successive
// reprintings, but this is deprecated in favor of a single <ReprintDetail>
// instance [or a single instance per language] and use of the XHTML <dl> list
// structure.)
type ReprintDetail struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// ResourceContentType represents an ONIX code indicating the type of content
// carried in a supporting resource. Mandatory in each occurrence of the
// <SupportingResource> composite, and non-repeating.
type ResourceContentType struct {
	Value string `xml:",innerxml"` // List158
	GeneralAttributes
}

// ResourceFeature represents a repeatable group of data elements which
// together describe a feature of a supporting resource which is common to all
// versions in which the resource is offered. Optional.
type ResourceFeature struct {
	ResourceFeatureType ResourceFeatureType
	FeatureValue        *FeatureValue
	FeatureNote         []FeatureNote
	GeneralAttributes
}

// ResourceFeatureType represents an ONIX code which specifies the feature
// described by an instance of the <ResourceFeature> composite.
// Mandatory in each occurrence of the composite, and non-repeating.
type ResourceFeatureType struct {
	Value string `xml:",innerxml"` // List160
	GeneralAttributes
}

// ResourceForm represents an ONIX code indicating the form of a version of a
// supporting resource. Mandatory in each occurrence of the <ResourceVersion>
// composite, and non-repeating.
type ResourceForm struct {
	Value string `xml:",innerxml"` // List161
	GeneralAttributes
}

// ResourceLink represents a URI which provides a link to a supporting
// resource. Mandatory in each occurrence of the <ResourceVersion> composite,
// and repeatable if the resource can be linked in more than one way, eg by
// URL or DOI, or where a supporting resource is available in multiple parallel
// languages. Where multiple languages are used, all repeats must carry the
// language attribute.
type ResourceLink struct {
	Value string `xml:",innerxml"` // dt.NonEmptyURI
	GeneralAttributes
	LanguageAttribute
}

// ResourceMode represents an ONIX code indicating the mode of the supporting
// resource, eg audio, video. Mandatory in each occurrence of the
// <SupportingResource> composite, and non-repeating.
type ResourceMode struct {
	Value string `xml:",innerxml"` // List159
	GeneralAttributes
}

// ResourceVersion represents a repeatable group of data elements which
// together describe a version of a supporting resource, for example a
// particular format of a cover image. At least one instance is mandatory in
// each occurrence of the <SupportingResource> composite.
type ResourceVersion struct {
	ResourceForm           ResourceForm
	ResourceVersionFeature []ResourceVersionFeature
	ResourceLink           []ResourceLink
	ContentDate            []ContentDate
	GeneralAttributes
}

// ResourceVersionFeature represents a repeatable group of data elements which
// together describe a feature of a supporting resource which is specific to a
// version in which the resource is offered. Formally optional, but it is
// unlikely that a supporting resource version could be adequately described
// without specifying some of its features.
type ResourceVersionFeature struct {
	ResourceVersionFeatureType ResourceVersionFeatureType
	FeatureValue               *FeatureValue
	FeatureNote                []FeatureNote
	GeneralAttributes
}

// ResourceVersionFeatureType represents an ONIX code which specifies a feature
// described by an instance of the <ResourceVersionFeature> composite.
// Mandatory in each occurrence of the composite, and non-repeating.
type ResourceVersionFeatureType struct {
	Value string `xml:",innerxml"` // List162
	GeneralAttributes
}

// ROWSalesRightsType represents an ONIX code describing the sales rights
// applicable in territories not specifically associated with a sales right
// within an occurrence of the <SalesRights> composite. Optional, but
// required in all cases where no sales rights type is associated with the
// region ‘WORLD’, and in all cases where a sales rights type is associated
// with ‘WORLD’ but with exclusions that are not themselves associated with a
// sales rights type. Not repeatable. Note the value ‘00’ should be used to
// indicate where sales rights are genuinely unknown, or are unstated for any
// reason – in this case, data recipients must not assume anything about the
// rights that are applicable.
type ROWSalesRightsType struct {
	Value string `xml:",innerxml"` // List46
	GeneralAttributes
}

// ReturnsCode represents a returns conditions code from the scheme specified
// in <ReturnsCodeType>. Mandatory in each occurrence of the
// <ReturnsConditions> composite, and non-repeating.
type ReturnsCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// ReturnsCodeType represents an ONIX code identifying the scheme from which
// the returns conditions code in <ReturnsCode> is taken. Mandatory in each
// occurrence of the <ReturnsConditions> composite, and non-repeating.
type ReturnsCodeType struct {
	Value string `xml:",innerxml"` // List53
	GeneralAttributes
}

// ReturnsCodeTypeName represents a name which identifies a proprietary returns
// code scheme. Must be used when, and only when, the code in the
// <ReturnsCodeType> element indicates a proprietary scheme, eg a wholesaler’s
// own code. Optional and non-repeating.
type ReturnsCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ReturnsConditions represents an optional and repeatable group of data
// elements which together allow the supplier’s returns conditions to be
// specified in coded form.
type ReturnsConditions struct {
	ReturnsCodeType     ReturnsCodeType
	ReturnsCodeTypeName *ReturnsCodeTypeName
	ReturnsCode         ReturnsCode
	ReturnsNote         []ReturnsNote
	GeneralAttributes
}

// ReturnsNote represents a free text note explaining the returns
// conditions or special returns instructions, where the code alone is not
// sufficient. Optional, and repeatable if parallel text is provided in
// multiple languages. The language attribute is optional for a single instance
// of <ReturnsNote>, but must be included in each instance if <ReturnsNote>
// is repeated.
type ReturnsNote struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SalesOutlet represents an optional and repeatable group of data elements
// which together identify a sales outlet to which a restriction is linked.
// Each occurrence of the composite must include a <SalesOutletIdentifier>
// composite or a <SalesOutletName> or both.
type SalesOutlet struct {
	SalesOutletIdentifier []SalesOutletIdentifier
	SalesOutletName       *SalesOutletName
	GeneralAttributes
}

// SalesOutletIdentifier represents an optional and repeatable group of data
// elements which together represent a coded identification of an organization,
// used here to identify a sales outlet.
type SalesOutletIdentifier struct {
	SalesOutletIDType SalesOutletIDType
	IDTypeName        *IDTypeName
	IDValue           IDValue
	GeneralAttributes
}

// SalesOutletIDType represents an ONIX code which identifies the scheme from
// which the value in the <IDValue> element is taken. Mandatory in each
// occurrence of the <SalesOutletIdentifier> composite, and non-repeating.
type SalesOutletIDType struct {
	Value string `xml:",innerxml"` // List102
	GeneralAttributes
}

// SalesOutletName represents the name of a wholesale or retail sales outlet
// to which a sales restriction is linked. Non-repeating.
type SalesOutletName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SalesRestriction [DEPRECATED] represents a group of data elements which
// together identify a non-territorial sales restriction which a publisher
// applies to a product. Optional and repeatable, but deprecated in this
// context, in favor of using <SalesRestriction> within <SalesRights>
// (P.21.5a to P.21.5h).
type SalesRestriction struct {
	SalesRestrictionType SalesRestrictionType
	SalesOutlet          []SalesOutlet
	SalesRestrictionNote []SalesRestrictionNote
	StartDate            *StartDate
	EndDate              *EndDate
	GeneralAttributes
}

// SalesRestrictionNote represents a free text field describing an
// ‘unspecified’ restriction, or giving more explanation of a coded restriction
// type. Optional, and repeatable if parallel notes are provided in multiple
// languages. The language attribute is optional for a single instance of
// <SalesRestrictionNote>, but must be included in each instance if
// <SalesRestrictionNote> is repeated.
type SalesRestrictionNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// SalesRestrictionType represents an ONIX code which identifies a non-
// territorial sales restriction. Mandatory in each occurrence of the
// <SalesRestriction> composite, and non-repeating.
type SalesRestrictionType struct {
	Value string `xml:",innerxml"` // List71
	GeneralAttributes
}

// SalesRights represents an optional and repeatable group of data
// elements which together identify territorial sales rights which a
// publisher chooses to exercise in a product. When specifying a territory in
// which the product is not for sale, the publisher and product ID for an
// edition which is available in the specified territory can optionally be
// included. (In previous releases, this functionality was provided in a
// <NotForSale> composite, which is now redundant and has been deleted.)
type SalesRights struct {
	SalesRightsType   SalesRightsType
	Territory         *Territory
	SalesRestriction  []SalesRestriction
	ProductIdentifier []ProductIdentifier
	PublisherName     *PublisherName
	GeneralAttributes
}

// SalesRightsType represents an ONIX code which identifies the type of sales
// right or exclusion which applies in the territories which are associated
// with it. Mandatory in each occurrence of the <SalesRights> composite, and
// non-repeating. Values include: for sale with exclusive rights, for sale with
// non-exclusive rights, not for sale.
type SalesRightsType struct {
	Value string `xml:",innerxml"` // List46
	GeneralAttributes
}

// ScriptCode represents a code identifying the script in which the language is
// represented. Optional and non-repeating.
type ScriptCode struct {
	Value string `xml:",innerxml"` // List121
	GeneralAttributes
}

// Sender represents a group of data elements which together specify the sender
// of an ONIX for Books message. Mandatory in any ONIX for Books message, and
// non-repeating.
type Sender struct {
	SenderName       *SenderName
	ContactName      *ContactName
	EmailAddress     *EmailAddress
	SenderIdentifier []SenderIdentifier
	GeneralAttributes
}

// SenderIDType represents an ONIX code identifying a scheme from which an
// identifier in the <IDValue> element is taken. Mandatory in each occurrence
// of the <SenderIdentifier> composite, and non-repeating.
type SenderIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

// SenderIdentifier represents a group of data elements which together define an
// identifier of the sender. The composite is optional, and repeatable if more
// than one identifier of different types is sent; but either a <SenderName> or
// a <SenderIdentifier> must be included.
type SenderIdentifier struct {
	SenderIDType SenderIDType
	IDTypeName   *IDTypeName
	IDValue      IDValue
	GeneralAttributes
}

// SenderName represents the name of the sender organization, which should
// always be stated in a standard form agreed with the addressee. Optional and
// non-repeating; but either a <SenderName> element or a <SenderIdentifier>
// composite must be included.
type SenderName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// SentDateTime represents the date on which the message is sent. Optionally,
// the time may be added, using the 24-hour clock, with an explicit indication
// of the time zone if required, in a format based on ISO 8601. Mandatory and
// non-repeating.
type SentDateTime struct {
	Value string `xml:",innerxml"` // dt.DateOrDateTime
	GeneralAttributes
}

// SequenceNumber represents a number which specifies a single overall sequence
// of title elements, which is the preferred order for display of the various
// title elements when constructing a complete title. Optional and non-
// repeating. It is strongly recommended that each occurrence of the
// <TitleElement> composite should carry a <SequenceNumber>.
type SequenceNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

// SourceName represents the name of a collection source. If the
// <CollectionType> code indicates an ascribed collection (ie a collection
// which has been identified and described by a supply chain organization other
// than the publisher), this element may be used to carry the name of the
// organization responsible. Optional and non-repeating.
type SourceName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// SourceTitle represents the title of a publication from which the text sent
// in the <Text> element was taken, eg if it is a review quote. Optional, and
// repeatable to provide the title in multiple languages. The language
// attribute is optional for a single instance of <SourceTitle>, but must be
// included in each instance if <SourceTitle> is repeated.
type SourceTitle struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SourceType represents an ONIX code indicating the type of source from which
// the cited material originated, eg radio, TV. Optional, and non-repeating.
type SourceType struct {
	Value string `xml:",innerxml"` // List157
	GeneralAttributes
}

// StartDate represents the date from which a sales restriction is effective.
// Optional and non-repeating.
type StartDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// Stock represents a repeatable group of data elements which together specify
// a quantity of stock and, where a supplier has more than one warehouse, a
// supplier location. Optional
type Stock struct {
	LocationIdentifier []LocationIdentifier
	LocationName       []LocationName
	OnOrderDetail      []OnOrderDetail
	Velocity           []Velocity
	OnHand             OnHand
	Proximity          *Proximity
	StockQuantityCoded []StockQuantityCoded
	GeneralAttributes
}

// StockQuantityCode represents a code value taken from the scheme specified in
// the <StockQuantityCodeType> element. Mandatory in each occurrence of the
// <StockQuantityCoded> composite, and non-repeating.
type StockQuantityCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// StockQuantityCoded represents a group of data elements which together
// specify coded stock level without stating the exact quantity of stock.
// Either <StockQuantityCoded> or <OnHand> is mandatory in each occurrence of
// the <Stock> composite, even if the quantity on hand is zero. Repeteable, so
// that it is possible to provide quantities on hand, quantities on order etc
// separately.
type StockQuantityCoded struct {
	StockQuantityCodeType     StockQuantityCodeType
	StockQuantityCodeTypeName *StockQuantityCodeTypeName
	StockQuantityCode         StockQuantityCode
	GeneralAttributes
}

// StockQuantityCodeType represents an ONIX code identifying the scheme from
// which the value in the <StockQuantityCode> element is taken. Mandatory in
// each occurrence of the <StockQuantityCoded> composite, and non-repeating.
type StockQuantityCodeType struct {
	Value string `xml:",innerxml"` // List70
	GeneralAttributes
}

// StockQuantityCodeTypeName represents a name that identifies a proprietary
// stock quantity coding scheme (ie a scheme which is not a standard and for
// which there is no individual ID type code). Must be used when, and only when
// the code in the <StockQuantityCodeType> element indicates a proprietary
// scheme, eg a wholesaler’s own code. Optional, and non-repeating.
type StockQuantityCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// StudyBibleType represents an ONIX code identifying a particular study
// version of a Bible or selected Biblical text, for example ‘Life
// Application’. Optional and non-repeating. Some study Bibles are available in
// different editions based on different text versions.
type StudyBibleType struct {
	Value string `xml:",innerxml"` // List84
	GeneralAttributes
}

// Subject represents an optional and repeatable group of data elements which
// together specify a subject classification or subject heading.
type Subject struct {
	MainSubject             *MainSubject
	SubjectSchemeIdentifier SubjectSchemeIdentifier
	SubjectSchemeName       *SubjectSchemeName
	SubjectSchemeVersion    *SubjectSchemeVersion
	SubjectCode             *SubjectCode
	SubjectHeadingText      []SubjectHeadingText
	GeneralAttributes
}

// SubjectCode represents a subject class or category code from the scheme
// specified in the <SubjectSchemeIdentifier> element. Either <SubjectCode> or
// <SubjectHeadingText> or both must be present in each occurrence of the
// <Subject> composite. Non-repeating.
type SubjectCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// SubjectDate represents an optional and repeatable group of data elements
// which together specify a date associated with the person or organization
// identified in an occurrence of the <NameAsSubject> composite, eg birth or
// death.
type SubjectDate struct {
	SubjectDateRole SubjectDateRole
	DateFormat      *DateFormat
	Date            Date
	GeneralAttributes
}

// SubjectDateRole represents an ONIX code indicating the significance of the
// date in relation to the subject name. Mandatory in each occurrence of the
// <SubjectDate> composite, and non-repeating.
type SubjectDateRole struct {
	Value string `xml:",innerxml"` // List177
	GeneralAttributes
}

// SubjectHeadingText represents the text of a subject heading taken from the
// scheme specified in the <SubjectSchemeIdentifier> element, or of free
// language keywords if the scheme is specified as ‘keywords’; or the text
// equivalent to the <SubjectCode> value, if both code and text are sent.
// Either <SubjectCode> or <SubjectHeadingText> or both must be present in each
// occurrence of the <Subject> composite.
//
// Optional, and repeatable if the text is sent in multiple languages. The
// language attribute is optional for a single instance of
// <SubjectHeadingText>, but must be included in each instance if
// <SubjectHeadingText> is repeated.
type SubjectHeadingText struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SubjectSchemeIdentifier represents an ONIX code which identifies the
// category scheme which is used in an occurrence of the <Subject> composite.
// Mandatory in each occurrence of the composite, and non-repeating. For
// category schemes that use code values, use the associated <SubjectCode>
// element to carry the value (if so required, the <SubjectHeadingText> element
// can be used simultaneously to carry the text equivalent of the code). For
// schemes that use text headings, use the <SubjectHeadingText> element to
// carry the text of the category heading.
type SubjectSchemeIdentifier struct {
	Value string `xml:",innerxml"` // List27
	GeneralAttributes
}

// SubjectSchemeName represents a name identifying a proprietary subject scheme
// (ie a scheme which is not a standard and for which there is no individual
// identifier code) when <SubjectSchemeIdentifier> is coded ‘24’. Optional and
// non-repeating.
type SubjectSchemeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SubjectSchemeVersion represents a number which identifies a version or
// edition of the subject scheme specified in the associated
// <SubjectSchemeIdentifier> element. Optional and non-repeating.
type SubjectSchemeVersion struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// Subtitle represents the text of a subtitle, if any. ‘Subtitle‘ means any
// added words which appear with the title element given in an occurrence of
// the <TitleElement> composite, and which amplify and explain the title
// element, but which are not considered to be part of the title element
// itself. Optional and non-repeating.
type Subtitle struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

// SuffixToKey represents the sixth part of a structured name of a person who
// contributed to the creation of the product: a suffix following a person’s
// key name(s), eg ‘Jr’ or ‘III’. Optional and non-repeating.
type SuffixToKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// Supplier represents A group of data elements which together define a
// supplier. Mandatory in each occurrence of the <SupplyDetail> composite, and
// not repeatable.
type Supplier struct {
	SupplierRole       SupplierRole
	SupplierIdentifier []SupplierIdentifier
	SupplierName       *SupplierName
	TelephoneNumber    []TelephoneNumber
	FaxNumber          []FaxNumber
	EmailAddress       []EmailAddress
	Website            []Website
	GeneralAttributes
}

// SupplierCodeType represents an ONIX code identifying the type of a supplier
// own code. Mandatory in each occurrence of the <SupplierOwnCoding> composite,
// and non-repeating.
type SupplierCodeType struct {
	Value string `xml:",innerxml"` // List165
	GeneralAttributes
}

// SupplierCodeTypeName represents a name which identifies the proprietary
// coding scheme used. Optional and non-repeating.
type SupplierCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SupplierCodeValue represents a supplier-defined code of the type specified
// in the <SupplierCodeType> element. Mandatory in each occurrence of the
// <SupplierOwnCoding> composite, and non-repeating.
type SupplierCodeValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// SupplierIdentifier represents a repeatable group of data elements which
// together define the identity of a supplier in accordance with a specified
// scheme, and allowing different types of supplier identifier to be included
// without defining additional data elements. Optional, but each occurrence of
// the <NewSupplier> composite must carry either at least one supplier
// identifier, or a <SupplierName>, or both.
type SupplierIdentifier struct {
	SupplierIDType SupplierIDType
	IDTypeName     *IDTypeName
	IDValue        IDValue
	GeneralAttributes
}

// SupplierIDType represents an ONIX code identifying the scheme from which
// the identifier in the <IDValue> element is taken. Mandatory in each
// occurrence of the <SupplierIdentifier> composite, and non-repeating.
type SupplierIDType struct {
	Value string `xml:",innerxml"` // List92
	GeneralAttributes
}

// SupplierName represents the name of a new supplier. Optional and non-
// repeating; required if no supplier identifier is sent in an occurrence of
// the <NewSupplier> composite.
type SupplierName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// SupplierOwnCoding represents an optional and repeatable group of data
// elements which together allow a supplier to send coded data of a specified
// type, using its own coding schemes.
type SupplierOwnCoding struct {
	SupplierCodeType     SupplierCodeType
	SupplierCodeTypeName *SupplierCodeTypeName
	SupplierCodeValue    SupplierCodeValue
	GeneralAttributes
}

// SupplierRole represents an ONIX code identifying the role of a supplier in
// relation to the product, eg Publisher, Publisher’s exclusive distributor,
// etc. Mandatory in each occurrence of the <Supplier> composite, and non-
//repeating.
type SupplierRole struct {
	Value string `xml:",innerxml"` // List93
	GeneralAttributes
}

// SupplyDate represents an optional and repeatable group of data elements
// which together specify a date associated with the supply status of the
// product, eg expected ship date.
type SupplyDate struct {
	SupplyDateRole SupplyDateRole
	DateFormat     *DateFormat
	Date           Date
	GeneralAttributes
}

// SupplyDateRole represents an ONIX code indicating the significance of the
// date. Mandatory in each occurrence of the <SupplyDate> composite, and non-
// repeating.
type SupplyDateRole struct {
	Value string `xml:",innerxml"` // List166
	GeneralAttributes
}

// SupplyDetail represents a group of data elements which together give details
// of a supply source, and price and availability from that source. Mandatory
// in each occurrence of the <ProductSupply> block and repeatable.
type SupplyDetail struct {
	Supplier            Supplier
	SupplierOwnCoding   []SupplierOwnCoding
	ReturnsConditions   []ReturnsConditions
	ProductAvailability ProductAvailability
	SupplyDate          []SupplyDate
	OrderTime           *OrderTime
	NewSupplier         *NewSupplier
	Stock               []Stock
	PackQuantity        *PackQuantity
	Reissue             *Reissue
	UnpricedItemType    *UnpricedItemType
	Price               []Price
	GeneralAttributes
}

// SupportingResource [DEPRECATED] represents an optional and repeatable group
// of data elements which together specify a supporting resource, used here to
// indicate that there is a new cover or jacket image, or other supporting
// resource, for a forthcoming reissue. Deprecated in this context.
type SupportingResource struct {
	ResourceContentType ResourceContentType
	ContentAudience     []ContentAudience
	Territory           *Territory
	ResourceMode        ResourceMode
	ResourceFeature     []ResourceFeature
	ResourceVersion     []ResourceVersion
	GeneralAttributes
}

// Tax represents a repeatable group of data elements which together specify tax
// applicable to a price amount. Optional, and used only when <PriceType>
// indicates an inc-tax price. For items to which different taxes or tax rates
// apply (eg mixed media products in the UK which are partly taxed at standard
// rate and partly at zero rate), the composite is repeated for each separate
// tax or tax rate. Although only one of <TaxRatePercent> or <TaxAmount> is
// mandatory within the composite, it is recommended that all tax elements in
// the composite should be explicitly populated.
//
// If the tax regime requires separate tax rates and amounts linked explicitly
// to particular product parts (eg in Germany), the <ProductIdentifier>
// composite may be included in each <Tax> composite. Where tax is payable on
// multiple product parts, each should have its own instance of the <Tax>
// composite.
type Tax struct {
	ProductIdentifier []ProductIdentifier
	TaxType           *TaxType
	TaxRateCode       *TaxRateCode
	TaxRatePercent    TaxRatePercent
	TaxableAmount     *TaxableAmount
	TaxAmount         *TaxAmount
	GeneralAttributes
}

// TaxAmount represents the amount of tax chargeable at the rate specified in
// an occurrence of the <Tax> composite. Optional and non-repeating; but either
// <TaxRatePercent> or <TaxAmount> or both must be present in each occurrence
// of the <Tax> composite.
type TaxAmount struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

// TaxRateCode represents an ONIX code which specifies a tax rate. Optional and
// non-repeating.
type TaxRateCode struct {
	Value string `xml:",innerxml"` // List62
	GeneralAttributes
}

// TaxRatePercent represents a tax rate expressed numerically as a percentage.
// Optional and non-repeating; but either <TaxRatePercent> or <TaxAmount> or
// both must be present in each occurrence of the <Tax> composite.
type TaxRatePercent struct {
	Value string `xml:",innerxml"` // dt.PercentDecimal
	GeneralAttributes
}

// TaxableAmount represents the amount of the unit price of the product,
// excluding tax, which is taxable at the rate specified in an occurrence of
// the <Tax> composite. Optional and non-repeating; but required if tax is
// charged on part of the price. Omission of this element implies that tax is
// charged on the full amount of the price.
type TaxableAmount struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

// TaxType represents an ONIX code identifying a tax type, eg VAT. Optional,
// and non-repeating.
type TaxType struct {
	Value string `xml:",innerxml"` // List171
	GeneralAttributes
}

// TelephoneNumber [DEPRECATED] represents a telephone number of an agent or
// local // publisher. Optional and repeatable. Deprecated in this context, in
// favor of providing contact details in the <ProductContact> composite.
type TelephoneNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

// Territory represents a group of data elements which together identify a
// territory in which the price stated in an occurrence of the <Price>
// composite is applicable. Optional and non-repeating. Additional guidance
// on the description of price territories in ONIX 3.0 will be found in a
// separate document ONIX for Books Product Information Message: How to
// Specify Markets and Suppliers in ONIX 3.
type Territory struct {
	RegionsExcluded   *RegionsExcluded
	CountriesIncluded *CountriesIncluded
	RegionsIncluded   *RegionsIncluded
	CountriesExcluded *CountriesExcluded
	GeneralAttributes
}

// Text represents the text specified in the <TextType> element. Mandatory in
// each occurrence of the <TextContent> composite, and repeatable when
// essentially identical text is supplied in multiple languages. The language
// attribute is optional for a single instance of <Text>, but must be included
// in each instance if <Text> is repeated.
type Text struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// TextAuthor represents the name of an author of text sent in the <Text>
// element, eg if it is a review or promotional quote. Optional, and repeatable
// if the text is jointly authored.
type TextAuthor struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// TextContent represents a group of data elements which together carry text
// related to a content item. Optional and repeatable.
type TextContent struct {
	TextType            TextType
	ContentAudience     []ContentAudience
	Territory           *Territory
	Text                []Text
	ReviewRating        *ReviewRating
	TextAuthor          []TextAuthor
	TextSourceCorporate *TextSourceCorporate
	SourceTitle         []SourceTitle
	ContentDate         []ContentDate
	GeneralAttributes
}

// TextItem represents a group of data elements which are specific to text
// content. The composite must occur once and only once in a <ContentItem>
// composite which describes a text content item. (Similar composites may be
// defined for other media, and the occurrence of one of them will be mandatory
// in any <ContentItem> composite.)
type TextItem struct {
	TextItemType       TextItemType
	TextItemIdentifier []TextItemIdentifier
	PageRun            []PageRun
	NumberOfPages      *NumberOfPages
	GeneralAttributes
}

// TextItemIDType represents an ONIX code identifying the scheme from which the
// identifier in <IDValue> is taken. Mandatory in each occurrence of the
// <TextItemIdentifier> composite, and non-repeating.
type TextItemIDType struct {
	Value string `xml:",innerxml"` // List43
	GeneralAttributes
}

// TextItemIdentifier represents a repeatable group of data elements which
// together define an identifier of a text item in accordance with a specified
// scheme. The composite is optional.
type TextItemIdentifier struct {
	TextItemIDType TextItemIDType
	IDTypeName     *IDTypeName
	IDValue        IDValue
	GeneralAttributes
}

// TextItemType represents an ONIX code which identifies the nature of a text
// item. Mandatory in each occurrence of the <TextItem> composite, and non-
// repeatable.
type TextItemType struct {
	Value string `xml:",innerxml"` // List42
	GeneralAttributes
}

// TextSourceCorporate represents the name of a company or corporate body
// responsible for the text sent in the <Text> element. Optional and non-
// repeating.
type TextSourceCorporate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// TextType represents an ONIX code which identifies the type of text which is
// sent in the <Text> element. Mandatory in each occurrence of the
// <TextContent> composite, and non-repeating.
type TextType struct {
	Value string `xml:",innerxml"` // List153
	GeneralAttributes
}

// ThesisPresentedTo represents the name of an academic institution to which a
// thesis was presented. Optional and non-repeating, but if this element is
// present, <ThesisType> must also be present.
type ThesisPresentedTo struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

// ThesisType represents an ONIX code identifying a thesis type, when the ONIX
// record describes an item which was originally presented as an academic
// thesis or dissertation. Optional and non-repeating.
type ThesisType struct {
	Value string `xml:",innerxml"` // List72
	GeneralAttributes
}

// ThesisYear represents the year in which a thesis was presented. Optional and
// non-repeating, but if this element is present, <ThesisType> must also be
// present.
type ThesisYear struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

// TitleDetail represents a repeatable group of data elements which together
// give the text of a title of a content item and specify its type. Optional.
type TitleDetail struct {
	TitleType      TitleType
	TitleElement   []TitleElement
	TitleStatement *TitleStatement
	GeneralAttributes
}

// TitleElement represents a repeatable group of data elements which together
// represent an element of a collection title. At least one title element is
// mandatory in each occurrence of the <TitleDetail> composite. An instance of
// the <TitleElement> composite must include at least one of: <PartNumber>;
// <YearOfAnnual>; <TitleText>, <NoPrefix/> together with <TitleWithoutPrefix>,
// or <TitlePrefix> together with <TitleWithoutPrefix>. In other words it must
// carry either the text of a title element or a part or year designation; and
// it may carry both. A title element must be designated as belonging to
// product level, collection level, or subcollection level (the first of these
// may not occur in a title element representing a collective identity, and the
// last-named may only occur in the case of a multi-level collection).
//
// In the simplest case, title detail sent in a <Collection> composite will
// consist of a single title element, at collection level. However, the
// composite structure in ONIX 3.0 allows more complex combinations of titles
// and part designations in multi-level collections to be correctly represented.
type TitleElement struct {
	SequenceNumber     *SequenceNumber
	TitleElementLevel  TitleElementLevel
	PartNumber         *PartNumber
	YearOfAnnual       *YearOfAnnual
	TitlePrefix        *TitlePrefix
	NoPrefix           *NoPrefix
	TitleWithoutPrefix *TitleWithoutPrefix
	TitleText          *TitleText
	Subtitle           *Subtitle
	GeneralAttributes
}

// TitleElementLevel represents an ONIX code indicating the level of a title
// element: collection level, subcollection level, or product level. Mandatory
// in each occurrence of the <TitleElement> composite, and non-repeating.
type TitleElementLevel struct {
	Value string `xml:",innerxml"` // List149
	GeneralAttributes
}

// TitlePrefix represents text at the beginning of a title element which is to
// be ignored for alphabetical sorting. Optional and non-repeating; can only be
// used when <TitleText> is omitted, and if the <TitleWithoutPrefix> element is
// also present. These two elements may be used in combination in applications
// where it is necessary to distinguish an initial word or character string
// which is to be ignored for filing purposes, eg in library systems and in
// some bookshop databases.
type TitlePrefix struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

// TitlesAfterNames represents titles following a person’s names, eg ‘Duke of
// Edinburgh’. Optional and non-repeating.
type TitlesAfterNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// TitlesBeforeNames represents qualifications and/or titles preceding a
// person’s names, eg ‘Professor’ or ‘HRH Prince’ or ‘Saint’. Optional and
// non-repeating: see Group P.7 introductory text for valid options.
type TitlesBeforeNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

// TitleStatement represents af ree text showing how the collection title
// should be presented in any display, particularly when a standard
// concatenation of individual title elements from Group P.5 (in the order
// specified by the <SequenceNumber> data elements) would not give a
// satisfactory result. Optional and non-repeating. When this field is sent,
// the recipient should use it to replace all title detail sent in Group P.5
// for display purposes only. The individual collection title element detail
// must also be sent, for indexing and retrieval.
type TitleStatement struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// TitleText represents the text of a title element, excluding any subtitle.
// Optional and non-repeating, may only be used where <TitlePrefix>,
// <NoPrefix/> and <TitleWithoutPrefix> are not used. This element is intended
// to be used when the sending system cannot reliably provide prefixes that are
// ignored for sorting purposes in a separate data element. If the system can
// reliably separate prefixes, it should state whether a prefix is present
// (using <TitlePrefix> and <TitleWithoutPrefix>) or absent (using <NoPrefix/>
// and <TitleWithoutPrefix>).
type TitleText struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

// TitleType represents an ONIX code indicating the type of a title. Mandatory
// in each occurrence of the <TitleDetail> composite, and non-repeating.
type TitleType struct {
	Value string `xml:",innerxml"` // List15
	GeneralAttributes
}

// TitleWithoutPrefix represents the text of a title element without the title
// prefix; and excluding any subtitle. Optional and non-repeating; can only be
// used if one of the <NoPrefix/> or <TitlePrefix> elements is also present.
type TitleWithoutPrefix struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

// ToLanguage is used only when the <ContributorRole> code value is B06, B08 or
// B10 indicating a translator, to specify the target language into which the
// translation was made. This element makes it possible to specify a
// translator’s exact responsibility when a work involves translation into two
// or more languages. Optional, and repeatable in the event that a single
// person has been responsible for translation to two or more languages.
type ToLanguage struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

// ToQuantity represents a maximum order quantity eligible for a specified
// discount, used only in the case of ‘progressive’ discounts. Optional, but
// where used, must be preceded by a minimum qualifying order quantity (even
// if that minimum is 1). For the special case where there is no maximum (ie
// in the repeat of the <Discount> composite that specifies the highest
// progressive discount), use zero.
type ToQuantity struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

// TradeCategory represents an ONIX code which indicates a trade category
// which is somewhat related to, but not properly an attribute of, product
// form. Optional and non-repeating.
type TradeCategory struct {
	Value string `xml:",innerxml"` // List12
	GeneralAttributes
}

// UnnamedPersons represents an ONIX code allowing a positive indication to
// be given when authorship is unknown or anonymous, or when as a matter of
// editorial policy only a limited number of contributors are named. Optional
// and non-repeating: see Group P.7 introductory text for valid options. Use
// here in preference to P.7.47, where it is deprecated.
type UnnamedPersons struct {
	Value string `xml:",innerxml"` // List19
	GeneralAttributes
}

// UnpricedItemType represents an ONIX code which specifies the product is
// free of charge, or a reason why a price is not sent. If code value 02 is
// used to send advance information without giving a price, the price must be
// confirmed as soon as possible. Optional and non-repeating, but required if
// the <SupplyDetail> composite does not carry a price. Use here in preference
// to P.26.70a when the product is available only under free of charge or
// unpriced terms from the supplier.
type UnpricedItemType struct {
	Value string `xml:",innerxml"` // List57
	GeneralAttributes
}

// Velocity represents an optional group of data elements which together
// specify the rate of stock depletion – or equally, a rate of accumulation
// of backorders. Repeatable if the rate of depletion is specified using more
// than one metric (eg specifying both a minimum and maximum daily sale).
type Velocity struct {
	VelocityMetric VelocityMetric
	Rate           Rate
	Proximity      *Proximity
	GeneralAttributes
}

// VelocityMetric represents an ONIX code that specifies how the rate of stock
// depletion is measured. Mandatory within the <Velocity> composite, and non-
// repeating.
type VelocityMetric struct {
	Value string `xml:",innerxml"` // List216
	GeneralAttributes
}

// Website represents an optional and repeatable group of data elements which
// together identify and provide pointers to a website which is related to the
// person or organization identified in an occurrence of the <Contributor>
// composite.
type Website struct {
	WebsiteRole        *WebsiteRole
	WebsiteDescription []WebsiteDescription
	WebsiteLink        WebsiteLink
	GeneralAttributes
}

// WebsiteDescription represents a free text describing the nature of the
// website which is linked through the <WebsiteLink> element. Optional, and
// repeatable to provide parallel descriptive text in multiple languages. The
// language attribute is optional for a single instance of
// <WebsiteDescription>, but must be included in each instance if
// <WebsiteDescription> is repeated.
type WebsiteDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

// WebsiteLink represents the URL for the website. Mandatory in each occurrence
// of the <Website> composite, and non-repeating.
type WebsiteLink struct {
	Value string `xml:",innerxml"` // dt.NonEmptyURI
	GeneralAttributes
}

// WebsiteRole represents an ONIX code which identifies the role or purpose of
// the website which is linked through the <WebsiteLink> element. Optional and
// non-repeating.
type WebsiteRole struct {
	Value string `xml:",innerxml"` // List73
	GeneralAttributes
}

// WorkIdentifier represents a repeatable group of data elements which together
// define an identifier of a work in accordance with a specified scheme.
// Mandatory in each occurrence of the <RelatedWork> composite. Repeatable only
// if two or more identifiers for the same work are sent using different
// identifier schemes (eg ISTC and DOI).
type WorkIdentifier struct {
	WorkIDType WorkIDType
	IDTypeName *IDTypeName
	IDValue    IDValue
	GeneralAttributes
}

// WorkIDType represents an ONIX code identifying the scheme from which the
// identifier in the <IDValue> element is taken. Mandatory in each occurrence
// of the <WorkIdentifier> composite, and non-repeating.
type WorkIDType struct {
	Value string `xml:",innerxml"` // List16
	GeneralAttributes
}

// WorkRelationCode represents an ONIX code which identifies the nature of the
// relationship between a product and a work. Mandatory in each occurrence of
// the <RelatedWork> composite, and non-repeating.
type WorkRelationCode struct {
	Value string `xml:",innerxml"` // List164
	GeneralAttributes
}

// YearOfAnnual is used when the year of an annual is part of a title. This
// field should be used to carry the year (or, if required, a spread of years
// such as 2009-2010). Optional and non-repeating.
type YearOfAnnual struct {
	Value string `xml:",innerxml"` // dt.YearOrYearRange
	GeneralAttributes
}

// From ONIX for Books Release 3.0, identifies the release of the ONIX format
// standard to which the message conforms. Used only in the top-level element
// <ONIXMessage> (short tag <ONIXmessage>), and is mandatory. The value will
// change with each new release, so that all messages will show explicitly the
// release to which they are intended to conform.
type ReleaseAttribute struct {
	Release string `xml:"release,attr,omitempty"`
}

// CollationkeyAttribute enables a data element to carry the key to be used for
// sorting, when the sort order is not inherent to the data itself. For example,
// with Chinese or Japanese contributor names, the collationkey attribute may
// carry phonetic information required to sort records by contributor.
type CollationkeyAttribute struct {
	Collationkey string `xml:"collationkey,attr,omitempty"` // xs:string
}

// GeneralAttributes represents three attributes that may be used with any ONIX
// element, and they are not noted individually for each data element in the
// specification:
type GeneralAttributes struct {
	Datestamp  string `xml:"datestamp,attr,omitempty"`  // dt.DateOrDateTime
	Sourcetype string `xml:"sourcetype,attr,omitempty"` // SourceTypeCode
	Sourcename string `xml:"sourcename,attr,omitempty"`
}

// DateformatAttribute is used with a range of date elements to specify the
// format of the date. Each data element on which this attribute may be used
// specifies a default dateformat if the attribute is not supplied – for most
// date elements, this is format ‘00’, YYYYMMDD. In some cases, the format of
// the date may be described via a <DateFormat> data element instead, but this
// is deprecated. If dateformat and <DateFormat> are both supplied,
// <DateFormat> should be ignored.
type DateformatAttribute struct {
	Dateformat string `xml:"dateformat,attr,omitempty"` // List55
}

// LanguageAttribute enables the language of any text element to be specified
// when it is not the expected default language of the message. The default
// language of the message (ie of the metadata) is generally set by agreement
// between sender and recipient, and is separate from (though usually identical
// to) the default language of the text used within the products described
// within the message (for the latter, see <DefaultLanguageOfText>). Many data
// elements that carry the language attribute are repeatable in order to allow
// parallel text to be provided in multiple languages. Data recipients able to
// support only a single language should select the repeat that carries the
// most appropriate language attribute.
//
// Generally, a language also implies a particular script. (There are a very
// few languages that are commonly written in more than one script.) However,
// names are not considered to be ‘in’ a particular language, but are commonly
// transliterated from one script to another, and a limited number of data
// elements carry the textscript attribute. These may also carry a language
// attribute in order that distinct transliterations that match the orthography
// norms of a particular language can be distinguished, cf the Cyrillic name
// ‘Александр Солженицын’, and the transliterations into Latin script
// ‘Aleksandr Solzhenitsyn’ (English), and ‘Alexandre Soljenitsyne’ (French).
type LanguageAttribute struct {
	Language string `xml:"language,attr,omitempty"` // List74
}

// TextcaseAttribute enables the case of the text of a title or subtitle to be
// specified. If not supplied, the default value is ‘00’, indicating the case
// is Undefined.
type TextcaseAttribute struct {
	Textcase string `xml:"textcase,attr,omitempty"` // TextCaseCode
}

// TextformatAttribute is used with a limited range of text elements that are
// allowed to contain formatted text, to enable the markup format such as XHTML
// to be specified. If not supplied, the default is ‘06’, indicating the text
// format is plain text in the character encoding declared in the XML
// declaration at the top of message, without additional markup. If the XML
// declaration does not specify a character encoding, the XML default character
// set should be assumed to be the basic ASCII characters. See the Character
// sets and special characters section below for further details of how a
// character encoding declaration is used.
type TextformatAttribute struct {
	Textformat string `xml:"textformat,attr,omitempty"` // TextFormatCode
}

// TextscriptAttribute is used with a limited range of text elements that are
// used to provide transliterated alternatives to names within the
// <AlternativeName> composite, and to titles.
type TextscriptAttribute struct {
	Textscript string `xml:"textscript,attr,omitempty"` // List121
}
