package response

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
)

//Availability Search response
type BaseAvailabilitySearchRsp struct {
	comrs.BaseSearchRsp
	FlightDetailsList    *FlightDetailsList      `xml:"FlightDetailsList,omitempty"`    //minOccurs="0"
	AirSegmentList       *AirSegmentList         `xml:"AirSegmentList,omitempty"`       //minOccurs="0"
	FareInfoList         *FareInfoList           `xml:"FareInfoList,omitempty"`         //minOccurs="0"
	FareRemarkList       *FareRemarkList         `xml:"FareRemarkList,omitempty"`       //minOccurs="0"
	AirItinerarySolution []*AirItinerarySolution `xml:"AirItinerarySolution,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	HostTokenList        *HostTokenList          `xml:"HostTokenList,omitempty"`        //minOccurs="0" maxOccurs="1"
	APISRequirementsList *APISRequirementsList   `xml:"APISRequirementsList,omitempty"` //minOccurs="0"
	DistanceUnits        com.TypeDistance        `xml:"DistanceUnits,attr,omitempty"`   //use="optional"
}

//Base Response for Air Search
type AirSearchRsp struct {
	BaseAvailabilitySearchRsp
	FareNoteList                  *FareNoteList                  `xml:"FareNoteList,omitempty"`                  //minOccurs="0"
	ExpertSolutionList            *ExpertSolutionList            `xml:"ExpertSolutionList,omitempty"`            //minOccurs="0"
	RouteList                     *RouteList                     `xml:"RouteList,omitempty"`                     //minOccurs="0"
	AlternateRouteList            *AlternateRouteList            `xml:"AlternateRouteList,omitempty"`            //minOccurs="0"
	AlternateLocationDistanceList *AlternateLocationDistanceList `xml:"AlternateLocationDistanceList,omitempty"` //minOccurs="0"
	FareInfoMessage               *FareInfoMessage               `xml:"FareInfoMessage,omitempty"`               //minOccurs="0" maxOccurs="99"
	//<xs:choice minOccurs="0">
	AirPricingSolution []*AirPricingSolution `xml:"AirPricingSolution,omitempty"` //minOccurs="0" maxOccurs="unbounded"
	AirPricePointList  *AirPricePointList    `xml:"AirPricePointList,omitempty"`  //minOccurs="0"
	//</xs:choice>
	//<xs:element ref="rail:RailSegmentList" minOccurs="0"/>
	//<xs:element ref="rail:RailJourneyList" minOccurs="0"/>
	//<xs:element ref="rail:RailFareNoteList" minOccurs="0"/>
	//<xs:element ref="rail:RailFareIDList" minOccurs="0"/>
	//<xs:element ref="rail:RailFareList" minOccurs="0"/>
	//<xs:element ref="rail:RailPricingSolution" minOccurs="0" maxOccurs="unbounded"/>
}
