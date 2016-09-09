package ncip

import "encoding/xml"

type SchemeValue struct {
	Scheme string `xml:",attr,omitempty"`
	Value  string `xml:",innerxml"`
}

type AcknowledgedItemUseRestrictionType SchemeValue
type AgencyUserPrivilegeType SchemeValue
type AuthenticationDataFormatType SchemeValue
type ComponentIdentifierType SchemeValue
type ConsortiumAgreement SchemeValue
type RequestIdentifierType SchemeValue
type UserIdentifierType SchemeValue
type ApplicationProfileSupportedType SchemeValue
type BibliographicRecordIdentifierCode SchemeValue
type PickupLocation SchemeValue
type RequestType SchemeValue
type ToSystemId SchemeValue
type UserAddressRoleType SchemeValue
type AgencyElementType SchemeValue
type CirculationStatus SchemeValue
type ElectronicDataFormatType SchemeValue
type FiscalActionType SchemeValue
type ProblemType SchemeValue
type RequestStatusType SchemeValue
type ApplicationProfileType SchemeValue
type AuthenticationInputType SchemeValue
type ElectronicAddressType SchemeValue
type PhysicalAddressType SchemeValue
type RequestedActionType SchemeValue
type AgencyAddressRoleType SchemeValue
type ItemIdentifierType SchemeValue
type PhysicalConditionType SchemeValue
type AuthenticationPromptType SchemeValue
type BibliographicLevel SchemeValue
type BlockOrTrapType SchemeValue
type ItemElementType SchemeValue
type MediumType SchemeValue
type SecurityMarker SchemeValue
type UserPrivilegeStatusType SchemeValue
type ItemDescriptionLevel SchemeValue
type ItemUseRestrictionType SchemeValue
type NoticeType SchemeValue
type OrganizationNameType SchemeValue
type RequestElementType SchemeValue
type UserLanguage SchemeValue
type BibliographicItemIdentifierCode SchemeValue
type CurrencyCode SchemeValue
type FiscalTransactionType SchemeValue
type FromSystemId SchemeValue
type Language SchemeValue
type LocationType SchemeValue
type PaymentMethodType SchemeValue
type RequestScopeType SchemeValue
type RequiredItemUseRestrictionType SchemeValue
type UnstructuredAddressType SchemeValue

type AcceptItem struct {
	XMLName             xml.Name
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	RequestId           RequestId
	RequestedActionType RequestedActionType
	UserId              *UserId
	ItemId              *ItemId
	// xs:choice ->
	DateForReturn               string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	NonReturnableFlag           *NonReturnableFlag
	// <- xs:choice
	RenewalNotPermitted          *RenewalNotPermitted
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	PickupLocation               *PickupLocation
	PickupExpiryDate             string `xml:",omitempty"` // xs:dateTime
	Ext                          *Ext
}

type AcceptItemResponse struct {
	XMLName        xml.Name
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type AgencyCreated struct {
	XMLName                         xml.Name
	InitiationHeader                *InitiationHeader
	AgencyId                        SchemeValue
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []ApplicationProfileSupportedType
	ConsortiumAgreement             []ConsortiumAgreement
	AgencyUserPrivilegeType         []AgencyUserPrivilegeType
	Ext                             *Ext
}

type AgencyCreatedResponse struct {
	XMLName        xml.Name
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type AgencyUpdated struct {
	InitiationHeader   *InitiationHeader
	AgencyId           SchemeValue
	DeleteAgencyFields *DeleteAgencyFields
	AddAgencyFields    *AddAgencyFields
	Ext                *Ext
}

type AgencyUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type CancelRecallItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	ItemElementType  []ItemElementType
	UserElementType  []SchemeValue
	Ext              *Ext
}

type CancelRecallItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	ItemId                       ItemId
	UserId                       *UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type CancelRequestItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	RequestType           RequestType
	RequestScopeType      *RequestScopeType
	AcknowledgedFeeAmount *AcknowledgedFeeAmount
	PaidFeeAmount         *PaidFeeAmount
	ItemElementType       []ItemElementType
	UserElementType       []SchemeValue
	Ext                   *Ext
}

type CancelRequestItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	UserId                       UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type CheckInItem struct {
	XMLName          xml.Name
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	ItemElementType  []ItemElementType
	UserElementType  []SchemeValue
	Ext              *Ext
}

type CheckInItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	ItemId                       ItemId
	UserId                       *UserId
	RoutingInformation           *RoutingInformation
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type CheckOutItem struct {
	XMLName          xml.Name
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	ItemId                             ItemId
	RequestId                          *RequestId
	AcknowledgedFeeAmount              *AcknowledgedFeeAmount
	PaidFeeAmount                      *PaidFeeAmount
	AcknowledgedItemUseRestrictionType []AcknowledgedItemUseRestrictionType
	ShippingInformation                *ShippingInformation
	ResourceDesired                    *ResourceDesired
	DesiredDateDue                     string `xml:",omitempty"` // xs:dateTime
	ItemElementType                    []ItemElementType
	UserElementType                    []SchemeValue
	Ext                                *Ext
}

type CheckOutItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	// xs:sequence ->
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []RequiredItemUseRestrictionType
	// <- xs:sequence
	// xs:sequence ->
	ItemId ItemId
	UserId UserId
	// xs:choice ->
	DateDue                     string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	NonReturnableFlag           *NonReturnableFlag
	// <- xs:choice
	RenewalCount                 *int // xs:nonnegativeInteger
	ElectronicResource           *ElectronicResource
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type CirculationStatusChangeReported struct {
	InitiationHeader *InitiationHeader
	ItemId           ItemId
	UserId           UserId
	// xs:choice ->
	ItemReportedReturned      *ItemReportedReturned
	ItemReportedNeverBorrowed *ItemReportedNeverBorrowed
	ItemReportedLost          *ItemReportedLost
	ItemReportedMissing       *ItemReportedMissing
	// <- xs:choice
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	FiscalTransactionInformation *FiscalTransactionInformation
	Ext                          *Ext
}

type CirculationStatusChangeReportedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type CirculationStatusUpdated struct {
	InitiationHeader  *InitiationHeader
	ItemId            ItemId
	CirculationStatus CirculationStatus
	Ext               *Ext
}

type CirculationStatusUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type CreateAgency struct {
	InitiationHeader                *InitiationHeader
	MandatedAction                  *MandatedAction
	AgencyId                        *SchemeValue
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []ApplicationProfileSupportedType
	ConsortiumAgreement             []ConsortiumAgreement
	AgencyUserPrivilegeType         []AgencyUserPrivilegeType
	Ext                             *Ext
}

type CreateAgencyResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem  []Problem
	AgencyId *SchemeValue
	// <- xs:choice
	Ext *Ext
}

type CreateItem struct {
	InitiationHeader         *InitiationHeader
	MandatedAction           *MandatedAction
	ItemId                   *ItemId
	RequestId                *RequestId
	BibliographicDescription BibliographicDescription
	ItemUseRestrictionType   []ItemUseRestrictionType
	CirculationStatus        *CirculationStatus
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SecurityMarker
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type CreateItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	ItemId  *ItemId
	// <- xs:choice
	Ext *Ext
}

type CreateUser struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	NameInformation        NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []UserLanguage
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type CreateUserResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	UserId  *UserId
	// <- xs:choice
	Ext *Ext
}

type CreateUserFiscalTransaction struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

type CreateUserFiscalTransactionResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	UserId                       UserId
	FiscalTransactionReferenceId FiscalTransactionReferenceId
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type DeleteItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	Ext              *Ext
}

type DeleteItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	ItemId  *ItemId
	// <- xs:choice
	Ext *Ext
}

type DeleteUser struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	Ext *Ext
}

type DeleteUserResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	UserId  *UserId
	// <- xs:choice
	Ext *Ext
}

type InitiationHeader struct {
	FromSystemId             *FromSystemId
	FromSystemAuthentication string `xml:",omitempty"`
	FromAgencyId             FromAgencyId
	FromAgencyAuthentication string `xml:",omitempty"`
	OnBehalfOfAgency         *OnBehalfOfAgency
	ToSystemId               *ToSystemId
	ToAgencyId               ToAgencyId
	ApplicationProfileType   *ApplicationProfileType
	Ext                      *Ext
}

