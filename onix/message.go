package onix

type Addressee struct {
	ContactName         *ContactName
	EmailAddress        *EmailAddress
	AddresseeIdentifier []AddresseeIdentifier
	AddresseeName       *AddresseeName
	GeneralAttributes
}

type AddresseeIdentifier struct {
	AddresseeIDType AddresseeIDType
	IDTypeName      *IDTypeName
	IDValue         IDValue
	GeneralAttributes
}

type AddresseeIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type AddresseeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type Affiliation struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type AgentIdentifier struct {
	AgentIDType AgentIDType
	IDTypeName  *IDTypeName
	IDValue     IDValue
	GeneralAttributes
}

type AgentIDType struct {
	Value string `xml:",innerxml"` // List92
	GeneralAttributes
}

type AgentName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type AgentRole struct {
	Value string `xml:",innerxml"` // List69
	GeneralAttributes
}

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

type AncillaryContent struct {
	AncillaryContentType        AncillaryContentType
	AncillaryContentDescription []AncillaryContentDescription
	Number                      *Number
	GeneralAttributes
}

type AncillaryContentType struct {
	Value string `xml:",innerxml"` // List25
	GeneralAttributes
}

type AncillaryContentDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type Audience struct {
	AudienceCodeType     AudienceCodeType
	AudienceCodeTypeName *AudienceCodeTypeName
	AudienceCodeValue    AudienceCodeValue
	GeneralAttributes
}

type AudienceCode struct {
	Value string `xml:",innerxml"` // List28
	GeneralAttributes
}

type AudienceCodeType struct {
	Value string `xml:",innerxml"` // List29
	GeneralAttributes
}

type AudienceCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type AudienceCodeValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type AudienceDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type AudienceRange struct {
	AudienceRangeQualifier AudienceRangeQualifier
	AudienceRangePrecision AudienceRangePrecision
	AudienceRangeValue     AudienceRangeValue
	GeneralAttributes
}

type AudienceRangePrecision struct {
	Value string `xml:",innerxml"` // List31
	GeneralAttributes
}

type AudienceRangeQualifier struct {
	Value string `xml:",innerxml"` // List30
	GeneralAttributes
}

type AudienceRangeValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type Barcode struct {
	BarcodeType       BarcodeType
	PositionOnProduct *PositionOnProduct
	GeneralAttributes
}

type BarcodeType struct {
	Value string `xml:",innerxml"` // List141
	GeneralAttributes
}

type BatchBonus struct {
	BatchQuantity BatchQuantity
	FreeQuantity  FreeQuantity
	GeneralAttributes
}

type BatchQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

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

type BibleContents struct {
	Value string `xml:",innerxml"` // List82
	GeneralAttributes
}

type BiblePurpose struct {
	Value string `xml:",innerxml"` // List85
	GeneralAttributes
}

type BibleReferenceLocation struct {
	Value string `xml:",innerxml"` // List87
	GeneralAttributes
}

type BibleTextFeature struct {
	Value string `xml:",innerxml"` // List97
	GeneralAttributes
}

type BibleTextOrganization struct {
	Value string `xml:",innerxml"` // List86
	GeneralAttributes
}

type BibleVersion struct {
	Value string `xml:",innerxml"` // List83
	GeneralAttributes
}

type BiographicalNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type BookClubAdoption struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type CBO struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

type CitationNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type CitedContent struct {
	CitedContentType CitedContentType
	ContentAudience  []ContentAudience
	Territory        *Territory
	SourceType       *SourceType
	CitationNote     []CitationNote
	ResourceLink     []ResourceLink
	ContentDate      []ContentDate
	ReviewRating     ReviewRating
	SourceTitle      []SourceTitle
	ListName         []ListName
	PositionOnList   *PositionOnList
	GeneralAttributes
}

type CitedContentType struct {
	Value string `xml:",innerxml"` // List156
	GeneralAttributes
}

type CityOfPublication struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type CollateralDetail struct {
	TextContent        []TextContent
	CitedContent       []CitedContent
	SupportingResource []SupportingResource
	Prize              []Prize
	GeneralAttributes
}

type Collection struct {
	CollectionType       CollectionType
	SourceName           *SourceName
	CollectionIdentifier []CollectionIdentifier
	CollectionSequence   []CollectionSequence
	TitleDetail          []TitleDetail
	GeneralAttributes
}

type CollectionIdentifier struct {
	CollectionIDType CollectionIDType
	IDTypeName       *IDTypeName
	IDValue          IDValue
	GeneralAttributes
}

type CollectionIDType struct {
	Value string `xml:",innerxml"` // List13
	GeneralAttributes
}

type CollectionSequence struct {
	CollectionSequenceType     CollectionSequenceType
	CollectionSequenceTypeName *CollectionSequenceTypeName
	CollectionSequenceNumber   CollectionSequenceNumber
	GeneralAttributes
}

type CollectionSequenceNumber struct {
	Value string `xml:",innerxml"` // dt.MultiLevelNumber
	GeneralAttributes
}

type CollectionSequenceType struct {
	Value string `xml:",innerxml"` // List197
	GeneralAttributes
}

type CollectionSequenceTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type CollectionType struct {
	Value string `xml:",innerxml"` // List148
	GeneralAttributes
}

type ComparisonProductPrice struct {
	ProductIdentifier []ProductIdentifier
	PriceType         *PriceType
	PriceAmount       PriceAmount
	CurrencyCode      *CurrencyCode
	GeneralAttributes
}

