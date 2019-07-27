package response

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
)

//The additional titles associated to the brand or optional service. Providers: ACH, RCH, 1G, 1V, 1P, 1J.
type Title TextElement

//Type of Text, Eg-'Upsell','Marketing Agent','Marketing Consumer','Strapline','Rule'.
type Text TextElement

//The tax information for a
type TaxInfo comrs.TaxInfo

//A generic type of fee for those charges which are incurred by the passenger, but not necessarily shown on tickets
type FeeInfo comrs.FeeInfo

//An Air marketable travel segment.
type AirSegment BaseAirSegment

//Reference to optional service
type OptionalServiceRef com.TypeRef

//Information related to Embargo
type EmbargoInfo BaseBaggageAllowanceInfo

//The unique identifier of the brand
type TypeBrandId string
