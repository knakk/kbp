package ncip

import "encoding/xml"

type SchemeValue struct {
	Scheme string `xml:",attr,omitempty"`
	Value  string `xml:",innerxml"`
}

type AcceptItem struct {
	XMLName                      xml.Name
	InitiationHeader             *InitiationHeader
	MandatedAction               *MandatedAction
	RequestId                    RequestId
	RequestedActionType          SchemeValue
	UserId                       *UserId
	ItemId                       *ItemId
	DateForReturn                string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag  *IndeterminateLoanPeriodFlag
	NonReturnableFlag            *NonReturnableFlag
	RenewalNotPermitted          *RenewalNotPermitted
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	PickupLocation               *SchemeValue
	PickupExpiryDate             string `xml:",omitempty"` // xs:dateTime
	Ext                          *Ext
}

type AcceptItemResponse struct {
	XMLName        xml.Name
	ResponseHeader *ResponseHeader
	Problem        []Problem
	RequestId      *RequestId
	ItemId         *ItemId
	Ext            *Ext
}

type AgencyCreated struct {
	XMLName                         xml.Name
	InitiationHeader                *InitiationHeader
	AgencyId                        SchemeValue
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []SchemeValue
	ConsortiumAgreement             []SchemeValue
	AgencyUserPrivilegeType         []SchemeValue
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
	ItemElementType  []SchemeValue
	UserElementType  []SchemeValue
	Ext              *Ext
}