type Complexity struct {
	ComplexitySchemeIdentifier ComplexitySchemeIdentifier
	ComplexityCode             ComplexityCode
	GeneralAttributes
}

type ComplexityCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ComplexitySchemeIdentifier struct {
	Value string `xml:",innerxml"` // List32
	GeneralAttributes
}

type ComponentNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ComponentTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

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

type ConferenceAcronym struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ConferenceDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type ConferenceName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ConferenceNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type ConferencePlace struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ConferenceRole struct {
	Value string `xml:",innerxml"` // List20
	GeneralAttributes
}

type ConferenceSponsor struct {
	GeneralAttributes
}

type ConferenceSponsorIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type ConferenceSponsorIdentifier struct {
	ConferenceSponsorIDType ConferenceSponsorIDType
	IDTypeName              *IDTypeName
	IDValue                 IDValue
	GeneralAttributes
}

type ConferenceTheme struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type ContactName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ContentAudience struct {
	Value string `xml:",innerxml"` // List154
	GeneralAttributes
}

type ContentDate struct {
	ContentDateRole ContentDateRole
	DateFormat      *DateFormat
	Date            Date
	GeneralAttributes
}

type ContentDateRole struct {
	Value string `xml:",innerxml"` // List155
	GeneralAttributes
}

type ContentDetail struct {
	ContentItem []ContentItem
	GeneralAttributes
}

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
	BiographicalNote        []BiographicalNote
	Website                 []Website
	ContributorDescription  []ContributorDescription
	ContributorPlace        []ContributorPlace
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
	UnnamedPersons          UnnamedPersons
	GeneralAttributes
}

type ContributorDate struct {
	ContributorDateRole ContributorDateRole
	DateFormat          *DateFormat
	Date                Date
	GeneralAttributes
}

type ContributorDateRole struct {
	Value string `xml:",innerxml"` // List177
	GeneralAttributes
}

type ContributorDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type ContributorPlace struct {
	ContributorPlaceRelator ContributorPlaceRelator
	LocationName            []LocationName
	CountryCode             CountryCode
	RegionCode              *RegionCode
	GeneralAttributes
}

type ContributorPlaceRelator struct {
	Value string `xml:",innerxml"` // List151
	GeneralAttributes
}

type ContributorRole struct {
	Value string `xml:",innerxml"` // List17
	GeneralAttributes
}

type ContributorStatement struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type CopiesSold struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type CopyrightOwner struct {
	GeneralAttributes
}

type CopyrightOwnerIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type CopyrightOwnerIdentifier struct {
	CopyrightOwnerIDType CopyrightOwnerIDType
	IDTypeName           *IDTypeName
	IDValue              IDValue
	GeneralAttributes
}

type CopyrightStatement struct {
	CopyrightType  *CopyrightType
	CopyrightYear  []CopyrightYear
	CopyrightOwner []CopyrightOwner
	GeneralAttributes
}

type CopyrightType struct {
	Value string `xml:",innerxml"` // List219
	GeneralAttributes
}

type CopyrightYear struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type CorporateName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type CorporateNameInverted struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type CountriesIncluded struct {
	Value string `xml:",innerxml"` // CountryCodeList
	GeneralAttributes
}

type CountriesExcluded struct {
	Value string `xml:",innerxml"` // CountryCodeList
	GeneralAttributes
}

type CountryCode struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

type CountryOfManufacture struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

type CountryOfPublication struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

type CurrencyCode struct {
	Value string `xml:",innerxml"` // List96
	GeneralAttributes
}

type CurrencyZone struct {
	Value string `xml:",innerxml"` // List172
	GeneralAttributes
}

type Date struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type DateFormat struct {
	Value string `xml:",innerxml"` // List55
	GeneralAttributes
}

type DefaultCurrencyCode struct {
	Value string `xml:",innerxml"` // List96
	GeneralAttributes
}

type DefaultLanguageOfText struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

type DefaultPriceType struct {
	Value string `xml:",innerxml"` // List58
	GeneralAttributes
}

type DeletionText struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

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
	NoCollection            NoCollection
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

type Discount struct {
	DiscountType    *DiscountType
	DiscountPercent DiscountPercent
	DiscountAmount  *DiscountAmount
	GeneralAttributes
}

type DiscountAmount struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

type DiscountCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type DiscountCodeType struct {
	Value string `xml:",innerxml"` // List100
	GeneralAttributes
}

type DiscountCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type DiscountCoded struct {
	DiscountCodeType     DiscountCodeType
	DiscountCodeTypeName *DiscountCodeTypeName
	DiscountCode         DiscountCode
	GeneralAttributes
}

type DiscountType struct {
	Value string `xml:",innerxml"` // List170
	GeneralAttributes
}

type EditionNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type EditionStatement struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type EditionType struct {
	Value string `xml:",innerxml"` // List21
	GeneralAttributes
}

type EditionVersionNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type EmailAddress struct {
	Value string `xml:",innerxml"` // dt.EmailString
	GeneralAttributes
}

type EndDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type EpubLicense struct {
	EpubLicenseName       []EpubLicenseName
	EpubLicenseExpression []EpubLicenseExpression
	GeneralAttributes
}

