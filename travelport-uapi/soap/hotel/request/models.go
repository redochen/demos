package request

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
)

//The hotel property
type HotelProperty struct {
	PropertyAddress       *UnstructuredAddress       `xml:"hot:PropertyAddress,omitempty"`        //minOccurs="0"
	PhoneNumber           []*comrq.PhoneNumber       `xml:"com:PhoneNumber,omitempty"`            //minOccurs="0" maxOccurs="unbounded"
	CoordinateLocation    *comrq.CoordinateLocation  `xml:"com:CoordinateLocation,omitempty"`     //minOccurs="0"
	Distance              *comrq.Distance            `xml:"com:Distance,omitempty"`               //minOccurs="0"
	HotelRating           []*HotelRating             `xml:"hot:HotelRating,omitempty"`            //minOccurs="0" maxOccurs="unbounded"
	Amenities             *Amenities                 `xml:"hot:Amenities,omitempty"`              //minOccurs="0"
	MarketingMessage      *MarketingMessage          `xml:"hot:MarketingMessage,omitempty"`       //MarketingMessage
	HotelChain            com.TypeHotelChainCode     `xml:"HotelChain,attr"`                      //use="required"
	HotelCode             com.TypeHotelCode          `xml:"HotelCode,attr"`                       //use="required"
	HotelLocation         hot.TypeHotelLocationCode  `xml:"HotelLocation,attr,omitempty"`         //use="optional" //The location code for this entity. IATA code or TRM code.
	Name                  string                     `xml:"Name,attr,omitempty"`                  //use="optional"
	VendorLocationKey     string                     `xml:"VendorLocationKey,attr,omitempty"`     //use="optional" //The VendorLocationKey for this HotelProperty.
	HotelTransportation   com.TypeOTACode            `xml:"HotelTransportation,attr,omitempty"`   //use="optional" //OTA Transporation code. Transportation available to hotel.
	ReserveRequirement    com.TypeReserveRequirement `xml:"ReserveRequirement,attr,omitempty"`    //use="optional"
	ParticipationLevel    com.StringLength1          `xml:"ParticipationLevel,attr,omitempty"`    //use="optional" //2=Best Available Rate 1G, 1V,  4=Lowest Possible Rate 1G, 1V, 1P, 1J
	Availability          string                     `xml:"Availability,attr,omitempty"`          //use="optional"
	Key                   com.TypeRef                `xml:"Key,attr,omitempty"`                   //use="optional"
	PreferredOption       bool                       `xml:"PreferredOption,attr,omitempty"`       //use="optional" //This attribute is used to indicate if the vendors responsible for the fare or rate being returned have been determined to be ‘preferred’ based on the associated policy settings.
	MoreRates             bool                       `xml:"MoreRates,attr,omitempty"`             //When true, more rates are available for this hotel property.Applicable only for HotelDetails and HotelSuperShopper. Supported Providers: 1G, 1V.
	MoreRatesToken        string                     `xml:"MoreRatesToken,attr,omitempty"`        //use="optional" //HS3 Token to identify the Rates for a property. Supported Providers 1G,1V.
	NetTransCommissionInd string                     `xml:"NetTransCommissionInd,attr,omitempty"` //use="optional" //This attribute indicates whether hotel property is tracking through net trans commission indicator.
	NumOfRatePlans        uint                       `xml:"NumOfRatePlans,attr,omitempty"`        //use="optional" //The specific number of RatePlanTypes for each property responded on the message, integer 1 - 999. Supported provider: HotelSuperShopper message only.
}

//Location information for the hotel.
type HotelSearchLocation struct {
	HotelLocation      *HotelLocation               `xml:"hot:HotelLocation,omitempty"`      //minOccurs="0" //Date and Location information for the Hotel.
	ProviderLocation   *ProviderLocation            `xml:"hot:ProviderLocation,omitempty"`   //minOccurs="0" //Provider specific Hotel location. Applicable for TRM only.
	VendorLocation     *comrq.VendorLocation        `xml:"com:VendorLocation,omitempty"`     //minOccurs="0" maxOccurs="99"
	HotelAddress       *comrq.TypeStructuredAddress `xml:"com:HotelAddress,omitempty"`       //minOccurs="0" //Search by address or postal code. Applicable for 1G, 1V, 1P, 1J
	ReferencePoint     *HotelReferencePoint         `xml:"hot:ReferencePoint,omitempty"`     //minOccurs="0" //Search for hotels near a reference point. HotelLocation/Location is mandatory for aggregated scenario if ReferencePoint is used. Applicable for 1G,1V,1P,1J. Country/State are only applicable for 1P/1J
	CoordinateLocation *comrq.CoordinateLocation    `xml:"com:CoordinateLocation,omitempty"` //minOccurs="0" //Search using latitude and longitude. Applicable for 1G, 1V only. Not applicable for HotelSuperShopper
	Distance           *comrq.Distance              `xml:"com:Distance,omitempty"`           //minOccurs="0"
}

//Date and Location information for the Hotel. Location can be optional if a Reference Point is provided.
type HotelLocation struct {
	Location     com.TypeIATACode `xml:"Location,attr,omitempty"`     //use="optional" //IATA city/airport code
	LocationType string           `xml:"LocationType,attr,omitempty"` //default="Airport"
}

