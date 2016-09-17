package ncip

import "encoding/xml"

type SchemeValue struct {
	Scheme string `xml:",attr,omitempty"`
	Value  string `xml:",innerxml"`
}

// AcceptItem represents a requests that the responding application accept an
// Item to be circulated to a User. The responding  application may be a third
// party that has no prior knowledge of either the User or the Item. The
// initiation message identifies the action the responding agency is requested
// to take when it receives the item. The request may include a date by which
// the initiating application requires the Item to be returned, an indication
// that the user may apply to renew the loan of the Item, and financial data
// relating to the supply of the Item. If there is a possibility that the
// responding application has no prior // knowledge of either the User or the
// Item, the request may optionally include data about the User and/or the
// Item.
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

// AcceptItemResponse represents a response to an AcceptItem request. The
// responder accepts the Item for the action specified. It returns the Request
// Id. It may optionally provide an Item Id that the initiating application may
// use to reference the Item in subsequent messages.
type AcceptItemResponse struct {
	XMLName        xml.Name
	ResponseHeader *ResponseHeader
	Problem        []Problem
	RequestId      *RequestId
	ItemId         *ItemId
	Ext            *Ext
}

// AgencyCreated represents a request that informs the responding application
// that a new Agency has been created. The initiating // application provides
// the Agency Id and details about the Agency.
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

// AgencyCreatedResponse represents a response to a AgencyCreated request. The
// responding application replies that it understood the notification message.
type AgencyCreatedResponse struct {
	XMLName        xml.Name
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// AgencyUpdated represents a request that informs the responding application
// that one or more data elements of an Agency have been updated. The
// initiating application provides elements that have been deleted with the
// value for each element and/or elements that have been added with the value
// for each element. If the value for an existing element has changed, the
// message must include that element in both the delete and add sections of the
// message.
type AgencyUpdated struct {
	InitiationHeader   *InitiationHeader
	AgencyId           SchemeValue
	DeleteAgencyFields *DeleteAgencyFields
	AddAgencyFields    *AddAgencyFields
	Ext                *Ext
}

// AgencyUpdatedResponse represents a response to an AgencyUpdated request.
// The responding application replies that it understood the notification
// message.
type AgencyUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// CancelRecallItem represents a request that the responding application cancel
// the recall of an Item that it had previously been asked to recall. The
// initiating application may also request data about the User and/or Item
// involved with this recall cancellation.
type CancelRecallItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	ItemElementType  []SchemeValue
	UserElementType  []SchemeValue
	Ext              *Ext
}

// CancelRecallItemResponse represents a response to a CancelRecallItem request
// The responding application cancels the recall of the Item and takes
// appropriate fiscal actions. It may also supply the data elements requested.
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

// CancelRequestItem item represents a equest that the responding application
// cancel a previous request for an Item. The initiating application may also
// request data about the User and/or Item involved with this request
// cancellation.
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

// CancelRequestItemResponse represents a response to a CancelRequestItem
// request. The responding application cancels the request. It may provide
// updated fiscal data about the User. It may also supply the data elements
// requested.
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

// CheckInItem represents a request that the responding application check in an
// Item. It also permits the initiating application to request data about the
// User and/or Item involved with this check in.
type CheckInItem struct {
	XMLName          xml.Name
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	ItemElementType  []SchemeValue
	UserElementType  []SchemeValue
	Ext              *Ext
}

// CheckInItemResponse represents a response to a CheckInItem request. The
// responder checks in the Item and returns any routing data, fine/fee data,
// and/or requested User or Item data.
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

// CheckOutItem represents a request that the responding application check out
// an Item to a User. It also permits the initiating application to acknowledge
// the fee amount (if any) associated with the check out. The initiating
// application may also request data about the User and/or Item involved with this
// check out.
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

// CheckOutItemResponse represents a response to CheckOutItem request. The
// responding application checks out the Item to the User until the date
// indicated in the response. It may also supply the data elements requested.
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

