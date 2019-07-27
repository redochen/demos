package request

import (
	air "github.com/redochen/demos/travelport-uapi/soap/air"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	"time"
)

//[RQ] Search version of AirLeg used to specify search criteria
type SearchAirLeg struct {
	SearchOrigin      []*comrq.SearchLocation   `xml:"air:SearchOrigin"`                //maxOccurs="unbounded"
	SearchDestination []*comrq.SearchLocation   `xml:"air:SearchDestination,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	SearchDepTime     []*comrq.FlexibleTimeSpec `xml:"air:SearchDepTime"`               //maxOccurs="unbounded"
	SearchArvTime     []*comrq.TimeSpec         `xml:"air:SearchArvTime,omitempty"`     //maxOccurs="unbounded" //Specifies the preferred time within the time range. For 1G, 1V, 1P, 1J, it is supported for AvailabilitySearchReq (TimeRange must also be specified) and not supported for LowFareSearchReq. ACH does not support search by arrival time.
	AirLegModifiers   *AirLegModifiers          `xml:"air:AirLegModifiers,omitempty"`   //minOccurs="0"
}

//[RQ]
type AirLegModifiers struct {
	PermittedCabins                *PermittedCabins        `xml:"air:PermittedCabins,omitempty"`                 //minOccurs="0"
	PreferredCabins                *PreferredCabins        `xml:"air:PreferredCabins,omitempty"`                 //minOccurs="0"
	PermittedCarriers              *PermittedCarriers      `xml:"air:PermittedCarriers,omitempty"`               //minOccurs="0"
	ProhibitedCarriers             *ProhibitedCarriers     `xml:"air:ProhibitedCarriers,omitempty"`              //minOccurs="0"
	PreferredCarriers              *PreferredCarriers      `xml:"air:PreferredCarriers,omitempty"`               //minOccurs="0"
	PermittedConnectionPoints      *ConnectionPoints       `xml:"air:PermittedConnectionPoints,omitempty"`       //minOccurs="0" //This is the container to specify all permitted connection points. Applicable for 1G/1V/1P/1J.
	ProhibitedConnectionPoints     *ConnectionPoints       `xml:"air:ProhibitedConnectionPoints,omitempty"`      //minOccurs="0" //This is the container to specify all prohibited connection points. Applicable for 1G/1V/1P/1J.
	PreferredConnectionPoints      *ConnectionPoints       `xml:"air:PreferredConnectionPoints,omitempty"`       //minOccurs="0" //This is the container to specify all preferred connection points. Applicable for 1G/1V only.
	PermittedBookingCodes          *PermittedBookingCodes  `xml:"air:PermittedBookingCodes,omitempty"`           //minOccurs="0"
	ProhibitedBookingCodes         *ProhibitedBookingCodes `xml:"air:ProhibitedBookingCodes,omitempty"`          //minOccurs="0"
	PreferredBookingCodes          *PreferredBookingCodes  `xml:"air:PreferredBookingCodes,omitempty"`           //minOccurs="0"
	PreferredAlliances             *PreferredAlliances     `xml:"air:PreferredAlliances,omitempty"`              //minOccurs="0"
	DisfavoredAlliances            *DisfavoredAlliances    `xml:"air:DisfavoredAlliances,omitempty"`             //minOccurs="0"
	FlightType                     *FlightType             `xml:"air:FlightType,omitempty"`                      //minOccurs="0"
	AnchorFlightData               *TypeAnchorFlightData   `xml:"air:AnchorFlightData,omitempty"`                //minOccurs="0" maxOccurs="1"
	ProhibitOvernightLayovers      bool                    `xml:"ProhibitOvernightLayovers,attr,omitempty"`      //use="optional" default="false" //If true, excludes connections if arrival time of first flight and departure time of second flight is on 2 different calendar days. When used in conjunction with MaxConnectionTime, it would exclude all connections if the connecting flights wait time exceeds the time specified in MaxConnectionTime.
	MaxConnectionTime              int                     `xml:"MaxConnectionTime,attr,omitempty"`              //use="optional"
	ReturnFirstAvailableOnly       bool                    `xml:"ReturnFirstAvailableOnly,attr,omitempty"`       //use="optional" //If it is true then it will search for first available for the booking code designated or any booking code in same cabin.
	AllowDirectAccess              bool                    `xml:"AllowDirectAccess,attr,omitempty"`              //use="optional" default="false" //If it is true request will be sent directly to the carrier.
	ProhibitMultiAirportConnection bool                    `xml:"ProhibitMultiAirportConnection,attr,omitempty"` //Indicates whether to restrict multi-airport connections
	PreferNonStop                  bool                    `xml:"PreferNonStop,attr,omitempty"`                  //use="optional" default="false" //When non-stops are preferred, the distribution of search results should skew heavily toward non-stop flights while still returning some one stop flights for comparison and price competitiveness. The search request will ‘boost' the preference towards non-stops. If true then Non Stop flights will be preferred.
	OrderBy                        string                  `xml:"OrderBy,attr,omitempty"`                        //use="optional" //Indicates whether to sort by Journey Time, Deparature Time or Arrival Time
	MaxJourneyTime                 uint                    `xml:"MaxJourneyTime,attr,omitempty"`                 //use="optional" //Maximum Journey Time for this leg (in hours) 0-99. Supported Providers 1G,1V.
}

//[RQ]
type ConnectionPoints struct {
	ConnectionPoint []*comrq.ConnectionPoint `xml:"com:ConnectionPoint,omitempty"` //maxOccurs="unbounded"
}

//[RQ] Modifier to request flight type options example non-stop only, non-stop and direct only, include single online connection etc.
type FlightType struct {
	RequireSingleCarrier bool `xml:"RequireSingleCarrier,attr,omitempty"` //use="optional" default="false"
	MaxConnections       int  `xml:"MaxConnections,attr,omitempty"`       //use="optional" default="-1" //To support other provider like 1P maxConnection values ranged from -1 to 3. Here -1 means user don't know about connection information.  In case of maxConnection -1, maxStop value will be considered.
	MaxStops             int  `xml:"MaxStops,attr,omitempty"`             //use="optional" default="-1" //To support other provider like 1P maxStop values ranged from -1 to 3. Here -1 means provider can ignore the maxStop value.
	NonStopDirects       bool `xml:"NonStopDirects,attr,omitempty"`       //use="optional"
	StopDirects          bool `xml:"StopDirects,attr,omitempty"`          //use="optional"
	SingleOnlineCon      bool `xml:"SingleOnlineCon,attr,omitempty"`      //use="optional"
	DoubleOnlineCon      bool `xml:"DoubleOnlineCon,attr,omitempty"`      //use="optional"
	TripleOnlineCon      bool `xml:"TripleOnlineCon,attr,omitempty"`      //use="optional"
	SingleInterlineCon   bool `xml:"SingleInterlineCon,attr,omitempty"`   //use="optional"
	DoubleInterlineCon   bool `xml:"DoubleInterlineCon,attr,omitempty"`   //use="optional"
	TripleInterlineCon   bool `xml:"TripleInterlineCon,attr,omitempty"`   //use="optional"
}

//[RQ]
type SearchSpecificAirSegment struct {
	DepartureTime string `xml:"DepartureTime,attr"` //use="required" //The date and time at which this entity departs. This does not include time zone information since it can be derived from the origin location.
	Carrier       string `xml:"Carrier,attr"`       //use="required" //The carrier that is marketing this segment
	FlightNumber  string `xml:"FlightNumber,attr"`  //use="required" //The flight number under which the marketing carrier is marketing this flight
	Origin        string `xml:"Origin,attr"`        //use="required" //The IATA location code for this origination of this entity.
	Destination   string `xml:"Destination,attr"`   //use="required" //The IATA location code for this destination of this entity.
	SegmentIndex  int    `xml:"SegmentIndex,attr"`  //use="optional" //The sequential AirSegment number that this segment connected to.
}

//[RQ] Controls and switches for the Air Search request
type AirSearchModifiers struct {
	DisfavoredProviders            *DisfavoredProviders    `xml:"air:DisfavoredProviders,omitempty"`             //minOccurs="0"
	PreferredProviders             *PreferredProviders     `xml:"air:PreferredProviders,omitempty"`              //minOccurs="0"
	DisfavoredCarriers             *DisfavoredCarriers     `xml:"air:DisfavoredCarriers,omitempty"`              //minOccurs="0"
	PermittedCarriers              *PermittedCarriers      `xml:"air:PermittedCarriers,omitempty"`               //minOccurs="0"
	ProhibitedCarriers             *ProhibitedCarriers     `xml:"air:ProhibitedCarriers,omitempty"`              //minOccurs="0"
	PreferredCarriers              *PreferredCarriers      `xml:"air:PreferredCarriers,omitempty"`               //minOccurs="0"
	PermittedCabins                *PermittedCabins        `xml:"air:PermittedCabins,omitempty"`                 //minOccurs="0"
	PreferredCabins                *PreferredCabins        `xml:"air:PreferredCabins,omitempty"`                 //minOccurs="0"
	PreferredAlliances             *PreferredAlliances     `xml:"air:PreferredAlliances,omitempty"`              //minOccurs="0" maxOccurs="unbounded"
	DisfavoredAlliances            *DisfavoredAlliances    `xml:"air:DisfavoredAlliances,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	PermittedBookingCodes          *PermittedBookingCodes  `xml:"air:PermittedBookingCodes,omitempty"`           //minOccurs="0" maxOccurs="unbounded" //This is the container to specify all permitted booking codes
	PreferredBookingCodes          *PreferredBookingCodes  `xml:"air:PreferredBookingCodes,omitempty"`           //minOccurs="0"
	ProhibitedBookingCodes         *ProhibitedBookingCodes `xml:"air:ProhibitedBookingCodes,omitempty"`          //minOccurs="0" //This is the container to specify all prohibited booking codes
	FlightType                     *FlightType             `xml:"air:FlightType,omitempty,omitempty"`            //minOccurs="0"
	MaxLayoverDuration             *MaxLayoverDurationType `xml:"air:MaxLayoverDuration,omitempty,omitempty"`    //minOccurs="0" //This is the maximum duration the layover may have for each trip in the request. Supported providers 1P, 1J.
	NativeSearchModifier           *NativeSearchModifier   `xml:"air:NativeSearchModifier,omitempty,omitempty"`  //minOccurs="0" //Container for Native command modifiers. Providers supported : 1P
	DistanceType                   string                  `xml:"DistanceType,attr,omitempty"`                   //use="optional" default="MI"
	IncludeFlightDetails           bool                    `xml:"IncludeFlightDetails,attr,omitempty"`           //use="optional" default="true"
	AllowChangeOfAirport           bool                    `xml:"AllowChangeOfAirport,attr,omitempty"`           //use="optional" default="true"
	ProhibitOvernightLayovers      bool                    `xml:"ProhibitOvernightLayovers,attr,omitempty"`      //use="optional" default="false" //If true, excludes connections if arrival time of first flight and departure time of second flight is on 2 different calendar days. When used in conjunction with MaxConnectionTime, it would exclude all connections if the connecting flights wait time exceeds the time specified in MaxConnectionTime.
	MaxSolutions                   int                     `xml:"MaxSolutions,attr,omitempty"`                   //use="optional" //The maximum number of solutions to return. Decreasing this number
	MaxConnectionTime              int                     `xml:"MaxConnectionTime,attr,omitempty"`              //use="optional" //The maximum anount of time (in minutes) that a  solution can contain for connections between flights.
	SearchWeekends                 bool                    `xml:"SearchWeekends,attr,omitempty"`                 //A value of true indicates that search should be expanded to include weekend combinations, if applicable.
	IncludeExtraSolutions          bool                    `xml:"IncludeExtraSolutions,attr,omitempty"`          //If true, indicates that search should be made for returning more solutions, if available. For example, for certain providers, premium members may have the facility to get more solutions. This attribute may have to be combined with other applicable modifiers (like SearchWeekends) to return more results.
	ProhibitMultiAirportConnection bool                    `xml:"ProhibitMultiAirportConnection,attr,omitempty"` //Indicates whether to restrict multi-airport connections
	PreferNonStop                  bool                    `xml:"PreferNonStop,attr,omitempty"`                  //use="optional" default="false" //When non-stops are preferred, the distribution of search results should skew heavily toward non-stop flights while still returning some one stop flights for comparison and price competitiveness. The search request will ‘boost' the preference towards non-stops. If true then Non Stop flights will be preferred.
	OrderBy                        string                  `xml:"OrderBy,attr,omitempty"`                        //use="optional" //Indicates whether to sort by Journey Time, Deparature Time or Arrival Time. Applicable to air availability only.
	ExcludeOpenJawAirport          bool                    `xml:"ExcludeOpenJawAirport,attr,omitempty"`          //use="optional" default="false" //This option ensures that travel into/out of each location will be into/out of the same airport of that location. Values are true or false. Default value is 'false'. If value is true then open jaws are exclude. If false the open jaws are included. The supported providers: 1P, 1J
	ExcludeGroundTransportation    bool                    `xml:"ExcludeGroundTransportation,attr,omitempty"`    //use="optional" default="false" //Indicates whether to allow the user to exclude ground transportation or not. Default value is 'false'. If value is true then ground transportations are excluded. If false then ground transportations are included. The supported providers: 1P, 1J
	MaxJourneyTime                 uint                    `xml:"MaxJourneyTime,attr,omitempty"`                 //use="optional" //Maximum Journey Time for all legs (in hours) 0-99. For LFS Supported Providers are 1G,1V,1P,1J. For AirAvail Supported Providers are 1G,1V.
}

//[RQ]
type DisfavoredProviders struct {
	Provider []*comrq.Provider `xml:"com:Provider"` //maxOccurs="unbounded"
}

//[RQ]
type PreferredProviders struct {
	Provider []*comrq.Provider `xml:"com:Provider"` //maxOccurs="unbounded"
}

//[RQ]
type DisfavoredCarriers struct {
	Carrier []*comrq.Carrier `xml:"com:Carrier"` //maxOccurs="unbounded"
}

//[RQ]
type PreferredCarriers struct {
	Carrier []*comrq.Carrier `xml:"com:Carrier"` //maxOccurs="unbounded"
}

//[RQ]
type ProhibitedCarriers struct {
	Carrier []*comrq.Carrier `xml:"com:Carrier"` //maxOccurs="unbounded"
}

//[RQ]
type PermittedCarriers struct {
	Carrier []*comrq.Carrier `xml:"com:Carrier"` //maxOccurs="unbounded"
}

//[RQ]
type PermittedCabins struct {
	CabinClass []*comrq.CabinClass `xml:"com:CabinClass"` //maxOccurs="5"
}

//[RQ]
type PreferredCabins struct {
	CabinClass *comrq.CabinClass `xml:"com:CabinClass"`
}

//[RQ]
type ProhibitedCabins struct {
	CabinClass []*comrq.CabinClass `xml:"com:CabinClass"` //maxOccurs="3"
}

//[RQ] The Booking Code (Class of Service) for a segment
type BookingCode struct {
	Code string `xml:"Code,attr"` //use="required"
}

//[RQ]
type PermittedBookingCodes struct {
	BookingCode []*BookingCode `xml:"air:BookingCode"` //maxOccurs="unbounded"
}

//[RQ] This is the container to specify all preferred booking codes
type PreferredBookingCodes struct {
	BookingCode []*BookingCode `xml:"air:BookingCode"` //maxOccurs="unbounded"
}

