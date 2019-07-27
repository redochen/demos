package enum

type EnumItineraryCode int

const (
	International EnumItineraryCode = iota //Indicates the itinerary is International
	Domestic                               //Indicates the itinerary is domestic

)

//获取EnumItineraryCode的字符串值
func (this EnumItineraryCode) String() string {
	switch this {
	case International:
		return "International"
	case Domestic:
		return "Domestic"
	}
	return ""
}

//获取EnumItineraryCode的整数值
func (this EnumItineraryCode) Value() int {
	if this >= International && this <= Domestic {
		return int(this)
	} else {
		return -1
	}
}
