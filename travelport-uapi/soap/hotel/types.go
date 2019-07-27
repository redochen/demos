package hotel

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
)

//Transportation Type
type TypeTransportationType string

//Numeric hotel rating
type TypeSimpleHotelRating int

//Bedding types
type TypeBedding string

//The 1 – 3  byte Open Travel Alliance amenity code.
type TypeAmenity int

//Date without time zones YYYY-MM-DD
type TypeDate string

//hotel location specified as IATA code or TRM custom codes
type TypeHotelLocationCode string

type TypeSourceLink bool

//Hotel specific rate categories (e.g. Standard, Military)
type TypeHotelRateCategory string

type TypeRateRuleDetail string

//Indicates the supplier of the rate. Maybe required for hotels provided by aggregators.
type TypeRateOfferId string

//Type definition of Room view code
type TypeRoomView string

//max length:30
type TypePromoCodeString string

//Indicates the commission during Hotel reservation. Provider supoorted 1P and 1J.
type HotelCommission string

//Hotel Confirmation Number from a Third Party System. max length:32
type TypeHotelConfirmationNumber string

type HotelSpecialRequest com.TypeGeneralText