//[RQ]
type ProhibitedBookingCodes struct {
	BookingCode []*BookingCode `xml:"air:BookingCode"` //maxOccurs="unbounded"
}

//[RQ] The possible values are *A for Star Alliance,*O for One world,*S for Sky team etc.
type Alliance struct {
	Code string `xml:"Code,attr"` //use="required"
}

//[RQ]
type PreferredAlliances struct {
	Alliance []*Alliance `xml:"air:Alliance"` //maxOccurs="unbounded"
}

//[RQ]
type DisfavoredAlliances struct {
	Alliance []*Alliance `xml:"air:Alliance"` //maxOccurs="unbounded"
}

//[RQ] User can specify its attribute's value in Minutes. Maximum size of each attribute is 4.
type MaxLayoverDurationType struct {
	Domestic      air.MaxLayoverDurationRangeType `xml:"Domestic,attr,omitempty"`      //use="optional" //It will be applied for all Domestic-to-Domestic connections.
	Gateway       air.MaxLayoverDurationRangeType `xml:"Gateway,attr,omitempty"`       //It will be applied for all Domestic to International and International to Domestic connections.
	International air.MaxLayoverDurationRangeType `xml:"International,attr,omitempty"` //It will be applied for all International-to-International connections.
}

//[RQ]
type NativeSearchModifier struct {
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr"` //use="required"
}

//[RQ] Performs journey aware air availability
type JourneyData struct {
	AirSegment []*AirSegment `xml:"air:AirSegment"` //maxOccurs="99"
}

//[RQ]
type BaseAirSegment struct {
	comrq.Segment
	SponsoredFltInfo             *SponsoredFltInfo               `xml:"air:SponsoredFltInfo,omitempty"`             //minOccurs="0"
	CodeshareInfo                *CodeshareInfo                  `xml:"air:CodeshareInfo,omitempty"`                //minOccurs="0"
	AirAvailInfo                 []*AirAvailInfo                 `xml:"air:AirAvailInfo,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	FlightDetails                []*FlightDetails                `xml:"air:FlightDetails,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	FlightDetailsRef             []*FlightDetailsRef             `xml:"air:FlightDetailsRef,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	AlternateLocationDistanceRef []*AlternateLocationDistanceRef `xml:"air:AlternateLocationDistanceRef,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	Connection                   *Connection                     `xml:"air:Connection,omitempty"`                   //minOccurs="0"
	SellMessage                  []*com.SellMessage              `xml:"com:SellMessage,omitempty"`                  //minOccurs="0" maxOccurs="unbounded"
	RailCoachDetails             *RailCoachDetails               `xml:"air:RailCoachDetails,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	OpenSegment                  bool                            `xml:"OpenSegment,attr,omitempty"`                 //use="optional" //Indicates OpenSegment when True
	Group                        int                             `xml:"Group,attr"`                                 //use="required" //The Origin Destination Grouping of this segment
	Carrier                      string                          `xml:"Carrier,attr"`                               //The carrier that is marketing this segment
	CabinClass                   string                          `xml:"CabinClass,attr,omitempty"`                  //Specifies Cabin class for a group of class of services. Cabin class is not identified if it is not present.
	FlightNumber                 com.TypeFlightNumber            `xml:"FlightNumber,attr"`                          //The flight number under which the marketing carrier is marketing this flight
	//<xs:attributeGroup name="attrOrigDestDepatureInfo">
	Origin        com.TypeIATACode `xml:"Origin,attr"`                  //use="required" //The IATA location code for this origination of this entity.
	Destination   com.TypeIATACode `xml:"Destination,attr"`             //use="required" //The IATA location code for this destination of this entity.
	DepartureTime string           `xml:"DepartureTime,attr,omitempty"` //use="optional" //The date and time at which this entity departs. This does not include time zone information since it can be derived from the origin location.
	ArrivalTime   string           `xml:"ArrivalTime,attr,omitempty"`   //use="optional" //The date and time at which this entity arrives at the destination. This does not include time zone information since it can be derived from the origin location.
	//</xs:attributeGroup>
	//<xs:attributeGroup ref="common:attrFlightTimes">
	FlightTime int `xml:"FlightTime,attr,omitempty"` //use="optional" //Time spent (minutes) traveling in flight, including airport taxi time.
	TravelTime int `xml:"TravelTime,attr,omitempty"` //use="optional" //Total time spent (minutes) traveling including flight time and ground time.
	Distance   int `xml:"Distance,attr,omitempty"`   //use="optional" //The distance traveled. Units are specified in the parent response element.
	//</xs:attributeGroup>
	//<xs:attributeGroup name="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional"
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional"
	//</xs:attributeGroup>
	ParticipantLevel                  string                     `xml:"ParticipantLevel,attr,omitempty"`                  //use="optional" //Type of sell agreement between host and link carrier.
	LinkAvailability                  bool                       `xml:"LinkAvailability,attr,omitempty"`                  //use="optional" //Indicates if carrier has link (carrier specific) display option.
	PolledAvailabilityOption          string                     `xml:"PolledAvailabilityOption,attr,omitempty"`          //use="optional" //Indicates if carrier has Inside (polled)Availability option.
	AvailabilityDisplayType           string                     `xml:"AvailabilityDisplayType,attr,omitempty"`           //use="optional" //The type of availability from which the segment is sold.Possible Values (List): G - General S - Flight Specific L - Carrier Specific/Direct Access M - Manual Sell F - Fare Shop/Optimal Shop Q - Fare Specific Fare Quote unbooked R - Redemption Availability used to complete the sell. Supported Providers: 1G,1V.
	ClassOfService                    com.TypeClassOfService     `xml:"ClassOfService,attr,omitempty"`                    //use="optional"
	ETicketability                    string                     `xml:"ETicketability,attr,omitempty"`                    //use="optional" //Identifies if this particular segment is E-Ticketable
	Equipment                         air.TypeEquipment          `xml:"Equipment,attr,omitempty"`                         //use="optional" //Identifies the equipment that this segment is operating under.
	MarriageGroup                     int                        `xml:"MarriageGroup,attr,omitempty"`                     //use="optional" //Identifies this segment as being a married segment. It is paired with other segments of the same value.
	NumberOfStops                     int                        `xml:"NumberOfStops,attr,omitempty"`                     //use="optional" //Identifies the number of stops for each within the segment.
	Seamless                          bool                       `xml:"Seamless,attr,omitempty"`                          //use="optional" //Identifies that this segment was sold via a direct access channel to the marketing carrier.
	ChangeOfPlane                     bool                       `xml:"ChangeOfPlane,attr,omitempty"`                     //use="optional" default="false" //Indicates the traveler must change planes between flights.
	GuaranteedPaymentCarrier          string                     `xml:"GuaranteedPaymentCarrier,attr,omitempty"`          //use="optional" //Identifies that this segment has Guaranteed Payment Carrier.
	HostTokenRef                      string                     `xml:"HostTokenRef,attr,omitempty"`                      //use="optional" //Identifies that this segment has Guaranteed Payment Carrier.
	ProviderReservationInfoRef        com.TypeRef                `xml:"ProviderReservationInfoRef,attr,omitempty"`        //use="optional" //Provider reservation reference key.
	PassiveProviderReservationInfoRef com.TypeRef                `xml:"PassiveProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	OptionalServicesIndicator         bool                       `xml:"OptionalServicesIndicator,attr,omitempty"`         //use="optional" //Indicates true if flight provides optional services.
	AvailabilitySource                air.TypeAvailabilitySource `xml:"AvailabilitySource,attr,omitempty"`                //use="optional" //Indicates Availability source of AirSegment.
	APISRequirementsRef               string                     `xml:"APISRequirementsRef,attr,omitempty"`               //use="optional" //Reference to the APIS Requirements for this AirSegment.
	BlackListed                       bool                       `xml:"BlackListed,attr,omitempty"`                       //use="optional" //Indicates blacklisted carriers which are banned from servicing points to, from and within the European Community.
	OperationalStatus                 string                     `xml:"OperationalStatus,attr,omitempty"`                 //use="optional" //Refers to the flight operational status for the segment. This attribute will only be returned in the AvailabilitySearchRsp and not used/returned in any other request/responses. If this attribute is not returned back in the response, it means the flight is operational and not past scheduled departure.
	NumberInParty                     uint                       `xml:"NumberInParty,attr,omitempty"`                     //Number of person traveling in this air segment excluding the number of infants on lap.
	RailCoachNumber                   string                     `xml:"RailCoachNumber,attr,omitempty"`                   //use="optional" //Coach number for which rail seatmap/coachmap is returned.
	BookingDate                       string                     `xml:"BookingDate,attr,omitempty"`                       //use="optional" //time.Time //Used for rapid reprice. The date the booking was made. Providers: 1G/1V/1P/1S/1A
	FlownSegment                      bool                       `xml:"FlownSegment,attr,omitempty"`                      //use="optional" default="false" //Used for rapid reprice. Tells whether or not the air segment has been flown. Providers: 1G/1V/1P/1S/1A
	ScheduleChange                    bool                       `xml:"ScheduleChange,attr,omitempty"`                    //use="optional" default="false" //Used for rapid reprice. Tells whether or not the air segment had a schedule change by the carrier. This tells rapid reprice that the change in the air segment was involuntary and because of a schedule change, not because the user is changing the segment. Providers: 1G/1V/1P/1S/1A
	BrandIndicator                    string                     `xml:"BrandIndicator,attr,omitempty"`                    //use="optional" //Value “B” specifies that the carrier supports Rich Content and Branding.  The Brand Indicator is only returned in the availability search response.  Provider: 1G, 1V, 1P, 1J, ACH
}

//[RQ] This describes whether the segment is determined to be a sponsored flight. The SponsoredFltInfo node will only come back for Travelport UIs and not for other customers.
type SponsoredFltInfo struct {
	SponsoredLNB uint   `xml:"SponsoredLNB,attr"` //use="required" //The line number of the sponsored flight item
	NeutralLNB   uint   `xml:"NeutralLNB,attr"`   //use="required" //The neutral line number for the flight item.
	FltKey       string `xml:"FltKey,attr"`       //use="required" //The unique identifying key for the sponsored flight.
}

//[RQ] Describes the codeshare disclosure (simple text string) or the specific operating flight information (as attributes).
type CodeshareInfo struct {
	Value                 string               `xml:",innerxml"`
	OperatingCarrier      com.TypeCarrier      `xml:"OperatingCarrier,attr,omitempty"`      //use="optional" //The actual carrier that is operating the flight.
	OperatingFlightNumber com.TypeFlightNumber `xml:"OperatingFlightNumber,attr,omitempty"` //use="optional" //The actual flight number of the carrier that is operating the flight.
}

//[RQ] Matches class of service information with availability counts. Only provided on search results.
type AirAvailInfo struct {
	BookingCodeInfo []*BookingCodeInfo   `xml:"air:BookingCodeInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	FareTokenInfo   []*FareTokenInfo     `xml:"air:FareTokenInfo,omitempty"`   //minOccurs="0" maxOccurs="unbounded"
	ProviderCode    com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`   //use="optional"
	HostTokenRef    string               `xml:"HostTokenRef,attr,omitempty"`   //use="optional"
}

//[RQ] Details Cabin class info and class of service information with availability counts. Only provided on search results and grouped by Cabin class
type BookingCodeInfo struct {
	CabinClass    string `xml:"CabinClass,attr,omitempty"`    //use="optional" //Specifies Cabin class for a group of class of services. Cabin class is not identified if it is not present.
	BookingCounts string `xml:"BookingCounts,attr,omitempty"` //use="optional" //Lists class of service and their counts for specific cabin class
}

//[RQ] Associates Fare with HostToken
type FareTokenInfo struct {
	FareInfoRef  string `xml:"FareInfoRef,attr"`  //use="required"
	HostTokenRef string `xml:"HostTokenRef,attr"` //use="required"
}

//[RQ] Specific details within a flight segment.
type FlightDetails struct {
	Connection *Connection `xml:"air:Connection,omitempty"` //minOccurs="0"
	Meals      []*string   `xml:"air:Meals,omitempty"`      // minOccurs="0" maxOccurs="unbounded"
	//InFlightServices    []InFlightServices `xml:"InFlightServices,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	Key                 com.TypeRef       `xml:"Key,attr"`                           //use="required"
	Origin              com.TypeIATACode  `xml:"Origin,attr"`                        //use="required" //The IATA location code for this origination of this entity.
	Destination         com.TypeIATACode  `xml:"Destination,attr"`                   //use="required" //The IATA location code for this destination of this entity.
	DepartureTime       string            `xml:"DepartureTime,attr,omitempty"`       //use="optional" //The date and time at which this entity departs. This does not include time zone information since it can be derived from the origin location.
	ArrivalTime         string            `xml:"ArrivalTime,attr,omitempty"`         //use="optional" //The date and time at which this entity arrives at the destination. This does not include time zone information since it can be derived from the origin location.
	FlightTime          int               `xml:"FlightTime,attr,omitempty"`          //use="optional" //Time spent (minutes) traveling in flight, including airport taxi time.
	TravelTime          int               `xml:"TravelTime,attr,omitempty"`          //use="optional" //Total time spent (minutes) traveling including flight time and ground time.
	Distance            int               `xml:"Distance,attr,omitempty"`            //use="optional" //The distance traveled. Units are specified in the parent response element.
	Equipment           air.TypeEquipment `xml:"Equipment,attr,omitempty"`           //use="optional"
	OnTimePerformance   int               `xml:"OnTimePerformance,attr,omitempty"`   //use="optional" //Represents flight on time performance as a percentage from 0 to 100
	OriginTerminal      string            `xml:"OriginTerminal,attr,omitempty"`      //use="optional"
	DestinationTerminal string            `xml:"DestinationTerminal,attr,omitempty"` //use="optional"
	AutomatedCheckin    bool              `xml:"AutomatedCheckin,attr,omitempty"`    //use="optional" default="false" //“True” indicates that the flight allows automated check-in. The default is “False”.
	//<xs:simpleType name="typeElementStatus">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:simpleType>
}

//[RQ] Flight Connection Information
type Connection struct {
	FareNote                   *FareNote `xml:"air:FareNote,omitempty"`                    //minOccurs="0"
	ChangeOfPlane              bool      `xml:"ChangeOfPlane,attr,omitempty"`              //use="optional" default="false" //Indicates the traveler must change planes between flights.
	ChangeOfTerminal           bool      `xml:"ChangeOfTerminal,attr,omitempty"`           //use="optional" default="false" //Indicates the traveler must change terminals between flights.
	ChangeOfAirport            bool      `xml:"ChangeOfAirport,attr,omitempty"`            //use="optional" default="false" //Indicates the traveler must change airports between flights.
	StopOver                   bool      `xml:"StopOver,attr,omitempty"`                   //use="optional" default="false" //Indicates that there is a significant delay between flights (usually 12 hours or more)
	MinConnectionTime          int       `xml:"MinConnectionTime,attr,omitempty"`          //use="optional" //The minimum time needed to connect between the two different destinations.
	Duration                   int       `xml:"Duration,attr,omitempty"`                   //use="optional" //The actual duration (in minutes) between flights.
	SegmentIndex               int       `xml:"SegmentIndex,attr,omitempty"`               //use="optional" //The sequential AirSegment number that this connection information applies to.
	FlightDetailsIndex         int       `xml:"FlightDetailsIndex,attr,omitempty"`         //use="optional" //The sequential FlightDetails number that this connection information applies to.
	IncludeStopOverToFareQuote string    `xml:"IncludeStopOverToFareQuote,attr,omitempty"` //use="optional" //The field determines to quote fares with or without stop overs,the values can be NoStopOver,StopOver and IgnoreSegment.
}

