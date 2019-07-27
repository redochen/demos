package response

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
)

//A single hotel availabilty result.
type HotelSearchResult struct {
	VendorLocation      []*comrs.VendorLocation      `xml:"VendorLocation,omitempty"`      //minOccurs="0" maxOccurs="unbounded"
	HotelProperty       []*HotelProperty             `xml:"HotelProperty"`                 //minOccurs="1" maxOccurs="unbounded" //he hotel property. Multiple property can only be supported with TRM and GDS property aggrigation.
	HotelSearchError    []*HotelSearchError          `xml:"HotelSearchError,omitempty"`    //minOccurs="0" maxOccurs="unbounded"
	CorporateDiscountID []*comrs.CorporateDiscountID `xml:"CorporateDiscountID,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	RateInfo            []*RateInfo                  `xml:"RateInfo,omitempty"`            //minOccurs="0" maxOccurs="unbounded"
	MediaItem           *comrs.MediaItem             `xml:"MediaItem,omitempty"`           //minOccurs="0"
	HotelType           *HotelType                   `xml:"HotelType,omitempty"`           //minOccurs="0" //Supported Providers:1P/1J
	PropertyDescription *PropertyDescription         `xml:"PropertyDescription,omitempty"` //minOccurs="0" //Hotel Property description. Maximum of 100 words returned. Supported Providers: TRM
}

//The hotel property
type HotelProperty struct {
	PropertyAddress       *UnstructuredAddress       `xml:"PropertyAddress,omitempty"`            //minOccurs="0"
	PhoneNumber           []*comrs.PhoneNumber       `xml:"PhoneNumber,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	CoordinateLocation    *comrs.CoordinateLocation  `xml:"CoordinateLocation,omitempty"`         //minOccurs="0"
	Distance              *comrs.Distance            `xml:"Distance,omitempty"`                   //minOccurs="0"
	HotelRating           []*HotelRating             `xml:"HotelRating,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	Amenities             *Amenities                 `xml:"Amenities,omitempty"`                  //minOccurs="0"
	MarketingMessage      *MarketingMessage          `xml:"MarketingMessage,omitempty"`           //MarketingMessage
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

//Hotel rating information
type HotelRating struct {
	//<xs:choice>
	Rating      []hot.TypeSimpleHotelRating `xml:"Rating"`      //maxOccurs="unbounded" //Hotel rating value
	RatingRange *RatingRange                `xml:"RatingRange"` //Search for a range of ratings
	//</xs:choice>
	RatingProvider string `xml:"RatingProvider,attr"` //use="required" //Rating providers, ie AAA, NTM
}

//Search for a range of rating
type RatingRange struct {
	MinimumRating hot.TypeSimpleHotelRating `xml:"MinimumRating,attr,omitempty"` //use="optional"
	MaximumRating hot.TypeSimpleHotelRating `xml:"MaximumRating,attr,omitempty"` //use="optional"
}

//Amenity information
type Amenities struct {
	Amenity []*Amenity `xml:"Amenity,omitempty"` //minOccurs="0" maxOccurs="8"
}

type Amenity struct {
	Code        hot.TypeAmenity `xml:"Code,attr"`                  //use="required"
	AmenityType string          `xml:"AmenityType,attr,omitempty"` //use="optional" //Amenity type code. “HA” (Hotel Property Amenity) or “RA” (Room Amenity). Defaults to “HA” if no value is sent.
}

type RateInfo struct {
	RoomDispatch                 *TypeRoomDispatch          `xml:"RoomDispatch,omitempty"`                      //minOccurs="0" //Breakdown for the aggregator’s multiple room offer by number of rooms and guest (adult/child capacity) for each room.  Supported Providers TRM.
	PolicyCodesList              *comrs.PolicyCodesList     `xml:"PolicyCodesList,omitempty"`                   //minOccurs="0" //A list of codes that indicate why an item was determined to be ‘out of policy’.
	MinimumAmount                com.TypeMoney              `xml:"MinimumAmount,attr,omitempty"`                //use="optional" //The low end of the nightly price range
	ApproximateMinimumAmount     com.TypeMoney              `xml:"ApproximateMinimumAmount,attr,omitempty"`     //use="optional" //The low end of the nightly price range in another currency
	MinAmountRateChanged         bool                       `xml:"MinAmountRateChanged,attr,omitempty"`         //use="optional" //Indicates the low end price range changes over the requested stay
	MaximumAmount                com.TypeMoney              `xml:"MaximumAmount,attr,omitempty"`                //use="optional" //The high end of the nightly price range
	ApproximateMaximumAmount     com.TypeMoney              `xml:"ApproximateMaximumAmount,attr,omitempty"`     //use="optional" //The high end of the nightly price range in another currency
	MaxAmountRateChanged         bool                       `xml:"MaxAmountRateChanged,attr,omitempty"`         //use="optional" //Indicates the high end price range changes over the requested stay
	MinimumStayAmount            com.TypeMoney              `xml:"MinimumStayAmount,attr,omitempty"`            //use="optional" //The low end of the price range for the entire stay
	ApproximateMinimumStayAmount com.TypeMoney              `xml:"ApproximateMinimumStayAmount,attr,omitempty"` //use="optional" //The low end of the price range for the entire stay in another currency
	Commission                   string                     `xml:"Commission,attr,omitempty"`                   //Commission information for this rate supplier
	RateSupplier                 com.TypeThirdPartySupplier `xml:"RateSupplier,attr,omitempty"`                 //use="optional" //Indicates the supplier of the rate.
	RateSupplierLogo             string                     `xml:"RateSupplierLogo,attr,omitempty"`             //use="optional" //Url of the supplier's logo
	PaymentType                  string                     `xml:"PaymentType,attr,omitempty"`                  //use="optional" //Payment type. “PrePay” rates require advance payment to complete the booking. “PostPay” rates allow payment after booking, typically after the hotel stay is completed. By default, all payment types are returned. Supported Providers TRM.
	//<xs:attributeGroup name="attrPolicyMarkingMaxMinPolicyCodes">
	MinInPolicy bool `xml:"MinInPolicy,attr,omitempty"` //use="optional" //This attribute will be used to indicate if the minimum fare or rate has been determined to be ‘in policy’ based on the associated policy settings.
	MaxInPolicy bool `xml:"MaxInPolicy,attr,omitempty"` //use="optional" //This attribute will be used to indicate if the maximum fare or rate has been determined to be ‘in policy’ based on the associated policy settings.
	//</xs:attributeGroup>
	ApproxAvgNightlyAmt com.TypeMoney `xml:"ApproxAvgNightlyAmt,attr,omitempty"` //use="optional" //The calculated average nightly price of the minimum stay amount in the user’s preferred currency. Supported Providers TRM.
	TaxesIncluded       bool          `xml:"TaxesIncluded,attr,omitempty"`       //use="optional" //Indicates if taxes are included in the rate. Supported Providers TRM.
	AmountConverted     bool          `xml:"AmountConverted,attr,omitempty"`     //use="optional" //If true, Amount is converted to user’s provisioned or preferred currency. If false, Amount returned in the approximate currency.  Supported Providers TRM.
	MultipleRoom        string        `xml:"MultipleRoom,attr,omitempty"`        //use="optional" //Indicates if multiple rooms functionality is supported for this aggregator and offer. Possible values are Unsupported, Unknown, and Available. Supported Providers TRM.
	PackageOffer        string        `xml:"PackageOffer,attr,omitempty"`        //use="optional" //Indicates if the aggregator supports a package offer for the multiple room result. Example: True/False or AnyType  Supported Providers TRM.
}

type HotelType struct {
	SourceLink hot.TypeSourceLink `xml:"SourceLink,attr,omitempty"` //use="optional" //Indicates whether results are returned from the vendor or from the database. If true, vendor results were returned. Supported providers:1G, 1V
}

type PropertyDescription struct {
	Value        string               `xml:",innerxml"`
	ProviderCode com.TypeProviderCode `xml:"ProviderCode,attr"` //use="required" //The host associated with this token
}

type HotelReferencePoint struct {
	Value   com.TypeReferencePoint `xml:",innerxml"`
	Country com.TypeCountry        `xml:"Country,attr,omitempty"` //Country code.
	State   com.TypeState          `xml:"State,attr,omitempty"`   //State or Province Code.
}

type HotelSearchError struct {
	comrs.ResultMessage
	RateSupplier com.TypeThirdPartySupplier `xml:"RateSupplier,attr,omitempty"` //use="optional" //Indicates the supplier of the rate.
}

//A simple unstructured address (e.g. 123 South State Avenue, Chicago, IL 60612)
type UnstructuredAddress struct {
	Address []string `xml:"Address"` //maxOccurs="6"
}

type HotelPropertyWithMediaItems struct {
	HotelProperty      *HotelProperty         `xml:"HotelProperty"`
	MediaItem          []*comrs.MediaItem     `xml:"MediaItem,omitempty"`          //minOccurs="0" maxOccurs="unbounded" //Photos and other media urls for the item referenced above.
	MediaResultMessage []*comrs.ResultMessage `xml:"MediaResultMessage,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Errors, Warnings and informational messages for the property referenced above.
}

