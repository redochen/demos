package enum

//The valid rule types
type EnumFareRuleType int

const (
	FrtNone EnumFareRuleType = iota
	FrtShort
	FrtLong
)

//获取EnumFareRuleType的字符串值
func (this EnumFareRuleType) String() string {
	switch this {
	case FrtNone:
		return "none"
	case FrtShort:
		return "short"
	case FrtLong:
		return "long"
	}
	return ""
}

//获取EnumFareRuleType的整数值
func (this EnumFareRuleType) Value() int {
	if this >= FrtNone && this <= FrtLong {
		return int(this)
	} else {
		return -1
	}
}