//[RQ] A simple textual fare note. Used within several other objects.
type FareNote struct {
	Value              string      `xml:",innerxml"`
	Key                string      `xml:"Key,attr"`                          //use="required"
	Precedence         int         `xml:"Precedence,attr,omitempty"`         //use="optional"
	NoteName           string      `xml:"NoteName,attr,omitempty"`           //use="optional"
	FareInfoMessageRef com.TypeRef `xml:"FareInfoMessageRef,attr,omitempty"` //use="optional"
	//<xs:simpleType name="typeElementStatus">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:simpleType>
}

//[RQ] A reference to a fare note from a shared list. Used to minimize xml results.
type FareNoteRef struct {
	Key string `xml:"Key,attr"` //use="required"
}

//[RQ] Reference to a complete FlightDetails from a shared list
type FlightDetailsRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RQ] Reference to a AlternateLocationDistance
type AlternateLocationDistanceRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RQ]
type RailCoachDetails struct {
	RailCoachNumber         string `xml:"RailCoachNumber,attr"`         //Rail coach number for the returned coach details.
	AvailableRailSeats      string `xml:"AvailableRailSeats,attr"`      //Number of available seats present in this rail coach.
	RailSeatMapAvailability bool   `xml:"RailSeatMapAvailability,attr"` //Indicates if seats are available in this rail coach which can be mapped.
}

//[RQ] Specify pseudo City
type PCC struct {
	OverridePCC  *comrq.OverridePCC   `xml:"com:OverridePCC,omitempty"`  //minOccurs="0"
	PointOfSale  []*comrq.PointOfSale `xml:"com:PointOfSale,omitempty"`  //minOccurs="0" maxOccurs="5"
	TicketAgency *TicketAgency        `xml:"air:TicketAgency,omitempty"` //minOccurs="0"
}

//[RQ] Fare Rules Filter if requested will return rules for requested category in the response. Applicable for providers 1G,1V,1P,1J.
type FareRulesFilterCategory struct {
	CategoryCode []string `xml:"air:CategoryCode"` //maxOccurs="10" //Fare Rules Filter category can be requested. Currently only ‘CHG’ is supported. Applicable for Providers 1G,1V,1P,1J.
}

//[RQ] This modifier will override the pseudo of the ticketing agency found in the AAT (TKAG). Used for all plating carrier validation.
type TicketAgency struct {
	ProviderCode   string `xml:"ProviderCode,attr"`   //use="required" //The code of the Provider (e.g. 1G, 1P)
	PseudoCityCode string `xml:"PseudoCityCode,attr"` //use="required" //The PCC of the host system.
}

//Reference to a complete AirItinerarySolution from a shared list
type AirItinerarySolutionRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RQ] Reference to a complete AirSegment from a shared list
type AirSegmentRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RQ] Reference to a Leg
type LegRef struct {
	Key com.TypeRef `xml:"Key,attr"`
}

//[RQ] The pricing container for an air travel itinerary
type AirPricingSolution struct {
	AirSegment              []*AirSegment              `xml:"air:AirSegment,omitempty"`              //minOccurs="0" maxOccurs="unbounded"
	AirSegmentRef           []*AirSegmentRef           `xml:"air:AirSegmentRef,omitempty"`           //minOccurs="0" maxOccurs="unbounded"
	Journey                 []*Journey                 `xml:"air:Journey,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	AirLegRefSegment        []*LegRef                  `xml:"air:LegRef,omitempty"`                  //minOccurs="0" maxOccurs="unbounded"
	AirPricingInfo          []*AirPricingInfo          `xml:"air:AirPricingInfo,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	FareNote                []*FareNote                `xml:"air:FareNote,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	FareNoteRef             []*FareNoteRef             `xml:"air:FareNoteRef,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	Connection              []*Connection              `xml:"air:Connection,omitempty"`              //minOccurs="0" maxOccurs="unbounded"
	MetaData                []*comrq.MetaData          `xml:"com:MetaData,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	AirPricingResultMessage []*comrq.ResultMessage     `xml:"com:AirPricingResultMessage,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	FeeInfo                 []*FeeInfo                 `xml:"air:FeeInfo,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	TaxInfo                 []*TaxInfo                 `xml:"air:TaxInfo,omitempty"`                 //minOccurs="0" maxOccurs="unbounded" //Itinerary level taxes
	AirItinerarySolutionRef []*AirItinerarySolutionRef `xml:"air:AirItinerarySolutionRef,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	HostToken               []*comrq.HostToken         `xml:"com:HostToken,omitempty"`               //minOccurs="0" maxOccurs="unbounded"
	OptionalServices        *OptionalServices          `xml:"air:OptionalServices,omitempty"`        //minOccurs="0"
	AvailableSSR            *AvailableSSR              `xml:"air:AvailableSSR,omitempty"`            //minOccurs="0"
	PricingDetails          *PricingDetails            `xml:"air:PricingDetails,omitempty"`          //minOccurs="0" maxOccurs="1"
	Key                     com.TypeRef                `xml:"Key,attr"`                              //use="required"
	CompleteItinerary       bool                       `xml:"CompleteItinerary,attr,omitempty"`      //use="optional" default="true" //This attribute is used to return whether complete Itinerary is present in the AirPricingSolution structure or not. If set to true means AirPricingSolution contains the result for full requested itinerary.
	QuoteDate               string                     `xml:"QuoteDate,attr,omitempty"`              //use="optional" //time.Time //This date will be equal to the date of the transaction unless the request included a modified ticket date.
	//<xs:attributeGroup ref="common:attrPrices">
	TotalPrice            com.TypeMoney `xml:"TotalPrice,attr,omitempty"`            //use="optional" //The total price for this entity including base price and all taxes.
	BasePrice             com.TypeMoney `xml:"BasePrice,attr,omitempty"`             //use="optional" //Represents the base price for this entity. This does not include any taxes or surcharges.
	ApproximateTotalPrice com.TypeMoney `xml:"ApproximateTotalPrice,attr,omitempty"` //use="optional" //The Converted total price in Default Currency for this entity including base price and all taxes.
	ApproximateBasePrice  com.TypeMoney `xml:"ApproximateBasePrice,attr,omitempty"`  //use="optional" //The Converted base price in Default Currency for this entity. This does not include any taxes or surcharges.
	EquivalentBasePrice   com.TypeMoney `xml:"EquivalentBasePrice,attr,omitempty"`   //use="optional" //Represents the base price in the related currency for this entity. This does not include any taxes or surcharges.
	Taxes                 com.TypeMoney `xml:"Taxes,attr,omitempty"`                 //use="optional" //The aggregated amount of all the taxes that are associated with this entity. See the associated TaxInfo array for a breakdown of the individual taxes.
	Fees                  com.TypeMoney `xml:"Fees,attr,omitempty"`                  //use="optional" //The aggregated amount of all the fees that are associated with this entity. See the associated FeeInfo array for a breakdown of the individual fees.
	Services              com.TypeMoney `xml:"Services,attr,omitempty"`              //use="optional" //The total cost for all optional services.
	ApproximateTaxes      com.TypeMoney `xml:"ApproximateTaxes,attr,omitempty"`      //use="optional" //The Converted tax amount in Default Currency.
	ApproximateFees       com.TypeMoney `xml:"ApproximateFees,attr,omitempty"`       //use="optional" //The Converted fee amount in Default Currency. </xs:attribute>
	//<xs:attributeGroup>
	Itinerary string `xml:"Itinerary,attr,omitempty"` //use="optional"
}

//[RQ] Information about all connecting segment list and total traveling time
type Journey struct {
	AirSegmentRef []*AirSegmentRef `xml:"air:AirSegmentRef"`         //maxOccurs="unbounded"
	TravelTime    string           `xml:"TravelTime,attr,omitempty"` //use="optional" //xs:duration格式：PnYnMnDTnHnMnS //otal traveling time that is difference between the departure time of the first segment and the arrival time of the last segments for that particular entire set of connection.
}

//[RQ] Per traveler type pricing breakdown. This will reflect the pricing for all travelers of the specified type.
type AirPricingInfo struct {
	FareInfo                   []*FareInfo                   `xml:"air:FareInfo,omitempty"`                   //minOccurs="0" maxOccurs="unbounded"
	FareStatus                 *FareStatus                   `xml:"air:FareStatus,omitempty"`                 //minOccurs="0"
	FareInfoRef                []*FareInfoRef                `xml:"air:FareInfoRef,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	BookingInfo                []*BookingInfo                `xml:"air:BookingInfo,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	TaxInfo                    []*TaxInfo                    `xml:"air:TaxInfo,omitempty"`                    //minOccurs="0" maxOccurs="unbounded"
	FareCalc                   air.FareCalc                  `xml:"air:FareCalc,omitempty"`                   //minOccurs="0"
	PassengerType              []*PassengerType              `xml:"air:PassengerType,omitempty"`              //minOccurs="0" maxOccurs="unbounded"
	BookingTravelerRef         []*comrq.BookingTravelerRef   `xml:"com:BookingTravelerRef,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	WaiverCode                 *WaiverCode                   `xml:"air:WaiverCode,omitempty"`                 //minOccurs="0"
	PaymentRef                 []*PaymentRef                 `xml:"air:PaymentRef,omitempty"`                 //minOccurs="0" maxOccurs="unbounded" //The reference to the Payment if Air Pricing is charged
	ChangePenalty              *TypeFarePenalty              `xml:"air:ChangePenalty,omitempty"`              //minOccurs="0" //The penalty (if any) to change the itinerary
	CancelPenalty              *TypeFarePenalty              `xml:"air:CancelPenalty,omitempty"`              //minOccurs="0" //The penalty (if any) to cancel the fare
	FeeInfo                    []*FeeInfo                    `xml:"air:FeeInfo,omitempty"`                    //minOccurs="0" maxOccurs="unbounded"
	Adjustment                 []*Adjustment                 `xml:"air:Adjustment,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	Yield                      []*Yield                      `xml:"air:Yield,omitempty"`                      //minOccurs="0" maxOccurs="unbounded"
	AirPricingModifiers        *AirPricingModifiers          `xml:"air:AirPricingModifiers,omitempty"`        //minOccurs="0" maxOccurs="1"
	TicketingModifiersRef      []*TicketingModifiersRef      `xml:"air:TicketingModifiersRef,omitempty"`      //minOccurs="0" maxOccurs="unbounded"
	AirSegmentPricingModifiers []*AirSegmentPricingModifiers `xml:"air:AirSegmentPricingModifiers,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	FlightOptionsList          *FlightOptionsList            `xml:"air:FlightOptionsList,omitempty"`          //minOccurs="0"
	BaggageAllowances          *BaggageAllowances            `xml:"air:BaggageAllowances,omitempty"`          //minOccurs="0" maxOccurs="1"
	FareRulesFilter            *FareRulesFilter              `xml:"air:FareRulesFilter,omitempty"`            //minOccurs="0"
	PolicyCodesList            []*PolicyCodesList            `xml:"air:PolicyCodesList,omitempty"`            //minOccurs="0" //A list of codes that indicate why an item was determined to be ‘out of policy’
	PriceChange                []*PriceChangeType            `xml:"air:PriceChange,omitempty"`                //minOccurs="0" maxOccurs="99" //Indicates a price change is found in Fare Control Manager
	ActionDetails              *ActionDetails                `xml:"air:ActionDetails,omitempty"`              //minOccurs="0"
	Key                        com.TypeRef                   `xml:"Key,attr"`                                 //use="required"
	CommandKey                 string                        `xml:"CommandKey,attr,omitempty"`                //use="optional" //The command identifier used when this is in response to an AirPricingCommand. Not used in any request processing.
	//<xs:attributeGroup ref="common:attrPrices">
	TotalPrice            com.TypeMoney `xml:"TotalPrice,attr,omitempty"`            //use="optional" //The total price for this entity including base price and all taxes.
	BasePrice             com.TypeMoney `xml:"BasePrice,attr,omitempty"`             //use="optional" //Represents the base price for this entity. This does not include any taxes or surcharges.
	ApproximateTotalPrice com.TypeMoney `xml:"ApproximateTotalPrice,attr,omitempty"` //use="optional" //The Converted total price in Default Currency for this entity including base price and all taxes.
	ApproximateBasePrice  com.TypeMoney `xml:"ApproximateBasePrice,attr,omitempty"`  //use="optional" //The Converted base price in Default Currency for this entity. This does not include any taxes or surcharges.
	EquivalentBasePrice   com.TypeMoney `xml:"EquivalentBasePrice,attr,omitempty"`   //use="optional" //Represents the base price in the related currency for this entity. This does not include any taxes or surcharges.
	Taxes                 com.TypeMoney `xml:"Taxes,attr,omitempty"`                 //use="optional" //The aggregated amount of all the taxes that are associated with this entity. See the associated TaxInfo array for a breakdown of the individual taxes.
	Fees                  com.TypeMoney `xml:"Fees,attr,omitempty"`                  //use="optional" //The aggregated amount of all the fees that are associated with this entity. See the associated FeeInfo array for a breakdown of the individual fees.
	Services              com.TypeMoney `xml:"Services,attr,omitempty"`              //use="optional" //The total cost for all optional services.
	ApproximateTaxes      com.TypeMoney `xml:"ApproximateTaxes,attr,omitempty"`      //use="optional" //The Converted tax amount in Default Currency.
	ApproximateFees       com.TypeMoney `xml:"ApproximateFees,attr,omitempty"`       //use="optional" //The Converted fee amount in Default Currency. </xs:attribute>
	//<xs:attributeGroup>
	//<xs:attributeGroup name="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional"
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional"
	//</xs:attributeGroup>
	AmountType                 com.StringLength1to32 `xml:"AmountType,attr,omitempty"`                 //use="optional" //This field displays type of payment amount when it is non-monetary. Presently available/supported value is "Flight Pass Credits".
	IncludesVAT                bool                  `xml:"IncludesVAT,attr,omitempty"`                //use="optional" //Indicates whether the Base Price includes VAT.
	ExchangeAmount             com.TypeMoney         `xml:"ExchangeAmount,attr,omitempty"`             //use="optional" //The amount to pay to cover the exchange of the fare (includes penalties).
	ForfeitAmount              com.TypeMoney         `xml:"ForfeitAmount,attr,omitempty"`              //use="optional" //The amount forfeited when the fare is exchanged.
	Refundable                 bool                  `xml:"Refundable,attr,omitempty"`                 //use="optional" //Indicates whether the fare is refundable
	Exchangeable               bool                  `xml:"Exchangeable,attr,omitempty"`               //use="optional" //Indicates whether the fare is exchangeable
	LatestTicketingTime        string                `xml:"LatestTicketingTime,attr,omitempty"`        //use="optional" //The latest date/time at which this pricing information is valid
	PricingMethod              string                `xml:"PricingMethod,attr"`                        //use="required"
	Checksum                   string                `xml:"Checksum,attr,omitempty"`                   //use="optional" //A security value used to guarantee that the pricing data sent in matches the pricing data previously returned
	ETicketability             string                `xml:"ETicketability,attr,omitempty"`             //use="optional" //The E-Ticketability of this AirPricing
	PlatingCarrier             com.TypeCarrier       `xml:"PlatingCarrier,attr,omitempty"`             //use="optional" //The Plating Carrier for this journey
	ProviderReservationInfoRef com.TypeRef           `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	AirPricingInfoGroup        int                   `xml:"AirPricingInfoGroup,attr,omitempty"`        //use="optional" //This attribute is added to support multiple store fare in Host. All AirPricingInfo with same group number will be stored together.
	TotalNetPrice              com.TypeMoney         `xml:"TotalNetPrice,attr,omitempty"`              //use="optional" //The total price of a negotiated fare.
	Ticketed                   bool                  `xml:"Ticketed,attr,omitempty"`                   //use="optional" //Indicates if the associated stored fare is ticketed or not.
	PricingType                string                `xml:"PricingType,attr,omitempty"`                //use="optional" //Indicates the Pricing Type used. The possible values are TicketRecord, StoredFare, PricingInstruction.
	TrueLastDateToTicket       string                `xml:"TrueLastDateToTicket,attr,omitempty"`       //use="optional" //This date indicates the true last date/time to ticket for the fare. This date comes from the filed fare . There is no guarantee the fare will still be available on that date or that the fare amount may change. It is merely the last date to purchase a ticket based on the carriers fare rules at the time the itinerary was quoted and stored
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	//<xs:attributeGroup name="attrPolicyMarking">
	InPolicy        bool `xml:"InPolicy,attr,omitempty"`        //use="optional" //This attribute will be used to indicate if a fare or rate has been determined to be ‘in policy’ based on the associated policy settings.
	PreferredOption bool `xml:"PreferredOption,attr,omitempty"` //use="optional" //This attribute is used to indicate if the vendors responsible for the fare or rate being returned have been determined to be ‘preferred’ based on the associated policy settings.
	//</xs:attributeGroup>
	FareCalculationInd string `xml:"FareCalculationInd,attr,omitempty"` //use="optional" //Fare calculation that was used to price the itinerary.
}