// CirculationStatusChangeReported represents a request that the responding
// application mark an Item as reported returned, missing, lost, or never
// borrowed. The initiating application may also request data about the User
// and/or Item involved with this status change.
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

// CirculationStatusChangeReportedResponse represents a response to a
// CirculationStatusChange request. The responding application marks the Item
// as specified by the initiator. It may also supply the data elements
// requested.
type CirculationStatusChangeReportedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// CirculationStatusUpdated represents a request that informs the responding
// application that the circulation status of an Item has been updated. The
// initiating application provides the unique Id of the item Id and the new
// circulation status. Only the following three values from Circulation Status
// are valid here: Renew Still Pending, Item Not Renewed, or Item Overdue. No
// other values of Circulation Status are permitted for this service.
type CirculationStatusUpdated struct {
	InitiationHeader  *InitiationHeader
	ItemId            ItemId
	CirculationStatus SchemeValue
	Ext               *Ext
}

// CirculationStatusUpdatedResponse represents a response to a
// CirculationStatusUpdate request. The responding application replies that it
// understood the notification message.
type CirculationStatusUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// CreateAgency represents a request that the responding application create a
// record for an agency. The initiating application supplies the data elements
// to be used to create the Agency. It may optionally supply a proposed Agency
// Id.
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

// CreateAgencyResponse represents a response to a CreateAgency request.
// The responding application creates a record for an agency and returns a
// Agency Id.
type CreateAgencyResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	AgencyId       *SchemeValue
	Ext            *Ext
}

// CreateItem represents a request that the responding application create a
// record for an item. The initiating application supplies the data elements to
// be used to create the Item. It may optionally supply a proposed Item Id.
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

// CreateItemResponse represents a response to a CreateItem request. The
// responding application creates a record for an item and returns a Item Id.
type CreateItemResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

// CreateUser represents a request that the responding application create a
// record for a user. The initiating application supplies the data elements to
// be used to create the User. It may optionally supply a proposed User Id.
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

// CreateUserResponse represents a response to CreateUser request. The
// responding application creates a record for a user and returns a User Id.
type CreateUserResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         *UserId
	Ext            *Ext
}

// CreateUserFiscalTransaction represents a request that the responding
// application create a new fiscal transaction for a service provided to a
// User. The initiating application provides the type of update and fiscal
// details about the update it wants done.
type CreateUserFiscalTransaction struct {
	InitiationHeader             *InitiationHeader
	MandatedAction               *MandatedAction
	UserId                       *UserId
	AuthenticationInput          []AuthenticationInput
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

// CreateUserFiscalTransactionResponse represents a response to a
// CreateUserFiscalTransaction request. The responding application creates
// fiscal transaction data for the User.
type CreateUserFiscalTransactionResponse struct {
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	UserId                       UserId
	FiscalTransactionReferenceId FiscalTransactionReferenceId
	Ext                          *Ext
}

// DeleteItem represents a request that asks that the responding application
//delete an item record. This may also result in deletion of the associated
// bibliographic record, depending on the rules of the responding application.
// The purpose of this service is to enable applications to remove unnecessary
// items in weeding and binding situations as well as cleaning up after
// creation of a temporary item. The initiating application provides an
// identifier and the responder is not required to return any data other than
// an acknowledgement or an error.
type DeleteItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	Ext              *Ext
}

// DeleteItemResponse represents a response to a DeleteItem request. The
// responding application deletes the item specified by the initiating
// application.
type DeleteItemResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

// DeleteUser represents a request that asks that the responding application
// delete a User. This service is intended to enable resource sharing systems
// and similar applications to remove temporary users created during request
// processes. The initiating application provides an identifier and the
// responder is not required to return any data other than an acknowledgement
// or an error.
type DeleteUser struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	Ext                 *Ext
}

// DeleteUserResponse represents a response to a DeleteUser request. The
// responding application deletes the user record specified by the initiating
// application.
type DeleteUserResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         *UserId
	Ext            *Ext
}