type EpubLicenseExpression struct {
	EpubLicenseExpressionType     EpubLicenseExpressionType
	EpubLicenseExpressionTypeName *EpubLicenseExpressionTypeName
	EpubLicenseExpressionLink     EpubLicenseExpressionLink
	GeneralAttributes
}

type EpubLicenseExpressionLink struct {
	Value string `xml:",innerxml"` // dt.NonEmptyURI
	GeneralAttributes
}

type EpubLicenseExpressionType struct {
	Value string `xml:",innerxml"` // List218
	GeneralAttributes
}

type EpubLicenseExpressionTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type EpubLicenseName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type EpubTechnicalProtection struct {
	Value string `xml:",innerxml"` // List144
	GeneralAttributes
}

type EpubUsageConstraint struct {
	EpubUsageType   EpubUsageType
	EpubUsageStatus EpubUsageStatus
	EpubUsageLimit  []EpubUsageLimit
	GeneralAttributes
}

type EpubUsageLimit struct {
	Quantity      Quantity
	EpubUsageUnit EpubUsageUnit
	GeneralAttributes
}

type EpubUsageStatus struct {
	Value string `xml:",innerxml"` // List146
	GeneralAttributes
}

type EpubUsageType struct {
	Value string `xml:",innerxml"` // List145
	GeneralAttributes
}

type EpubUsageUnit struct {
	Value string `xml:",innerxml"` // List147
	GeneralAttributes
}

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

type EventAcronym struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type EventDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type EventName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type EventNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type EventPlace struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type EventRole struct {
	Value string `xml:",innerxml"` // List20
	GeneralAttributes
}

type EventSponsor struct {
	GeneralAttributes
}

type EventSponsorIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type EventSponsorIdentifier struct {
	EventSponsorIDType EventSponsorIDType
	IDTypeName         *IDTypeName
	IDValue            IDValue
	GeneralAttributes
}

type EventTheme struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ExpectedDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type Extent struct {
	ExtentType       ExtentType
	ExtentUnit       ExtentUnit
	ExtentValue      ExtentValue
	ExtentValueRoman *ExtentValueRoman
	GeneralAttributes
}

type ExtentType struct {
	Value string `xml:",innerxml"` // List23
	GeneralAttributes
}

type ExtentUnit struct {
	Value string `xml:",innerxml"` // List24
	GeneralAttributes
}

type ExtentValue struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

type ExtentValueRoman struct {
	Value string `xml:",innerxml"` // dt.RomanNumeralString
	GeneralAttributes
}

type FaxNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type FeatureNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type FeatureValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type FirstPageNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type FreeQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type FromLanguage struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

type Funding struct {
	FundingIdentifier []FundingIdentifier
	GeneralAttributes
}

type FundingIdentifier struct {
	FundingIDType FundingIDType
	IDTypeName    *IDTypeName
	IDValue       IDValue
	GeneralAttributes
}

type FundingIDType struct {
	Value string `xml:",innerxml"` // List228
	GeneralAttributes
}

type Gender struct {
	Value string `xml:",innerxml"` // List229
	GeneralAttributes
}

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

type IDTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type IDValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type Illustrated struct {
	Value string `xml:",innerxml"` // List152
	GeneralAttributes
}

type IllustrationsNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type Imprint struct {
	GeneralAttributes
}

type ImprintIdentifier struct {
	ImprintIDType ImprintIDType
	IDTypeName    *IDTypeName
	IDValue       IDValue
	GeneralAttributes
}

type ImprintIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type ImprintName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type InitialPrintRun struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type KeyNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type Language struct {
	LanguageRole LanguageRole
	LanguageCode LanguageCode
	CountryCode  *CountryCode
	ScriptCode   *ScriptCode
	GeneralAttributes
}

type LanguageCode struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

type LanguageRole struct {
	Value string `xml:",innerxml"` // List22
	GeneralAttributes
}

type LastPageNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type LatestReprintNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type LettersAfterNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type LevelSequenceNumber struct {
	Value string `xml:",innerxml"` // dt.MultiLevelNumber
	GeneralAttributes
}

type ListName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type LocationIdentifier struct {
	LocationIDType LocationIDType
	IDTypeName     *IDTypeName
	IDValue        IDValue
	GeneralAttributes
}

type LocationIDType struct {
	Value string `xml:",innerxml"` // List92
	GeneralAttributes
}

type LocationName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type MainSubject struct {
	GeneralAttributes
}

type MapScale struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type Market struct {
	Territory        Territory
	SalesRestriction []SalesRestriction
	GeneralAttributes
}

type MarketDate struct {
	MarketDateRole MarketDateRole
	DateFormat     *DateFormat
	Date           Date
	GeneralAttributes
}

type MarketDateRole struct {
	Value string `xml:",innerxml"` // List163
	GeneralAttributes
}

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

type MarketPublishingStatus struct {
	Value string `xml:",innerxml"` // List68
	GeneralAttributes
}

type MarketPublishingStatusNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type Measure struct {
	MeasureType     MeasureType
	Measurement     Measurement
	MeasureUnitCode MeasureUnitCode
	GeneralAttributes
}

type Measurement struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

type MeasureType struct {
	Value string `xml:",innerxml"` // List48
	GeneralAttributes
}

type MeasureUnitCode struct {
	Value string `xml:",innerxml"` // List50
	GeneralAttributes
}

type MessageNote struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type MessageNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type MessageRepeat struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type MinimumOrderQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

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

