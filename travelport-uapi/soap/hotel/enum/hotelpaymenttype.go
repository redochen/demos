package enum

//PrePay, PostPay
type EnumHotelPaymentType int

const (
	HptPrePay EnumHotelPaymentType = iota
	HptPostPay
)

//获取EnumHotelPaymentType的字符串值
func (this EnumHotelPaymentType) String() string {
	switch this {
	case HptPrePay:
		return "PrePay"
	case HptPostPay:
		return "PostPay"
	}
	return ""
}

//获取EnumHotelPaymentType的整数值
func (this EnumHotelPaymentType) Value() int {
	if this >= HptPrePay && this <= HptPostPay {
		return int(this)
	} else {
		return -1
	}
}
