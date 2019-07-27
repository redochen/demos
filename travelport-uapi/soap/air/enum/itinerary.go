package enum

//For an exchange request this tells if the itinerary is the original one or new one. A value of Original will only apply to 1G/1V/1P/1S/1A. A value of New will apply to 1G/1V/1P/1S/1A/ACH.
type EnumItinerary int

const (
	ItNew EnumItinerary = iota
	ItOriginal
)

//获取EnumItinerary的字符串值
func (this EnumItinerary) String() string {
	switch this {
	case ItNew:
		return "New"
	case ItOriginal:
		return "Original"
	}
	return ""
}

//获取EnumItinerary的整数值
func (this EnumItinerary) Value() int {
	if this >= ItNew && this <= ItOriginal {
		return int(this)
	} else {
		return -1
	}
}
