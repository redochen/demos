package enum

//Identifies the type of payment. This can be for an itinerary, a traveler, or a service fee for example.
type EnumPaymentType int

const (
	PtAirlineFee EnumPaymentType = iota
	PtDeliveryFee
	PtItinerary
	PtPassenger
	PtServiceFee
	PtOptionalService
	PtTicketFee
)

//获取EnumPaymentType的字符串值
func (this EnumPaymentType) String() string {
	switch this {
	case PtAirlineFee:
		return "AirlineFee"
	case PtDeliveryFee:
		return "DeliveryFee"
	case PtItinerary:
		return "Itinerary"
	case PtPassenger:
		return "Passenger"
	case PtServiceFee:
		return "ServiceFee"
	case PtOptionalService:
		return "OptionalService"
	case PtTicketFee:
		return "TicketFee"
	}
	return ""
}

//获取EnumPaymentType的整数值
func (this EnumPaymentType) Value() int {
	if this >= PtAirlineFee && this <= PtTicketFee {
		return int(this)
	} else {
		return -1
	}
}