type NameIdentifier struct {
	NameIDType NameIDType
	IDTypeName *IDTypeName
	IDValue    IDValue
	GeneralAttributes
}

type NameIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type NamesAfterKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type NamesBeforeKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type NameType struct {
	Value string `xml:",innerxml"` // List18
	GeneralAttributes
}

type NewSupplier struct {
	TelephoneNumber    []TelephoneNumber
	FaxNumber          []FaxNumber
	EmailAddress       []EmailAddress
	SupplierIdentifier []SupplierIdentifier
	SupplierName       *SupplierName
	GeneralAttributes
}

type NoCollection struct {
	GeneralAttributes
}

type NoContributor struct {
	GeneralAttributes
}

type NoEdition struct {
	GeneralAttributes
}

type NoPrefix struct {
	GeneralAttributes
}

type NoProduct struct {
	GeneralAttributes
}

type NotificationType struct {
	Value string `xml:",innerxml"` // List1
	GeneralAttributes
}

type Number struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

type NumberOfCopies struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type NumberOfIllustrations struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

type NumberOfItemsOfThisForm struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type NumberOfPages struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type ONIXMessage struct {
	Header    Header
	NoProduct NoProduct
	Product   []Product
	GeneralAttributes
	ReleaseAttribute
}

type OnHand struct {
	Value string `xml:",innerxml"` // dt.Integer
	GeneralAttributes
}

type OnOrder struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

type OnOrderDetail struct {
	OnOrder      OnOrder
	Proximity    *Proximity
	ExpectedDate ExpectedDate
	GeneralAttributes
}

type OrderQuantityMinimum struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type OrderQuantityMultiple struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type OrderTime struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

type PackQuantity struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type PageRun struct {
	FirstPageNumber FirstPageNumber
	LastPageNumber  *LastPageNumber
	GeneralAttributes
}

type PartNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
	TextscriptAttribute
}

type Percent struct {
	Value string `xml:",innerxml"` // dt.PercentDecimal
	GeneralAttributes
}

type DiscountPercent struct {
	Value string `xml:",innerxml"` // dt.PercentDecimal
	GeneralAttributes
}

type PersonName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type PersonNameInverted struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type PositionOnList struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type PositionOnProduct struct {
	Value string `xml:",innerxml"` // List142
	GeneralAttributes
}

type PrefixToKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

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
	CurrencyCode           *CurrencyCode
	Territory              *Territory
	CurrencyZone           *CurrencyZone
	ComparisonProductPrice []ComparisonProductPrice
	PriceDate              []PriceDate
	PrintedOnProduct       *PrintedOnProduct
	PositionOnProduct      *PositionOnProduct
	Tax                    []Tax
	PriceAmount            PriceAmount
	PriceCoded             PriceCoded
	UnpricedItemType       UnpricedItemType
	GeneralAttributes
}

type PriceAmount struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

type PriceCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type PriceCoded struct {
	PriceCodeType     PriceCodeType
	PriceCodeTypeName *PriceCodeTypeName
	PriceCode         PriceCode
	GeneralAttributes
}

type PriceCodeType struct {
	Value string `xml:",innerxml"` // List179
	GeneralAttributes
}

type PriceCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type PriceCondition struct {
	PriceConditionType     PriceConditionType
	PriceConditionQuantity []PriceConditionQuantity
	ProductIdentifier      []ProductIdentifier
	GeneralAttributes
}

type PriceConditionQuantity struct {
	PriceConditionQuantityType PriceConditionQuantityType
	Quantity                   Quantity
	QuantityUnit               QuantityUnit
	GeneralAttributes
}

type PriceConditionQuantityType struct {
	Value string `xml:",innerxml"` // List168
	GeneralAttributes
}

type PriceConditionType struct {
	Value string `xml:",innerxml"` // List167
	GeneralAttributes
}

type PriceDate struct {
	PriceDateRole PriceDateRole
	DateFormat    *DateFormat
	Date          Date
	GeneralAttributes
}

type PriceDateRole struct {
	Value string `xml:",innerxml"` // List173
	GeneralAttributes
}

type PriceIdentifier struct {
	PriceIDType PriceIDType
	IDTypeName  *IDTypeName
	IDValue     IDValue
	GeneralAttributes
}

type PriceIDType struct {
	Value string `xml:",innerxml"` // List217
	GeneralAttributes
}

type PricePer struct {
	Value string `xml:",innerxml"` // List60
	GeneralAttributes
}

type PriceQualifier struct {
	Value string `xml:",innerxml"` // List59
	GeneralAttributes
}

type PriceStatus struct {
	Value string `xml:",innerxml"` // List61
	GeneralAttributes
}

type PriceType struct {
	Value string `xml:",innerxml"` // List58
	GeneralAttributes
}

type PriceTypeDescription struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type PriceConstraint struct {
	PriceConstraintType   PriceConstraintType
	PriceConstraintStatus PriceConstraintStatus
	PriceConstraintLimit  []PriceConstraintLimit
	GeneralAttributes
}

type PriceConstraintLimit struct {
	Quantity            Quantity
	PriceConstraintUnit PriceConstraintUnit
	GeneralAttributes
}

type PriceConstraintStatus struct {
	Value string `xml:",innerxml"` // List146
	GeneralAttributes
}

type PriceConstraintType struct {
	Value string `xml:",innerxml"` // List230
	GeneralAttributes
}

type PriceConstraintUnit struct {
	Value string `xml:",innerxml"` // List147
	GeneralAttributes
}