//Returns Hotel rate details for requested hotel property.
type RequestedHotelDetails struct {
	TypeHotelDetails
	HotelType *HotelType `xml:"HotelType,omitempty"` //minOccurs="0" //Supported Providers:1G/1V/1P/1J.
}

//Alternate Properties returned by some Vendors if the requested property is not available
type HotelAlternateProperties struct {
	HotelProperty []*HotelProperty `xml:"HotelProperty"` //maxOccurs="unbounded"
}

//Textual information about the hotel
type HotelDetailItem struct {
	Text []string `xml:"Text"`      //maxOccurs="99"
	Name string   `xml:"Name,attr"` //use="required"
}

//Returns hotel rate details during the stay if rates are available for requested property.
type HotelRateDetail struct {
	PolicyCodesList      *comrs.PolicyCodesList       `xml:"PolicyCodesList,omitempty"`      //minOccurs="0" //A list of codes that indicate why an item was determined to be ‘out of policy’.
	RoomRateDescription  []*HotelRateDescription      `xml:"RoomRateDescription,omitempty"`  //minOccurs="0" maxOccurs="9999"
	HotelRateByDate      []*HotelRateByDate           `xml:"HotelRateByDate,omitempty"`      //minOccurs="0" maxOccurs="9999"
	CorporateDiscountID  []*comrs.CorporateDiscountID `xml:"CorporateDiscountID,omitempty"`  //minOccurs="0" maxOccurs="9999" //Corporate Discount IDs and Negotiate rate codes associated with this rate
	AcceptedPayment      []*AcceptedPayment           `xml:"AcceptedPayment,omitempty"`      //minOccurs="0" maxOccurs="99" //Form of payment accepted by the hotel supplier (chain or property). For credit cards, the two-character code for the credit card type is used.
	Commission           *Commission                  `xml:"Commission,omitempty"`           //minOccurs="0" maxOccurs="1" //Commission associated with the Rate Plan, as a percentage or flat amount.
	RateMatchIndicator   []*RateMatchIndicator        `xml:"RateMatchIndicator,omitempty"`   //minOccurs="0" maxOccurs="9999" //Returns "Match" Indicators for certain request parameters for Hotel Rate returned in response.
	TaxDetails           *TaxDetails                  `xml:"TaxDetails,omitempty"`           //minOccurs="0"
	CancelInfo           *CancelInfo                  `xml:"CancelInfo,omitempty"`           //minOccurs="0"
	GuaranteeInfo        *GuaranteeInfo               `xml:"GuaranteeInfo,omitempty"`        //minOccurs="0" //Guarantee, deposit, and prepayment information
	SupplementalRateInfo string                       `xml:"SupplementalRateInfo,omitempty"` //minOccurs="0" //Supplemental rate information provided by the aggregator. Supported Providers TRM.
	RoomCapacity         *RoomCapacity                `xml:"RoomCapacity,omitempty"`         //minOccurs="0" //The maximum number of guests for a room or for each room in a package. Provider: TRM.
	RatePlanType         com.TypeRatePlanType         `xml:"RatePlanType,attr"`              //use="required"
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
	Capacity  []uint `xml:"Capacity,omitempty"`       //minOccurs="0" maxOccurs="99" //The maximum number of guests per room. Provider: TRM.
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
	Tax []*Tax `xml:"Tax"` //minOccurs="1" maxOccurs="unbounded"
}

