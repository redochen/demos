package hotelavailex

import (
	avmodel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavailex"
	mimodel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelmedia"
)

//getHotelMediaParamFromHotelAvailResult 根据HotelAvail接口结果获取HotelMedia参数
func getHotelMediaParamFromHotelAvailResult(hotelAvailResult *avmodel.HotelAvailResult, param *HotelAvailExParam) *mimodel.HotelMediaParam {
	if nil == hotelAvailResult {
		return nil
	}

	hotelMediaParam := &mimodel.HotelMediaParam{}
	if param != nil {
		hotelMediaParam.GdsAccount = param.GdsAccount
		hotelMediaParam.ProviderCode = param.ProviderCode
		hotelMediaParam.RandomText = param.RandomText
	}

	if hotelAvailResult.Hotels != nil && len(hotelAvailResult.Hotels) > 0 {
		hotelMediaParam.HotelProperties = make([]*mimodel.HotelPropertyParam, 0)

		for _, h := range hotelAvailResult.Hotels {
			if nil == h.Properties || len(h.Properties) <= 0 {
				continue
			}

			for _, p := range h.Properties {
				param := &mimodel.HotelPropertyParam{
					Chain: p.Chain,
					Code:  p.Code,
				}

				hotelMediaParam.HotelProperties = append(hotelMediaParam.HotelProperties, param)
			}
		}
	}

	return hotelMediaParam
}
