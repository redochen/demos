package request

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	uni "github.com/redochen/demos/travelport-uapi/soap/universal"
)

//Provider Reservation informations
type ProviderReservationInfo struct {
	ProviderReservationDetails            *ProviderReservationDetails            `xml:"univ:ProviderReservationDetails,omitempty"`            //minOccurs="0"
	ProviderReservationDisplayDetailsList *ProviderReservationDisplayDetailsList `xml:"univ:ProviderReservationDisplayDetailsList,omitempty"` //minOccurs="0"
	ExternalReservationInfo               *ExternalReservationInfo               `xml:"univ:ExternalReservationInfo,omitempty"`               //minOccurs="0" maxOccurs="1"
	Key                                   com.TypeRef                            `xml:"Key,attr"`                                             //use="required" //Key value of the provider reservation
	ProviderCode                          com.TypeProviderCode                   `xml:"ProviderCode,attr"`                                    //use="required" //Contains the Provider Code of the entity housing the actual reservation in the event this is a passive one.
	LocatorCode                           com.TypeProviderLocatorCode            `xml:"LocatorCode,attr"`                                     //use="required" //Contains the Locator Code of the actual reservation in the event this is a passive reservation.
	CreateDate                            string                                 `xml:"CreateDate,attr"`                                      //use="required" //The date and time that this reservation was created.
	HostCreateDate                        string                                 `xml:"HostCreateDate,attr,omitempty"`                        //type="xs:date" use="optional" //The date that this reservation was created in the host system.
	HostCreateTime                        string                                 `xml:"HostCreateTime,attr,omitempty"`                        //type="xs:time" use="optional" //The Time that this reservation was created in the host system for 1P and 1J.
	ModifiedDate                          string                                 `xml:"ModifiedDate,attr"`                                    //use="required" //The date and time that this reservation was last modified for any reason.
	Imported                              bool                                   `xml:"Imported,attr,omitempty"`                              //use="optional" //Identifies a reservation that was originally created elsewhere and imported into a Universal Record.
	TicketingModifiersRef                 com.TypeRef                            `xml:"TicketingModifiersRef,attr,omitempty"`                 //use="optional" //Reference to a Ticketing Modifers which is attached to this PNR. Ticketing Modifers referred  by this Key is a Primary Ticketing Modifers. Worldspan Primary DI line will be supported using this feature.
	InQueueMode                           bool                                   `xml:"InQueueMode,attr,omitempty"`                           //use="optional" //Identifies whether the gds pnr is being processed from the GDS queue.
	GroupRef                              com.TypeRef                            `xml:"GroupRef,attr,omitempty"`                              //use="optional" //Represents a traveler group for Group booking and all their accompanying data. SUPPORTED PROVIDER: Worldspan and JAL.
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
	DisplayDetails  []*DisplayDetails  `xml:"univ:DisplayDetails"`            //maxOccurs="unbounded"
	TravelerNameNum []*TravelerNameNum `xml:"univ:TravelerNameNum,omitempty"` //minOccurs="0" maxOccurs="unbounded"
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
	DisplayDetail   []*DisplayDetail     `xml:"univ:DisplayDetail,omitempty"`   //maxOccurs="unbounded" //xs:choice
	DisplayContents *uni.DisplayContents `xml:"univ:DisplayContents,omitempty"` //xs:choice
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
