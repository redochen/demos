package enum

type EnumFarePull int

const (
	FpReverseOfOriginDestination EnumFarePull = iota
	FpSameAsOriginDestination
)

//获取EnumFarePull的字符串值
func (this EnumFarePull) String() string {
	switch this {
	case FpReverseOfOriginDestination:
		return "Base"
	case FpSameAsOriginDestination:
		return "Total"
	}
	return ""
}

//获取EnumFarePull的整数值
func (this EnumFarePull) Value() int {
	if this >= FpReverseOfOriginDestination &&
		this <= FpSameAsOriginDestination {
		return int(this)
	} else {
		return -1
	}
}
