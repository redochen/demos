package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
)

//AirPriceReqBody AirPrice请求的SOAP体
type AirPriceReqBody struct {
	Request *AirPriceReq `xml:"air:AirPriceReq"`
}

//AirPriceReq Request to price an itinerary in one to many ways. Pricing commands can be specified globally, or specifically per command.
type AirPriceReq struct {
	soap.NameSpace
	airrq.BaseAirPriceReq
}

//AirPriceRspEnvelope AirPrice响应的SOAP信封
type AirPriceRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *AirPriceRspBody `xml:"Body"`
}

//AirPriceRspBody AirPrice响应的SOAP体
type AirPriceRspBody struct {
	Fault    *AirPriceRspFault `xml:"Fault,omitempty"`
	Response *AirPriceRsp      `xml:"AirPriceRsp"`
}

//AirPriceRspFault 错误信息
type AirPriceRspFault struct {
	soap.Fault
	//Error *airrs.AvailabilityErrorInfo `xml:"detail>AvailabilityErrorInfo,omitempty"`
}

//AirPriceRsp 响应
type AirPriceRsp struct {
	BaseAirPriceRsp
}

//BaseAirPriceRsp ...
type BaseAirPriceRsp struct {
	comrs.BaseRsp
	AirItinerary   *airrs.AirItinerary     `xml:"AirItinerary"`   //Provider: 1G,1V,1P,1J,ACH.
	AirPriceResult []*airrs.AirPriceResult `xml:"AirPriceResult"` //maxOccurs="16" //Provider: 1G,1V,1P,1J,ACH.
}

//NewAirPriceReqBody 创建AirPrice请求体
func NewAirPriceReqBody(branchCode string) *AirPriceReqBody {
	body := &AirPriceReqBody{
		Request: &AirPriceReq{},
	}

	body.Request.TargetBranch = branchCode
	body.Request.BillingPointOfSaleInfo = &comrq.BillingPointOfSaleInfo{
		OriginApplication: "UAPI",
	}

	//使用到了air和com命名空间
	body.Request.SetAirNameSpace()
	body.Request.SetComNameSpace()

	return body
}
