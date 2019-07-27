package enum

type EnumDisplayCategory int

const (
	DcWithItineraryPricing EnumDisplayCategory = iota
	DcStore
	DcSpecialService
)

//获取EnumDisplayCategory的字符串值
func (this EnumDisplayCategory) String() string {
	switch this {
	case DcWithItineraryPricing:
		return "With Itinerary Pricing"
	case DcStore:
		return "Store"
	case DcSpecialService:
		return "SpecialService"
	}
	return ""
}

//获取EnumDisplayCategory的整数值
func (this EnumDisplayCategory) Value() int {
	if this >= DcWithItineraryPricing && this <= DcSpecialService {
		return int(this)
	} else {
		return -1
	}
}
