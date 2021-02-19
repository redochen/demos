package createpnr

import (
	"fmt"

	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	CcTime "github.com/redochen/tools/time"
)

//CreatePnrParam 创建PNR参数
type CreatePnrParam struct {
	models.BaseParam
	Travelers         []*Traveler         `xml:"travelers" json:"travelers"`                                   //旅客列表
	OtherServiceInfos []*OtherServiceInfo `xml:"otherServiceInfo,omitempty" json:"otherServiceInfo,omitempty"` //其他服务信息列表
	Air               *AirParam           `xml:"air,omitempty" json:"air,omitempty"`                           //机票相关
	Hotel             *HotelParam         `xml:"hotel,omitempty" json:"hotel,omitempty"`                       //酒店相关
	Payments          []*FormOfPayment    `xml:"payments,omitempty" json:"payments,omitempty"`                 //支付类型列表
}

//AirParam 机票参数类
type AirParam struct {
	Segments         []*Segment `xml:"segments" json:"segments"`                 //航段列表
	TicketDate       string     `xml:"ticketDate" json:"ticketDate"`             //开票日期，格式为：yyyy-MM-ddTHH:mm:ss
	TicketingCarrier string     `xml:"ticketingCarrier" json:"ticketingCarrier"` //开票航司
}

//HotelParam 酒店参数类
type HotelParam struct {
	Property     *HotelProperty  `xml:"property" json:"property"`                     //酒店物业
	Rates        []*HotelRate    `xml:"rates" json:"rates"`                           //价格列表
	Rooms        int             `xml:"rooms,omitempty" json:"rooms,omitempty"`       //房间数量，默认为1
	Adults       int             `xml:"adults,omitempty" json:"adults,omitempty"`     //成人数量，默认为1
	Children     int             `xml:"children,omitempty" json:"children,omitempty"` //儿童数量，默认为0
	CheckinDate  string          `xml:"checkin,omitempty" json:"checkin,omitempty"`   //入住日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	CheckoutDate string          `xml:"checkout,omitempty" json:"checkout,omitempty"` //退房日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	HostToken    *HostTokenParam `xml:"hostToken,omitempty" json:"hostToken,omitempty"`
}

//HostTokenParam 主机口令。TRM使用，由Hotel Search接口返回，有效期为15分钟
type HostTokenParam struct {
	Host  string `xml:"host" json:"host"`   //主机
	Key   string `xml:"key" json:"key"`     //键
	Token string `xml:"token" json:"token"` //口令
}

//PreCheck 预检查
func (createPnrParam *CreatePnrParam) PreCheck() {
	createPnrParam.ServiceName = "CreatePnr"

	//if createPnrParam.LogContext == "" {
	passengerNum := 0
	if createPnrParam.Travelers != nil {
		passengerNum = len(createPnrParam.Travelers)
	}

	//清空日志上下文
	createPnrParam.LogContext = ""

	if createPnrParam.Air != nil {
		if createPnrParam.Air.Segments != nil && len(createPnrParam.Air.Segments) > 0 {
			for _, segment := range createPnrParam.Air.Segments {
				createPnrParam.LogContext += fmt.Sprintf("[%s%d-%s-%s-%s-%s][AS-%s]",
					segment.Carrier,
					segment.FlightNumber,
					segment.BookingCode,
					CcTime.GetShortDate(segment.DepartureTime),
					segment.DepartureAirport,
					segment.ArrivalAirport,
					segment.AvailabilitySource)
			}
		}

		createPnrParam.LogContext += fmt.Sprintf("[P-%d]", passengerNum)
	}

	if createPnrParam.Hotel != nil {
		if createPnrParam.Hotel.Property != nil {
			createPnrParam.LogContext += fmt.Sprintf("[%s-%s-%s]",
				createPnrParam.Hotel.Property.Chain,
				createPnrParam.Hotel.Property.Code,
				createPnrParam.Hotel.Property.Location)
		}
	}
	//}

	createPnrParam.BaseParam.PreCheck()
}
