package lowfare

import (
	"errors"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/air/lowfare"
	. "github.com/redochen/demos/travelport-uapi/services/air"
	"github.com/redochen/demos/travelport-uapi/soap"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	. "github.com/redochen/tools/time"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *LowFareParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if nil == param.Routes || len(param.Routes) == 0 {
		return nil, errors.New("should contains at least one route")
	}

	body := airproxy.NewLowFareReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create LowFareReqBody")
	}

	req := body.Request

	//结果返回形式。True - AirPricingSolution; False - AirPricingPointList
	req.SolutionResult = false

	//要求返回退改签
	req.FareRulesFilterCategory = &airrq.FareRulesFilterCategory{
		CategoryCode: make([]string, 0),
	}

	//目前只支持"CHG"
	req.FareRulesFilterCategory.CategoryCode = append(req.FareRulesFilterCategory.CategoryCode, "CHG")

	//设置航路信息
	req.SearchAirLeg = make([]*airrq.SearchAirLeg, 0)
	for _, r := range param.Routes {
		leg := &airrq.SearchAirLeg{
			SearchOrigin:      make([]*comrq.SearchLocation, 0),
			SearchDestination: make([]*comrq.SearchLocation, 0),
			SearchDepTime:     make([]*comrq.FlexibleTimeSpec, 0),
			AirLegModifiers:   &airrq.AirLegModifiers{},
		}

		//出发城市或机场
		org := &comrq.SearchLocation{
			CityOrAirport: &comrq.CityOrAirport{
				Code:       r.Origin,
				PreferCity: true,
			},
		}
		leg.SearchOrigin = append(leg.SearchOrigin, org)

		//到达城市或机场
		dst := &comrq.SearchLocation{
			CityOrAirport: &comrq.CityOrAirport{
				Code:       r.Destination,
				PreferCity: true,
			},
		}
		leg.SearchDestination = append(leg.SearchDestination, dst)

		//日期时间
		dep := &comrq.FlexibleTimeSpec{}
		dep.PreferredTime = CcTime.AddDateSeparator(r.DepartureDate, "-", true)
		leg.SearchDepTime = append(leg.SearchDepTime, dep)

		//舱位等级
		leg.AirLegModifiers.SetCabinClass(param.CabinClass)

		req.SearchAirLeg = append(req.SearchAirLeg, leg)
	}

	req.SearchPassenger = make([]*comrq.SearchPassenger, 0)

	//设置查询选项
	if param.MetaOption != "" {
		req.MetaOptionIdentifier = "99" // 目前有bug，只能传99 //param.MetaOption
	}

	//设置旅客类型
	for adult := 0; adult < param.AdultCount; adult++ {
		passenger := &comrq.SearchPassenger{}
		passenger.Code = "ADT"
		passenger.PricePTCOnly = true
		req.SearchPassenger = append(req.SearchPassenger, passenger)
	}

	for child := 0; child < param.ChildCount; child++ {
		passenger := &comrq.SearchPassenger{}
		passenger.Code = "CNN"
		passenger.Age = 8 //儿童要带年龄
		passenger.PricePTCOnly = true
		req.SearchPassenger = append(req.SearchPassenger, passenger)
	}

	//获取查询选项
	req.AirSearchModifiers = GetAirSearchModifiers(param.Modifiers, param.ProviderCode)

	//设置最多返回结果数
	if param.MaxSolutions > 0 {
		req.AirSearchModifiers.MaxSolutions = param.MaxSolutions
	}

	//获取价格选项
	req.AirPricingModifiers = getAirPricingModifiers(param)

	return soap.NewReqEnvelope(body), nil
}

//getAirPricingModifiers 获取AirPricingModifiers
func getAirPricingModifiers(param *LowFareParam) *airrq.AirPricingModifiers {
	if nil == param {
		return nil
	}

	modifiers := &airrq.AirPricingModifiers{
		ETicketability:       "Required", //Yes, No, Required, Ticketless
		AccountCodeFaresOnly: false,
	}

	//设置货币代码
	if len(param.Currency) > 0 {
		modifiers.CurrencyType = com.TypeCurrency(param.Currency)
	}

	//设置开票航司
	if len(param.PlatingCarrier) > 0 {
		modifiers.PlatingCarrier = com.TypeCarrier(param.PlatingCarrier)
	}

	return modifiers
}

//getResult 解析结果
func getResult(body *airproxy.LowFareRspBody) (*LowFareResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("LowFareSearchRsp is nil")
	}

	if nil == rsp.AirPricePointList || nil == rsp.AirPricePointList.AirPricePoint ||
		len(rsp.AirPricePointList.AirPricePoint) <= 0 {
		return nil, errors.New("AirPricePointList is nil")
	}

	result := &LowFareResult{
		Fares: make([]*FlightFare, 0),
	}

	for _, point := range rsp.AirPricePointList.AirPricePoint {
		if nil == point || nil == point.AirPricingInfo || len(point.AirPricingInfo) == 0 {
			continue
		}

		flightFare := &FlightFare{
			Prices:   make([]*PricingInfo, 0),
			Segments: make([]*Segment, 0),
		}

		for _, info := range point.AirPricingInfo {
			segments := getSegments(info, rsp.RouteList, rsp.AirSegmentList, rsp.FlightDetailsList)
			flightFare.Segments = segments

			pricingInfo := GetPricingInfo(info, rsp.FareInfoList, string(point.Key))
			if nil == pricingInfo {
				continue
			}

			flightFare.Prices = append(flightFare.Prices, pricingInfo)
		}

		result.Fares = append(result.Fares, flightFare)
	}

	return result, nil
}

//getSegments 获取航段列表
func getSegments(info *airrs.AirPricingInfo, routes *airrs.RouteList, segments *airrs.AirSegmentList, details *airrs.FlightDetailsList) []*Segment {
	if nil == info || nil == info.FlightOptionsList ||
		nil == info.FlightOptionsList.FlightOption {
		return nil
	}

	if nil == routes || nil == segments {
		return nil
	}

	list := make([]*Segment, 0)

	for _, flight := range info.FlightOptionsList.FlightOption {
		if nil == flight || nil == flight.Option ||
			len(flight.Option) == 0 {
			continue
		}

		leg := routes.GetLeg(flight.LegRef)
		if nil == leg {
			continue
		}

		for _, option := range flight.Option {
			if nil == option.BookingInfo || len(option.BookingInfo) == 0 {
				continue
			}

			for segIndex, booking := range option.BookingInfo {
				airSegment := segments.GetSegment(booking.SegmentRef)
				if nil == airSegment {
					continue
				}

				isConn := option.IsConnection(segIndex)
				segment := GetSegmentResult(airSegment, segIndex, isConn, string(option.Key), details)
				if nil == segment {
					continue
				}

				segment.BookingCode = booking.BookingCode
				segment.BookingCount = booking.BookingCount

				if len(segment.CabinClass) <= 0 {
					segment.CabinClass = booking.CabinClass
				}

				list = append(list, segment)
			}
		}
	}

	return list
}
