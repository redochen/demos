package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
)

//HotelDetailsReqBody HotelDetails请求的SOAP体
type HotelDetailsReqBody struct {
	Request *HotelDetailsReq `xml:"hot:HotelDetailsReq"`
}

//HotelDetailsReq Request to retrieve the details of a hotel property
type HotelDetailsReq struct {
	soap.NameSpace
	hotrq.BaseHotelDetailsReq
	HostToken           *comrq.HostToken           `xml:"com:HostToken,omitempty"`           //minOccurs="0"
	NextResultReference *comrq.NextResultReference `xml:"com:NextResultReference,omitempty"` //minOccurs="0"
	ReturnMediaLinks    bool                       `xml:"ReturnMediaLinks,attr,omitempty"`   //use="optional" default="false" //If true, return the media links. Not supported by all providers
	ReturnGuestReviews  bool                       `xml:"ReturnGuestReviews,attr,omitempty"` //use="optional" default="false" //If true, return reviews and comments for the hotel property. Not supported by all providers
	PolicyReference     com.TypePolicyReference    `xml:"PolicyReference,attr,omitempty"`    //use="optional //This attribute will be used to pass in a value on the request which would be used to link to a ‘Policy Group’ in a policy engine external to UAPI.
}

//HotelDetailsRspEnvelope HotelDetails响应的SOAP信封
type HotelDetailsRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *HotelDetailsRspBody `xml:"Body"`
}

//HotelDetailsRspBody HotelDetails响应的SOAP体
type HotelDetailsRspBody struct {
	Response *HotelDetailsRsp `xml:"HotelDetailsRsp"`
}

//HotelDetailsRsp Response showing details of a given hotel property
type HotelDetailsRsp struct {
	comrs.BaseRsp
	NextResultReference *comrs.NextResultReference `xml:"NextResultReference,omitempty"` //minOccurs="0"
	HostToken           *comrs.HostToken           `xml:"HostToken,omitempty"`           //minOccurs="0"
	//<xs:choice>
	RequestedHotelDetails    *hotrs.RequestedHotelDetails    `xml:"RequestedHotelDetails,omitempty"`  //minOccurs="0" //Supported Provider GDS – 1G, 1V, 1P, 1J.
	AggregatorHotelDetails   []*hotrs.AggregatorHotelDetails `xml:"AggregatorHotelDetails,omitempty"` //minOccurs="0" maxOccurs="99" //Supported Provider TRM.
	HotelAlternateProperties *hotrs.HotelAlternateProperties `xml:"HotelAlternateProperties,omitempty"`
	//</xs:choice>
	GuestReviews *hotrs.GuestReviews `xml:"GuestReviews,omitempty"` //minOccurs="0"
}

//NewHotelDetailsReqBody 创建HotelDetails请求体
func NewHotelDetailsReqBody(branchCode string) *HotelDetailsReqBody {
	body := &HotelDetailsReqBody{
		Request: &HotelDetailsReq{},
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