//Provider specific Hotel location. Applicable for TRM only.
type ProviderLocation struct {
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr"` //use="required" //Code for the provider.
	Location     string               `xml:"Location,attr"`     //use="required" //6 character location code.
}

//Controls and switches for the Hotel Search request
type HotelSearchModifiers struct {
	PermittedChains           *HotelChains                 `xml:"hot:PermittedChains,omitempty"`            //minOccurs="0"
	ProhibitedChains          *HotelChains                 `xml:"hot:ProhibitedChains,omitempty"`           //minOccurs="0"
	PermittedProviders        *comrq.PermittedProviders    `xml:"com:PermittedProviders,omitempty"`         //minOccurs="0"
	PermittedAggregators      *PermittedAggregators        `xml:"hot:PermittedAggregators,omitempty"`       //minOccurs="0"
	LoyaltyCard               []*comrq.LoyaltyCard         `xml:"com:LoyaltyCard,omitempty"`                //minOccurs="0" maxOccurs="4"
	HotelName                 string                       `xml:"hot:HotelName,omitempty"`                  //minOccurs="0" //There can be at most one Hotel Name to be requested
	CorporateDiscountID       []*comrq.CorporateDiscountID `xml:"com:CorporateDiscountID,omitempty"`        //minOccurs="0" maxOccurs="unbounded" //Search with corporate discount IDs or negotiated rate codes. 1G/1V allows a max of 4. 1P/1J allows a max of 1 corporate discount ID and up to 30 negotiated rate codes. Support for this function is hotel supplier dependent.
	RateCategory              []com.TypeOTACode            `xml:"com:RateCategory,omitempty"`               //minOccurs="0" maxOccurs="8" //Search for specific rate categories
	HotelRating               []*HotelRating               `xml:"hot:HotelRating,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	SearchPriority            *SearchPriority              `xml:"hot:SearchPriority,omitempty"`             //minOccurs="0"
	HotelBedding              []*HotelBedding              `xml:"hot:HotelBedding,omitempty"`               //minOccurs="0" maxOccurs="4"
	Amenities                 *Amenities                   `xml:"hot:Amenities,omitempty"`                  //minOccurs="0"
	HotelPaymentType          string                       `xml:"hot:HotelPaymentType,omitempty"`           //minOccurs="0" //Payment type. “PrePay” rates require advance payment to complete the booking. “PostPay” rates allow payment after booking, typically after the hotel stay is completed. By default, all payment types are returned. Supported Provider: TRM.
	NumberOfChildren          *NumberOfChildren            `xml:"hot:NumberOfChildren,omitempty"`           //minOccurs="0"
	HotelTransportation       *TransportationType          `xml:"hot:HotelTransportation,omitempty"`        //minOccurs="0" //OTA Transportation code. Search for specific transportation. Supported providers: 1G/1V.  Only CourtesyBus '7' supported by 1P/1J.
	BookingGuestInformation   *BookingGuestInformation     `xml:"hot:BookingGuestInformation,omitempty"`    //minOccurs="0"
	NumberOfAdults            int                          `xml:"NumberOfAdults,attr,omitempty"`            //use="optional" //The total number of adult guests per booking. Defaults to ‘1’. Supported Providers: 1G, 1V, 1P, 1J. Also for aggregated request (GDS provider + TRM).
	NumberOfRooms             int                          `xml:"NumberOfRooms,attr,omitempty"`             //use="optional" //The number of rooms per booking. Defaults to ‘1’. Supported Providers 1G, 1V, 1P, 1J. Also for aggregated request (GDS provider + TRM) for single room only.
	MaxProperties             int                          `xml:"MaxProperties,attr,omitempty"`             //use="optional" //The maximum number of hotel properties to return in one message. Supported by TRM only. Valid values 1 to 9999. Default is 99.
	IsRelaxed                 bool                         `xml:"IsRelaxed,attr,omitempty"`                 //use="optional" //Default is true. If false, only the results matching all the criteria returned.
	PreferredCurrency         com.TypeCurrency             `xml:"PreferredCurrency,attr,omitempty"`         //use="optional" //Requested currency for target rate.
	AvailableHotelsOnly       bool                         `xml:"AvailableHotelsOnly,attr,omitempty"`       //use="optional" //Set to true to request only available hotels. Default is false and all results from the provider are returned.
	MaxWait                   uint                         `xml:"MaxWait,attr,omitempty"`                   //Maximum wait time in milliseconds for hotel search results. Supported provider: TRMAggregateResults
	AggregateResults          bool                         `xml:"AggregateResults,attr,omitempty"`          //use="optional" default="false" //Indicator to identify TRM and GDS property match required or not.
	ReturnPropertyDescription bool                         `xml:"ReturnPropertyDescription,attr,omitempty"` //use="optional" default="false" //Request hotel property description. Valid Values are "true" or "false". Default "false". Supported Providers: TRM
	NumOfRatePlans            uint                         `xml:"NumOfRatePlans,attr,omitempty"`            //use="optional" //The specific number of RatePlanTypes for each property responded on the message, integer 1 - 999. Supported provider: HotelSuperShopper message only.
	ReturnAmenities           bool                         `xml:"ReturnAmenities,attr,omitempty"`           //use="optional" default="false" //If hotel amenities are desired set as 'true', else default 'false' for no amenity support.
}

//Controls and switches for the Hotel Details request
type HotelDetailsModifiers struct {
	PermittedProviders      *comrq.PermittedProviders    `xml:"com:PermittedProviders,omitempty"`  //minOccurs="0"
	LoyaltyCard             []*comrq.LoyaltyCard         `xml:"com:LoyaltyCard,omitempty"`         //minOccurs="0" maxOccurs="4"
	CorporateDiscountID     []*comrq.CorporateDiscountID `xml:"com:CorporateDiscountID,omitempty"` //minOccurs="0" maxOccurs="9999" //Search with corporate discount IDs or negotiated rate codes. 1G/1V allows a max of 4. 1P/1J allows a max of 1 corporate discount ID and up to 30 negotiated rate codes. Support for this function is hotel supplier dependent.
	HotelStay               *HotelStay                   `xml:"hot:HotelStay"`
	NumberOfChildren        *NumberOfChildren            `xml:"hot:NumberOfChildren,omitempty"`        //minOccurs="0"
	HotelBedding            []*HotelBedding              `xml:"hot:HotelBedding,omitempty"`            //minOccurs="0" maxOccurs="4"
	RateCategory            []com.TypeOTACode            `xml:"hot:RateCategory,omitempty"`            //minOccurs="0" maxOccurs="8" //Specify Rate Category
	PermittedAggregators    *PermittedAggregators        `xml:"hot:PermittedAggregators,omitempty"`    //minOccurs="0" //TRM Only
	BookingGuestInformation *BookingGuestInformation     `xml:"hot:BookingGuestInformation,omitempty"` //minOccurs="0" //Information about requested rooms and allocation of guests per room. Provider: TRM.
	NumberOfAdults          int                          `xml:"NumberOfAdults,attr,omitempty"`         //use="optional" //The total number of adult guests per booking. Defaults to ‘1’. GDS Providers: 1G, 1V, 1P, 1J. Also for aggregated request (GDS provider + TRM).
	RateRuleDetail          hot.TypeRateRuleDetail       `xml:"RateRuleDetail,attr"`                   //default="None"
	NumberOfRooms           int                          `xml:"NumberOfRooms,attr,omitempty"`          //use="optional" default="1" //The number of rooms per booking. Defaults to ‘1’. GDS Providers 1G, 1V, 1P, 1J. Also for aggregated request (GDS provider + TRM) for single room only.
	Key                     com.TypeRef                  `xml:"Key,attr,omitempty"`                    //use="optional
	PreferredCurrency       com.TypeCurrency             `xml:"PreferredCurrency,attr,omitempty"`      //use="optional //Alternate currency
	TotalOccupants          int                          `xml:"TotalOccupants,attr,omitempty"`         //use="optional" //Number of guests for the room. Supported Providers: 1P/1J
	ProcessAllNegoRatesInd  bool                         `xml:"ProcessAllNegoRatesInd,attr,omitempty"` //use="optional" default="false" //When false, we will process the request with all the provided negotiated rates in a single request. The request will fail when the number of negotiated rates have exceeded for that hotel chain. When true, this allows to process a request for all provided negotiated rates that may exceed the hotel chain limit. Supported for 1P only.
	MaxWait                 uint                         `xml:"MaxWait,attr"`                          //Maximum wait time in milliseconds for hotel detail results. Supported provider: TRM
}

type HotelChains struct {
	HotelChain []*HotelChain `xml:"hot:HotelChain"` //maxOccurs="unbounded"
}

//The hotel chain code
type HotelChain struct {
	Code com.TypeHotelChainCode `xml:"Code,attr"` //use="required"
}

//Hotel specific rate categories (e.g. Standard, Military)
type HotelRateCategory struct {
	Value string `xml:",innerxml"`
}

//OTA Transportation code. Search for specific transportation. Supported providers: 1G/1V.  Only CourtesyBus '7' supported by 1P/1J.
type TransportationType struct {
	Type com.TypeOTACode `xml:"Type,attr"` //use="required" //Transportation type code
}

//Hotel rating information
type HotelRating struct {
	//<xs:choice>
	Rating      []hot.TypeSimpleHotelRating `xml:"hot:Rating"`      //maxOccurs="unbounded" //Hotel rating value
	RatingRange *RatingRange                `xml:"hot:RatingRange"` //Search for a range of ratings
	//</xs:choice>
	RatingProvider string `xml:"RatingProvider,attr"` //use="required" //Rating providers, ie AAA, NTM
}

//Search for a range of rating
type RatingRange struct {
	MinimumRating hot.TypeSimpleHotelRating `xml:"MinimumRating,attr,omitempty"` //use="optional"
	MaximumRating hot.TypeSimpleHotelRating `xml:"MaximumRating,attr,omitempty"` //use="optional"
}

//Override the search order for hotel availability request
type SearchPriority struct {
	Criteria []*Criteria `xml:"hot:Criteria"` //maxOccurs="8"
}

type Criteria struct {
	Order int    `xml:,"Order,attr"` //use="required" //Criteria order for hotel search, Highest Priority=1 Lowest Priority=7
	Type  string `xml:"Type,attr"`   //use="required" //Search type
}

//Specify desired bedding
type HotelBedding struct {
	Type         string        `xml:"Type,attr"`                   //use="required" //Queen, King, double, etc
	NumberOfBeds int           `xml:"NumberOfBeds,attr,omitempty"` //use="optional" //Number of beds of desired Type in room. Use '0' to delete the hotel Optional Beds ( Only RA RC CR )
	Amount       com.TypeMoney `xml:"Amount,attr,omitempty"`       //use="optional" //Fee for bed type. Providers: 1g/1v/1p/1j
	Content      string        `xml:"Content,attr,omitempty"`      //use="optional" //Additional information Providers: 1p/1j
}

//Amenity information
type Amenities struct {
	Amenity []*Amenity `xml:"hot:Amenity,omitempty"` //minOccurs="0" maxOccurs="8"
}

type Amenity struct {
	Code        hot.TypeAmenity `xml:"Code,attr"`                  //use="required"
	AmenityType string          `xml:"AmenityType,attr,omitempty"` //use="optional" //Amenity type code. “HA” (Hotel Property Amenity) or “RA” (Room Amenity). Defaults to “HA” if no value is sent.
}

//Number of Adults
type NumberOfAdults struct {
	Value       string        `xml:",innerxml"`
	ExtraAdults int           `xml:"ExtraAdults,attr,omitempty"` //use="optional" //The number of extra adults in the room ,use '0' to delete the extra adults
	Amount      com.TypeMoney `xml:"Amount,attr,omitempty"`      //use="optional" //Fee for extra adults.  Providers: 1g/1v/1p/1j
	Content     string        `xml:"Content,attr,omitempty"`     //use="optional" //Additional information.  Providers 1p/1j
}

//Number of Children
type NumberOfChildren struct {
	Age    []int         `xml:"hot:Age,omitempty"`    //minOccurs="0" maxOccurs="99" //The Ages of the Children. . The defined age of a Child traveler may vary by supplier, but is typically 1 to 17 years. Supported Providers 1G/1V.
	Count  int           `xml:"Count,attr"`           //use="required" //The total number of children in the booking. Supported Providers 1P/1J.
	Amount com.TypeMoney `xml:"Amountattr,omitempty"` //use="optional" //Fee per child. Providers: 1g/1v
}

//Arrival and Departure dates
type HotelStay struct {
	CheckinDate  hot.TypeDate `xml:"hot:CheckinDate"`
	CheckoutDate hot.TypeDate `xml:"hot:CheckoutDate"`
	Key          com.TypeRef  `xml:"Key,attr,omitempty"` //use="optional"
}

type HotelReferencePoint struct {
	Value   com.TypeReferencePoint `xml:",innerxml"`
	Country com.TypeCountry        `xml:"Country,attr,omitempty"` //Country code.
	State   com.TypeState          `xml:"State,attr,omitempty"`   //State or Province Code.
}

//A simple unstructured address (e.g. 123 South State Avenue, Chicago, IL 60612)
type UnstructuredAddress struct {
	Address []string `xml:"hot:Address"` //maxOccurs="6"
}

//Returns hotel rate details during the stay if rates are available for requested property.
type HotelRateDetail struct {
	PolicyCodesList      *comrq.PolicyCodesList       `xml:"hot:PolicyCodesList,omitempty"`      //minOccurs="0" //A list of codes that indicate why an item was determined to be ‘out of policy’.
	RoomRateDescription  []*HotelRateDescription      `xml:"hot:RoomRateDescription,omitempty"`  //minOccurs="0" maxOccurs="9999"
	HotelRateByDate      []*HotelRateByDate           `xml:"hot:HotelRateByDate,omitempty"`      //minOccurs="0" maxOccurs="9999"
	CorporateDiscountID  []*comrq.CorporateDiscountID `xml:"com:CorporateDiscountID,omitempty"`  //minOccurs="0" maxOccurs="9999" //Corporate Discount IDs and Negotiate rate codes associated with this rate
	AcceptedPayment      []*AcceptedPayment           `xml:"hot:AcceptedPayment,omitempty"`      //minOccurs="0" maxOccurs="99" //Form of payment accepted by the hotel supplier (chain or property). For credit cards, the two-character code for the credit card type is used.
	Commission           *Commission                  `xml:"hot:Commission,omitempty"`           //minOccurs="0" maxOccurs="1" //Commission associated with the Rate Plan, as a percentage or flat amount.
	RateMatchIndicator   []*RateMatchIndicator        `xml:"hot:RateMatchIndicator,omitempty"`   //minOccurs="0" maxOccurs="9999" //Returns "Match" Indicators for certain request parameters for Hotel Rate returned in response.
	TaxDetails           *TaxDetails                  `xml:"hot:TaxDetails,omitempty"`           //minOccurs="0"
	CancelInfo           *CancelInfo                  `xml:"hot:CancelInfo,omitempty"`           //minOccurs="0"
	GuaranteeInfo        *GuaranteeInfo               `xml:"hot:GuaranteeInfo,omitempty"`        //minOccurs="0" //Guarantee, deposit, and prepayment information
	SupplementalRateInfo string                       `xml:"hot:SupplementalRateInfo,omitempty"` //minOccurs="0" //Supplemental rate information provided by the aggregator. Supported Providers TRM.
	RoomCapacity         *RoomCapacity                `xml:"hot:RoomCapacity,omitempty"`         //minOccurs="0" //The maximum number of guests for a room or for each room in a package. Provider: TRM.
	RatePlanType         com.TypeRatePlanType         `xml:"RatePlanType,attr"`                  //use="required"
	//<xs:attributeGroup name="attrHotelRate"> //Attributes used to describe Hotel Rates
	Base      com.TypeMoney `xml:"Base,attr,omitempty"`      //use="optional" //This attribute is used to describe the Hotel Supplier Base Rate
	Tax       com.TypeMoney `xml:"Tax,attr,omitempty"`       //use="optional" //This attribute used to describe Tax associated with the room
	Total     com.TypeMoney `xml:"Total,attr,omitempty"`     //use="optional" //This attribute used to describe Hotel Supplier Total Rate
	Surcharge com.TypeMoney `xml:"Surcharge,attr,omitempty"` //use="optional" //This attribute used to describe Surcharge associated with the room
	//</xs:attributeGroup>
	ApproximateBase           com.TypeMoney              `xml:"ApproximateBase,attr,omitempty"`           //use="optional" //Hotel base rate expressed in another currency
	ApproximateTax            com.TypeMoney              `xml:"ApproximateTax,attr,omitempty"`            //use="optional" //Taxes expressed in another currency
	ApproximateTotal          com.TypeMoney              `xml:"ApproximateTotal,attr,omitempty"`          //use="optional" //Hotel total rate expressed in another currency. For TRM, ApproximateTotal is without fees
	ApproximateSurcharge      com.TypeMoney              `xml:"ApproximateSurcharge,attr,omitempty"`      //use="optional" //Surcharge associated with the room expressed in another currency. Applicable for TRM only.
	RateGuaranteed            bool                       `xml:"RateGuaranteed,attr,omitempty"`            //use="optional"
	ApproximateRateGuaranteed bool                       `xml:"ApproximateRateGuaranteed,attr,omitempty"` //use="optional" //If true, approximate rate is guaranteed by vendor. Supported Providers: 1G,1V
	RateCategory              com.TypeOTACode            `xml:"RateCategory,attr,omitempty"`              //use="optional" //An enumerated type that allows the query to specify a rate category type, and provides major categories for comparison across brands. Refer to OpenTravel Code List Rate Plan Type (RPT). Encode/decode data in Util ReferenceDataRetrieveReq TypeCode=“HotelRateCategory".
	Key                       com.TypeRef                `xml:"Key,attr,omitempty"`                       //use="optional"
	RateSupplier              com.TypeThirdPartySupplier `xml:"RateSupplier,attr,omitempty"`              //use="optional" //Indicates the source of the rate. Provider: TRM.
	BookableQuantity          int                        `xml:"BookableQuantity,attr,omitempty"`          //use="optional" //The number of rooms which can be booked on the rate returned in HotelRateDetails.  When the aggregator responds ‘IsPackage’= true (pricing for all rooms together), the BookableQuantity value will be ‘1’. Supported Providers TRM.
	RateOfferId               hot.TypeRateOfferId        `xml:"RateOfferId,attr,omitempty"`               //use="optional" //Offer Identifier. Maybe required for hotels provided by aggregators. Supported Provider TRM.
	InPolicy                  bool                       `xml:"InPolicy,attr,omitempty"`                  //use="optional" //This attribute will be used to indicate if a fare or rate has been determined to be ‘in policy’ based on the associated policy settings.
	RateChangeIndicator       com.TypeTrinary            `xml:"RateChangeIndicator,attr,omitempty"`       //use="optional" //Determines if the rate changes during the length of stay. Enumerated values are true, false, and unknown.
	ExtraFeesIncluded         com.TypeTrinary            `xml:"ExtraFeesIncluded,attr,omitempty"`         //use="optional" //When true, total amounts includes additional fees or charges." Enumerated values are true, false, and unknown
}

//The maximum number of guests for a room or for each room in a package. Provider: TRM.
type RoomCapacity struct {
	Capacity  []uint `xml:"hot:Capacity,omitempty"`   //minOccurs="0" maxOccurs="99" //The maximum number of guests per room. Provider: TRM.
	IsPackage bool   `xml:"IsPackage,attr,omitempty"` //use="optional" //If true, the rooms are offered as a package by the aggregator. Supported Providers TRM.
}

type HotelRateDescription struct {
	Text []string `xml:"Text"`                //maxOccurs="unbounded"
	Name string   `xml:"Name,attr,omitempty"` //use="optional" //Optional context name of the text block being returned i.e. Room details
}

//The daily rate details
type HotelRateByDate struct {
	EffectiveDate string `xml:"EffectiveDate,attr,omitempty"` //use="optional"
	ExpireDate    string `xml:"ExpireDate,attr,omitempty"`    //use="optional"
	//<xs:attributeGroup name="attrHotelRate"> //Attributes used to describe Hotel Rates
	Base      com.TypeMoney `xml:"Base,attr,omitempty"`      //use="optional" //This attribute is used to describe the Hotel Supplier Base Rate
	Tax       com.TypeMoney `xml:"Tax,attr,omitempty"`       //use="optional" //This attribute used to describe Tax associated with the room
	Total     com.TypeMoney `xml:"Total,attr,omitempty"`     //use="optional" //This attribute used to describe Hotel Supplier Total Rate
	Surcharge com.TypeMoney `xml:"Surcharge,attr,omitempty"` //use="optional" //This attribute used to describe Surcharge associated with the room
	//</xs:attributeGroup>
	ApproximateBase  com.TypeMoney `xml:"ApproximateBase,attr,omitempty"`  //use="optional" //Hotel base rate expressed in another currency
	ApproximateTotal com.TypeMoney `xml:"ApproximateTotal,attr,omitempty"` //use="optional" //Hotel total rate expressed in another currency. Supported Providers: 1P,1J
	Contents         string        `xml:"Contents,attr,omitempty"`         //use="optional" //Contents will be representing all unformatted data returned by HOST, those are not uAPI supported. Support provider 1P and 1J.
}

//"Match" Indicators for certain request parameters, e.g. Child Count, Extra Adults etc.
type RateMatchIndicator struct {
	Type   string `xml:"Type,attr"`
	Status string `xml:"Status,attr"`
	Value  string `xml:"Value,attr,omitempty"`
}

type TaxDetails struct {
	Tax []*Tax `xml:"hot:Tax"` //minOccurs="1" maxOccurs="unbounded"
}

type Tax struct {
	//<xs:choice>
	Amount     com.TypeMoney `xml:"com:Amount,omitempty"`     //minOccurs="1"
	Percentage float32       `xml:"hot:Percentage,omitempty"` //minOccurs="1"
	//</xs:choice>
	Code           com.TypeOTACode `xml:"Code,attr"`                     //use="required" //Code identifying fee (e.g. agency fee, bed tax etc.). Refer to OPEN Travel Code List for Fee Tax Type. Possible values are OTA Code against FTT.
	EffectiveDate  string          `xml:"EffectiveDate,attr,omitempty"`  //type="xs:date" use="optional"
	ExpirationDate string          `xml:"ExpirationDate,attr,omitempty"` //type="xs:date" use="optional"
	Term           string          `xml:"Term,attr,omitempty"`           //use="optional" //Indicates how the tax is applied. Values can be PerPerson, PerNight and PerStay
	CollectionFreq string          `xml:"CollectionFreq,attr,omitempty"` //use="optional" //Indicates how often the tax is collected. Values can be Once or Daily
}

//Returns cancellation information for certain hotel returned in response. This information is available through GDS transactions
type CancelInfo struct {
	CancellationPolicy            string          `xml:"hot:CancellationPolicy,omitempty"`             //minOccurs="0" //Return cancellation policy text by the aggregator. Provider: TRM.
	Text                          []string        `xml:"hot:Text,omitempty"`                           //minOccurs="0" maxOccurs="99" //The informational text provided by the supplier to cancel the booking, if @Method="INFO". For all other values of @Method, Text is not returned. Provider: TRM.
	NonRefundableStayIndicator    com.TypeTrinary `xml:"NonRefundableStayIndicator,attr,omitempty"`    //use="optional" //True if Deposit or Payment is non-refundable
	CancelDeadline                string          `xml:"CancelDeadline,attr,omitempty"`                //type="xs:dateTime" use="optional" //Last date/time the reservation can be canceled without penalty.
	TaxInclusive                  bool            `xml:"TaxInclusive,attr,omitempty"`                  //use="optional" //Indicates whether or not the Penalty amount includes taxes.
	FeeInclusive                  bool            `xml:"FeeInclusive,attr,omitempty"`                  //use="optional" //Indicates whether or not the Penalty amount includes fees.
	CancelPenaltyAmount           com.TypeMoney   `xml:"CancelPenaltyAmount,attr,omitempty"`           //use="optional" //This will contain the cancellation penalty amount.
	NumberOfNights                uint            `xml:"NumberOfNights,attr,omitempty"`                //use="optional" //This will contain the number of nights that will be assessed as the cancelation penalty.
	CancelPenaltyPercent          float32         `xml:"CancelPenaltyPercent,attr,omitempty"`          //use="optional" //This will contain the cancellation penalty expressed as a percentage.
	CancelPenaltyPercentAppliesTo string          `xml:"CancelPenaltyPercentAppliesTo,attr,omitempty"` //use="optional" //This will contain the cost qualifier that explains what the percentage is applied to in order to calculate the cancel penalty.
	Method                        string          `xml:"Method,attr,omitempty"`                        //use="optional" //Cancellation method, either "API", "URL", "INFO", or "NONE". Supported Providers TRM.
	Supported                     bool            `xml:"Supported,attr,omitempty"`                     //use="optional" //If true, the booking can be canceled. If false, the booking cannot be canceled. Provider: TRM.
	URL                           string          `xml:"URL,attr,omitempty"`                           //use="optional"  //The URL provided by the supplier to cancel the booking, if @Method="URL". For all other values of @Method, @URL is not returned. Provider: TRM.
	//<xs:attributeGroup name="DeadlineGroup"> //The absolute deadline or amount of offset time before a deadline for a payment of cancel goes into effect.
	OffsetTimeUnit       com.StringLength1to16 `xml:"OffsetTimeUnit,attr,omitempty"`       //use="optional" //The units of time, e.g: days, hours, etc that apply to the deadline. Enumerated values are “Year”, “Month”, “Day”, and “Hour”.
	OffsetUnitMultiplier com.Numeric0to999     `xml:"OffsetUnitMultiplier,attr,omitempty"` //use="optional" //The number of units of DeadlineTimeUnit.
	OffsetDropTime       com.StringLength1to20 `xml:"OffsetDropTime,attr,omitempty"`       //use="optional" //An enumerated type indicating when the deadline drop time goes into effect. Enumerated values are “AfterBooking” and “BeforeArrival”.
	//</xs:attributeGroup>
}

type GuaranteeInfo struct {
	//<xs:choice minOccurs="0">
	DepositAmount  *DepositAmount `xml:"hot:DepositAmount,omitempty"`  //minOccurs="0" //Amount required for deposit/prepayment
	DepositNights  int            `xml:"hot:DepositNights,omitempty"`  //Number of Nights required for deposit/prepayment
	DepositPercent int            `xml:"hot:DepositPercent,omitempty"` //Percentage of stay required for deposit/prepayment
	//</xs:choice>
	GuaranteePaymentType []*GuaranteePaymentType `xml:"hot:GuaranteePaymentType,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Accepted payment types
	AbsoluteDeadline     string                  `xml:"AbsoluteDeadline,attr,omitempty"`    //type="xs:dateTime" use="optional" //Latest date/time when deposit/payment/guarantee is required.Not Supported by TRM.
	CredentialsRequired  bool                    `xml:"CredentialsRequired,attr,omitempty"` //use="optional" //Identification required at booking/checkin. Not supported by 1P,1J and TRM.
	HoldTime             string                  `xml:"HoldTime,attr,omitempty"`            //use="optional" //Expiration time for room reservation held without deposit/guarantee. Not supported by TRM.
	GuaranteeType        string                  `xml:"GuaranteeType,attr,omitempty"`       //use="optional" //Deposit, Guarantee, or Prepayment required to hold/book the room. Applicable only for HotelSupershopper, Hotel Details and Hotel rules
	//<xs:attributeGroup name="DeadlineGroup"> //The absolute deadline or amount of offset time before a deadline for a payment of cancel goes into effect.
	OffsetTimeUnit       com.StringLength1to16 `xml:"OffsetTimeUnit,attr,omitempty"`       //use="optional" //The units of time, e.g: days, hours, etc that apply to the deadline. Enumerated values are “Year”, “Month”, “Day”, and “Hour”.
	OffsetUnitMultiplier com.Numeric0to999     `xml:"OffsetUnitMultiplier,attr,omitempty"` //use="optional" //The number of units of DeadlineTimeUnit.
	OffsetDropTime       com.StringLength1to20 `xml:"OffsetDropTime,attr,omitempty"`       //use="optional" //An enumerated type indicating when the deadline drop time goes into effect. Enumerated values are “AfterBooking” and “BeforeArrival”.
	//</xs:attributeGroup>
}