type ItemCheckedIn struct {
	InitiationHeader             *InitiationHeader
	UserId                       *UserId
	ItemId                       ItemId
	RoutingInformation           *RoutingInformation
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type ItemCheckedInResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemCheckedOut struct {
	InitiationHeader *InitiationHeader
	UserId           UserId
	ItemId           ItemId
	RequestId        *RequestId
	// xs:choice ->
	DateDue                     string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	NonReturnableFlag           *NonReturnableFlag
	// <- xs:choice
	ElectronicResourceProvidedFlag *ElectronicResourceProvidedFlag
	RenewalCount                   *int // xs:nonnegativeInteger
	FiscalTransactionInformation   *FiscalTransactionInformation
	ShippingInformation            *ShippingInformation
	ItemOptionalFields             *ItemOptionalFields
	UserOptionalFields             *UserOptionalFields
	Ext                            *Ext
}

type ItemCheckedOutResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemCreated struct {
	InitiationHeader         *InitiationHeader
	ItemId                   ItemId
	RequestId                *RequestId
	BibliographicDescription BibliographicDescription
	ItemUseRestrictionType   []ItemUseRestrictionType
	CirculationStatus        *CirculationStatus
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SecurityMarker
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type ItemCreatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemRecallCancelled struct {
	InitiationHeader             *InitiationHeader
	UserId                       *UserId
	ItemId                       ItemId
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type ItemRecallCancelledResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemRecalled struct {
	InitiationHeader    *InitiationHeader
	UserId              *UserId
	ItemId              ItemId
	DateDue             string // xs:dateTime
	ShippingInformation *ShippingInformation
	ItemOptionalFields  *ItemOptionalFields
	UserOptionalFields  *UserOptionalFields
	Ext                 *Ext
}

type ItemRecalledResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemReceived struct {
	InitiationHeader   *InitiationHeader
	ItemId             ItemId
	UserId             *UserId
	RequestId          *RequestId
	DateReceived       string // xs:dateTime
	ItemOptionalFields *ItemOptionalFields
	UserOptionalFields *UserOptionalFields
	Ext                *Ext
}

type ItemReceivedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemRenewed struct {
	InitiationHeader             *InitiationHeader
	UserId                       *UserId
	ItemId                       ItemId
	DateDue                      string // xs:dateTime
	RenewalCount                 *int   // xs:nonnegativeInteger
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type ItemRenewedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemRequestCancelled struct {
	InitiationHeader *InitiationHeader
	UserId           UserId
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	RequestType                  RequestType
	RequestScopeType             *RequestScopeType
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type ItemRequestCancelledResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemRequestUpdated struct {
	InitiationHeader *InitiationHeader
	// xs:choice ->
	// xs:sequence ->
	UserId      UserId
	ItemId      ItemId
	RequestType RequestType
	// <- xs:sequence
	RequestId *RequestId
	// <- xs:choice
	DeleteRequestFields *DeleteRequestFields
	AddRequestFields    *AddRequestFields
	ItemOptionalFields  *ItemOptionalFields
	UserOptionalFields  *UserOptionalFields
	Ext                 *Ext
}

type ItemRequestUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemRequested struct {
	InitiationHeader *InitiationHeader
	UserId           UserId
	// xs:choice ->
	ItemId *ItemId
	// xs:sequence ->
	BibliographicId BibliographicId
	RequestId       RequestId
	// <- xs:sequence
	// <- xs:choice
	RequestType         RequestType
	RequestScopeType    RequestScopeType
	ShippingInformation *ShippingInformation
	EarliestDateNeeded  string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate      string `xml:",omitempty"` // xs:dateTime
	PickupLocation      *PickupLocation
	PickupExpiryDate    string `xml:",omitempty"` // xs:dateTime
	DateOfUserRequest   string `xml:",omitempty"` // xs:dateTime
	DateAvailable       string `xml:",omitempty"` // xs:dateTime
	ItemOptionalFields  *ItemOptionalFields
	UserOptionalFields  *UserOptionalFields
	Ext                 *Ext
}

type ItemRequestedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemShipped struct {
	InitiationHeader *InitiationHeader
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	UserId              *UserId
	DateShipped         string // xs:dateTime
	ShippingInformation ShippingInformation
	ItemOptionalFields  *ItemOptionalFields
	UserOptionalFields  *UserOptionalFields
	Ext                 *Ext
}

type ItemShippedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type ItemUpdated struct {
	InitiationHeader *InitiationHeader
	ItemId           ItemId
	DeleteItemFields *DeleteItemFields
	AddItemFields    *AddItemFields
	Ext              *Ext
}

type ItemUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type LookupAgency struct {
	InitiationHeader  *InitiationHeader
	AgencyId          SchemeValue
	AgencyElementType []AgencyElementType
	Ext               *Ext
}

type LookupAgencyResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	AgencyId                        SchemeValue
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []ApplicationProfileSupportedType
	ConsortiumAgreement             []ConsortiumAgreement
	AgencyUserPrivilegeType         []AgencyUserPrivilegeType
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type LookupItem struct {
	InitiationHeader *InitiationHeader
	// xs:choice ->
	ItemId    *ItemId
	RequestId *RequestId
	// <- xs:choice
	ItemElementType          []ItemElementType
	CurrentBorrowerDesired   *CurrentBorrowerDesired
	CurrentRequestersDesired *CurrentRequestersDesired
	Ext                      *Ext
}

type LookupItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	HoldPickupDate     string `xml:",omitempty"` // xs:dateTime
	DateRecalled       string `xml:",omitempty"` // xs:dateTime
	ItemTransaction    *ItemTransaction
	ItemOptionalFields *ItemOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type LookupRequest struct {
	InitiationHeader *InitiationHeader
	// xs:choice ->
	// xs:sequence ->
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	ItemId      ItemId
	RequestType RequestType
	// <- xs:sequence
	RequestId *RequestId
	// <- xs:choice
	RequestElementType []RequestElementType
	ItemElementType    []ItemElementType
	UserElementType    []SchemeValue
	Ext                *Ext
}

type LookupRequestResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	UserId                *UserId
	RequestType           *RequestType
	RequestScopeType      *RequestScopeType
	RequestStatusType     *RequestStatusType
	HoldQueuePosition     *int // xs:positiveInteger
	ShippingInformation   *ShippingInformation
	EarliestDateNeeded    string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate        string `xml:",omitempty"` // xs:dateTime
	PickupDate            string `xml:",omitempty"` // xs:dateTime
	PickupLocation        *PickupLocation
	PickupExpiryDate      string `xml:",omitempty"` // xs:dateTime
	DateOfUserRequest     string `xml:",omitempty"` // xs:dateTime
	DateAvailable         string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount *AcknowledgedFeeAmount
	PaidFeeAmount         *PaidFeeAmount
	ItemOptionalFields    *ItemOptionalFields
	UserOptionalFields    *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type LookupUser struct {
	XMLName          xml.Name
	InitiationHeader *InitiationHeader
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	UserElementType          []SchemeValue
	LoanedItemsDesired       *LoanedItemsDesired       `xml:",omitempty"`
	RequestedItemsDesired    *RequestedItemsDesired    `xml:",omitempty"`
	UserFiscalAccountDesired *UserFiscalAccountDesired `xml:",omitempty"`
	Ext                      *Ext
}

type LookupUserResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	UserId              UserId
	UserFiscalAccount   []UserFiscalAccount
	LoanedItemsCount    []LoanedItemsCount
	LoanedItem          []LoanedItem
	RequestedItemsCount []RequestedItemsCount
	RequestedItem       []RequestedItem
	UserOptionalFields  *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type RecallItem struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	ItemId              ItemId
	DesiredDateDue      string `xml:",omitempty"` // xs:dateTime
	ShippingInformation *ShippingInformation
	ItemElementType     []ItemElementType
	UserElementType     []SchemeValue
	Ext                 *Ext
}

type RecallItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	ItemId                       ItemId
	UserId                       *UserId
	DateDue                      string `xml:",omitempty"` // xs:dateTime
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type RenewItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	ItemId                             ItemId
	ItemElementType                    []ItemElementType
	UserElementType                    []SchemeValue
	DesiredDateDue                     string `xml:",omitempty"` // xs:dateTime
	DesiredDateForReturn               string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount              *AcknowledgedFeeAmount
	PaidFeeAmount                      *PaidFeeAmount
	AcknowledgedItemUseRestrictionType []AcknowledgedItemUseRestrictionType
	Ext                                *Ext
}

type RenewItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	// xs:sequence ->
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []RequiredItemUseRestrictionType
	// <- xs:sequence
	Pending *Pending
	// xs:sequence ->
	ItemId                       ItemId
	UserId                       *UserId
	DateDue                      string `xml:",omitempty"` // xs:dateTime
	DateForReturn                string `xml:",omitempty"` // xs:dateTime
	RenewalCount                 *int   // xs:nonnegativeInteger
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type ReportCirculationStatusChange struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	// xs:choice ->
	ItemReportedReturned      *ItemReportedReturned
	ItemReportedNeverBorrowed *ItemReportedNeverBorrowed
	ItemReportedLost          *ItemReportedLost
	ItemReportedMissing       *ItemReportedMissing
	// <- xs:choice
	ItemElementType []ItemElementType
	UserElementType []SchemeValue
	Ext             *Ext
}

type ReportCirculationStatusChangeResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	ItemId             ItemId
	UserId             *UserId
	ItemOptionalFields *ItemOptionalFields
	UserOptionalFields *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type RequestItem struct {
	XMLName          xml.Name
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	// xs:choice ->
	//ItemId []ItemId
	// xs:sequence ->
	BibliographicId []BibliographicId
	ItemId          []ItemId
	// <- xs:sequence
	// <- xs:choice
	RequestId                          *RequestId
	RequestType                        RequestType
	RequestScopeType                   RequestScopeType
	ItemOptionalFields                 *ItemOptionalFields
	ShippingInformation                *ShippingInformation
	EarliestDateNeeded                 string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate                     string `xml:",omitempty"` // xs:dateTime
	PickupLocation                     *PickupLocation
	PickupExpiryDate                   string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount              *AcknowledgedFeeAmount
	PaidFeeAmount                      *PaidFeeAmount
	AcknowledgedItemUseRestrictionType []AcknowledgedItemUseRestrictionType
	ItemElementType                    []ItemElementType
	UserElementType                    []SchemeValue
	Ext                                *Ext
}

type RequestItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	// xs:sequence ->
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []RequiredItemUseRestrictionType
	// <- xs:sequence
	// xs:sequence ->
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	UserId                       UserId
	RequestType                  RequestType
	RequestScopeType             RequestScopeType
	ShippingInformation          *ShippingInformation
	DateAvailable                string `xml:",omitempty"` // xs:dateTime
	HoldPickupDate               string `xml:",omitempty"` // xs:dateTime
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type ResponseHeader struct {
	FromSystemId             *FromSystemId
	FromSystemAuthentication string `xml:",omitempty"`
	FromAgencyId             FromAgencyId
	FromAgencyAuthentication string `xml:",omitempty"`
	ToSystemId               *ToSystemId
	ToAgencyId               ToAgencyId
	Ext                      *Ext
}

type SendUserNotice struct {
	InitiationHeader  *InitiationHeader
	MandatedAction    *MandatedAction
	UserId            UserId
	DateToSend        string `xml:",omitempty"` // xs:dateTime
	UserNoticeDetails UserNoticeDetails
	Ext               *Ext
}

type SendUserNoticeResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	UserId UserId
	// xs:choice ->
	DateSent     string `xml:",omitempty"` // xs:dateTime
	DateWillSend string `xml:",omitempty"` // xs:dateTime
	// <- xs:choice
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type UndoCheckOutItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	RequestId *RequestId
	Ext       *Ext
}

type UndoCheckOutItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	// xs:sequence ->
	ItemId                       ItemId
	UserId                       *UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type UpdateAgency struct {
	InitiationHeader   *InitiationHeader
	MandatedAction     *MandatedAction
	AgencyId           SchemeValue
	DeleteAgencyFields *DeleteAgencyFields
	AddAgencyFields    *AddAgencyFields
	Ext                *Ext
}

type UpdateAgencyResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem  []Problem
	AgencyId *SchemeValue
	// <- xs:choice
	Ext *Ext
}

