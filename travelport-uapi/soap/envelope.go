package soap

import (
	"encoding/xml"
)

const (
	//NSSoapEnv Soap命名空间
	NSSoapEnv string = "http://schemas.xmlsoap.org/soap/envelope/"
)

//ReqEnvelope 请求SOAP信封
type ReqEnvelope struct {
	XMLName   xml.Name    `xml:"SOAP:Envelope"`
	Namespace string      `xml:"xmlns:SOAP,attr,omitempty"`
	Header    *Header     `xml:"SOAP:Header,omitempty"`
	Body      interface{} `xml:"SOAP:Body"`
}

//BaseRspEnvelope 响应SOAP信封基类
type BaseRspEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  *Header  `xml:"Header,omitempty"`
}

//Header SOAP头
type Header struct {
}

//Fault SOAP错误信息
type Fault struct {
	Code string `xml:"faultcode,omitempty"`
	Text string `xml:"faultstring,omitempty"`
}

//NewReqEnvelope 创建新的请求实例
func NewReqEnvelope(body interface{}) *ReqEnvelope {
	envelope := &ReqEnvelope{
		Namespace: NSSoapEnv,
		Header:    &Header{},
		Body:      body,
	}

	return envelope
}
