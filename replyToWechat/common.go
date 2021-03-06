package replyToWechat

import "encoding/xml"

//Model is the common fields
type Model struct {
	XMLName        xml.Name `xml:"xml"`
	AppID          CString  `xml:"ToUserName"`
	UserOpenID     CString  `xml:"FromUserName"`
	CreateTimeUnix int64    `xml:"CreateTime"`
	MsgType        CString  `xml:"MsgType"`
}

//CString will be marshalled to CDATA
type CString string

//MarshalXML ...
func (c CString) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}
