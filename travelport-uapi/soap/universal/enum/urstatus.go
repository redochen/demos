package enum

//Status of the UniversalRecord or the SavedTrip.
type EnumURStatus int

const (
	UrsActive EnumURStatus = iota
	UrsArchived
	UrsRetained
)

//获取EnumURStatus的字符串值
func (this EnumURStatus) String() string {
	switch this {
	case UrsActive:
		return "Active"
	case UrsArchived:
		return "Archived"
	case UrsRetained:
		return "Retained"
	}
	return ""
}

//获取EnumURStatus的整数值
func (this EnumURStatus) Value() int {
	if this >= UrsActive && this <= UrsRetained {
		return int(this)
	} else {
		return -1
	}
}
