package enum

//Used in Low Fare Search to better target the results
type EnumTripType int

const (
	TtCheapest EnumTripType = iota
	TtQuickest
	TtMostConvenient
	TtLeisure
	TtBusiness
	TtLuxury
	TtPreferFirst
	TtBusinessOrFirst
	TtNoPenalty
)

//获取EnumTripType的字符串值
func (tt EnumTripType) String() string {
	switch tt {
	case TtCheapest:
		return "Cheapest"
	case TtQuickest:
		return "Quickest"
	case TtMostConvenient:
		return "MostConvenient"
	case TtLeisure:
		return "Leisure"
	case TtBusiness:
		return "Business"
	case TtLuxury:
		return "Luxury"
	case TtPreferFirst:
		return "PreferFirst"
	case TtBusinessOrFirst:
		return "BusinessOrFirst"
	case TtNoPenalty:
		return "NoPenalty"
	}
	return ""
}

//获取EnumTripType的整数值
func (tt EnumTripType) Value() int {
	if tt >= TtCheapest && tt <= TtNoPenalty {
		return int(tt)
	} else {
		return -1
	}
}
