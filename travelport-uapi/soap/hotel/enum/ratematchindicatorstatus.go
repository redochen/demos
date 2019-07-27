package enum

type EnumRateMatchIndicatorStatus int

const (
	RmisAvailable         EnumRateMatchIndicatorStatus = iota //Number of items requested for the IndicatorType is available
	RmisNotAvailable                                          //Number of items requested for the IndicatorType is not available. The actual number available is provided against Value
	RmisSubstituteOffered                                     //A substitute has been offered for the originally requested number and/or type. The substituted available is provided in against Value
	RmisMaximumExceeded                                       //Number of items requested for the IndicatorType exceeds the maximum applicable value. The substituted available is provided in against Value
)

//获取EnumRateMatchIndicatorStatus的字符串值
func (this EnumRateMatchIndicatorStatus) String() string {
	switch this {
	case RmisAvailable:
		return "Available"
	case RmisNotAvailable:
		return "NotAvailable"
	case RmisSubstituteOffered:
		return "SubstituteOffered"
	case RmisMaximumExceeded:
		return "MaximumExceeded"
	}
	return ""
}

//获取EnumRateMatchIndicatorStatus的整数值
func (this EnumRateMatchIndicatorStatus) Value() int {
	if this >= RmisAvailable && this <= RmisMaximumExceeded {
		return int(this)
	} else {
		return -1
	}
}
