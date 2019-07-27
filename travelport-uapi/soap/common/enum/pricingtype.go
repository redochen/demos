package enum

type EnumPricingType int

const (
	PtClassBooked EnumPricingType = iota
	PtLowestClass
	PtLowestQuote
)

//获取EnumPricingType的字符串值
func (this EnumPricingType) String() string {
	switch this {
	case PtClassBooked:
		return "ClassBooked"
	case PtLowestClass:
		return "LowestClass"
	case PtLowestQuote:
		return "LowestQuote"
	}
	return ""
}

//获取EnumPricingType的整数值
func (this EnumPricingType) Value() int {
	if this >= PtClassBooked && this <= PtLowestQuote {
		return int(this)
	} else {
		return -1
	}
}
