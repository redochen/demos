package lowfare

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
)

//LowFareParam LowFareSearch接口参数
type LowFareParam struct {
	models.BaseParam
	Routes         []*Route         `xml:"routes" json:"routes"`                                     //路线列表
	CabinClass     string           `xml:"cabin,omitempty" json:"cabin,omitempty"`                   //舱位等级：F-头等舱；C-商务舱；Y-经济舱。 默认为Y
	AdultCount     int              `xml:"adult,omitempty" json:"adult,omitempty"`                   //成人数量，默认为1
	ChildCount     int              `xml:"child,omitempty" json:"child,omitempty"`                   //儿童数量，默认为0
	Modifiers      *SearchModifiers `xml:"modifiers,omitempty" json:"modifiers,omitempty"`           //查询选项
	Currency       string           `xml:"currency,omitempty" json:"currency,omitempty"`             //货币代码
	MetaOption     string           `xml:"metaOption,omitempty" json:"metaOption,omitempty"`         //有效值为：00~99，或者D（默认值）。为空则表示不启用MetaSearch
	SaleCity       string           `xml:"saleCity,omitempty" json:"saleCity,omitempty"`             //销售城市（暂不起作用）
	BillingCity    string           `xml:"billingCity,omitempty" json:"billingCity,omitempty"`       //开票城市（暂不起作用）
	PlatingCarrier string           `xml:"platingCarrier,omitempty" json:"platingCarrier,omitempty"` //开票航司
	MaxSolutions   int              `xml:"maxSolutions,omitempty" json:"maxSolutions,omitempty"`     //返回多少条结果
}

//PreCheck 预检查
func (lowFareParam *LowFareParam) PreCheck() {
	lowFareParam.ServiceName = "LowFare"

	if len(lowFareParam.CabinClass) <= 0 {
		lowFareParam.CabinClass = "Y"
	}

	//清空日志上下文
	lowFareParam.LogContext = ""

	//if lowFareParam.LogContext == "" {
	if lowFareParam.Routes != nil && len(lowFareParam.Routes) > 0 {
		for _, route := range lowFareParam.Routes {
			if route != nil {
				lowFareParam.LogContext += route.GetLogContext()
			}
		}
	}

	lowFareParam.LogContext += fmt.Sprintf("[CoS-%s]",
		lowFareParam.CabinClass)

	lowFareParam.LogContext += fmt.Sprintf("[%d-%d]",
		lowFareParam.AdultCount,
		lowFareParam.ChildCount)

	if lowFareParam.Modifiers != nil {
		lowFareParam.LogContext += lowFareParam.Modifiers.GetLogContext()
	}

	if len(lowFareParam.Currency) > 0 {
		lowFareParam.LogContext += fmt.Sprintf("[%s]",
			lowFareParam.Currency)
	}

	if len(lowFareParam.PlatingCarrier) > 0 {
		lowFareParam.LogContext += fmt.Sprintf("[PC-%s]",
			lowFareParam.PlatingCarrier)
	}
	//}

	lowFareParam.BaseParam.PreCheck()
}
