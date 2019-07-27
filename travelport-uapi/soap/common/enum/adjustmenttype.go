package enum

type EnumAdjustmentType int

const (
	AtAmount EnumAdjustmentType = iota
	AtPercentage
)

//获取EnumAdjustmentType的字符串值
func (this EnumAdjustmentType) String() string {
	switch this {
	case AtAmount:
		return "Amount"
	case AtPercentage:
		return "Percentage"
	}
	return ""
}

//获取EnumAdjustmentType的整数值
func (this EnumAdjustmentType) Value() int {
	if this >= AtAmount && this <= AtPercentage {
		return int(this)
	} else {
		return -1
	}
}
