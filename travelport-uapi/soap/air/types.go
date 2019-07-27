package air

//Layover duration range of valid values 0-9999
type MaxLayoverDurationRangeType int

//3 Letter equipment code (sometimes vary by carrier)
type TypeEquipment string

//Availability Source value for Sell.
type TypeAvailabilitySource string

//Available InFlight Services.
//They are: 'Movie', 'Telephone', 'Telex', 'AudioProgramming', 'Television' ,
//'ResvBookingService' ,'DutyFreeSales' ,'Smoking' ,'NonSmoking' ,'ShortFeatureVideo' ,
//'NoDutyFree' ,'InSeatPowerSource' ,'InternetAccess' ,'Email' ,'Library' ,'LieFlatSeat' ,
//'Additional service(s) exists' ,'WiFi' ,'Lie-Flat seat first' ,'Lie-Flat seat business' ,
//'Lie-Flat seat premium economy' ,'Amenities subject to change' etc..
//These follow the IATA standard. Please see the IATA standards for a more complete list.
type InFlightServices string

//Car code value. Maximum 15 characters. Applicable provider is 1G and 1V
type TypeCarCode string

//Value code value. Maximum 15 characters. Applicable provider is 1G and 1V
type TypeValueCode string

//The complete fare calculation line.
type FareCalc string

//Tour code value. Maximum 15 characters
type TypeTourCode string

//Ticket Designator type.Size can be up to 20 characters
type TypeTicketDesignator string

//ATPCO fare type code (e.g. XPN)
type TypeFareTypeCode string

//Currently returned: FullyRefundable (1G,1V), RefundableWithPenalty (1G,1V), Refundable (1P,1J),  NonRefundable (1G,1V,1P,1J).
type TypeRefundabilityValue string

//Reason codes for Fare rule failure. Following values will be supported.
/*                        MinimumStayFailure,
AdvPurchaseFailure,
PICTypeFailure [Passenger Identification Code Failure],
StopoverTransferFailure,
DateSeasonalityFailure,
RoutingFailure,
MileageFailure,
DayTimeFailure,
OpenJawUsageFailure,
IndirectTravelProvision,
SalesRestrictionNotMet,
FICNAFare,
HIFFailure [Higher Intermediate Fare/Point or Mileage Exceptions failures],
IntlSurfaceSector,
CurrencyUsageFailure,
DiscountApplFailure,
FootNoteFailure,
DayTimeApplCatNotMet,
DayTimeApplCatIncomplete,
SeasonalityCatNotMet,
SeasonalityCatIncomplete,
FlightApplCatNotMet,
FlightApplCatIncomplete,
AdvResvAdvTicketCatNotMet,
AdvResvAdvTicketCatIncomplete,
BookingClassFailure,
MinStayCatNotMet,
MinStayCatIncomplete,
StopoverCatNotMet,
StopoverCatIncomplete,
PermittedCombinationCatNotMet,
PermittedCombinationCatNotIncomplete,
BlackoutCatNotMet,
BlackoutCatNotIncomplete,
AccomTvlReqCatNotMet,
AccomTvlReqCatIncomplete,
SalesRestCatNotMet,
SalesRestCatIncomplete,
EligibilityCatNotMet,
EligibilityCatIncomplete,
TransfersCatNotMet,
TransfersCatIncomplete,
TransfersRoutingFailure,
HIPMileageCatNotMet [Higher Intermediate Point or Mileage Exception categories not met],
HIPMileageCatIncomplete [Higher Intermediate Point or Mileage Exceptions categories incomplete],
ChildDiscountCatNotMet,
ChildDiscountCatIncomplete,
TourConductorDiscCatNotMet,
TourConductorDiscCatIncomplete,
AgentDiscountCatNotMet,
AgentDiscountCatIncomplete,
OtherDiscountCatNotMet,
OtherDiscountCatIncomplete,
MiscFareTagCatNotMet,
MiscFareTagCatIncomplete,
FareByRuleCatNotMet,
FareByRuleCatIncomplete,
VisitCountryCatNotMet,
VisitCountryCatIncomplete,
NegFaresCatNotMet,
NegFaresCatIncomplete,
OthFailurePTCFailed,
OthFailureRecFailed,
CombineabilityFailure,
TravelRestrictionNotMet,
SurchargesNotMet,
MaximumStayFailure,
FareUsageFailure,
IATAFareNotValid,
RecordOneFailure,
Cat01EligibilityFailure,
FlightApplicationsFailure,
FootNoteFailure,
Cat11BlackOutfailure,
Cat13AccompaniedTravelRequirementFailure,
Cat19ChildDiscountFailure,
Cat20TourDiscountFailure,
Cat21AgentDiscountApplicationFail,
YYSuppressionTableFailure,
HostUseOnly87,
HostUseOnly88,
HostUseOnly89,
HostUseOnly90,
HostUseOnly99,
HostUseOnly100,
HostUseOnly105,
HostUseOnly108,
HostUseOnly111,
HostUseOnly112,
HostUseOnly113,
HostUseOnly114,
HostUseOnly121,
HostUseOnly122,
SpareForFutureIndicators,
YYSuppressionTableFailed,
HIPCheckFailed,
TourCodeFail,
AbonnmentFareFailure,
FailedSurfaceSector,
IndirectTravelFailed,
FailedCurrencyUsage,
CAT12NotMet
*/
type TypeFareRuleFailureInfoReason string

//The type of FeeApplication
type TypeFeeApplication string

//The type of AssessIndicator
type TypeAssessIndicator string

//"xs:duration" :PnYnMnDTnHnMnS
//P 表示周期
//nY 表示年数
//nM 表示月数
//nD 表示天数
//T 表示时间部分的起始 （如果您打算规定小时、分钟和秒，则此选项为必需）
//nH 表示小时数
//nM 表示分钟数
//nS 表示秒数
type Duration string

//获取总分钟数
func (this *Duration) GetTotalMinutes() int {
	return 0
}

//Version of the Universal record. Required with any request to modify the existing Universal record.
type TypeURVersion int

//The unique identifier of the brand
type TypeBrandId string
