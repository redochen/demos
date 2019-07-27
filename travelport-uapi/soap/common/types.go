package common

//5 Character Provider Code
type TypeProviderCode string

//A Locator Code that uniquely identifies a Provider Reservation or searches for one.
type TypeProviderLocatorCode string

//1 to 5 Character Supplier code
type TypeSupplierCode string

//flight number type.
type TypeFlightNumber int

//Valid 3 letter IATA city or airport code
type TypeIATACode string

//Class of service code (Booking code) usually one letter, rarely two.
type TypeClassOfService string

//Reference type
type TypeRef string

//2 Letter Carrier code
type TypeCarrier string

//Sell Message from Vendor. This is applicable in response messages only, any input in request message will be ignored.
type SellMessage string

//Passenger Type Code (ADT, A2B5)
type TypePTC string

//The gender of a person.  Data is defined in Ref Pub
type TypeGender string

//A type for unique party identifiers of any party role.
type TypeProfileID int

//At least one character data in Next Result Reference
type TypeNonBlanks string

//Used for Character Strings, length 1.
type StringLength1 string

//Used for Character Strings, length 1 to 8.
type StringLength1to8 string

//Used for Character Strings, length 1 to 13.
type StringLength1to13 string

//Used for Character Strings, length 1 to 16.
type StringLength1to16 string

//Used for Character Strings, length 1 to 20.
type StringLength1to20 string

//Used for Character Strings, length 1 to 32.
type StringLength1to32 string

//Used for Character Strings, length 1 to 64.
type StringLength1to64 string

//Used for Character Strings, length 1 to 128.
type StringLength1to128 string

//Used for Character Strings, length 1 to 255.
type StringLength1to255 string

//Used for Numeric values, from 0 to 999 inclusive.
type Numeric0to999 string

//Loyalty Card number with maximum length as 36.
type TypeCardNumber string

//[a-zA-Z0-9]{1,1}
type TypePriorityCode string

//2 Letter Country code
type TypeCountry string

//Defines the State code.
type TypeState string

//3 Letter City Code
type TypeCity string

//3 Letter Airport Code
type TypeAirport string

//2 to 10 Letter Pseudo City Code
type TypePCC string

//ARC/IATA code that represents a branch/agency.
type TypeIATA string

//Type for PolicyReference attribute.
type TypePolicyReference string

//3 Letter Currency Code
type TypeCurrency string

//A monetary value (valid to req/rsp Currency type) Format : Currency Code + Amount(USD123.10)
type TypeMoney string

//An alpha-numeric string which denotes fare family. Some carriers may return this in lieu of or in addition to the CabinClass.
type TypeFareFamily string

//Type for PolicyCode attribute.
type TypePolicyCode int

//A code that indicates why the minimum fare or rate was determined to be ‘out of policy’.
type MinPolicyCode string

//A code that indicates why the maximum fare or rate was determined to be ‘out of policy’.
type MaxPolicyCode string

//Tour code value. Maximum 15 characters
type TypeTourCode string

//A percentage that can include up to two decimal places ([0-9]{1,2}|100)\.[0-9]{1,2}
type TypePercentageWithDecimal string

//Common type for general textual information,max length:250
type TypeGeneralText string

//SSR Code, exactly 4 characters (e.g. DEAF, NSST, etc.)
type TypeSSRCode string

//SSR Free Text
type TypeSSRFreeText string

//2 Letter distance unit code
type TypeDistance string

//Endorsement type.Size can be up to 256 characters
type TypeEndorsement string

//An identifier that labels this Merchandising Service (Baggage, Nomiles,GroundTransportation etc)
type TypeMerchandisingService string

//The purchase windows available for merchandising service
type TypePurchaseWindow string

//A Locator Code that uniquely identifies a Record or searches for one.
type TypeLocatorCode string

//Type for Traveler Last Name.
type TypeTravelerLastName string

//An identifier that labels this email address (Personal, Business, Agency, etc)
type TypeEmailType string

//Mail comment is used to include one line of freeform information.
type TypeEmailComment string

//Defines the GDS remark types
//Only below mentioned values are Supported as typeGdsRemark
//    Alpha
//    Basic
//    Historical
//    Invoice
//    Itinerary
//    Vendor
//    Confidential
//    FOPComment (Currently this is only used by Worldspan and JAL.)
type TypeGdsRemark string

//Valid 2 letter Status Code
type TypeStatusCode string

//Valid 4 letter Seat Type Code
type TypeSeatTypeCode string

//Version of the Universal record. Required with any request to modify the existing Universal record.
type TypeURVersion string

//Defines the GDS Accounting remark types
//Only below mentioned values are Supported as typeGdsAccountingRemark
// Fare
// Canned
// Ticket
// Account
// Other
// InvoiceLayout
// ServiceFee
// AgentSign
// TourCode (1P)
// Endorsement (1P)
// CorporateTrackingId (1P)
// ItineraryInvoicePerTraveler (1P)
// ItineraryInvoicePerSurname (1P)
// DividerCard (1P)
// NetFare/VC/CAR (1P/1J)
// MarketCode (1V)
// BranchLocationOverride (1V)
// BookingAgentOverride (1V)
// SellingAgentOverride (1V)
// ProductTypeOverride (1V)
// TicketingAgent (1V)
// Sort (1V)
// PurchaseOrder (1V)
// ItineraryWithFare (1V)
// ItineraryWithoutFare (1V)
// ItineraryWithAmount (1V)
type TypeGdsAccountingRemark string

//Agency Branch Code Identifier
type TypeBranchCode string

//External reference string for Client application to identify the Form of Payment. Element will be a max of 32 hex characters alpha-numeric.
type TypeExternalReference string

//2 Letter Hotel Chain Code
type TypeHotelChainCode string

//Unique hotel identifier for the channel
type TypeHotelCode string

//Third Party Content Provider name.
type TypeThirdPartySupplier string

//
type TypeReferencePoint string

//Type of payment required to reserve travel i.e. Hotel Reservation requirement
type TypeReserveRequirement string

//Represents the rate plan code of room type for specified hotel property.
type TypeRatePlanType string

//2 letter Credit/Debit Card merchant type
type TypeCardMerchantType string

//The associated credit/debit card number without spaces or dashes.length range: 13-128
type TypeCreditCardNumber string

//Percentage value xs:nonNegativeInteger
type TypeIntegerPercentage int

//Value of the Duration in P[NumberOfDays]D format.Ranges Permitted are P001D to P366D.
type TypeDurationYearInDays string

//The identifying number for the actual ticket
type TicketNumber StringLength1to13

//2 Letter ISO Language code
type TypeLanguage string

//Refers to Open Travel Code
type TypeOTACode uint

//Extension of boolean, that allows for unknown values.:true/false/unknown
type TypeTrinary string

//EnumTypeFarePull
type TypeFarePull string
