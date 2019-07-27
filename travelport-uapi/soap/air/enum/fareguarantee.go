package enum

//The status of a fare
type EnumFareGuarantee int

const (
	FgAuto       EnumFareGuarantee = iota //Automatically generated
	FgManual                              //Agent has overridden default(s)
	FgManualFare                          //Fare has been constructed by agent
	FgGuaranteed                          //Fare is guaranteed
	FgInvalid                             //Invalid fare, e.g. due to name or itinerary change
	FgRestored                            //Ticketed stored fare has been restored
	FgTicketed
	FgUnticketable                      //Unable to ticket
	FgReprice                           //Need requote to ticket
	FgExpired                           //Expired fare due to older fare guarantee date typically older than 7 days
	FgAutoUsingPrivateFare              //Agency private fares that are not guaranteed
	FgGuaranteedUsingAirlinePrivateFare //Guaranteed fare using Airline private fare that was filed with a fare distributor.
	FgAirline                           //Fare guaranteed by Airline.
	FgGuaranteeExpired                  //Guaranteed fare recently got expired as ticketing hadn't been done within a time frame typically midnight local time of POS .
	FgAgencyPrivateFareNoOverride       //Agency Private Fare with no rules override
	FgUnknown                           //To handle new enumerations added by provider but currently not recognized by API
)

//获取EnumFareGuarantee的字符串值
func (this EnumFareGuarantee) String() string {
	switch this {
	case FgAuto:
		return "Auto"
	case FgManual:
		return "Manual"
	case FgManualFare:
		return "ManualFare"
	case FgGuaranteed:
		return "Guaranteed"
	case FgInvalid:
		return "Invalid"
	case FgRestored:
		return "Restored"
	case FgTicketed:
		return "Ticketed"
	case FgUnticketable:
		return "Unticketable"
	case FgReprice:
		return "Reprice"
	case FgExpired:
		return "Expired"
	case FgAutoUsingPrivateFare:
		return "AutoUsingPrivateFare"
	case FgGuaranteedUsingAirlinePrivateFare:
		return "GuaranteedUsingAirlinePrivateFare"
	case FgAirline:
		return "Airline"
	case FgGuaranteeExpired:
		return "GuaranteeExpired"
	case FgAgencyPrivateFareNoOverride:
		return "AgencyPrivateFareNoOverride"
	case FgUnknown:
		return "Unknown"
	}
	return ""
}

//获取EnumFareGuarantee的整数值
func (this EnumFareGuarantee) Value() int {
	if this >= FgAuto && this <= FgUnknown {
		return int(this)
	} else {
		return -1
	}
}