type PrimaryContentType struct {
	Value string `xml:",innerxml"` // List81
	GeneralAttributes
}

type PrimaryPart struct {
	GeneralAttributes
}

type PrintedOnProduct struct {
	Value string `xml:",innerxml"` // List174
	GeneralAttributes
}

type Prize struct {
	PrizeName      []PrizeName
	PrizeYear      *PrizeYear
	PrizeCountry   *PrizeCountry
	PrizeCode      *PrizeCode
	PrizeStatement []PrizeStatement
	PrizeJury      []PrizeJury
	GeneralAttributes
}

type PrizeCode struct {
	Value string `xml:",innerxml"` // List41
	GeneralAttributes
}

type PrizeCountry struct {
	Value string `xml:",innerxml"` // List91
	GeneralAttributes
}

type PrizeJury struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type PrizeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type PrizeStatement struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type PrizeYear struct {
	Value string `xml:",innerxml"` // dt.Year
	GeneralAttributes
}

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

type ProductAvailability struct {
	Value string `xml:",innerxml"` // List65
	GeneralAttributes
}

type ProductClassification struct {
	ProductClassificationType ProductClassificationType
	ProductClassificationCode ProductClassificationCode
	Percent                   *Percent
	GeneralAttributes
}

type ProductClassificationCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ProductClassificationType struct {
	Value string `xml:",innerxml"` // List9
	GeneralAttributes
}

type ProductComposition struct {
	Value string `xml:",innerxml"` // List2
	GeneralAttributes
}

type ProductContact struct {
	ProductContactRole       ProductContactRole
	ContactName              *ContactName
	EmailAddress             *EmailAddress
	ProductContactIdentifier []ProductContactIdentifier
	ProductContactName       *ProductContactName
	GeneralAttributes
}

type ProductContactIdentifier struct {
	ProductContactIDType ProductContactIDType
	IDTypeName           *IDTypeName
	IDValue              IDValue
	GeneralAttributes
}

type ProductContactIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type ProductContactName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ProductContactRole struct {
	Value string `xml:",innerxml"` // List198
	GeneralAttributes
}

type ProductContentType struct {
	Value string `xml:",innerxml"` // List81
	GeneralAttributes
}

type ProductForm struct {
	Value string `xml:",innerxml"` // List150
	GeneralAttributes
}

type ProductFormDescription struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ProductFormDetail struct {
	Value string `xml:",innerxml"` // List175
	GeneralAttributes
}

type ProductFormFeature struct {
	ProductFormFeatureType        ProductFormFeatureType
	ProductFormFeatureValue       *ProductFormFeatureValue
	ProductFormFeatureDescription []ProductFormFeatureDescription
	GeneralAttributes
}

type ProductFormFeatureDescription struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ProductFormFeatureType struct {
	Value string `xml:",innerxml"` // List79
	GeneralAttributes
}

type ProductFormFeatureValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ProductIdentifier struct {
	ProductIDType ProductIDType
	IDTypeName    *IDTypeName
	IDValue       IDValue
	GeneralAttributes
}

type ProductIDType struct {
	Value string `xml:",innerxml"` // List5
	GeneralAttributes
}

type ProductPackaging struct {
	Value string `xml:",innerxml"` // List80
	GeneralAttributes
}

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

type ProductRelationCode struct {
	Value string `xml:",innerxml"` // List51
	GeneralAttributes
}

type ProductSupply struct {
	Market                 []Market
	MarketPublishingDetail *MarketPublishingDetail
	SupplyDetail           []SupplyDetail
	GeneralAttributes
}

type ProfessionalAffiliation struct {
	GeneralAttributes
}

type ProfessionalPosition struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type PromotionCampaign struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type PromotionContact struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type Proximity struct {
	Value string `xml:",innerxml"` // List215
	GeneralAttributes
}

type Publisher struct {
	PublishingRole      PublishingRole
	Funding             []Funding
	Website             []Website
	PublisherIdentifier []PublisherIdentifier
	PublisherName       *PublisherName
	GeneralAttributes
}

type PublisherIdentifier struct {
	PublisherIDType PublisherIDType
	IDTypeName      *IDTypeName
	IDValue         IDValue
	GeneralAttributes
}

type PublisherIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type PublisherName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

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

type PublishingDate struct {
	PublishingDateRole PublishingDateRole
	DateFormat         *DateFormat
	Date               Date
	GeneralAttributes
}

type PublishingDateRole struct {
	Value string `xml:",innerxml"` // List163
	GeneralAttributes
}

type PublishingDetail struct {
	CityOfPublication    []CityOfPublication
	CountryOfPublication *CountryOfPublication
	ProductContact       []ProductContact
	Imprint              []Imprint
	Publisher            []Publisher
	PublishingDate       []PublishingDate
	LatestReprintNumber  *LatestReprintNumber
	CopyrightStatement   []CopyrightStatement
	SalesRestriction     []SalesRestriction
	GeneralAttributes
}

type PublishingRole struct {
	Value string `xml:",innerxml"` // List45
	GeneralAttributes
}

type PublishingStatus struct {
	Value string `xml:",innerxml"` // List64
	GeneralAttributes
}

type PublishingStatusNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type Quantity struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

type QuantityUnit struct {
	Value string `xml:",innerxml"` // List169
	GeneralAttributes
}

