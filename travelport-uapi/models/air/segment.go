package air

import (
	"fmt"

	CcTime "github.com/redochen/tools/time"
)

//Segment 航段类
type Segment struct {
	Group                     int          `xml:"group" json:"group"`                                               //组号
	Index                     int          `xml:"index" json:"index"`                                               //航段索引号
	OptionKey                 string       `xml:"optionKey,omitempty" json:"optionKey,omitempty"`                   //Option键（一个行程的多个航段具体相同的OptionKey）
	SegmentKey                string       `xml:"segmentKey,omitempty" json:"segmentKey,omitempty"`                 //航段键（每个航段具有独特的SegmentKey）
	FlightDetailsKeys         []string     `xml:"flightDetailsKeys,omitempty" json:"flightDetailsKeys,omitempty"`   //航班详情键列表（经停的航段有多个FlightDetails）
	Carrier                   string       `xml:"carrier" json:"carrier"`                                           //销售航司
	FlightNumber              int          `xml:"flightNo" json:"flightNo"`                                         //航班号
	OperatingCarrier          string       `xml:"operatingCarrier,omitempty" json:"operatingCarrier,omitempty"`     //承运航司
	OperatingFlightNumber     int          `xml:"operatingFlightNo,omitempty" json:"operatingFlightNo,omitempty"`   //承运航班号
	DepartureAirport          string       `xml:"depAirport" json:"depAirport"`                                     //出发机场
	DepartureTerminal         string       `xml:"depTerminal" json:"depTerminal"`                                   //出发航站楼
	DepartureTime             string       `xml:"depTime" json:"depTime"`                                           //起飞时间，格式：yyyy-MM-ddTHH:mm:ss.fffzzz
	ArrivalAirport            string       `xml:"arrAirport" json:"arrAirport"`                                     //抵达机场
	ArrivalTerminal           string       `xml:"arrTerminal" json:"arrTerminal"`                                   //抵达航站楼
	ArrivalTime               string       `xml:"arrTime" json:"arrTime"`                                           //降落时间，格式：yyyy-MM-ddTHH:mm:ss.fffzzz
	FlightTime                int          `xml:"flightTime,omitempty" json:"flightTime,omitempty"`                 //飞行时间（分钟数）
	TravelTime                int          `xml:"travelTime,omitempty" json:"travelTime,omitempty"`                 //旅行时间（分钟数）
	Distance                  int          `xml:"distance,omitempty" json:"distance,omitempty"`                     //飞行距离
	Equipment                 string       `xml:"equipment" json:"equipment"`                                       //机型
	ETicket                   bool         `xml:"eTicket,omitempty" json:"eTicket,omitempty"`                       //是否电子客票
	BookingCode               string       `xml:"bookingCode,omitempty" json:"bookingCode,omitempty"`               //预订舱位
	BookingCount              string       `xml:"bookingCount,omitempty" json:"bookingCount,omitempty"`             //余座数
	CabinClass                string       `xml:"cabinClass,omitempty" json:"cabinClass,omitempty"`                 //舱位等级
	Stopovers                 []*Stopover  `xml:"stopovers,omitempty" json:"stopovers,omitempty"`                   //经停列表
	MarriageGroup             int          `xml:"marriageGroup,omitempty" json:"marriageGroup,omitempty"`           //联姻组号
	ChangeOfPlane             bool         `xml:"changePlane,omitempty" json:"changePlane,omitempty"`               //是否需要换机？？
	AvailInfos                []*AvailInfo `xml:"availInfos,omitempty" json:"availInfos,omitempty"`                 //余座信息
	Status                    string       `xml:"status,omitempty" json:"status,omitempty"`                         //状态
	TravelOrder               int          `xml:"travelOrder,omitempty" json:"travelOrder,omitempty"`               //航段序号
	ReservationInfoRef        string       `xml:"reservationInfoRef,omitempty" json:"reservationInfoRef,omitempty"` //预订信息索引
	Seamless                  bool         `xml:"seamless,omitempty" json:"seamless,omitempty"`                     //是否直连销售
	IsConnection              bool         `xml:"isConnection,omitempty" json:"isConnection,omitempty"`             //是否为Connection
	LinkAvailability          bool         `xml:"avLink,omitempty" json:"avLink,omitempty"`                         //是否链接AV
	AvailabilityDisplayType   string       `xml:"avDisplayType,omitempty" json:"avDisplayType,omitempty"`           //AV显示类型
	AvailabilitySource        string       `xml:"avSource,omitempty" json:"avSource,omitempty"`                     //AV来源
	OptionalServicesIndicator bool         `xml:"optionalSvcInd,omitempty" json:"optionalSvcInd,omitempty"`         //可选服务指示器
	ParticipantLevel          string       `xml:"participantLevel,omitempty" json:"participantLevel,omitempty"`     //参与级别
	PolledAvailabilityOption  string       `xml:"polledAvOption,omitempty" json:"polledAvOption,omitempty"`         //AV缓存选项
	ProviderCode              string       `xml:"providerCode,omitempty" json:"providerCode,omitempty"`             //供应商代码
}

//GetLogContext 获取日志上下文
func (segment *Segment) GetLogContext() string {
	logContext := fmt.Sprintf("[%s%d-%s-%s-%s]",
		segment.Carrier,
		segment.FlightNumber,
		segment.DepartureAirport,
		segment.ArrivalAirport,
		CcTime.RemoveDateSeparator(segment.DepartureTime))

	return logContext
}
