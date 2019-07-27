package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
)

//HotelKeywordReqBody HotelKeyword请求的SOAP体
type HotelKeywordReqBody struct {
	Request *HotelKeywordReq `xml:"hot:HotelKeywordReq"`
}

//HotelKeywordReq Request to retrieve the hotel keyword details of a hotel chain or property
type HotelKeywordReq struct {
	soap.NameSpace
	comrq.BaseReq
	Keyword            []*comrq.Keyword          `xml:"com:Keyword,omitempty"`            //minOccurs="0" maxOccurs="15" //Used to request specific keyword details.
	PermittedProviders *comrq.PermittedProviders `xml:"com:PermittedProviders,omitempty"` //minOccurs="0"
	HotelChain         com.TypeHotelChainCode    `xml:"HotelChain,attr"`                  //use="required"
	HotelCode          com.TypeHotelCode         `xml:"HotelCode,attr,omitempty"`         //use="optional"
	CheckinDate        string                    `xml:"CheckinDate,attr,omitempty"`       //use="optional" //type="typeDate"
	ReturnKeywordList  bool                      `xml:"ReturnKeywordList,attr,omitempty"` //When true, a list of keyword names should be returned. If false then list of keyword details should be returned
}

//HotelKeywordRspEnvelope HotelKeyword响应的SOAP信封
type HotelKeywordRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *HotelKeywordRspBody `xml:"Body"`
}

//HotelKeywordRspBody HotelKeyword响应的SOAP体
type HotelKeywordRspBody struct {
	Response *HotelKeywordRsp `xml:"HotelKeywordRsp"`
}

//HotelKeywordRsp Response showing keyword details of a given hotel chain or property
type HotelKeywordRsp struct {
	comrs.BaseRsp
	MarketingInformation *comrs.MarketingInformation `xml:"MarketingInformation,omitempty"` //minOccurs="0" maxOccurs="1"
	Keyword              []*comrs.Keyword            `xml:"Keyword,omitempty"`              //minOccurs="0" maxOccurs="unbounded" //A word that a vendor uses to describe corporate policy/information.
}

//NewHotelKeywordReqBody 创建HotelKeyword请求体
func NewHotelKeywordReqBody(branchCode string) *HotelKeywordReqBody {
	body := &HotelKeywordReqBody{
		Request: &HotelKeywordReq{},
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
