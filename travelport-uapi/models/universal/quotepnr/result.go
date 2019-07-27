package quotepnr

import (
	ap "github.com/redochen/demos/travelport-uapi/models/air/airprice"
	cp "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	rp "github.com/redochen/demos/travelport-uapi/models/universal/retrievepnr"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//QuotePnrResult 询价PNR结果
type QuotePnrResult struct {
	cp.CreatePnrResult
	AirPrices []*ap.PriceResult `xml:"airPrices" json:"airPrices"` //机票价格列表
}

//FillinByRetrievePnrResult 根据RetrievePnrResult填充结果
func (quotePnrResult *QuotePnrResult) FillinByRetrievePnrResult(result *rp.RetrievePnrResult) {
	if nil == result {
		return
	}

	if result.Status != 0 {
		quotePnrResult.Status = result.Status
		quotePnrResult.Message = result.Message
	}

	quotePnrResult.UniversalRecord = result.UniversalRecord
	quotePnrResult.Travelers = result.Travelers
	quotePnrResult.OtherServiceInfos = result.OtherServiceInfos
	quotePnrResult.ProviderInfos = result.ProviderInfos
	quotePnrResult.AirReservations = result.AirReservations
	quotePnrResult.HotelReservations = result.HotelReservations
	quotePnrResult.AgencyInfos = result.AgencyInfos
	quotePnrResult.Payments = result.Payments
}

//FillinByAirPriceResult 根据AirPriceResult填充结果
func (quotePnrResult *QuotePnrResult) FillinByAirPriceResult(result *ap.AirPriceResult) {
	if nil == result {
		return
	}

	if result.Status != 0 {
		quotePnrResult.Status = result.Status
		quotePnrResult.Message = result.Message
	}

	quotePnrResult.AirPrices = result.Prices
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *QuotePnrResult {
	result := &QuotePnrResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (quotePnrResult *QuotePnrResult) SetErrorCode(code ErrorCode) *QuotePnrResult {
	quotePnrResult.BaseResult.SetErrorCode(code)
	return quotePnrResult
}
