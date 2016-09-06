package ncip

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func mustDecodeRequest(s string) Request {
	req, err := DecodeRequest(bytes.NewBufferString(s))
	if err != nil {
		panic("mustDecodeRequest: " + err.Error())
	}
	return req
}

func mustEncodeRequest(req Request) string {
	output, err := xml.MarshalIndent(&req, "", "  ")
	if err != nil {
		panic("mustEncodeRequest: " + err.Error())
	}

	return string(output)
}

// Verify that a XML request can be decoded, and encoded again into an identical XML request.
func TestDecodeEncodeRequest(t *testing.T) {
	requests := []string{
		`<RequestItem>
  <InitiationHeader>
    <FromAgencyId>
      <AgencyId>
        <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
        <Value>DK-190101</Value>
      </AgencyId>
    </FromAgencyId>
    <FromAgencyAuthentication>[PASSWORD]</FromAgencyAuthentication>
    <ToAgencyId>
      <AgencyId>
        <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
        <Value>DK-715700</Value>
      </AgencyId>
    </ToAgencyId>
  </InitiationHeader>
  <UserId>
    <AgencyId>
      <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
      <Value>DK-715700</Value>
    </AgencyId>
    <UserIdentifierValue>1231231230</UserIdentifierValue>
  </UserId>
  <BibliographicId>
    <BibliographicRecordId>
      <BibliographicRecordIdentifier>44397315</BibliographicRecordIdentifier>
      <BibliographicRecordIdentifierCode>
        <Scheme>http://biblstandard.dk/ncip/schemes/faust/1.0/</Scheme>
        <Value>FAUST</Value>
      </BibliographicRecordIdentifierCode>
    </BibliographicRecordId>
  </BibliographicId>
  <RequestId>
    <AgencyId>
      <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
      <Value>DK-190101</Value>
    </AgencyId>
    <RequestIdentifierValue>23770051</RequestIdentifierValue>
  </RequestId>
  <RequestType>
    <Scheme>http://www.niso.org/ncip/v1_0/imp1/schemes/requesttype/requesttype.scm</Scheme>
    <Value>Hold</Value>
  </RequestType>
  <RequestScopeType>
    <Scheme>http://www.niso.org/ncip/v1_0/imp1/schemes/requestscopetype/requestscopetype.scm</Scheme>
    <Value>Bibliographic Item</Value>
  </RequestScopeType>
  <NeedBeforeDate>2008-09-29T00:00+01:00</NeedBeforeDate>
</RequestItem>`,
	}

	for i, reqXML := range requests {
		if req := mustDecodeRequest(reqXML); mustEncodeRequest(req) != reqXML {
			t.Errorf("%d got:\n%s", i, mustEncodeRequest(req))
		}
	}
}
