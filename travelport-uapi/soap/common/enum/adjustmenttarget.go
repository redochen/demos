package enum

type EnumAdjustmentTarget int

const (
	AtBase EnumAdjustmentTarget = iota
	AtTotal
	AtOther
)

//获取EnumAdjustmentTarget的字符串值
func (this EnumAdjustmentTarget) String() string {
	switch this {
	case AtBase:
		return "Base"
	case AtTotal:
		return "Total"
	case AtOther:
		return "Other"
	}
	return ""
}

//获取EnumAdjustmentTarget的整数值
func (this EnumAdjustmentTarget) Value() int {
	if this >= AtBase && this <= AtOther {
		return int(this)
	} else {
		return -1
	}
}
