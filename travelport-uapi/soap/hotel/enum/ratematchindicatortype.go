package enum

//"Match" Indicators for certain request parameters, e.g. Child Count, Bed Type, Extra Adults etc.
type EnumRateMatchIndicatorType int

const (
	RmitRateCategory EnumRateMatchIndicatorType = iota
	RmitRoomCount
	RmitAdultCount
	RmitChildCount
	RmitAdultRollaway
	RmitChildRollaway
	RmitCrib
)

//获取EnumRateMatchIndicatorType的字符串值
func (this EnumRateMatchIndicatorType) String() string {
	switch this {
	case RmitRateCategory:
		return "RateCategory"
	case RmitRoomCount:
		return "RoomCount"
	case RmitAdultCount:
		return "AdultCount"
	case RmitChildCount:
		return "ChildCount"
	case RmitAdultRollaway:
		return "AdultRollaway"
	case RmitChildRollaway:
		return "ChildRollaway"
	case RmitCrib:
		return "Crib"
	}
	return ""
}

//获取EnumRateMatchIndicatorType的整数值
func (this EnumRateMatchIndicatorType) Value() int {
	if this >= RmitRateCategory && this <= RmitCrib {
		return int(this)
	} else {
		return -1
	}
}
