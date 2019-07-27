package response

import (
	com "github.com/redochen/demos/travelport-uapi/soap/common"
)

//GetRoute 根据引用KEY获取航路信息
func (this *RouteList) GetRoute(routeRef com.TypeRef) *Route {
	if "" == string(routeRef) {
		return nil
	}

	if nil == this.Route || len(this.Route) == 0 {
		return nil
	}

	for _, route := range this.Route {
		if route.Key == routeRef {
			return route
		}
	}

	return nil
}

//GetLeg 根据引用KEY获取航路信息
func (this *RouteList) GetLeg(legRef com.TypeRef) *Leg {
	if nil == this.Route || len(this.Route) == 0 {
		return nil
	}

	for _, route := range this.Route {
		leg := route.GetLeg(legRef)
		if leg != nil {
			return leg
		}
	}

	return nil
}

//GetLeg 根据引用KEY获取航段信息
func (this *Route) GetLeg(legRef com.TypeRef) *Leg {
	if "" == string(legRef) {
		return nil
	}

	if nil == this.Leg || len(this.Leg) == 0 {
		return nil
	}

	for _, leg := range this.Leg {
		if leg.Key == legRef {
			return leg
		}
	}

	return nil
}

//GetFlightDetails 根据引用KEY获取航班详情
func (this *FlightDetailsList) GetFlightDetails(flightDetailsRef com.TypeRef) *FlightDetails {
	if "" == string(flightDetailsRef) {
		return nil
	}

	if nil == this.FlightDetails || len(this.FlightDetails) <= 0 {
		return nil
	}

	for _, details := range this.FlightDetails {
		if details.Key == flightDetailsRef {
			return details
		}
	}

	return nil
}

//GetAllFlightDetails 根据引用KEY列表获取航班详情列表
func (this *FlightDetailsList) GetAllFlightDetails(flightDetailsRefs []*FlightDetailsRef) []*FlightDetails {
	if nil == flightDetailsRefs || len(flightDetailsRefs) <= 0 {
		return nil
	}

	flightDetailsArray := make([]*FlightDetails, 0)

	for _, flightDetailsRef := range flightDetailsRefs {
		if nil == flightDetailsRef {
			continue
		}

		flightDetails := this.GetFlightDetails(flightDetailsRef.Key)
		if flightDetails != nil {
			flightDetailsArray = append(flightDetailsArray, flightDetails)
		}
	}

	return flightDetailsArray
}

//GetFlightDetailsOfSegment 获取航段的航班详情列表
func (this *FlightDetailsList) GetFlightDetailsOfSegment(segment *AirSegment) []*FlightDetails {
	if nil == segment || nil == segment.FlightDetailsRef ||
		len(segment.FlightDetailsRef) == 0 {
		return nil
	}

	flightDetails := make([]*FlightDetails, 0)

	for _, fdr := range segment.FlightDetailsRef {
		details := this.GetFlightDetails(fdr.Key)
		if nil == details {
			continue
		}

		flightDetails = append(flightDetails, details)
	}

	return flightDetails
}

//GetSegment 根据引用KEY获取航段信息
func (this *AirSegmentList) GetSegment(segmentRef com.TypeRef) *AirSegment {
	if "" == string(segmentRef) {
		return nil
	}

	if nil == this.AirSegment || len(this.AirSegment) <= 0 {
		return nil
	}

	for _, segment := range this.AirSegment {
		if segment.Key == segmentRef {
			return segment
		}
	}

	return nil
}

//GetFareInfo 根据引用KEY获取运价信息
func (this *FareInfoList) GetFareInfo(fareRef com.TypeRef) *FareInfo {
	if "" == string(fareRef) {
		return nil
	}

	if nil == this.FareInfo || len(this.FareInfo) <= 0 {
		return nil
	}

	for _, fare := range this.FareInfo {
		if fare.Key == fareRef {
			return fare
		}
	}

	return nil
}

//GetPrivateFareType 获取私有运价类型
func (this *FareInfoList) GetPrivateFareType(fareRef com.TypeRef) string {
	fare := this.GetFareInfo(fareRef)
	if fare != nil && len(fare.PrivateFare) > 0 {
		return fare.PrivateFare
	}

	return ""
}

//IsConnection 判断是否为Connection
func (this *Option) IsConnection(segmentIndex int) bool {
	if nil == this.Connection || len(this.Connection) <= 0 {
		return false
	}

	for _, op := range this.Connection {
		if op.SegmentIndex == segmentIndex {
			//根据以下文档的说明，StopOver=true的不视为Connection
			//https://support.travelport.com/webhelp/uapi/uAPI.htm#Air/Shared_Air_Topics/AirSegmentConnectionLogic.htm
			return !op.StopOver
		}
	}

	return false
}