type CancelRecallItemResponse struct {
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	ItemId                       ItemId
	UserId                       *UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type CancelRequestItem struct {
	XMLName               xml.Name
	InitiationHeader      *InitiationHeader
	MandatedAction        *MandatedAction
	UserId                *UserId
	AuthenticationInput   []AuthenticationInput
	RequestId             *RequestId
	ItemId                *ItemId
	RequestType           SchemeValue
	RequestScopeType      *SchemeValue
	AcknowledgedFeeAmount *AcknowledgedFeeAmount
	PaidFeeAmount         *PaidFeeAmount
	ItemElementType       []SchemeValue
	UserElementType       []SchemeValue
	Ext                   *Ext
}

type CancelRequestItemResponse struct {
	XMLName                      xml.Name
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	RequestId                    *RequestId
	ItemId                       *ItemId
	UserId                       *UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type CheckInItem struct {
	XMLName          xml.Name
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	ItemElementType  []SchemeValue
	UserElementType  []SchemeValue
	Ext              *Ext
}

type CheckInItemResponse struct {
	XMLName                      xml.Name
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	ItemId                       *ItemId
	UserId                       *UserId
	RoutingInformation           *RoutingInformation
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type CheckOutItem struct {
	XMLName                            xml.Name
	InitiationHeader                   *InitiationHeader
	MandatedAction                     *MandatedAction
	UserId                             *UserId
	AuthenticationInput                []AuthenticationInput
	ItemId                             ItemId
	RequestId                          *RequestId
	AcknowledgedFeeAmount              *AcknowledgedFeeAmount
	PaidFeeAmount                      *PaidFeeAmount
	AcknowledgedItemUseRestrictionType []SchemeValue
	ShippingInformation                *ShippingInformation
	ResourceDesired                    *ResourceDesired
	DesiredDateDue                     string `xml:",omitempty"` // xs:dateTime
	ItemElementType                    []SchemeValue
	UserElementType                    []SchemeValue
	Ext                                *Ext
}

type CheckOutItemResponse struct {
	XMLName                        xml.Name
	ResponseHeader                 *ResponseHeader
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []SchemeValue
	ItemId                         *ItemId
	UserId                         *UserId
	DateDue                        string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag    *IndeterminateLoanPeriodFlag
	NonReturnableFlag              *NonReturnableFlag
	RenewalCount                   *int // xs:nonnegativeInteger
	ElectronicResource             *ElectronicResource
	FiscalTransactionInformation   *FiscalTransactionInformation
	ItemOptionalFields             *ItemOptionalFields
	UserOptionalFields             *UserOptionalFields
	Ext                            *Ext
}

type CirculationStatusChangeReported struct {
	InitiationHeader             *InitiationHeader
	ItemId                       ItemId
	UserId                       UserId
	ItemReportedReturned         *ItemReportedReturned
	ItemReportedNeverBorrowed    *ItemReportedNeverBorrowed
	ItemReportedLost             *ItemReportedLost
	ItemReportedMissing          *ItemReportedMissing
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
	CirculationStatus SchemeValue
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
	ApplicationProfileSupportedType []SchemeValue
	ConsortiumAgreement             []SchemeValue
	AgencyUserPrivilegeType         []SchemeValue
	Ext                             *Ext
}

type CreateAgencyResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	AgencyId       *SchemeValue
	Ext            *Ext
}

type CreateItem struct {
	InitiationHeader         *InitiationHeader
	MandatedAction           *MandatedAction
	ItemId                   *ItemId
	RequestId                *RequestId
	BibliographicDescription BibliographicDescription
	ItemUseRestrictionType   []SchemeValue
	CirculationStatus        *SchemeValue
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SchemeValue
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type CreateItemResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

type CreateUser struct {
	InitiationHeader       *InitiationHeader
	MandatedAction         *MandatedAction
	UserId                 *UserId
	AuthenticationInput    []AuthenticationInput
	NameInformation        NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []SchemeValue
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type CreateUserResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         *UserId
	Ext            *Ext
}

type CreateUserFiscalTransaction struct {
	InitiationHeader             *InitiationHeader
	MandatedAction               *MandatedAction
	UserId                       *UserId
	AuthenticationInput          []AuthenticationInput
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

type CreateUserFiscalTransactionResponse struct {
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	UserId                       UserId
	FiscalTransactionReferenceId FiscalTransactionReferenceId
	Ext                          *Ext
}

type DeleteItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	Ext              *Ext
}

type DeleteItemResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

type DeleteUser struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	Ext                 *Ext
}

type DeleteUserResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         *UserId
	Ext            *Ext
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
	InitiationHeader               *InitiationHeader
	UserId                         UserId
	ItemId                         ItemId
	RequestId                      *RequestId
	DateDue                        string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag    *IndeterminateLoanPeriodFlag
	NonReturnableFlag              *NonReturnableFlag
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
	ItemUseRestrictionType   []SchemeValue
	CirculationStatus        *SchemeValue
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SchemeValue
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
	InitiationHeader             *InitiationHeader
	UserId                       UserId
	RequestId                    RequestId
	ItemId                       *ItemId
	RequestType                  SchemeValue
	RequestScopeType             *SchemeValue
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
	InitiationHeader    *InitiationHeader
	UserId              UserId
	ItemId              ItemId
	RequestType         SchemeValue
	RequestId           *RequestId
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
	InitiationHeader    *InitiationHeader
	UserId              UserId
	ItemId              *ItemId
	BibliographicId     BibliographicId
	RequestId           RequestId
	RequestType         SchemeValue
	RequestScopeType    SchemeValue
	ShippingInformation *ShippingInformation
	EarliestDateNeeded  string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate      string `xml:",omitempty"` // xs:dateTime
	PickupLocation      *SchemeValue
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
	InitiationHeader    *InitiationHeader
	RequestId           RequestId
	ItemId              *ItemId
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
	AgencyElementType []SchemeValue
	Ext               *Ext
}

type LookupAgencyResponse struct {
	ResponseHeader                  *ResponseHeader
	Problem                         []Problem
	AgencyId                        SchemeValue
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []SchemeValue
	ConsortiumAgreement             []SchemeValue
	AgencyUserPrivilegeType         []SchemeValue
	Ext                             *Ext
}

type LookupItem struct {
	InitiationHeader         *InitiationHeader
	ItemId                   *ItemId
	RequestId                *RequestId
	ItemElementType          []SchemeValue
	CurrentBorrowerDesired   *CurrentBorrowerDesired
	CurrentRequestersDesired *CurrentRequestersDesired
	Ext                      *Ext
}

type LookupItemResponse struct {
	ResponseHeader     *ResponseHeader
	Problem            []Problem
	RequestId          RequestId
	ItemId             *ItemId
	HoldPickupDate     string `xml:",omitempty"` // xs:dateTime
	DateRecalled       string `xml:",omitempty"` // xs:dateTime
	ItemTransaction    *ItemTransaction
	ItemOptionalFields *ItemOptionalFields
	Ext                *Ext
}

type LookupRequest struct {
	InitiationHeader    *InitiationHeader
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	ItemId              ItemId
	RequestType         SchemeValue
	RequestId           *RequestId
	RequestElementType  []SchemeValue
	ItemElementType     []SchemeValue
	UserElementType     []SchemeValue
	Ext                 *Ext
}

type LookupRequestResponse struct {
	ResponseHeader        *ResponseHeader
	Problem               []Problem
	RequestId             RequestId
	ItemId                *ItemId
	UserId                *UserId
	RequestType           *SchemeValue
	RequestScopeType      *SchemeValue
	RequestStatusType     *SchemeValue
	HoldQueuePosition     *int // xs:positiveInteger
	ShippingInformation   *ShippingInformation
	EarliestDateNeeded    string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate        string `xml:",omitempty"` // xs:dateTime
	PickupDate            string `xml:",omitempty"` // xs:dateTime
	PickupLocation        *SchemeValue
	PickupExpiryDate      string `xml:",omitempty"` // xs:dateTime
	DateOfUserRequest     string `xml:",omitempty"` // xs:dateTime
	DateAvailable         string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount *AcknowledgedFeeAmount
	PaidFeeAmount         *PaidFeeAmount
	ItemOptionalFields    *ItemOptionalFields
	UserOptionalFields    *UserOptionalFields
	Ext                   *Ext
}

type LookupUser struct {
	XMLName                  xml.Name
	InitiationHeader         *InitiationHeader
	UserId                   *UserId
	AuthenticationInput      []AuthenticationInput
	UserElementType          []SchemeValue
	LoanedItemsDesired       *LoanedItemsDesired       `xml:",omitempty"`
	RequestedItemsDesired    *RequestedItemsDesired    `xml:",omitempty"`
	UserFiscalAccountDesired *UserFiscalAccountDesired `xml:",omitempty"`
	Ext                      *Ext
}

type LookupUserResponse struct {
	XMLName             xml.Name
	ResponseHeader      *ResponseHeader
	Problem             []Problem
	UserId              *UserId
	UserFiscalAccount   []UserFiscalAccount
	LoanedItemsCount    []LoanedItemsCount
	LoanedItem          []LoanedItem
	RequestedItemsCount []RequestedItemsCount
	RequestedItem       []RequestedItem
	UserOptionalFields  *UserOptionalFields
	Ext                 *Ext
}

type RecallItem struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	ItemId              ItemId
	DesiredDateDue      string `xml:",omitempty"` // xs:dateTime
	ShippingInformation *ShippingInformation
	ItemElementType     []SchemeValue
	UserElementType     []SchemeValue
	Ext                 *Ext
}