//[RQ] Information about this fare component
type FareInfo struct {
	FareTicketDesignator []*FareTicketDesignator `xml:"air:FareTicketDesignator,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	FareSurcharge        []*FareSurcharge        `xml:"air:FareSurcharge,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	AccountCode          []*comrq.AccountCode    `xml:"com:AccountCode,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	ContractCode         []*ContractCode         `xml:"air:ContractCode,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	Endorsement          []*comrq.Endorsement    `xml:"com:Endorsement,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	BaggageAllowance     *BaggageAllowance       `xml:"air:BaggageAllowance,omitempty"`     //minOccurs="0"
	FareRuleKey          *FareRuleKey            `xml:"air:FareRuleKey,omitempty"`          //minOccurs="0"
	FareRuleFailureInfo  *FareRuleFailureInfo    `xml:"air:FareRuleFailureInfo,omitempty"`  //minOccurs="0"
	FareRemarkRef        []*FareRemarkRef        `xml:"air:FareRemarkRef,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	Brand                *Brand                  `xml:"air:Brand,omitempty"`                //minOccurs="0"
	Commission           *comrq.Commission       `xml:"com:Commission,omitempty"`           //minOccurs="0" //Specifies the Commission for Agency for a particular Fare component. Apllicable Providers are 1G and 1V.
	Key                  com.TypeRef             `xml:"Key,attr"`                           //use="required"
	FareBasis            string                  `xml:"FareBasis,attr"`                     //use="required" //The fare basis code for this fare
	PassengerTypeCode    com.TypePTC             `xml:"PassengerTypeCode,attr"`             //use="required" //The PTC that is associated with this fare.
	Origin               com.TypeIATACode        `xml:"Origin,attr"`                        //use="required" //Returns the airport or city code that defines the origin market for this fare.
	Destination          com.TypeIATACode        `xml:"Destination,attr"`                   //use="required" //Returns the airport or city code that defines the destination market for this fare.
	EffectiveDate        string                  `xml:"EffectiveDate,attr"`                 //use="required" //Returns the date on which this fare was quoted
	TravelDate           string                  `xml:"TravelDate,attr,omitempty"`          //use="optional" //Returns the departure date of the first segment that uses this fare.
	DepartureDate        string                  `xml:"DepartureDate,attr,omitempty"`       //use="optional" //Returns the departure date of the first segment of the journey.
	Amount               com.TypeMoney           `xml:"Amount,attr,omitempty"`              //use="optional"
	PrivateFare          string                  `xml:"PrivateFare,attr,omitempty"`         //use="optional"
	NegotiatedFare       bool                    `xml:"NegotiatedFare,attr,omitempty"`      //use="optional" //Identifies the fare as a Negotiated Fare.
	TourCode             air.TypeTourCode        `xml:"TourCode,attr,omitempty"`            //use="optional"
	WaiverCode           string                  `xml:"WaiverCode,attr,omitempty"`          //use="optional"
	NotValidBefore       string                  `xml:"NotValidBefore,attr,omitempty"`      //use="optional" //Fare not valid before this date.
	NotValidAfter        string                  `xml:"NotValidAfter,attr,omitempty"`       //use="optional" //Fare not valid after this date.
	PseudoCityCode       com.TypePCC             `xml:"PseudoCityCode,attr,omitempty"`      //use="optional" //Provider PseudoCityCode associated with private fare.
	FareFamily           com.TypeFareFamily      `xml:"FareFamily,attr,omitempty"`          //use="optional" //An alpha-numeric string which denotes fare family. Some carriers may return this in lieu of or in addition to the CabinClass.
	PromotionalFare      bool                    `xml:"PromotionalFare,attr,omitempty"`     //use="optional" //Boolean to describe whether the Fare is Promotional fare or not.
	CarCode              air.TypeCarCode         `xml:"CarCode,attr,omitempty"`             //use="optional"
	ValueCode            air.TypeValueCode       `xml:"ValueCode,attr,omitempty"`           //use="optional"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"` //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr"`      //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	BulkTicket    bool                 `xml:"BulkTicket,attr,omitempty"`    //use="optional" //Whether the ticket can be issued as bulk for this fare. Providers supported: Worldspan and JAL
	InclusiveTour bool                 `xml:"InclusiveTour,attr,omitempty"` //use="optional" //Whether the ticket can be issued as part of included package for this fare. Providers supported: Worldspan and JAL
	Value         string               `xml:"Value,attr,omitempty"`         //use="optional" //Used in rapid reprice
	SupplierCode  com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"`  //use="optional" //Code of the provider returning this fare info
	TaxAmount     com.TypeMoney        `xml:"TaxAmount,attr,omitempty"`     //use="optional" //Currency code and value for the approximate tax amount for this fare component.
}

//[RQ] Controls and switches for a Air Search request that contains Pricing Information
type AirPricingModifiers struct {
	ProhibitedRuleCategories     *ProhibitedRuleCategories  `xml:"air:ProhibitedRuleCategories,omitempty"`      //minOccurs="0"
	AccountCodes                 *AccountCodes              `xml:"air:AccountCodes,omitempty"`                  //minOccurs="0" //Used to get negotiated pricing. Provider:ACH.
	PermittedCabins              *PermittedCabins           `xml:"air:PermittedCabins,omitempty"`               //minOccurs="0"
	ContractCodes                *ContractCodes             `xml:"air:ContractCodes,omitempty"`                 //minOccurs="0"
	ExemptTaxes                  *ExemptTaxes               `xml:"air:ExemptTaxes,omitempty"`                   //minOccurs="0"
	PenaltyFareInformation       *PenaltyFareInformation    `xml:"air:PenaltyFareInformation,omitempty"`        //minOccurs="0" //Request Fares with specific Penalty Information.
	DiscountCard                 []*comrq.DiscountCard      `xml:"com:DiscountCard,omitempty"`                  //minOccurs="0" maxOccurs="9" //Discount request for rail.
	PromoCodes                   *PromoCodes                `xml:"air:PromoCodes,omitempty"`                    //minOccurs="0"
	ManualFareAdjustment         []*ManualFareAdjustment    `xml:"air:ManualFareAdjustment,omitempty"`          //minOccurs="0" maxOccurs="unbounded" //Represents increment/discount applied manually by agent.
	PointOfSale                  *comrq.PointOfSale         `xml:"com:PointOfSale,omitempty"`                   //minOccurs="0" //User can use this node to send a specific PCC to access fares allowed only for that PCC. This node gives the capability for fare redistribution at stored fare level. As multiple UAPI AirPricingInfos (all having same AirPricingInfoGroup) can converge to a single stored fare, UAPI will map PoinOfSale information from the first available one from each group
	BrandModifiers               *BrandModifiers            `xml:"air:BrandModifiers,omitempty"`                //minOccurs="0" //Used to specify the level of branding requested.
	MultiGDSSearchIndicator      []*MultiGDSSearchIndicator `xml:"air:MultiGDSSearchIndicator,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	ProhibitMinStayFares         bool                       `xml:"ProhibitMinStayFares,attr,omitempty"`         //use="optional" default="false"
	ProhibitMaxStayFares         bool                       `xml:"ProhibitMaxStayFares,attr,omitempty"`         //use="optional" default="false"
	CurrencyType                 com.TypeCurrency           `xml:"CurrencyType,attr,omitempty"`                 //use="optional"
	ProhibitAdvancePurchaseFares bool                       `xml:"ProhibitAdvancePurchaseFares,attr,omitempty"` //use="optional" default="false"
	ProhibitNonRefundableFares   bool                       `xml:"ProhibitNonRefundableFares,attr,omitempty"`   //use="optional" default="false"
	ProhibitRestrictedFares      bool                       `xml:"ProhibitRestrictedFares,attr,omitempty"`      //use="optional" default="false"
	FaresIndicator               string                     `xml:"FaresIndicator,attr,omitempty"`               //use="optional" //Indicates whether only public fares should be returned or specific type of private fares
	FiledCurrency                com.TypeCurrency           `xml:"FiledCurrency,attr,omitempty"`                //use="optional" //Currency in which Fares/Prices will be filed if supported by the supplier else approximated to.
	PlatingCarrier               com.TypeCarrier            `xml:"PlatingCarrier,attr,omitempty"`               //use="optional" //The Plating Carrier for this journey.
	OverrideCarrier              com.TypeCarrier            `xml:"OverrideCarrier,attr,omitempty"`              //use="optional" //The Plating Carrier for this journey.
	ETicketability               string                     `xml:"ETicketability,attr,omitempty"`               //use="optional" //Request a search based on whether only E-ticketable fares are required.
	AccountCodeFaresOnly         bool                       `xml:"AccountCodeFaresOnly,attr,omitempty"`         //use="optional" //Indicates whether or not the private fares returned should be restricted to only those specific to the input account code and contract code.
	Key                          com.TypeRef                `xml:"Key,attr,omitempty"`                          //use="optional"
	ProhibitNonExchangeableFares bool                       `xml:"ProhibitNonExchangeableFares,attr,omitempty"` //use="optional" default="false"
	ForceSegmentSelect           bool                       `xml:"ForceSegmentSelect,attr,omitempty"`           //use="optional" default="false" //This indicator allows agent to force segment select option in host while selecting all air segments to store price on a PNR. This is relevent only when agent selects all air segmnets to price. if agent selects specific segments to price then this attribute will be ignored by the system. This is currently used by Worldspan only.
	InventoryRequestType         string                     `xml:"InventoryRequestType,attr,omitempty"`         //use="optional" //This allows user to make request for a particular source of inventory for pricing modifier purposes. This is currently used by Worldspan only.
	OneWayShop                   bool                       `xml:"OneWayShop,attr,omitempty"`                   //default="false" //Via this attribute one way shop can be requested. Applicable provider is 1G
	ProhibitUnbundledFareTypes   bool                       `xml:"ProhibitUnbundledFareTypes,attr,omitempty"`   //use="optional" default="false" //A "True" value wiill remove fares with EOU and ERU fare types from consideration. A "False" value is the same as no value.  Default is no value. Applicable providers:  1P/1J/1G/1V
	ReturnServices               bool                       `xml:"ReturnServices,attr,omitempty"`               //use="optional" default="true" //When set to false, ATPCO filed Optional Services will not be returned. Default is true. Provider: 1G, 1V, 1P, 1J
	ChannelId                    string                     `xml:"ChannelId,attr,omitempty"`                    //use="optional" //A Channel ID is 2 to 4 alpha-numeric characters used to activate the Search Control Console filter for a specific group of travelers being served by the agency credential.
}

//[RQ] Provides controls and switches for the Exchange process
type AirExchangeModifiers struct {
	ContractCodes              *ContractCodes      `xml:"air:ContractCodes,omitempty"`               //minOccurs="0"
	BookingDate                string              `xml:"BookingDate,attr,omitempty"`                //use="optional"
	TicketingDate              string              `xml:"TicketingDate,attr,omitempty"`              //use="optional"
	AccountCode                string              `xml:"AccountCode,attr,omitempty"`                //use="optional"
	TicketDesignator           string              `xml:"TicketDesignator,attr,omitempty"`           //use="optional"
	AllowPenaltyFares          bool                `xml:"AllowPenaltyFares,attr,omitempty"`          //use="optional" default="true"
	PrivateFaresOnly           bool                `xml:"PrivateFaresOnly,attr,omitempty"`           //use="optional" default="false"
	UniversalRecordLocatorCode com.TypeLocatorCode `xml:"UniversalRecordLocatorCode,attr,omitempty"` //use="optional" //Which UniversalRecord should this new reservation be applied to. If blank, then a new one is created.
	ProviderLocatorCode        com.TypeLocatorCode `xml:"ProviderLocatorCode,attr,omitempty"`        //use="optional" //Which Provider reservation does this reservation get added to.
	ProviderCode               string              `xml:"ProviderCode,attr,omitempty"`               //use="optional" //To be used with ProviderLocatorCode, which host the reservation being added to belongs to.
}

//[RQ]
type ContractCodes struct {
	ContractCode []*ContractCode `xml:"air:BookingCode"` //maxOccurs="unbounded"
}

//[RQ]
type ProhibitedRuleCategories struct {
	FareRuleCategory []*FareRuleCategory `xml:"air:FareRuleCategory"` //maxOccurs="unbounded"
}

//[RQ]
type AccountCodes struct {
	AccountCode []*comrq.AccountCode `xml:"com:AccountCode"` //maxOccurs="unbounded"
}

//[RQ]
type PromoCodes struct {
	PromoCode []*PromoCode `xml:"air:PromoCode"` //minOccurs="1" maxOccurs="unbounded"
}

//[RQ] Reference to a complete FareInfo from a shared list
type FareInfoRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RQ] Links segments and fares together
type BookingInfo struct {
	BookingCode             string      `xml:"BookingCode,attr"`                       //use="required"
	BookingCount            string      `xml:"BookingCount,attr,omitempty"`            //use="optional" //Seat availability of the BookingCode
	CabinClass              string      `xml:"CabinClass,attr,omitempty"`              //use="optional"
	FareInfoRef             com.TypeRef `xml:"FareInfoRef,attr"`                       //use="required"
	SegmentRef              com.TypeRef `xml:"SegmentRef,attr,omitempty"`              //use="optional"
	CouponRef               com.TypeRef `xml:"CouponRef,attr,omitempty"`               //use="optional" //The coupon to which that booking is relative (if applicable)
	AirItinerarySolutionRef com.TypeRef `xml:"AirItinerarySolutionRef,attr,omitempty"` //use="optional" //Reference to an Air Itinerary Solution
	HostTokenRef            com.TypeRef `xml:"HostTokenRef,attr,omitempty"`            //use="optional" //HostToken Reference for this segment and fare combination.
}

//[RQ] The passenger type details associated to a fare.
type PassengerType struct {
	comrq.PassengerType
	FareGuaranteeInfo *FareGuaranteeInfo `xml:"air:FareGuaranteeInfo,omitempty"` //minOccurs="0"
}

//[RQ] The information related to fare guarantee details.
type FareGuaranteeInfo struct {
	GuaranteeDate string `xml:"GuaranteeDate,attr,omitempty"` //use="optional" //The date till which the fare is guaranteed.
	GuaranteeType string `xml:"GuaranteeType,attr"`           //use="required" //Determines the status of a fare for a passenger.

}

//[RQ] Reference to one of the air reservation payments
type PaymentRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RQ] Penalty applicable on a Fare for change/ cancellation etc- expressed in both Money and Percentage.
type TypeFarePenalty struct {
	Amount     com.TypeMoney                 `xml:"com:Amount,omitempty"`     //minOccurs="0" //The penalty (if any) - expressed as the actual amount of money. Both Amount and Percentage can be present.
	Percentage com.TypePercentageWithDecimal `xml:"com:Percentage,omitempty"` //minOccurs="0" //The penalty (if any) - expressed in percentage. Both Amount and Percentage can be present.
}

//[RQ] Rule Categories to filter on.
type FareRuleCategory struct {
	Category int `xml:"Category,attr"` //use="required"
}

//[RQ] Some private fares (non-ATPCO) are secured to a contract code.
type ContractCode struct {
	Code        string `xml:"Code,attr"`        //use="required" //The 1-64 character string which uniquely identifies a Contract.
	CompanyName string `xml:"CompanyName,attr"` //Providers supported : ACH
	//<xs:attributeGroup name="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional"
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional"
	//</xs:attributeGroup>
}

//[RQ] Request tax exemption for specific tax category and/or all taxes of a specific country
type ExemptTaxes struct {
	CountryCode  []*com.TypeCountry `xml:"com:CountryCode,omitempty"`   //minOccurs="0" maxOccurs="unbounded" //Specify ISO country code for which tax exemption is requested.
	TaxCategory  []string           `xml:"air:TaxCategory,omitempty"`   //minOccurs="0" maxOccurs="unbounded" //Specify tax category for which tax exemption is requested.
	AllTaxes     bool               `xml:"AllTaxes,attr,omitempty"`     //use="optional" //Request exemption of all taxes.
	CompanyName  string             `xml:"CompanyNa,attr,omitempty"`    //use="optional" //The federal government body name must be provided in this element. This field is required by AC
	TaxTerritory string             `xml:"TaxTerritory,attr,omitempty"` //use="optional" //exemption is achieved by sending in the TaxTerritory in the tax exempt price request.
}

//[RQ]
type PenaltyFareInformation struct {
	PenaltyInfo          *TypeFarePenalty `xml:"air:PenaltyInfo,omitempty"` //minOccurs="0" //Penalty Limit if requested.
	ProhibitPenaltyFares bool             `xml:"ProhibitPenaltyFares,attr"` //use="required" //Indicates whether user wants penalty fares to be returned.
}

//[RQ] A container to specify Promotional code with Provider code and Supplier code.
type PromoCode struct {
	Code         string               `xml:"Code,attr"`         //use="required" //To be used to specify Promotional Code.
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr"` //use="required" //To be used to specify Provider Code.
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr"` //use="required" //To be used to specify Supplier Code.
}

