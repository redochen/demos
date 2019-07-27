package request

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
)

//BaseCoreReq ...
type BaseCoreReq struct {
	BillingPointOfSaleInfo *BillingPointOfSaleInfo `xml:"com:BillingPointOfSaleInfo"`
	AgentIDOverride        []*AgentIDOverride      `xml:"com:AgentIDOverride,omitempty"`     //minOccurs="0" maxOccurs="unbounded"
	TerminalSessionInfo    string                  `xml:"com:TerminalSessionInfo,omitempty"` //minOccurs="0" maxOccurs="1"
	TraceID                string                  `xml:"TraceId,attr,omitempty"`            //use="optional" //Unique identifier for this atomic transaction traced by the user. Use is optional.
	TokenID                string                  `xml:"TokenId,attr,omitempty"`            //use="optional" //Authentication Token ID used when running in statefull operation. Obtained from the LoginRsp. Use is optional.
	AuthorizedBy           string                  `xml:"AuthorizedBy,attr,omitempty"`       //use="optional" //Used in showing who authorized the request. Use is optional.
	TargetBranch           string                  `xml:"TargetBranch,attr,omitempty"`       //use="optional" //Used for Emulation - If authorised will execute the request as if the agent's parent branch is the TargetBranch specified.
	OverrideLogging        string                  `xml:"OverrideLogging,attr,omitempty"`    //use="optional" //Use to override the default logging level
	LanguageCode           string                  `xml:"LanguageCode,attr,omitempty"`       //use="optional" //ISO 639 two-character language codes are used to retrieve specific information in the requested language. For Rich Content and Branding, language codes ZH-HANT (Chinese Traditional), ZH-HANS (Chinese Simplified), FR-CA (French Canadian) and PT-BR (Portuguese Brazil) can also be used. For RCH, language codes ENGB, ENUS, DEDE, DECH can also be used. Only certain services support this attribute. Providers: ACH, RCH, 1G, 1V, 1P, 1J.
}

//BaseCoreSearchReq Base Request for Air Search
type BaseCoreSearchReq struct {
	BaseCoreReq
	NextResultReference []*NextResultReference `xml:"com:NextResultReference,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//BaseReq ...
type BaseReq struct {
	BaseCoreReq
	OverridePCC                        *OverridePCC `xml:"com:OverridePCC,omitempty"`                         //minOccurs="0" maxOccurs="1"
	RetrieveProviderReservationDetails bool         `xml:"RetrieveProviderReservationDetails,attr,omitempty"` //use="optional" default="false"
}

//BaseSearchReq ...
type BaseSearchReq struct {
	BaseReq
	NextResultReference []*NextResultReference `xml:"com:NextResultReference,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//BaseCreateReservationReq ...
type BaseCreateReservationReq struct {
	BaseReq
	LinkedUniversalRecord      []*LinkedUniversalRecord `xml:"com:LinkedUniversalRecord,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	BookingTraveler            []*BookingTraveler       `xml:"com:BookingTraveler,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	OSI                        []*OSI                   `xml:"com:OSI,omitempty"`                         //minOccurs="0" maxOccurs="unbounded"
	AccountingRemark           []*AccountingRemark      `xml:"com:AccountingRemark,omitempty"`            //minOccurs="0" maxOccurs="unbounded"
	GeneralRemark              []*GeneralRemark         `xml:"com:GeneralRemark,omitempty"`               //minOccurs="0" maxOccurs="unbounded"
	XMLRemark                  []*XMLRemark             `xml:"com:XMLRemark,omitempty"`                   //minOccurs="0" maxOccurs="unbounded"
	UnassociatedRemark         []*UnassociatedRemark    `xml:"com:UnassociatedRemark,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	Postscript                 *Postscript              `xml:"com:Postscript,omitempty"`                  //minOccurs="0" maxOccurs="1"
	PassiveInfo                *PassiveInfo             `xml:"com:PassiveInfo,omitempty"`                 //minOccurs="0"
	ContinuityCheckOverride    *ContinuityCheckOverride `xml:"com:ContinuityCheckOverride,omitempty"`     //minOccurs="0" maxOccurs="1" //This element will be used if user wants to override segment continuity check.
	AgencyContactInfo          *AgencyContactInfo       `xml:"com:AgencyContactInfo,omitempty"`           //minOccurs="0"
	CustomerID                 *CustomerID              `xml:"com:CustomerID,omitempty"`                  //minOccurs="0"
	FileFinishingInfo          *FileFinishingInfo       `xml:"com:FileFinishingInfo,omitempty"`           //minOccurs="0"
	CommissionRemark           *CommissionRemark        `xml:"com:CommissionRemark,omitempty"`            //minOccurs="0"
	ConsolidatorRemark         *ConsolidatorRemark      `xml:"com:ConsolidatorRemark,omitempty"`          //minOccurs="0"
	InvoiceRemark              []*InvoiceRemark         `xml:"com:InvoiceRemark,omitempty"`               //minOccurs="0" maxOccurs="unbounded"
	SSR                        []*SSR                   `xml:"com:SSR,omitempty"`                         //minOccurs="0" maxOccurs="unbounded" //SSR element outside Booking Traveler without any Segment Ref or Booking Traveler Ref.
	EmailNotification          *EmailNotification       `xml:"com:EmailNotification,omitempty"`           //minOccurs="0" maxOccurs="1"
	QueuePlace                 *QueuePlace              `xml:"com:QueuePlace,omitempty"`                  //minOccurs="0" maxOccurs="1" //Allow queue placement of a PNR at the time of booking in AirCreateReservationReq,HotelCreateReservationReq,PassiveCreateReservationReq and VehicleCreateReservationReq for providers 1G,1V,1P and 1J.
	RuleName                   string                   `xml:"RuleName,attr,omitempty"`                   //use="optional" //his attribute is meant to attach a mandatory custom check rule name to a PNR. A non-mandatory custom check rule too can be attached to a PNR.
	UniversalRecordLocatorCode com.TypeLocatorCode      `xml:"UniversalRecordLocatorCode,attr,omitempty"` //use="optional" //Which UniversalRecord should this new reservation be applied to.  If blank, then a new one is created.
	ProviderLocatorCode        com.TypeLocatorCode      `xml:"ProviderLocatorCode,attr,omitempty"`        //use="optional" //Which Provider reservation does this reservation get added to.
	ProviderCode               string                   `xml:"ProviderCode,attr,omitempty"`               //use="optional" //To be used with ProviderLocatorCode, which host the reservation being added to belongs to.
	CustomerNumber             string                   `xml:"CustomerNumber,attr,omitempty"`             //use="optional" //Optional client centric customer identifier
	Version                    int                      `xml:"Version,attr,omitempty"`                    //use="optional"
}

//BaseCreateWithFormOfPaymentReq Container for BaseCreateReservation along with Form Of Payment
type BaseCreateWithFormOfPaymentReq struct {
	BaseCreateReservationReq
	FormOfPayment []*FormOfPayment `xml:"com:FormOfPayment,omitempty"` //minOccurs="0" maxOccurs="unbounded" //Provider:1G,1V,1P,1J,ACH,SDK
}
