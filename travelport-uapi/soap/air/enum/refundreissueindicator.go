package enum

//An attribute specifying whether the service is refundable or reissuable.
type EnumRefundReissueIndicator int

const (
	RriRefundable EnumRefundReissueIndicator = iota
	RriNonRefundable
	RriReuse
)

//获取EnumRefundReissueIndicator的字符串值
func (this EnumRefundReissueIndicator) String() string {
	switch this {
	case RriRefundable:
		return "Refundable"
	case RriNonRefundable:
		return "NonRefundable"
	case RriReuse:
		return "Reuse"
	}
	return ""
}

//获取EnumRefundReissueIndicator的整数值
func (this EnumRefundReissueIndicator) Value() int {
	if this >= RriRefundable && this <= RriReuse {
		return int(this)
	} else {
		return -1
	}
}
