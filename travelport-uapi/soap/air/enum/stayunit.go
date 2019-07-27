package enum

//Units for the Length of Stay
type EnumStayUnit int

const (
	SuMinutes EnumStayUnit = iota
	SuHours
	SuDays
	SuMonths
	SuMonday
	SuTuesday
	SuWednesday
	SuThursday
	SuFriday
	SuSaturday
	SuSunday
)

//获取EnumStayUnit的字符串值
func (this EnumStayUnit) String() string {
	switch this {
	case SuMinutes:
		return "Minutes"
	case SuHours:
		return "Hours"
	case SuDays:
		return "Days"
	case SuMonths:
		return "Months"
	case SuMonday:
		return "Monday"
	case SuTuesday:
		return "Tuesday"
	case SuWednesday:
		return "Wednesday"
	case SuThursday:
		return "Thursday"
	case SuFriday:
		return "Friday"
	case SuSaturday:
		return "Saturday"
	case SuSunday:
		return "Sunday"
	}
	return ""
}

//获取EnumStayUnit的整数值
func (this EnumStayUnit) Value() int {
	if this >= SuMinutes && this <= SuSunday {
		return int(this)
	} else {
		return -1
	}
}