//Amount required for deposit/prepayment
type DepositAmount struct {
	Amount            com.TypeMoney `xml:"Amount,attr,omitempty"`            //use="optional" //Supplier deposit amount required for deposit/prepayment.Supported by all Providers when supported by supplier
	ApproximateAmount com.TypeMoney `xml:"ApproximateAmount,attr,omitempty"` //use="optional" //Approximate deposit amount required for deposit/prepayment.Supports TRM provider only.
}

//Accepted payment types. Applicable only for HotelSupershopper, Hotel Details and Hotel rules.
type GuaranteePaymentType struct {
	Value       string `xml:",innerxml"`
	Type        string `xml:"Type,attr"`                  //use="required" //Accepted payment types: CreditCard, AgencyIATA/ARC, FrequentGuest, SpecialIndustry, CDNumber, HomeAddress, CompanyAddress, Override, Other, or None
	Description string `xml:"Description,attr,omitempty"` //use="optional"
}

type RoomView struct {
	Code com.TypeOTACode `xml:"Code,attr,omitempty"` //use="optional" //OTA code represents different hotel room views.
}

type PromotionCode struct {
	Value string      `xml:",innerxml"`
	Key   com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
}

//The information like number of rooms ,number of adults,children to be provided while booking the  hotel
type GuestInformation struct {
	NumberOfAdults   *NumberOfAdults   `xml:"hot:NumberOfAdults,omitempty"`   //minOccurs="0"
	NumberOfChildren *NumberOfChildren `xml:"hot:NumberOfChildren,omitempty"` //minOccurs="0"
	ExtraChild       *ExtraChild       `xml:"hot:ExtraChild,omitempty"`       //minOccurs="0" //Providers: 1p/1j
	NumberOfRooms    int               `xml:"NumberOfRooms,attr,omitempty"`   //use="optional"
}

