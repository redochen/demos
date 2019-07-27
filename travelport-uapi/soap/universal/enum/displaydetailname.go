package enum

//Retain the Reservation (do not cancel) in the the event of a schedule or price change
type EnumDisplayDetailName int

const (
	DdnOwningAgencyPCC              EnumDisplayDetailName = iota //The pseudo city code of the agency owning the PNR
	DdnCreatingAgentSignOn                                       //Sign on code of the agent created the PNR
	DdnCreatingAgentDuty                                         //Duty code of the agent created the PNR
	DdnCreatingAgencyIATA                                        //IATA number of the agency  created the PNR
	DdnPrepaidTicketAdviceIndicator                              //Indicates whether Prepaid Ticket Advice is attached to PNR
	DdnBFBorrowed                                                //Indicates the current PNR as borrowed
	DdnGlobalPNR                                                 //Indicates whether current PNR is Global PNR
	DdnReadOnlyPNR                                               //Indicates PNR as read only
	DdnPastDateQuickRetrievedPNR                                 //Indicates PNR retrieved from offline archival process
	DdnSuperPNR                                                  //Indicates whether current PNR is Super PNR
	DdnPNRPurgeDate                                              //Purge date of the PNR
	DdnOriginalReceivedFieldValue                                //The original data in the received field
	DdnDivPsgrName                                               //Passenger Name in a divided PNR
	DdnTruncInd                                                  //Indicator for Truncated names
	DdnDivTypeInd                                                //Divide Type Indicator if it pertains to child or parent booking
	DdnRLoc                                                      //Record Locator of parent or child booking
	DdnDivDt                                                     //Date of divide action
	DdnDivTm                                                     //Time of day of divide action
	DdnTravelOrder                                               //Travel order of the segment
	DdnSegmentStatus                                             //Status of associated segment
	DdnStartDate                                                 //Date of departure of the rail segment
	DdnDayChange                                                 //Indicates arrival date as number of days before or after departure date
	DdnVendor                                                    //Vendor code of the segment
	DdnStartTime                                                 // Start Time of the segment
	DdnEndTime                                                   //End Time of the segment
	DdnBookingCode                                               //The booking code of the rail segment seating area
	DdnTrainNumber                                               //Denotes Train number
	DdnNumberOfSeats                                             //Number of seats sold for the trip
	DdnSellType                                                  //The mode of selling the segment
	DdnTariffType                                                //Type of Tariff for the segment
	DdnConfirmationNumber                                        //The confirmation number of the segment
	DdnBoardingInformation                                       //Information related to boarding point of the segment
	DdnArrivalInformation                                        //Information related to arrival point of the segment
	DdnText                                                      //Additional information
)

//获取EnumDisplayDetailName的字符串值
func (this EnumDisplayDetailName) String() string {
	switch this {
	case DdnOwningAgencyPCC:
		return "OwningAgencyPCC"
	case DdnCreatingAgentSignOn:
		return "CreatingAgentSignOn"
	case DdnCreatingAgentDuty:
		return "CreatingAgentDuty"
	case DdnCreatingAgencyIATA:
		return "CreatingAgencyIATA"
	case DdnPrepaidTicketAdviceIndicator:
		return "PrepaidTicketAdviceIndicator"
	case DdnBFBorrowed:
		return "BFBorrowed"
	case DdnGlobalPNR:
		return "GlobalPNR"
	case DdnReadOnlyPNR:
		return "ReadOnlyPNR"
	case DdnPastDateQuickRetrievedPNR:
		return "PastDateQuickRetrievedPNR"
	case DdnSuperPNR:
		return "SuperPNR"
	case DdnPNRPurgeDate:
		return "PNRPurgeDate"
	case DdnOriginalReceivedFieldValue:
		return "OriginalReceivedFieldValue"
	case DdnDivPsgrName:
		return "DivPsgrName"
	case DdnTruncInd:
		return "TruncInd"
	case DdnDivTypeInd:
		return "DivTypeInd"
	case DdnRLoc:
		return "RLoc"
	case DdnDivDt:
		return "DivDt"
	case DdnDivTm:
		return "DivTm"
	case DdnTravelOrder:
		return "TravelOrder"
	case DdnSegmentStatus:
		return "SegmentStatus"
	case DdnStartDate:
		return "StartDate"
	case DdnDayChange:
		return "DayChange"
	case DdnVendor:
		return "Vendor"
	case DdnStartTime:
		return "StartTime"
	case DdnEndTime:
		return "EndTime"
	case DdnBookingCode:
		return "BookingCode"
	case DdnTrainNumber:
		return "TrainNumber"
	case DdnNumberOfSeats:
		return "NumberOfSeats"
	case DdnSellType:
		return "SellType"
	case DdnTariffType:
		return "TariffType"
	case DdnConfirmationNumber:
		return "ConfirmationNumber"
	case DdnBoardingInformation:
		return "BoardingInformation"
	case DdnArrivalInformation:
		return "ArrivalInformation"
	case DdnText:
		return "Text"
	}
	return ""
}

//获取EnumDisplayDetailName的整数值
func (this EnumDisplayDetailName) Value() int {
	if this >= DdnOwningAgencyPCC && this <= DdnText {
		return int(this)
	} else {
		return -1
	}
}