type RecallItemResponse struct {
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	ItemId                       ItemId
	UserId                       *UserId
	DateDue                      string `xml:",omitempty"` // xs:dateTime
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

type RenewItem struct {
	XMLName                            xml.Name
	InitiationHeader                   *InitiationHeader
	MandatedAction                     *MandatedAction
	UserId                             *UserId
	AuthenticationInput                []AuthenticationInput
	ItemId                             ItemId
	ItemElementType                    []SchemeValue
	UserElementType                    []SchemeValue
	DesiredDateDue                     string `xml:",omitempty"` // xs:dateTime
	DesiredDateForReturn               string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount              *AcknowledgedFeeAmount
	PaidFeeAmount                      *PaidFeeAmount
	AcknowledgedItemUseRestrictionType []SchemeValue
	Ext                                *Ext
}

type RenewItemResponse struct {
	XMLName                        xml.Name
	ResponseHeader                 *ResponseHeader
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []SchemeValue
	Pending                        *Pending
	ItemId                         *ItemId
	UserId                         *UserId
	DateDue                        string `xml:",omitempty"` // xs:dateTime
	DateForReturn                  string `xml:",omitempty"` // xs:dateTime
	RenewalCount                   *int   // xs:nonnegativeInteger
	FiscalTransactionInformation   *FiscalTransactionInformation
	ItemOptionalFields             *ItemOptionalFields
	UserOptionalFields             *UserOptionalFields
	Ext                            *Ext
}

type ReportCirculationStatusChange struct {
	InitiationHeader          *InitiationHeader
	MandatedAction            *MandatedAction
	ItemId                    ItemId
	UserId                    *UserId
	AuthenticationInput       []AuthenticationInput
	ItemReportedReturned      *ItemReportedReturned
	ItemReportedNeverBorrowed *ItemReportedNeverBorrowed
	ItemReportedLost          *ItemReportedLost
	ItemReportedMissing       *ItemReportedMissing
	ItemElementType           []SchemeValue
	UserElementType           []SchemeValue
	Ext                       *Ext
}

type ReportCirculationStatusChangeResponse struct {
	ResponseHeader     *ResponseHeader
	Problem            []Problem
	ItemId             ItemId
	UserId             *UserId
	ItemOptionalFields *ItemOptionalFields
	UserOptionalFields *UserOptionalFields
	Ext                *Ext
}

type RequestItem struct {
	XMLName                            xml.Name
	InitiationHeader                   *InitiationHeader
	MandatedAction                     *MandatedAction
	UserId                             *UserId
	AuthenticationInput                []AuthenticationInput
	BibliographicId                    []BibliographicId
	ItemId                             []ItemId
	RequestId                          *RequestId
	RequestType                        SchemeValue
	RequestScopeType                   SchemeValue
	ItemOptionalFields                 *ItemOptionalFields
	ShippingInformation                *ShippingInformation
	EarliestDateNeeded                 string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate                     string `xml:",omitempty"` // xs:dateTime
	PickupLocation                     *SchemeValue
	PickupExpiryDate                   string `xml:",omitempty"` // xs:dateTime
	AcknowledgedFeeAmount              *AcknowledgedFeeAmount
	PaidFeeAmount                      *PaidFeeAmount
	AcknowledgedItemUseRestrictionType []SchemeValue
	ItemElementType                    []SchemeValue
	UserElementType                    []SchemeValue
	Ext                                *Ext
}

type RequestItemResponse struct {
	XMLName                        xml.Name
	ResponseHeader                 *ResponseHeader
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []SchemeValue
	RequestId                      *RequestId
	ItemId                         *ItemId
	UserId                         *UserId
	RequestType                    *SchemeValue
	RequestScopeType               *SchemeValue
	ShippingInformation            *ShippingInformation
	DateAvailable                  string `xml:",omitempty"` // xs:dateTime
	HoldPickupDate                 string `xml:",omitempty"` // xs:dateTime
	FiscalTransactionInformation   *FiscalTransactionInformation
	ItemOptionalFields             *ItemOptionalFields
	UserOptionalFields             *UserOptionalFields
	Ext                            *Ext
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
	Problem        []Problem
	UserId         UserId
	DateSent       string `xml:",omitempty"` // xs:dateTime
	DateWillSend   string `xml:",omitempty"` // xs:dateTime
	Ext            *Ext
}

type UndoCheckOutItem struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	ItemId              ItemId
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	RequestId           *RequestId
	Ext                 *Ext
}

