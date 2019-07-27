package enum

type EnumUnits int

const (
	MI EnumUnits = iota
	KM
)

//获取EnumUnits的字符串值
func (this EnumUnits) String() string {
	switch this {
	case MI:
		return "MI"
	case KM:
		return "KM"
	}
	return ""
}

//获取EnumUnits的整数值
func (this EnumUnits) Value() int {
	if this >= MI && this <= KM {
		return int(this)
	} else {
		return -1
	}
}
