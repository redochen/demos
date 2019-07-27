package airavail

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
)

//AirAvailParam AirAvail接口参数
type AirAvailParam struct {
	models.BaseParam
	Routes      []*Route         `xml:"routes,omitempty" json:"routes,omitempty"`           //路线列表（查询功能）
	Segments    []*Segment       `xml:"segments,omitempty" json:"segments,omitempty"`       //航段列表（验舱功能）
	Modifiers   *SearchModifiers `xml:"modifiers,omitempty" json:"modifiers,omitempty"`     //查询选项
	MaxPageDown int              `xml:"maxPageDown,omitempty" json:"maxPageDown,omitempty"` //最多翻页次数，默认为0，即不翻页
}

//PreCheck 预检查
func (airAvailParam *AirAvailParam) PreCheck() {
	airAvailParam.ServiceName = "AirAvail"

	//清空日志上下文
	airAvailParam.LogContext = ""

	//if this.LogContext == "" {

	if airAvailParam.Segments != nil && len(airAvailParam.Segments) > 0 {
		for _, segment := range airAvailParam.Segments {
			if segment != nil {
				airAvailParam.LogContext += segment.GetLogContext()
			}
		}
	} else if airAvailParam.Routes != nil && len(airAvailParam.Routes) > 0 {
		for _, route := range airAvailParam.Routes {
			if route != nil {
				airAvailParam.LogContext += route.GetLogContext()
			}
		}
	}

	if airAvailParam.Modifiers != nil {
		airAvailParam.LogContext += airAvailParam.Modifiers.GetLogContext()
	}

	if airAvailParam.MaxPageDown > 0 {
		airAvailParam.LogContext += fmt.Sprintf("[PN-%d]",
			airAvailParam.MaxPageDown)
	}
	//}

	airAvailParam.BaseParam.PreCheck()
}
