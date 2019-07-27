package enum

//Types of possible commission.
type EnumCommissionType int

const (
	CtFlat EnumCommissionType = iota
	CtPercentBase
	CtPercentTotal
)

//获取EnumCommissionType的字符串值
func (this EnumCommissionType) String() string {
	switch this {
	case CtFlat:
		return "Flat"
	case CtPercentBase:
		return "PercentBase"
	case CtPercentTotal:
		return "PercentTotal"
	}
	return ""
}

//获取EnumCommissionType的整数值
func (this EnumCommissionType) Value() int {
	if this >= CtFlat && this <= CtPercentTotal {
		return int(this)
	} else {
		return -1
	}
}
