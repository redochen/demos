package hotelrules

import (
	"errors"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelrules"
	hotdetsvc "github.com/redochen/demos/travelport-uapi/services/hotel/hoteldetails"
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	. "github.com/redochen/tools/time"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *HotelRulesParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	body := hotproxy.NewHotelRulesReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create NewHotelRulesReqBody")
	}

	req := body.Request

	if len(param.UrlCode) > 0 {
		//PNR编码
		req.HotelReservationLocatorCode = com.TypeLocatorCode(param.UrlCode)
	} else {
		req.HotelRulesLookup = translateHotelRulesLookupParam(param)
	}

	return soap.NewReqEnvelope(body), nil
}

//translateHotelRulesLookupParam 转换查询参数
func translateHotelRulesLookupParam(param *HotelRulesParam) *hotrq.HotelRulesLookup {
	if nil == param {
		return nil
	}

	lookup := &hotrq.HotelRulesLookup{
		Base:         com.TypeMoney(fmt.Sprintf("%s%.2f", param.Currency, param.BaseAmount)),
		RatePlanType: param.RatePlanType,
	}

	//设置酒店属性
	lookup.HotelProperty = &hotrq.HotelProperty{
		HotelChain: com.TypeHotelChainCode(param.Chain),
		HotelCode:  com.TypeHotelCode(param.Code),
		Name:       param.Name,
	}

	//设置入住时间
	lookup.HotelStay = &hotrq.HotelStay{
		CheckinDate:  hot.TypeDate(CcTime.AddDateSeparator(param.CheckinDate, "-", true)),
		CheckoutDate: hot.TypeDate(CcTime.AddDateSeparator(param.CheckoutDate, "-", true)),
	}

	//设置请求类型
	if param.RequestType == 1 {
		lookup.RulesDetailReqd = "Details"
	} else if param.RequestType == 2 {
		lookup.RulesDetailReqd = "Rules"
	} else {
		lookup.RulesDetailReqd = "All"
	}

	//设置供应商
	if len(param.Provider) > 0 {
		if nil == lookup.HotelRulesModifiers {
			lookup.HotelRulesModifiers = &hotrq.HotelRulesModifiers{}
		}
		lookup.HotelRulesModifiers.PermittedProviders = &comrq.PermittedProviders{
			Provider: &comrq.Provider{
				Code: param.Provider,
			},
		}
	}

	//设置成人数量
	if param.Adults > 0 {
		if nil == lookup.HotelRulesModifiers {
			lookup.HotelRulesModifiers = &hotrq.HotelRulesModifiers{}
		}
		lookup.HotelRulesModifiers.NumberOfAdults = param.Adults
	}

	//设置儿童数量
	if param.Children > 0 {
		if nil == lookup.HotelRulesModifiers {
			lookup.HotelRulesModifiers = &hotrq.HotelRulesModifiers{}
		}
		lookup.HotelRulesModifiers.NumberOfChildren = &hotrq.NumberOfChildren{
			Count: param.Children,
		}
	}

	//设置房间数量
	if param.Rooms > 0 {
		if nil == lookup.HotelRulesModifiers {
			lookup.HotelRulesModifiers = &hotrq.HotelRulesModifiers{}
		}
		lookup.HotelRulesModifiers.NumberOfRooms = param.Rooms
	}

	return lookup
}

//getResult 转换结果
func getResult(body *hotproxy.HotelRulesRspBody) (*HotelRulesResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("HotelRulesRsp is empty")
	}

	result := &HotelRulesResult{}

	//解析HotelRateDetail
	if rsp.HotelRateDetail != nil && len(rsp.HotelRateDetail) > 0 {
		result.Rates = make([]*HotelRateDetail, 0)
		for _, hrd := range rsp.HotelRateDetail {
			rate := hotdetsvc.TranslateHotelRateDetailResult(hrd)
			if rate != nil {
				result.Rates = append(result.Rates, rate)
			}
		}
	}

	//解析HotelRuleItem
	if rsp.HotelRuleItem != nil && len(rsp.HotelRuleItem) > 0 {
		result.Rules = make([]*HotelRuleItem, 0)
		for _, hri := range rsp.HotelRuleItem {
			item := &HotelRuleItem{
				Name: hri.Name,
				Text: hri.Text,
			}

			result.Rules = append(result.Rules, item)
		}
	}

	//解析HotelType
	if rsp.HotelType != nil {
		result.SourceLink = bool(rsp.HotelType.SourceLink)
	}

	return result, nil
}
