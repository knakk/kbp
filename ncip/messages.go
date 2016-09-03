package ncip

import "time"

type maybeString string

type InitiationHeader struct {
	// Mandatory
	FromAgencyId FromAgencyId
	ToAgencyId   ToAgencyId

	// Optional
	FromSystemAuthentication string `xml:",omitempty"`
	FromAgencyAuthentication string `xml:",omitempty"`
	FromSystemId             *SchemeValue
	OnBehalfOfAgency         *AgencyId
	ToSystemId               *SchemeValue
	ApplicationProfileType   *SchemeValue
}

type ResponseHeader struct {
	// Mandatory
	FromAgencyId AgencyId
	ToAgencyId   AgencyId

	// Optional fields
	FromSystemAuthentication string `xml:",omitempty"`
	FromAgencyAuthentication string `xml:",omitempty"`
	FromSystemId             *SchemeValue
	ToSystemId               *SchemeValue
	Problem                  []Problem
}

type FromAgencyId struct {
	AgencyId AgencyId
}

type ToAgencyId struct {
	AgencyId AgencyId
}

type AgencyId struct {
	SchemeValue
}

type SchemeValue struct {
	Scheme string
	Value  string
}

type Problem struct {
	// ProblemDetail
	// ProblemElement
	// ProblemValue

	ProblemType SchemeValue
}

type RequestId struct {
	AgencyId               AgencyId
	RequestIdentifierValue string
	RequestIdentifierType  *SchemeValue
}

type MandateAction struct {
	DateEventOccurred time.Time
}

type RequestType struct {
	SchemeValue
}

type UserId struct {
	UserIdentifierValue string
	UserIdentifierType  *SchemeValue
	AgencyId            *AgencyId
}

type AuthenticationInput struct {
	AuthenticationInputData      string
	AuthenticationDataFormatType SchemeValue
	AuthenticationInputType      SchemeValue
}

type BibliographicId struct {
	// Mandatory one-off:
	BibliographicItemId   *BibliographicItemId
	BibliographicRecordId *BibliographicRecordId
}

type BibliographicItemId struct {
	// TODO
}

type BibliographicRecordId struct {
	// Mandatory
	BibliographicRecordIdentifier string

	// Mandatory one-off
	AgencyId                          *AgencyId
	BibliographicRecordIdentifierCode *SchemeValue
}

// Requests:

type AcceptItem struct {
	// Mandatory
	RequestId RequestId

	// Optional
	InitiationHeader *InitiationHeader
	MandateAction    *MandateAction
}

type RequestItem struct {
	// Mandatory
	RequestScopeType SchemeValue
	RequestType      RequestType

	// One-of-mandatory
	UserId              *UserId
	AuthenticationInput *AuthenticationInput

	// Optional
	InitiationHeader *InitiationHeader
	MandateAction    *MandateAction
	NeedBeforeDate   string `xml:",omitempty"`
	BibliographicId  *BibliographicId
	RequestId        *RequestId
}

func (r RequestItem) Type() requestType { return TypeRequestItem }

// Responses
