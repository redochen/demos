package request

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
)

//Base hotel Search Request
type BaseHotelSearchReq struct {
	comrq.BaseSearchReq
	HotelSearchLocation  *HotelSearchLocation    `xml:"hot:HotelSearchLocation,omitempty"`  //minOccurs="0"
	HotelSearchModifiers *HotelSearchModifiers   `xml:"hot:HotelSearchModifiers,omitempty"` //minOccurs="0"
	HotelStay            *HotelStay              `xml:"hot:HotelStay"`
	PointOfSale          *comrq.PointOfSale      `xml:"com:PointOfSale,omitempty"`      //minOccurs="0"
	PolicyReference      com.TypePolicyReference `xml:"PolicyReference,attr,omitempty"` //use="optional" //This attribute will be used to pass in a value on the request which would be used to link to a ‘Policy Group’ in a policy engine external to UAPI.
}

//Base request for all hotel details search request.
type BaseHotelDetailsReq struct {
	comrq.BaseReq
	HotelProperty         *HotelProperty         `xml:"hot:HotelProperty"`
	HotelDetailsModifiers *HotelDetailsModifiers `xml:"hot:HotelDetailsModifiers,omitempty"` //minOccurs="0"
	PointOfSale           *comrq.PointOfSale     `xml:"com:PointOfSale,omitempty"`           //minOccurs="0"
}
