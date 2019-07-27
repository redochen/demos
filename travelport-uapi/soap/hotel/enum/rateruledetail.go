package enum

//Availability status of hotel
type EnumRateRuleDetail int

const (
	RrdNone          EnumRateRuleDetail = iota //'None' returns hotel property descriptive information-supported for 1p/1j,1g/1v.
	RrdComplete                                //'Complete' returns the complete hotel and room rate information-supported for 1p/1j,1g/1v
	RrdRatePlansOnly                           //'RatePlansOnly' returns hotel rate information only - supported for 1p/1j.

)

//获取EnumRateRuleDetail的字符串值
func (this EnumRateRuleDetail) String() string {
	switch this {
	case RrdNone:
		return "None"
	case RrdComplete:
		return "Complete"
	case RrdRatePlansOnly:
		return "RatePlansOnly"
	}
	return ""
}

//获取EnumRateRuleDetail的整数值
func (this EnumRateRuleDetail) Value() int {
	if this >= RrdNone && this <= RrdRatePlansOnly {
		return int(this)
	} else {
		return -1
	}
}
