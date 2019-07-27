package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
)

//HotelAvailReqBody HotelAvail请求的SOAP体
type HotelAvailReqBody struct {
	Request *HotelSearchAvailabilityReq `xml:"hot:HotelSearchAvailabilityReq"`
}

//HotelSearchAvailabilityReq Request to search for hotel availability.
type HotelSearchAvailabilityReq struct {
	soap.NameSpace
	hotrq.BaseHotelSearchReq
}

//HotelAvailRspEnvelope HotelAvail响应的SOAP信封
type HotelAvailRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *HotelAvailRspBody `xml:"Body"`
}

//HotelAvailRspBody HotelAvail响应的SOAP体
type HotelAvailRspBody struct {
	Response *HotelSearchAvailabilityRsp `xml:"HotelSearchAvailabilityRsp"`
}

//HotelSearchAvailabilityRsp Hotel availablity search response.
type HotelSearchAvailabilityRsp struct {
	hotrs.BaseHotelSearchRsp
}

//NewHotelAvailReqBody 创建HotelAvail请求的SOAP信封
func NewHotelAvailReqBody(branchCode string) *HotelAvailReqBody {
	body := &HotelAvailReqBody{
		Request: &HotelSearchAvailabilityReq{},
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
