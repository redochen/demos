package response

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
)

//Container for error data when there is an application error.
type ErrorInfo struct {
	Code           string   `xml:"Code,omitempty"`
	Service        string   `xml:"Service,omitempty"`
	Type           string   `xml:"Type,omitempty"`
	Description    string   `xml:"Description,omitempty"`
	TransactionId  string   `xml:"TransactionId,omitempty"`
	TraceId        string   `xml:"TraceId,omitempty"`        //minOccurs="0"
	CommandHistory string   `xml:"CommandHistory,omitempty"` //minOccurs="0"
	Auxdata        *Auxdata `xml:"Auxdata,omitempty"`        //minOccurs="0"
	StackTrace     string   `xml:"StackTrace,omitempty"`     //minOccurs="0"
}

type Auxdata struct {
	Entry []*Entry `xml:"Entry,omitempty"` //maxOccurs="unbounded"
}

type Entry struct {
	Reason      string `xml:"Reason,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type SearchLocation struct {
	Airport            *Airport            `xml:"Airport,omitempty"`            //minOccurs="0"
	City               *City               `xml:"City,omitempty"`               //minOccurs="0"
	CityOrAirport      *CityOrAirport      `xml:"CityOrAirport,omitempty"`      //minOccurs="0"
	CoordinateLocation *CoordinateLocation `xml:"CoordinateLocation,omitempty"` //minOccurs="0"
	RailLocation       *RailLocation       `xml:"RailLocation,omitempty"`       //minOccurs="0"
	Distance           *Distance           `xml:"Distance,omitempty"`           //minOccurs="0"
}

//Airport identifier
type Airport struct {
	Code string `xml:"Code,attr"` //use="required"
}

//City identifier
type City struct {
	Code string `xml:"Code,attr"` //use="required"
}

//This element can be used when it is not known whether the value is an airport or a city code.
type CityOrAirport struct {
	Code       string `xml:"Code,attr"`                 //use="required" //The airport or city IATA code.
	PreferCity bool   `xml:"PreferCity,attr,omitempty"` //use="optional" default="false" //Indicates that the search should prefer city results over airport results.
}

//Specific lat/long location, usually associated with a Distance
type CoordinateLocation struct {
	Latitude  float64 `xml:"latitude,attr"`  //use="required"
	Longitude float64 `xml:"longitude,attr"` //use="required"
}

//RCH specific location code (a.k.a UCodes) which uniquely identifies a train station.
type RailLocation struct {
	Code string `xml:"Code,attr"` //use="required"
}

type Location struct {
	Airport       *Airport       `xml:"Airport,omitempty"`       //minOccurs="0"
	City          *City          `xml:"City,omitempty"`          //minOccurs="0"
	CityOrAirport *CityOrAirport `xml:"CityOrAirport,omitempty"` //minOccurs="0"
}

//Container to encapsulate the a distance value with its unit of measure.
type Distance struct {
	Units     string `xml:"Units,attr,omitempty"` //use="optional" default="MI"
	Value     int    `xml:"Value,attr"`           //use="required"
	Direction string `xml:"Direction,attr"`
}

//Container to return/send additional retrieve/request additional search results
type NextResultReference struct {
	Value        string `xml:",innerxml"`
	ProviderCode string `xml:"ProviderCode,attr,omitempty"` //use="optional
}

//Point of Sale information for Billing
type BillingPointOfSaleInfo struct {
	OriginApplication string `xml:"OriginApplication,attr"`    //use="required"
	CIDBNumber        int    `xml:"CIDBNumber,attr,omitempty"` //use="optional"
}

//Vendor specific agent identifier overrides to be used to access vendor systems.
type AgentIDOverride struct {
	SupplierCode string `xml:"SupplierCode,attr"` //use="optional" //Supplier code to determine which vendor this AgentId belongs to.
	ProviderCode string `xml:"ProviderCode,attr"` //use="required" //Provider code to route the AgentId to proper provider.
	AgentID      int    `xml:"AgentID,attr"`      //use="required" //The Agent ID for the applicable supplier/vendor
}

//A type which can be used for flexible date/time specification -extends the generic type typeTimeSpec to provide extra options for search.
type FlexibleTimeSpec struct {
	TimeSpec
	SearchExtraDays *SearchExtraDays `xml:"SearchExtraDays,omitempty"` //minOccurs="0"
}

//Specifies times as either specific times, or a time range
type TimeSpec struct {
	TimeRange     *TimeRange    `xml:"TimeRange,omitempty"`          //minOccurs="0"
	SpecificTime  *SpecificTime `xml:"SpecificTime,omitempty"`       //minOccurs="0"
	PreferredTime string        `xml:"PreferredTime,attr,omitempty"` //use="optional" //Specifies a time that would be preferred within the time range specified.
}

//Specify a range of times.
type TimeRange struct {
	EarliestTime string `xml:"EarliestTime,attr"` //use="required"
	LatestTime   string `xml:"LatestTime,attr"`   //use="required"
}

//Specify exact times. System will automatically convert to a range according to agency configuration.
type SpecificTime struct {
	Time string `xml:"Time,attr"` //use="required"
}

//
type SearchExtraDays struct {
	DaysBefore int `xml:"DaysBefore,attr,omitempty"` //use="optional" //Number of days to search before the specified date
	DaysAfter  int `xml:"DaysAfter,attr,omitempty"`  //use="optional" //Number of days to search after the specified date
}

//Requests cabin class (First, Business and Economy, etc.) as supported by the provider or supplier.
type CabinClass struct {
	Type string `xml:"Type,attr"` //use="required"
}

//Carrier identifier
type Carrier struct {
	Code string `xml:"Code,attr"` //use="required"
}

//Provider identifier
type Provider struct {
	Code string `xml:"Code,attr"` //use="required"
}

//Used to emulate to another PCC or SID.  Providers: 1G, 1V, 1P, 1J.
type OverridePCC struct {
	ProviderCode   string `xml:"ProviderCode,attr"`   //use="required"
	PseudoCityCode string `xml:"PseudoCityCode,attr"` //use="required"
}

//Alliance Code
type Alliance struct {
	Code string `xml:"Code,attr"` //use="required" //The possible values are *A for Star Alliance,*O for One world,*S for Sky team etc.
}

//The base segment type
type Segment struct {
	SegmentRemark        []*SegmentRemark `xml:"SegmentRemark,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	Key                  com.TypeRef      `xml:"Key,attr"`                            //use="required"
	Status               string           `xml:"Status,attr,omitempty"`               //use="optional" //Status of this segment.
	Passive              bool             `xml:"Passive,attr,omitempty"`              //use="optional"
	TravelOrder          int              `xml:"TravelOrder,attr,omitempty"`          //use="optional" //To identify the appropriate travel sequence for Air/Car/Hotel segments/reservations based on travel dates. This ordering is applicable across the UR not provider or traveler specific
	ProviderSegmentOrder int              `xml:"ProviderSegmentOrder,attr,omitempty"` //use="optional" //To identify the appropriate travel sequence for Air/Car/Hotel/Rail segments/reservations in the provider reservation.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//A textual remark container to hold any printable text. (max 512 chars)
type SegmentRemark struct {
	Value string      `xml:",innerxml"`
	Key   com.TypeRef `xml:"Key,attr"` //use="required"
}

//Passenger type with code and optional age information
type SearchPassenger struct {
	PassengerType
	Key com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
}

//Passenger type code with optional age information
type PassengerType struct {
	Name                 *Name              `xml:"Name,omitempty"`                      //minOccurs="0" //Optional passenger Name with associated LoyaltyCard may provide benefit when pricing itineraries using Low Cost Carriers. In general, most carriers do not consider passenger LoyalyCard information when initially pricing itineraries.
	LoyaltyCard          []*LoyaltyCard     `xml:"LoyaltyCard,omitempty"`               //minOccurs="0" maxOccurs="unbounded"
	DiscountCard         []*DiscountCard    `xml:"DiscountCard,omitempty"`              //minOccurs="0" maxOccurs="9"
	PersonalGeography    *PersonalGeography `xml:"PersonalGeography,omitempty"`         //minOccurs="0" //Passenger personal geography detail to be sent to Host for accessing location specific fares
	Code                 com.TypePTC        `xml:"Code,attr"`                           //use="required" //The 3-char IATA passenger type code
	Age                  int                `xml:"Age,attr,omitempty"`                  //use="optional"
	DOB                  string             `xml:"DOB,attr,omitempty"`                  //use="optional" //Passenger Date of Birth
	Gender               com.TypeGender     `xml:"Gender,attr,omitempty"`               //use="optional" //The passenger gender type
	PricePTCOnly         bool               `xml:"PricePTCOnly,attr,omitempty"`         //use="optional"
	BookingTravelerRef   string             `xml:"BookingTravelerRef,attr,omitempty"`   //use="optional" //This value should be set for Multiple Passengers in the request.
	AccompaniedPassenger bool               `xml:"AccompaniedPassenger,attr,omitempty"` //use="optional" default="false" //Container to identify accompanied passenger. Set true means this passenger is accompanied
	ResidencyType        string             `xml:"ResidencyType,attr,omitempty"`        //use="optional" //The passenger residence type.
}

//Complete name fields
type Name struct {
	Prefix            string            `xml:"Prefix,attr,omitempty"`            //use="optional" //Name prefix. Size can be up to 20 characters
	First             string            `xml:"First,attr"`                       //use="required" //First Name. Size can be up to 256 characters
	Middle            string            `xml:"Middle,attr,omitempty"`            //use="optional" //Midle name. Size can be up to 256 characters
	Last              string            `xml:"Last,attr"`                        //use="required" //Last Name. Size can be up to 256 characters
	Suffix            string            `xml:"Suffix,attr,omitempty"`            //use="optional" //Name suffix. Size can be up to 256 characters
	TravelerProfileId com.TypeProfileID `xml:"TravelerProfileId,attr,omitempty"` //Traveler Applied Profile ID.
}

//Provider loyalty card information
type LoyaltyCard struct {
	ProviderReservationSpecificInfo []*ProviderReservationSpecificInfo `xml:"ProviderReservationSpecificInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	Key                             com.TypeRef                        `xml:"Key,attr,omitempty"`                        //use="optional"
	SupplierCode                    com.TypeCarrier                    `xml:"SupplierCode,attr,omitempty"`               //use="optional" //Carrier Code
	AllianceLevel                   string                             `xml:"AllianceLevel,attr,omitempty"`              //use="optional"
	MembershipProgram               com.StringLength1to32              `xml:"MembershipProgram,attr,omitempty"`          //use="optional" //Loyalty Program membership Id of the traveler specific to Amtrak(2V) Guest Rewards
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	CardNumber        com.TypeCardNumber   `xml:"CardNumber,attr"`                  //use="required"
	Status            string               `xml:"Status,attr,omitempty"`            //use="optional"
	MembershipStatus  string               `xml:"MembershipStatus,attr,omitempty"`  //use="optional"
	FreeText          string               `xml:"FreeText,attr,omitempty"`          //use="optional"
	SupplierType      string               `xml:"SupplierType,attr,omitempty"`      //use="optional"
	Level             string               `xml:"Level,attr,omitempty"`             //use="optional"
	PriorityCode      com.TypePriorityCode `xml:"PriorityCode,attr,omitempty"`      //use="optional"
	VendorLocationRef string               `xml:"VendorLocationRef,attr,omitempty"` //use="optional"
}

//
type ProviderReservationSpecificInfo struct {
	OperatedBy                 []*OperatedBy               `xml:"OperatedBy,omitempty"`                    //minOccurs="0" maxOccurs="unbounded" //Cross accrual carrier info
	ProviderReservationInfoRef *ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"`    //minOccurs="0" //Tagging provider reservation info with LoyaltyCard.
	ProviderReservationLevel   bool                        `xml:"ProviderReservationLevel,attr,omitempty"` //use="optional" //If true means Loyalty card is applied at ProviderReservation level.
	ReservationLevel           bool                        `xml:"ReservationLevel,attr,omitempty"`         //use="optional" //If true means Loyalty card is applied at Universal Record Reservation level e.g. Hotel Reservation, Vehicle Reservation etc.
}

type OperatedBy struct {
	Value string `xml:",innerxml"` //This is the carrier code to support Cross Accrual
}

//[RS] Container for Provider reservation reference key.
type ProviderReservationInfoRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Rail Discount Card Information
type DiscountCard struct {
	Key         com.TypeRef            `xml:"Key,attr,omitempty"`         //use="optional"
	Code        com.StringLength1to8   `xml:"StringLength1to8,attr"`      //use="required"
	Description com.StringLength1to255 `xml:"Description,attr,omitempty"` //use="optional"
	Number      com.TypeCardNumber     `xml:"Number,attr,omitempty"`      //use="optional"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Personal geography details of the associated passenger.
type PersonalGeography struct {
	CountryCode       com.TypeCountry `xml:"CountryCode,omitempty"`       //minOccurs="0" //Passenger country code.
	StateProvinceCode com.TypeState   `xml:"StateProvinceCode,omitempty"` //minOccurs="0" //Passenger state province code.
	CityCode          com.TypeCity    `xml:"CityCode,omitempty"`          //minOccurs="0" //Passenger state Passenger city code.
}

//User can use this node to send a specific PCC to access fares allowed only for that PCC. This node gives the capability for fare redistribution at UR level.  For fare redistribution at the stored fare level see AirPricingSolution/AirPricingInfo/AirPricingModifiers/PointOfSale.
type PointOfSale struct {
	ProviderCode   com.TypeProviderCode `xml:"ProviderCode,attr"`   //use="required" //The provider in which the PCC is defined.
	PseudoCityCode com.TypePCC          `xml:"PseudoCityCode,attr"` //use="required" //The PCC in the host system.
	Key            com.TypeRef          `xml:"Key,attr,omitempty"`  //use="optional"
	IATA           com.TypeIATA         `xml:"IATA,attr,omitempty"` //use="optional" //Used for rapid reprice. This field is the IATA associated to this Point of Sale PCC. Providers: 1G/1V
}

type TaxInfo struct {
	TaxDetail              []*TaxDetail           `xml:"TaxDetail,omitempty"`                   //minOccurs="0" maxOccurs="unbounded"
	IncludedInBase         *IncludedInBase        `xml:"IncludedInBase,omitempty"`              //minOccurs="0" maxOccurs="1"
	Key                    com.TypeRef            `xml:"Key,attr,omitempty"`                    //use="optional" //The tax key represents a valid key of tax
	Category               string                 `xml:"Category,attr"`                         //use="required" //The tax category represents a valid IATA tax code.
	CarrierDefinedCategory string                 `xml:"CarrierDefinedCategory,attr,omitemtpy"` //use="optional" //Optional category, where a carrier has used a non-standard IATA tax category. The tax category will be set to "DU"
	SegmentRef             com.TypeRef            `xml:"SegmentRef,attr,omitempty"`             //use="optional" //The segment to which that tax is relative (if applicable)
	FlightDetailsRef       com.TypeRef            `xml:"FlightDetailsRef,attr,omitempty"`       //use="optional" //The flight details that this tax is relative to (if applicable)
	CouponRef              com.TypeRef            `xml:"CouponRef,attr,omitempty"`              //use="optional" //The coupon to which that tax is relative (if applicable)
	Amount                 com.TypeMoney          `xml:"Amount,attr"`                           //use="required"
	OriginAirport          com.TypeAirport        `xml:"OriginAirport,attr,omitempty"`          //use="optional"
	DestinationAirport     com.TypeAirport        `xml:"DestinationAirport,attr,omitempty"`     //use="optional"
	CountryCode            string                 `xml:"CountryCode,attr,omitemtpy"`            //use="optional"
	FareInfoRef            com.TypeRef            `xml:"FareInfoRef,attr,omitemtpy"`            //use="optional"
	TaxExempted            bool                   `xml:"TaxExempted,attr,omitempty"`            //use="optional" //This indicates whether the tax specified by tax category is exempted.
	ProviderCode           com.TypeProviderCode   `xml:"ProviderCode,attr,omitempty"`           //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode           com.TypeSupplierCode   `xml:"SupplierCode,attr,omitempty"`           //use="optional" //Code of the supplier returning this TaxInfo.
	Text                   com.StringLength1to128 `xml:"Text,attr,omitempty"`                   //use="optional" //Additional Information returned from Supplier.(ACH  only)
}

//The tax detail information for a fare quote tax.
type TaxDetail struct {
	Amount             com.TypeMoney   `xml:"Amount,attr"`                       //use="required"
	OriginAirport      com.TypeAirport `xml:"OriginAirport,attr,omitempty"`      //use="optional"
	DestinationAirport com.TypeAirport `xml:"DestinationAirport,attr,omitempty"` //use="optional"
	CountryCode        string          `xml:"CountryCode,attr,omitemtpy"`        //use="optional"
	FareInfoRef        com.TypeRef     `xml:"FareInfoRef,attr,omitemtpy"`        //use="optional"
}

//A generic type of fee for those charges which are incurred by the passenger, but not necessarily shown on tickets
type FeeInfo struct {
	TaxInfoRef         []*TaxInfoRef         `xml:"TaxInfoRef,omitempty"`              //minOccurs="0" maxOccurs="unbounded" //This reference elements will associate relevant taxes to this fee<
	Key                com.TypeRef           `xml:"Key,attr"`                          //use="required"
	IncludedInBase     *IncludedInBase       `xml:"IncludedInBase,omitempty"`          //minOccurs="0" maxOccurs="1"
	BaseAmount         com.TypeMoney         `xml:"BaseAmount,attr,omitempty"`         //use="optional"
	Description        string                `xml:"Description,attr,omitempty"`        //use="optional"
	SubCode            string                `xml:"SubCode,attr,omitempty"`            //use="optional"
	Amount             com.TypeMoney         `xml:"Amount,attr"`                       //use="required"
	Code               string                `xml:"Code,attr"`                         //use="required"
	FeeToken           string                `xml:"FeeToken,attr,omitempty"`           //use="optional"
	PaymentRef         com.TypeRef           `xml:"PaymentRef,attr,omitempty"`         //use="optional" //The reference to the one of the air reservation payments if fee included in charge
	BookingTravelerRef com.TypeRef           `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //Reference to booking traveler
	PassengerTypeCode  com.TypePTC           `xml:"PassengerTypeCode,attr,omitempty"`
	Text               com.StringLength1to64 `xml:"Text,attr,omitempty"` //use="optional" //Additional Information returned from Supplier.(ACH  only)
	//<xs:attributeGroup ref="attrProviderSupplier"> //attrProviderSupplier
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional" //Code of the supplier returning this TaxInfo.
	//</xs:attributeGroup>
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//This reference elements will associate relevant taxes to this fee
type TaxInfoRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Extra data to elaborate the parent element. This data is primarily informative and is not persisted.
type MetaData struct {
	Key   int `xml:"Key,attr"`   //use="required"
	Value int `xml:"Value,attr"` //use="required"
}

//A simple textual fare note. Used within several other objects.
type ResponseMessage struct {
	Value string `xml:",innerxml"`
	Code  int    `xml:"Code,attr"`           //use="required"
	Type  string `xml:"Type,attr,omitempty"` //use="optional" //Indicates the type of message (Warning, Error, Info)
	//<xs:attributeGroup ref="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional" //Code of the supplier returning this TaxInfo.
	//</xs:attributeGroup>
}

//[RS] Used to identify the results of a requests
type ResultMessage struct {
	Value string `xml:",innerxml"`
	Code  int    `xml:"Code,attr"`           //use="required"
	Type  string `xml:"Type,attr,omitempty"` //use="optional" //Indicates the type of message (Warning, Error, Info)
}

//This is a host token. It contains some kind of payload we got from a host
//that must be passed in on successive calls they know who you are as our system
//does not maintain state. The format of this string isn't important as long as it
//is not altered in any way between calls. Since a host token is only valid on
//the host it is assocated with, there is also an attribute called Host so we know
//how to route the command(s). You can have multiple active sessions between
//one or more hosts
type HostToken struct {
	Value string               `xml:",innerxml"`
	Host  com.TypeProviderCode `xml:"Host,attr"` //The host associated with this token
	Key   string               `xml:"Key,attr"`  //Unique identifier for this token - use this key when a single HostToken is shared by multiple elements.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Reference Element for Booking Traveler and Loyalty cards
type BookingTravelerRef struct {
	LoyaltyCardRef    []*LoyaltyCardRef  `xml:"LoyaltyCardRef,omitempty"`    //minOccurs="0" maxOccurs="unbounded"
	DriversLicenseRef *DriversLicenseRef `xml:"DriversLicenseRef,omitempty"` //minOccurs="0"
	DiscountCardRef   []*DiscountCardRef `xml:"DiscountCardRef,omitempty"`   //minOccurs="0" maxOccurs="9"
	PaymentRef        []*PaymentRef      `xml:"PaymentRef,omitempty"`        //minOccurs="0" maxOccurs="3"
	Key               com.TypeRef        `xml:"Key,attr"`                    //use="required"
}

//Reference Element for Loyalty cards
type LoyaltyCardRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Reference Element for Drivers Licenses
type DriversLicenseRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Reference Element for Discount Card
type DiscountCardRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Reference Element for Payment
type PaymentRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Payment information - must be used in conjunction with credit card info
type Payment struct {
	Key                com.TypeRef           `xml:"Key,attr,omitempty"`
	Type               string                `xml:"Type,attr"`                         //Identifies the type of payment. This can be for an itinerary, a traveler, or a service fee for example.
	FormOfPaymentRef   com.TypeRef           `xml:"FormOfPaymentRef,attr"`             //use="required" //The credit card that is will be used to make this payment.
	BookingTravelerRef com.TypeRef           `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //If the type represents a per traveler payment, then this will reference the traveler this payment refers to.
	Amount             com.TypeMoney         `xml:"Amount,attr"`                       //use="required"
	AmountType         com.StringLength1to32 `xml:"AmountType,attr,omitempty"`         //use="optional" //This field displays type of payment amount when it is non-monetary. Presently available/supported value is "Flight Pass Credits".
	ApproximateAmount  com.TypeMoney         `xml:"ApproximateAmount,attr,omitempty"`  //use="optional" //It stores the converted payment amount in agency's default currency
	Status             string                `xml:"Status,attr,omitempty"`             //use="optional" //Status to indicate the business association of the payment element.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Account Code is used to get Private Fares.If ProviderCode or SupplierCode is not specified,
//it will be considered a default AccounCode to be sent to all the Providers or Suppliers.
type AccountCode struct {
	Code string `xml:"Code,attr,omitempty"` //use="optional"
	//<xs:attributeGroup ref="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional" //Code of the supplier returning this TaxInfo.
	//</xs:attributeGroup>
	Type string `xml:"Type,attr,omitempty"` //use="optional" //An identifier to categorize this account code. For example, FlightPass for AC Flight Pass or RFB for AC corporate Rewards for Business.
}

type LoyaltyProgram struct {
	//<xs:attributeGroup name="attrLoyalty">
	Key               com.TypeRef           `xml:"Key,attr,omitempty"`               //use="optional"
	SupplierCode      com.TypeCarrier       `xml:"SupplierCode,attr"`                //use="required" //The code used to identify the Loyalty supplier, e.g. AA, ZE, MC
	AllianceLevel     string                `xml:"AllianceLevel,attr,omitempty"`     //use="optional"
	MembershipProgram com.StringLength1to32 `xml:"MembershipProgram,attr,omitempty"` //use="optional" //Loyalty Program membership Id of the traveler specific to Amtrak(2V) Guest Rewards
	//</xs:attributeGroup>
	Level string `xml:"Level,attr,omitempty"` //use="optional
}

//[RS] Special serivces like wheel chair, or pet carrier.
type SSR struct {
	Key                        com.TypeRef         `xml:"Key,attr,omitempty"`                        //use="optional"
	SegmentRef                 com.TypeRef         `xml:"SegmentRef,attr,omitempty"`                 //use="optional" //Reference to the air segment. May be required for some Types.
	PassiveSegmentRef          com.TypeRef         `xml:"PassiveSegmentRef,attr,omitempty"`          //use="optional" //Reference to the passive segment.
	ProviderReservationInfoRef com.TypeRef         `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	Type                       com.TypeSSRCode     `xml:"Type,attr"`                                 //use="required" //Programmatic SSRs use codes recognized by the provider/supplier (example, VGML=vegetarian meal code). Manual SSRs do not have an associated programmatic code.
	Status                     string              `xml:"Status,attr,omitempty"`                     //use="optional"
	FreeText                   com.TypeSSRFreeText `xml:"FreeText,attr,omitempty"`                   //use="optional" //Certain SSR types will require a free text message. For example MAAS (Meet and assist).
	Carrier                    com.TypeCarrier     `xml:"Carrier,attr,omitempty"`                    //use="optional"
	CarrierSpecificText        string              `xml:"CarrierSpecificText,attr,omitempty"`        //Carrier specific information which are not captured in the FreeText field(not present in IATA's standard SSR DOCO format). An example is VISA Expiration Date.
	Description                string              `xml:"Description,attr,omitempty"`                //use="optional"
	ProviderDefinedType        string              `xml:"ProviderDefinedType,attr,omitempty"`        //use="optional" //Original Type as sent by the provider
	SSRRuleRef                 com.TypeRef         `xml:"SSRRuleRef,attr,omitempty"`                 //use="optional" //UniqueID to associate a rule to the SSR
	URL                        string              `xml:"URL,attr"`                                  //xs:anyURI
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	ProfileID                 string      `xml:"ProfileID,attr,omitempty"`                 //use="optional" //Key assigned for Secure Flight Document value from the specified profile
	ProfileSecureFlightDocKey com.TypeRef `xml:"ProfileSecureFlightDocKey,attr,omitempty"` //use="optional" //Unique ID of Booking Traveler's Profile that contains the Secure flight Detail
}

//Contains the rules for applying service rules
type ServiceRuleType struct {
	ApplicationRules   *ApplicationRules        `xml:"ApplicationRules,omitempty"`   //minOccurs="0"
	ApplicationLevel   *ApplicationLevel        `xml:"ApplicationLevel,omitempty"`   //minOccurs="0"
	ModifyRules        *ModifyRules             `xml:"ModifyRules,omitempty"`        //minOccurs="0"
	SecondaryTypeRules *SecondaryTypeRules      `xml:"SecondaryTypeRules,omitempty"` //minOccurs="0"
	Remarks            []*FormattedTextTextType `xml:"Remarks,omitempty"`            //minOccurs="0" maxOccurs="99" //Adds text remarks / rules for the optional / additional service
	Key                com.TypeRef              `xml:"Key,attr"`                     //use="required" //Unique ID to identify an optional service rule
}

//The rules to apply the rule to the itinerary
type ApplicationRules struct {
	RequiredForAllTravelers     bool `xml:"RequiredForAllTravelers,attr"`     //Indicates if the option needs to be applied to all travelers in the itinerary if selected
	RequiredForAllSegments      bool `xml:"RequiredForAllSegments,attr"`      //Indicates if the option needs to be applied to all segments in the itinerary if selected
	RequiredForAllSegmentsInOD  bool `xml:"RequiredForAllSegmentsInOD,attr"`  //Indicates if the option needs to be applied to all segments in a origin / destination (connection flights) if selected for one segment in the OD
	UnselectedOptionRequired    bool `xml:"UnselectedOptionRequired,attr"`    //If an UnselectedOption is present in the option, then the Unselected option  needs to be selected even if the option is not selected when this flag is set to true
	SecondaryOptionCodeRequired bool `xml:"SecondaryOptionCodeRequired,attr"` //If set to true, the secondary option code is required for this option
}

//Lists the levels where the option is applied in the itinerary. Some options are applied for the entire itinerary, some for entire segments, etc.
type ApplicationLevel struct {
	ApplicationLimits               *ApplicationLimits `xml:"ApplicationLimits,omitempty"`                    //minOccurs="0"
	ServiceData                     []*ServiceData     `xml:"ServiceData,omitempty"`                          //minOccurs="0" maxOccurs="unbounded"
	ApplicableLevels                string             `xml:"ApplicableLevels,attr"`                          //<xs:list itemType="OptionalServiceApplicabilityType"/> //Indicates the level in the itinerary when the option is applied.
	ProviderDefinedApplicableLevels string             `xml:"ProviderDefinedApplicableLevels,attr,omitempty"` // use="optional" //Indicates the actual provider defined ApplicableLevels which is mapped to Other
}

//Adds the limits on the number of options that can be selected for a particular type
type ApplicationLimits struct {
	ApplicationLimit []*OptionalServiceApplicationLimitType `xml:"ApplicationLimit"` //maxOccurs="10" //The application limits for a particular level
}

//Groups the modification rules for the Option
type ModifyRules struct {
	ModifyRule                      []*ModifyRule `xml:"ModifyRule"`                                     //maxOccurs="unbounded"
	SupportedModifications          string        `xml:"SupportedModifications,attr"`                    //<xs:list itemType="ModificationType"/> //Lists the supported modifications for the itinerary.
	ProviderDefinedModificationType string        `xml:"ProviderDefinedModificationType,attr,omitempty"` //use="optional"
}

//Indicates modification rules for the particular modification type.
type ModifyRule struct {
	//<xs:attributeGroup name="ModificationRulesGroup">
	Modification                    string `xml:"Modification,attr"`                              //use="required" //The modificaiton for which this rule group applies.
	AutomaticallyAppliedOnAdd       bool   `xml:"AutomaticallyAppliedOnAdd,attr,omitempty"`       //use="optional" default="false" //Indicates if the option will be automatically added to new segments / passengers in the itinerary.
	CanDelete                       bool   `xml:"CanDelete,attr,omitempty"`                       //use="optional" //Indicates if the option can be deleted from the itinerary without segment or passenger modifications
	CanAdd                          bool   `xml:"CanAdd,attr,omitempty"`                          //use="optional" //Indicates if the option can be added to the itinerary without segment or passenger modification
	Refundable                      bool   `xml:"Refundable,attr,omitempty"`                      //use="optional" //Indicates if the price of the option is refundable.
	ProviderDefinedModificationType string `xml:"ProviderDefinedModificationType,attr,omitempty"` //use="optional" //Indicates the actual provider defined modification type which is mapped to Other
	//</xs:attributeGroup>
}

//Lists the supported Secondary Codes for the optional / additional service.
type SecondaryTypeRules struct {
	SecondaryTypeRule []*SecondaryTypeRule `xml:"ModifyRule"` //maxOccurs="unbounded"
}

//Lists a single secondary code for the optional / additional service.
type SecondaryTypeRule struct {
	ApplicationLimit []*OptionalServiceApplicationLimitType `xml:"ApplicationLimit,omitempty"` //minOccurs="0" maxOccurs="10"
	SecondaryType    com.TypeRef                            `xml:"SecondaryType,attr"`         //use="required" //The unique type to associate a secondary type in an optional service
}

//The optional service application limit
type OptionalServiceApplicationLimitType struct {
	//<xs:attributeGroup name="OptionalServiceApplicabilityLimitGroup">
	ApplicableLevel                 string `xml:"ApplicableLevel,attr"`                           //use="required" //Indicates the applicable level for the option
	ProviderDefinedApplicableLevels string `xml:"ProviderDefinedApplicableLevels,attr,omitempty"` //use="optional" //Indicates the actual provider defined ApplicableLevels which is mapped to Other
	MaximumQuantity                 uint   `xml:"MaximumQuantity,attr"`                           //use="required" //The maximum quantity allowed for the type
	MinimumQuantity                 uint   `xml:"MinimumQuantity,attr,omitempty"`                 //use="optional" //Indicates the minimum number of the option that can be selected.
	//</xs:attributeGroup>
}

//
type ServiceData struct {
	SeatAttributes     *SeatAttributes      `xml:"SeatAttributes,omitempty"`          //minOccurs="0"
	CabinClass         *CabinClass          `xml:"CabinClass,omitempty"`              //minOccurs="0"
	SSRRef             []*KeyBasedReference `xml:"SSRRef,omitempty"`                  //minOccurs="0" maxOccurs="unbounded" //References to the related SSRs. At present, only reference to ASVC SSR is supported. Supported providers are 1G/1V/1P/1J
	Data               string               `xml:"Data,attr,omitempty"`               //use="optional" //Data that specifies the details of the merchandising offering (e.g. seat number for seat service)
	AirSegmentRef      com.TypeRef          `xml:"AirSegmentRef,attr,omitempty"`      //use="optional" //Reference to a segment if the merchandising offering only pertains to that segment. If no segment reference is present this means this offering is for the whole itinerary.
	BookingTravelerRef com.TypeRef          `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //Reference to a passenger if the merchandising offering only pertains to that passenger. If no passenger reference is present this means this offering is for all passengers.
	StopOver           bool                 `xml:"StopOver,attr,omitempty"`           //use="optional" default="false" //Indicates that there is a significant delay between flights (usually 12 hours or more)
	TravelerType       com.TypePTC          `xml:"TravelerType,attr,omitempty"`       //use="optional" //Passenger Type Code.
	EMDSummaryRef      com.TypeRef          `xml:"EMDSummaryRef,attr,omitempty"`      //use="optional" //Reference to the corresponding EMD issued. Supported providers are 1G/1V/1P/1J
	EMDCouponRef       com.TypeRef          `xml:"EMDCouponRef,attr,omitempty"`       //use="optional" //Reference to the corresponding EMD coupon issued. Supported providers are 1G/1V/1P/1J
}

//Identifies the seat attribute of the service.
type SeatAttributes struct {
	SeatAttribute []*SeatAttribute `xml:"SeatAttribute,omitempty"` //minOccurs="0" maxOccurs="10"
}

//Identifies the seat attribute of the service.
type SeatAttribute struct {
	Value string `xml:"Value,attr"` // use="required"
}

//Generic type to be used for Key based reference
type KeyBasedReference struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Provides text and indicates whether it is formatted or not.
type FormattedTextTextType struct {
	Value     string `xml:",innerxml"`
	Formatted bool   `xml:"Formatted,attr,omitempty"` //use="optional" //Textual information, which may be formatted as a line of information, or unformatted, as a paragraph of text.
	//<xs:attributeGroup name="LanguageGroup">
	Language string `xml:"Language,attr,omitempty"` //use="optional" type="xs:language" //Language identification.
	//</xs:attributeGroup>
	TextFormat string `xml:"TextFormat,attr,omitempty"` //use="optional" //Indicates the format of text used in the description e.g. unformatted  or html.
}

//Indicates Carrier Supports this industry standard.
type IndustryStandardSSR struct {
	Code string `xml:"Code,attr,omitempty"` //This code indicates which Standard of SSR's they support. Sucha as the 'AIRIMP' standard identified by 'IATA.org'
}

//Restrictions or instructions about the fare or ticket
type Endorsement struct {
	Value com.TypeEndorsement `xml:"Value,attr"` //use="required"
}

//Identifies the agency commission
type Commission struct {
	Key                com.TypeRef                   `xml:"Key,attr,omitempty"`                //use="optional"
	Level              string                        `xml:"Level,attr"`                        //use="required" //The commission percentage level.
	Type               string                        `xml:"Type,attr"`                         //use="required" //The commission type.
	Modifier           string                        `xml:"Modifier,attr,omitempty"`           //use="optional" //Optional commission modifier.
	Amount             com.TypeMoney                 `xml:"Amount,attr,omitempty"`             //use="optional" //The monetary amount of the commission.
	Value              string                        `xml:"Value,attr,omitempty"`              //use="optional" //Contains alphanumeric or alpha characters intended as 1G Value Code as applicable by BSP of client.
	Percentage         com.TypePercentageWithDecimal `xml:"Percentage,attr,omitempty"`         //use="optional" //The percent of the commission.
	BookingTravelerRef com.TypeRef                   `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //A reference to a passenger.
	CommissionOverride bool                          `xml:"CommissionOverride,attr,omitempty"` //use="optional" //default="false" //This is enabled to override CAT-35 commission error during air ticketing. PROVIDER SUPPORTED:Worldspan and JAL
}

type ServiceInfo struct {
	Description []string     `xml:"Description"`              //maxOccurs="unbounded" //Description of the Service.  Usually used in tandem with  one or more media items.
	MediaItem   []*MediaItem `xml:"MediaItem,attr,omitempty"` //minOccurs="0" maxOccurs="3"
}

//A textual remark container to hold any printable text. (max 512 chars)
type Remark struct {
	Value string      `xml:",innerxml"`
	Key   com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
}

//Photos and other media urls for the property referenced above.
type MediaItem struct {
	Caption  string `xml:"caption,attr"`
	Height   uint   `xml:"height,attr"`
	Width    uint   `xml:"width,attr"`
	Type     string `xml:"type,attr"`
	Url      string `xml:"url,attr"`  //xs:anyURI
	Icon     string `xml:"icon,attr"` //xs:anyURI
	SizeCode string `xml:"sizeCode,attr"`
}

//[RS]
type LinkedUniversalRecord struct {
	LocatorCode com.TypeLocatorCode `xml:"LocatorCode,attr"`   //use="required" //A Universal Record that need to be linked to the current Universal Record.
	Key         com.TypeRef         `xml:"Key,attr,omitempty"` //use="optional"
	//<xs:attributeGroup attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"` //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr"`      //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

/*<xs:simpleType name="typeResponseImageSize">
<xs:annotation>
<xs:documentation>Allowable images sizes in response
</xs:documentation>
</xs:annotation>
<xs:union memberTypes="typeImageSize typeOtherImageSize"/>
</xs:simpleType>*/

//[RS] A traveler and all their accompanying data.
type BookingTraveler struct {
	//<xs:group name="BaseBookingTravelerInfoA">
	BookingTravelerName *BookingTravelerName `xml:"BookingTravelerName"`
	DeliveryInfo        []*DeliveryInfo      `xml:"DeliveryInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	PhoneNumber         []*PhoneNumber       `xml:"PhoneNumber,omitempty"`  //minOccurs="0" maxOccurs="unbounded"
	Email               []*Email             `xml:"Email,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	LoyaltyCard         []*LoyaltyCard       `xml:"LoyaltyCard,omitempty"`  //minOccurs="0" maxOccurs="unbounded"
	DiscountCard        []*DiscountCard      `xml:"DiscountCard,omitempty"` //minOccurs="0" maxOccurs="9"
	//</xs:group>
	SSR               []*SSR               `xml:"SSR,omitempty"`               //minOccurs="0" maxOccurs="unbounded"
	NameRemark        []*NameRemark        `xml:"NameRemark,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	AirSeatAssignment []*AirSeatAssignment `xml:"AirSeatAssignment,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	//RailSeatAssignment []*RailSeatAssignment `xml:"RailSeatAssignment,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	//<xs:group name="BaseBookingTravelerInfoB">
	EmergencyInfo        string                   `xml:"EmergencyInfo,omitempty"`        //minOccurs="0"
	Address              []*TypeStructuredAddress `xml:"Address,omitempty"`              //minOccurs="0" maxOccurs="unbounded"
	DriversLicense       []*DriversLicense        `xml:"DriversLicense,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	AppliedProfile       []*AppliedProfile        `xml:"AppliedProfile,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	CustomizedNameData   []*CustomizedNameData    `xml:"CustomizedNameData,omitempty"`   //minOccurs="0" maxOccurs="unbounded"
	TravelComplianceData []*TravelComplianceData  `xml:"TravelComplianceData,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Travel Compliance and Preferred Supplier information of the booking traveler specific to a segment. Not applicable to Saved Trip.
	TravelInfo           *TravelInfo              `xml:"TravelInfo,omitempty"`           //minOccurs="0"
	//</xs:group>
	//<xs:attributeGroup name="attrBookingTravelerGrp">
	Key          com.TypeRef     `xml:"Key,attr,omitempty"`          //use="optional"
	TravelerType com.TypePTC     `xml:"TravelerType,attr,omitempty"` //use="optional" //Defines the type of traveler used for booking which could be a non-defining type (Companion, Web-fare, etc), or a standard type (Adult, Child, etc).
	Age          int             `xml:"Age,attr,omitempty"`          //use="optional" //BookingTraveler age
	VIP          bool            `xml:"VIP,attr,omitempty"`          //use="optional" //default="false" //When set to True indicates that the Booking Traveler is a VIP based on agency/customer criteria
	DOB          string          `xml:"DOB,attr,omitempty"`          //use="optional" //type="xs:date" //Traveler Date of Birth
	Gender       com.TypeGender  `xml:"Gender,attr,omitempty"`       //use="optional" //The BookingTraveler gender type
	Nationality  com.TypeCountry `xml:"Nationality,attr,omitempty"`  //use="optional" //Specify ISO country code for nationality of the Booking Traveler
	//</xs:attributeGroup>
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Complete name fields
type BookingTravelerName struct {
	//<xs:attributeGroup name="attrBookingTravelerName">
	Prefix string                   `xml:"Prefix,attr,omitempty"` //use="optional" //Name prefix.
	First  string                   `xml:"First,attr"`            //use="required" //First Name.
	Middle string                   `xml:"Middle,attr,omitempty"` //use="optional" //Midle name.
	Last   com.TypeTravelerLastName `xml:"Last,attr"`             //use="required" //Last Name.
	Suffix string                   `xml:"Suffix,attr,omitempty"` //use="optional" //Name suffix.
	//</xs:attributeGroup>
}

//[RS] Container to encapsulate all delivery related information
type DeliveryInfo struct {
	ShippingAddress            *ShippingAddress              `xml:"ShippingAddress,omitempty"`            //minOccurs="0"
	PhoneNumber                *PhoneNumber                  `xml:"PhoneNumber,omitempty"`                //minOccurs="0"
	Email                      *Email                        `xml:"Email,omitempty"`                      //minOccurs="0"
	GeneralRemark              []*GeneralRemark              `xml:"GeneralRemark,omitempty"`              //minOccurs="0" maxOccurs="unbounded"
	ProviderReservationInfoRef []*ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Tagging provider reservation info with Delivery Info.
	Type                       string                        `xml:"Type,attr,omitempty"`                  //use="optional" //An arbitrary identifier to categorize this delivery info
	SignatureRequired          string                        `xml:"SignatureRequired,attr,omitempty"`     //use="optional" //Indicates whether a signature shoud be required in order to make the delivery.
	TrackingNumber             string                        `xml:"TrackingNumber,attr,omitempty"`        //use="optional" //The tracking number of the shipping company making the delivery.
}

//[RS]
type ShippingAddress struct {
	TypeStructuredAddress
}

//[RS] A fully structured address
type TypeStructuredAddress struct {
	AddressName                string                        `xml:"AddressName,omitempty"`                //minOccurs="0"
	Street                     []string                      `xml:"Street,omitempty"`                     //minOccurs="0" maxOccurs="5" //The Address street and number, e.g. 105 Main St.
	City                       string                        `xml:"City,omitempty"`                       //minOccurs="0" //The city name for the requested address, e.g. Atlanta.
	State                      *State                        `xml:"State,omitempty"`                      //minOccurs="0" //The State or Province of address requested, e.g. CA, Ontario.
	PostalCode                 string                        `xml:"PostalCode,omitempty"`                 //minOccurs="0" //The 5-15 alphanumeric postal Code for the requested address, e.g. 90210.
	Country                    string                        `xml:"Country,omitempty"`                    //minOccurs="0" //The Full country name or two letter ISO country code e.g. US, France. A two letter country code is required for a Postal Code Searches.
	ProviderReservationInfoRef []*ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"` //minOccurs="0" maxOccurs="99" //Tagging provider reservation info with Address.
	Key                        com.TypeRef                   `xml:"Key,attr,omitempty"`                   //use="optional" //Key for update/delete of the element
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Consists of type (office, home, fax), location (city code), the country code, the number, and an extension.
type PhoneNumber struct {
	ProviderReservationInfoRef []*ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Tagging provider reservation info with PhoneNumber.
	Key                        com.TypeRef                   `xml:"Key,attr,omitempty"`                   //use="optional"
	Type                       string                        `xml:"Type,attr,omitempty"`                  //use="optional"
	Location                   string                        `xml:"Location,attr,omitempty"`              //use="optional" //IATA code for airport or city
	CountryCode                string                        `xml:"CountryCode,attr,omitempty"`           //use="optional" //Hosts/providers will expect this to be international dialing digits
	AreaCode                   string                        `xml:"AreaCode,attr,omitempty"`              //use="optional"
	Number                     string                        `xml:"Number,attr"`                          //use="required" //The local phone number
	Extension                  string                        `xml:"Extension,attr,omitempty"`             //use="optional"
	Text                       string                        `xml:"Text,attr,omitempty"`                  //use="optional"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Container for an email address with a type specifier (max 128 chars)
type Email struct {
	ProviderReservationInfoRef []*ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Tagging provider reservation info with Email.
	Type                       com.TypeRef                   `xml:"Key,attr,omitempty"`                   //use="optional"
	Key                        com.TypeEmailType             `xml:"Type,attr,omitempty"`                  //use="optional"
	Comment                    com.TypeEmailComment          `xml:"Comment,attr,omitempty"`               //use="optional"
	EmailID                    string                        `xml:"EmailID,attr"`                         //use="required"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Container to house the state code for an address
type State struct {
	Value string `xml:",innerxml"`
}

//[RS]
type SeatAssignment struct {
	Key              com.TypeRef          `xml:"Key,attr,omitempty"`              //use="optional"
	Status           com.TypeStatusCode   `xml:"Status,attr"`                     //use="required"
	Seat             string               `xml:"Seat,attr"`                       //use="required"
	SeatTypeCode     com.TypeSeatTypeCode `xml:"SeatTypeCode,attr,omitempty"`     //use="optional" //The 4 letter SSR code like SMSW,NSSW,SMST etc.
	SegmentRef       com.TypeRef          `xml:"SegmentRef,attr,omitempty"`       //use="optional"
	FlightDetailsRef com.TypeRef          `xml:"FlightDetailsRef,attr,omitempty"` //use="optional"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	RailCoachNumber string `xml:"RailCoachNumber,attr,omitempty"` //use="optional" //Coach number for which rail seatmap/coachmap is returned.
}

//[RS] Identifies the seat assignment for a passenger.
type AirSeatAssignment struct {
	SeatAssignment
}

//[RS] A remark container to hold an XML document. (max 1024 chars) This will be encoded with xml encoding.
type XMLRemark struct {
	Value    string      `xml:",innerxml"`
	Key      com.TypeRef `xml:"Key,attr,omitempty"`      //use="optional"
	Category string      `xml:"Category,attr,omitempty"` //use="optional" //A category to group and organize the various remarks. This is not required, but it is recommended.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] A textual remark container to hold any printable text. (max 512 chars)
type GeneralRemark struct {
	RemarkData                 string            `xml:"RemarkData"`                                //Actual remarks data.
	BookingTravelerRef         []com.TypeRef     `xml:"BookingTravelerRef"`                        //minOccurs="0" maxOccurs="unbounded" //Reference to Booking Traveler.
	Type                       com.TypeRef       `xml:"Key,attr,omitempty"`                        //use="optional"
	Category                   string            `xml:"Category,attr,omitempty"`                   //use="optional" //A category to group and organize the various remarks. This is not required, but it is recommended.
	TypeInGds                  com.TypeGdsRemark `xml:"TypeInGds,attr,omitempty"`                  //use="optional"
	SupplierType               string            `xml:"SupplierType,attr,omitempty"`               //use="optional" //The type of product this reservation is relative to
	ProviderReservationInfoRef com.TypeRef       `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	//<xs:attributeGroup name="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional"
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional"
	//</xs:attributeGroup>
	Direction             string `xml:"Direction,attr,omitempty"`             //use="optional" //Direction Incoming or Outgoing of the GeneralRemark.
	CreateDate            string `xml:"CreateDate,attr,omitempty"`            //use="optional" //type="xs:dateTime" //The date and time that this GeneralRemark was created.
	UseProviderNativeMode bool   `xml:"UseProviderNativeMode,attr,omitempty"` //use="optional" default="false" //Will be true when terminal process required, else false
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS]
type RemarkWithTravelerRef struct {
	RemarkData                 string               `xml:"RemarkData"`                                //Actual remarks data.
	BookingTravelerRef         []com.TypeRef        `xml:"BookingTravelerRef,omitempty"`              //minOccurs="0" maxOccurs="unbounded" //Reference to Booking Traveler.
	ProviderReservationInfoRef com.TypeRef          `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this element is used
}

//[RS] A textual remark container to hold non-associated itinerary remarks
type UnassociatedRemark struct {
	RemarkWithTravelerRef
	Key com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Text that support Name Remarks.
type NameRemark struct {
	RemarkData                 string                        `xml:"RemarkData"`                           //Actual remarks data.
	ProviderReservationInfoRef []*ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Tagging provider reservation info with NameRemark.
	Key                        com.TypeRef                   `xml:"Key,attr,omitempty"`                   //use="optional"
	Category                   string                        `xml:"Category,attr,omitempty"`              //use="optional" //A category to group and organize the various remarks. This is not required, but it is recommended.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Details of drivers license
type DriversLicense struct {
	Key           com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
	LicenseNumber string      `xml:"LicenseNumber,attr"` //use="required" //The driving license number of the booking traveler.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] A simple container to specify the profiles that were applied to a reservation.
type AppliedProfile struct {
	Key                 com.TypeRef `xml:"Key,attr,omitempty"`                 //use="optional" //Key for update/delete of the element
	TravelerID          string      `xml:"TravelerID,attr"`                    //The ID of the TravelerProfile that was applied
	TravelerName        string      `xml:"TravelerName,attr"`                  //The name from the TravelerProfile that was applied
	AccountID           string      `xml:"AccountID,attr"`                     //The ID of the AccountProfile that was applied
	AccountName         string      `xml:"AccountName,attr"`                   //The name from the AccountProfile that was applied
	ImmediateParentID   string      `xml:"ImmediateParentID,attr,omitempty"`   //use="optional" //The ID of the immediate parent that was applied
	ImmediateParentName string      `xml:"ImmediateParentName,attr,omitempty"` //use="optional" //The name of the immediate parent that was applied
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Customized Name Data is used to print customized name on the different documents.
type CustomizedNameData struct {
	Value                      string      `xml:",innerxml"`
	Key                        com.TypeRef `xml:"Key,attr,omitempty"`                        //use="optional"
	ProviderReservationInfoRef com.TypeRef `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional"
}

//[RS] Travel Compliance and Preferred Supplier information of the traveler specific to a segment.
type TravelComplianceData struct {
	PolicyCompliance      []*PolicyCompliance   `xml:"PolicyCompliance,omitempty"`           //minOccurs="0" maxOccurs="2"
	ContractCompliance    []*ContractCompliance `xml:"ContractCompliance,omitempty"`         //minOccurs="0" maxOccurs="2"
	PreferredSupplier     []*PreferredSupplier  `xml:"PreferredSupplier,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	Key                   com.TypeRef           `xml:"Key,attr,omitempty"`                   //use="optional" //System generated key, returned back in the response. This can be used to modify or delete a saved TravelComplianceData.
	AirSegmentRef         com.TypeRef           `xml:"AirSegmentRef,attr,omitempty"`         //use="optional" //Refers to Air Segment. Applicable only for Air. Ignored for others.
	PassiveSegmentRef     com.TypeRef           `xml:"PassiveSegmentRef,attr,omitempty"`     //use="optional" //Refers to Passive Segment. Applicable only for Passive. Ignored for others.
	RailSegmentRef        com.TypeRef           `xml:"RailSegmentRef,attr,omitempty"`        //use="optional" //Refers to Rail Segment. Applicable only for Rail. Ignored for others.
	ReservationLocatorRef com.TypeLocatorCode   `xml:"ReservationLocatorRef,attr,omitempty"` //use="optional" //This is returned in the response. Any input will be ignored for this attribute. This represents the association of Travel Compliance Data with the uAPI reservation locator code, mainly relevant to Hotel and Vehicle.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS]
type PolicyCompliance struct {
	InPolicy    bool                   `xml:"InPolicy,attr"`              //use="required" //Policy Compliance Indicator. For In-Policy set to 'true', For Out-Of-Policy set to 'false''.
	PolicyToken com.StringLength1to128 `xml:"PolicyToken,attr,omitempty"` //use="optional" //Optional text message to set the rule or token for which it's In Policy or Out Of Policy.
}

//[RS]
type ContractCompliance struct {
	InContract    bool                   `xml:"InContract,attr"`              //use="required" //Contract Compliance Indicator. For In-Contract set to 'true', For Out-Of-Contract set to 'false'.
	ContractToken com.StringLength1to128 `xml:"ContractToken,attr,omitempty"` //use="optional" //Optional text message to set the rule or token for which it's In Contract or Out Of Contract.
}

//[RS]
type PreferredSupplier struct {
	Preferred   bool   `xml:"Preferred,attr"`   //use="required" //referred Supplier - 'true', 'false'.
	ProfileType string `xml:"ProfileType,attr"` //use="required" //Indicate profile type. e.g. if Agency Preferred then pass Agency, if Traveler Preferred then pass Traveler.
}

//[RS] Traveler information details like Travel Purpose and Trip Name
type TravelInfo struct {
	TripName      string `xml:"TripName,attr,omitempty"`      //use="optional" //Trip Name
	TravelPurpose string `xml:"TravelPurpose,attr,omitempty"` //use="optional" //Purpose of the trip
}

//[RS]
type BaseReservation struct {
	AccountingRemark []*AccountingRemark `xml:"AccountingRemark,omitempty"`    //minOccurs="0" maxOccurs="unbounded"
	GeneralRemark    []*GeneralRemark    `xml:"GeneralRemark,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	Restriction      []*Restriction      `xml:"Restriction,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	PassiveInfo      *PassiveInfo        `xml:"PassiveInfo,omitempty"`         //minOccurs="0"
	LocatorCode      com.TypeLocatorCode `xml:"LocatorCode,attr"`              //use="required" //The unique identifier for this reservation. If this is this View Only UR LocatorCode is '999999'.
	CreateDate       string              `xml:"CreateDate,attr"`               //use="required" //The date and time that this reservation was created.
	ModifiedDate     string              `xml:"ModifiedDate,attr"`             //use="required" //The date and time that this reservation was last modified for any reason.
	CustomerNumber   string              `xml:"CustomerNumber,attr,omitempty"` //use="optional"
}

//[RS] An accounting remark container to hold any printable text.
type AccountingRemark struct {
	RemarkData                 string                      `xml:"RemarkData"`                                //Actual remarks data.
	BookingTravelerRef         com.TypeRef                 `xml:"BookingTravelerRef,omitempty"`              //minOccurs="0" maxOccurs="unbounded" //Reference to Booking Traveler.
	Key                        com.TypeRef                 `xml:"Key,attr,omitempty"`                        //use="optional"
	Category                   string                      `xml:"Category,attr,omitempty"`                   //use="optional" //A category to group and organize the various remarks. This is not required, but it is recommended.
	TypeInGds                  com.TypeGdsAccountingRemark `xml:"TypeInGds,attr,omitempty"`                  //use="optional"
	ProviderReservationInfoRef com.TypeRef                 `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode        `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this accounting remark is used
	UseProviderNativeMode      bool                        `xml:"UseProviderNativeMode,attr,omitempty"`      //use="optional" default="false" //Will be true when terminal process required, else false
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS] Which activities are supported for a particular element
type Restriction struct {
	Operation string `xml:"Operation,attr"`        //use="required" //The operation that is restricted
	Reason    string `xml:"Reason,attr,omitempty"` //use="optional" //The reason it is restricted
}

//[RS] Used by CreateReservationReq for passing in elements normally found post-booking
type PassiveInfo struct {
	TicketNumber        []string    `xml:"TicketNumber,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	ConfirmationNumber  []string    `xml:"ConfirmationNumber,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	Commission          *Commission `xml:"Commission,omitempty"`               //minOccurs="0"
	ProviderCode        string      `xml:"ProviderCode,attr,omitempty"`        //use="optional"
	ProviderLocatorCode string      `xml:"ProviderLocatorCode,attr,omitempty"` //use="optional"
	SupplierCode        string      `xml:"SupplierCode,attr,omitempty"`        //use="optional"
	SupplierLocatorCode string      `xml:"SupplierLocatorCode,attr,omitempty"` //use="optional"
}

//[RS] Locator code on the host carrier system
type SupplierLocator struct {
	SegmentRef                 []*GeneralReference `xml:"SegmentRef,omitempty"`                      //minOccurs="0" maxOccurs="unbounded" //Air/Passive Segment Reference
	SupplierCode               com.TypeCarrier     `xml:"SupplierCode,attr"`                         //use="required"
	SupplierLocatorCode        string              `xml:"SupplierLocatorCode,attr"`                  //use="required" //Carrier reservation locator code
	ProviderReservationInfoRef com.TypeRef         `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider Reservation  reference
	CreateDateTime             string              `xml:"CreateDateTime,attr,omitempty"`             //use="optional" //The Date and Time which the reservation is received from the Vendor as a SupplierLocator creation Da
}

//[RS]
type GeneralReference struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//Tracks the various agent/agency information
type AgencyInfo struct {
	AgentAction []*AgentAction `xml:"AgentAction"` //maxOccurs="unbounded"
}

//Depending on context, this will represent information about which agent perform different actions.
type AgentAction struct {
	ActionType string             `xml:"ActionType,attr"`          //use="required" //The type of action the agent performed.
	AgentCode  string             `xml:"AgentCode,attr"`           //use="required" //The AgenctCode who performed the action.
	BranchCode com.TypeBranchCode `xml:"BranchCode,attr"`          //use="required" //The BranchCode of the branch (working branch, branchcode used for the request. If nothing specified, branchcode for the agent) who performed the action.
	AgencyCode string             `xml:"AgencyCode,attr"`          //use="required" //The AgencyCode of the agent who performed the action.
	AgentSine  string             `xml:"AgentSine,attr,omitempty"` //use="optional" //The sign in user name of the agent logged into the terminal. PROVIDER SUPPORTED: ACH
	EventTime  string             `xml:"EventTime,attr"`           //use="required" //Date and time at which this event took place.
	//<xs:attributeGroup name="attrAgentOverride">
	AgentOverride string `xml:"AgentOverride,attr,omitempty"` //use="optional" //AgentSine value that was used during PNR creation or End Transact.
	//</xs:attributeGroup>
}

//[RS] Generic agency contact information container. It must contain at least one phone number to be used by an agency
type AgencyContactInfo struct {
	PhoneNumber []*PhoneNumber `xml:"PhoneNumber"`        //maxOccurs="unbounded"
	Key         com.TypeRef    `xml:"Key,attr,omitempty"` //use="optional"
}

//[RS] Status of the action that will happen or has happened to the air reservation. One Action status for each provider reservation
type ActionStatus struct {
	Remark                     *Remark           `xml:"Remark,omitempty"`                          //minOccurs="0"
	Type                       string            `xml:"Type,attr"`                                 //use="required" //Identifies the type of action (if any) to take on this air reservation. Only TTL, TAU, TAX and TAW can be set by the user.
	TicketDate                 string            `xml:"TicketDate,attr,omitempty"`                 //use="optional" //Identifies when the action type will happen, or has happened according to the type.
	Key                        com.TypeRef       `xml:"Key,attr,omitempty"`                        //use="optional" //Identifies when the action type will happen, or has happened according to the type.
	ProviderReservationInfoRef com.TypeRef       `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	QueueCategory              com.TypeNonBlanks `xml:"QueueCategory,attr,omitempty"`              //use="optional" //Add Category placement to ticketing queue (required in 1P - default is 00)
	AirportCode                com.TypeAirport   `xml:"AirportCode,attr,omitempty"`                //use="optional" //Used with Time Limit to specify the airport location where the ticket is to be issued.
	//<xs:attributeGroup ref="attrProviderSupplier">
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional" //Code of the supplier returning this TaxInfo.
	//</xs:attributeGroup>
	PseudoCityCode com.TypePCC `xml:"PseudoCityCode,attr,omitempty"` //use="optional" //The Branch PCC in the host system where PNR can be queued for ticketing. When used with TAU it will auto queue to Q10. When used with TAW agent performs manual move to Q.
	AccountCode    string      `xml:"AccountCode,attr,omitempty"`    //use="optional" //Used with TAW. Used to specify a corporate or in house account code to the PNR as part of ticketing arrangement field.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//[RS]
type SegmentRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RS] Marketing text or Notices for Suppliers
type MarketingInformation struct {
	Text []string `xml:"Text"` //maxOccurs="unbounded"
}

//Location definition specific to a Vendor in a specific provider (e.g. 1G) system.
type VendorLocation struct {
	ProviderCode     com.TypeProviderCode `xml:"ProviderCode,attr"`               //use="required" //The code of the provider (e.g. 1G, 1S)
	VendorCode       com.TypeSupplierCode `xml:"VendorCode,attr"`                 //use="required" //The code of the vendor (e.g. HZ, etc.)
	PreferredOption  bool                 `xml:"PreferredOption,attr,omitempty"`  //use="optional" //Preferred Option marker for Location.
	VendorLocationID string               `xml:"VendorLocationID,attr,omitempty"` //use="optional" //Location identifier
	Key              com.TypeRef          `xml:"Key,attr,omitempty"`              //use="optional" //Key which maps vendor location with vehicles
}

//These are zero or more negotiated rate codes
type CorporateDiscountID struct {
	Value              string `xml:",innerxml"`
	NegotiatedRateCode bool   `xml:"NegotiatedRateCode,attr,omitempty"` //use="optional" //When set to true, the data in the CorporateDiscountID is a negotiated rate code. Otherwise, this data is a Corporate Discount ID rate.
}

//The attributes and elements in a SubKey.
type TypeSubKey struct {
	Text        []string `xml:"Text,omitempty"`             //minOccurs="0" maxOccurs="999" //Information for a sub key.
	Name        string   `xml:"Name,attr"`                  //use="required" //A subkey to identify the specific information within this keyword
	Description string   `xml:"Description,attr,omitempty"` //A brief description of a subkey.
}

//A complexType for keyword information.
type TypeKeyword struct {
	SubKey      []*TypeSubKey `xml:"SubKey,omitempty"`           //minOccurs="0" maxOccurs="99" //A further breakdown of a keyword.
	Text        []string      `xml:"Text,omitempty"`             //minOccurs="0" maxOccurs="unbounded" //Information for a keyword.
	Name        string        `xml:"Name,attr"`                  //use="required" //The keyword name.
	Number      string        `xml:"Number,attr,omitempty"`      //The number for this keyword.
	Description string        `xml:"Description,attr,omitempty"` //A brief description of the keyword
}

//The attributes and elements in a SubKey.
type TypeOTASubKey struct {
	Text        []string        `xml:"Text,omitempty"`             //minOccurs="0" maxOccurs="999" //Information for a sub key.
	Name        com.TypeOTACode `xml:"Name,attr"`                  //use="required" //A subkey to identify the special equipment codes. Applicable when Policy/@Name is EQUIP. Uses OTA CODE "EQP". 1P/1J.
	Description string          `xml:"Description,attr,omitempty"` //A brief description of a subkey.
}

//A complexType for keyword information.
type TypeOTAKeyword struct {
	SubKey      []*TypeOTASubKey `xml:"SubKey,omitempty"`           //minOccurs="0" maxOccurs="99" //A further breakdown of a keyword.
	Text        []string         `xml:"Text,omitempty"`             //minOccurs="0" maxOccurs="unbounded" //Information for a keyword.
	Name        string           `xml:"Name,attr"`                  //use="required" //The keyword name.
	Number      string           `xml:"Name,attr,omitempty"`        //The number for this keyword.
	Description string           `xml:"Description,attr,omitempty"` //A brief description of the keyword
}

//Detail information of keywords.
type Keyword struct {
	TypeKeyword
	Text []string `xml:"Text,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Information for a keyword.
}

type TypeSegmentRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

type TypeNonAirReservationRef struct {
	LocatorCode com.TypeLocatorCode `xml:"LocatorCode,attr"` //use="required"
}

//Third party supplier locator information. Specifically applicable for SDK booking.
type ThirdPartyInformation struct {
	SegmentRef                 []*GeneralReference        `xml:"GeneralReference,omitempty"`                //minOccurs="0" maxOccurs="unbounded" //Air/Passive Segment Reference
	ThirdPartyCode             string                     `xml:"ThirdPartyCode,attr"`                       //Third party supplier code.
	ThirdPartyLocatorCode      string                     `xml:"ThirdPartyLocatorCode,attr"`                //Confirmation number for third party supplier.
	ThirdPartyName             com.TypeThirdPartySupplier `xml:"ThirdPartyName,attr"`                       //Third party supplier name.
	ProviderReservationInfoRef com.TypeRef                `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider Reservation  reference
	Key                        com.TypeRef                `xml:"Key,attr"`                                  //Unique identifier of the third party supplier. Key can be used to modify or delete saved third party information.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Travel Agency Service Fees (TASF) are charged by the agency through BSP or Airline Reporting Corporation (ARC).
type ServiceFeeInfo struct {
	FormOfPayment *FormOfPayment `xml:"FormOfPayment,omitempty"` //minOccurs="0"
	//ServiceFeeTaxInfo *ServiceFeeTaxInfo `xml:"ServiceFeeTaxInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	//CreditCardAuth    *CreditCardAuth    `xml:"CreditCardAuth,omitempty"`    //minOccurs="0"
	Payment                           *Payment      `xml:"Payment,omitempty"`                                //minOccurs="0"
	Status                            string        `xml:"Status,attr,omitempty"`                            //use="optional" //Status of the service fee. Possible Values - Issued,ReadyToIssue,IssueLater.
	Description                       string        `xml:"Description,attr,omitempty"`                       //use="optional" //The description of the service fee.
	Key                               com.TypeRef   `xml:"Key,attr,omitempty"`                               //use="optional"
	Confirmation                      string        `xml:"Confirmation,attr,omitempty"`                      //use="optional" /The confirmation number of the service fee in the merchant host system.
	TicketNumber                      string        `xml:"TicketNumber,attr,omitempty"`                      //use="optional" //The ticket that this fee was issued in connection with.
	BookingTravelerRef                com.TypeRef   `xml:"BookingTravelerRef,attr,omitempty"`                //use="optional" //A reference to a passenger.
	ProviderReservationInfoRef        com.TypeRef   `xml:"ProviderReservationInfoRef,attr,omitempty"`        //use="optional" //A reference to the provider reservation info to which the service is tied.
	PassiveProviderReservationInfoRef com.TypeRef   `xml:"PassiveProviderReservationInfoRef,attr,omitempty"` //use="optional" //A reference to the passive provider reservation info to which the service is tied.
	TotalAmount                       com.TypeMoney `xml:"TotalAmount,attr,omitempty"`                       //use="optional" //The total amount for this Service Fee including base amount and all taxes.
	BaseAmount                        com.TypeMoney `xml:"BaseAmount,attr,omitempty"`                        //use="optional" //Represents the base price for this entity. This does not include any taxes.
	Taxes                             com.TypeMoney `xml:"Taxes,attr,omitempty"`                             //use="optional" //The aggregated amount of all the taxes that are associated with this entity. See the associated Service Fee TaxInfo array for a breakdown of the individual taxes.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	BookingTravelerName string `xml:"BookingTravelerName,attr,omitempty"` //use="optional"  //The name of the passenger.
}

//Other Service information sent to the carriers during air bookings
type OSI struct {
	Key                        com.TypeRef          `xml:"Key,attr,omitempty"`                        //use="optional"
	Carrier                    com.TypeCarrier      `xml:"Carrier,attr"`                              //use="required"
	Code                       string               `xml:"Code,attr,omitempty"`                       //use="optional"
	Text                       string               `xml:"Text,attr"`                                 //use="required"
	ProviderReservationInfoRef com.TypeRef          `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this OSI is used
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Represents a traveler group for Group booking and all their accompanying data. SUPPORTED PROVIDER: Worldspan and JAL.
type Group struct {
	Name               *Name                  `xml:"Name"`                         //Name of the group in group booking.
	DeliveryInfo       *DeliveryInfo          `xml:"DeliveryInfo,omitempty"`       //minOccurs="0" maxOccurs="1"
	PhoneNumber        []*PhoneNumber         `xml:"PhoneNumber,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	SSRRef             []*SSRRef              `xml:"SSRRef,omitempty"`             //minOccurs="0" maxOccurs="unbounded" //Reference Element for SSR.
	Address            *TypeStructuredAddress `xml:"Address,omitempty"`            //minOccurs="0" maxOccurs="1"
	BookingTravelerRef []*BookingTravelerRef  `xml:"BookingTravelerRef,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Reference Element for Booking Traveler.
	Key                com.TypeRef            `xml:"Key,attr,omitempty"`           //use="optional"
	TravelerType       com.TypePTC            `xml:"TravelerType,attr,omitempty"`  //use="optional" //Defines the type of traveler used for booking which could be a non-defining type (Companion, Web-fare, etc), or a standard type (Adult, Child, etc).
	GroupSize          uint                   `xml:"GroupSize,attr"`               //use="required" //Represents size of the group
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Reference Element for SSR.
type SSRRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//[RS] A Form of Payment used to purchase all or part of a booking.
type FormOfPayment struct {
	//<xs:choice>
	CreditCard *CreditCard `xml:"CreditCard,omitempty"` //minOccurs="0"
	DebitCard  *DebitCard  `xml:"DebitCard,omitempty"`  //minOccurs="0"
	EnettVan   *EnettVan   `xml:"EnettVan,omitempty"`   //minOccurs="0"
	//<xs:group name="FormOfPaymentSequenceGroup">
	//<xs:choice>
	Certificate       *Certificate       `xml:"Certificate,omitempty"`       //minOccurs="0"
	TicketNumber      com.TicketNumber   `xml:"TicketNumber,omitempty"`      //minOccurs="0"
	Check             *Check             `xml:"Check,omitempty"`             //minOccurs="0"
	Requisition       *Requisition       `xml:"Requisition,omitempty"`       //minOccurs="0"
	MiscFormOfPayment *MiscFormOfPayment `xml:"MiscFormOfPayment,omitempty"` //minOccurs="0"
	AgencyPayment     *AgencyPayment     `xml:"AgencyPayment,omitempty"`     //minOccurs="0"
	UnitedNations     *UnitedNations     `xml:"UnitedNations,omitempty"`     //minOccurs="0"
	DirectPayment     *DirectPayment     `xml:"DirectPayment,omitempty"`     //minOccurs="0"
	AgentVoucher      *AgentVoucher      `xml:"AgentVoucher,omitempty"`      //minOccurs="0"
	PaymentAdvice     *PaymentAdvice     `xml:"PaymentAdvice,omitempty"`     //minOccurs="0"
	//</xs:choice>
	//</xs:group>
	//</xs:choice>
	ProviderReservationInfoRef []*ProviderReservationInfoRef `xml:"ProviderReservationInfoRef,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	SegmentRef                 []*GeneralReference           `xml:"SegmentRef,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	Key                        com.TypeRef                   `xml:"Key,attr,omitempty"`                   //use="optional"
	Type                       string                        `xml:"Type,attr"`                            //use="required" //Allowable values are "Certificate" "Cash" "Credit" "Check" "Ticket" "Debit" "Invoice" "Requisition" "MiscFormOfPayment" "AgencyPayment" "DirectBill" "UnitedNations" "DirectPayment" "AgentVoucher" "AccountReceivable" "AgentNonRefundable" "Enett"
	FulfillmentType            string                        `xml:"FulfillmentType,attr,omitempty"`       //use="optional" //Defines how the client wishes to receive travel documents. Type does not define where or how payment is made. The supported values are "Ticket on Departure","Travel Agency","Courier","Standard Mail","Ticketless","Ticket Office","Express Mail","Corporate Kiosk","Train Station Service Desk","Direct Printing of Ticket","Ticket by Email","Digital Printing of Ticket at Home","Retrieve Ticket at Eurostar in London" Collect booking ticket at a Kiosk, print in agency.
	FulfillmentLocation        string                        `xml:"FulfillmentLocation,attr,omitempty"`   //use="optional" //Information about the location of the printer.
	FulfillmentIDType          string                        `xml:"FulfillmentIDType,attr,omitempty"`     //use="optional" //Identification type, e.g. credit card, to define how the customer will identify himself when collecting the ticket
	FulfillmentIDNumber        string                        `xml:"FulfillmentIDNumber,attr,omitempty"`   //use="optional" //Identification number, e.g. card number, to define how the customer will identify himself when collecting the ticket
	IsAgentType                bool                          `xml:"IsAgentType,attr,omitempty"`           //use="optional" //default="false" //If this is true then FormOfPayment mention in Type is anAgent type FormOfPayment.
	AgentText                  string                        `xml:"AgentText,attr,omitempty"`             //use="optional" //This is only relevent when IsAgentType is specified as true. Otherwise this will be ignored.
	ReuseFOP                   com.TypeRef                   `xml:"ReuseFOP,attr,omitempty"`              //use="optional" //Key of the FOP Key to be reused as this Form of Payment.Only Credit and Debit Card will be supported for FOP Reuse.
	ExternalReference          com.TypeExternalReference     `xml:"ExternalReference,attr,omitempty"`     //use="optional"
	Reusable                   bool                          `xml:"Reusable,attr,omitempty"`              //use="optional" default="false" //Indicates whether the form of payment can be reused or not. Currently applicable for Credit and Debit form of payment
	ProfileID                  string                        `xml:"ProfileID,attr,omitempty"`             //The unique ID of the profile that contains the payment details to use.
	ProfileKey                 com.TypeRef                   `xml:"ProfileKey,attr,omitempty"`            //The Key assigned to the payment details value from the specified profile.
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Container for all eNett Van information.
type EnettVan struct {
	MinPercentage com.TypeIntegerPercentage  `xml:"MinPercentage,attr,omitempty"` //use="optional" //The minimum percentage that will be applied on the Total price and sent to enett,which will denote the minimum authorized amount approved by eNett.uApi will default this to zero for multi-use Van's.
	MaxPercentage com.TypeIntegerPercentage  `xml:"MaxPercentage,attr,omitempty"` //use="optional" //The maximum percentage that will be applied on the Total price and sent to enett, which will denote the maximum authorized amount as approved by eNett. This value will be ignored and not used for Multi-Use VAN’s.
	ExpiryDays    com.TypeDurationYearInDays `xml:"ExpiryDays,attr,omitempty"`    //use="optional" //The number of days from the VAN generation date that the VAN will be active for, after which the VAN cannot be used.
}

//Container for all debit card information.
type DebitCard struct {
	PaymentCard
	IssueNumber string `xml:"IssueNumber,attr,omitempty"` //use="optional" //Verification number for Debit Cards //maxLength value="8"
	//<xs:attributeGroup name="attrAppliedProfilePaymentInfo"> //ProfileID and Key are required in order to reference a payment method from a profile.
	ProfileID string      `xml:"ProfileID,attr,omitempty"` //The unique ID of the profile that contains the payment details to use.
	Key       com.TypeRef `xml:"Key,attr,omitempty"`       //The Key assigned to the payment details value from the specified profile.
	//</xs:attributeGroup>
}

//Container for all credit card information.
type CreditCard struct {
	CreditCardType
	//<xs:attributeGroup name="attrAppliedProfilePaymentInfo"> //ProfileID and Key are required in order to reference a payment method from a profile.
	ProfileID string      `xml:"ProfileID,attr,omitempty"` //The unique ID of the profile that contains the payment details to use.
	Key       com.TypeRef `xml:"Key,attr,omitempty"`       //The Key assigned to the payment details value from the specified profile.
	//</xs:attributeGroup>
}

type CreditCardType struct {
	PaymentCard
	ExtendedPayment    string          `xml:"ExtendedPayment,attr,omitempty"`    //use="optional" //Used for American Express cards.
	CustomerReference  string          `xml:"CustomerReference,attr,omitempty"`  //use="optional" //Agencies use this to pass the traveler information to the credit card company.
	AcceptanceOverride bool            `xml:"AcceptanceOverride,attr,omitempty"` //use="optional" //Override airline restriction on the credit card.
	ThirdPartyPayment  bool            `xml:"ThirdPartyPayment,attr,omitempty"`  //use="optional" //default="false" //If true, this indicates that the credit card holder is not one of the passengers.
	BankName           string          `xml:"BankName,attr,omitempty"`           //Issuing bank name for this credit card
	BankCountryCode    com.TypeCountry `xml:"BankCountryCode,attr,omitempty"`    //ISO Country code associated with the issuing bank
	BankStateCode      com.TypeState   `xml:"BankStateCode,attr,omitempty"`      //State code associated with the issuing bank.
	Enett              bool            `xml:"Enett,attr,omitempty"`              //use="optional" //default="false" //Acceptable values are true or false. If set to true it will denote that the credit card used has been issued through Enett. For all other credit card payments this value will be set to false.
}

//Container for all credit and debit card information.
type PaymentCard struct {
	PhoneNumber    *PhoneNumber             `xml:"PhoneNumber,omitempty"`       //minOccurs="0"
	BillingAddress *TypeStructuredAddress   `xml:"BillingAddress,omitempty"`    //minOccurs="0" //The address to where the billing statements for this card are sent. Used for address verification purposes.
	Type           com.TypeCardMerchantType `xml:"Type,attr,omitempty"`         //use="optional"
	Number         com.TypeCreditCardNumber `xml:"Number,attr,omitempty"`       //use="optional"
	ExpDate        string                   `xml:"ExpDate,attr,omitempty"`      //use="optional" //The Expiration date of this card in YYYY-MM format.
	Name           string                   `xml:"Name,attr,omitempty"`         //use="optional" //The name as it appears on the card.
	CVV            string                   `xml:"CVV,attr,omitempty"`          //use="optional" //Card Verification Code.length = 4
	ApprovalCode   string                   `xml:"ApprovalCode,attr,omitempty"` //use="optional" //This code is required for an authorization process from the Credit Card company directly,required for some of the CCH carriers.This attribute is also used for EMD retrieve and issuance transactions.
}

//Certificate Form of Payment
type Certificate struct {
	Number             string        `xml:"Number,attr"`                       //use="required" //The Certificate number
	Amount             com.TypeMoney `xml:"Amount,attr,omitempty"`             //use="optional" //The monetary value of the certificate.
	DiscountAmount     com.TypeMoney `xml:"DiscountAmount,attr,omitempty"`     //use="optional" //The monetary discount amount of this certificate.
	DiscountPercentage int           `xml:"DiscountPercentage,attr,omitempty"` //use="optional" //The percentage discount value of this certificate.
	NotValidBefore     string        `xml:"NotValidBefore,attr,omitempty"`     //use="optional" type="xs:date" //The date that this certificate becomes valid.
	NotValidAfter      string        `xml:"NotValidAfter,attr,omitempty"`      //use="optional" type="xs:date" //The date that this certificate expires.
}

//Check Form of Payment
type Check struct {
	MICRNumber    string `xml:"MICRNumber,attr,omitempty"`    //use="optional" //Magnetic Ink Character Reader Number of check.
	RoutingNumber string `xml:"RoutingNumber,attr,omitempty"` //use="optional" //The bank routing number of the check.
	AccountNumber string `xml:"AccountNumber,attr,omitempty"` //use="optional" //The account number of the check
	CheckNumber   string `xml:"CheckNumber,attr,omitempty"`   //use="optional" //The sequential check number of the check.
}

//Requisition Form of Payment
type Requisition struct {
	Number   string `xml:"Number,attr,omitempty"`   //use="optional" //Requisition number used for accounting
	Category string `xml:"Category,attr,omitempty"` //use="optional" //Classification Category for the requisition payment: Government,Other
	Type     string `xml:"Type,attr,omitempty"`     //use="optional" //Type can be Cash or Credit for category as Government: Cash,Credit
}

//Miscellaneous Form of Payments
type MiscFormOfPayment struct {
	CreditCardType     string                   `xml:"CreditCardType,attr,omitempty"`     //use="optional" //The 2 letter credit/ debit card type or code which may not have been issued using the standard bank card types - i.e. an airline issued card
	CreditCardNumber   com.TypeCreditCardNumber `xml:"CreditCardNumber,attr,omitempty"`   //use="optional"
	ExpDate            string                   `xml:"ExpDate,attr,omitempty"`            //use="optional" type="xs:gYearMonth" //The Expiration date of this card in YYYY-MM format.
	Text               string                   `xml:"Text,attr,omitempty"`               //use="optional" //Any free form text which may be associated with the Miscellaneous Form of Payment. This text may be provider or GDS specific
	Category           string                   `xml:"Category,attr"`                     //use="required" //Indicates what Category the Miscellaneous Form Of Payment is being used for payment - The category may vary by GDS.Allowable values are "Text" "Credit" "CreditCard" "FreeFormCreditCard" "Invoice" "NonRefundable" "MultipleReceivables" "Exchange" "Cash"
	AcceptanceOverride bool                     `xml:"AcceptanceOverride,attr,omitempty"` //use="optional" //Override airline restriction on the credit card.
}

//Type for Agency Payment.
type AgencyPayment struct {
	AgencyBillingIdentifier string `xml:"AgencyBillingIdentifier,attr"`         //use="required" //Value of the billing id
	AgencyBillingNumber     string `xml:"AgencyBillingNumber,attr,omitempty"`   //use="optional" //Value of billing number
	AgencyBillingPassword   string `xml:"AgencyBillingPassword,attr,omitempty"` //use="optional" //Value of billing password
}

//United Nations Form of Payments
type UnitedNations struct {
	Number string `xml:"Number,attr"` //use="required"
}

//Direct Payment Form of Payments
type DirectPayment struct {
	Text string `xml:"Text,attr,omitempty"` //use="optional"
}

//Agent Voucher Form of Payments
type AgentVoucher struct {
	Number string `xml:"Number,attr"` //use="required"
}

//Contains other form of payment for Cruise Reservations
type PaymentAdvice struct {
	Type           string           `xml:"Type,attr"`                  //use="required" //Other Payment Yype. Possible Values: AGC - Agency Check, AGG - Agency Guarantee, AWC - Award Check, CSH - Cash Equivalent, DBC - Denied Boarding Compensation, MCO - Miscellaneous Charge Order, TOO - Tour Order, TOV - Tour Voucher
	DocumentNumber string           `xml:"DocumentNumber,attr"`        //use="required" //Payment Document Number Examples: 1234567890, R7777
	IssueDate      string           `xml:"IssueDate,attr"`             //use="required" type="xs:date" //Document Issuance date
	IssueCity      com.TypeIATACode `xml:"IssueCity,attr"`             //use="required" //City code of document issuance
	OriginalFOP    string           `xml:"OriginalFOP,attr,omitempty"` //use="optional" //Original form of payment Examples: CHECK 3500
}

//Review Booking or Queue Minders is to add the reminders in the Provider Reservation along with the date time and Queue details. On the date time defined in reminders, the message along with the PNR goes to the desired Queue.
type ReviewBooking struct {
	Key                        com.TypeRef          `xml:"Key,attr,omitempty"`                        //use="optional" //Returned in response. Use it for update of saved review booking.
	Queue                      int                  `xml:"Queue,attr"`                                //use="required" //Queue number, Must be numeric and less than 100.
	QueueCategory              string               `xml:"QueueCategory,attr,omitempty"`              //use="optional" //Queue Category, 2 Character Alpha or Numeric.
	DateTime                   string               `xml:"DateTime,attr"`                             //use="required" //Date and Time to place message on designated Queue, Should be prior to the last segment date in the PNR.
	PseudoCityCode             com.TypePCC          `xml:"PseudoCityCode,attr,omitempty"`             //use="optional" //Input PCC optional value for placing the PNR into Queue. If not passed, will add as default PNR's Pseudo.
	ProviderCode               com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`               //use="optional" //The code of the Provider (e.g 1G,1V).
	ProviderReservationInfoRef com.TypeRef          `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider Reservation reference. Returned in the response. Use it for update of saved Review Booking.
	Remarks                    string               `xml:"Remarks,attr"`                              //use="required" //Remark or reminder message. It can be truncated depending on the provider. max length = 300
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Authorization remark for Consolidator access to a PNR. Contains PCC information created by retail agent to allow a consolidator or other Axess users to service their PNR. PROVIDER SUPPORTED: Worldspan and JAL.
type ConsolidatorRemark struct {
	PseudoCityCode             []*PseudoCityCode    `xml:"PseudoCityCode"`                            //minOccurs="1" maxOccurs="5"
	Key                        com.TypeRef          `xml:"Key,attr,omitempty"`                        //use="optional" //Key to be used for internal processing.
	ProviderReservationInfoRef com.TypeRef          `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this element is used
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

type InvoiceRemark struct {
	AssociatedRemark
	//<xs:choice>
	AirSegmentRef         *TypeSegmentRef           `xml:"AirSegmentRef,omitempty"`         //minOccurs="0" maxOccurs="1" //Reference to AirSegment from an Air Reservation.
	HotelReservationRef   *TypeNonAirReservationRef `xml:"HotelReservationRef,omitempty"`   //minOccurs="0" maxOccurs="1" //Specify the locator code of Hotel reservation.
	VehicleReservationRef *TypeNonAirReservationRef `xml:"VehicleReservationRef,omitempty"` //minOccurs="0" maxOccurs="1" //Specify the locator code of Vehicle reservation.
	PassiveSegmentRef     *TypeSegmentRef           `xml:"PassiveSegmentRef,omitempty"`     //minOccurs="0" maxOccurs="1" //Reference to PassiveSegment from a Passive Reservation.
	//</xs:choice>
}

//Contains the PCC name which can access or modify the PNR.
type PseudoCityCode struct {
	Value string `xml:",innerxml"`
}

//A textual remark container to hold Associated itinerary remarks
type AssociatedRemark struct {
	RemarkWithTravelerRef
	Key com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Identifies the agency commision remarks. Specifically used for Worldspan.
type CommissionRemark struct {
	//<xs:choice>
	ProviderReservationLevel *ProviderReservationLevel `xml:"ProviderReservationLevel,omitempty"` //Specify commission which is applicable to PNR level.
	PassengerTypeLevel       []*PassengerTypeLevel     `xml:"PassengerTypeLevel,omitempty"`       //Specify commission which is applicable to per PTC level.
	//</xs:choice>
	Key                        com.TypeRef          `xml:"Key,attr,omitempty"`                        //use="optional" //Key to be used for internal processing.
	ProviderReservationInfoRef com.TypeRef          `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this accounting remark is used
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//Specify commission which is applicable to PNR level.
type ProviderReservationLevel struct {
	//<xs:attributeGroup name="attrCommissionRemark">
	Amount        com.TypeMoney                 `xml:"Amount,attr,omitempty"`        //use="optional" //The monetary amount of the commission.
	Percentage    com.TypePercentageWithDecimal `xml:"Percentage,attr,omitempty"`    //use="optional" //The percent of the commission.
	CommissionCap com.TypeMoney                 `xml:"CommissionCap,attr,omitempty"` //use="optional" //Commission cap for the Airline.
	//</xs:attributeGroup>
}

//Specify commission which is applicable to per PTC level.
type PassengerTypeLevel struct {
	TravelerType com.TypePTC `xml:"TravelerType,attr"` //use="required"
	//<xs:attributeGroup name="attrCommissionRemark">
	Amount        com.TypeMoney                 `xml:"Amount,attr,omitempty"`        //use="optional" //The monetary amount of the commission.
	Percentage    com.TypePercentageWithDecimal `xml:"Percentage,attr,omitempty"`    //use="optional" //The percent of the commission.
	CommissionCap com.TypeMoney                 `xml:"CommissionCap,attr,omitempty"` //use="optional" //Commission cap for the Airline.
	//</xs:attributeGroup>
}

type TypeRemark struct {
	Value                      string               `xml:",innerxml"`
	ProviderReservationInfoRef com.TypeRef          `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this element is used
}

//Postscript Notes
type Postscript struct {
	TypeRemark
	Key com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
}

//A provider reservation field used to store customer information. It may be used to identify reservations which will/will not be available for access.
type CustomerID struct {
	TypeRemark
	Key com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
}

//List of invoices only for 1G/1V
type InvoiceData struct {
	BookingTravelerInformation []*BookingTravelerInformation `xml:"BookingTravelerInformation"` //maxOccurs="9"
	Key                        com.TypeRef                   `xml:"Key,attr,omitempty"`
	InvoiceNumber              string                        `xml:"InvoiceNumber,attr"`              //use="required" //Invoice number
	IssueDate                  string                        `xml:"IssueDate,attr,omitempty"`        //Invoice issue date
	ProviderReservationInfoRef com.TypeRef                   `xml:"ProviderReservationInfoRef,attr"` //use="required" //Provider reservation reference key.
}

//Booking Traveler information tied to invoice
type BookingTravelerInformation struct {
	Name               *Name       `xml:"Name"`
	BookingTravelerRef com.TypeRef `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //A reference to a passenger related to a ticket.
}

//Represents host ARNK segments.
type ProviderARNKSegment struct {
	PreviousSegment            *PreviousSegment `xml:"PreviousSegment,omitempty"`
	NextSegment                *NextSegment     `xml:"NextSegment,omitempty"`
	Key                        com.TypeRef      `xml:"Key,attr,omitempty"`                        //use="optional"
	ProviderReservationInfoRef com.TypeRef      `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderSegmentOrder       int              `xml:"ProviderSegmentOrder,attr,omitempty`        //use="optional" //To identify the appropriate travel sequence for Air/Car/Hotel/Rail segments/reservations in the provider reservation.
}

type PreviousSegment struct {
	//<xs:choice>
	AirSegmentRef         *TypeSegmentRef           `xml:"AirSegmentRef,omitempty"`         //minOccurs="0" maxOccurs="1" //Reference to AirSegment from an Air Reservation.
	HotelReservationRef   *TypeNonAirReservationRef `xml:"HotelReservationRef,omitempty"`   //minOccurs="0" maxOccurs="1" //Specify the locator code of Hotel reservation.
	VehicleReservationRef *TypeNonAirReservationRef `xml:"VehicleReservationRef,omitempty"` //minOccurs="0" maxOccurs="1" //Specify the locator code of Vehicle reservation.
	PassiveSegmentRef     *TypeSegmentRef           `xml:"PassiveSegmentRef,omitempty"`     //minOccurs="0" maxOccurs="1" //Reference to PassiveSegment from a Passive Reservation.
	//</xs:choice>
}

type NextSegment struct {
	//<xs:choice>
	AirSegmentRef         *TypeSegmentRef           `xml:"AirSegmentRef,omitempty"`         //minOccurs="0" maxOccurs="1" //Reference to AirSegment from an Air Reservation.
	HotelReservationRef   *TypeNonAirReservationRef `xml:"HotelReservationRef,omitempty"`   //minOccurs="0" maxOccurs="1" //Specify the locator code of Hotel reservation.
	VehicleReservationRef *TypeNonAirReservationRef `xml:"VehicleReservationRef,omitempty"` //minOccurs="0" maxOccurs="1" //Specify the locator code of Vehicle reservation.
	PassiveSegmentRef     *TypeSegmentRef           `xml:"PassiveSegmentRef,omitempty"`     //minOccurs="0" maxOccurs="1" //Reference to PassiveSegment from a Passive Reservation.
	//</xs:choice>
}

//Payment Guarantee: Guarantee, Deposit
type Guarantee struct {
	//<xs:choice>
	CreditCard         *CreditCard         `xml:"CreditCard,omitempty"`         //minOccurs="0"
	OtherGuaranteeInfo *OtherGuaranteeInfo `xml:"OtherGuaranteeInfo,omitempty"` //minOccurs="0"
	//</xs:choice>
	Type              string                    `xml:"Type,attr"`                        //use="required" //Guarantee only or Deposit
	Key               com.TypeRef               `xml:"Type,attr,omitempty"`              //use="optional" //Key for update/delete of the element
	ReuseFOP          com.TypeRef               `xml:"ReuseFOP,attr,omitempty"`          //use="optional" //Key of the FOP Key to be reused as this Form of Payment.Only Credit and Debit Card will be supported for FOP Reuse.
	ExternalReference com.TypeExternalReference `xml:"ExternalReference,attr,omitempty"` //use="optional"
	Reusable          bool                      `xml:"Reusable,attr,omitempty"`          //use="optional" default="false" //Indicates whether the form of payment can be reused or not. Currently applicable for Credit and Debit form of payment
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

type OtherGuaranteeInfo struct {
	Value string `xml:",innerxml"`
	Type  string `xml:"Type,attr"` //use="required" //1) IATA/ARC Number 2) Agency Address 2) Deposit Taken 3) Others
}

//Container to represent reservation name as appears in GDS booking
type ReservationName struct {
	//<xs:choice>
	BookingTravelerRef *BookingTravelerRef `xml:"BookingTravelerRef,omitempty"`
	NameOverride       *NameOverride       `xml:"NameOverride,omitempty"` //To be used if the reservation name is other than booking travelers in the PNR
	//</xs:choice>
}

//To be used if the name is different from booking travelers in the PNR
type NameOverride struct {
	First string `xml:"First,attr"`         //use="required" //First Name.
	Last  string `xml:"Last,attr"`          //use="required" //Last Name.
	Age   int    `xml:"Age,attr,omitempty"` //use="optional" //Age.
}

//Used as an override in a booking.
type BookingSource struct {
	Code string `xml:"Code,attr"` //use="required" //Alternate booking source code or number.
	Type string `xml:"Type,attr"` //use="required" //Type of booking source sent in the Code attribute. Possible values are "PseudoCityCode"," ArcNumber"," IataNumber", "CustomerId" and "BookingSourceOverrride". "BookingSourceOverrride" is only applicable in VehicleCreateReservationReq. 1P/1J.
}

//A textual remark container to hold Associated itinerary remarks with segment association
type AssociatedRemarkWithSegmentRef struct {
	AssociatedRemark
	SegmentRef com.TypeRef `xml:"SegmentRef,attr,omitempty"` //use="optional" //Reference to an Air/Passive Segment
}

type PolicyCodesList struct {
	PolicyCode    []com.TypePolicyCode `xml:"PolicyCode,omitempty"`    //minOccurs="0" maxOccurs="10" //A code that indicates why an item was determined to be 'out of policy'.
	MinPolicyCode []com.MinPolicyCode  `xml:"MinPolicyCode,omitempty"` //minOccurs="0" maxOccurs="10" //A code that indicates why the minimum fare or rate was determined to be 'out of policy'.
	MaxPolicyCode []com.MaxPolicyCode  `xml:"MaxPolicyCode,omitempty"` //minOccurs="0" maxOccurs="10" //A code that indicates why the maximum fare or rate was determined to be 'out of policy'.
}

//Shows the taxes and fees included in the base fare. (ACH only)
type IncludedInBase struct {
	Amount com.TypeMoney `xml:"Amount,attr,omitempty"` //use="optional" //this attribute shows the amount included in the base fare for the specific fee or tax
}

//Provides results of a exchange quote
type AirExchangeInfo struct {
	TotalPenaltyTaxInfo      *TotalPenaltyTaxInfo `xml:"TotalPenaltyTaxInfo,omitempty"`           //minOccurs="0" maxOccurs="1"
	PaidTax                  []*TypeTax           `xml:"PaidTax,omitempty"`                       //minOccurs="0" maxOccurs="unbounded"
	TicketFeeInfo            []*TicketFeeInfo     `xml:"TicketFeeInfo,omitempty"`                 //minOccurs="0" maxOccurs="unbounded" //Used for rapid reprice. Providers: 1G/1V/1P/1S/1A
	Reason                   []string             `xml:"Reason,omitempty"`                        //minOccurs="0" maxOccurs="unbounded" //Used for rapid reprice. The reason code or text is returned if the PricingTag is not equal to A, and explains why A was not returned. Providers: 1G/1V/1P/1S/1A
	FeeInfo                  []*TypeFeeInfo       `xml:"FeeInfo,omitempty"`                       //minOccurs="0" maxOccurs="unbounded"
	TaxInfo                  []*TypeTaxInfo       `xml:"TaxInfo,omitempty"`                       //minOccurs="0" maxOccurs="unbounded" //Itinerary level taxes
	ExchangeAmount           com.TypeMoney        `xml:"ExchangeAmount,attr"`                     //use="required"
	BaseFare                 com.TypeMoney        `xml:"BaseFare,attr,omitempty"`                 //use="optional"
	EquivalentBaseFare       com.TypeMoney        `xml:"EquivalentBaseFare,attr,omitempty"`       //use="optional"
	Taxes                    com.TypeMoney        `xml:"Taxes,attr,omitempty"`                    //use="optional"
	ChangeFee                com.TypeMoney        `xml:"ChangeFee,attr,omitempty"`                //use="optional"
	ForfeitAmount            com.TypeMoney        `xml:"ForfeitAmount,attr,omitempty"`            //use="optional"
	Refundable               bool                 `xml:"Refundable,attr,omitempty"`               //use="optional"
	Exchangeable             bool                 `xml:"Exchangeable,attr,omitempty"`             //use="optional"
	FirstClassUpgrade        bool                 `xml:"FirstClassUpgrade,attr,omitempty"`        //use="optional"
	TicketByDate             string               `xml:"TicketByDate,attr,omitempty"`             //use="optional"
	PricingTag               string               `xml:"PricingTag,attr,omitempty"`               //use="optional"
	EquivalentChangeFee      com.TypeMoney        `xml:"EquivalentChangeFee,attr,omitempty"`      //use="optional"
	EquivalentExchangeAmount com.TypeMoney        `xml:"EquivalentExchangeAmount,attr,omitempty"` //use="optional"
	AddCollection            com.TypeMoney        `xml:"AddCollection,attr,omitempty"`            //use="optional"
	ResidualValue            com.TypeMoney        `xml:"ResidualValue,attr,omitempty"`            //use="optional"
	TotalResidualValue       com.TypeMoney        `xml:"TotalResidualValue,attr,omitempty"`       //use="optional"
	OriginalFlightValue      com.TypeMoney        `xml:"OriginalFlightValue,attr,omitempty"`      //use="optional"
	FlownSegmentValue        com.TypeMoney        `xml:"FlownSegmentValue,attr,omitempty"`        //use="optional"
	BulkTicketAdvisory       bool                 `xml:"BulkTicketAdvisory,attr,omitempty"`       //use="optional"
	FarePull                 com.TypeFarePull     `xml:"FarePull,attr,omitempty"`                 //use="optional"
	Refund                   com.TypeMoney        `xml:"Refund,attr,omitempty"`                   //use="optional" //Total refund amount.
}

type TotalPenaltyTaxInfo struct {
	PenaltyTaxInfo  []*TypeTax    `xml:"PenaltyTaxInfo,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	TotalPenaltyTax com.TypeMoney `xml:"TotalPenaltyTax,attr,omitempty"` // use="optional"
}

//Used for rapid reprice. Providers: 1G/1V/1P/1S/1A
type TicketFeeInfo struct {
	Base  com.TypeMoney `xml:"Base,atrr,omitempty"`  //use="optional"
	Tax   com.TypeMoney `xml:"Tax,atrr,omitempty"`   //use="optional"
	Total com.TypeMoney `xml:"Total,atrr,omitempty"` //use="optional"
}

type TypeTax struct {
	Amount com.TypeMoney `xml:"Amount,atrr,omitempty"` //use="optional"
	Code   string        `xml:"Code,atrr,omitempty"`   //use="optional"
}

//A generic type of fee for those charges which are incurred by the passenger, but not necessarily shown on tickets
type TypeFeeInfo struct {
	TaxInfoRef         []*TaxInfoRef         `xml:"TaxInfoRef,omitempty"`              //minOccurs="0" maxOccurs="unbounded" //This reference elements will associate relevant taxes to this fee
	IncludedInBase     *IncludedInBase       `xml:"IncludedInBase,omitempty"`          //minOccurs="0" maxOccurs="1"
	BaseAmount         com.TypeMoney         `xml:"BaseAmount,attr,omitempty"`         //use="optional"
	Description        string                `xml:"Description,attr,omitempty"`        //use="optional"
	SubCode            string                `xml:"SubCode,attr,omitempty"`            //use="optional"
	Key                com.TypeRef           `xml:"Key,attr,omitempty"`                //use="required"
	Amount             com.TypeMoney         `xml:"Amount,attr,omitempty"`             //use="required"
	Code               string                `xml:"Code,attr,omitempty"`               //use="required"
	FeeToken           string                `xml:"FeeToken,attr,omitempty"`           //use="optional"
	PaymentRef         com.TypeRef           `xml:"PaymentRef,attr,omitempty"`         //use="optional"//The reference to the one of the air reservation payments if fee included in charge
	BookingTravelerRef com.TypeRef           `xml:"BookingTravelerRef,attr,omitempty"` //use="optional" //Reference to booking travele
	PassengerTypeCode  com.TypePTC           `xml:"PassengerTypeCode,attr,omitempty"`
	Text               com.StringLength1to64 `xml:"Text,attr,omitempty"` //use="optional" //Additional Information returned from Supplier.(ACH only)
	//<xs:attributeGroup ref="attrProviderSupplier"> //attrProviderSupplier
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr,omitempty"` //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode com.TypeSupplierCode `xml:"SupplierCode,attr,omitempty"` //use="optional" //Code of the supplier returning this TaxInfo.
	//</xs:attributeGroup>
	//<xs:attributeGroup name="attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"`      //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr,omitempty"` //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>

}

type TypeTaxInfo struct {
	TaxDetail              []*TaxDetail    `xml:"TaxDetail,omitempty"`                   //minOccurs="0" maxOccurs="unbounded"
	IncludedInBase         *IncludedInBase `xml:"IncludedInBase,omitempty"`              //minOccurs="0" maxOccurs="1"
	Key                    com.TypeRef     `xml:"Key,attr,omitempty"`                    //use="optional" //The tax key represents a valid key of tax
	Category               string          `xml:"Category,attr,omitempty"`               //use="required" //The tax category represents a valid IATA tax code.
	CarrierDefinedCategory string          `xml:"CarrierDefinedCategory,attr,omitempty"` //use="optional" //Optional category, where a carrier has used a non-standard IATA tax category. The tax category will be set to "DU"
	SegmentRef             com.TypeRef     `xml:"SegmentRef,attr,omitempty"`             //use="optional" //The segment to which that tax is relative (if applicable)
	FlightDetailsRef       com.TypeRef     `xml:"FlightDetailsRef,attr,omitempty"`       //use="optional" //The flight details that this tax is relative to (if applicable)
	CouponRef              com.TypeRef     `xml:"CouponRef,attr,omitempty"`              //use="optional" //The coupon to which that tax is relative (if applicable)
	//<xs:attributeGroup name="attrTaxDetail"> //Holds fare quote tax information.
	Amount             com.TypeMoney   `xml:"Amount,attr"`                       //use="required"
	OriginAirport      com.TypeAirport `xml:"OriginAirport,attr,omitempty"`      //use="optional"
	DestinationAirport com.TypeAirport `xml:"DestinationAirport,attr,omitempty"` //use="optional"
	CountryCode        string          `xml:"CountryCode,attr,omitempty"`        //use="optional"
	FareInfoRef        com.TypeRef     `xml:"FareInfoRef,attr,omitempty"`        //use="optional"
	//</xs:attributeGroup>
	TaxExempted  bool                   `xml:"TaxExempted,attr,omitempty"`  //use="optional" //This indicates whether the tax specified by tax category is exempted.
	ProviderCode com.TypeProviderCode   `xml:"ProviderCode,attr,omitempty"` //use="optional" //Code of the provider returning this TaxInfo.
	SupplierCode com.TypeSupplierCode   `xml:"SupplierCode,attr,omitempty"` //use="optional" //Code of the supplier returning this TaxInfo.
	Text         com.StringLength1to128 `xml:"Text,attr,omitempty"`         //use="optional" //Additional Information returned from Supplier.(ACH only)
}

//Exchange penalty information
type Penalty struct {
	CancelRefund        bool          `xml:"CancelRefund,attr,omitempty"`        //use="optional"
	NonRefundable       bool          `xml:"NonRefundable,attr,omitempty"`       //use="optional"
	NonExchangeable     bool          `xml:"NonExchangeable,attr,omitempty"`     //use="optional"
	CancelationPenalty  bool          `xml:"CancelationPenalty,attr,omitempty"`  //use="optional"
	ReissuePenalty      bool          `xml:"ReissuePenalty,attr,omitempty"`      //use="optional"
	NonReissuePenalty   bool          `xml:"NonReissuePenalty,attr,omitempty"`   //use="optional"
	TicketRefundPenalty bool          `xml:"TicketRefundPenalty,attr,omitempty"` //use="optional"
	ChargeApplicable    bool          `xml:"ChargeApplicable,attr,omitempty"`    //use="optional"
	ChargePortion       bool          `xml:"ChargePortion,attr,omitempty"`       //use="optional"
	PenaltyAmount       com.TypeMoney `xml:"PenaltyAmount,attr,omitempty"`       //use="optional"
}
