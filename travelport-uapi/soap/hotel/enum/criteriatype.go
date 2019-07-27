package enum

type EnumCriteriaType int

const (
	CtCorporateDiscountID EnumCriteriaType = iota
	CtPermittedChains
	CtHotelName
	CtDistance
	CtRateCategory
	CtHotelRating
	CtAmenities
	CtHotelTransportation
)

//获取EnumCriteriaType的字符串值
func (this EnumCriteriaType) String() string {
	switch this {
	case CtCorporateDiscountID:
		return "CorporateDiscountID"
	case CtPermittedChains:
		return "PermittedChains"
	case CtHotelName:
		return "HotelName"
	case CtDistance:
		return "Distance"
	case CtRateCategory:
		return "RateCategory"
	case CtHotelRating:
		return "HotelRating"
	case CtAmenities:
		return "Amenities"
	case CtHotelTransportation:
		return "HotelTransportation"
	}
	return ""
}

//获取EnumCriteriaType的整数值
func (this EnumCriteriaType) Value() int {
	if this >= CtCorporateDiscountID && this <= CtHotelTransportation {
		return int(this)
	} else {
		return -1
	}
}
