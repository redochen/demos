package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	unirs "github.com/redochen/demos/travelport-uapi/soap/universal/response"
)

//AirCreatePnrReqBody AirCreatePnr请求的SOAP体
type AirCreatePnrReqBody struct {
	Request *AirCreateReservationReq `xml:"univ:AirCreateReservationReq"`
}

//AirCreateReservationReq Request to store an air itinerary and create initial PNR.
type AirCreateReservationReq struct {
	soap.NameSpace
	comrq.BaseCreateWithFormOfPaymentReq
	SupplierLocator       []*comrq.SupplierLocator       `xml:"com:SupplierLocator,omitempty"`       //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH,SDK.
	ThirdPartyInformation []*comrq.ThirdPartyInformation `xml:"com:ThirdPartyInformation,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Provider: SDK.
	PointOfSale           *comrq.PointOfSale             `xml:"com:PointOfSale,omitempty"`           //minOccurs="0" //Provider: 1G,1V.
	AirPricingSolution    *airrq.AirPricingSolution      `xml:"air:AirPricingSolution"`              //Provider: 1G,1V,1P,1J,ACH,SDK.
	ActionStatus          []*comrq.ActionStatus          `xml:"com:ActionStatus,omitempty"`          //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH,SDK.
	Payment               []*comrq.Payment               `xml:"com:Payment,omitempty"`               //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH.
	//DeliveryInfo *comrq.DeliveryInfo `xml:"com:DeliveryInfo,omitempty"` //minOccurs="0" //Provider: ACH.
	AutoSeatAssignment           []*airrq.AutoSeatAssignment           `xml:"air:AutoSeatAssignment,omitempty"`           //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J.
	SpecificSeatAssignment       []*airrq.SpecificSeatAssignment       `xml:"air:SpecificSeatAssignment,omitempty"`       //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH.
	AssociatedRemark             []*airrq.AssociatedRemark             `xml:"air:AssociatedRemark,omitempty"`             //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH,SDK.
	PocketItineraryRemark        []*airrq.PocketItineraryRemark        `xml:"air:PocketItineraryRemark,omitempty"`        //minOccurs="0" maxOccurs="unbounded" //Provider: 1P,1J.
	ReviewBooking                []*comrq.ReviewBooking                `xml:"com:ReviewBooking,omitempty"`                //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V-Review Booking or Queue Minders is to add the reminders in the Provider Reservation along with the date time and Queue details. On the date time defined in reminders, the message along with the PNR goes to the desired Queue.
	AirPricingTicketingModifiers []*airrq.AirPricingTicketingModifiers `xml:"air:AirPricingTicketingModifiers,omitempty"` //minOccurs="0" maxOccurs="99" //AirPricing TicketingModifier information used to associate Ticketing Modifiers with one or more AirPricingInfos/ProviderReservationInfo for 1G,1V,1P,1J
	RetainReservation            string                                `xml:"RetainReservation,attr,omitempty"`           //use="optional" default="None" //Provider: 1G,1V,1P, 1J, ACH.
	Source                       string                                `xml:"Source,attr,omitempty"`                      //use="optional"
	OverrideMCT                  bool                                  `xml:"OverrideMCT,attr,omitempty"`                 //use="optional" default="false" //Provider: 1G,1V.
	RestrictWaitlist             bool                                  `xml:"RestrictWaitlist,attr,omitempty"`            //use="optional" default="false" //Restrict Book if it sells a Waitlisted AirSegment. Provider: 1G,1V
	CreatePassiveForHold         bool                                  `xml:"CreatePassiveForHold,attr,omitempty"`        //use="optional" default="false" //Creates a background passive segment for an ACH hold booking.
	ChannelID                    string                                `xml:"ChannelId,attr,omitempty"`                   //use="optional" //A Channel ID is 4 alpha-numeric characters used to activate the Search Control Console filter for a specific group of travelers being served by the agency credential.
	NSCC                         string                                `xml:"NSCC,attr,omitempty"`                        //use="optional" //Allows the agency to bypass/override the Search Control Console rule.
}

//AirCreatePnrRspEnvelope AirCreatePnr响应的SOAP信封
type AirCreatePnrRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *AirCreatePnrRspBody `xml:"Body"`
}

//AirCreatePnrRspBody AirCreatePnr响应的SOAP体
type AirCreatePnrRspBody struct {
	Fault    *AirCreatePnrRspFault    `xml:"Fault,omitempty"`
	Response *AirCreateReservationRsp `xml:"AirCreateReservationRsp,omitempty"`
}

//AirCreatePnrRspFault 错误信息
type AirCreatePnrRspFault struct {
	soap.Fault
	Error *airrs.AvailabilityErrorInfo `xml:"detail>AvailabilityErrorInfo,omitempty"`
}

//AirCreateReservationRsp Provider: 1G,1V,1P,1J,ACH.
type AirCreateReservationRsp struct {
	comrs.BaseRsp
	UniversalRecord *unirs.UniversalRecord `xml:"UniversalRecord,omitempty"` //minOccurs="0"
	//AirSolutionChangedInfo    []*airrs.AirSolutionChangedInfo  `xml:"AirSolutionChangedInfo,omitempty"`    //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH.
	AirSegmentSellFailureInfo *airrs.AirSegmentSellFailureInfo `xml:"AirSegmentSellFailureInfo,omitempty"` //minOccurs="0" //Provider: 1G,1V,1P,1J,ACH.
}

//NewAirCreatePnrReqBody 创建AirCreatePnr请求体
func NewAirCreatePnrReqBody(branchCode string) *AirCreatePnrReqBody {
	body := &AirCreatePnrReqBody{
		Request: &AirCreateReservationReq{},
	}

	body.Request.TargetBranch = branchCode
	body.Request.BillingPointOfSaleInfo = &comrq.BillingPointOfSaleInfo{
		OriginApplication: "UAPI",
	}

	//使用到了air、com和univ命名空间
	body.Request.SetAirNameSpace()
	body.Request.SetComNameSpace()
	body.Request.SetUnivNameSpace()

	return body
}
