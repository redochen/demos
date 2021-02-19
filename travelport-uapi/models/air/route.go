package air

import (
	"fmt"

	CcTime "github.com/redochen/tools/time"
)

//Route 路线参数
type Route struct {
	Origin        string `xml:"origin" json:"origin"`           //起始地（城市或机场）
	Destination   string `xml:"destination" json:"destination"` //目的地（城市或机场）
	DepartureDate string `xml:"depDate" json:"depDate"`         //出发日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	//ArrivalDate         string `xml:"arr,omitempty" json:"arr,omitempty"` //到达日期（可选），格式为“yyyy-MM-dd”或“yyyyMMdd”
}

//GetLogContext 获取日志上下文
func (route *Route) GetLogContext() string {
	logContext := fmt.Sprintf("[%s-%s-%s]",
		route.Origin,
		route.Destination,
		CcTime.RemoveDateSeparator(route.DepartureDate))

	return logContext
}
