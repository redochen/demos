package quotepnr

import (
	. "github.com/redochen/demos/travelport-uapi/models/air"
	ap "github.com/redochen/demos/travelport-uapi/models/air/airprice"
	cp "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	. "github.com/redochen/demos/travelport-uapi/models/universal/quotepnr"
	rp "github.com/redochen/demos/travelport-uapi/models/universal/retrievepnr"
)

//getRetrievePnrParam 将QuotePnrParam转换为RetrievePnrParam
func getRetrievePnrParam(param *QuotePnrParam) *rp.RetrievePnrParam {
	if nil == param {
		return nil
	}

	result := &rp.RetrievePnrParam{
		UrlCode: param.UrlCode,
	}

	result.GdsAccount = param.GdsAccount
	result.ProviderCode = param.ProviderCode
	result.RandomText = param.RandomText

	return result
}

//getAirPriceParam 从RetrievePnrResult结果中获取AirPriceParam参数
func getAirPriceParam(result *rp.RetrievePnrResult, quote *QuotePnrParam) *ap.AirPriceParam {
	if nil == result {
		return nil
	}

	param := &ap.AirPriceParam{
		CheckInventory:     false, //不检查库存
		SpecifyBookingCode: true,  //指定舱位
	}

	//转换机票航段数据
	if result.AirReservations != nil && len(result.AirReservations) > 0 {
		param.Segments = getSegments(result.AirReservations[0])
	}

	if nil == param.Segments || len(param.Segments) == 0 {
		return nil
	}

	//转换乘机人数据
	if result.Travelers != nil && len(result.Travelers) > 0 {
		param.Passengers = getPassengerParams(result.Travelers)
	}

	//其他信息
	if quote != nil {
		param.Carrier = quote.Carrier
		param.Currency = quote.Currency
	}

	param.GdsAccount = quote.GdsAccount
	param.ProviderCode = quote.ProviderCode
	param.RandomText = quote.RandomText

	return param
}

//getSegments 从CreatePnr.AirReservationResult中获取AirPrice.SegmentParam
func getSegments(reservation *cp.AirReservationResult) []*Segment {
	if nil == reservation || nil == reservation.Segments ||
		len(reservation.Segments) == 0 {
		return nil
	}

	return reservation.Segments
}

//getPassengerParams 从CreatePnr.Traveler中获取AirPrice.PassengerParam
func getPassengerParams(travelers []*cp.Traveler) []*ap.PassengerParam {
	if nil == travelers || len(travelers) == 0 {
		return nil
	}

	passengers := make([]*ap.PassengerParam, 0)
	for _, t := range travelers {
		passenger := getPassengerParam(t)
		if passenger != nil {
			passengers = append(passengers, passenger)
		}
	}

	return passengers
}

//getPassengerParam 将CreatePnr.Traveler转换为AirPrice.PassengerParam
func getPassengerParam(traveler *cp.Traveler) *ap.PassengerParam {
	if nil == traveler {
		return nil
	}

	param := &ap.PassengerParam{
		PassengerType: traveler.PassengerType,
		ReferenceKey:  traveler.Key,
	}

	return param
}