// ItemCheckedIn represents a request that informs the responding application
// that an Item has been checked in. The initiating application provides the
// Unique Item Id and may optionally provide other details associated with the
// check in.
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

// ItemCheckedInResponse represents a response to a ItemCheckedIn request. The
// responding application replies that it understood the notification message.
type ItemCheckedInResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemCheckedOut represents a request that informs the responding application
// that an Item has been checked out to a User. The initiating application
// provides the Item Id and User Id involved in the check out, and may
// optionally provide other details associated with the check out.
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

// ItemCheckedOutResponse represents a response to an ItenCheckedOut request.
// The responding application replies that it understood the notification
// message.
type ItemCheckedOutResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemCreated represents a request that informs the responding application
// that a new Item has been created. The initiating application provides the
// Item Id and details about the Item, and may optionally provide other details
// associated with the Item.
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

// ItemCreatedResponse represents a response to an ItemCreated request.
// The responding application replies that it understood the notification
// message.
type ItemCreatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemRecallCancelled represents a request that informs the responding
// application that a recall has been cancelled. The initiating application
// provides the Item Id, and may optionally provide the User Id and other data
// related to cancellation of the recall.
type ItemRecallCancelled struct {
	InitiationHeader             *InitiationHeader
	UserId                       *UserId
	ItemId                       ItemId
	FiscalTransactionInformation *FiscalTransactionInformation
	ItemOptionalFields           *ItemOptionalFields
	UserOptionalFields           *UserOptionalFields
	Ext                          *Ext
}

// ItemRecallCancelledResponse represents a response to an ItemRecallCanceled
// request. The responding application replies that it understood the
// notification message
type ItemRecallCancelledResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemRecalled represents a request that informs the responding application
// that an Item has been recalled. The initiating application provides the Item
// Id of the recalled item and a new due date, and may optionally provide other
// details associated with the recall.
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

// ItemRecalledResponse represents a response to an ItemRecalled request. The
// responding application replies that it understood the notification message.
type ItemRecalledResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemReceived represents a request that informs the responding application
// that an agency has received an item. If the item has been received on behalf
// of a user, then data about that User should also be included in the message.
// The initiating application may optionally provide other details associated
// with the receipt of the Item.
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

// ItemReceivedResponse represents a response to an ItemReceived request. The
// responding application replies that it understood the notification message.
type ItemReceivedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemRenewed represents a request that informs the responding application
// that an Item has been renewed. The initiating application provides the Item
// Id and the new due date, and may optionally supply other details associated
// with the renewal.
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

// ItemRenewedResponse represents a response to an ItemRenewed request. The
// responding application replies that it understood the notification message.
type ItemRenewedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemRequestCancelled represents a request that informs the responding
// application that a request for an Item has been cancelled. The initiating
// application provides the Item Id and/or Request Id, and the User Id, and the
// type of the request being cancelled. It may optionally provide other details
// associated with the request being cancelled.
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

// ItemRequestCancelledResponse represents a response to an ItemRequestCancelled
// request. The responding application replies that it understood the notification
// message.
type ItemRequestCancelledResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemRequestUpdated represents a request that informs the responding
// application that a request for an Item has been updated. The initiating
// application provides the Item Id and the User Id, or the Request  Id. The
// initiating application provides elements that have been deleted with the
// value for each element and/or elements that have been added with the value
// for each element. If the value for an existing element has changed, the
// message must include that element in both the delete and add sections of the
// message. Other data may also be provided.
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

// ItemRequestUpdatedResponse represents a response to an ItemRequestUpdated
// request. The responding application replies that it understood the
// notification message.
type ItemRequestUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemRequested represents a request that informs the responding application
// that a request for an Item, which may not be immediately available, has been
// made. The initiating application provides the Item Id, or the Request Id and
// Bibliographic Id. It also provides the User Id and the type of request being
// made. It may also optionally provide other details related to the request.
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

