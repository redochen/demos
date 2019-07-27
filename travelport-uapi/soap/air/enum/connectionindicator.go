package enum

//Types of connection indicator
type EnumConnectionIndicator int

const (
	CiAvailabilityAndPricing EnumConnectionIndicator = iota //Specified availability and pricing connection
	CiTurnAround                                            //Specified turn around
	CiStopover                                              //Specified stopover
)

//获取EnumConnectionIndicator的字符串值
func (this EnumConnectionIndicator) String() string {
	switch this {
	case CiAvailabilityAndPricing:
		return "AvailabilityAndPricing"
	case CiTurnAround:
		return "TurnAround"
	case CiStopover:
		return "Stopover"

	}
	return ""
}

//获取EnumConnectionIndicator的整数值
func (this EnumConnectionIndicator) Value() int {
	if this >= CiAvailabilityAndPricing && this <= CiStopover {
		return int(this)
	} else {
		return -1
	}
}