//[RQ]
type ManualFareAdjustment struct {
	AppliedOn        string                   `xml:"AppliedOn,attr"`                  //use="required" //Represents pricing component upon which manual increment/discount to be applied. Presently supported values are Base and Total. Other is present as a future place holder but presently no request processing logic is available for value Other
	AdjustmentType   string                   `xml:"AdjustmentType,attr"`             //use="required" //Represents process used for applying manual discount/increment. Presently supported values are Flat, Percentage.
	Value            float64                  `xml:"Value,attr"`                      //use="required" //Represents value of increment/discount applied. Negative value is considered as discount whereas positive value represents increment
	PassengerRef     com.TypeRef              `xml:"PassengerRef,attr,omitempty"`     //use="optional" //Represents passenger association.
	TicketDesignator air.TypeTicketDesignator `xml:"TicketDesignator,attr,omitempty"` //use="optional" //Providers: 1p/1j
	FareType         air.TypeFareTypeCode     `xml:"FareType,attr,omitempty"`         //use="optional" //Providers: 1p/1j
}

//[RQ] Wrapper for Brand modifiers
type BrandModifiers struct {
	ModifierType string `xml:"ModifierType,attr"` //use="required" //Type of Brand modifiers. e.g FareFamilyDisplay or BasicDetailOnly.
}

//[RQ] This is the container for a set of modifiers which allow the user to perform a special kind of low fare search, depicted as flex explore, based on different parameters like Area, Zone, Country, State, Specific locations, Distance around the actual destination of the itinerary. Applicable for providers 1G,1V,1P
type FlexExploreModifiers struct {
	Destination []com.TypeIATACode `xml:"com:Destination,omitempty"` //minOccurs="0" maxOccurs="59" //List of specific destinations for performing flex explore. Applicable only with flex explore type - Destination
	Type        string             `xml:"Type,attr"`                 //use="required" //Type of flex explore to be performed
	Radius      int                `xml:"Radius,attr,omitempty"`     //use="optional" //Radius around the destination of actual itinerary in which the search would be performed. Supported only with types - DistanceInMiles and DistanceInKilometers
	GroupName   string             `xml:"GroupName,attr,omitempty"`  //Group name for a set of destinations to be searched.  Use with Type=Group. Group names are defined in the Search Control Console. Supported Providers:  1G/1V/1P
}

//[RQ] Provides the capability to group the results into differnt trip type and diversification strategies.
type Enumeration struct {
	SolutionGroup []*SolutionGroup `xml:"air:SolutionGroup"` //maxOccurs="unbounded"
}

//[RQ] Specifies the trip type and diversity of all or a subset of the result solutions.
type SolutionGroup struct {
	PermittedAccountCodes  *PermittedAccountCodes  `xml:"air:PermittedAccountCodes,omitempty"`  //minOccurs="0"
	PreferredAccountCodes  *PreferredAccountCodes  `xml:"air:PreferredAccountCodes,omitempty"`  //minOccurs="0"
	ProhibitedAccountCodes *ProhibitedAccountCodes `xml:"air:ProhibitedAccountCodes,omitempty"` //minOccurs="0"
	PermittedPointOfSales  *PermittedPointOfSales  `xml:"air:PermittedPointOfSales,omitempty"`  //minOccurs="0"
	ProhibitedPointOfSales *ProhibitedPointOfSales `xml:"air:ProhibitedPointOfSales,omitempty"` //minOccurs="0"
	Count                  int                     `xml:"Count,attr,omitempty"`                 //use="optional" //The number of solution to include in this group. If only one group specified, this can be left blank. If multiple groups specified, all counts must add up to the MaxResults of the request.
	TripType               string                  `xml:"TripType,attr,omitempty"`              //use="required" //Specifies the trip type for this group of results. Allows targeting a result set to a particular set of characterists.
	Diversification        string                  `xml:"Diversification,attr,omitempty"`       //use="optional" //Specifies the diversification of this group of results, if specified. Allows targeting a result set to ensure they contain more unique results.
	Tag                    string                  `xml:"Tag,attr,"`                            //use="optional" //An arbitrary name for this group of solutions. Will be returned with the solution for idetification.
	Primary                bool                    `xml:"Primary,attr,omitempty"`               //use="optional" default="false" //Indicates that this is a primary SolutionGroup when using alternate pricing concepts
}

//[RQ]
type PermittedAccountCodes struct {
	AccountCode []*comrq.AccountCode `xml:"com:AccountCode"` //maxOccurs="unbounded"
}

//[RQ]
type PreferredAccountCodes struct {
	AccountCode []*comrq.AccountCode `xml:"com:AccountCode"` //maxOccurs="unbounded"
}

//[RQ]
type ProhibitedAccountCodes struct {
	AccountCode []*comrq.AccountCode `xml:"com:AccountCode"` //maxOccurs="unbounded"
}

//[RQ]
type PermittedPointOfSales struct {
	PointOfSale []*comrq.PointOfSale `xml:"com:PointOfSale"` //maxOccurs="unbounded"
}

//[RQ]
type ProhibitedPointOfSales struct {
	PointOfSale []*comrq.PointOfSale `xml:"com:PointOfSale"` //maxOccurs="unbounded"
}

//[RQ] Ticket Designator used to further qualify a Fare
type FareTicketDesignator struct {
	Value air.TypeTicketDesignator `xml:"Value,attr,omitempty"` //use="optional"
}