// ItemRequestedResponse represents a response to a ItemRequested request. The
// responding application replies that it understood the notification message.
type ItemRequestedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemShipped represents a request that informs the responding application
// that an Agency has shipped an Item. The date when the Item was shipped and
// the address to which it was shipped must be included in the message. When
// shipment supports a user request, User data should be supplied with the
// message. The initiating application may also provide other details
// associated with shipment of the Item.
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

// ItemShippedResponse represents a response to an ItemShipped request. The
// responding application replies that it understood the notification message.
type ItemShippedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// ItemUpdated represents a request that informs the responding application
// that one or more data elements of an Item have been updated. The initiating
// application provides elements that have been deleted with the value for each
// element and/or elements that have been added with the value for each element.
// If the value for an existing element has changed, the message must include
// that element in both the delete and add sections of the message. This
// service may not be used to report changes to circulation status.
type ItemUpdated struct {
	InitiationHeader *InitiationHeader
	ItemId           ItemId
	DeleteItemFields *DeleteItemFields
	AddItemFields    *AddItemFields
	Ext              *Ext
}

// ItemUpdatedResponse represents a response to an ItemUpdated request. The
// responding application replies that it understood the notification message.
type ItemUpdatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// LookupAgency represents a request that requests data about a particular
// Agency known to the responding application. The initiating application
// provides the Id of the Agency and the list of elements for which data is
// requested.
type LookupAgency struct {
	InitiationHeader  *InitiationHeader
	AgencyId          SchemeValue
	AgencyElementType []SchemeValue
	Ext               *Ext
}

// LookupAgencyResponse represents a response to a LookupAgency request. The
// responding application returns the requested data to the initiating application.
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

// LookupItem represents a request that data about a particular Item known to
// the responding application. The initiator provides the Id of the Item and a
// list of elements for which data is requested.
type LookupItem struct {
	InitiationHeader         *InitiationHeader
	ItemId                   *ItemId
	RequestId                *RequestId
	ItemElementType          []SchemeValue
	CurrentBorrowerDesired   *CurrentBorrowerDesired
	CurrentRequestersDesired *CurrentRequestersDesired
	Ext                      *Ext
}

// LookupItemResponse represents a response to a LookupItem request. The
// responding application returns the requested data to the initiating
// application.
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

// LookupRequest represents a request that data about a particular Request
// known to the responding application. The initiator provides the Id of the
// Request and a list of elements for which data is requested.
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

// LookupRequestResponse represents a response to a LookupRequest request. The
// responding application returns the requested data to the initiating
// application.
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

// LookupUser represents a request that data about a particular User known to
// the responding application. The initiator provides the Id of the User and a
// list of elements for which data is requested.
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

// LookupUserResponse represents a response to a LookupUser request. The
// responding application returns the requested data to the initiating
// application.
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

// RecallItem represents a request that the responding application recall an
// Item from a User. The initiating application may propose a new date due for
// the Item. The initiating application may also request data about the User
// and/or Item involved with this recall.
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

// RecallItemResponse represents a response to a RecallItem request. The
// responding application agrees to recall the Item and to take appropriate
// action. It may supply the new date on which the Item is now due. It may
// supply the data elements requested.
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

// RenewItem represents a request that the responding application renew an Item
// for a User. The initiating application may include a suggested revision of
// the due date and an acknowledgement of fee amount. The initiating
// application may also request data about the User and/or Item involved with
// this renewal.
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

// RenewItemResponse represents a response to a RenewItem request. The
// responding application renews the Item to the User and provides a
// revised date on which the Item now is due. It may also supply the data
// elements requested.
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

// ReportCirculationStatusChange represents a request that the responding
// application mark an Item as reported returned, missing, lost, or never
// borrowed. The initiating application may also request data about the User
// and/or Item involved with this status change.
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

