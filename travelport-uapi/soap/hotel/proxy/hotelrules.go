package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
)

//HotelRulesReqBody HotelRules请求的SOAP体
type HotelRulesReqBody struct {
	Request *HotelRulesReq `xml:"hot:HotelRulesReq"`
}

//HotelRulesReq Retrieves hotel rules using hotel property code, chain code and hotel room rate type.
type HotelRulesReq struct {
	soap.NameSpace
	comrq.BaseReq
	//<xs:choice>
	HotelReservationLocatorCode com.TypeLocatorCode     `xml:"com:HotelReservationLocatorCode,omitempty"` //Request hotel rules using Locator code of existing hotel reservation.
	HotelRulesLookup            *hotrq.HotelRulesLookup `xml:"hot:HotelRulesLookup,omitempty"`            //Details to request Hotel rate rules post shopping request.
	//</xs:choice>
}

//HotelRulesRspEnvelope HotelRules响应的SOAP信封
type HotelRulesRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *HotelRulesRspBody `xml:"Body"`
}

//HotelRulesRspBody HotelRules响应的SOAP体
type HotelRulesRspBody struct {
	Response *HotelRulesRsp `xml:"HotelRulesRsp"`
}

//HotelRulesRsp Response showing rule details of a given hotel property and room rate code
type HotelRulesRsp struct {
	comrs.BaseRsp
	HotelRateDetail []*hotrs.HotelRateDetail `xml:"HotelRateDetail,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Returns hotels rate and rule details.
	HotelRuleItem   []*hotrs.HotelRuleItem   `xml:"HotelRuleItem,omitempty"`   //minOccurs="0" maxOccurs="unbounded" //Return rules and policies related to the property (Like Cancellation, Accepted FOP etc.).
	HotelType       *hotrs.HotelType         `xml:"HotelType,omitempty"`       //minOccurs="0" //Supported Providers:1G/1V/1P/1J.
}

//NewHotelRulesReqBody 创建HotelRules请求体
func NewHotelRulesReqBody(branchCode string) *HotelRulesReqBody {
	body := &HotelRulesReqBody{
		Request: &HotelRulesReq{},
	}

	body.Request.TargetBranch = branchCode
	body.Request.BillingPointOfSaleInfo = &comrq.BillingPointOfSaleInfo{
		OriginApplication: "UAPI",
	}

	//使用到了hot和com命名空间
	body.Request.SetHotNameSpace()
	body.Request.SetComNameSpace()

	return body
}
