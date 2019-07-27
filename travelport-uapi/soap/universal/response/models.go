package response

import (
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
	uni "github.com/redochen/demos/travelport-uapi/soap/universal"
)

//[RS] Universal Record holds one or more provider reservations
type UniversalRecord struct {
	LinkedUniversalRecord   []*comrs.LinkedUniversalRecord `xml:"LinkedUniversalRecord,omitempty"`   //minOccurs="0" maxOccurs="unbounded"
	Group                   *comrs.Group                   `xml:"Group,omitempty"`                   //minOccurs="0" maxOccurs="1"
	BookingTraveler         []*comrs.BookingTraveler       `xml:"BookingTraveler,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	ServiceFeeInfo          []*comrs.ServiceFeeInfo        `xml:"ServiceFeeInfo,omitempty"`          //minOccurs="0" maxOccurs="unbounded" //Travel Agency Service Fees (TASF) are charged by the agency through BSP or Airline Reporting Corporation (ARC). FOP will appear directly inside UniversalRecord
	OSI                     []*comrs.OSI                   `xml:"OSI,omitempty"`                     //minOccurs="0" maxOccurs="unbounded"
	ActionStatus            []*comrs.ActionStatus          `xml:"ActionStatus,omitempty"`            //minOccurs="0" maxOccurs="unbounded"
	ProviderReservationInfo []*ProviderReservationInfo     `xml:"ProviderReservationInfo,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	AirReservation          []*airrs.AirReservation        `xml:"AirReservation,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	HotelReservation        []*hotrs.HotelReservation      `xml:"HotelReservation,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	//VehicleReservation      []*vcl.VehicleReservation      `xml:"VehicleReservation,omitempty"`      //minOccurs="0" maxOccurs="unbounded"
	//PassiveReservation      []*psv.PassiveReservation      `xml:"PassiveReservation,omitempty"`      //minOccurs="0" maxOccurs="unbounded"
	//RailReservation         []*rail.RailReservation        `xml:"RailReservation,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	//CruiseReservation       []*crs.CruiseReservation       `xml:"CruiseReservation,omitempty"`       //minOccurs="0" maxOccurs="unbounded" //The parent container for all cruise booking data. Supported Providers :1V
	//EMDSummaryInfo      []*airrs.EMDSummaryInfo      `xml:"EMDSummaryInfo,omitempty"`      //minOccurs="0" maxOccurs="unbounded" //List of EMDs to be shown as part of UR. Supported providers are 1V/1G/1P/1J
	ProviderARNKSegment   []*comrs.ProviderARNKSegment `xml:"ProviderARNKSegment,omitempty"`       //minOccurs="0" maxOccurs="unbounded"
	SegmentContinuityInfo *SegmentContinuityInfo       `xml:"SegmentContinuityInfo,omitempty"`     //minOccurs="0" maxOccurs="1"
	XMLRemark             []*comrs.XMLRemark           `xml:"XMLRemark,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	GeneralRemark         []*comrs.GeneralRemark       `xml:"GeneralRemark,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	AccountingRemark      []*comrs.AccountingRemark    `xml:"AccountingRemark,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	UnassociatedRemark    []*comrs.UnassociatedRemark  `xml:"UnassociatedRemark,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	Postscript            []*comrs.Postscript          `xml:"Postscript,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	AgencyInfo            *comrs.AgencyInfo            `xml:"AgencyInfo,omitempty"`                //minOccurs="0"
	AppliedProfile        *comrs.AppliedProfile        `xml:"AppliedProfile,omitempty"`            //minOccurs="0"
	AgencyContactInfo     *comrs.AgencyContactInfo     `xml:"AgencyContactInfo,omitempty"`         //minOccurs="0"
	CustomerID            []*comrs.CustomerID          `xml:"CustomerID,omitempty"`                //minOccurs="0" maxOccurs="unbounded"
	CommissionRemark      []*comrs.CommissionRemark    `xml:"CommissionRemark,omitempty"`          //minOccurs="0" maxOccurs="unbounded"
	ConsolidatorRemark    []*comrs.ConsolidatorRemark  `xml:"ConsolidatorRemark,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	UnmaskedDataRemark    []*UnmaskedDataRemark        `xml:"UnmaskedDataRemark,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	InvoiceRemark         []*comrs.InvoiceRemark       `xml:"InvoiceRemark,omitempty"`             //minOccurs="0" maxOccurs="unbounded"
	ReviewBooking         []*comrs.ReviewBooking       `xml:"ReviewBooking,omitempty"`             //minOccurs="0" maxOccurs="unbounded" //Review Booking or Queue Minders is to add the reminders in the Provider Reservation along with the date time and Queue details. On the date time defined in reminders, the message along with the PNR goes to the desired Queue.
	SSR                   []*comrs.SSR                 `xml:"SSR,omitempty"`                       //minOccurs="0" maxOccurs="unbounded" //SSR's having no bookingTravelerRef need to add at providerReservation level outside bookingTraveler
	InvoiceData           []*comrs.InvoiceData         `xml:"InvoiceData,omitempty"`               //minOccurs="0" maxOccurs="unbounded"
	FormOfPayment         []*comrs.FormOfPayment       `xml:"FormOfPayment,omitempty"`             //minOccurs="0" maxOccurs="unbounded" //Provider: 1G,1V,1P,1J,ACH,SDK.Product:Air,Hotel,Vehicle,Cruise
	LocatorCode           com.TypeLocatorCode          `xml:"LocatorCode,attr"`                    //use="required" //Unique Identifier of a Universal Record. If this is ViewOnly UR then Locator Code is '999999'.
	SavedTripLocatorCode  com.TypeLocatorCode          `xml:"SavedTripLocatorCode,attr,omitempty"` //use="optional"
	LockReason            string                       `xml:"LockReason,attr,omitempty"`           //use="optional" //The reason for which the reservation is currently locked for modifications
	CreateDate            string                       `xml:"CreateDate,attr,omitempty"`           //use="optional" //The date and time that this reservation was created.
	Version               com.TypeURVersion            `xml:"Version,attr,omitempty"`              //use="optional"
	Status                string                       `xml:"Status,attr"`                         //use="required"
}

