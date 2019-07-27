package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	unirq "github.com/redochen/demos/travelport-uapi/soap/universal/request"
	unirs "github.com/redochen/demos/travelport-uapi/soap/universal/response"
)

//RetrievePnrReqBody RetrievePnr请求的SOAP体
type RetrievePnrReqBody struct {
	Request *UniversalRecordRetrieveReq `xml:"univ:UniversalRecordRetrieveReq"`
}

//UniversalRecordRetrieveReq Request to retrieve a summary information for reservations under a Universal Record
type UniversalRecordRetrieveReq struct {
	soap.NameSpace
	comrq.BaseReq
	//<xs:choice>
	UniversalRecordLocatorCode com.TypeLocatorCode            `xml:"univ:UniversalRecordLocatorCode,omitempty"` //minOccurs="1" maxOccurs="1" //Represents a valid Universal Recordlocator code
	ProviderReservationInfo    *unirq.ProviderReservationInfo `xml:"univ:ProviderReservationInfo,omitempty"`    //
	//</xs:choice>
	ViewOnlyInd      bool                     `xml:"ViewOnlyInd,attr,omitempty"`      //use="optional" default="false" //True-Retrieves the PNR in UR Format, but doesn't create an actual UR in UAPI. False-Creates and Retrieves an actual UR. Default false.
	TravelerLastName com.TypeTravelerLastName `xml:"TravelerLastName,attr,omitempty"` //use="optional" //Match Traveler Last Name.
}

//RetrievePnrRspEnvelope RetrievePnr响应的SOAP信封
type RetrievePnrRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *RetrievePnrRspBody `xml:"Body"`
}

//RetrievePnrRspBody RetrievePnr响应的SOAP体
type RetrievePnrRspBody struct {
	Response *UniversalRecordRetrieveRsp `xml:"UniversalRecordRetrieveRsp"`
}

//UniversalRecordRetrieveRsp Return a Universal Record
type UniversalRecordRetrieveRsp struct {
	comrs.BaseRsp
	UniversalRecord *unirs.UniversalRecord `xml:"UniversalRecord"`
	Updated         bool                   `xml:"Updated,attr,omitempty"` //use="optional" default="false" //Returns true if the underlying reservation has changed since it was last accessed
}

//NewRetrievePnrReqBody 创建RetrievePnr请求体
func NewRetrievePnrReqBody(branchCode string) *RetrievePnrReqBody {
	body := &RetrievePnrReqBody{
		Request: &UniversalRecordRetrieveReq{},
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
