package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	unirs "github.com/redochen/demos/travelport-uapi/soap/universal/response"
)

//CancelPnrReqBody CancelPnr请求的SOAP体
type CancelPnrReqBody struct {
	Request *UniversalRecordCancelReq `xml:"univ:UniversalRecordCancelReq"`
}

//UniversalRecordCancelReq Request to Cancel an Universal Record
type UniversalRecordCancelReq struct {
	soap.NameSpace
	comrq.BaseReq
	FileFinishingInfo          *comrq.FileFinishingInfo `xml:"com:FileFinishingInfo,omitempty"` //minOccurs="0"
	UniversalRecordLocatorCode com.TypeLocatorCode      `xml:"UniversalRecordLocatorCode,attr"` //use="required" //Represents a valid Universal Record locator code
	Version                    com.TypeURVersion        `xml:"Version,attr"`                    //use="required"
}

//CancelPnrRspEnvelope CancelPnr响应的SOAP信封
type CancelPnrRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *CancelPnrRspBody `xml:"Body"`
}

//CancelPnrRspBody CancelPnr响应的SOAP体
type CancelPnrRspBody struct {
	Response *UniversalRecordCancelRsp `xml:"UniversalRecordCancelRsp"`
}

//UniversalRecordCancelRsp Return status for each provider reservation
type UniversalRecordCancelRsp struct {
	comrs.BaseRsp
	ProviderReservationStatus []*unirs.ProviderReservationStatus `xml:"ProviderReservationStatus,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//NewCancelPnrReqBody 创建CancelPnr请求体
func NewCancelPnrReqBody(branchCode string) *CancelPnrReqBody {
	body := &CancelPnrReqBody{
		Request: &UniversalRecordCancelReq{},
	}

	body.Request.TargetBranch = branchCode
	body.Request.BillingPointOfSaleInfo = &comrq.BillingPointOfSaleInfo{
		OriginApplication: "UAPI",
	}

	//使用到了com和univ命名空间
	body.Request.SetComNameSpace()
	body.Request.SetUnivNameSpace()

	return body
}