//Provider Reservation informations
type ProviderReservationInfo struct {
	ProviderReservationDetails            *ProviderReservationDetails            `xml:"ProviderReservationDetails,omitempty"`            //minOccurs="0"
	ProviderReservationDisplayDetailsList *ProviderReservationDisplayDetailsList `xml:"ProviderReservationDisplayDetailsList,omitempty"` //minOccurs="0"
	ExternalReservationInfo               *ExternalReservationInfo               `xml:"ExternalReservationInfo,omitempty"`               //minOccurs="0" maxOccurs="1"
	Key                                   com.TypeRef                            `xml:"Key,attr"`                                        //use="required" //Key value of the provider reservation
	ProviderCode                          com.TypeProviderCode                   `xml:"ProviderCode,attr"`                               //use="required" //Contains the Provider Code of the entity housing the actual reservation in the event this is a passive one.
	LocatorCode                           com.TypeProviderLocatorCode            `xml:"LocatorCode,attr"`                                //use="required" //Contains the Locator Code of the actual reservation in the event this is a passive reservation.
	CreateDate                            string                                 `xml:"CreateDate,attr"`                                 //use="required" //The date and time that this reservation was created.
	HostCreateDate                        string                                 `xml:"HostCreateDate,attr,omitempty"`                   //type="xs:date" use="optional" //The date that this reservation was created in the host system.
	HostCreateTime                        string                                 `xml:"HostCreateTime,attr,omitempty"`                   //type="xs:time" use="optional" //The Time that this reservation was created in the host system for 1P and 1J.
	ModifiedDate                          string                                 `xml:"ModifiedDate,attr"`                               //use="required" //The date and time that this reservation was last modified for any reason.
	Imported                              bool                                   `xml:"Imported,attr,omitempty"`                         //use="optional" //Identifies a reservation that was originally created elsewhere and imported into a Universal Record.
	TicketingModifiersRef                 com.TypeRef                            `xml:"TicketingModifiersRef,attr,omitempty"`            //use="optional" //Reference to a Ticketing Modifers which is attached to this PNR. Ticketing Modifers referred  by this Key is a Primary Ticketing Modifers. Worldspan Primary DI line will be supported using this feature.
	InQueueMode                           bool                                   `xml:"InQueueMode,attr,omitempty"`                      //use="optional" //Identifies whether the gds pnr is being processed from the GDS queue.
	GroupRef                              com.TypeRef                            `xml:"GroupRef,attr,omitempty"`                         //use="optional" //Represents a traveler group for Group booking and all their accompanying data. SUPPORTED PROVIDER: Worldspan and JAL.
	//<xs:attributeGroup attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"` //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr"`      //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
	OwningPCC com.TypePCC `xml:"OwningPCC,attr,omitempty"` //use="optional" //Indentifies the owning PCC of the PNR. PROVIDER SUPPORTED: Worldspan,JAL,Galileo and Apollo
	//<xs:attributeGroup name="attrAgentOverride">
	AgentOverride string `xml:"AgentOverride,attr,omitempty"` //use="optional" //AgentSine value that was used during PNR creation or End Transact.
	//</xs:attributeGroup>
}

