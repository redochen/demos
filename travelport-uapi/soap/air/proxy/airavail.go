package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
)

//AirAvailReqBody AirAvail请求的SOAP体
type AirAvailReqBody struct {
	Request *AvailabilitySearchReq `xml:"air:AvailabilitySearchReq"`
}

//AvailabilitySearchReq Availability Search request.
type AvailabilitySearchReq struct {
	soap.NameSpace
	airrq.AirSearchReq
	SearchPassenger      []*comrq.SearchPassenger `xml:"com:SearchPassenger,omitempty"`       //minOccurs="0" maxOccurs="18" //Provider: 1G,1V,1P,1J,ACH-Maxinumber of passenger increased in to 18 to support 9 INF passenger along with 9 ADT,CHD,INS passenger
	PointOfSale          []*comrq.PointOfSale     `xml:"com:PointOfSale,omitempty"`           //minOccurs="0" maxOccurs="5" //Provider: ACH.
	ReturnBrandIndicator bool                     `xml:"ReturnBrandIndicator,attr,omitempty"` //use="optional" default="false" //When set to “true”, the Brand Indicator can be returned in the availability search response. Provider: 1G, 1V, 1P, 1J, ACH
	ChannelID            string                   `xml:"ChannelId,attr,omitempty"`            //use="optional" //A Channel ID is 4 alpha-numeric characters used to activate the Search Control Console filter for a specific group of travelers being served by the agency credential.
	NSCC                 string                   `xml:"NSCC,attr,omitempty"`                 //use="optional" //Allows the agency to bypass/override the Search Control Console rule.
}

//AirAvailRspEnvelope AirAvail响应的SOAP信封
type AirAvailRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *AirAvailRspBody `xml:"Body"`
}

//AirAvailRspBody AirAvail响应的SOAP体
type AirAvailRspBody struct {
	Response *AirAvailRsp `xml:"AvailabilitySearchRsp"`
}

//AirAvailRsp Provider: 1G,1V,1P,1J,ACH.
type AirAvailRsp struct {
	airrs.BaseAvailabilitySearchRsp
}

//NewAirAvailReqBody 创建AirAvail请求体
func NewAirAvailReqBody(branchCode string) *AirAvailReqBody {
	body := &AirAvailReqBody{
		Request: &AvailabilitySearchReq{},
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
