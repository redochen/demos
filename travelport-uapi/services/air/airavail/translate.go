package airavail

import (
	"errors"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/air/airavail"
	. "github.com/redochen/demos/travelport-uapi/services/air"
	"github.com/redochen/demos/travelport-uapi/soap"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	. "github.com/redochen/tools/string"
	. "github.com/redochen/tools/time"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *AirAvailParam, nextReqResultReferences []*comrq.NextResultReference, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if nil == param.Routes || len(param.Routes) == 0 {
		return nil, errors.New("should contains at least one route")
	}

	body := airproxy.NewAirAvailReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create AirAvailReqBody")
	}

	req := body.Request

	//前一次查询结果引用
	body.Request.NextResultReference = nextReqResultReferences

	req.SearchSpecificAirSegment = getSearchSpecificAirSegments(param.Segments)
	if nil == req.SearchSpecificAirSegment || len(req.SearchSpecificAirSegment) <= 0 {
		req.SearchAirLeg = getSearchAirLegs(param.Routes)
		if nil == req.SearchAirLeg || len(req.SearchAirLeg) <= 0 {
			return nil, errors.New("should contains at least one segment or one route")
		}
	}

	//获取查询选项
	req.AirSearchModifiers = GetAirSearchModifiers(param.Modifiers, param.ProviderCode)

	return soap.NewReqEnvelope(body), nil
}

//getSearchAirLegs 获取查询行程列表
func getSearchAirLegs(routes []*Route) []*airrq.SearchAirLeg {
	if nil == routes || len(routes) <= 0 {
		return nil
	}

	airLegs := make([]*airrq.SearchAirLeg, 0)

	for _, route := range routes {
		leg := &airrq.SearchAirLeg{
			SearchOrigin:      make([]*comrq.SearchLocation, 0),
			SearchDestination: make([]*comrq.SearchLocation, 0),
			SearchDepTime:     make([]*comrq.FlexibleTimeSpec, 0),
		}

		//出发城市或机场
		original := &comrq.SearchLocation{
			CityOrAirport: &comrq.CityOrAirport{
				Code:       route.Origin,
				PreferCity: true,
			},
		}
		leg.SearchOrigin = append(leg.SearchOrigin, original)

		//到达城市或机场
		destination := &comrq.SearchLocation{
			CityOrAirport: &comrq.CityOrAirport{
				Code:       route.Destination,
				PreferCity: true,
			},
		}
		leg.SearchDestination = append(leg.SearchDestination, destination)

		//日期时间
		date := &comrq.FlexibleTimeSpec{}
		date.PreferredTime = CcTime.AddDateSeparator(route.DepartureDate, "-", true)
		leg.SearchDepTime = append(leg.SearchDepTime, date)

		airLegs = append(airLegs, leg)
	}

	return airLegs
}

//getSearchSpecificAirSegments 获取查询航段列表
func getSearchSpecificAirSegments(segments []*Segment) []*airrq.SearchSpecificAirSegment {
	if nil == segments || len(segments) <= 0 {
		return nil
	}

	airSegments := make([]*airrq.SearchSpecificAirSegment, 0)

	for _, segment := range segments {
		segment := &airrq.SearchSpecificAirSegment{
			SegmentIndex:  segment.Index,
			Carrier:       segment.Carrier,
			FlightNumber:  CcStr.FormatInt(segment.FlightNumber),
			Origin:        segment.DepartureAirport,
			Destination:   segment.ArrivalAirport,
			DepartureTime: segment.DepartureTime,
		}

		airSegments = append(airSegments, segment)
	}

	return airSegments
}

//getResult 解析结果
func getResult(body *airproxy.AirAvailRspBody) (*AirAvailResult, []*comrs.NextResultReference, error) {
	if nil == body || nil == body.Response {
		return nil, nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, nil, errors.New("AirAvailRsp is nil")
	}

	if nil == rsp.FlightDetailsList {
		return nil, nil, errors.New("flight details list is nil")
	}

	if nil == rsp.AirSegmentList {
		return nil, nil, errors.New("segment list is nil")
	}

	if nil == rsp.AirItinerarySolution || len(rsp.AirItinerarySolution) == 0 {
		return nil, nil, errors.New("air itinerary solution is empty")
	}

	result := &AirAvailResult{
		Flights: make([]*Flight, 0),
	}

	for _, solution := range rsp.AirItinerarySolution {
		if nil == solution || nil == solution.AirSegmentRef || len(solution.AirSegmentRef) == 0 {
			continue
		}

		var lastIndex int
		var segmentCount int
		var lastFlight *Flight

		//遍历所有航段引用
		for segIdx, asr := range solution.AirSegmentRef {
			if nil == asr {
				continue
			}

			//当前行程的航段计数
			segmentCount++

			var curFlight *Flight

			//有前延航段
			if isConnection(solution.Connection, segIdx-1, false) {
				lastIndex++
				curFlight = lastFlight
			}

			if nil == curFlight {
				lastIndex = 0
				curFlight = &Flight{
					Segments: make([]*Segment, 0),
				}
			}

			//这里可能会有问题，有待观察 2016.03.02
			isConn := isConnection(solution.Connection, segIdx, true)

			//解析航段信息
			airSegment := rsp.AirSegmentList.GetSegment(asr.Key)
			if nil == airSegment {
				continue
			}

			segment := GetSegmentResult(airSegment, lastIndex, isConn, string(solution.Key), rsp.FlightDetailsList)
			if nil == segment {
				continue
			}

			curFlight.Segments = append(curFlight.Segments, segment)

			//有后延航段
			if isConnection(solution.Connection, segIdx, false) {
				lastFlight = curFlight
			} else {
				//行程中所有航段都有余座时才添加到列表中
				if curFlight != nil && curFlight.Segments != nil &&
					len(curFlight.Segments) == segmentCount {
					result.Flights = append(result.Flights, curFlight)
				}

				//重置变量及计数
				lastIndex = 0
				segmentCount = 0
				lastFlight = nil
			}
		}
	}

	return result, rsp.NextResultReference, nil
}

//isConnection 是否为连接点
func isConnection(connections []*airrs.Connection, index int, checkStopOver bool) bool {
	if nil == connections || len(connections) == 0 {
		return false
	}

	for _, conn := range connections {
		if conn.SegmentIndex == index {
			if !checkStopOver {
				return true
			}

			//根据以下文档的说明，StopOver=true的不视为Connection
			//https://support.travelport.com/webhelp/uapi/uAPI.htm#Air/Shared_Air_Topics/AirSegmentConnectionLogic.htm
			return !conn.StopOver
		}
	}

	return false
}