type ExtraChild struct {
	Count   int    `xml:"Count,attr"`   //The number of extra children in the room
	Content string `xml:"Content,attr"` //Additional information
}

type AssociatedRemark struct {
	comrq.AssociatedRemark
}

//Information about guest to book. Supported Providers TRM.
type TypeGuestRoomInformation struct {
	Adults             int                         `xml:"hot:Adults"`                       //The number of adult guests per room. Maximum number of adults may vary by supplier or aggregator.
	BookingTravelerRef []*comrq.BookingTravelerRef `xml:"hot:BookingTravelerRef,omitempty"` //minOccurs="0" maxOccurs="9" //Reference for the Booking Traveler. Used for Hotel Booking only. The value is arbitrary.
	Child              []*Child                    `xml:"hot:Child,omitempty"`              //minOccurs="0" maxOccurs="6" //Information about a child guest.
}

//Information about requested rooms and guests allocation. Supported Providers TRM.
type BookingGuestInformation struct {
	Room []*TypeGuestRoomInformation `xml:"hot:Room"` //maxOccurs="9" //Individual room. Multiple occurrences if there are multiple rooms in the request. Maximum number of rooms may vary by supplier or aggregator.
}

//Information about a child guest.
type Child struct {
	TypeGuestChildInformation
	BookingTravelerRef *comrq.BookingTravelerRef `xml:"hot:BookingTravelerRef,omitempty"` //minOccurs="0" //Reference for the Booking Traveler. Used for Hotel Booking only. The value is arbitrary.
}

