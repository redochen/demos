package request

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
)

//BaseAirSearchReq Base Request for Low fare air Search
type BaseAirSearchReq struct {
	comrq.BaseCoreSearchReq
	//<xs:choice>
	SearchAirLeg             []*SearchAirLeg             `xml:"air:SearchAirLeg,omitempty"`             //maxOccurs="9"
	SearchSpecificAirSegment []*SearchSpecificAirSegment `xml:"air:SearchSpecificAirSegment,omitempty"` //maxOccurs="unbounded"
	//</xs:choice>
	AirSearchModifiers   *AirSearchModifiers   `xml:"air:AirSearchModifiers,omitempty"`   //minOccurs="0"
	SplitTicketingSearch *SplitTicketingSearch `xml:"air:SplitTicketingSearch,omitempty"` //minOccurs="0" maxOccurs="1"
	JourneyData          *JourneyData          `xml:"air:JourneyData,omitempty"`          //minOccurs="0"
}

//AirSearchReq Base Request for Air Search
type AirSearchReq struct {
	comrq.BaseSearchReq
	PointOfCommencement *comrq.PointOfCommencement `xml:"air:PointOfCommencement,omitempty"` //minOccurs="0" maxOccurs="1"
	//<xs:choice>
	SearchAirLeg             []*SearchAirLeg             `xml:"air:SearchAirLeg,omitempty"`             //maxOccurs="16"
	SearchSpecificAirSegment []*SearchSpecificAirSegment `xml:"air:SearchSpecificAirSegment,omitempty"` //maxOccurs="unbounded"
	//</xs:choice>
	AirSearchModifiers *AirSearchModifiers `xml:"air:AirSearchModifiers,omitempty"` //minOccurs="0"
	JourneyData        *JourneyData        `xml:"air:JourneyData,omitempty"`        //minOccurs="0"
}

//BaseLowFareSearchReq Base Low Fare Search Request
type BaseLowFareSearchReq struct {
	BaseAirSearchReq
	SearchPassenger              []*comrq.SearchPassenger `xml:"com:SearchPassenger"`                         //maxOccurs="18" //Provider: 1G,1V,1P,1J,ACH-Maxinumber of passenger increased in to 18 to support 9 INF passenger along with 9 ADT,CHD,INS passenger
	AirPricingModifiers          *AirPricingModifiers     `xml:"air:AirPricingModifiers,omitempty"`           //minOccurs="0" //Provider: 1G,1V,1P,1J,ACH.
	Enumeration                  *Enumeration             `xml:"air:Enumeration,omitempty"`                   //minOccurs="0" //Provider: 1G,1V,1P,1J,ACH.
	AirExchangeModifiers         *AirExchangeModifiers    `xml:"air:AirExchangeModifiers,omitempty"`          //minOccurs="0" //Provider: ACH.
	FlexExploreModifiers         *FlexExploreModifiers    `xml:"air:FlexExploreModifiers,omitempty"`          //minOccurs="0" //This is the container for a set of modifiers which allow the user to perform a special kind of low fare search, depicted as flex explore, based on different parameters like Area, Zone, Country, State, Specific locations, Distance around the actual destination of the itinerary. Applicable for providers 1G,1V,1P.
	PCC                          *PCC                     `xml:"air:PCC,omitempty"`                           //minOccurs="0"
	FareRulesFilterCategory      *FareRulesFilterCategory `xml:"air:FareRulesFilterCategory,omitempty"`       //minOccurs="0"
	FormOfPayment                []*comrq.FormOfPayment   `xml:"com:FormOfPayment,omitempty"`                 //minOccurs="0" maxOccurs="99" //Provider: 1P,1J
	EnablePointToPointSearch     bool                     `xml:"EnablePointToPointSearch,attr,omitempty"`     //use="optional" default="false" //Provider: 1G,1V,1P,1J,ACH-Indicates that low cost providers should be queried for top connection options and the results returned with the search.
	EnablePointToPointAlternates bool                     `xml:"EnablePointToPointAlternates,attr,omitempty"` //use="optional" default="false" //Provider: 1G,1V,1P,1J,ACH-Indicates that suggestions for alternate connection cities for low cost providers should be returned with the search.
	MaxNumberOfExpertSolutions   int                      `xml:"MaxNumberOfExpertSolutions,attr,omitempty"`   //use="optional" default="0" //Provider: 1G,1V,1P,1J,ACH-Indicates the Maximum Number of Expert Solutions to be returned from the Knowledge Base for the provided search criteria
	SolutionResult               bool                     `xml:"SolutionResult,attr,omitempty"`               //use="optional" default="false" //Provider: 1G,1V,1P,1J,ACH-Indicates whether the response will contain Solution result (AirPricingSolution) or Non Solution Result (AirPricingPoints). The default value is false. This attribute cannot be combined with EnablePointToPointSearch, EnablePointToPointAlternates and MaxNumberOfExpertSolutions.
	PreferCompleteItinerary      bool                     `xml:"PreferCompleteItinerary,attr,omitempty"`      //use="optional" default="true" //Provider: ACH-This attribute is only supported for ACH .It works in conjunction with the @SolutionResult flag
	MetaOptionIdentifier         string                   `xml:"MetaOptionIdentifier,attr,omitempty"`         //Invoke Meta Search.  Valid values are 00 to 99, or D for the default meta search configuration.  When Meta Search not requested, normal LowFareSearch applies.  Supported Providers;  1g/1v/1p/1j
	ReturnUpsellFare             bool                     `xml:"ReturnUpsellFare,attr,omitempty"`             //use="optional" default="false" //When set to “true”, Upsell information will be returned in the shop response. Provider supported : 1G, 1V, 1P, 1J
	IncludeFareInfoMessages      bool                     `xml:"IncludeFareInfoMessages,attr,omitempty"`      //use="optional" default="false" //Set to True to return FareInfoMessageList. Providers supported: 1G/1V/1P/1J
	ReturnBrandedFares           bool                     `xml:"ReturnBrandedFares,attr,omitempty"`           //use="optional" default="true" //When ReturnBrandedFares is set to “false”, Rich Content and Branding will not be returned in the shop response.  When ReturnBrandedFares it is set to “true” or is not sent, Rich Content and Branding will be returned in the shop response.  Provider: 1P/1J/ACH.
	MultiGDSSearch               bool                     `xml:"MultiGDSSearch,attr,omitempty"`               //use="optional" default="false" //A "true" value indicates MultiGDSSearch. Specific provisioning is required.
	ReturnMM                     bool                     `xml:"ReturnMM,attr,omitempty"`                     //use="optional" default="false" //If this attribute is set to “true”, Fare Control Manager processing will be invoked.
	CheckOBFees                  string                   `xml:"CheckOBFees,attr,omitempty"`                  //use="optional" //A flag to return fees for ticketing and for various forms of payment. The default is “TicketingOnly” and will return only ticketing fees.  The value “All” will return ticketing fees and the applicable form of payment fees for the form of payment information specified in the request.  “FOPOnly” will return the applicable form of payment fees for the form of payment information specified in the request. Form of payment fees are never included in the total unless specific card details are in the request.Provider notes:ACH - CheckOBFees is valid only for LowFareSearch.  The valid values are “All”, “TicketingOnly” and “None” and the default value is “None”. 1P/1J -The valid values are “All”, “None” and “TicketingOnly”.1G – All four values are supported.1V/RCH – CheckOBFees are not supported.”
	NSCC                         string                   `xml:"NSCC,attr,omitempty"`                         //use="optional" //1 to 3 numeric that defines a Search Control Console filter.This attribute is used to override that filter.
}

