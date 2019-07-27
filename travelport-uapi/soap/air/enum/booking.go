package enum

//Type of booking
type EnumBooking int

const (
	SSR EnumBooking = iota
	AuxillarySegment
	AvailableForDisplayOrPricing
	ContactCarrierForBooking
	NoBookingRequired
	ApplyBookingPerService
)

//获取EnumBooking的字符串值
func (b EnumBooking) String() string {
	switch b {
	case SSR:
		return "SSR"
	case AuxillarySegment:
		return "Auxillary Segment"
	case AvailableForDisplayOrPricing:
		return "Available for Display/Pricing"
	case ContactCarrierForBooking:
		return "Contact Carrier for Booking"
	case NoBookingRequired:
		return "No Booking Required"
	case ApplyBookingPerService:
		return "Apply booking per service"
	}
	return ""
}

//获取EnumBooking的整数值
func (b EnumBooking) Value() int {
	if b >= SSR && b <= ApplyBookingPerService {
		return int(b)
	} else {
		return -1
	}
}
