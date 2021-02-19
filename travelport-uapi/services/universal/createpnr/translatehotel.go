package createpnr

import (
	"errors"
	"fmt"

	. "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
	CcStr "github.com/redochen/tools/string"
	CcTime "github.com/redochen/tools/time"
)

//getHotelReqEnvolpe 获取酒店请求参数
func getHotelReqEnvolpe(param *CreatePnrParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if nil == param.Travelers || len(param.Travelers) == 0 {
		return nil, errors.New("should contains at least one traveler")
	}

	if nil == param.Hotel || nil == param.Hotel.Property {
		return nil, errors.New("hotel property cannot be nil")
	}

	if nil == param.Hotel.Rates || len(param.Hotel.Rates) == 0 {
		return nil, errors.New("hotel rate cannot be nil")
	}

	body := uniproxy.NewHotelCreatePnrReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create HotelCreatePnrReqBody")
	}

	req := body.Request

	//设置旅客信息
	req.BookingTraveler = make([]*comrq.BookingTraveler, 0)
	for _, p := range param.Travelers {
		traveler := getBookingTraveler(p, "HOTEL", "")
		if traveler != nil {
			req.BookingTraveler = append(req.BookingTraveler, traveler)
		}
	}

	//设置HostToken
	if param.Hotel.HostToken != nil {
		req.HostToken = &comrq.HostToken{
			Host:  com.TypeProviderCode(param.Hotel.HostToken.Host),
			Key:   param.Hotel.HostToken.Key,
			Value: param.Hotel.HostToken.Token,
		}
	}

	//设置OSI
	if param.OtherServiceInfos != nil && len(param.OtherServiceInfos) > 0 {
		req.OSI = make([]*comrq.OSI, 0)
		for _, osi := range param.OtherServiceInfos {
			otherSvcInfo := getOSI(osi)
			if otherSvcInfo != nil {
				req.OSI = append(req.OSI, otherSvcInfo)
			}
		}
	}

	//设置支付信息
	if param.Payments != nil && len(param.Payments) > 0 {
		req.FormOfPayment = make([]*comrq.FormOfPayment, 0)
		for _, p := range param.Payments {
			payment := getFormOfPaymentParam(p)
			if payment != nil {
				req.FormOfPayment = append(req.FormOfPayment, payment)
			}
		}
	}

	//设置酒店信息
	req.HotelProperty = getHotelPropertyParam(param.Hotel.Property)

	//设置价格信息
	req.HotelRateDetail = make([]*hotrq.HotelRateDetail, 0)
	for _, r := range param.Hotel.Rates {
		rate := getHotelRateDetail(r)
		if rate != nil {
			req.HotelRateDetail = append(req.HotelRateDetail, rate)
		}
	}

	//设置入住时间
	if len(param.Hotel.CheckinDate) > 0 || len(param.Hotel.CheckoutDate) > 0 {
		req.HotelStay = &hotrq.HotelStay{
			CheckinDate:  hot.TypeDate(CcTime.AddDateSeparator(param.Hotel.CheckinDate, "-", true)),
			CheckoutDate: hot.TypeDate(CcTime.AddDateSeparator(param.Hotel.CheckoutDate, "-", true)),
		}
	}

	//设置人数及房间数
	req.GuestInformation = &hotrq.GuestInformation{
		NumberOfRooms: param.Hotel.Rooms,
		NumberOfAdults: &hotrq.NumberOfAdults{
			Value: CcStr.FormatInt(param.Hotel.Adults),
		},
	}

	if param.Hotel.Children > 0 {
		req.GuestInformation.NumberOfChildren = &hotrq.NumberOfChildren{
			Count: param.Hotel.Children,
		}
	}

	return soap.NewReqEnvelope(body), nil
}

//getHotelResult 解析酒店预订结果
func getHotelResult(body *uniproxy.HotelCreatePnrRspBody) (*CreatePnrResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("HotelCreateReservationRsp is empty")
	}

	if nil == rsp.UniversalRecord {
		return nil, errors.New("universal record is empty")
	}

	result := &CreatePnrResult{}

	//解析预订结果
	ParseUniversalRecord(result, rsp.UniversalRecord)

	return result, nil
}

