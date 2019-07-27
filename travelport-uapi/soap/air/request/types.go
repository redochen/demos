package request

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
)

//An Air marketable travel segment.
type AirSegment BaseAirSegment

//The additional titles associated to the brand or optional service. Providers: ACH, RCH, 1G, 1V, 1P, 1J.
type Title TextElement

//Type of Text, Eg-'Upsell','Marketing Agent','Marketing Consumer','Strapline','Rule'.
type Text TextElement

//The tax information for a
type TaxInfo comrq.TaxInfo

//A generic type of fee for those charges which are incurred by the passenger, but not necessarily shown on tickets
type FeeInfo comrq.FeeInfo

//Reference to optional service
type OptionalServiceRef com.TypeRef

//The unique identifier of the brand
type TypeBrandId string

//Information related to Embargo
type EmbargoInfo BaseBaggageAllowanceInfo
