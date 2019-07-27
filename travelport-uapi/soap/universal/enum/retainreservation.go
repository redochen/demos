package enum

//Retain the Reservation (do not cancel) in the the event of a schedule or price change
type EnumRetainReservation int

const (
	RrNone EnumRetainReservation = iota
	RrSchedule
	RrPrice
	RrBoth
)

//获取EnumRetainReservation的字符串值
func (this EnumRetainReservation) String() string {
	switch this {
	case RrNone:
		return "None"
	case RrSchedule:
		return "Schedule"
	case RrPrice:
		return "Price"
	case RrBoth:
		return "Both"
	}
	return ""
}

//获取EnumRetainReservation的整数值
func (this EnumRetainReservation) Value() int {
	if this >= RrNone && this <= RrBoth {
		return int(this)
	} else {
		return -1
	}
}
