package airprice

import (
	"fmt"

	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	CcTime "github.com/redochen/tools/time"
)

//AirPriceParam AirPrice接口参数
type AirPriceParam struct {
	models.BaseParam
	Segments           []*Segment        `xml:"segments" json:"segments"`                         //航段列表
	Passengers         []*PassengerParam `xml:"passengers,omitempty" json:"passengers,omitempty"` //乘机人列表
	Carrier            string            `xml:"carrier,omitempty" json:"carrier,omitempty"`       //开票航司
	Currency           string            `xml:"currency,omitempty" json:"currency,omitempty"`     //货币代码
	CabinClass         string            `xml:"cabinClass,omitempty" json:"cabinClass,omitempty"` //舱位等级
	CheckInventory     bool              `xml:"checkInventory" json:"checkInventory"`             //是否检查库存：验舱和抢票=true，获取PNR价格=false
	SpecifyBookingCode bool              `xml:"specifyBookingCode" json:"specifyBookingCode"`     //是否指定舱位：获取PNR价格和验舱=true，抢票=false
}

//PassengerParam 乘机人参数类
type PassengerParam struct {
	PassengerType string `xml:"passengerType" json:"passengerType"` //类型
	ReferenceKey  string `xml:"referenceKey" json:"referenceKey"`   //预订引用KEY
}

//PreCheck 预检查
func (airPriceParam *AirPriceParam) PreCheck() {
	airPriceParam.ServiceName = "AirPrice"

	//if airPriceParam.LogContext == "" {
	if len(airPriceParam.CabinClass) <= 0 {
		airPriceParam.CabinClass = "Y"
	}

	//清空日志上下文
	airPriceParam.LogContext = ""

	if airPriceParam.Segments != nil && len(airPriceParam.Segments) > 0 {
		for _, segment := range airPriceParam.Segments {
			airPriceParam.LogContext += fmt.Sprintf("[%s%d-%s-%s-%s-%s]",
				segment.Carrier,
				segment.FlightNumber,
				segment.BookingCode,
				CcTime.GetShortDate(segment.DepartureTime),
				segment.DepartureAirport,
				segment.ArrivalAirport)
		}
	}

	if airPriceParam.CheckInventory {
		airPriceParam.LogContext += fmt.Sprintf("[ChkInv]")
	}

	if airPriceParam.SpecifyBookingCode {
		airPriceParam.LogContext += fmt.Sprintf("[SpcCode]")
	}

	airPriceParam.LogContext += fmt.Sprintf("[%s]", airPriceParam.CabinClass)
	airPriceParam.LogContext += fmt.Sprintf("[%s]", airPriceParam.Carrier)
	//}

	airPriceParam.BaseParam.PreCheck()
}
