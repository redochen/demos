package enum

//Type of booking source sent in the Code attribute.
type EnumBookingSourceType int

const (
	BstPseudoCityCode EnumBookingSourceType = iota
	BstArcNumber
	BstIataNumber
	BstCustomerId
	BstBookingSourceOverride //The Booking Source Override is usually used when the car supplier has assigned a number (which can be alpha/numeric) to the agency/e-commerce to use in place of an IATA number. Supported provider(s) : 1P/1J
)

//获取EnumBookingSourceType的字符串值
func (this EnumBookingSourceType) String() string {
	switch this {
	case BstPseudoCityCode:
		return "PseudoCityCode"
	case BstArcNumber:
		return "ArcNumber"
	case BstIataNumber:
		return "IataNumber"
	case BstCustomerId:
		return "CustomerId"
	case BstBookingSourceOverride:
		return "BookingSourceOverride"
	}
	return ""
}

//获取EnumBookingSourceType的整数值
func (this EnumBookingSourceType) Value() int {
	if this >= BstPseudoCityCode && this <= BstBookingSourceOverride {
		return int(this)
	} else {
		return -1
	}
}
