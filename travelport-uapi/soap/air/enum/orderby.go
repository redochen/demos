package enum

type EnumOrderBy int

const (
	ObJourneyTime EnumOrderBy = iota
	ObDepartureTime
	ObArrivalTime
)

//获取EnumOrderBy的字符串值
func (this EnumOrderBy) String() string {
	switch this {
	case ObJourneyTime:
		return "JourneyTime"
	case ObDepartureTime:
		return "DepartureTime"
	case ObArrivalTime:
		return "ArrivalTime"
	}
	return ""
}

//获取EnumOrderBy的整数值
func (this EnumOrderBy) Value() int {
	if this >= ObJourneyTime && this <= ObArrivalTime {
		return int(this)
	} else {
		return -1
	}
}