//Indicates the type of content in PNR,to retrieve the display only content use ProviderReservationDisplayDetailsReq
type ProviderReservationDetails struct {
	//<xs:attributeGroup name="attrProviderReservationDetails">
	ProviderReservationDetail bool `xml:"ProviderReservationDetail,attr,omitempty"` //Provider Reservation data exists.
	CustomCheck               bool `xml:"CustomCheck,attr,omitempty"`               //Custom check data exists.
	ProviderProfile           bool `xml:"ProviderProfile,attr,omitempty"`           //Provider Profile data exists.
	DivideDetails             bool `xml:"DivideDetails,attr,omitempty"`             //Divide Details data exists.
	EnhancedItinModifiers     bool `xml:"EnhancedItinModifiers,attr,omitempty"`     //Enhanced itinerary modifiers data exists
	IntegratedContent         bool `xml:"IntegratedContent,attr,omitempty"`         //Integrated content data exists
	Cruise                    bool `xml:"Cruise,attr,omitempty"`                    //Cruise data exists.
	RailSegment               bool `xml:"RailSegment,attr,omitempty"`               //Rail Segment data exists.
	//</xs:attributeGroup>
}

//Response to display the addtional details of provider reservation information.
type ProviderReservationDisplayDetailsList struct {
	DisplayDetails  []*DisplayDetails  `xml:"DisplayDetails"`            //maxOccurs="unbounded"
	TravelerNameNum []*TravelerNameNum `xml:"TravelerNameNum,omitempty"` //minOccurs="0" maxOccurs="unbounded"
}

//Container for Host Passenger Details.Each TravelerNameNum represents a passenger present in the PNR. This is currently used by JAL(1J) only.
type TravelerNameNum struct {
	//<xs:attributeGroup name="attrBookingTravelerName">
	Prefix string                   `xml:"Prefix,attr,omitempty"` //use="optional" //Name prefix.
	First  string                   `xml:"First,attr"`            //use="required" //First Name.
	Middle string                   `xml:"Middle,attr,omitempty"` //use="optional" //Midle name.
	Last   com.TypeTravelerLastName `xml:"Last,attr"`             //use="required" //Last Name.
	Suffix string                   `xml:"Suffix,attr,omitempty"` //use="optional" //Name suffix.
	//</xs:attributeGroup>
	PaxNum string `xml:"PaxNum,attr"` //use="required" //Passenger Name Number in host. Should be in the format SurNameNum.GivenNameNum
}

//The container to display the contents of PNR in GDS format.
type DisplayDetails struct {
	DisplayDetail   []*DisplayDetail     `xml:"DisplayDetail,omitempty"`   //maxOccurs="unbounded" //xs:choice
	DisplayContents *uni.DisplayContents `xml:"DisplayContents,omitempty"` //xs:choice
	//<xs:attributeGroup name="attrProviderReservationDetails">
	ProviderReservationDetail bool `xml:"ProviderReservationDetail,attr,omitempty"` //Provider Reservation data exists.
	CustomCheck               bool `xml:"CustomCheck,attr,omitempty"`               //Custom check data exists.
	ProviderProfile           bool `xml:"ProviderProfile,attr,omitempty"`           //Provider Profile data exists.
	DivideDetails             bool `xml:"DivideDetails,attr,omitempty"`             //Divide Details data exists.
	EnhancedItinModifiers     bool `xml:"EnhancedItinModifiers,attr,omitempty"`     //Enhanced itinerary modifiers data exists
	IntegratedContent         bool `xml:"IntegratedContent,attr,omitempty"`         //Integrated content data exists
	Cruise                    bool `xml:"Cruise,attr,omitempty"`                    //Cruise data exists.
	RailSegment               bool `xml:"RailSegment,attr,omitempty"`               //Rail Segment data exists.
	//</xs:attributeGroup>
}

//Display the contents for requested MCO,Cruise etc. details
type DisplayDetail struct {
	Name  string `xml:"Name,attr"`  //use="required"
	Value string `xml:"Value,attr"` //use="required"
}

//Contains the details of the External PNR from which current PNR has neen copied. External PNR may reside in other gds system.This is a specific JAL functionality. PROVIDER SUPPORTED: JAL.
type ExternalReservationInfo struct {
	LocatorCode com.TypeLocatorCode `xml:"LocatorCode,attr"`       //use="required" //Locator Code of the External PNR.
	Carrier     com.TypeCarrier     `xml:"Carrier,attr,omitempty"` //use="optional" //Carrier associated with the External PNR.
}

