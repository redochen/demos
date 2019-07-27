package proxy

import (
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
	unirs "github.com/redochen/demos/travelport-uapi/soap/universal/response"
)

//HotelCreatePnrReqBody HotelCreatePnr请求的SOAP体
type HotelCreatePnrReqBody struct {
	Request *HotelCreateReservationReq `xml:"univ:HotelCreateReservationReq"`
}

//HotelCreateReservationReq Request to create a hotel reservation
type HotelCreateReservationReq struct {
	soap.NameSpace
	comrq.BaseCreateWithFormOfPaymentReq
	HotelRateDetail         []*hotrq.HotelRateDetail        `xml:"hot:HotelRateDetail"` //maxOccurs="99"
	HotelProperty           *hotrq.HotelProperty            `xml:"hot:HotelProperty"`
	ThirdPartyInformation   []*comrq.ThirdPartyInformation  `xml:"com:ThirdPartyInformation,omitempty"` //minOccurs="0"
	HotelStay               *hotrq.HotelStay                `xml:"hot:HotelStay"`
	Guarantee               *comrq.Guarantee                `xml:"com:Guarantee,omitempty"`               //minOccurs="0"
	HotelSpecialRequest     hot.HotelSpecialRequest         `xml:"hot:HotelSpecialRequest,omitempty"`     //minOccurs="0"
	PointOfSale             *comrq.PointOfSale              `xml:"com:PointOfSale,omitempty"`             //minOccurs="0"
	PromotionCode           *hotrq.PromotionCode            `xml:"hot:PromotionCode,omitempty"`           //minOccurs="0" //Used to specify promotional code include in the booking
	BookingSource           *comrq.BookingSource            `xml:"com:BookingSource,omitempty"`           //minOccurs="0" //Specify alternate booking source
	HotelBedding            []*hotrq.HotelBedding           `xml:"hot:HotelBedding,omitempty"`            //minOccurs="0" maxOccurs="4"
	GuestInformation        *hotrq.GuestInformation         `xml:"hot:GuestInformation,omitempty"`        //minOccurs="0"
	AssociatedRemark        []*hotrq.AssociatedRemark       `xml:"hot:AssociatedRemark,omitempty"`        //minOccurs="0" maxOccurs="9999"
	ReservationName         *comrq.ReservationName          `xml:"com:ReservationName,omitempty"`         //minOccurs="0" //If specified then it will be used for GDS reservation otherwise first booking traveler will be used.
	ActionStatus            *comrq.ActionStatus             `xml:"com:ActionStatus,omitempty"`            //minOccurs="0"
	HostToken               *comrq.HostToken                `xml:"com:HostToken,omitempty"`               //minOccurs="0"
	ReviewBooking           []*comrq.ReviewBooking          `xml:"com:ReviewBooking,omitempty"`           //minOccurs="0" maxOccurs="9999" //Review Booking or Queue Minders is to add the reminders in the Provider Reservation along with the date time and Queue details. On the date time defined in reminders, the message along with the PNR goes to the desired Queue.
	HotelCommission         hot.HotelCommission             `xml:"hot:HotelCommission,omitempty"`         //This element indicates hotel commission applied during hotel booking.  Provider supported 1P and 1J.
	BookingGuestInformation *hotrq.BookingGuestInformation  `xml:"hot:BookingGuestInformation,omitempty"` //minOccurs="0"
	UserAcceptance          bool                            `xml:"UserAcceptance,attr,omitempty"`         //use="optional" default="false" //If true, traveler has reviewed and accepted all policies, restrictions, and terms and conditions prior to booking. Default, false.
	MandatoryRateMatch      bool                            `xml:"MandatoryRateMatch,attr,omitempty"`     //use="optional" default="false" //If true, hotel will not be booked if there is a rate discrepancy.  Default is false. Supported providers: 1g,1v,1p,1j.
	StatusCode              com.TypeStatusCode              `xml:"StatusCode,attr,omitempty"`             //use="optional" //Hotel Segment Status Code.Supported Providers:1P/1J.
	BookingConfirmation     hot.TypeHotelConfirmationNumber `xml:"BookingConfirmation,attr,omitempty"`    //use="optional" //Hotel Booking Confirmation Number for passive hotel segment. Supported Providers:1P/1J.
}

//HotelCreatePnrRspEnvelope HotelCreatePnr响应的SOAP信封
type HotelCreatePnrRspEnvelope struct {
	soap.BaseRspEnvelope
	Body *HotelCreatePnrRspBody `xml:"Body"`
}

//HotelCreatePnrRspBody HotelCreatePnr响应的SOAP体
type HotelCreatePnrRspBody struct {
	Fault    *HotelCreatePnrRspFault    `xml:"Fault,omitempty"`
	Response *HotelCreateReservationRsp `xml:"HotelCreateReservationRsp,omitempty"`
}

//HotelCreatePnrRspFault 错误信息
type HotelCreatePnrRspFault struct {
	soap.Fault
	//Error *hotrs.AvailabilityErrorInfo `xml:"detail>AvailabilityErrorInfo,omitempty"`
}

//HotelCreateReservationRsp Provider: 1G,1V,1P,1J,ACH.
type HotelCreateReservationRsp struct {
	comrs.BaseRsp
	UniversalRecord      *unirs.UniversalRecord        `xml:"UniversalRecord,omitempty"`      //minOccurs="0"
	HotelRateChangedInfo []*hotrs.HotelRateChangedInfo `xml:"HotelRateChangedInfo,omitempty"` //Applicable for 1G, 1V, 1P, 1J
}

//NewHotelCreatePnrReqBody 创建HotelCreatePnr请求体
func NewHotelCreatePnrReqBody(branchCode string) *HotelCreatePnrReqBody {
	body := &HotelCreatePnrReqBody{
		Request: &HotelCreateReservationReq{},
	}

	body.Request.TargetBranch = branchCode
	body.Request.BillingPointOfSaleInfo = &comrq.BillingPointOfSaleInfo{
		OriginApplication: "UAPI",
	}

	//使用到了hot、com和univ命名空间
	body.Request.SetHotNameSpace()
	body.Request.SetComNameSpace()
	body.Request.SetUnivNameSpace()

	return body
}