// ReportCirculationStatusChangeResponse represents a response to a
// ReportCirculationStatusChange request.The responding application marks the
// Item as specified by the initiator. It may also supply the data elements
// requested.
type ReportCirculationStatusChangeResponse struct {
	ResponseHeader     *ResponseHeader
	Problem            []Problem
	ItemId             ItemId
	UserId             *UserId
	ItemOptionalFields *ItemOptionalFields
	UserOptionalFields *UserOptionalFields
	Ext                *Ext
}

// RequestItem represents a request that the responding application place a
// request on an Item for a User whether or not the Item is immediately
// available. The initiating application indicates the type of request being
// made. The initiating application may optionally provide an acknowledgement
// of the fee to be charged for the service. The initiating application may
// also request data about the User and/or Item involved with this request.
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

// RequestItemResponse represents a response to a RequestItem request. The
// responding application places the request and provides data about where the
// Item may be picked up and the date it expects the Item to be available. It
// may also supply the data elements requested.
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

// SendUserNotice represents a request that the responding application send a
// notice to a user. The initiating application provides the Id of the User and
// the type of notice to send. The initiating application may optionally
// provide a date on which it desires the notice be sent, and the content of
// the notice.
type SendUserNotice struct {
	InitiationHeader  *InitiationHeader
	MandatedAction    *MandatedAction
	UserId            UserId
	DateToSend        string `xml:",omitempty"` // xs:dateTime
	UserNoticeDetails UserNoticeDetails
	Ext               *Ext
}

// SendUserNoticeResponse represents a response to a SendUserNotice request.
// The responding application agrees to send a notice and may provide either
// the date the notice was sent or the date on which it will send the notice.
type SendUserNoticeResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         UserId
	DateSent       string `xml:",omitempty"` // xs:dateTime
	DateWillSend   string `xml:",omitempty"` // xs:dateTime
	Ext            *Ext
}

// UndoCheckOutItem represents a request that the responding application undo
// the check out that immediately preceded this message. If the responder
// agrees to the request, processing continues as if the preceding check out
// had never occurred. This service is typically used by self-check
// applications.
type UndoCheckOutItem struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	ItemId              ItemId
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	RequestId           *RequestId
	Ext                 *Ext
}

// UndoCheckOutItemResponse represents a response to a UndoCheckOutItem
// request. The immediately preceding check out is undone.
type UndoCheckOutItemResponse struct {
	ResponseHeader               *ResponseHeader
	Problem                      []Problem
	ItemId                       ItemId
	UserId                       *UserId
	FiscalTransactionInformation *FiscalTransactionInformation
	Ext                          *Ext
}

// UpdateAgency represents a  request that the responding application update
// data about an Agency. The initiating application provides elements to be
// deleted with the value for each element and/or elements to be added with the
// value for each element. If the initiating application is changing the value
// for an existing element, it must include that element in both the delete and
// add sections of the message.
type UpdateAgency struct {
	InitiationHeader   *InitiationHeader
	MandatedAction     *MandatedAction
	AgencyId           SchemeValue
	DeleteAgencyFields *DeleteAgencyFields
	AddAgencyFields    *AddAgencyFields
	Ext                *Ext
}

// UpdateAgencyResponse represents a response to an UpdateAgency request. The
// responding application updates data about the Agency as specified by the
// initiating application.
type UpdateAgencyResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	AgencyId       *SchemeValue
	Ext            *Ext
}

// UpdateCirculationStatus represents a request that the responding application
// update the circulation status of an Item. The initiating application
// provides the Id of the item Id and the new circulation status.
type UpdateCirculationStatus struct {
	InitiationHeader  *InitiationHeader
	MandatedAction    *MandatedAction
	ItemId            ItemId
	CirculationStatus SchemeValue
	Ext               *Ext
}

// UpdateCirculationStatusResponse represents a response to a
// UpdateCirculationStats request. The responding application updates the
// circulation status as specified by the initiating application.
type UpdateCirculationStatusResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