type UpdateCirculationStatus struct {
	InitiationHeader  *InitiationHeader
	MandatedAction    *MandatedAction
	ItemId            ItemId
	CirculationStatus CirculationStatus
	Ext               *Ext
}

type UpdateCirculationStatusResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	ItemId  *ItemId
	// <- xs:choice
	Ext *Ext
}

type UpdateItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	DeleteItemFields *DeleteItemFields
	AddItemFields    *AddItemFields
	Ext              *Ext
}

type UpdateItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	ItemId  *ItemId
	// <- xs:choice
	Ext *Ext
}

type UpdateRequestItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	// xs:sequence ->
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	ItemId      ItemId
	RequestType RequestType
	// <- xs:sequence
	RequestId *RequestId
	// <- xs:choice
	DeleteRequestFields *DeleteRequestFields
	AddRequestFields    *AddRequestFields
	ItemElementType     []ItemElementType
	UserElementType     []SchemeValue
	Ext                 *Ext
}

type UpdateRequestItemResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	// xs:sequence ->
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []RequiredItemUseRestrictionType
	// <- xs:sequence
	// xs:sequence ->
	ItemId                       ItemId
	UserId                       UserId
	RequestType                  RequestType
	RequestScopeType             RequestScopeType
	DateAvailable                string `xml:",omitempty"` // xs:dateTime
	HoldPickupDate               string `xml:",omitempty"` // xs:dateTime
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type UpdateUser struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	// xs:choice ->
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	// <- xs:choice
	DeleteUserFields *DeleteUserFields
	AddUserFields    *AddUserFields
	Ext              *Ext
}

