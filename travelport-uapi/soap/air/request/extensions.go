package request

import (
	//"strings"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	"github.com/redochen/demos/travelport-uapi/util"
)

//SetClassOfService 设置预订舱位
func (airPricingCommand *AirPricingCommand) SetClassOfService(segmentRefKey com.TypeRef, classOfService string) {
	bookingCode := &BookingCode{
		Code: classOfService,
	}

	modifier := &AirSegmentPricingModifiers{
		AirSegmentRef: segmentRefKey,
		PermittedBookingCodes: &PermittedBookingCodes{
			BookingCode: make([]*BookingCode, 0),
		},
	}
	modifier.PermittedBookingCodes.BookingCode = append(modifier.PermittedBookingCodes.BookingCode, bookingCode)

	if nil == airPricingCommand.AirSegmentPricingModifiers {
		airPricingCommand.AirSegmentPricingModifiers = make([]*AirSegmentPricingModifiers, 0)
	}

	airPricingCommand.AirSegmentPricingModifiers = append(airPricingCommand.AirSegmentPricingModifiers, modifier)
}

//SetCabinClass 设置舱位等级
func (airPricingModifiers *AirPricingModifiers) SetCabinClass(cabinClassLetter string) {
	airPricingModifiers.PermittedCabins = &PermittedCabins{
		CabinClass: make([]*comrq.CabinClass, 0),
	}

	cabinClass := &comrq.CabinClass{
		Type: util.GetCabinClass(cabinClassLetter),
	}

	airPricingModifiers.PermittedCabins.CabinClass = append(
		airPricingModifiers.PermittedCabins.CabinClass,
		cabinClass)
}

//SetCabinClass 设置舱位等级
func (airLegModifiers *AirLegModifiers) SetCabinClass(cabinClassLetter string) {
	airLegModifiers.PermittedCabins = &PermittedCabins{
		CabinClass: make([]*comrq.CabinClass, 0),
	}

	//在某些情况下会有导致数据查询不出来
	/*
		if strings.EqualFold(cabinClassLetter, "Y") {
			cabinClassEx := &comrq.CabinClass{
				Type: util.GetCabinClass("S"), //顺带查询超级经济舱
			}

			this.PermittedCabins.CabinClass = append(
				this.PermittedCabins.CabinClass,
				cabinClassEx)
		}
	*/

	cabinClass := &comrq.CabinClass{
		Type: util.GetCabinClass(cabinClassLetter),
	}

	airLegModifiers.PermittedCabins.CabinClass = append(
		airLegModifiers.PermittedCabins.CabinClass,
		cabinClass)
}

//IsConnection 判断是否为Connection
func (option *Option) IsConnection(segmentIndex int) bool {
	if nil == option.Connection || len(option.Connection) <= 0 {
		return false
	}

	for _, op := range option.Connection {
		if op.SegmentIndex == segmentIndex {
			//根据以下文档的说明，StopOver=true的不视为Connection
			//https://support.travelport.com/webhelp/uapi/uAPI.htm#Air/Shared_Air_Topics/AirSegmentConnectionLogic.htm
			return !op.StopOver
		}
	}

	return false
}
