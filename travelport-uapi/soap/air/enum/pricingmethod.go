package enum

//The method at which the pricing data was acquired
type EnumPricingMethod int

const (
	PmAuto       EnumPricingMethod = iota //Automatically generated
	PmManual                              //Agent has overridden default(s)
	PmManualFare                          //Fare has been constructed by agent
	PmGuaranteed                          //Fare is guaranteed
	PmInvalid                             //Invalid fare, e.g. due to name or itinerary change
	PmRestored                            //Ticketed stored fare has been restored
	PmTicketed
	PmUnticketable                      //Unable to ticket
	PmReprice                           //Need requote to ticket
	PmExpired                           //Expired fare, older than 7 days
	PmAutoUsingPrivateFare              //Agency private fares that are not guaranteed
	PmGuaranteedUsingAirlinePrivateFare //Guaranteed fare using Airline private fare that was filed with a fare distributor.
	PmAirline                           //Fare created as a result of Claim PNR which transfers data to GDS for ticketing purposes.
	PmAgentAssisted                     //Fare is created using Agent Asisted Pricing. Worldspan TKG FAX Line Documentation - AGENT ASSISTEDPRICED
	PmVerifyPrice                       //Verify existing saved price on PNR . Worldspan TKG FAX Line Documentation -  AWAITING PRICE VERIFICATION
	PmAltSegmentRemovedReprice          //ALT Segment removed, Reprice pricing.  Worldspan TKG FAX Line Documentation - AWAITING REPRICING ALT SEGS RMVD
	PmAuxiliarySegmentRemovedReprice    //AUX Segment removed, Reprice pricing. Worldspan TKG FAX Line Documentation -  AWAITING REPRICING AUX SEGS REMOVED
	PmDuplicateSegmentRemovedReprice    //Duplicate Segment removed, Reprice pricing.  Worldspan TKG FAX Line Documentation - AWAITING REPRICING DUPE SEGS REMOVED
	PmUnknown                           //Any other kind of Pricing Method which is not supported by API.
	PmGuaranteedUsingAgencyPrivateFare  //Guaranteed fare using Agency private fare that was filed with a fare distributor.
	PmAutoRapidReprice                  //Auto priced by rapid reprice. Provider 1P FCI code 4 .
)

//获取EnumPricingMethod的字符串值
func (this EnumPricingMethod) String() string {
	switch this {
	case PmAuto:
		return "Auto"
	case PmManual:
		return "Manual"
	case PmManualFare:
		return "ManualFare"
	case PmGuaranteed:
		return "Guaranteed"
	case PmInvalid:
		return "Invalid"
	case PmRestored:
		return "Restored"
	case PmTicketed:
		return "Ticketed"
	case PmUnticketable:
		return "Unticketable"
	case PmReprice:
		return "Reprice"
	case PmExpired:
		return "Expired"
	case PmAutoUsingPrivateFare:
		return "AutoUsingPrivateFare"
	case PmGuaranteedUsingAirlinePrivateFare:
		return "GuaranteedUsingAirlinePrivateFare"
	case PmAirline:
		return "Airline"
	case PmAgentAssisted:
		return "AgentAssisted"
	case PmVerifyPrice:
		return "VerifyPrice"
	case PmAltSegmentRemovedReprice:
		return "AltSegmentRemovedReprice"
	case PmAuxiliarySegmentRemovedReprice:
		return "AuxiliarySegmentRemovedReprice"
	case PmDuplicateSegmentRemovedReprice:
		return "DuplicateSegmentRemovedReprice"
	case PmUnknown:
		return "Unknown"
	case PmGuaranteedUsingAgencyPrivateFare:
		return "GuaranteedUsingAgencyPrivateFare"
	case PmAutoRapidReprice:
		return "AutoRapidReprice"
	}
	return ""
}

//获取EnumPricingMethod的整数值
func (this EnumPricingMethod) Value() int {
	if this >= PmAuto && this <= PmAutoRapidReprice {
		return int(this)
	} else {
		return -1
	}
}
