package ncip

import (
	"encoding/xml"
	"fmt"
	"io"
)

const (
	header = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE NCIPMessage xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="http://www.niso.org/ncip/v2_02/imp1/xsd/ncip_v2_02.xsd" version="http://www.niso.org/ncip/v2_02/imp1/xsd/ncip_v2_02.xsd">`
)

type NCIPRequestMessage struct {
	XMLName xml.Name `xml:"NCIPMessage"`
	Request
}

type NCIPResponseMessage struct {
	XMLName xml.Name `xml:"NCIPMessage"`
	Response
}

func DecodeRequest(r io.Reader) (Request, error) {
	dec := xml.NewDecoder(r)

	for {
		t, _ := dec.Token()
		if t == nil {
			break
		}
		switch elem := t.(type) {
		case xml.StartElement:
			var reqItem Request
			switch elem.Name.Local {
			case "NCIPMessage":
				continue
			case "RequestItem":
				reqItem = &RequestItem{}
			case "LookupUser":
				reqItem = &LookupUser{}
			case "CheckOutItem":
				reqItem = &CheckOutItem{}
			case "CheckInItem":
				reqItem = &CheckInItem{}
			case "AcceptItem":
				reqItem = &AcceptItem{}
			case "RenewItem":
				reqItem = &RenewItem{}
			case "CancelRequestItem":
				reqItem = &CancelRequestItem{}
			default:
				return reqItem, fmt.Errorf("TODO: Decode %s", elem.Name.Local)
			}
			err := dec.DecodeElement(reqItem, &elem)
			return reqItem, err
		}
	}

	return nil, io.EOF
}

func DecodeResponse(r io.Reader) (Response, error) {
	dec := xml.NewDecoder(r)

	for {
		t, _ := dec.Token()
		if t == nil {
			break
		}
		switch elem := t.(type) {
		case xml.StartElement:
			var respItem Response
			switch elem.Name.Local {
			case "NCIPMessage":
				continue
			case "LookupUserResponse":
				respItem = &LookupUserResponse{}
			case "CheckOutItemResponse":
				respItem = &CheckOutItemResponse{}
			case "CheckInItemResponse":
				respItem = &CheckInItemResponse{}
			case "AcceptItemResponse":
				respItem = &AcceptItemResponse{}
			case "RenewItemResponse":
				respItem = &RenewItemResponse{}
			case "RequestItemResponse":
				respItem = &RequestItemResponse{}
			case "CancelRequestItemResponse":
				respItem = &CancelRequestItemResponse{}
			default:
				return respItem, fmt.Errorf("TODO: Decode %s", elem.Name.Local)
			}
			err := dec.DecodeElement(respItem, &elem)
			return respItem, err
		}
	}

	return nil, io.EOF
}
