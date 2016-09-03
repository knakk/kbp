package ncip

type Request interface {
	Type() requestType
}

type Response interface {
	Type() responseType
}

type Handler interface {
	Process(Request) Response
}

type requestType int

const (
	illegalRequest requestType = iota
	TypeAcceptItem
	TypeAcceptItemResponsegencyCreated
	TypeAgencyUpdated
	TypeCancelRecallItem
	TypeCancelRequestItem
	TypeCheckInItem
	TypeCheckOutItem
	TypeCirculationStatusChangeReported
	TypeCirculationStatusUpdated
	TypeCreateAgency
	TypeCreateItem
	TypeCreateUser
	TypeCreateUserFiscalTransaction
	TypeDeleteItem
	TypeDeleteUser
	TypeItemCheckedIn
	TypeItemCheckedOut
	TypeItemCreated
	TypeItemRecallCancelled
	TypeItemRecalled
	TypeItemReceived
	TypeItemRenewed
	TypeItemRequestCancelled
	TypeItemRequestUpdated
	TypeItemRequested
	TypeItemShipped
	TypeItemUpdated
	TypeLookupAgency
	TypeLookupItem
	TypeLookupRequest
	TypeLookupUser
	TypeRecallItem
	TypeRenewItem
	TypeReportCirculationStatusChange
	TypeRequestItem
	TypeSendUserNotice
	TypeUndoCheckOutItem
	TypeUpdateAgency
	TypeUpdateCirculationStatus
	TypeUpdateItem
	TypeUpdateRequestItem
	TypeUpdateUser
	TypeUserCreated
	TypeUserFiscalTransactionCreated
	TypeUserNoticeSent
	TypeUserUpdated
)

type responseType int

const (
	illegalResponse responseType = iota
	TypeAgencyCreatedResponse
	TypeAgencyUpdatedResponse
	TypeCancelRecallItemResponse
	TypeCancelRequestItemResponse
	TypeCheckInItemResponse
	TypeCheckOutItemResponse
	TypeCirculationStatusChangeReportedResponse
	TypeCirculationStatusUpdatedResponse
	TypeCreateAgencyResponse
	TypeCreateItemResponse
	TypeCreateUserResponse
	TypeCreateUserFiscalTransactionResponse
	TypeDeleteItemResponse
	TypeDeleteUserResponse
	TypeItemCheckedInResponse
	TypeItemCheckedOutResponse
	TypeItemCreatedResponse
	TypeItemRecallCancelledResponse
	TypeItemRecalledResponse
	TypeItemReceivedResponse
	TypeItemRenewedResponse
	TypeItemRequestCancelledResponse
	TypeItemRequestUpdatedResponse
	TypeItemRequestedResponse
	TypeItemShippedResponse
	TypeItemUpdatedResponse
	TypeLookupAgencyResponse
	TypeLookupItemResponse
	TypeLookupRequestResponse
	TypeLookupUserResponse
	TypeRecallItemResponse
	TypeRenewItemResponse
	TypeReportCirculationStatusChangeResponse
	TypeRequestItemResponse
	TypeSendUserNoticeResponse
	TypeUndoCheckOutItemResponse
	TypeUpdateAgencyResponse
	TypeUpdateCirculationStatusResponse
	TypeUpdateItemResponse
	TypeUpdateRequestItemResponse
	TypeUpdateUserResponse
	TypeUserCreatedResponse
	TypeUserFiscalTransactionCreatedResponse
	TypeUserNoticeSentResponse
	TypeUserUpdatedResponse
)