type UndoCheckOutItemResponse struct {
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	ItemId                       ItemId
	UserId                       *UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	Ext                          *Ext
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
	Problem        []Problem
	AgencyId       *SchemeValue
	Ext            *Ext
}

type UpdateCirculationStatus struct {
	InitiationHeader  *InitiationHeader
	MandatedAction    *MandatedAction
	ItemId            ItemId
	CirculationStatus SchemeValue
	Ext               *Ext
}

type UpdateCirculationStatusResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
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
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

type UpdateRequestItem struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	ItemId              ItemId
	RequestType         SchemeValue
	RequestId           *RequestId
	DeleteRequestFields *DeleteRequestFields
	AddRequestFields    *AddRequestFields
	ItemElementType     []SchemeValue
	UserElementType     []SchemeValue
	Ext                 *Ext
}

type UpdateRequestItemResponse struct {
	ResponseHeader                 *ResponseHeader
	Problem                        []Problem
	RequiredFeeAmount              *RequiredFeeAmount
	RequiredItemUseRestrictionType []SchemeValue
	ItemId                         *ItemId
	UserId                         *UserId
	RequestType                    *SchemeValue
	RequestScopeType               *SchemeValue
	DateAvailable                  string `xml:",omitempty"` // xs:dateTime
	HoldPickupDate                 string `xml:",omitempty"` // xs:dateTime
	FiscalTransactionInformation   *FiscalTransactionInformation
	ItemOptionalFields             *ItemOptionalFields
	UserOptionalFields             *UserOptionalFields
	Ext                            *Ext
}