// UpdateItem represents a request that the responding application update data
// about an Item. The initiating application provides elements to be deleted
// with the value for each element and/or elements to be added with the value
// for each element. If the initiating application is changing the value for an
// existing element, it must include that element in both the delete and add
// sections of the message. This service cannot be used to change the
// circulation status of an Item.
type UpdateItem struct {
	InitiationHeader *InitiationHeader
	MandatedAction   *MandatedAction
	ItemId           ItemId
	DeleteItemFields *DeleteItemFields
	AddItemFields    *AddItemFields
	Ext              *Ext
}

// UpdateItemResponse represents a response to an UpdateItem request. The
// responding application updates data about the Item as specified by the
// initiating application.
type UpdateItemResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	ItemId         *ItemId
	Ext            *Ext
}

// UpdateRequestItem represents a request that asks the responding application
// to update data about a request for an Item made by a User. The initiating
// application provides elements to be deleted with the value for each element
// and/or elements to be added with the value for each element. If the
// initiating application is changing the value for an existing element, it
// must include that element in both the delete andadd sections of the message.
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

// UpdateRequestItemResponse represents a response to a UpdateRequestItem
// request. The responding application updates the Item request as specified by
// the initiating application.
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

// UpdateUser represents a request that the responding application update data
// about a User. The initiating application provides elements to be deleted
// with the value for each element and/or elements to be added with the value
// for each element. If the initiating application is changing the value for an
// existing element, it must include that element in both the delete and add
//sections of the message.
type UpdateUser struct {
	InitiationHeader    *InitiationHeader
	MandatedAction      *MandatedAction
	UserId              *UserId
	AuthenticationInput []AuthenticationInput
	DeleteUserFields    *DeleteUserFields
	AddUserFields       *AddUserFields
	Ext                 *Ext
}

// UpdateUserResponse represents a response to an UpdateUser request. The
// responding application updates data about the User as specified by the
// initiating application.
type UpdateUserResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	UserId         *UserId
	Ext            *Ext
}

// UserCreated represents a request that informs the responding application
// that a record for a user has been created. The initiating application
// provides Name Information and the User Id and may optionally provide
// additional data about the new User.
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

// UserCreatedResponse represents a response to an UserCreated request. The
// responding application replies that it understood the notification message.
type UserCreatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// UserFiscalTransactionCreated represents a request that informs the responding
// application that a new fiscal transaction has been applied to a User’s
// fiscal account.
type UserFiscalTransactionCreated struct {
	InitiationHeader             *InitiationHeader
	UserId                       UserId
	FiscalTransactionInformation FiscalTransactionInformation
	Ext                          *Ext
}

// UserFiscalTransactionCreatedResponse represents a response to an
// UserFiscalTransactionCreated request. The responding application replies
// that it understood the notification message.
type UserFiscalTransactionCreatedResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// UserNoticeSent represents a request that informs the responding application
// that a notice has been sent to a User. The initiating application provides
// the User Id, the type of notice, and details about the notice.
type UserNoticeSent struct {
	InitiationHeader  *InitiationHeader
	UserId            UserId
	DateSent          string // xs:dateTime
	UserNoticeDetails UserNoticeDetails
	Ext               *Ext
}

// UserNoticeSentResponse represents a response to an UserNoticeSent request.
// The responding application replies that it understood the notification
// message.
type UserNoticeSentResponse struct {
	ResponseHeader *ResponseHeader
	Problem        []Problem
	Ext            *Ext
}

// UserUpdated represents a request that informs the responding application
// that one or more data elements of a User have been updated. The initiating
// application provides elements that have been deleted with the value for each
// element and/or elements that have been added with the value for each
// element. If the value for an existing element has changed, the message must
// include that element in both the delete and add sections of the message.
type UserUpdated struct {
	InitiationHeader *InitiationHeader
	UserId           UserId
	DeleteUserFields *DeleteUserFields
	AddUserFields    *AddUserFields
	Ext              *Ext
}

// UserUpdatedResponse represents a response to an UserUpdated request. The
// responding application replies that it understood the notification message.
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
