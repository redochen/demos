package airprice

import (
	"errors"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/air/airprice"
	. "github.com/redochen/demos/travelport-uapi/services/air"
	"github.com/redochen/demos/travelport-uapi/soap"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	"strings"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *AirPriceParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if nil == param.Segments || len(param.Segments) == 0 {
		return nil, errors.New("should contains at least one segment")
	}

	body := airproxy.NewAirPriceReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create AirPriceReqBody")
	}

	req := body.Request

	//返回结果中包含航班详情
	req.FareRuleType = "none"
	req.CheckFlightDetails = true

	//设置航段信息
	req.AirItinerary = &airrq.AirItinerary{
		AirSegment: make([]*airrq.AirSegment, 0),
	}

	pricingCmd := &airrq.AirPricingCommand{}

	for segIndex, segment := range param.Segments {
		airSegment := GetAirSegmentParam(segment, segIndex, param.ProviderCode, param.SpecifyBookingCode, pricingCmd)

		//航段转换失败，则返回错误
		if nil == airSegment {
			return nil, errors.New("failed to get segment param")
		}

		req.AirItinerary.AirSegment = append(req.AirItinerary.AirSegment, airSegment)
	}

	//设置AirPricingCommand
	if pricingCmd != nil {
		if pricingCmd.AirPricingModifiers != nil || pricingCmd.AirSegmentPricingModifiers != nil {

			//Currently only one AirPricingCommand is processed
			req.AirPricingCommand = make([]*airrq.AirPricingCommand, 0)
			req.AirPricingCommand = append(req.AirPricingCommand, pricingCmd)
		}
	}

	//设置AirPricingModifiers
	req.AirPricingModifiers = &airrq.AirPricingModifiers{
		FaresIndicator: "AllFares", //返回所有运价
	}

	//设置检查库存
	if param.CheckInventory {
		req.AirPricingModifiers.InventoryRequestType = "DirectAccess" //Seamless,Basic
	} else {
		req.AirPricingModifiers.InventoryRequestType = "Basic"
	}

	//不指定要预订的舱位，只指定舱位等级
	if !param.SpecifyBookingCode && param.CabinClass != "" {
		req.AirPricingModifiers.SetCabinClass(param.CabinClass)
	}

	//设置出票航司
	if len(param.Carrier) > 0 {
		req.AirPricingModifiers.PlatingCarrier = com.TypeCarrier(param.Carrier)
	}

	//设置货币代码
	if len(param.Currency) > 0 {
		req.AirPricingModifiers.CurrencyType = com.TypeCurrency(param.Currency)
	}

	//设置旅客信息
	if param.Passengers != nil && len(param.Passengers) > 0 {
		req.SearchPassenger = make([]*comrq.SearchPassenger, 0)
		for _, pax := range param.Passengers {
			if nil == pax {
				continue
			}

			passenger := &comrq.SearchPassenger{}
			passenger.PricePTCOnly = true
			passenger.BookingTravelerRef = pax.ReferenceKey

			//更改儿童类型为CNN+年龄
			if strings.EqualFold(pax.PassengerType, "CHD") || strings.EqualFold(pax.PassengerType, "CNN") {
				passenger.Code = "CNN"
				passenger.Age = 8 //儿童要带年龄
			} else {
				passenger.Code = com.TypePTC(pax.PassengerType)
			}

			req.SearchPassenger = append(req.SearchPassenger, passenger)
		}
	}

	return soap.NewReqEnvelope(body), nil
}

//getResult 解析结果
func getResult(body *airproxy.AirPriceRspBody) (*AirPriceResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("AirPriceRsp is nil")
	}

	result := &AirPriceResult{
		Segments: make([]*Segment, 0),
		Prices:   make([]*PriceResult, 0),
	}

	//解析航段信息
	if rsp.AirItinerary != nil && rsp.AirItinerary.AirSegment != nil &&
		len(rsp.AirItinerary.AirSegment) > 0 {
		for idx, airSegment := range rsp.AirItinerary.AirSegment {
			segment := GetSegmentResult(airSegment, idx, false, "", nil)
			if segment != nil {
				result.Segments = append(result.Segments, segment)
			}
		}
	}

	//解析价格信息
	if rsp.AirPriceResult != nil && len(rsp.AirPriceResult) > 0 {
		for _, apr := range rsp.AirPriceResult {
			price := getPriceResult(apr)
			if price != nil {
				result.Prices = append(result.Prices, price)
			}
		}
	}

	return result, nil
}