type Rate struct {
	Value string `xml:",innerxml"` // dt.Integer
	GeneralAttributes
}

type ReviewRating struct {
	Rating      Rating
	RatingLimit *RatingLimit
	RatingUnits []RatingUnits
	GeneralAttributes
}

type Rating struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

type RatingLimit struct {
	Value string `xml:",innerxml"` // dt.PositiveInteger
	GeneralAttributes
}

type RatingUnits struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type RecordReference struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type RecordSourceIdentifier struct {
	RecordSourceIDType RecordSourceIDType
	IDTypeName         *IDTypeName
	IDValue            IDValue
	GeneralAttributes
}

type RecordSourceIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type RecordSourceName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type RecordSourceType struct {
	Value string `xml:",innerxml"` // List3
	GeneralAttributes
}

type RegionCode struct {
	Value string `xml:",innerxml"` // List49
	GeneralAttributes
}

type RegionsIncluded struct {
	Value string `xml:",innerxml"` // RegionCodeList
	GeneralAttributes
}

type RegionsExcluded struct {
	Value string `xml:",innerxml"` // RegionCodeList
	GeneralAttributes
}

type Reissue struct {
	ReissueDate        ReissueDate
	ReissueDescription *ReissueDescription
	Price              []Price
	SupportingResource []SupportingResource
	GeneralAttributes
}

type ReissueDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type ReissueDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type RelatedMaterial struct {
	RelatedWork    []RelatedWork
	RelatedProduct []RelatedProduct
	GeneralAttributes
}

type RelatedProduct struct {
	ProductRelationCode []ProductRelationCode
	ProductIdentifier   []ProductIdentifier
	GeneralAttributes
}

type RelatedWork struct {
	WorkRelationCode WorkRelationCode
	WorkIdentifier   []WorkIdentifier
	GeneralAttributes
}

type ReligiousText struct {
	GeneralAttributes
}

type ReligiousTextFeature struct {
	ReligiousTextFeatureType        ReligiousTextFeatureType
	ReligiousTextFeatureCode        ReligiousTextFeatureCode
	ReligiousTextFeatureDescription []ReligiousTextFeatureDescription
	GeneralAttributes
}

type ReligiousTextFeatureCode struct {
	Value string `xml:",innerxml"` // List90
	GeneralAttributes
}

type ReligiousTextFeatureDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type ReligiousTextFeatureType struct {
	Value string `xml:",innerxml"` // List89
	GeneralAttributes
}

type ReligiousTextIdentifier struct {
	Value string `xml:",innerxml"` // List88
	GeneralAttributes
}

type ReprintDetail struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type ResourceContentType struct {
	Value string `xml:",innerxml"` // List158
	GeneralAttributes
}

type ResourceFeature struct {
	ResourceFeatureType ResourceFeatureType
	FeatureValue        *FeatureValue
	FeatureNote         []FeatureNote
	GeneralAttributes
}

type ResourceFeatureType struct {
	Value string `xml:",innerxml"` // List160
	GeneralAttributes
}

type ResourceForm struct {
	Value string `xml:",innerxml"` // List161
	GeneralAttributes
}

type ResourceLink struct {
	Value string `xml:",innerxml"` // dt.NonEmptyURI
	GeneralAttributes
	LanguageAttribute
}

type ResourceMode struct {
	Value string `xml:",innerxml"` // List159
	GeneralAttributes
}

type ResourceVersion struct {
	ResourceForm           ResourceForm
	ResourceVersionFeature []ResourceVersionFeature
	ResourceLink           []ResourceLink
	ContentDate            []ContentDate
	GeneralAttributes
}

type ResourceVersionFeature struct {
	ResourceVersionFeatureType ResourceVersionFeatureType
	FeatureValue               *FeatureValue
	FeatureNote                []FeatureNote
	GeneralAttributes
}

type ResourceVersionFeatureType struct {
	Value string `xml:",innerxml"` // List162
	GeneralAttributes
}

type ROWSalesRightsType struct {
	Value string `xml:",innerxml"` // List46
	GeneralAttributes
}

type ReturnsCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type ReturnsCodeType struct {
	Value string `xml:",innerxml"` // List53
	GeneralAttributes
}

type ReturnsCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ReturnsConditions struct {
	ReturnsCodeType     ReturnsCodeType
	ReturnsCodeTypeName *ReturnsCodeTypeName
	ReturnsCode         ReturnsCode
	ReturnsNote         []ReturnsNote
	GeneralAttributes
}

type ReturnsNote struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SalesOutlet struct {
	SalesOutletIdentifier []SalesOutletIdentifier
	SalesOutletName       *SalesOutletName
	GeneralAttributes
}

type SalesOutletIdentifier struct {
	SalesOutletIDType SalesOutletIDType
	IDTypeName        *IDTypeName
	IDValue           IDValue
	GeneralAttributes
}

type SalesOutletIDType struct {
	Value string `xml:",innerxml"` // List102
	GeneralAttributes
}

type SalesOutletName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SalesRestriction struct {
	SalesRestrictionType SalesRestrictionType
	SalesOutlet          []SalesOutlet
	SalesRestrictionNote []SalesRestrictionNote
	StartDate            *StartDate
	EndDate              *EndDate
	GeneralAttributes
}

type SalesRestrictionNote struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type SalesRestrictionType struct {
	Value string `xml:",innerxml"` // List71
	GeneralAttributes
}

