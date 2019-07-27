package response

import (
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
)

//Base hotel Search Response
type BaseHotelSearchRsp struct {
	comrs.BaseSearchRsp
	ReferencePoint       *HotelReferencePoint        `xml:"ReferencePoint,omitempty"`            //minOccurs="0" //Hotel reference point. Applicable for 1G,1V,1P,1J.
	HotelSearchResult    []*HotelSearchResult        `xml:"HotelSearchResult,omitempty"`         //minOccurs="0" maxOccurs="unbounded"
	MarketingInformation *comrs.MarketingInformation `xml:"MarketingInformation,omitempty"`      //minOccurs="0" maxOccurs="1"
	HostToken            *comrs.HostToken            `xml:"HostToken,omitempty"`                 //minOccurs="0"
	AddressSearchQuality int                         `xml:"AddressSearchQuality,attr,omitempty"` //use="optional" //Indicates the address matching level success for hotel address or Postal Code searches. Valid values: "1"-"8". Providers 1G, 1V.
}