//[RQ] A container for an Air only travel itinerary.
type AirItinerary struct {
	AirSegment       []*AirSegment       `xml:"air:AirSegment"`                 //maxOccurs="unbounded"
	HostToken        []*comrq.HostToken  `xml:"com:HostToken,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	APISRequirements []*APISRequirements `xml:"air:APISRequirements,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//[RQ] Specific details for APIS Requirements.
type APISRequirements struct {
	Document            []*Document `xml:"air:Document,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	Key                 string      `xml:"Key,attr"`                           //Unique identifier for this APIS Requirements - use this key when a single APIS Requirements is shared by multiple elements.
	Level               string      `xml:"Level,attr,omitempty"`               //use="optional" //Applicability level of the Document. Required, Supported, API_Supported or Unknown
	GenderRequired      bool        `xml:"GenderRequired,attr,omitempty"`      //use="optional"
	DateOfBirthRequired bool        `xml:"DateOfBirthRequired,attr,omitempty"` //use="optional"
	RequiredDocuments   string      `xml:"RequiredDocuments,attr,omitempty"`   //use="optional" //What are required documents for the APIS Requirement. One, FirstAndOneOther or All
	NationalityRequired bool        `xml:"NationalityRequired,attr,omitempty"` //use="optional" //Nationality of the traveler is required for booking for some suppliers.
}

//[RQ] APIS Document Details.
type Document struct {
	Sequence int    `xml:"Sequence,attr,omitempty"` //use="optional" //Sequence number for the document.
	Type     string `xml:"Type,attr,omitempty"`     //use="optional" //Type of the Document. Visa, Passport, DriverLicense etc.
	Level    string `xml:"Level,attr,omitempty"`    //use="optional" //Applicability level of the Document. Required, Supported, API_Supported or Unknown.
}

//[RQ] A containter to identify individual pricing events. A pricing result will be returned for each pricing command according to its parameters.
type AirPricingCommand struct {
	AirPricingModifiers        *AirPricingModifiers          `xml:"air:AirPricingModifiers,omitempty"`        //minOccurs="0"
	AirSegmentPricingModifiers []*AirSegmentPricingModifiers `xml:"air:AirSegmentPricingModifiers,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	CommandKey                 string                        `xml:"CommandKey,attr,omitempty"`                //use="optional" //An identifier to link the pricing responses to the pricing commands. The value passed here will be returned in the resulting AirPricingInfo(s) from this command.
	CabinClass                 string                        `xml:"CabinClass,attr,omitempty"`                //use="optional" //Specify the cabin type to price the entire itinerary in. If segment level cabin selection is required, this attribute should not be used.
}

//[RQ] Specifies modifiers that a particular segment should be priced in. If this is used, then there must be one for each AirSegment in the AirItinerary.
type AirSegmentPricingModifiers struct {
	PermittedBookingCodes        *PermittedBookingCodes `xml:"air:PermittedBookingCodes,omitempty"`         //minOccurs="0"
	AirSegmentRef                com.TypeRef            `xml:"AirSegmentRef,attr,omitempty"`                //use="optional"
	CabinClass                   string                 `xml:"CabinClass,attr,omitempty"`                   //use="optional"
	AccountCode                  string                 `xml:"AccountCode,attr,omitempty"`                  //use="optional"
	ProhibitAdvancePurchaseFares bool                   `xml:"ProhibitAdvancePurchaseFares,attr,omitempty"` //use="optional" default="false"
	ProhibitNonRefundableFares   bool                   `xml:"ProhibitNonRefundableFares,attr,omitempty"`   //use="optional" default="false"
	ProhibitPenaltyFares         bool                   `xml:"ProhibitPenaltyFares,attr,omitempty"`         //use="optional" default="false"
	FareBasisCode                string                 `xml:"FareBasisCode,attr,omitempty"`                //use="optional" //The fare basis code to be used for pricing.
	FareBreak                    string                 `xml:"FareBreak,attr,omitempty"`                    //use="optional" //Fare break point modifier to instruct Fares where it should or should not break the fare.
	ConnectionIndicator          string                 `xml:"ConnectionIndicator,attr,omitempty"`          //use="optional" //ConnectionIndicator attribute will be used to map connection indicators AvailabilityAndPricing, TurnAround and Stopover. This attribute is for Wordspan/1P only.
}

type AssociatedRemark struct {
	comrq.AssociatedRemarkWithSegmentRef
}

type PocketItineraryRemark struct {
	comrq.AssociatedRemarkWithSegmentRef
}

//Request object used to request seats automatically by seat type
type AutoSeatAssignment struct {
	SegmentRef         com.TypeRef `xml:"SegmentRef,attr,omitempty"`         //use="optional" //The segment that this assignment belongs to
	Smoking            bool        `xml:"Smoking,attr,omitempty"`            //use="optional" default="false" //Indicates that the requested seat type should be a smoking seat.
	SeatType           string      `xml:"SeatType,attr"`                     //use="required" //The type of seat that is requested
	Group              bool        `xml:"Group,attr,omitempty"`              //use="optional" default="false" //Indicates that this seat request is for group seating for all passengers. If no SegmentRef is included, group seating will be requested for all segments.
	BookingTravelerRef com.TypeRef `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //The booking traveler that this seat assignment is for. If not entered, this applies to the primary booking traveler and other passengers are adjacent.
}

//Request object used to specify a specific seat
type SpecificSeatAssignment struct {
	BookingTravelerRef com.TypeRef `xml:"BookingTravelerRef,attr"`        //use="required" //The passenger that this seat assignment is for
	SegmentRef         com.TypeRef `xml:"SegmentRef,attr"`                //use="required" //The segment that we will assign this seat on
	FlightDetailRef    com.TypeRef `xml:"FlightDetailRef,attr,omitempty"` //use="optional" //The Flight Detail ref of the AirSegment used when requesting seats on Change of Guage flights
	SeatId             string      `xml:"SeatId,attr"`                    //use="required" //The actual seat ID that is being requested. Special Characters are not supported in this field.
	RailCoachNumber    string      `xml:"RailCoachNumber,attr,omitempty"` //use="optional" //Coach number for which rail seatmap/coachmap is returned.
}

//Indicates whether public fares and/or private fares should be returned.
type MultiGDSSearchIndicator struct {
	Type                string               `xml:Type,attr,omitempty`                  //use="optional" //Indicates whether only public fares or both public and private fares should be returned or a specific type of private fares. Examples of valid values are PublicFaresOnly, PrivateFaresOnly, AirlinePrivateFaresOnly, AgencyPrivateFaresOnly, PublicandPrivateFares, and NetFaresOnly.
	ProviderCode        com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`        //use="optional"
	DefaultProvider     bool                 `xml:"DefaultProvider,attr,omitempty"`     //use="optional" //Use the value “true” if the provider is the default (primary) provider.  Use the value “false” if the provider is the alternate (secondary).  Use of this attribute requires specifically provisioned credentials.
	PrivateFareCode     string               `xml:"PrivateFareCode,attr,omitempty"`     //use="optional" //The code of the corporate private fare.  This is the same as an account code.  Use of this attribute requires specifically provisioned credentials.
	PrivateFareCodeOnly bool                 `xml:"PrivateFareCodeOnly,attr,omitempty"` //use="optional"  //Indicates whether or not the private fares returned should be restricted to only those specific to the PrivateFareCode in the previous attribute.  This has the same validation as the AccountCodeFaresOnly attribute.  Use of this attribute requires specifically provisioned credentials.
}

//Identifies the AirReservation LocatorCode within the Universal Record
type AirReservationLocatorCode struct {
	Value string      `xml:",innerxml"`
	Key   com.TypeRef `xml:"Key,attr"` //use="required"
}

//Commercially recognized product offered by an airline
type Brand struct {
	Title                   []*Title              `xml:"air:Title,omitempty"`                    //minOccurs="0" maxOccurs="2" //The additional titles associated to the brand
	Text                    []*Text               `xml:"air:Text,omitempty"`                     //minOccurs="0" maxOccurs="5" //Text associated to the brand
	ImageLocation           []*ImageLocation      `xml:"air:ImageLocation,omitempty"`            //minOccurs="0" maxOccurs="3" //Images associated to the brand
	OptionalServices        *OptionalServices     `xml:"air:OptionalServices,omitempty"`         //minOccurs="0"
	Rules                   []*Rules              `xml:"air:Rules,omitempty"`                    //minOccurs="0" maxOccurs="99" //Brand rules
	ServiceAssociations     *ServiceAssociations  `xml:"air:ServiceAssociations,omitempty"`      //minOccurs="0" //Service associated with this brand
	UpsellBrand             *UpsellBrand          `xml:"air:UpsellBrand,omitempty"`              //minOccurs="0" //The unique identifier of the Upsell brand
	ApplicableSegment       []*ApplicableSegment  `xml:"air:ApplicableSegment,omitempty"`        //minOccurs="0" maxOccurs="99"
	DefaultBrandDetail      []*DefaultBrandDetail `xml:"air:DefaultBrandDetail,omitempty"`       //minOccurs="0" maxOccurs="99" //Default brand details.
	Key                     com.TypeRef           `xml:"Key,attr"`                               //Brand Key
	BrandID                 TypeBrandId           `xml:"BrandID,attr,omitempty"`                 //The unique identifier of the brand
	Name                    string                `xml:"Name,attr,omitempty"`                    //The Title of the brand
	AirItineraryDetailsRef  com.TypeRef           `xml:"AirItineraryDetailsRef,attr,omitempty"`  //AirItinerary associated with this brand
	UpSellBrandID           TypeBrandId           `xml:"UpSellBrandID,attr,omitempty"`           //The unique identifier of the upsell brand
	BrandFound              bool                  `xml:"BrandFound,attr,omitempty"`              //Indicates whether brand for the fare was found for carrier or not
	UpSellBrandFound        bool                  `xml:"UpSellBrandFound,attr,omitempty"`        //Indicates whether upsell brand for the fare was found for carrier or not
	BrandedDetailsAvailable bool                  `xml:"BrandedDetailsAvailable,attr,omitempty"` //Indicates if full details of the brand is available
	Carrier                 com.TypeCarrier       `xml:"Carrier,attr,omitempty"`                 //use="optional"
}

//Fare Rules Filter about this fare component. Applicable Providers are 1P,1J,1G,1V.
type FareRulesFilter struct {
	Refundability       *Refundability `xml:"air:Refundability,omitempty"`        //minOccurs="0"
	LatestTicketingTime string         `xml:"LatestTicketingTime,attr,omitempty"` // minOccurs="0" xs:dateTime //For Future Use
}

//Refundability/Penalty Fare Rules about this fare component.
type Refundability struct {
	Value air.TypeRefundabilityValue `xml:"Value,attr"` //use="required" //Currently returned: FullyRefundable (1G,1V), RefundableWithPenalty (1G,1V), Refundable (1P,1J),  NonRefundable (1G,1V,1P,1J).Refundable.
}

//To support Anchor flight search contain the anchor flight details. Supported providers 1P, 1J
type TypeAnchorFlightData struct {
	AirlineCode         com.TypeCarrier      `xml:"AirlineCode,attr"`                   //use="required" //Indicates Anchor flight carrier code
	FlightNumber        com.TypeFlightNumber `xml:"FlightNumber,attr"`                  //use="required" //Indicates Anchor flight number
	ConnectionIndicator bool                 `xml:"ConnectionIndicator,attr,omitempty"` //use="optional" //Indicates that the Anchor flight has any connecting flight or not
}

//[RQ] Surcharges for a fare component
type FareSurcharge struct {
	Key        com.TypeRef   `xml:"Key,attr,omitempty"`        //use="optional"
	Type       string        `xml:"Type,attr"`                 //use="required"
	Amount     com.TypeMoney `xml:"Amount,attr"`               //use="required"
	SegmentRef com.TypeRef   `xml:"SegmentRef,attr,omitempty"` //use="optional"
	CouponRef  com.TypeRef   `xml:"CouponRef,attr,omitempty"`  //use="optional" //The coupon to which that surcharge is relative (if applicable)
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"` //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr"`      //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RQ] Free Baggage Allowance
type BaggageAllowance struct {
	NumberOfPieces string  `xml:"air:NumberOfPieces,omitempty"` //minOccurs="0"
	MaxWeight      *Weight `xml:"air:MaxWeight,omitempty"`      //minOccurs="0"
}

//[RQ]
type Weight struct {
	Value int    `xml:"Value,attr"`
	Unit  string `xml:"Unit,attr"`
}

//[RQ] The Fare Rule requested using a Key. The key is typically a provider specific string which is required to make a following Air Fare Rule Request. This Key is returned in Low Fare Shop or Air Price Response
type FareRuleKey struct {
	Value        string               `xml:",innerxml"`
	FareInfoRef  string               `xml:"FareInfoRef,attr"`  //use="required" //The Fare Component to which this Rule Key applies
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr"` //use="required"
}

//[RQ] Returns fare rule failure reason codes when fare basis code is forced.
type FareRuleFailureInfo struct {
	Reason []air.TypeFareRuleFailureInfoReason `xml:"air:Reason"` //maxOccurs="unbounded"
}

//[RQ]
type FareRemarkRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

type TextElement struct {
	Value        string `xml:",innerxml"`
	Type         string `xml:"Type,attr"`                   //use="required"
	LanguageCode string `xml:"LanguageCode,attr,omitempty"` //type="xs:language" use="optional" //ISO 639 two-character language codes are used to retrieve specific information in the requested language. For Rich Content and Branding, language codes ZH-HANT (Chinese Traditional), ZH-HANS (Chinese Simplified), FR-CA (French Canadian) and PT-BR (Portuguese Brazil) can also be used. For RCH, language codes ENGB, ENUS, DEDE, DECH can also be used. Only certain services support this attribute. Providers: ACH, RCH, 1G, 1V, 1P, 1J.
}

type ImageLocation struct {
	Value       string `xml:",innerxml"`
	Type        string `xml:"Type,attr"`        //use="required" //Type of Image Location. E.g., "Agent", "Consumer".
	ImageWidth  uint   `xml:"ImageWidth,attr"`  //use="required" //The width of the image
	ImageHeight uint   `xml:"ImageHeight,attr"` //use="required" //The height of the image
}

//A wrapper for all the information regarding each of the Optional services
type OptionalServices struct {
	OptionalServicesTotal *OptionalServicesTotal   `xml:"air:OptionalServicesTotal,omitempty"` //minOccurs="0"
	OptionalService       []*OptionalService       `xml:"air:OptionalService"`                 //maxOccurs="unbounded"
	GroupedOptionInfo     []*GroupedOptionInfo     `xml:"air:GroupedOptionInfo,omitempty"`     //minOccurs="0" maxOccurs="unbounded" //Details about an unselected or "other" option when optional services are grouped together.
	OptionalServiceRules  []*comrq.ServiceRuleType `xml:"com:OptionalServiceRules,omitempty"`  //minOccurs="0" maxOccurs="unbounded" //Holds the rules for selecting the optional service in the itinerary
}

//The total fares, fees and taxes associated with the Optional Services
type OptionalServicesTotal struct {
	TaxInfo []*TaxInfo `xml:"air:TaxInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	FeeInfo []*FeeInfo `xml:"air:FeeInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	//<xs:attributeGroup ref="common:attrPrices">
	TotalPrice            com.TypeMoney `xml:"TotalPrice,attr,omitempty"`            //use="optional" //The total price for this entity including base price and all taxes.
	BasePrice             com.TypeMoney `xml:"BasePrice,attr,omitempty"`             //use="optional" //Represents the base price for this entity. This does not include any taxes or surcharges.
	ApproximateTotalPrice com.TypeMoney `xml:"ApproximateTotalPrice,attr,omitempty"` //use="optional" //The Converted total price in Default Currency for this entity including base price and all taxes.
	ApproximateBasePrice  com.TypeMoney `xml:"ApproximateBasePrice,attr,omitempty"`  //use="optional" //The Converted base price in Default Currency for this entity. This does not include any taxes or surcharges.
	EquivalentBasePrice   com.TypeMoney `xml:"EquivalentBasePrice,attr,omitempty"`   //use="optional" //Represents the base price in the related currency for this entity. This does not include any taxes or surcharges.
	Taxes                 com.TypeMoney `xml:"Taxes,attr,omitempty"`                 //use="optional" //The aggregated amount of all the taxes that are associated with this entity. See the associated TaxInfo array for a breakdown of the individual taxes.
	Fees                  com.TypeMoney `xml:"Fees,attr,omitempty"`                  //use="optional" //The aggregated amount of all the fees that are associated with this entity. See the associated FeeInfo array for a breakdown of the individual fees.
	Services              com.TypeMoney `xml:"Services,attr,omitempty"`              //use="optional" //The total cost for all optional services.
	ApproximateTaxes      com.TypeMoney `xml:"ApproximateTaxes,attr,omitempty"`      //use="optional" //The Converted tax amount in Default Currency.
	ApproximateFees       com.TypeMoney `xml:"ApproximateFees,attr,omitempty"`       //use="optional" //The Converted fee amount in Default Currency. </xs:attribute>
	//<xs:attributeGroup>
}

//[RQ] Describes a merchandising service of a given type (e.g. Seat) that can be purchased for the indicated amount. If this service is for only a given passenger or segment the references will indicate the restrictions. If this service applies to all passenger or the entire itinerary, the references will not be present.
type OptionalService struct {
	ServiceData     []*comrq.ServiceData `xml:"com:ServiceData,omitempty"`     //minOccurs="0" maxOccurs="unbounded"
	ServiceInfo     *comrq.ServiceInfo   `xml:"com:ServiceInfo,omitempty"`     //minOccurs="0"
	Remark          []*comrq.Remark      `xml:"com:Remark,omitempty"`          //minOccurs="0" maxOccurs="unbounded" //Information regarding any specific for this service.
	TaxInfo         []*TaxInfo           `xml:"air:TaxInfo,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	FeeInfo         []*FeeInfo           `xml:"air:FeeInfo,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	EMD             *EMD                 `xml:"air:EMD,omitempty"`             //minOccurs="0" maxOccurs="1"
	BundledServices *BundledServices     `xml:"air:BundledServices,omitempty"` //minOccurs="0" maxOccurs="1"
	AdditionalInfo  *AdditionalInfo      `xml:"air:AdditionalInfo,omitempty"`  //minOccurs="0" maxOccurs="16"
	FeeApplication  *FeeApplication      `xml:"air:FeeApplication,omitempty"`  //minOccurs="0" maxOccurs="1" //Specifies how the Optional Service fee is to be applied.  The choices are Per One Way, Per Round Trip, Per Item (Per Piece), Per Travel, Per Ticket, Per 1 Kilo, Per 5 Kilos.  Provider: 1G, 1V, 1P, 1J
	Text            *Text                `xml:"air:Text,omitempty"`            //minOccurs="0" maxOccurs="4"
	PriceRange      *PriceRange          `xml:"air:PriceRange,omitempty"`      //minOccurs="0" maxOccurs="5"
	TourCode        *TourCode            `xml:"air:TourCode,omitempty"`        //minOccurs="0"
	BrandingInfo    *BrandingInfo        `xml:"air:BrandingInfo,omitempty"`    //minOccurs="0"
	Title           []*Title             `xml:"air:Title,omitempty"`           //minOccurs="0" maxOccurs="2"
	//<xs:attributeGroup name="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional"
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional"
	//</xs:attributeGroup>
	OptionalServicesRuleRef com.TypeRef                  `xml:"OptionalServicesRuleRef,attr,omitempty"` //use="optional" //UniqueID to associate a rule to the Optional Service
	Type                    com.TypeMerchandisingService `xml:"Type,attr"`                              //use="required" //Specify the type of service offered (e.g. seats, baggage, etc.)
	Confirmation            string                       `xml:"Confirmation,attr,omitempty"`            //use="optional" //Confirmation number provided by the supplier
	SecondaryType           string                       `xml:"SecondaryType,attr,omitempty"`           //use="optional" //The secondary option code type required for certain options
	PurchaseWindow          com.TypePurchaseWindow       `xml:"PurchaseWindow,attr,omitempty"`          //use="optional" //Describes when the Service is available for confirmation or purchase (e.g. Booking Only, Check-in Only, Anytime, etc.)
	Priority                int                          `xml:"Priority,attr,omitempty"`                //use="optional" //Numeric value that represents the priority order of the Service
	Available               bool                         `xml:"Available,attr,omitempty"`               //use="optional" //Boolean to describe whether the Service is available for sale or not
	Entitled                bool                         `xml:"Entitled,attr,omitempty"`                //use="optional" //Boolean to describe whether the passenger is entitled for the service without charge or not
	PerTraveler             bool                         `xml:"PerTraveler,attr,omitempty"`             //use="optional" //default="true" //Boolean to describe whether the Amount on the Service is charged per traveler.
	CreateDate              time.Time                    `xml:"CreateDate,attr,omitempty"`              //use="optional" //Timestamp when this service/offer got created.
	PaymentRef              com.TypeRef                  `xml:"PaymentRef,attr,omitempty"`              //use="optional" //Reference to a payment for merchandising services.
	ServiceStatus           string                       `xml:"ServiceStatus,attr,omitempty"`           //use="optional" //Specify the service status (e.g. active, canceled, etc.)
	Quantity                int                          `xml:"Quantity,attr,omitempty"`                //use="optional" //The number of units availed for each optional service (e.g. 2 baggage availed will be specified as 2 in quantity for optional service BAGGAGE)
	SequenceNumber          int                          `xml:"SequenceNumber,attr,omitempty"`          //use="optional" //The sequence number associated with the OptionalService
	ServiceSubCode          string                       `xml:"ServiceSubCode,attr,omitempty"`          //use="optional" //The service subcode associated with the  OptionalService
	SSRCode                 com.TypeSSRCode              `xml:"SSRCode,attr,omitempty"`                 //use="optional" //The SSR Code associated with the OptionalService
	IssuanceReason          string                       `xml:"IssuanceReason,attr,omitempty"`          //use="optional" //A one-letter code specifying the reason for  issuance of the OptionalService
	ProviderDefinedType     string                       `xml:"ProviderDefinedType,attr,omitempty"`     //use="optional" //Original Type as sent by the provider
	//<xs:attributeGroup ref="common:attrPrices">
	TotalPrice            com.TypeMoney `xml:"TotalPrice,attr,omitempty"`            //use="optional" //The total price for this entity including base price and all taxes.
	BasePrice             com.TypeMoney `xml:"BasePrice,attr,omitempty"`             //use="optional" //Represents the base price for this entity. This does not include any taxes or surcharges.
	ApproximateTotalPrice com.TypeMoney `xml:"ApproximateTotalPrice,attr,omitempty"` //use="optional" //The Converted total price in Default Currency for this entity including base price and all taxes.
	ApproximateBasePrice  com.TypeMoney `xml:"ApproximateBasePrice,attr,omitempty"`  //use="optional" //The Converted base price in Default Currency for this entity. This does not include any taxes or surcharges.
	EquivalentBasePrice   com.TypeMoney `xml:"EquivalentBasePrice,attr,omitempty"`   //use="optional" //Represents the base price in the related currency for this entity. This does not include any taxes or surcharges.
	Taxes                 com.TypeMoney `xml:"Taxes,attr,omitempty"`                 //use="optional" //The aggregated amount of all the taxes that are associated with this entity. See the associated TaxInfo array for a breakdown of the individual taxes.
	Fees                  com.TypeMoney `xml:"Fees,attr,omitempty"`                  //use="optional" //The aggregated amount of all the fees that are associated with this entity. See the associated FeeInfo array for a breakdown of the individual fees.
	Services              com.TypeMoney `xml:"Services,attr,omitempty"`              //use="optional" //The total cost for all optional services.
	ApproximateTaxes      com.TypeMoney `xml:"ApproximateTaxes,attr,omitempty"`      //use="optional" //The Converted tax amount in Default Currency.
	ApproximateFees       com.TypeMoney `xml:"ApproximateFees,attr,omitempty"`       //use="optional" //The Converted fee amount in Default Currency. </xs:attribute>
	//<xs:attributeGroup>
	Key                  com.TypeRef             `xml:"Key,attr,omitempty"`                  //use="optional"
	AssessIndicator      air.TypeAssessIndicator `xml:"AssessIndicator,attr,omitempty"`      //use="optional" //Indicates whether price is assessed by mileage or currency or both
	Mileage              int                     `xml:"Mileage,attr,omitempty"`              //use="optional" //Indicates mileage fee/amount
	ApplicableFFLevel    int                     `xml:"ApplicableFFLevel,attr,omitempty"`    //use="optional" //Numerical value of the loyalty card level for which this service is available.
	Private              bool                    `xml:"Private,attr,omitempty"`              //use="optional" //Describes if service is private or not.
	SSRFreeText          com.TypeSSRFreeText     `xml:"SSRFreeText,attr,omitempty"`          //use="optional" //Certain SSR types sent in OptionalService SSRCode require a free text message. For example, PETC Pet in Cabin.
	IsPricingApproximate bool                    `xml:"IsPricingApproximate,attr,omitempty"` //use="optional" //When set to True indicates that the pricing returned is approximate. Supported providers are MCH/ACH
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"` //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr"`      //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	Chargeable                 string        `xml:"Chargeable,attr,omitempty"`                 //use="optional" //Indicates if the optional service is not offered, is available for a charge, or is included in the brand
	InclusiveOfTax             bool          `xml:"InclusiveOfTax,attr,omitempty"`             //use="optional" //Identifies if the service was filed with a fee that is inclusive of tax.
	InterlineSettlementAllowed bool          `xml:"InterlineSettlementAllowed,attr,omitempty"` //use="optional" //Identifies if the interline settlement is allowed in service .
	GeographySpecification     string        `xml:"GeographySpecification,attr,omitempty"`     //use="optional" //Sector, Portion, Journey.
	ExcessWeightRate           string        `xml:"ExcessWeightRate,attr,omitempty"`           //use="optional" //The cost of the bag per unit weight.
	Source                     string        `xml:"Source,attr,omitempty"`                     //use="optional" //The Source of the optional service. The source can be ACH, MCE, or MCH.
	ViewableOnly               bool          `xml:"ViewableOnly,attr,omitempty"`               //use="optional" //Describes if the OptionalService is viewable only or not. If viewable only then the service cannot be sold.
	DisplayText                string        `xml:"DisplayText,attr,omitempty"`                //use="optional" //Title of the Optional Service.  Provider: ACH
	WeightInExcess             string        `xml:"WeightInExcess,attr,omitempty"`             //use="optional" //The excess weight of a bag. Providers: 1G, 1V, 1P, 1J
	TotalWeight                string        `xml:"TotalWeight,attr,omitempty"`                //use="optional" //The total weight of a bag. Providers: 1G, 1V, 1P, 1J
	BaggageUnitPrice           com.TypeMoney `xml:"BaggageUnitPrice,attr,omitempty"`           //use="optional" //The per unit price of baggage. Providers: 1G, 1V, 1P, 1J
	FirstPiece                 int           `xml:"FirstPiece,attr,omitempty"`                 //use="optional" //Indicates the minimum occurrence of excess baggage.Provider: 1G, 1V, 1P, 1J.
	LastPiece                  int           `xml:"LastPiece,attr,omitempty"`                  //use="optional" //Indicates the maximum occurrence of excess baggage.Provider: 1G, 1V, 1P, 1J.
	Restricted                 bool          `xml:"Restricted,attr,omitempty"`                 //use="optional" default="false" //When set to “true”, the Optional Service is restricted by an embargo. Provider: 1G, 1V, 1P, 1J
	IsRepriceRequired          bool          `xml:"IsRepriceRequired,attr,omitempty"`          //default="false" //When set to “true”, the Optional Service must be re-priced. Provider: 1G, 1V, 1P, 1J
	BookedQuantity             string        `xml:"BookedQuantity,attr,omitempty"`             //use="optional" //Indicates the Optional Service quantity already booked. Provider: 1G, 1V, 1P, 1J
	Group                      string        `xml:"Group,attr,omitempty"`                      //use="optional" //Associates Optional Services with the same ServiceSub Code, Air Segment, Passenger, and EMD Associated Item. Provider:1G, 1V, 1P, 1J
	PseudoCityCode             com.TypePCC   `xml:"PseudoCityCode,attr,omitempty"`             //use="optional" //The PCC or SID that booked the Optional Service.
}

//Container to display the optional services which are coupled by business rules.
type GroupedOptionInfo struct {
	GroupedOption []*GroupedOption `xml:"air:GroupedOption"` //minOccurs="2" maxOccurs="unbounded"
}

type GroupedOption struct {
	OptionalServiceRef com.TypeRef `xml:"OptionalServiceRef,attr"` //use="required" //Reference to a optionalService which is paired with other optional services in the parent PairedOptions element.
}

type ServiceAssociations struct {
	ApplicableSegment []*ApplicableSegment `xml:"air:ApplicableSegment"` //maxOccurs="unbounded" //Applicable air segment associated with this brand.
}

//Upsell brand reference
type UpsellBrand struct {
	FareBasis   string `xml:"FareBasis,attr,omitempty"`   //use="optional"
	FareInfoRef string `xml:"FareInfoRef,attr,omitempty"` //use="optional"
}

type ApplicableSegment struct {
	ResponseMessage    *comrq.ResponseMessage `xml:"com:ResponseMessage,omitempty"`    //minOccurs="0"
	OptionalServiceRef *OptionalServiceRef    `xml:"air:OptionalServiceRef,omitempty"` //minOccurs="0"
	Key                com.TypeRef            `xml:"Key,attr"`                         //Applicable air segment key
}

//[RQ] A container for EMD elements related to an OptionalService
type EMD struct {
	FulfillmentType             int              `xml:"FulfillmentType,attr,omitempty"`             //use="optional" //A one digit code specifying how the service must be fulfilled. See FulfillmentTypeDescription for the description of this value.
	FulfillmentTypeDescription  string           `xml:"FulfillmentTypeDescription,attr,omitempty"`  //use="optional" //EMD description.
	AssociatedItem              string           `xml:"AssociatedItem,attr,omitempty"`              //use="optional" //The type of Optional Service.  The choices are Flight, Ticket, Merchandising, Rule Buster, Allowance, Chargeable Baggage, Carry On Baggage Allowance, Prepaid Baggage.  Provider: 1G, 1V, 1P, 1J
	AvailabilityChargeIndicator string           `xml:"AvailabilityChargeIndicator,attr,omitempty"` //use="optional" //A one-letter code specifying whether the service is available or if there is a charge associated with it. X = Service not available F = No charge for service (free) and an EMD is not issued to reflect free service E = No charge for service (free) and an EMD is issued to reflect the free service. G = No charge for service (free), booking is not required and an EMD is not issued to reflect free service H = No charge for service (free), booking is not required, and an EMD is issued to reflect the free service. Blank = No application. Charges apply according to the data in the Service Fee fields.
	RefundReissueIndicator      string           `xml:"RefundReissueIndicator,attr,omitempty"`      //use="optional"
	Commissionable              bool             `xml:"Commissionable,attr,omitempty"`              //use="optional" //True/False value to whether or not the service is comissionable.
	MileageIndicator            bool             `xml:"MileageIndicator,attr,omitempty"`            //use="optional" //True/False value to whether or not the service has miles.
	Location                    com.TypeIATACode `xml:"Location,attr,omitempty"`                    //use="optional" //3 letter location code where the service will be availed.
	Date                        string           `xml:"Date,attr,omitempty"`                        //use="optional" //The date at which the service will be used.
	Booking                     string           `xml:"Booking,attr,omitempty"`                     //use="optional" //Holds the booking description for the service, e.g., SSR.
	DisplayCategory             string           `xml:"DisplayCategory,attr,omitempty"`             //use="optional" //Describes when the service should be displayed.
	Reusable                    bool             `xml:"Reusable,attr,omitempty"`                    //use="optional" //Identifies if the service can be re-used towards a future purchase.
}

//[RQ]
type BundledServices struct {
	BundledService []*BundledService `xml:"air:BundledService,omitempty"` //minOccurs="0" maxOccurs="16
}

//[RQ] Displays the services bundled together
type BundledService struct {
	Carrier        com.TypeCarrier `xml:"Carrier,attr,omitempty"`        //use="optional" //Carrier the service is applicable.
	CarrierSubCode bool            `xml:"CarrierSubCode,attr,omitempty"` //use="optional" //Carrier sub code. True means the carrier used their own sub code. False means the carrier used an ATPCO sub code
	ServiceType    string          `xml:"ServiceType,attr,omitempty"`    //use="optional" //The type of service or what the service is used for, e.g. F type is flight type, meaning the service is used on a flight
	ServiceSubCode string          `xml:"ServiceSubCode,attr,omitempty"` //use="optional" //The sub code of the service, e.g. OAA for Pre paid baggage
	Name           string          `xml:"Name,attr,omitempty"`           //use="optional" //Name of the bundled service.
	Booking        string          `xml:"Booking,attr,omitempty"`        //Booking method for the bundled service, e..g SSR.
	Occurrence     uint            `xml:"Occurrence,attr,omitempty"`     //use="optional" //How many of the service are included in the bundled service.
}

//Displays additional text about the service.
type AdditionalInfo struct {
	Category string `xml:"Category,attr"` //use="require //The category code is the code the AdditionalInfo text came from, e.g. S5 or S7.
}

//[RQ] Describes how the fees are to be applied.
type FeeApplication struct {
	Value string `xml:",innerxml"`
	Code  string `xml:"Code,attr,omitempty"` //use="optional //The code associated to the fee application. The  choices are: 1, 2, 3, 4, 5, K, F
}

//[RQ]
type PriceRange struct {
	DefaultCurrency bool          `xml:"DefaultCurrency,attr"` //Indicates if the currency code of StartPrice / EndPrice is the default currency code
	StartPrice      com.TypeMoney `xml:"StartPrice,attr"`      //Price range start value
	EndPrice        com.TypeMoney `xml:"EndPrice,attr"`        //Price range end value
}

//[RQ] Tour Code Fare Basis
type TourCode struct {
	Value air.TypeTourCode `xml:"Value,attr"` //use="required"
}

//[RQ] Branding information for the Ancillary Service.  Returned in Seat Map only.  Providers: 1G, 1V, 1P, 1J, ACH
type BrandingInfo struct {
	PriceRange          []*PriceRange       `xml:"air:PriceRange,omitempty"`           //minOccurs="0" maxOccurs="5" //The price range of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
	Text                []*Text             `xml:"air:Text,omitempty"`                 //minOccurs="0" maxOccurs="99"
	Title               []*Title            `xml:"air:Title,omitempty"`                //minOccurs="0" maxOccurs="2" //The additional titles associated to the brand or optional service. Providers: ACH, 1G, 1V, 1P, 1J
	ImageLocation       []*ImageLocation    `xml:"air:ImageLocation,omitempty"`        //minOccurs="0" maxOccurs="3"
	ServiceGroup        *ServiceGroup       `xml:"air:ServiceGroup,omitempty"`         //minOccurs="0"
	AirSegmentRef       []*comrq.SegmentRef `xml:"com:AirSegmentRef"`                  //maxOccurs="99" //Specifies the AirSegment the branding information is for. Providers: ACH, 1G, 1V, 1P, 1J
	ExternalServiceName string              `xml:"ExternalServiceName,attr,omitempty"` //use="optional" //The external name of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
	ServiceType         string              `xml:"ServiceType,attr,omitempty"`         //use="optional" //The type of Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
	CommercialName      string              `xml:"CommercialName,attr"`                //use="required" //The commercial name of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
	Chargeable          string              `xml:"Chargeable,attr,omitempty"`          //use="optional" //Indicates if the optional service is not offered, is available for a charge, or is included in the brand.  Providers: 1G, 1V, 1P, 1J, ACH
}

//[RQ] The Service Group of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
type ServiceGroup struct {
	ServiceSubGroup []*ServiceSubGroup `xml:"air:ServiceSubGroup,omitempty"` //minOccurs="0" maxOccurs="15"
	Code            string             `xml:"Code,attr"`                     //use="required" //The Service Group Code of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
}

//[RQ] The Service Sub Group of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
type ServiceSubGroup struct {
	Code string `xml:"Code,attr"` //The Service Sub Group Code of the Ancillary Service.  Providers: 1G, 1V, 1P, 1J, ACH
}

type Rules struct {
	RulesText string `xml:"air:RulesText,omitempty"` //minOccurs="0" //Rules text
}

//Denotes the status of a particular fare.
type FareStatus struct {
	FareStatusFailureInfo *FareStatusFailureInfo `xml:"air:FareStatusFailureInfo,omitempty"` //minOccurs="0"
	Code                  string                 `xml:"Code,attr"`                           //use="required" //The status of the fare.
}

//Denotes the failure reason of a particular fare.
type FareStatusFailureInfo struct {
	Code   string `xml:"Code,attr"`             //use="required"  //The failure code of the fare.
	Reason string `xml:"Reason,attr,omitempty"` //use="optional" //The reason for the failure.
}

//Waiver code to override fare validations
type WaiverCode struct {
	TourCode         air.TypeTourCode         `xml:"TourCode,attr,omitempty"`         //use="optional"
	TicketDesignator air.TypeTicketDesignator `xml:"TicketDesignator,attr,omitempty"` //use="optional"
	Endorsement      string                   `xml:"Endorsement,attr,omitempty"`      //use="optional" //Endorsement. Size can be up to 100 characters
}

//An indentifier which indentifies adjustment made on original pricing. It can a flat amount or percentage of original price. The value of Amount/Percent can be negetive. Negative value implies a discount.
type Adjustment struct {
	Amount                        com.TypeMoney `xml:"Amount,attr,omitempty"`                        //xs:choice //Implies a flat amount to be adjusted. Negetive value implies a discount.
	Percent                       float32       `xml:"Percent,attr,omitempty"`                       //xs:choice //Implies an adjustment to be made on original price. Negetive value implies a discount.
	AdjustedTotalPrice            com.TypeMoney `xml:"AdjustedTotalPrice,attr"`                      //use="required" //The adjusted price after applying adjustment on Total price
	ApproximateAdjustedTotalPrice com.TypeMoney `xml:"ApproximateAdjustedTotalPrice,attr,omitempty"` //use="optional" //The Converted adjusted total price in Default Currency for this entity.
	BookingTravelerRef            com.TypeRef   `xml:"BookingTravelerRef,attr,omitempty"`            //use="optional" //Reference to a booking traveler for which adjustment is applied.
}

//An identifier which identifies yield made on original pricing. It can be a flat amount of original price. The value of Amount can be negative. Negative value implies a discount.
type Yield struct {
	Amount             com.TypeMoney `xml:"AdjustedTotalPrice,attr,omitempty"` //use="optional" //Yield per passenger level in Default Currency for this entity.
	BookingTravelerRef com.TypeRef   `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //Reference to a booking traveler for which Yield is applied.
}

//Reference to a shared list of Ticketing Modifers
type TicketingModifiersRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//List of segment and fare available for the search air leg.
type Option struct {
	BookingInfo []*BookingInfo `xml:"air:BookingInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	Connection  []*Connection  `xml:"air:Connection,omitempty"`  //minOccurs="0" maxOccurs="unbounded"
	Key         com.TypeRef    `xml:"Key,attr"`                  //use="required"
	TravelTime  air.Duration   `xml:"TravelTime,attr,omitempty"` //use="optional" "xs:duration" //Total traveling time that is difference between the departure time of the first segment and the arrival time of the last segments for that particular entire set of connection.
}

//List of Options available for any search air leg.
type FlightOption struct {
	Option      []*Option        `xml:"air:Option"`            //minOccurs="1" maxOccurs="unbounded" //List of BookingInfo available for the search air leg.
	LegRef      com.TypeRef      `xml:"LegRef,attr,omitempty"` //use="optional" //Identifies the Leg Reference for this Flight Option.
	Origin      com.TypeIATACode `xml:"Origin,attr"`           //use="required" //The IATA location code for this origination of this entity.
	Destination com.TypeIATACode `xml:"Destination,attr"`      //use="required" //The IATA location code for this destination of this entity.
}

//List of Flight Options for the itinerary.
type FlightOptionsList struct {
	FlightOption []*FlightOption `xml:"air:FlightOption,omitempty"` // minOccurs="0" maxOccurs="unbounded"
}

//This contains common elements that are used for Baggage Allowance info, carry-on allowance info and embargo Info. Supported providers are 1V/1G/1P/1J
type BaseBaggageAllowanceInfo struct {
	URLInfo     []*URLInfo       `xml:"air:URLInfo,omitempty"`      //minOccurs="0" maxOccurs="unbounded" //Contains the text and URL information as published by carrier.
	TextInfo    []*TextInfo      `xml:"air:TextInfo,omitempty"`     //minOccurs="0" maxOccurs="unbounded" //Text information as published by carrier.
	Origin      com.TypeIATACode `xml:"Origin,attr,omitempty"`      //use="optional"
	Destination com.TypeIATACode `xml:"Destination,attr,omitempty"` //use="optional"
	Carrier     com.TypeCarrier  `xml:"Carrier,attr,omitempty"`     //use="optional"
}

//Details of Baggage allowance
type BaggageAllowances struct {
	BaggageAllowanceInfo []*BaggageAllowanceInfo `xml:"air:BaggageAllowanceInfo"`           //minOccurs="1" maxOccurs="unbounded"
	CarryOnAllowanceInfo []*CarryOnAllowanceInfo `xml:"air:CarryOnAllowanceInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	EmbargoInfo          []*EmbargoInfo          `xml:"air:EmbargoInfo,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
}

//Information related to Baggage allowance like URL,Height,Weight etc.
type BaggageAllowanceInfo struct {
	BaseBaggageAllowanceInfo
	BagDetails   []*BagDetails `xml:"air:BagDetails,omitempty"`    //minOccurs="0" maxOccurs="unbounded"
	TravelerType com.TypePTC   `xml:"TravelerType,attr,omitempty"` //use="optional"
	FareInfoRef  com.TypeRef   `xml:"FareInfoRef,attr,omitempty"`  //use="optional"
}

//Information related to Carry-On allowance like URL, pricing etc
type CarryOnAllowanceInfo struct {
	BaseBaggageAllowanceInfo
	CarryOnDetails []*CarryOnDetails `xml:"air:CarryOnDetails,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//Information related to Carry-On Bag details .
type CarryOnDetails struct {
	BaggageRestriction    []*BaggageRestriction `xml:"air:BaggageRestriction",omitempty"` //minOccurs="0" maxOccurs="99"
	ApplicableCarryOnBags string                `xml:"ApplicableCarryOnBags,attr"`        //pplicable Carry-On baggage "First", "Second", "Third" etc
	BasePrice             com.TypeMoney         `xml:"BasePrice,attr"`
	ApproximateBasePrice  com.TypeMoney         `xml:"ApproximateBasePrice,attr"`
	Taxes                 com.TypeMoney         `xml:"Taxes,attr"`
	TotalPrice            com.TypeMoney         `xml:"TotalPrice,attr"`
	ApproximateTotalPrice com.TypeMoney         `xml:"ApproximateTotalPrice,attr"`
}

//Contains the text and URL of baggage as published by carrier.
type URLInfo struct {
	Text []com.TypeGeneralText `xml:"com:Text,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	URL  []string              `xml:"air:URL,omitempty"`  //minOccurs="0" maxOccurs="unbounded" //xs:anyURI
}

//Information on baggage as published by carrier.
type TextInfo struct {
	Text  []com.TypeGeneralText `xml:"com:Text,omitempty"`   //minOccurs="0" maxOccurs="unbounded"
	Title string                `xml:"Title,attr,omitempty"` //use="optional"
}

//Information related to Bag details .
type BagDetails struct {
	BaggageRestriction    []*BaggageRestriction `xml:"air:BaggageRestriction,omitempty"`     //minOccurs="0" maxOccurs="unbounded"
	AvailableDiscount     []*AvailableDiscount  `xml:"air:AvailableDiscount,omitempty"`      //minOccurs="0" maxOccurs="unbounded"
	ApplicableBags        string                `xml:"ApplicableBags,attr"`                  //use="required" //Applicable baggage like Carry-On,1st Check-in,2nd Check -in etc.
	BasePrice             com.TypeMoney         `xml:"BasePrice,attr,omitempty"`             //use="optional"
	ApproximateBasePrice  com.TypeMoney         `xml:"ApproximateBasePrice,attr,omitempty"`  //use="optional"
	Taxes                 com.TypeMoney         `xml:"Taxes,attr,omitempty"`                 //use="optional"
	TotalPrice            com.TypeMoney         `xml:"TotalPrice,attr,omitempty"`            //use="optional"
	ApproximateTotalPrice com.TypeMoney         `xml:"ApproximateTotalPrice,attr,omitempty"` //use="optional"
}

//Information related to  Baggage restriction rules .
type BaggageRestriction struct {
	Dimension []*Dimension   `xml:"air:Dimension,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	MaxWeight *UnitOfMeasure `xml:"air:MaxWeight,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	TextInfo  []*TextInfo    `xml:"air:TextInfo,omitempty"`  //minOccurs="0" maxOccurs="unbounded"
}

//Information related to Length,Height,Width of a baggage.
type Dimension struct {
	UnitOfMeasure
	Type string `xml:"type,attr,omitempty"` //use="optional" //Type denotes Length,Height,Width of a baggage.
}

//
type UnitOfMeasure struct {
	Value float32 `xml:"Value,attr"`
	Unit  string  `xml:"Unit,attr"` //Unit values would be lb,Lb,kg etc.
}

type AvailableDiscount struct {
	LoyaltyProgram    []*comrq.LoyaltyProgram       `xml:"com:LoyaltyProgram,omitempty"`     //minOccurs="0" maxOccurs="unbounded" //Customer Loyalty Program Detail.
	Amount            com.TypeMoney                 `xml:"Amount,attr,omitempty"`            //use="optional"
	Percent           com.TypePercentageWithDecimal `xml:"Percent,attr,omitempty"`           //use="optional"
	Description       string                        `xml:"Description,attr,omitempty"`       //use="optional"
	DiscountQualifier string                        `xml:"DiscountQualifier,attr,omitempty"` //use="optional"
}

//A wrapper for all the information regarding each of the available SSR
type AvailableSSR struct {
	SSR                 []*comrq.SSR                 `xml:"com:SSR,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	SSRRules            []*comrq.ServiceRuleType     `xml:"com:SSRRules,omitempty"`            //minOccurs="0" maxOccurs="unbounded" //Holds the rules for selecting the SSR in the itinerary
	IndustryStandardSSR []*comrq.IndustryStandardSSR `xml:"com:IndustryStandardSSR,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//Used for rapid reprice. This is a response element.  Additional information about how pricing was obtain, messages, etc.  Providers: 1G/1V/1P/1S/1A
type PricingDetails struct {
	AdvisoryMessage        []string         `xml:"air:AdvisoryMessage,omitempty"`         //minOccurs="0" maxOccurs="unbounded" //Advisory messages returned from the host.
	EndorsementText        []string         `xml:"air:EndorsementText,omitempty"`         //minOccurs="0" maxOccurs="unbounded" //Endorsement text returned from the host.
	WaiverText             string           `xml:"air:WaiverText,omitempty"`              //minOccurs="0" maxOccurs="1" //Waiver text returned from the host.
	LowFarePricing         bool             `xml:"LowFarePricing,attr,omitempty"`         //use="optional" default="false" //This tells if Low Fare Finder was used.
	LowFareFound           bool             `xml:"LowFareFound,attr,omitempty"`           //use="optional" default="false" //This tells if the lowest fare was found.
	PenaltyApplies         bool             `xml:"PenaltyApplies,attr,omitempty"`         //use="optional" default="false" //This tells if penalties apply.
	DiscountApplies        bool             `xml:"DiscountApplies,attr,omitempty"`        //use="optional" default="false"  //This tells if a discount applies.
	ItineraryType          string           `xml:"ItineraryType,attr,omitempty"`          //use="optional" //Values allowed are International or Domestic. This tells if the itinerary is international or domestic.
	ValidatingVendorCode   com.TypeCarrier  `xml:"ValidatingVendorCode,attr,omitempty"`   //use="optional" //The vendor code of the validating carrier.
	ForTicketingOnDate     string           `xml:"ForTicketingOnDate,attr,omitempty"`     //use="optional" //type="xs:date" //The ticketing date of the itinerary.
	LastDateToTicket       string           `xml:"LastDateToTicket,attr,omitempty"`       //use="optional" //type="xs:date" //The last date to issue the ticket.
	FormOfRefund           string           `xml:"FormOfRefund,attr,omitempty"`           //use="optional" //How the refund will be issued. Values will be MCO or FormOfPayment
	AccountCode            string           `xml:"AccountCode,attr,omitempty"`            //use="optional"
	BankersSellingRate     float32          `xml:"BankersSellingRate,attr,omitempty"`     //use="optional" //The selling rate at time of quote.
	PricingType            string           `xml:"PricingType,attr,omitempty"`            //use="optional" //Ties with the RepricingModifiers sent in the request and tells how the itinerary was priced.
	ConversionRate         float32          `xml:"ConversionRate,attr,omitempty"`         //use="optional" //The conversion rate at the time of quote.
	RateOfExchange         float32          `xml:"RateOfExchange,attr,omitempty"`         //use="optional" //The exchange rate at time of quote.
	OriginalTicketCurrency com.TypeCurrency `xml:"OriginalTicketCurrency,attr,omitempty"` //use="optional" //The currency of the original ticket.
}

//Applicable air segment.
type DefaultBrandDetail struct {
	Text              []*Text              `xml:"air:Text,omitempty"`              //minOccurs="0" maxOccurs="4" //Text associated to the brand
	ImageLocation     []*ImageLocation     `xml:"air:ImageLocation,omitempty"`     //minOccurs="0" maxOccurs="3" //Images associated to the brand
	ApplicableSegment []*ApplicableSegment `xml:"air:ApplicableSegment,omitempty"` //minOccurs="0" maxOccurs="99" //Defines for which air segment the brand is applicable.
	BrandID           air.TypeBrandId      `xml:"BrandID,attr,omitempty"`          //use="optional" //The unique identifier of the brand
}

type PolicyCodesList struct {
	PolicyCode []com.TypePolicyCode `xml:"PolicyCode"` //minOccurs="1" maxOccurs="10" //A code that indicates why an item was determined to be ‘out of policy’.
}

//AirPricing TicketingModifier information - used to associate Ticketing Modifiers with one or more AirPricingInfos/ProviderReservationInfo
type AirPricingTicketingModifiers struct {
	AirPricingInfoRef []*AirPricingInfoRef `xml:"air:AirPricingInfoRef,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	//TicketingModifiers *TicketingModifiers  `xml:"air:TicketingModifiers"`
}

//Reference to a AirPricing from a shared list
type AirPricingInfoRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Indicates a price change is found in Fare Control Manager
type PriceChangeType struct {
	Value      string `xml:",innerxml"`
	Amount     string `xml:"Amount,attr"`               //use="required" //Contains the currency and amount information. Assume the amount is added unless a hyphen is present to indicate subtraction.
	Carrier    string `xml:"Carrier,attr,omitempty"`    //use="optional" //Contains carrier code information
	SegmentRef string `xml:"SegmentRef,attr,omitempty"` //use="optional" //Contains segment reference information
}

//SplitTicketingSearch is optional. Used to return both One-Way and Roundtrip fares in a single search response.
//Applicable to 1G, 1V, 1P only, the price points results path, and a simple roundtrip search only.
//Cannot be used in combination with Flex options.
type SplitTicketingSearch struct {
	RoundTrip int `xml:"RoundTrip,attr,omitempty"` //use="optional" //Percentage of Roundtrip price points to be returned in the search response. This should be an even number. The One-Way price points returned in the response would be evenly distributed between the outbound and the inbound.
}

//Information related to the storing of the fare: Agent, Date and Action for Provider: 1P/1J
type ActionDetails struct {
	PseudoCityCode com.TypePCC `xml:"PseudoCityCode,attr,omitempty"` //use="optional" //PCC in the host of the agent who stored the fare for Provider: 1P/1J
	AgentSine      string      `xml:"AgentSine,attr,omitempty"`      //use="optional" //The sign in of the user who stored the fare for Provider: 1P/1J
	EventDate      string      `xml:"EventDate,attr,omitempty"`      //type="xs:date" use="optional" //Date at which the fare was stored for Provider: 1P/1J
	EventTime      string      `xml:"EventTime,attr,omitempty"`      //type="xs:time" use="optional" //Time at which the fare was stored for Provider: 1P/1J
	Text           string      `xml:"Text,attr,omitempty"`           //use="optional" //The type of action the agent performed for Provider: 1P/1J
}