type SalesRights struct {
	SalesRightsType   SalesRightsType
	Territory         Territory
	SalesRestriction  []SalesRestriction
	ProductIdentifier []ProductIdentifier
	PublisherName     *PublisherName
	GeneralAttributes
}

type SalesRightsType struct {
	Value string `xml:",innerxml"` // List46
	GeneralAttributes
}

type ScriptCode struct {
	Value string `xml:",innerxml"` // List121
	GeneralAttributes
}

type Sender struct {
	ContactName      *ContactName
	EmailAddress     *EmailAddress
	SenderIdentifier []SenderIdentifier
	SenderName       *SenderName
	GeneralAttributes
}

type SenderIDType struct {
	Value string `xml:",innerxml"` // List44
	GeneralAttributes
}

type SenderIdentifier struct {
	SenderIDType SenderIDType
	IDTypeName   *IDTypeName
	IDValue      IDValue
	GeneralAttributes
}

type SenderName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type SentDateTime struct {
	Value string `xml:",innerxml"` // dt.DateOrDateTime
	GeneralAttributes
}

type SequenceNumber struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveInteger
	GeneralAttributes
}

type SourceName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type SourceTitle struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SourceType struct {
	Value string `xml:",innerxml"` // List157
	GeneralAttributes
}

type StartDate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

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

type StockQuantityCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type StockQuantityCoded struct {
	StockQuantityCodeType     StockQuantityCodeType
	StockQuantityCodeTypeName *StockQuantityCodeTypeName
	StockQuantityCode         StockQuantityCode
	GeneralAttributes
}

type StockQuantityCodeType struct {
	Value string `xml:",innerxml"` // List70
	GeneralAttributes
}

type StockQuantityCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type StudyBibleType struct {
	Value string `xml:",innerxml"` // List84
	GeneralAttributes
}

type Subject struct {
	MainSubject             *MainSubject
	SubjectSchemeIdentifier SubjectSchemeIdentifier
	SubjectSchemeName       *SubjectSchemeName
	SubjectSchemeVersion    *SubjectSchemeVersion
	SubjectCode             SubjectCode
	SubjectHeadingText      []SubjectHeadingText
	GeneralAttributes
}

type SubjectCode struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type SubjectDate struct {
	SubjectDateRole SubjectDateRole
	DateFormat      *DateFormat
	Date            Date
	GeneralAttributes
}

type SubjectDateRole struct {
	Value string `xml:",innerxml"` // List177
	GeneralAttributes
}

type SubjectHeadingText struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SubjectSchemeIdentifier struct {
	Value string `xml:",innerxml"` // List27
	GeneralAttributes
}

type SubjectSchemeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SubjectSchemeVersion struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type Subtitle struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

type SuffixToKey struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type Supplier struct {
	SupplierRole       SupplierRole
	TelephoneNumber    []TelephoneNumber
	FaxNumber          []FaxNumber
	EmailAddress       []EmailAddress
	Website            []Website
	SupplierIdentifier []SupplierIdentifier
	SupplierName       *SupplierName
	GeneralAttributes
}

type SupplierCodeType struct {
	Value string `xml:",innerxml"` // List165
	GeneralAttributes
}

type SupplierCodeTypeName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SupplierCodeValue struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type SupplierIdentifier struct {
	SupplierIDType SupplierIDType
	IDTypeName     *IDTypeName
	IDValue        IDValue
	GeneralAttributes
}

type SupplierIDType struct {
	Value string `xml:",innerxml"` // List92
	GeneralAttributes
}

type SupplierName struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type SupplierOwnCoding struct {
	SupplierCodeType     SupplierCodeType
	SupplierCodeTypeName *SupplierCodeTypeName
	SupplierCodeValue    SupplierCodeValue
	GeneralAttributes
}

type SupplierRole struct {
	Value string `xml:",innerxml"` // List93
	GeneralAttributes
}

type SupplyDate struct {
	SupplyDateRole SupplyDateRole
	DateFormat     *DateFormat
	Date           Date
	GeneralAttributes
}

type SupplyDateRole struct {
	Value string `xml:",innerxml"` // List166
	GeneralAttributes
}

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
	UnpricedItemType    UnpricedItemType
	Price               []Price
	GeneralAttributes
}

type SupportingResource struct {
	ResourceContentType ResourceContentType
	ContentAudience     []ContentAudience
	Territory           *Territory
	ResourceMode        ResourceMode
	ResourceFeature     []ResourceFeature
	ResourceVersion     []ResourceVersion
	GeneralAttributes
}

type Tax struct {
	ProductIdentifier []ProductIdentifier
	TaxType           *TaxType
	TaxRateCode       *TaxRateCode
	TaxRatePercent    TaxRatePercent
	TaxableAmount     *TaxableAmount
	TaxAmount         *TaxAmount
	GeneralAttributes
}

type TaxAmount struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

type TaxRateCode struct {
	Value string `xml:",innerxml"` // List62
	GeneralAttributes
}

type TaxRatePercent struct {
	Value string `xml:",innerxml"` // dt.PercentDecimal
	GeneralAttributes
}

type TaxableAmount struct {
	Value string `xml:",innerxml"` // dt.StrictPositiveDecimal
	GeneralAttributes
}

type TaxType struct {
	Value string `xml:",innerxml"` // List171
	GeneralAttributes
}

type TelephoneNumber struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
}

