package air

import "fmt"

//SearchModifiers 查询选项
type SearchModifiers struct {
	NumOfStops           int      `xml:"stop,omitempty" json:"stop,omitempty"`                                 //中转次数，-1表示不限制，默认为0（仅限直飞）
	IncludeAirlines      []string `xml:"includeAirlines,omitempty" json:"includeAirlines,omitempty"`           //包含的航司列表
	ExcludeAirlines      []string `xml:"excludeAirlines,omitempty" json:"excludeAirlines,omitempty"`           //排除的航司列表
	SameAirline          int      `xml:"sameAirline,omitempty" json:"sameAirline,omitempty"`                   //是否同航司：1 - 同航司；2 - 不同航司；其他 - 未指定
	MaxConnectionMintues int      `xml:"maxConnectionMintues,omitempty" json:"maxConnectionMintues,omitempty"` //最大转机时间（单位：分钟）
	MaxJourneyHours      int      `xml:"maxJourneyHours,omitempty" json:"maxJourneyHours,omitempty"`           //最大旅行时间（单位：小时）
}

//GetSameAirline 获取请求参数
func (modifiers *SearchModifiers) GetSameAirline() string {
	var sameAirline string

	switch modifiers.NumOfStops {
	case 1:
		if modifiers.SameAirline == 1 {
			sameAirline = "SO"
		} else if modifiers.SameAirline == 2 {
			sameAirline = "SI"
		}
		break
	case 2:
		if modifiers.SameAirline == 1 {
			sameAirline = "DO"
		} else if modifiers.SameAirline == 2 {
			sameAirline = "DI"
		}
		break
	case 3:
		if modifiers.SameAirline == 1 {
			sameAirline = "TO"
		} else if modifiers.SameAirline == 2 {
			sameAirline = "TI"
		}
	}

	return sameAirline
}

//GetLogContext 获取日志上下文
func (modifiers *SearchModifiers) GetLogContext() string {
	var logContext string

	logContext = fmt.Sprintf("[NoS-%d]", modifiers.NumOfStops)

	sameAirline := modifiers.GetSameAirline()
	if len(sameAirline) > 0 {
		logContext = fmt.Sprintf("[%s]", modifiers.GetSameAirline())
	}

	var includeAirlines string
	if modifiers.IncludeAirlines != nil && len(modifiers.IncludeAirlines) > 0 {
		for _, inc := range modifiers.IncludeAirlines {
			includeAirlines += inc
		}
	}

	if len(includeAirlines) > 0 {
		logContext += fmt.Sprintf("[%s]", includeAirlines)
	}

	var excludeAirlines string
	if modifiers.ExcludeAirlines != nil && len(modifiers.ExcludeAirlines) > 0 {
		for _, inc := range modifiers.ExcludeAirlines {
			excludeAirlines += inc
		}
	}

	if len(excludeAirlines) > 0 {
		logContext += fmt.Sprintf("[~%s]", excludeAirlines)
	}

	if modifiers.MaxConnectionMintues > 0 {
		logContext += fmt.Sprintf("[MCM-%d]", modifiers.MaxConnectionMintues)
	}

	if modifiers.MaxJourneyHours > 0 {
		logContext += fmt.Sprintf("[MJH-%d]", modifiers.MaxJourneyHours)
	}

	return logContext
}
