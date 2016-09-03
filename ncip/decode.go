package ncip

import (
	"encoding/xml"
	"io"
)

func DecodeRequest(r io.Reader) (req Request, err error) {
	dec := xml.NewDecoder(r)

	for {
		t, _ := dec.Token()
		if t == nil {
			break
		}
		switch elem := t.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "RequestItem":
				var reqItem RequestItem
				err = dec.DecodeElement(&reqItem, &elem)
				req = reqItem
			default:
				continue
			}
		}
	}

	return req, err
}