type Territory struct {
	RegionsExcluded   *RegionsExcluded
	CountriesIncluded CountriesIncluded
	RegionsIncluded   RegionsIncluded
	CountriesExcluded *CountriesExcluded
	GeneralAttributes
}

type Text struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type TextAuthor struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

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

type TextItem struct {
	TextItemType       TextItemType
	TextItemIdentifier []TextItemIdentifier
	PageRun            []PageRun
	NumberOfPages      *NumberOfPages
	GeneralAttributes
}

type TextItemIDType struct {
	Value string `xml:",innerxml"` // List43
	GeneralAttributes
}

type TextItemIdentifier struct {
	TextItemIDType TextItemIDType
	IDTypeName     *IDTypeName
	IDValue        IDValue
	GeneralAttributes
}

type TextItemType struct {
	Value string `xml:",innerxml"` // List42
	GeneralAttributes
}

type TextSourceCorporate struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type TextType struct {
	Value string `xml:",innerxml"` // List153
	GeneralAttributes
}

type ThesisPresentedTo struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	LanguageAttribute
}

type ThesisType struct {
	Value string `xml:",innerxml"` // List72
	GeneralAttributes
}

type ThesisYear struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	DateformatAttribute
}

type TitleDetail struct {
	TitleType      TitleType
	TitleElement   []TitleElement
	TitleStatement *TitleStatement
	GeneralAttributes
}

type TitleElement struct {
	SequenceNumber     *SequenceNumber
	TitleElementLevel  TitleElementLevel
	Subtitle           *Subtitle
	PartNumber         PartNumber
	YearOfAnnual       *YearOfAnnual
	TitleWithoutPrefix TitleWithoutPrefix
	TitlePrefix        TitlePrefix
	NoPrefix           NoPrefix
	TitleText          TitleText
	GeneralAttributes
}

type TitleElementLevel struct {
	Value string `xml:",innerxml"` // List149
	GeneralAttributes
}

type TitlePrefix struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

type TitlesAfterNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type TitlesBeforeNames struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	TextscriptAttribute
	LanguageAttribute
}

type TitleStatement struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type TitleText struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

type TitleType struct {
	Value string `xml:",innerxml"` // List15
	GeneralAttributes
}

type TitleWithoutPrefix struct {
	Value string `xml:",innerxml"` // dt.NonEmptyString
	GeneralAttributes
	CollationkeyAttribute
	LanguageAttribute
	TextscriptAttribute
	TextcaseAttribute
}

type ToLanguage struct {
	Value string `xml:",innerxml"` // List74
	GeneralAttributes
}

type ToQuantity struct {
	Value string `xml:",innerxml"` // dt.PositiveDecimal
	GeneralAttributes
}

type TradeCategory struct {
	Value string `xml:",innerxml"` // List12
	GeneralAttributes
}

type UnnamedPersons struct {
	Value string `xml:",innerxml"` // List19
	GeneralAttributes
}

type UnpricedItemType struct {
	Value string `xml:",innerxml"` // List57
	GeneralAttributes
}

type Velocity struct {
	VelocityMetric VelocityMetric
	Rate           Rate
	Proximity      *Proximity
	GeneralAttributes
}

type VelocityMetric struct {
	Value string `xml:",innerxml"` // List216
	GeneralAttributes
}

type Website struct {
	WebsiteRole        *WebsiteRole
	WebsiteDescription []WebsiteDescription
	WebsiteLink        WebsiteLink
	GeneralAttributes
}

type WebsiteDescription struct {
	Value string `xml:",innerxml"` // HTML
	GeneralAttributes
	LanguageAttribute
	TextformatAttribute
}

type WebsiteLink struct {
	Value string `xml:",innerxml"` // dt.NonEmptyURI
	GeneralAttributes
}

type WebsiteRole struct {
	Value string `xml:",innerxml"` // List73
	GeneralAttributes
}

type WorkIdentifier struct {
	WorkIDType WorkIDType
	IDTypeName *IDTypeName
	IDValue    IDValue
	GeneralAttributes
}

type WorkIDType struct {
	Value string `xml:",innerxml"` // List16
	GeneralAttributes
}

type WorkRelationCode struct {
	Value string `xml:",innerxml"` // List164
	GeneralAttributes
}

type YearOfAnnual struct {
	Value string `xml:",innerxml"` // dt.YearOrYearRange
	GeneralAttributes
}

type ReleaseAttribute struct {
	Release string `xml:"release,attr,omitempty"`
}

type CollationkeyAttribute struct {
	Collationkey string `xml:"collationkey,attr,omitempty"` // xs:string
}

type GeneralAttributes struct {
	Datestamp  string `xml:"datestamp,attr,omitempty"`  // dt.DateOrDateTime
	Sourcetype string `xml:"sourcetype,attr,omitempty"` // SourceTypeCode
	Sourcename string `xml:"sourcename,attr,omitempty"`
}

type DateformatAttribute struct {
	Dateformat string `xml:"dateformat,attr,omitempty"` // List55
}

type LanguageAttribute struct {
	Language string `xml:"language,attr,omitempty"` // List74
}

type TextcaseAttribute struct {
	Textcase string `xml:"textcase,attr,omitempty"` // TextCaseCode
}

type TextformatAttribute struct {
	Textformat string `xml:"textformat,attr,omitempty"` // TextFormatCode
}

type TextscriptAttribute struct {
	Textscript string `xml:"textscript,attr,omitempty"` // List121
}