//Infomration about the Child guest.
type TypeGuestChildInformation struct {
	Age uint `xml:"Age,attr,omitempty"` //use="optional" //Age of the Child.
}

//Supported Provider TRM.
type PermittedAggregators struct {
	Aggregator []*Aggregator `xml:"hot:Aggregator"` //maxOccurs="99" //Supported Provider TRM.

}

//Supported Provider TRM.
type Aggregator struct {
	Name com.TypeThirdPartySupplier `xml:"Name,attr"` //use="required" //Supported Provider TRM, 2 byte aggrgator code.
}

//Details to request Hotel rate rules post shopping request.
type HotelRulesLookup struct {
	HotelProperty       *HotelProperty       `xml:"hot:HotelProperty"`
	HotelStay           *HotelStay           `xml:"hot:HotelStay"`
	HotelRulesModifiers *HotelRulesModifiers `xml:"hot:HotelRulesModifiers,omitempty"` //minOccurs="0"
	RatePlanType        string               `xml:"RatePlanType,attr"`                 //use="required" //This is room rate plan type for a particular rate plan
	Base                com.TypeMoney        `xml:"Base,attr"`                         //use="required" //This is the Base Amount for the selected rate plan type as received in Hotel Details/Book/Modify Response.
	RulesDetailReqd     string               `xml:"RulesDetailReqd,attr,omitempty"`    //use="optional" //Request details for Rules, Details, or All. Default is All. Applicable for 1p/1j.
}