//[RS] Status of the cancellation for this provider reservation.
type ProviderReservationStatus struct {
	CancelInfo   *comrs.ResultMessage        `xml:"CancelInfo,omitempty"` // minOccurs="0" //If the provider reservation was not successfully cancelled or cancelled with warnings the provider system might provides some textual information describing the reason.
	CreateDate   string                      `xml:"CreateDate,attr"`      //use="required" type="xs:dateTime" //The date and time that this reservation was created.
	ModifiedDate string                      `xml:"ModifiedDate,attr"`    //use="required" type="xs:dateTime" //The date and time that this reservation was last modified for any reason.
	ProviderCode com.TypeProviderCode        `xml:"ProviderCode,attr"`    //use="required" //Contains the Provider Code of the entity housing the actual reservation in the event this is a passive one.
	LocatorCode  com.TypeProviderLocatorCode `xml:"LocatorCode,attr"`     //use="required" //Contains the Locator Code of the actual reservation in the event this is a passive reservation.
	Cancelled    bool                        `xml:"Cancelled,attr"`       //use="required" //Will be true if the reservation was successfuly cancelled on the provider system.
}

//Security remark to a PNR .If a PCC is a Bridge-Branch/5-CP-Consolidator AND is also listed in this field,it can retrieve the PNR data UNMASKED and UNSUPPRESSED.PROVIDER SUPPORTED:JAL.
type UnmaskedDataRemark struct {
	PseudoCityCode             []*comrs.PseudoCityCode `xml:"PseudoCityCode"`                            //minOccurs="1" maxOccurs="11"
	Key                        com.TypeRef             `xml:"Key,attr,omitempty"`                        //use="optional" //Key to be used for internal processing.
	ProviderReservationInfoRef com.TypeRef             `xml:"ProviderReservationInfoRef,attr,omitempty"` //use="optional" //Provider reservation reference key.
	ProviderCode               com.TypeProviderCode    `xml:"ProviderCode,attr,omitempty"`               //use="optional" //Contains the Provider Code of the provider for which this element is used
	//<xs:attributeGroup attrElementKeyResults">
	ElStat      string `xml:"ElStat,attr,omitempty"` //use="optional" //This attribute is used to show the action results of an element. Possible values are "A" (when elements have been added to the UR) and "M" (when existing elements have been modified). Response only.
	KeyOverride bool   `xml:"KeyOverride,attr"`      //If a duplicate key is found where we are adding elements in some cases like URAdd, then instead of erroring out set this attribute to true.
	//</xs:attributeGroup>
}

//This container holds Arnks and segment continuity remarks
type SegmentContinuityInfo struct {
	ArvlUnknSegment            []*ArvlUnknSegment          `xml:"ArvlUnknSegment,omitempty"`                 //minOccurs="0" maxOccurs="unbounded"
	ContinuityOverrideRemark   []*ContinuityOverrideRemark `xml:"ContinuityOverrideRemark,omitempty"`        //minOccurs="0" maxOccurs="unbounded"
	ArrivalUnknownSegmentCount int                         `xml:"ArrivalUnknownSegmentCount,attr,omitempty"` //use="optional"
}

//An ARNK segment that identifies a missing travel information
type ArvlUnknSegment struct {
	BookingTravelerRef []*BookingTravelerRef `xml:"BookingTravelerRef,omitempty"` //minOccurs="0" maxOccurs="255"
	Key                com.TypeRef           `xml:"Key,attr"`                     //use="required"
	Origin             com.TypeIATACode      `xml:"Origin,attr,omitempty"`        //use="optional" //The IATA CITY code for this origination of this entity.
	Destination        com.TypeIATACode      `xml:"Destination,attr,omitempty"`   //use="optional" //The IATA CITY code for this destination of this entity.
	TravelOrder        int                   `xml:"TravelOrder,attr,omitempty"`   //use="optional" //To identify the appropriate travel sequence for Air/Car/Hotel segments/reservations based on travel dates. This ordering is applicable across the UR not provider or traveler specific
}

//Reference Element for Booking Traveler
type BookingTravelerRef struct {
	Key com.TypeRef `xml:"Key,attr"` //use="required"
}

//A textual remark container to hold any printable text. (max 512 chars)
type ContinuityOverrideRemark struct {
	Value    string `xml:",innerxml"`
	Category string `xml:"Category,attr,omitempty"` //use="optional" default="MCT" //This is remark category is always MCT. 'Minimum Connect Time'
}
