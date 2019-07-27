package enum

type EnumReqSeat int

const (
	RsAny EnumReqSeat = iota
	RsAisle
	RsBulkhead
	RsExit
	RsWindow
	RsMiddle
)

//获取EnumReqSeat的字符串值
func (this EnumReqSeat) String() string {
	switch this {
	case RsAny:
		return "Any"
	case RsAisle:
		return "Aisle"
	case RsBulkhead:
		return "Bulkhead"
	case RsExit:
		return "Exit"
	case RsWindow:
		return "Window"
	case RsMiddle:
		return "Middle"
	}
	return ""
}

//获取EnumReqSeat的整数值
func (this EnumReqSeat) Value() int {
	if this >= RsAny && this <= RsMiddle {
		return int(this)
	} else {
		return -1
	}
}