//Controls and switches for the Hotel Rules request
type HotelRulesModifiers struct {
	PermittedProviders  *comrq.PermittedProviders    `xml:"com:PermittedProviders,omitempty"`  //minOccurs="0"
	NumberOfChildren    *NumberOfChildren            `xml:"hot:NumberOfChildren,omitempty"`    //minOccurs="0"
	HotelBedding        []*HotelBedding              `xml:"hot:HotelBedding,omitempty"`        //minOccurs="0" maxOccurs="4"
	CorporateDiscountID []*comrq.CorporateDiscountID `xml:"com:CorporateDiscountID,omitempty"` //minOccurs="0" maxOccurs="2"
	NumberOfAdults      int                          `xml:"NumberOfAdults,attr,omitempty"`     //use="optional" //Defaults to 1 if not specified
	NumberOfRooms       int                          `xml:"NumberOfRooms,attr,omitempty"`      //use="optional" default="1" /The numbers of rooms,defaults to 1 if not specified
	TotalOccupants      int                          `xml:"TotalOccupants,attr,omitempty"`     //use="optional" //Number of guests for the room. Supported Providers: 1P/1J
	Key                 com.TypeRef                  `xml:"Key,attr,omitempty"`                //use="optional"
}

//Marketing information provided by the supplier
type MarketingMessage struct {
	Text []string `xml:"hot:Text,omitempty"` //minOccurs="0" maxOccurs="99"
}

