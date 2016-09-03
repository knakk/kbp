package ncip

import (
	"bytes"
	"encoding/xml"
	"testing"
)

func mustParseRequest(s string) Request {
	req, err := DecodeRequest(bytes.NewBufferString(s))
	if err != nil {
		panic("mustParseRequest: " + err.Error())
	}
	return req
}

func encodeRequest(req Request) string {
	output, err := xml.MarshalIndent(&req, "", "  ")
	if err != nil {
		panic("decodeAndEncodeRequest: " + err.Error())
	}

	return string(output)
}

func TestParseRequest(t *testing.T) {
	var tests = []struct {
		xml  string
		want Request
	}{
		{xml: `<RequestItem>
  <RequestScopeType>
    <Scheme>http://www.niso.org/ncip/v1_0/imp1/schemes/requestscopetype/requestscopetype.scm</Scheme>
    <Value>Bibliographic Item</Value>
  </RequestScopeType>
  <RequestType>
    <Scheme>http://www.niso.org/ncip/v1_0/imp1/schemes/requesttype/requesttype.scm</Scheme>
    <Value>Hold</Value>
  </RequestType>
  <UserId>
    <UserIdentifierValue>1231231230</UserIdentifierValue>
    <AgencyId>
      <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
      <Value>DK-715700</Value>
    </AgencyId>
  </UserId>
  <InitiationHeader>
    <FromAgencyId>
      <AgencyId>
        <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
        <Value>DK-190101</Value>
      </AgencyId>
    </FromAgencyId>
    <ToAgencyId>
      <AgencyId>
        <Scheme>http://biblstandard.dk/isil/schemes/1.1/</Scheme>
        <Value>DK-715700</Value>
      </AgencyId>
    </ToAgencyId>
    <FromAgencyAuthentication>secret</FromAgencyAuthentication>
  </InitiationHeader>
  <NeedBeforeDate>2008-09-29T00:00+01:00</NeedBeforeDate>
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
</RequestItem>`,
			want: RequestItem{
				UserId: &UserId{
					UserIdentifierValue: "1231231230",
					AgencyId: &AgencyId{
						SchemeValue{
							Scheme: "http://biblstandard.dk/isil/schemes/1.1/",
							Value:  "DK-715700",
						},
					},
				},
				RequestScopeType: SchemeValue{
					Scheme: "http://www.niso.org/ncip/v1_0/imp1/schemes/requestscopetype/requestscopetype.scm",
					Value:  "Bibliographic Item",
				},
				RequestType: RequestType{
					SchemeValue{
						Scheme: "http://www.niso.org/ncip/v1_0/imp1/schemes/requesttype/requesttype.scm",
						Value:  "Hold",
					},
				},
				RequestId: &RequestId{
					AgencyId: AgencyId{
						SchemeValue{
							Scheme: "http://biblstandard.dk/isil/schemes/1.1/",
							Value:  "DK-190101",
						},
					},
					RequestIdentifierValue: "23770051",
				},
				InitiationHeader: &InitiationHeader{
					FromAgencyId: FromAgencyId{
						AgencyId{
							SchemeValue{
								Scheme: "http://biblstandard.dk/isil/schemes/1.1/",
								Value:  "DK-190101",
							},
						},
					},
					ToAgencyId: ToAgencyId{
						AgencyId{
							SchemeValue{
								Scheme: "http://biblstandard.dk/isil/schemes/1.1/",
								Value:  "DK-715700",
							},
						},
					},
					FromAgencyAuthentication: "secret",
				},
				NeedBeforeDate: "2008-09-29T00:00+01:00",
				BibliographicId: &BibliographicId{
					BibliographicRecordId: &BibliographicRecordId{
						BibliographicRecordIdentifier: "44397315",
						BibliographicRecordIdentifierCode: &SchemeValue{
							Scheme: "http://biblstandard.dk/ncip/schemes/faust/1.0/",
							Value:  "FAUST",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		got := encodeRequest(test.want)
		if got != test.xml {
			t.Errorf("got:\n%d:%s\nwant:\n%d:%s", len(got), got, len(test.xml), test.xml)
		}
	}
}