//BaseAirPriceReq ...
type BaseAirPriceReq struct {
	comrq.BaseCoreReq
	AirItinerary              *AirItinerary              `xml:"air:AirItinerary"`                        //Provider: 1G,1V,1P,1J,ACH.
	AirPricingModifiers       *AirPricingModifiers       `xml:"air:AirPricingModifiers,omitempty"`       //minOccurs="0" //Provider: 1G,1V,1P,1J,ACH.
	SearchPassenger           []*comrq.SearchPassenger   `xml:"com:SearchPassenger"`                     //maxOccurs="18" //Provider: 1G,1V,1P,1J,ACH-Maxinumber of passenger increased in to 18 to support 9 INF passenger along with 9 ADT,CHD,INS passenger
	AirPricingCommand         []*AirPricingCommand       `xml:"air:AirPricingCommand"`                   //maxOccurs="16" //Provider: 1G,1V,1P,1J,ACH.
	AirReservationLocatorCode *AirReservationLocatorCode `xml:"air:AirReservationLocatorCode,omitempty"` //minOccurs="0" //Provider: ACH,1P,1J
	//OptionalServices          *OptionalServices          `xml:"air:OptionalServices,omitempty"`          //minOccurs="0" //Provider: ACH.
	FormOfPayment      []*comrq.FormOfPayment `xml:"com:FormOfPayment,omitempty"`       //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH.
	PCC                *PCC                   `xml:"air:PCC,omitempty"`                 //minOccurs="0"
	CheckOBFees        string                 `xml:"CheckOBFees,attr,omitempty"`        //use="optional" //A flag to return fees for ticketing and for various forms of payment. The default is “TicketingOnly” and will return only ticketing fees.  The value “All” will return ticketing fees and the applicable form of payment fees for the form of payment information specified in the request.  “FOPOnly” will return the applicable form of payment fees for the form of payment information specified in the request. Form of payment fees are never included in the total unless specific card details are in the request.Provider notes:ACH - CheckOBFees is valid only for LowFareSearch.  The valid values are “All”, “TicketingOnly” and “None” and the default value is “None”. 1P/1J -The valid values are “All”, “None” and “TicketingOnly”.1G – All four values are supported.1V/RCH – CheckOBFees are not supported.”
	FareRuleType       string                 `xml:"FareRuleType,attr,omitempty"`       //use="optional" default="none" //Provider: 1G,1V,1P,1J,ACH.
	SupplierCode       com.TypeSupplierCode   `xml:"SupplierCode,attr,omitempty"`       //Specifies the supplier/ vendor for vendor specific price requests
	TicketDate         string                 `xml:"TicketDate,attr,omitempty"`         //use="optional" //type="xs:date" //YYYY-MM-DD Using a date in the past is a request for an historical fare
	CheckFlightDetails bool                   `xml:"CheckFlightDetails,attr,omitempty"` //default="false" //To Include FlightDetails in Response set to “true” the Default value is “false”.
	ReturnMM           bool                   `xml:"ReturnMM,attr,omitempty"`           //use="optional" default="false" //If this attribute is set to “true”, Fare Control Manager processing will be invoked.
	NSCC               string                 `xml:"NSCC,attr,omitempty"`               //use="optional" //1 to 3 numeric that defines a Search Control Console filter.This attribute is used to override that filter.
	SplitPricing       bool                   `xml:"SplitPricing,attr,omitempty"`       //use="optional" default="false" //Indicates whether the AirSegments should be priced together or separately. Set ‘true’ for split pricing. Set ‘false’ for pricing together.
}
