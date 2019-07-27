package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
)

//HotelMediaReqBody HotelMedia请求的SOAP体
type HotelMediaReqBody struct {
	Request *HotelMediaLinksReq `xml:"hot:HotelMediaLinksReq"`
}

//HotelMediaLinksReq Retrieves all image links from the Galileo Web Services Image Viewer eBL for up to 20 properties. Only the attributes of the HotelProperty are used in this request.
type HotelMediaLinksReq struct {
	soap.NameSpace
	comrq.BaseReq
	PermittedProviders *comrq.PermittedProviders `xml:"com:PermittedProviders,omitempty"` //minOccurs="0"
	HotelProperty      []*hotrq.HotelProperty    `xml:"hot:HotelProperty,omitempty"`      //maxOccurs="20"
	SecureLinks        bool                      `xml:"SecureLinks,attr,omitempty"`       //use="optional" default="true" //Request URLs returned use secured site (https) references. Default is true
	SizeCode           string                    `xml:"SizeCode,attr,omitempty"`          //use="optional" default="A" //Requested image size. Default is to get ALL images
	RichMedia          bool                      `xml:"RichMedia,attr,omitempty"`         //use="optional" default="true" //Request the Rich Media link. Default is true
	Gallery            bool                      `xml:"Gallery,attr,omitempty"`           //use="optional" default="true" //Request the Image Gallery link. Default is true
}

//HotelMediaRspEnvelope HotelMedia响应的SOAP信封
type HotelMediaRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *HotelMediaRspBody `xml:"Body"`
}

//HotelMediaRspBody HotelMedia响应的SOAP体
type HotelMediaRspBody struct {
	Response *HotelMediaLinksRsp `xml:"HotelMediaLinksRsp"`
}

//HotelMediaLinksRsp ...
type HotelMediaLinksRsp struct {
	comrs.BaseRsp
	HotelPropertyWithMediaItems []*hotrs.HotelPropertyWithMediaItems `xml:"HotelPropertyWithMediaItems,omitempty"` //maxOccurs="20"
}

//NewHotelMediaReqBody 创建HotelMedia请求体
func NewHotelMediaReqBody(branchCode string) *HotelMediaReqBody {
	body := &HotelMediaReqBody{
		Request: &HotelMediaLinksReq{},
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
