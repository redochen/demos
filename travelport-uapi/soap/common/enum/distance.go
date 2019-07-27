package enum

type EnumDistance int

const (
	DistMI EnumDistance = iota
	DistKM
)

//获取EnumDistance的字符串值
func (this EnumDistance) String() string {
	switch this {
	case DistMI:
		return "MI"
	case DistKM:
		return "KM"
	}
	return ""
}

//获取EnumDistance的整数值
func (this EnumDistance) Value() int {
	if this >= DistMI && this <= DistKM {
		return int(this)
	} else {
		return -1
	}
}