//getHotelReservationResult 转换HotelReservation结果
func getHotelReservationResult(result *hotrs.HotelReservation) *HotelReservationResult {
	if nil == result {
		return nil
	}

	reservation := &HotelReservationResult{}
	reservation.LocatorCode = string(result.LocatorCode)
	reservation.CreateDate = result.CreateDate
	reservation.ModifiedDate = result.ModifiedDate
	reservation.BookingTravelerKey = parseBookingTravelerRefs(result.BookingTravelerRef)

	//解析预订姓名
	if result.ReservationName != nil {
		if result.ReservationName.BookingTravelerRef != nil {
			reservation.TravelerKey = string(result.ReservationName.BookingTravelerRef.Key)
		}

		if result.ReservationName.NameOverride != nil {
			reservation.Surname = result.ReservationName.NameOverride.Last
			reservation.GivenName = result.ReservationName.NameOverride.First
		}
	}

	//解析酒店物业信息
	reservation.Property = getHotelPropertyResult(result.HotelProperty)

	//解析酒店价格信息
	if result.HotelRateDetail != nil && len(result.HotelRateDetail) > 0 {
		reservation.Rates = make([]*HotelRate, 0)
		for _, hrd := range result.HotelRateDetail {
			rate := getHotelRate(hrd)
			if rate != nil {
				reservation.Rates = append(reservation.Rates, rate)
			}
		}
	}

	//解析入住时间
	if result.HotelStay != nil {
		reservation.CheckinDate = CcTime.AddDateSeparator(string(result.HotelStay.CheckinDate), "-", true)
		reservation.CheckoutDate = CcTime.AddDateSeparator(string(result.HotelStay.CheckoutDate), "-", true)
	}

	//解析预订来源
	if result.BookingSource != nil {
		reservation.SourceType = result.BookingSource.Type
		reservation.SourceCode = result.BookingSource.Code
	}

	//解析人数及房间数
	if result.GuestInformation != nil {
		reservation.Rooms = result.GuestInformation.NumberOfRooms

		if result.GuestInformation.NumberOfAdults != nil {
			reservation.Adults = CcStr.ParseInt(result.GuestInformation.NumberOfAdults.Value)
		}

		if result.GuestInformation.NumberOfChildren != nil {
			reservation.Children = result.GuestInformation.NumberOfChildren.Count
		}
	}

	//解析售卖消息
	if result.SellMessage != nil && len(result.SellMessage) > 0 {
		reservation.SellMessage = make([]string, 0)
		for _, sm := range result.SellMessage {
			reservation.SellMessage = append(reservation.SellMessage, string(sm))
		}
	}

	return reservation
}

//getHotelPropertyParam 转换HotelProperty参数
func getHotelPropertyParam(param *HotelProperty) *hotrq.HotelProperty {
	if nil == param {
		return nil
	}

	property := &hotrq.HotelProperty{
		HotelChain:         com.TypeHotelChainCode(param.Chain),
		HotelCode:          com.TypeHotelCode(param.Code),
		HotelLocation:      hot.TypeHotelLocationCode(param.Location),
		Name:               param.Name,
		ParticipationLevel: com.StringLength1(param.ParticipationLevel),
	}

	//设置酒店地址
	if param.Addresses != nil && len(param.Addresses) > 0 {
		property.PropertyAddress = &hotrq.UnstructuredAddress{
			Address: make([]string, 0),
		}

		for _, addr := range param.Addresses {
			property.PropertyAddress.Address = append(property.PropertyAddress.Address, addr)
		}
	}

	//设置酒店电话
	if param.Phones != nil && len(param.Phones) > 0 {
		property.PhoneNumber = make([]*comrq.PhoneNumber, 0)
		for _, ph := range param.Phones {
			phone := getPhoneNumber(ph)
			if phone != nil {
				property.PhoneNumber = append(property.PhoneNumber, phone)
			}
		}
	}

	return property
}

//getHotelPropertyResult 转换HotelProperty结果
func getHotelPropertyResult(result *hotrs.HotelProperty) *HotelProperty {
	if nil == result {
		return nil
	}

	property := &HotelProperty{
		Chain:              string(result.HotelChain),
		Code:               string(result.HotelCode),
		Location:           string(result.HotelLocation),
		Name:               result.Name,
		ParticipationLevel: string(result.ParticipationLevel),
		Availability:       result.Availability,
	}

	//设置酒店地址
	if result.PropertyAddress != nil && result.PropertyAddress.Address != nil &&
		len(result.PropertyAddress.Address) > 0 {
		property.Addresses = make([]string, 0)
		for _, addr := range result.PropertyAddress.Address {
			property.Addresses = append(property.Addresses, addr)
		}
	}

	//设置酒店电话
	if result.PhoneNumber != nil && len(result.PhoneNumber) > 0 {
		property.Phones = make([]*Phone, 0)
		for _, ph := range result.PhoneNumber {
			phone := getPhone(ph)
			if phone != nil {
				property.Phones = append(property.Phones, phone)
			}
		}
	}

	return property
}