type UpdateUserResponse struct {
	ResponseHeader *ResponseHeader
	// xs:choice ->
	Problem []Problem
	UserId  *UserId
	// <- xs:choice
	Ext *Ext
}

type UserCreated struct {
	InitiationHeader       *InitiationHeader
	UserId                 UserId
	NameInformation        NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []UserLanguage
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type UserCreatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type UserFiscalTransactionCreated struct {
	InitiationHeader             *InitiationHeader
	UserId                       UserId
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

type UserFiscalTransactionCreatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type UserNoticeSent struct {
	InitiationHeader  *InitiationHeader
	UserId            UserId
	DateSent          string // xs:dateTime
	UserNoticeDetails UserNoticeDetails
	Ext               *Ext
}

type UserNoticeSentResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type UserUpdated struct {
	InitiationHeader *InitiationHeader
	UserId           UserId
	DeleteUserFields *DeleteUserFields
	AddUserFields    *AddUserFields
	Ext              *Ext
}

type UserUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

type AccountBalance struct {
	CurrencyCode  CurrencyCode
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type AccountDetails struct {
	AccrualDate                  string // xs:dateTime
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

type AcknowledgedFeeAmount struct {
	CurrencyCode  CurrencyCode
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type AddAgencyFields struct {
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []ApplicationProfileSupportedType
	ConsortiumAgreement             []ConsortiumAgreement
	AgencyUserPrivilegeType         []AgencyUserPrivilegeType
	Ext                             *Ext
}

type AddItemFields struct {
	BibliographicDescription *BibliographicDescription
	ItemUseRestrictionType   []ItemUseRestrictionType
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SecurityMarker
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type AddRequestFields struct {
	UserId                *UserId
	ItemId                *ItemId
	RequestType           *RequestType
	RequestScopeType      *RequestScopeType
	RequestStatusType     *RequestStatusType
	ShippingInformation   *ShippingInformation
	EarliestDateNeeded    string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate        string `xml:",omitempty"` // xs:dateTime
	PickupLocation        *PickupLocation
	PickupExpiryDate      string `xml:",omitempty"` // xs:dateTime
	DateOfUserRequest     string `xml:",omitempty"` // xs:dateTime
	DateAvailable         string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount *AcknowledgedFeeAmount
	PaidFeeAmount         *PaidFeeAmount
	Ext                   *Ext
}

type AddUserFields struct {
	AuthenticationInput    []AuthenticationInput
	NameInformation        *NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []UserLanguage
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type AgencyAddressInformation struct {
	AgencyAddressRoleType AgencyAddressRoleType
	ValidFromDate         string `xml:",omitempty"` // xs:dateTime
	ValidToDate           string `xml:",omitempty"` // xs:dateTime
	// xs:choice ->
	PhysicalAddress   *PhysicalAddress
	ElectronicAddress *ElectronicAddress
	// <- xs:choice
	Ext *Ext
}

type Amount struct {
	CurrencyCode  CurrencyCode
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type AuthenticationInput struct {
	AuthenticationInputData      string
	AuthenticationDataFormatType AuthenticationDataFormatType
	AuthenticationInputType      AuthenticationInputType
	Ext                          *Ext
}

type AuthenticationPrompt struct {
	PromptOutput PromptOutput
	PromptInput  PromptInput
	Ext          *Ext
}

type BibliographicDescription struct {
	Author                     string `xml:",omitempty"`
	AuthorOfComponent          string `xml:",omitempty"`
	BibliographicItemId        *BibliographicItemId
	BibliographicRecordId      []BibliographicRecordId
	ComponentId                *ComponentId
	Edition                    string `xml:",omitempty"`
	Pagination                 string `xml:",omitempty"`
	PlaceOfPublication         string `xml:",omitempty"`
	PublicationDate            string `xml:",omitempty"`
	PublicationDateOfComponent string `xml:",omitempty"`
	Publisher                  string `xml:",omitempty"`
	SeriesTitleNumber          string `xml:",omitempty"`
	Title                      string `xml:",omitempty"`
	TitleOfComponent           string `xml:",omitempty"`
	BibliographicLevel         *BibliographicLevel
	SponsoringBody             string `xml:",omitempty"`
	ElectronicDataFormatType   *ElectronicDataFormatType
	Language                   *Language
	MediumType                 *MediumType
	Ext                        *Ext
}

type BibliographicId struct {
	// xs:choice ->
	BibliographicItemId   *BibliographicItemId
	BibliographicRecordId *BibliographicRecordId
	// <- xs:choice
	Ext *Ext
}

type BibliographicItemId struct {
	BibliographicItemIdentifier     string
	BibliographicItemIdentifierCode *BibliographicItemIdentifierCode
	Ext                             *Ext
}

type BibliographicRecordId struct {
	BibliographicRecordIdentifier string
	// xs:choice ->
	AgencyId                          *SchemeValue
	BibliographicRecordIdentifierCode *BibliographicRecordIdentifierCode
	// <- xs:choice
	Ext *Ext
}

type BlockOrTrap struct {
	AgencyId        SchemeValue
	BlockOrTrapType BlockOrTrapType
	ValidFromDate   string `xml:",omitempty"` // xs:dateTime
	ValidToDate     string `xml:",omitempty"` // xs:dateTime
	Ext             *Ext
}

type ChronologyLevelInstance struct {
	// xs:choice ->
	//ChronologyCaption string `xml:",omitempty"`
	// xs:sequence ->
	ChronologyLevel   int    // xs:positiveInteger
	ChronologyCaption string `xml:",omitempty"`
	// <- xs:sequence
	// <- xs:choice
	ChronologyValue string
	Ext             *Ext
}

type ComponentId struct {
	ComponentIdentifierType ComponentIdentifierType
	ComponentIdentifier     string
	Ext                     *Ext
}

type CurrentBorrower struct {
	UserId UserId
	Ext    *Ext
}

type CurrentBorrowerDesired struct {
}

type CurrentRequester struct {
	UserId UserId
	Ext    *Ext
}

type CurrentRequestersDesired struct {
}

type DeleteAgencyFields struct {
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []ApplicationProfileSupportedType
	ConsortiumAgreement             []ConsortiumAgreement
	AgencyUserPrivilegeType         []AgencyUserPrivilegeType
	Ext                             *Ext
}

type DeleteItemFields struct {
	BibliographicDescription *BibliographicDescription
	ItemUseRestrictionType   []ItemUseRestrictionType
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SecurityMarker
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type DeleteRequestFields struct {
	UserId                *UserId
	ItemId                *ItemId
	RequestType           *RequestType
	RequestScopeType      *RequestScopeType
	RequestStatusType     *RequestStatusType
	ShippingInformation   *ShippingInformation
	EarliestDateNeeded    string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate        string `xml:",omitempty"` // xs:dateTime
	PickupLocation        *PickupLocation
	PickupExpiryDate      string `xml:",omitempty"` // xs:dateTime
	DateOfUserRequest     string `xml:",omitempty"` // xs:dateTime
	DateAvailable         string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount *AcknowledgedFeeAmount
	PaidFeeAmount         *PaidFeeAmount
	Ext                   *Ext
}

type DeleteUserFields struct {
	AuthenticationInput    []AuthenticationInput
	NameInformation        *NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []UserLanguage
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type Destination struct {
	// xs:choice ->
	//Location *Location
	// xs:sequence ->
	BinNumber string
	Location  *Location
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type ElectronicAddress struct {
	ElectronicAddressType ElectronicAddressType
	ElectronicAddressData string
	Ext                   *Ext
}

type ElectronicResource struct {
	// xs:choice ->
	ElectronicDataFormatType *ElectronicDataFormatType
	ActualResource           string `xml:",omitempty"`
	// <- xs:sequence
	ReferenceToResource string
	Ext                 *Ext
}

type ElectronicResourceProvidedFlag struct {
}

type EnumerationLevelInstance struct {
	// xs:choice ->
	//EnumerationCaption string `xml:",omitempty"`
	// xs:sequence ->
	EnumerationLevel   int    // xs:positiveInteger
	EnumerationCaption string `xml:",omitempty"`
	Ext                *Ext
	// <- xs:sequence
	// <- xs:choice
	EnumerationValue string
	//Ext              *Ext
}

type Ext struct {
}

type FiscalTransactionInformation struct {
	FiscalActionType                    FiscalActionType
	FiscalTransactionReferenceId        *FiscalTransactionReferenceId
	RelatedFiscalTransactionReferenceId []RelatedFiscalTransactionReferenceId
	FiscalTransactionType               FiscalTransactionType
	ValidFromDate                       string `xml:",omitempty"` // xs:dateTime
	ValidToDate                         string `xml:",omitempty"` // xs:dateTime
	Amount                              Amount
	PaymentMethodType                   *PaymentMethodType
	FiscalTransactionDescription        string `xml:",omitempty"`
	RequestId                           *RequestId
	ItemDetails                         *ItemDetails
	Ext                                 *Ext
}

type FiscalTransactionReferenceId struct {
	AgencyId                         SchemeValue
	FiscalTransactionIdentifierValue string
	Ext                              *Ext
}

type FromAgencyId struct {
	AgencyId SchemeValue
	Ext      *Ext
}

type HoldingsChronology struct {
	ChronologyLevelInstance []ChronologyLevelInstance
	Ext                     *Ext
}

type HoldingsEnumeration struct {
	EnumerationLevelInstance []EnumerationLevelInstance
	Ext                      *Ext
}

type HoldingsInformation struct {
	// xs:choice ->
	StructuredHoldingsData   []StructuredHoldingsData
	UnstructuredHoldingsData string `xml:",omitempty"`
	// <- xs:choice
	Ext *Ext
}

type IndeterminateLoanPeriodFlag struct {
}

type ItemDescription struct {
	CallNumber           string `xml:",omitempty"`
	CopyNumber           string `xml:",omitempty"`
	ItemDescriptionLevel *ItemDescriptionLevel
	HoldingsInformation  *HoldingsInformation
	NumberOfPieces       *int // xs:positiveInteger
	Ext                  *Ext
}

type ItemDetails struct {
	ItemId                   ItemId
	BibliographicDescription BibliographicDescription
	DateCheckedOut           string   `xml:",omitempty"` // xs:dateTime
	DateRenewed              []string `xml:",omitempty"` // xs:dateTime
	// xs:choice ->
	DateDue                     string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	NonReturnableFlag           *NonReturnableFlag
	// <- xs:choice
	DateReturned string `xml:",omitempty"` // xs:dateTime
	Ext          *Ext
}

type ItemId struct {
	AgencyId            *SchemeValue
	ItemIdentifierType  *ItemIdentifierType
	ItemIdentifierValue string
	Ext                 *Ext
}

type ItemOptionalFields struct {
	BibliographicDescription *BibliographicDescription
	ItemUseRestrictionType   []ItemUseRestrictionType
	CirculationStatus        *CirculationStatus
	HoldQueueLength          *int // xs:nonnegativeInteger
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	ElectronicResource       *ElectronicResource
	SecurityMarker           *SecurityMarker
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type ItemReportedLost struct {
}

type ItemReportedMissing struct {
}

type ItemReportedNeverBorrowed struct {
}

type ItemReportedReturned struct {
	DateReportedReturned string // xs:dateTime
	Ext                  *Ext
}

type ItemTransaction struct {
	CurrentBorrower  *CurrentBorrower
	CurrentRequester []CurrentRequester
	Ext              *Ext
}

type LoanedItem struct {
	ItemId        ItemId
	ReminderLevel int // xs:positiveInteger
	// xs:choice ->
	DateDue                     string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	// <- xs:choice
	Amount     Amount
	Title      string `xml:",omitempty"`
	MediumType *MediumType
	Ext        *Ext
}

type LoanedItemsCount struct {
	// xs:choice ->
	CirculationStatus      *CirculationStatus
	ItemUseRestrictionType *ItemUseRestrictionType
	// <- xs:choice
	LoanedItemCountValue int // xs:nonnegativeInteger
	Ext                  *Ext
}

type LoanedItemsDesired []byte // TODO

type Location struct {
	LocationType  LocationType
	LocationName  LocationName
	ValidFromDate string `xml:",omitempty"` // xs:dateTime
	ValidToDate   string `xml:",omitempty"` // xs:dateTime
	Ext           *Ext
}

type LocationName struct {
	LocationNameInstance []LocationNameInstance
	Ext                  *Ext
}

type LocationNameInstance struct {
	LocationNameLevel int // xs:positiveInteger
	LocationNameValue string
	Ext               *Ext
}

type MandatedAction struct {
	DateEventOccurred string // xs:dateTime
	Ext               *Ext
}

type NameInformation struct {
	// xs:choice ->
	PersonalNameInformation     *PersonalNameInformation
	OrganizationNameInformation []OrganizationNameInformation
	// <- xs:choice
	Ext *Ext
}

type NonReturnableFlag struct {
}

type NoticeItem struct {
	ItemDetails ItemDetails
	Amount      *Amount
	Ext         *Ext
}

type OnBehalfOfAgency struct {
	AgencyId SchemeValue
	Ext      *Ext
}

type OrganizationNameInformation struct {
	OrganizationNameType OrganizationNameType
	OrganizationName     string
	Ext                  *Ext
}

type PaidFeeAmount struct {
	CurrencyCode  CurrencyCode
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type Pending struct {
	DateOfExpectedReply string `xml:",omitempty"` // xs:dateTime
	Ext                 *Ext
}

type PersonalNameInformation struct {
	// xs:choice ->
	//UnstructuredPersonalUserName string `xml:",omitempty"`
	// xs:sequence ->
	StructuredPersonalUserName   StructuredPersonalUserName
	UnstructuredPersonalUserName string `xml:",omitempty"`
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type PhysicalAddress struct {
	// xs:choice ->
	StructuredAddress   *StructuredAddress
	UnstructuredAddress *UnstructuredAddress
	// <- xs:choice
	PhysicalAddressType PhysicalAddressType
	Ext                 *Ext
}

type PhysicalCondition struct {
	PhysicalConditionType    PhysicalConditionType
	PhysicalConditionDetails string `xml:",omitempty"`
	Ext                      *Ext
}

type PreviousUserId struct {
	AgencyId            SchemeValue
	UserIdentifierValue string
	ValidFromDate       string `xml:",omitempty"` // xs:dateTime
	ValidToDate         string `xml:",omitempty"` // xs:dateTime
	Ext                 *Ext
}

type Problem struct {
	ProblemType    ProblemType
	ProblemDetail  string `xml:",omitempty"`
	ProblemElement string `xml:",omitempty"`
	ProblemValue   string `xml:",omitempty"`
	Ext            *Ext
}

type PromptInput struct {
	AuthenticationInputType      AuthenticationInputType
	AuthenticationDataFormatType AuthenticationDataFormatType
	SensitiveDataFlag            *SensitiveDataFlag
	Ext                          *Ext
}

type PromptOutput struct {
	AuthenticationPromptData string
	AuthenticationPromptType AuthenticationPromptType
	Ext                      *Ext
}

type RelatedFiscalTransactionReferenceId struct {
	AgencyId                         SchemeValue
	FiscalTransactionIdentifierValue string
	Ext                              *Ext
}

type RenewalNotPermitted struct {
}

type RequestedItemsDesired []byte // TODO

type RequestId struct {
	AgencyId               *SchemeValue
	RequestIdentifierType  *RequestIdentifierType
	RequestIdentifierValue string
	Ext                    *Ext
}

type RequestedItem struct {
	// xs:choice ->
	//ItemId *ItemId
	// xs:sequence ->
	RequestId RequestId
	ItemId    *ItemId
	// <- xs:sequence
	// <- xs:choice
	RequestType       RequestType
	RequestStatusType RequestStatusType
	DatePlaced        string // xs:dateTime
	PickupDate        string `xml:",omitempty"` // xs:dateTime
	PickupLocation    *PickupLocation
	PickupExpiryDate  string `xml:",omitempty"` // xs:dateTime
	ReminderLevel     *int   // xs:positiveInteger
	HoldQueuePosition *int   // xs:positiveInteger
	Title             string `xml:",omitempty"`
	MediumType        *MediumType
	Ext               *Ext
}

type RequestedItemsCount struct {
	// xs:choice ->
	RequestType            *RequestType
	CirculationStatus      *CirculationStatus
	ItemUseRestrictionType *ItemUseRestrictionType
	// <- xs:choice
	RequestedItemCountValue int // xs:nonnegativeInteger
	Ext                     *Ext
}

type RequiredFeeAmount struct {
	CurrencyCode  CurrencyCode
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type ResourceDesired struct {
}

type RoutingInformation struct {
	RoutingInstructions string
	Destination         Destination
	RequestType         *RequestType
	UserId              *UserId
	NameInformation     *NameInformation
	Ext                 *Ext
}

type SensitiveDataFlag struct {
}

type SensitizationFlag struct {
}

type ShippingInformation struct {
	ShippingInstructions string `xml:",omitempty"`
	ShippingNote         string `xml:",omitempty"`
	// xs:choice ->
	PhysicalAddress   *PhysicalAddress
	ElectronicAddress *ElectronicAddress
	// <- xs:choice
	Ext *Ext
}

type StructuredAddress struct {
	// xs:choice ->
	LocationWithinBuilding string `xml:",omitempty"`
	HouseName              string `xml:",omitempty"`
	// xs:choice ->
	//District string `xml:",omitempty"`
	// xs:sequence ->
	//PostOfficeBox string
	//District      string `xml:",omitempty"`
	// <- xs:sequence
	// xs:sequence ->
	Street        string
	PostOfficeBox string `xml:",omitempty"`
	District      string `xml:",omitempty"`
	// <- xs:sequence
	// <- xs:choice
	// <- xs:sequence
	// xs:sequence ->
	Line1 string
	Line2 string `xml:",omitempty"`
	// <- xs:sequence
	Locality   string `xml:",omitempty"`
	Region     string `xml:",omitempty"`
	Country    string `xml:",omitempty"`
	PostalCode string `xml:",omitempty"`
	CareOf     string `xml:",omitempty"`
	Ext        *Ext
}

type StructuredHoldingsData struct {
	// xs:choice ->
	//HoldingsChronology *HoldingsChronology
	// xs:sequence ->
	HoldingsEnumeration HoldingsEnumeration
	HoldingsChronology  *HoldingsChronology
	// <- xs:sequence
	// <- xs:choice
	Ext *Ext
}

type StructuredPersonalUserName struct {
	Prefix    string `xml:",omitempty"`
	GivenName string `xml:",omitempty"`
	Initials  string `xml:",omitempty"`
	Surname   string
	Suffix    string `xml:",omitempty"`
	Ext       *Ext
}

type ToAgencyId struct {
	AgencyId SchemeValue
	Ext      *Ext
}

type UnstructuredAddress struct {
	UnstructuredAddressType UnstructuredAddressType
	UnstructuredAddressData string
	Ext                     *Ext
}

type UserAddressInformation struct {
	UserAddressRoleType UserAddressRoleType
	ValidFromDate       string `xml:",omitempty"` // xs:dateTime
	ValidToDate         string `xml:",omitempty"` // xs:dateTime
	// xs:choice ->
	PhysicalAddress   *PhysicalAddress
	ElectronicAddress *ElectronicAddress
	// <- xs:choice
	Ext *Ext
}

type UserElementEnum struct {
	// xs:sequence ->
	AccountBalance AccountBalance
	AccountDetails []AccountDetails
	Ext            *Ext
	// <- xs:sequence
}

type UserFiscalAccountDesired []byte

type UserId struct {
	AgencyId            *SchemeValue
	UserIdentifierType  *UserIdentifierType
	UserIdentifierValue string
	Ext                 *Ext
}

type UserNoticeDetails struct {
	NoticeType    NoticeType
	NoticeContent string `xml:",omitempty"`
	// xs:choice ->
	NoticeItem        []NoticeItem
	UserFiscalAccount *UserFiscalAccount
	UserPrivilege     *UserPrivilege
	// <- xs:choice
	Ext *Ext
}

type UserOptionalFields struct {
	NameInformation        *NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []UserLanguage
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	UserId                 []UserId
	PreviousUserId         []PreviousUserId
	Ext                    *Ext
}

type UserPrivilege struct {
	AgencyId                 SchemeValue
	AgencyUserPrivilegeType  AgencyUserPrivilegeType
	ValidFromDate            string `xml:",omitempty"` // xs:dateTime
	ValidToDate              string `xml:",omitempty"` // xs:dateTime
	UserPrivilegeFee         *UserPrivilegeFee
	UserPrivilegeStatus      *UserPrivilegeStatus
	UserPrivilegeDescription string `xml:",omitempty"`
	Ext                      *Ext
}

type UserPrivilegeFee struct {
	Amount            Amount
	PaymentMethodType *PaymentMethodType
	Ext               *Ext
}

type UserPrivilegeStatus struct {
	UserPrivilegeStatusType   UserPrivilegeStatusType
	DateOfUserPrivilegeStatus string `xml:",omitempty"` // xs:dateTime
	Ext                       *Ext
}

// TODO complete theese, missing from autogenerated output
type UserFiscalAccount struct{}

// TODO Interface implementations
func (r LookupUser) Type() requestType   { return TypeLookupUser }
func (r RequestItem) Type() requestType  { return TypeRequestItem }
func (r CheckOutItem) Type() requestType { return TypeCheckOutItem }
func (r CheckInItem) Type() requestType  { return TypeCheckInItem }
