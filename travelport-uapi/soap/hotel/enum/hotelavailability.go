package enum

//Availability status of hotel
type EnumHotelAvailability int

const (
	HaAvailable              EnumHotelAvailability = iota //Hotel is Available
	HaNotAvailable                                        //NotAvailable
	HaAvailableForOtherRates                              //Available, but not for the rates requested
	HaOnRequest                                           //On request
	HaUnknown                                             //Unknown
)

//获取EnumHotelAvailability的字符串值
func (this EnumHotelAvailability) String() string {
	switch this {
	case HaAvailable:
		return "Available"
	case HaNotAvailable:
		return "NotAvailable"
	case HaAvailableForOtherRates:
		return "AvailableForOtherRates"
	case HaOnRequest:
		return "OnRequest"
	case HaUnknown:
		return "Unknown"
	}
	return ""
}

//获取EnumHotelAvailability的整数值
func (this EnumHotelAvailability) Value() int {
	if this >= HaAvailable && this <= HaUnknown {
		return int(this)
	} else {
		return -1
	}
}