//将AirPriceResult转换为PriceResult
func getPriceResult(price *airrs.AirPriceResult) *PriceResult {
	if nil == price {
		return nil
	}

	result := &PriceResult{}

	if price.AirPricingSolution != nil && len(price.AirPricingSolution) > 0 {
		result.AirPricingSolutions = make([]*AirPricingSolution, 0)
		for _, aps := range price.AirPricingSolution {
			solution := getAirPricingSolution(aps, nil)
			if solution != nil {
				result.AirPricingSolutions = append(result.AirPricingSolutions, solution)
			}
		}
	}

	return result
}

//转换AirPricingSolution
func getAirPricingSolution(solution *airrs.AirPricingSolution, fareInfos *airrs.FareInfoList) *AirPricingSolution {
	if nil == solution {
		return nil
	}

	result := &AirPricingSolution{
		Key:               string(solution.Key),
		TotalPrice:        solution.TotalPrice.GetFloatAmount(),
		TotalCurrency:     solution.TotalPrice.GetCurrency(),
		AppTotalPrice:     solution.ApproximateTotalPrice.GetFloatAmount(),
		AppTotalCurrency:  solution.ApproximateTotalPrice.GetCurrency(),
		BasePrice:         solution.BasePrice.GetFloatAmount(),
		BaseCurrency:      solution.BasePrice.GetCurrency(),
		AppBasePrice:      solution.ApproximateBasePrice.GetFloatAmount(),
		AppBaseCurrency:   solution.ApproximateBasePrice.GetCurrency(),
		EquivBasePrice:    solution.EquivalentBasePrice.GetFloatAmount(),
		EquivBaseCurrency: solution.EquivalentBasePrice.GetCurrency(),
		Taxes:             solution.Taxes.GetFloatAmount(),
		TaxesCurrency:     solution.Taxes.GetCurrency(),
		AppTaxes:          solution.ApproximateTaxes.GetFloatAmount(),
		AppTaxesCurrency:  solution.ApproximateTaxes.GetCurrency(),
		QuoteDate:         solution.QuoteDate,
	}

	if solution.AirPricingInfo != nil && len(solution.AirPricingInfo) > 0 {
		result.PricingInfos = make([]*PricingInfo, 0)
		for _, info := range solution.AirPricingInfo {
			pricingInfo := GetPricingInfo(info, fareInfos, string(solution.Key))
			if pricingInfo != nil {
				result.PricingInfos = append(result.PricingInfos, pricingInfo)
			}
		}
	}

	return result
}

/*
解析错误
*/
func pasreFault(fault *airproxy.AirPriceRspFault) string {
	if nil == fault {
		return "unkown error"
	}

	/*
		if fault.Error != nil {
			if fault.Error.AirSegmentError != nil &&
				len(fault.Error.AirSegmentError) > 0 {

				var errorMsg string

				for _, segError := range fault.Error.AirSegmentError {
					if segError.AirSegment != nil {
						return fmt.Sprintf("%s%d %s",
							segError.AirSegment.Carrier,
							int(segError.AirSegment.FlightNumber),
							segError.ErrorMessage)
					} else if len(errorMsg) <= 0 &&
						len(segError.ErrorMessage) > 0 {
						errorMsg = segError.ErrorMessage
						break
					}
				}

				if len(errorMsg) > 0 {
					return errorMsg
				}

			} else if len(fault.Error.Description) > 0 {
				return fault.Error.Description
			}
		}
	*/

	//return fault.Error.Description
	return fmt.Sprintf("%s(%s)", fault.Code, fault.Text)
}
