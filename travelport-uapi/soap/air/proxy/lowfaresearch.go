package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
)

//LowFareReqBody LowFare请求的SOAP体
type LowFareReqBody struct {
	Request *LowFareSearchReq `xml:"air:LowFareSearchReq"`
}

//LowFareSearchReq Low Fare Search request.Provider: 1G,1V,1P,1J,ACH.
type LowFareSearchReq struct {
	soap.NameSpace
	airrq.BaseLowFareSearchReq
	PolicyReference com.TypePolicyReference `xml:"PolicyReference,attr,omitempty"` //use="optional" //This attribute will be used to pass in a value on the request which would be used to link to a ‘Policy Group’ in a policy engine external to UAPI.
}

//LowFareRspEnvelope LowFareSearch响应的SOAP信封
type LowFareRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *LowFareRspBody `xml:"Body"`
}

//LowFareRspBody LowFareSearch响应的SOAP体
type LowFareRspBody struct {
	Response *LowFareSearchRsp `xml:"LowFareSearchRsp"`
}

//LowFareSearchRsp Low Fare Search Response
type LowFareSearchRsp struct {
	airrs.AirSearchRsp
	BrandList    []*airrs.BrandList `xml:"BrandList,omitempty"`         //minOccurs="0"
	CurrencyType com.TypeCurrency   `xml:"CurrencyType,attr,omitempty"` //use="required" //Provider: 1G,1V,1P,1J,ACH.
}

//NewLowFareReqBody 创建LowFare请求体
func NewLowFareReqBody(branchCode string) *LowFareReqBody {
	body := &LowFareReqBody{
		Request: &LowFareSearchReq{},
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