type UpdateUser struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	DeleteUserFields    *DeleteUserFields
	AddUserFields       *AddUserFields
	Ext                 *Ext
}

type UpdateUserResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         *UserId
	Ext            *Ext
}

type UserCreated struct {
	InitiationHeader       *InitiationHeader
	UserId                 UserId
	NameInformation        NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []SchemeValue
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

// End of request/response structs. The following types are embedded in above structs:

type InitiationHeader struct {
	FromSystemId             *SchemeValue
	FromSystemAuthentication string `xml:",omitempty"`
	FromAgencyId             FromAgencyId
	FromAgencyAuthentication string `xml:",omitempty"`
	OnBehalfOfAgency         *OnBehalfOfAgency
	ToSystemId               *SchemeValue
	ToAgencyId               ToAgencyId
	ApplicationProfileType   *SchemeValue
	Ext                      *Ext
}

type ResponseHeader struct {
	FromSystemId             *SchemeValue
	FromSystemAuthentication string `xml:",omitempty"`
	FromAgencyId             FromAgencyId
	FromAgencyAuthentication string `xml:",omitempty"`
	ToSystemId               *SchemeValue
	ToAgencyId               ToAgencyId
	Ext                      *Ext
}

type AccountBalance struct {
	CurrencyCode  SchemeValue
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type AccountDetails struct {
	AccrualDate                  string // xs:dateTime
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

type AcknowledgedFeeAmount struct {
	CurrencyCode  SchemeValue
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type AddAgencyFields struct {
	OrganizationNameInformation     []OrganizationNameInformation
	AgencyAddressInformation        []AgencyAddressInformation
	AuthenticationPrompt            []AuthenticationPrompt
	ApplicationProfileSupportedType []SchemeValue
	ConsortiumAgreement             []SchemeValue
	AgencyUserPrivilegeType         []SchemeValue
	Ext                             *Ext
}

type AddItemFields struct {
	BibliographicDescription *BibliographicDescription
	ItemUseRestrictionType   []SchemeValue
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SchemeValue
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type AddRequestFields struct {
	UserId                *UserId
	ItemId                *ItemId
	RequestType           *SchemeValue
	RequestScopeType      *SchemeValue
	RequestStatusType     *SchemeValue
	ShippingInformation   *ShippingInformation
	EarliestDateNeeded    string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate        string `xml:",omitempty"` // xs:dateTime
	PickupLocation        *SchemeValue
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
	UserLanguage           []SchemeValue
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type AgencyAddressInformation struct {
	AgencyAddressRoleType SchemeValue
	ValidFromDate         string `xml:",omitempty"` // xs:dateTime
	ValidToDate           string `xml:",omitempty"` // xs:dateTime
	PhysicalAddress       *PhysicalAddress
	ElectronicAddress     *ElectronicAddress
	Ext                   *Ext
}

type Amount struct {
	CurrencyCode  SchemeValue
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type AuthenticationInput struct {
	AuthenticationInputData      string
	AuthenticationDataFormatType SchemeValue
	AuthenticationInputType      SchemeValue
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
	BibliographicLevel         *SchemeValue
	SponsoringBody             string `xml:",omitempty"`
	ElectronicDataFormatType   *SchemeValue
	Language                   *SchemeValue
	MediumType                 *SchemeValue
	Ext                        *Ext
}

type BibliographicId struct {
	BibliographicItemId   *BibliographicItemId
	BibliographicRecordId *BibliographicRecordId
	Ext                   *Ext
}

type BibliographicItemId struct {
	BibliographicItemIdentifier     string
	BibliographicItemIdentifierCode *SchemeValue
	Ext                             *Ext
}

type BibliographicRecordId struct {
	BibliographicRecordIdentifier     string
	AgencyId                          *SchemeValue
	BibliographicRecordIdentifierCode *SchemeValue
	Ext                               *Ext
}

type BlockOrTrap struct {
	AgencyId        SchemeValue
	BlockOrTrapType SchemeValue
	ValidFromDate   string `xml:",omitempty"` // xs:dateTime
	ValidToDate     string `xml:",omitempty"` // xs:dateTime
	Ext             *Ext
}

type ChronologyLevelInstance struct {
	//ChronologyCaption string `xml:",omitempty"`
	ChronologyLevel   int    // xs:positiveInteger
	ChronologyCaption string `xml:",omitempty"`
	ChronologyValue   string
	Ext               *Ext
}

type ComponentId struct {
	ComponentIdentifierType SchemeValue
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
	ApplicationProfileSupportedType []SchemeValue
	ConsortiumAgreement             []SchemeValue
	AgencyUserPrivilegeType         []SchemeValue
	Ext                             *Ext
}

type DeleteItemFields struct {
	BibliographicDescription *BibliographicDescription
	ItemUseRestrictionType   []SchemeValue
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	SecurityMarker           *SchemeValue
	SensitizationFlag        *SensitizationFlag
	Ext                      *Ext
}

type DeleteRequestFields struct {
	UserId                *UserId
	ItemId                *ItemId
	RequestType           *SchemeValue
	RequestScopeType      *SchemeValue
	RequestStatusType     *SchemeValue
	ShippingInformation   *ShippingInformation
	EarliestDateNeeded    string `xml:",omitempty"` // xs:dateTime
	NeedBeforeDate        string `xml:",omitempty"` // xs:dateTime
	PickupLocation        *SchemeValue
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
	UserLanguage           []SchemeValue
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	Ext                    *Ext
}

type Destination struct {
	BinNumber string
	Location  *Location
	Ext       *Ext
}

type ElectronicAddress struct {
	ElectronicAddressType SchemeValue
	ElectronicAddressData string
	Ext                   *Ext
}

type ElectronicResource struct {
	ElectronicDataFormatType *SchemeValue
	ActualResource           string `xml:",omitempty"`
	ReferenceToResource      string
	Ext                      *Ext
}

type ElectronicResourceProvidedFlag struct {
}

type EnumerationLevelInstance struct {
	EnumerationLevel   int    // xs:positiveInteger
	EnumerationCaption string `xml:",omitempty"`
	Ext                *Ext
	EnumerationValue   string
}

type Ext struct {
	V []byte `xml:",innerxml"`
}

type FiscalTransactionInformation struct {
	FiscalActionType                    SchemeValue
	FiscalTransactionReferenceId        *FiscalTransactionReferenceId
	RelatedFiscalTransactionReferenceId []RelatedFiscalTransactionReferenceId
	FiscalTransactionType               SchemeValue
	ValidFromDate                       string `xml:",omitempty"` // xs:dateTime
	ValidToDate                         string `xml:",omitempty"` // xs:dateTime
	Amount                              Amount
	PaymentMethodType                   *SchemeValue
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
	StructuredHoldingsData   []StructuredHoldingsData
	UnstructuredHoldingsData string `xml:",omitempty"`
	Ext                      *Ext
}

type IndeterminateLoanPeriodFlag struct {
}

type ItemDescription struct {
	CallNumber           string `xml:",omitempty"`
	CopyNumber           string `xml:",omitempty"`
	ItemDescriptionLevel *SchemeValue
	HoldingsInformation  *HoldingsInformation
	NumberOfPieces       *int // xs:positiveInteger
	Ext                  *Ext
}

type ItemDetails struct {
	ItemId                      ItemId
	BibliographicDescription    BibliographicDescription
	DateCheckedOut              string   `xml:",omitempty"` // xs:dateTime
	DateRenewed                 []string `xml:",omitempty"` // xs:dateTime
	DateDue                     string   `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	NonReturnableFlag           *NonReturnableFlag
	DateReturned                string `xml:",omitempty"` // xs:dateTime
	Ext                         *Ext
}

type ItemId struct {
	AgencyId            *SchemeValue
	ItemIdentifierType  *SchemeValue
	ItemIdentifierValue string
	Ext                 *Ext
}

type ItemOptionalFields struct {
	BibliographicDescription *BibliographicDescription
	ItemUseRestrictionType   []SchemeValue
	CirculationStatus        *SchemeValue
	HoldQueueLength          *int // xs:nonnegativeInteger
	ItemDescription          *ItemDescription
	Location                 []Location
	PhysicalCondition        *PhysicalCondition
	ElectronicResource       *ElectronicResource
	SecurityMarker           *SchemeValue
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
	ItemId                      ItemId
	ReminderLevel               int    // xs:positiveInteger
	DateDue                     string `xml:",omitempty"` // xs:dateTime
	IndeterminateLoanPeriodFlag *IndeterminateLoanPeriodFlag
	Amount                      Amount
	Title                       string `xml:",omitempty"`
	MediumType                  *SchemeValue
	Ext                         *Ext
}

type LoanedItemsCount struct {
	CirculationStatus      *SchemeValue
	ItemUseRestrictionType *SchemeValue
	LoanedItemCountValue   int // xs:nonnegativeInteger
	Ext                    *Ext
}

type LoanedItemsDesired []byte // TODO

type Location struct {
	LocationType  SchemeValue
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
	PersonalNameInformation     *PersonalNameInformation
	OrganizationNameInformation []OrganizationNameInformation
	Ext                         *Ext
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
	OrganizationNameType SchemeValue
	OrganizationName     string
	Ext                  *Ext
}

type PaidFeeAmount struct {
	CurrencyCode  SchemeValue
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type Pending struct {
	DateOfExpectedReply string `xml:",omitempty"` // xs:dateTime
	Ext                 *Ext
}

type PersonalNameInformation struct {
	//UnstructuredPersonalUserName string `xml:",omitempty"`
	StructuredPersonalUserName   StructuredPersonalUserName
	UnstructuredPersonalUserName string `xml:",omitempty"`
	Ext                          *Ext
}

type PhysicalAddress struct {
	StructuredAddress   *StructuredAddress
	UnstructuredAddress *UnstructuredAddress
	PhysicalAddressType SchemeValue
	Ext                 *Ext
}

type PhysicalCondition struct {
	PhysicalConditionType    SchemeValue
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
	ProblemType    SchemeValue
	ProblemDetail  string `xml:",omitempty"`
	ProblemElement string `xml:",omitempty"`
	ProblemValue   string `xml:",omitempty"`
	Ext            *Ext
}

type PromptInput struct {
	AuthenticationInputType      SchemeValue
	AuthenticationDataFormatType SchemeValue
	SensitiveDataFlag            *SensitiveDataFlag
	Ext                          *Ext
}

type PromptOutput struct {
	AuthenticationPromptData string
	AuthenticationPromptType SchemeValue
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
	RequestIdentifierType  *SchemeValue
	RequestIdentifierValue string
	Ext                    *Ext
}

type RequestedItem struct {
	RequestId         RequestId
	ItemId            *ItemId
	RequestType       SchemeValue
	RequestStatusType SchemeValue
	DatePlaced        string // xs:dateTime
	PickupDate        string `xml:",omitempty"` // xs:dateTime
	PickupLocation    *SchemeValue
	PickupExpiryDate  string `xml:",omitempty"` // xs:dateTime
	ReminderLevel     *int   // xs:positiveInteger
	HoldQueuePosition *int   // xs:positiveInteger
	Title             string `xml:",omitempty"`
	MediumType        *SchemeValue
	Ext               *Ext
}

type RequestedItemsCount struct {
	RequestType             *SchemeValue
	CirculationStatus       *SchemeValue
	ItemUseRestrictionType  *SchemeValue
	RequestedItemCountValue int // xs:nonnegativeInteger
	Ext                     *Ext
}

type RequiredFeeAmount struct {
	CurrencyCode  SchemeValue
	MonetaryValue int // xs:integer
	Ext           *Ext
}

type ResourceDesired struct {
}

type RoutingInformation struct {
	RoutingInstructions string
	Destination         Destination
	RequestType         *SchemeValue
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
	PhysicalAddress      *PhysicalAddress
	ElectronicAddress    *ElectronicAddress
	Ext                  *Ext
}

type StructuredAddress struct {
	LocationWithinBuilding string `xml:",omitempty"`
	HouseName              string `xml:",omitempty"`
	Street                 string `xml:",omitempty"`
	PostOfficeBox          string `xml:",omitempty"`
	District               string `xml:",omitempty"`
	Line1                  string
	Line2                  string `xml:",omitempty"`
	Locality               string `xml:",omitempty"`
	Region                 string `xml:",omitempty"`
	Country                string `xml:",omitempty"`
	PostalCode             string `xml:",omitempty"`
	CareOf                 string `xml:",omitempty"`
	Ext                    *Ext
}

type StructuredHoldingsData struct {
	HoldingsEnumeration HoldingsEnumeration
	HoldingsChronology  *HoldingsChronology
	Ext                 *Ext
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
	UnstructuredAddressType SchemeValue
	UnstructuredAddressData string
	Ext                     *Ext
}

type UserAddressInformation struct {
	UserAddressRoleType SchemeValue
	ValidFromDate       string `xml:",omitempty"` // xs:dateTime
	ValidToDate         string `xml:",omitempty"` // xs:dateTime
	PhysicalAddress     *PhysicalAddress
	ElectronicAddress   *ElectronicAddress
	Ext                 *Ext
}

type UserElementEnum struct {
	AccountBalance AccountBalance
	AccountDetails []AccountDetails
	Ext            *Ext
}

type UserFiscalAccountDesired []byte

type UserId struct {
	AgencyId            *SchemeValue
	UserIdentifierType  *SchemeValue
	UserIdentifierValue string
	Ext                 *Ext
}

type UserNoticeDetails struct {
	NoticeType        SchemeValue
	NoticeContent     string `xml:",omitempty"`
	NoticeItem        []NoticeItem
	UserFiscalAccount *UserFiscalAccount
	UserPrivilege     *UserPrivilege
	Ext               *Ext
}

type UserOptionalFields struct {
	NameInformation        *NameInformation
	UserAddressInformation []UserAddressInformation
	DateOfBirth            string `xml:",omitempty"` // xs:dateTime
	UserLanguage           []SchemeValue
	UserPrivilege          []UserPrivilege
	BlockOrTrap            []BlockOrTrap
	UserId                 []UserId
	PreviousUserId         []PreviousUserId
	Ext                    *Ext
}

type UserPrivilege struct {
	AgencyId                 SchemeValue
	AgencyUserPrivilegeType  SchemeValue
	ValidFromDate            string `xml:",omitempty"` // xs:dateTime
	ValidToDate              string `xml:",omitempty"` // xs:dateTime
	UserPrivilegeFee         *UserPrivilegeFee
	UserPrivilegeStatus      *UserPrivilegeStatus
	UserPrivilegeDescription string `xml:",omitempty"`
	Ext                      *Ext
}

type UserPrivilegeFee struct {
	Amount            Amount
	PaymentMethodType *SchemeValue
	Ext               *Ext
}

type UserPrivilegeStatus struct {
	UserPrivilegeStatusType   SchemeValue
	DateOfUserPrivilegeStatus string `xml:",omitempty"` // xs:dateTime
	Ext                       *Ext
}

// TODO complete theese, missing from autogenerated output
type UserFiscalAccount struct{}

func (r LookupUser) Type() requestType        { return TypeLookupUser }
func (r RequestItem) Type() requestType       { return TypeRequestItem }
func (r CheckOutItem) Type() requestType      { return TypeCheckOutItem }
func (r CheckInItem) Type() requestType       { return TypeCheckInItem }
func (r AcceptItem) Type() requestType        { return TypeAcceptItem }
func (r RenewItem) Type() requestType         { return TypeRenewItem }
func (r CancelRequestItem) Type() requestType { return TypeCancelRequestItem }

func (r LookupUserResponse) Type() responseType        { return TypeLookupUserResponse }
func (r RequestItemResponse) Type() responseType       { return TypeRequestItemResponse }
func (r CheckOutItemResponse) Type() responseType      { return TypeCheckOutItemResponse }
func (r CheckInItemResponse) Type() responseType       { return TypeCheckInItemResponse }
func (r RenewItemResponse) Type() responseType         { return TypeRenewItemResponse }
func (r CancelRequestItemResponse) Type() responseType { return TypeCancelRequestItemResponse }
func (r AcceptItemResponse) Type() responseType        { return TypeAcceptItemResponse }
