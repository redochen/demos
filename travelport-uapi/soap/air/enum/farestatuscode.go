package enum

//For an exchange request this tells if the itinerary is the original one or new one. A value of Original will only apply to 1G/1V/1P/1S/1A. A value of New will apply to 1G/1V/1P/1S/1A/ACH.
type EnumFareStatusCode int

const (
	FscReadyToTicket  EnumFareStatusCode = iota //Fare is enabled and available for ticketing
	FscUnableToTicket                           //Fare could not be ticketed
	FscReprice                                  //Fare needs to be repriced
	FscTicketed                                 //Fare is ticketed
	FscUnable                                   //Fare is not enabled
	FscUnknown                                  //To handle new enumerations added by provider but currently not recognized by API
)

//获取EnumFareStatusCode的字符串值
func (this EnumFareStatusCode) String() string {
	switch this {
	case FscReadyToTicket:
		return "ReadyToTicket"
	case FscUnableToTicket:
		return "UnableToTicket"
	case FscReprice:
		return "Reprice"
	case FscTicketed:
		return "Ticketed"
	case FscUnable:
		return "Unable"
	case FscUnknown:
		return "Unknown"
	}
	return ""
}

//获取EnumFareStatusCode的整数值
func (this EnumFareStatusCode) Value() int {
	if this >= FscReadyToTicket && this <= FscUnknown {
		return int(this)
	} else {
		return -1
	}
}