//getHotelRateDetail 转换HotelRateDetail参数
func getHotelRateDetail(param *HotelRate) *hotrq.HotelRateDetail {
	if nil == param {
		return nil
	}

	rate := &hotrq.HotelRateDetail{
		RatePlanType:   com.TypeRatePlanType(param.RatePlanType),
		RateGuaranteed: param.Guaranteed,
	}

	if param.Category > 0 {
		rate.RateCategory = com.TypeOTACode(param.Category)
	}

	if param.Base > 0 {
		rate.Base = com.TypeMoney(fmt.Sprintf("%s%.2f", param.BaseCurrency, param.Base))
	}
	if param.Tax > 0 {
		rate.Tax = com.TypeMoney(fmt.Sprintf("%s%.2f", param.TaxCurrency, param.Tax))
	}
	if param.Total > 0 {
		rate.Total = com.TypeMoney(fmt.Sprintf("%s%.2f", param.TotalCurrency, param.Total))
	}

	//转换价格描述
	if param.Descriptions != nil && len(param.Descriptions) > 0 {
		rate.RoomRateDescription = make([]*hotrq.HotelRateDescription, 0)
		rate.RoomRateDescription = append(rate.RoomRateDescription, &hotrq.HotelRateDescription{
			Text: param.Descriptions,
		})
	}

	//转换限期价格
	if param.RateByDates != nil && len(param.RateByDates) > 0 {
		rate.HotelRateByDate = make([]*hotrq.HotelRateByDate, 0)
		for _, rbd := range param.RateByDates {
			dateRate := getHotelRateByDateParam(rbd)
			if dateRate != nil {
				rate.HotelRateByDate = append(rate.HotelRateByDate, dateRate)
			}
		}
	}

	return rate
}

//getHotelRate 转换HotelRateDetail结果
func getHotelRate(result *hotrs.HotelRateDetail) *HotelRate {
	if nil == result {
		return nil
	}

	rate := &HotelRate{
		RatePlanType:  string(result.RatePlanType),
		Base:          result.Base.GetFloatAmount(),
		BaseCurrency:  result.Base.GetCurrency(),
		Tax:           result.Tax.GetFloatAmount(),
		TaxCurrency:   result.Tax.GetCurrency(),
		Total:         result.Total.GetFloatAmount(),
		TotalCurrency: result.Total.GetCurrency(),
		Category:      int(result.RateCategory),
		Guaranteed:    result.RateGuaranteed,
	}

	//解析价格描述
	if result.RoomRateDescription != nil && len(result.RoomRateDescription) > 0 {
		rate.Descriptions = make([]string, 0)
		for _, desc := range result.RoomRateDescription {
			if nil == desc || nil == desc.Text || len(desc.Text) == 0 {
				continue
			}

			for _, text := range desc.Text {
				rate.Descriptions = append(rate.Descriptions, text)
			}
		}
	}

	//解析限期价格
	if result.HotelRateByDate != nil && len(result.HotelRateByDate) > 0 {
		rate.RateByDates = make([]*HotelRateByDate, 0)
		for _, rbd := range result.HotelRateByDate {
			dateRate := getHotelRateByDateResult(rbd)
			if dateRate != nil {
				rate.RateByDates = append(rate.RateByDates, dateRate)
			}
		}
	}

	return rate
}

//getHotelRateByDateParam 转换HotelRateByDate参数
func getHotelRateByDateParam(param *HotelRateByDate) *hotrq.HotelRateByDate {
	if nil == param {
		return nil
	}

	dateRate := &hotrq.HotelRateByDate{
		Base:          com.TypeMoney(fmt.Sprintf("%s%.2f", param.BaseCurrency, param.Base)),
		Tax:           com.TypeMoney(fmt.Sprintf("%s%.2f", param.TaxCurrency, param.Tax)),
		Total:         com.TypeMoney(fmt.Sprintf("%s%.2f", param.TotalCurrency, param.Total)),
		EffectiveDate: param.EffectiveDate,
		ExpireDate:    param.ExpireDate,
		Contents:      param.Contents,
	}

	return dateRate
}

//getHotelRateByDateResult 转换HotelRateByDate结果
func getHotelRateByDateResult(result *hotrs.HotelRateByDate) *HotelRateByDate {
	if nil == result {
		return nil
	}

	dateRate := &HotelRateByDate{
		Base:          result.Base.GetFloatAmount(),
		BaseCurrency:  result.Base.GetCurrency(),
		Tax:           result.Tax.GetFloatAmount(),
		TaxCurrency:   result.Tax.GetCurrency(),
		Total:         result.Total.GetFloatAmount(),
		TotalCurrency: result.Total.GetCurrency(),
		EffectiveDate: result.EffectiveDate,
		ExpireDate:    result.ExpireDate,
		Contents:      result.Contents,
	}

	return dateRate
}

//parseHotelFault 解析酒店错误
func parseHotelFault(fault *uniproxy.HotelCreatePnrRspFault) string {
	if nil == fault {
		return "unkown error"
	}

	//return fault.Error.Description
	return fmt.Sprintf("%s(%s)", fault.Code, fault.Text)
}
