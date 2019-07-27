package system

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
)

//PingReqBody Ping请求的SOAP体
type PingReqBody struct {
	Request *PingReq `xml:"sys:PingReq"`
}

//PingReq A simple request to test connectivity to the system without imposing any actions
type PingReq struct {
	soap.NameSpace
	comrq.BaseReq
	Payload Payload `xml:"sys:Payload,omitempty"` //minOccurs="0"
}

//PingRspEnvelope Ping响应的SOAP信封
type PingRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *PingRspBody `xml:"Body"`
}

//PingRspBody ...
type PingRspBody struct {
	PingRsp *PingRsp `xml:"PingRsp"`
}

//PingRsp Response to the PingReq. Will contain the exact payload (if any) that was sent in
type PingRsp struct {
	comrs.BaseRsp
	Payload Payload `xml:"Payload,omitempty"` //minOccurs="0"
}

//NewPingReqBody 创建Ping请求体
func NewPingReqBody(branchCode string, payload string) *PingReqBody {
	body := &PingReqBody{
		Request: &PingReq{
			Payload: Payload(payload),
		},
	}

	body.Request.TargetBranch = branchCode
	body.Request.BillingPointOfSaleInfo = &comrq.BillingPointOfSaleInfo{
		OriginApplication: "UAPI",
	}

	//使用到了sys和com命名空间
	body.Request.SetSysNameSpace()
	body.Request.SetComNameSpace()

	return body
}