type Tax struct {
	//<xs:choice>
	Amount     com.TypeMoney `xml:"Amount,omitempty"`     //minOccurs="1"
	Percentage float32       `xml:"Percentage,omitempty"` //minOccurs="1"
	//</xs:choice>
	Code           com.TypeOTACode `xml:"Code,attr"`                     //use="required" //Code identifying fee (e.g. agency fee, bed tax etc.). Refer to OPEN Travel Code List for Fee Tax Type. Possible values are OTA Code against FTT.
	EffectiveDate  string          `xml:"EffectiveDate,attr,omitempty"`  //type="xs:date" use="optional"
	ExpirationDate string          `xml:"ExpirationDate,attr,omitempty"` //type="xs:date" use="optional"
	Term           string          `xml:"Term,attr,omitempty"`           //use="optional" //Indicates how the tax is applied. Values can be PerPerson, PerNight and PerStay
	CollectionFreq string          `xml:"CollectionFreq,attr,omitempty"` //use="optional" //Indicates how often the tax is collected. Values can be Once or Daily
}

//Returns cancellation information for certain hotel returned in response. This information is available through GDS transactions
type CancelInfo struct {
	CancellationPolicy            string          `xml:"CancellationPolicy,omitempty"`                 //minOccurs="0" //Return cancellation policy text by the aggregator. Provider: TRM.
	Text                          []string        `xml:"Text,omitempty"`                               //minOccurs="0" maxOccurs="99" //The informational text provided by the supplier to cancel the booking, if @Method="INFO". For all other values of @Method, Text is not returned. Provider: TRM.
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
	DepositAmount  *DepositAmount `xml:"DepositAmount,omitempty"`  //minOccurs="0" //Amount required for deposit/prepayment
	DepositNights  int            `xml:"DepositNights,omitempty"`  //Number of Nights required for deposit/prepayment
	DepositPercent int            `xml:"DepositPercent,omitempty"` //Percentage of stay required for deposit/prepayment
	//</xs:choice>
	GuaranteePaymentType []*GuaranteePaymentType `xml:"GuaranteePaymentType,omitempty"`     //minOccurs="0" maxOccurs="unbounded" //Accepted payment types
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

//Comments and Reviews from hotel guests
type GuestReviews struct {
	Comments []*Comments `xml:"Comments"` //maxOccurs="999"
}

type Comments struct {
	Value             string           `xml:",innerxml"`
	CommentId         com.TypeRef      `xml:"CommentId,attr,omitempty"`         //use="optional" //Unique comment identifier. For internal Travelport use only. Provider: TRM.
	Date              string           `xml:"Date,attr,omitempty"`              //use="optional" //date that the comment was entered. Supported Providers TRM.
	CommenterLanguage com.TypeLanguage `xml:"CommenterLanguage,attr,omitempty"` //use="optional" //Language of the commenter. Supported Providers TRM.
	Source            string           `xml:"Source,attr,omitempty"`            //use="optional" //Source code of the comment entry. Example: 'NB' for Nightsbridge, ‘RG’, ‘AG’ for Agrialla,‘TO’. Supported Providers TRM.
	CommentSourceName string           `xml:"CommentSourceName,attr,omitempty"` //use="optional" //Name of the source for the comment. Supported Providers TRM.
	Commenter         string           `xml:"Commenter,attr,omitempty"`         //use="optional" //Name of the comment's poster. Supported Providers TRM.
}

//Hotel Details Type Supported Providers TRM.
type TypeHotelDetails struct {
	HotelProperty   *HotelProperty     `xml:"HotelProperty"`
	HotelDetailItem []*HotelDetailItem `xml:"HotelDetailItem,omitempty"` //minOccurs="0" maxOccurs="999"
	HotelRateDetail []*HotelRateDetail `xml:"HotelRateDetail,omitempty"` //minOccurs="0" maxOccurs="999" //Returns hotel rate details during the stay if rates are available for requested property
	MediaItem       []*comrs.MediaItem `xml:"MediaItem,omitempty"`       //minOccurs="0" maxOccurs="999"
}

//Supported Provider TRM.
type AggregatorHotelDetails struct {
	TypeHotelDetails
	//<xs:attributeGroup name="attrAggregatorHotelDetail">
	Aggregator          string `xml:"Aggregator,attr,omitempty"`          //use="optional" //Two-character aggregator code. Provider: TRM.
	AggregatorName      string `xml:"AggregatorName,attr,omitempty"`      //use="optional" //The TRM aggregator name. Provider: TRM.
	RulesAndRestriction string `xml:"RulesAndRestriction,attr,omitempty"` //use="optional" //The URL for the aggregator’s rules and restrictions. Provider: TRM.
	TermsAndConditions  string `xml:"TermsAndConditions,attr,omitempty"`  //use="optional" //The URL for the aggregator’s terms and conditions. Provider: TRM.
	SupportsPayment     bool   `xml:"SupportsPayment,attr,omitempty"`     //use="optional" //If true, the aggregator supports booking for this supplier (hotel property or chain). Provider: TRM.
	CommissionModel     string `xml:"CommissionModel,attr,omitempty"`     //use="optional" //The aggregator’s commission (marketing fee) model. “Default” is Travelport Accounts Payable, in which Travelport settles the commission. Provider: TRM.
	MultiRoomSupport    string `xml:"MultiRoomSupport,attr,omitempty"`    //use="optional" //The type of multi-room support type. SINGLE (one room), SAME_TYPE (same room types for all rooms), ANY_TYPE (different room types for multiple rooms). Provider: TRM.
	SupportsChildren    bool   `xml:"SupportsChildren,attr,omitempty"`    //use="optional" //If true, the aggregator supports children. If false, the aggregator does not support children and may treat child as an adult in the booking. Provider: TRM.
	ChildStartAge       int    `xml:"ChildStartAge,attr,omitempty"`       //use="optional" //The minimum age of a child as defined by the aggregator. Children under this age may be booked as an infant by the aggregator. Provider: TRM.
	AdultStartAge       int    `xml:"AdultStartAge,attr,omitempty"`       //use="optional" //The minimum age of an adult as defined by the aggregator. Guests under this age may be booked as a child by the aggregator. Provider: TRM.
	MaxChildrenPerRoom  uint   `xml:"MaxChildrenPerRoom,attr,omitempty"`  //use="optional" //The maximum number of children per room. If the requested number of children is more than MaxChildrenPerRoom, the aggregator may book excess children as adults. Provider: TRM.
	PhoneNumber         string `xml:"PhoneNumber,attr,omitempty"`         //use="optional" //Aggregator Phone number. Supported Providers TRM.
	AreaCode            string `xml:"AreaCode,attr,omitempty"`            //use="optional" //Aggregator Phone number area code. Supported Providers TRM.
	CountryCode         string `xml:"CountryCode,attr,omitempty"`         //use="optional" //Aggregator Phone number country code. Supported Providers TRM.
	//</xs:attributeGroup>
}

//Breakdown for the aggregator’s multiple room offer by number of rooms and guest (adult/child capacity) for each room. Supporrted Providers TRM.
type TypeRoomDispatch struct {
	Room []*Room `xml:"Room,omitempty"` //minOccurs="0" maxOccurs="9" //Room Details. Supported providers TRM.
}

//Supported providers TRM.
type Room struct {
	Capacity uint `xml:"Capacity,attr,omitempty"` //use="optional" //The maximum number of guests in the room available for the multiple room result.  Supported Providers TRM.
	Quantity uint `xml:"Quantity,attr,omitempty"` //use="optional" //The number of rooms available for the multiple room result. Supported providers TRM.
}

//Applicable for 1G, 1V, 1P, 1J
type HotelRateChangedInfo struct {
	HotelProperty   *HotelProperty   `xml:"HotelProperty"`
	HotelRateDetail *HotelRateDetail `xml:"HotelRateDetail,omitempty"` //minOccurs="0"
	Reason          string           `xml:"Reason,attr"`               //Reason to represent whether rate change in hotel rules.Applicable for 1G, 1V, 1P, 1J
}

//The complete Hotel Reservation
type HotelReservation struct {
	comrs.BaseReservation
	BookingTravelerRef []*comrs.BookingTravelerRef `xml:"BookingTravelerRef,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	//<xs:group name="BaseHotelReservationGroup">
	ReservationName            *comrs.ReservationName       `xml:"ReservationName,omitempty"`       //minOccurs="0"
	ThirdPartyInformation      *comrs.ThirdPartyInformation `xml:"ThirdPartyInformation,omitempty"` //minOccurs="0"
	HotelProperty              *HotelProperty               `xml:"HotelProperty"`
	HotelRateDetail            []*HotelRateDetail           `xml:"HotelRateDetail"` //maxOccurs="99"
	HotelStay                  *HotelStay                   `xml:"HotelStay"`
	HotelSpecialRequest        hot.HotelSpecialRequest      `xml:"HotelSpecialRequest,omitempty"`        //minOccurs="0"
	Guarantee                  *comrs.Guarantee             `xml:"Guarantee,omitempty"`                  //minOccurs="0"
	PromotionCode              *PromotionCode               `xml:"PromotionCode,omitempty"`              //minOccurs="0" //Specifies promotional code used in hotel booking
	BookingSource              *comrs.BookingSource         `xml:"BookingSource,omitempty"`              //minOccurs="0" //Specify alternate booking source
	HotelBedding               []*HotelBedding              `xml:"HotelBedding,omitempty"`               //minOccurs="0" maxOccurs="4"
	GuestInformation           *GuestInformation            `xml:"GuestInformation,omitempty"`           //minOccurs="0"
	AssociatedRemark           []*AssociatedRemark          `xml:"AssociatedRemark,omitempty"`           //minOccurs="0" maxOccurs="9999"
	SellMessage                []com.SellMessage            `xml:"SellMessage,omitempty"`                //minOccurs="0" maxOccurs="9999"
	HotelCommission            hot.HotelCommission          `xml:"HotelCommission,omitempty"`            //minOccurs="0" //HotelCommission text indicates commision while hotel reservation. Provider supported 1P and 1J.
	BookingGuestInformation    *BookingGuestInformation     `xml:"BookingGuestInformation,omitempty"`    //minOccurs="0"
	RoomConfirmationCodes      *RoomConfirmationCodes       `xml:"RoomConfirmationCodes,omitempty"`      //minOccurs="0" //Individual room confirmation codes. Returns when rooms are booked as a package. Supported Providers TRM.
	CancelInfo                 *CancelInfo                  `xml:"CancelInfo,omitempty"`                 //minOccurs="0"
	TotalReservationPrice      *TotalReservationPrice       `xml:"TotalReservationPrice,omitempty"`      //minOccurs="0"
	HotelDetailItem            []*HotelDetailItem           `xml:"HotelDetailItem,omitempty"`            //minOccurs="0" maxOccurs="99"
	AdaptedRoomGuestAllocation *AdaptedRoomGuestAllocation  `xml:"AdaptedRoomGuestAllocation,omitempty"` //minOccurs="0" //This element defines how the aggregators or hotel property have allocated the guests to the rooms. Only displayed when Requested guest allocation is different from the Adapted room guest allocation. Supported Providers TRM.
	//</xs:group>
	Status                            string      `xml:"Status,attr"`                                      //use="required" //Reservation IATA status code, 2 byte.
	AggregatorBookingStatus           string      `xml:"AggregatorBookingStatus,attr,omitempty"`           //use="optional" //Aggregator reservation status response. Supported Provider TRM.
	BookingConfirmation               string      `xml:"BookingConfirmation,attr,omitempty"`               //use="optional"
	CancelConfirmation                string      `xml:"CancelConfirmation,attr,omitempty"`                //use="optional"
	ProviderReservationInfoRef        com.TypeRef `xml:"ProviderReservationInfoRef,attr,omitempty"`        //use="optional" //Provider reservation reference key.
	TravelOrder                       int         `xml:"TravelOrder,attr,omitempty"`                       //use="optional" //To identify the appropriate sequence for Air/Car/Hotel segments based on travel dates.
	ProviderSegmentOrder              int         `xml:"ProviderSegmentOrder,attr,omitempty"`              //use="optional" //To identify the appropriate travel sequence for Air/Car/Hotel/Rail segments/reservations in the provider reservation.
	PassiveProviderReservationInfoRef com.TypeRef `xml:"PassiveProviderReservationInfoRef,attr,omitempty"` //use="optional" //Passive Provider reservation reference key.
}

//Arrival and Departure dates
type HotelStay struct {
	CheckinDate  hot.TypeDate `xml:"CheckinDate"`
	CheckoutDate hot.TypeDate `xml:"CheckoutDate"`
	Key          com.TypeRef  `xml:"Key,attr,omitempty"` //use="optional"
}

type PromotionCode struct {
	Value string      `xml:",innerxml"`
	Key   com.TypeRef `xml:"Key,attr,omitempty"` //use="optional"
}

//Specify desired bedding
type HotelBedding struct {
	Type         string        `xml:"Type,attr"`                   //use="required" //Queen, King, double, etc
	NumberOfBeds int           `xml:"NumberOfBeds,attr,omitempty"` //use="optional" //Number of beds of desired Type in room. Use '0' to delete the hotel Optional Beds ( Only RA RC CR )
	Amount       com.TypeMoney `xml:"Amount,attr,omitempty"`       //use="optional" //Fee for bed type. Providers: 1g/1v/1p/1j
	Content      string        `xml:"Content,attr,omitempty"`      //use="optional" //Additional information Providers: 1p/1j
}

//The information like number of rooms ,number of adults,children to be provided while booking the  hotel
type GuestInformation struct {
	NumberOfAdults   *NumberOfAdults   `xml:"NumberOfAdults,omitempty"`     //minOccurs="0"
	NumberOfChildren *NumberOfChildren `xml:"NumberOfChildren,omitempty"`   //minOccurs="0"
	ExtraChild       *ExtraChild       `xml:"ExtraChild,omitempty"`         //minOccurs="0" //Providers: 1p/1j
	NumberOfRooms    int               `xml:"NumberOfRooms,attr,omitempty"` //use="optional"
}

type ExtraChild struct {
	Count   int    `xml:"Count,attr"`   //The number of extra children in the room
	Content string `xml:"Content,attr"` //Additional information
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
	Age    []int         `xml:"Age,omitempty"`        //minOccurs="0" maxOccurs="99" //The Ages of the Children. . The defined age of a Child traveler may vary by supplier, but is typically 1 to 17 years. Supported Providers 1G/1V.
	Count  int           `xml:"Count,attr"`           //use="required" //The total number of children in the booking. Supported Providers 1P/1J.
	Amount com.TypeMoney `xml:"Amountattr,omitempty"` //use="optional" //Fee per child. Providers: 1g/1v
}

type AssociatedRemark struct {
	comrs.AssociatedRemark
}

//Individual room confirmation codes. Returns when rooms are booked as a package. Supported Providers TRM.
type RoomConfirmationCodes struct {
	ConfirmationCode []string `xml:"ConfirmationCode"` //maxOccurs="9" //Individual room confirmation codes. Supported Providers TRM
}

//The total price for the entire stay, including fees, for all rooms in the booking. Provider: TRM.
type TotalReservationPrice struct {
	RoomRateDescription []*TypeHotelRateDescription `xml:"RoomRateDescription,omitempty"`   //minOccurs="0" maxOccurs="99"
	TotalPrice          com.TypeMoney               `xml:"TotalPrice,attr,omitempty"`       //use="optional" //The amount of the total price, including fees for all rooms in the booking. Provider: TRM.
	ApproxTotalPrice    com.TypeMoney               `xml:"ApproxTotalPrice,attr,omitempty"` //use="optional" //The approximate amount of the total hotel price, including fees, in another currency. Provider: TRM.
}

//This element defines how the aggregators or hotel property have allocated the guests to the rooms. Only displayed when Requested guest allocation is different from the Adapted room guest allocation. Supported Providers TRM.
type AdaptedRoomGuestAllocation struct {
	Room []*TypeAdaptedRoomGuestAllocation `xml:"Room"` //maxOccurs="9" //Individual room. Multiple occurrences if there are multiple rooms in the request. Maximum number of rooms may vary by supplier or aggregator.
}

//Information about guest to book. Supported Providers TRM.
type TypeGuestRoomInformation struct {
	Adults             int                         `xml:"Adults"`                       //The number of adult guests per room. Maximum number of adults may vary by supplier or aggregator.
	BookingTravelerRef []*comrs.BookingTravelerRef `xml:"BookingTravelerRef,omitempty"` //minOccurs="0" maxOccurs="9" //Reference for the Booking Traveler. Used for Hotel Booking only. The value is arbitrary.
	Child              []*Child                    `xml:"Child,omitempty"`              //minOccurs="0" maxOccurs="6" //Information about a child guest.
}

//Information about requested rooms and guests allocation. Supported Providers TRM.
type BookingGuestInformation struct {
	Room []*TypeGuestRoomInformation `xml:"Room"` //maxOccurs="9" //Individual room. Multiple occurrences if there are multiple rooms in the request. Maximum number of rooms may vary by supplier or aggregator.
}

//Information about a child guest.
type Child struct {
	TypeGuestChildInformation
	BookingTravelerRef *comrs.BookingTravelerRef `xml:"BookingTravelerRef,omitempty"` //minOccurs="0" //Reference for the Booking Traveler. Used for Hotel Booking only. The value is arbitrary.
}

//Infomration about the Child guest.
type TypeGuestChildInformation struct {
	Age uint `xml:"Age,attr,omitempty"` //use="optional" //Age of the Child.
}

type TypeHotelRateDescription struct {
	Text []string `xml:"Text"`                //maxOccurs="unbounded"
	Name string   `xml:"Name,attr,omitempty"` //use="optional" //Optional context name of the text block being returned i.e. Room details
}

//The allocation of guests per room assigned by the aggregator or supplier (hotel property). Returned only when the requested guest allocation is different from the provider or supplier’s adapted guest allocation. Supported Providers TRM.
type TypeAdaptedRoomGuestAllocation struct {
	Child          []*TypeGuestChildInformation `xml:"Child,omitempty"`     //minOccurs="0" maxOccurs="6" //Information about a child guest.
	NumberOfAdults uint                         `xml:"NumberOfAdults,attr"` //The number of adult guests per room. Maximum number of adults may vary by supplier or aggregator.
}

//Textual information about the hotel rule
type HotelRuleItem struct {
	Text []string `xml:"Text,omitempty"` //maxOccurs="unbounded"
	Name string   `xml:"Name,attr"`      //use="required"
}

//Marketing information provided by the supplier
type MarketingMessage struct {
	Text []string `xml:"Text,omitempty"` //minOccurs="0" maxOccurs="99"
}

type AcceptedPayment struct {
	PaymentCode *com.TypeCardMerchantType `xml:"PaymentCode,attr"` //The issuer of the form of payment, such as the credit card bank.
}

type Commission struct {
	Indicator                    com.TypeTrinary `xml:"Indicator,attr"`                    //Indicates if the Rate Plan is commissionable.True: Rate is commissionable.False: Rate is not commissionable.Unknown: Commission indicator is not returned by the hotel supplier (chain or property).
	Percent                      string          `xml:"Percent,attr"`                      //The percentage applied to the commissionable amount to determine the payable commission amount.
	CommissionAmount             com.TypeMoney   `xml:"CommissionAmount,attr"`             //The commission amount in the aggregator’s or supplier’s currency. For TRM, this amount may also include additional fees. TRM only.
	ApproxCommissionAmount       com.TypeMoney   `xml:"ApproxCommissionAmount,attr"`       //The approximate commission amount in the agency’s provisioned or requested currency. For TRM, this amount may also include additional fees. TRM only.
	CommissionOnSurcharges       com.TypeMoney   `xml:"CommissionOnSurcharges,attr"`       //Commission on surcharges in the aggregator’s or supplier’s currency. TRM only.
	ApproxCommissionOnSurcharges com.TypeMoney   `xml:"ApproxCommissionOnSurcharges,attr"` //The approximate commission on surcharges in the agency’s provisioned or requested currency. TRM only.
}