type Commission struct {
	Indicator                    com.TypeTrinary `xml:"Indicator,attr"`                    //Indicates if the Rate Plan is commissionable.True: Rate is commissionable.False: Rate is not commissionable.Unknown: Commission indicator is not returned by the hotel supplier (chain or property).
	Percent                      string          `xml:"Percent,attr"`                      //The percentage applied to the commissionable amount to determine the payable commission amount.
	CommissionAmount             com.TypeMoney   `xml:"CommissionAmount,attr"`             //The commission amount in the aggregator’s or supplier’s currency. For TRM, this amount may also include additional fees. TRM only.
	ApproxCommissionAmount       com.TypeMoney   `xml:"ApproxCommissionAmount,attr"`       //The approximate commission amount in the agency’s provisioned or requested currency. For TRM, this amount may also include additional fees. TRM only.
	CommissionOnSurcharges       com.TypeMoney   `xml:"CommissionOnSurcharges,attr"`       //Commission on surcharges in the aggregator’s or supplier’s currency. TRM only.
	ApproxCommissionOnSurcharges com.TypeMoney   `xml:"ApproxCommissionOnSurcharges,attr"` //The approximate commission on surcharges in the agency’s provisioned or requested currency. TRM only.
}

type AcceptedPayment struct {
	PaymentCode com.TypeCardMerchantType `xml:"PaymentCode,attr"` //The issuer of the form of payment, such as the credit card bank.
}
