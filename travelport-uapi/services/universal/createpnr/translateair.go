package createpnr

import (
	"errors"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	. "github.com/redochen/demos/travelport-uapi/services/air"
	"github.com/redochen/demos/travelport-uapi/soap"
	airrq "github.com/redochen/demos/travelport-uapi/soap/air/request"
	airrs "github.com/redochen/demos/travelport-uapi/soap/air/response"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
)

//getAirReqEnvolpe 获取机票请求参数
func getAirReqEnvolpe(param *CreatePnrParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if nil == param.Travelers || len(param.Travelers) == 0 {
		return nil, errors.New("should contains at least one traveler")
	}

	if nil == param.Air || nil == param.Air.Segments || len(param.Air.Segments) == 0 {
		return nil, errors.New("should contains at least one segment")
	}

	body := uniproxy.NewAirCreatePnrReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create AirCreatePnrReqBody")
	}

	req := body.Request

	//设置价格、时间变化时是否保留预订：None, Schedule, Price, Both
	req.RetainReservation = "Both"

	//设置出票日期
	req.ActionStatus = make([]*comrq.ActionStatus, 0)
	actionStatus := &comrq.ActionStatus{
		Type:         "TAU",
		TicketDate:   param.Air.TicketDate,
		ProviderCode: "1G",
	}

	req.ActionStatus = append(req.ActionStatus, actionStatus)

	//设置旅客信息
	req.BookingTraveler = make([]*comrq.BookingTraveler, 0)
	for _, p := range param.Travelers {
		traveler := getBookingTraveler(p, "AIR", param.Air.TicketingCarrier)
		if traveler != nil {
			req.BookingTraveler = append(req.BookingTraveler, traveler)
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

	//设置航段信息
	req.AirPricingSolution = &airrq.AirPricingSolution{
		AirSegment: make([]*airrq.AirSegment, 0),
	}

	for segIndex, segment := range param.Air.Segments {
		airSegment := GetAirSegmentParam(segment, segIndex, param.ProviderCode, false, nil)

		//航段转换失败，则返回错误
		if nil == airSegment {
			return nil, errors.New("failed to get segment param")
		}

		req.AirPricingSolution.AirSegment = append(req.AirPricingSolution.AirSegment, airSegment)
	}

	//备注价格信息，格式需与供应商沟通
	/*
		priceRemark := &comrq.GeneralRemark{}
		//priceRemark.Category = "AIR"
		priceRemark.TypeInGds = com.TypeGdsRemark("Basic")
		priceRemark.RemarkData = ""

		req.GeneralRemark = make([]*comrq.GeneralRemark, 0)
		req.GeneralRemark = append(req.GeneralRemark, priceRemark)
	*/

	return soap.NewReqEnvelope(body), nil
}

//getAirResult 解析机票预订结果
func getAirResult(body *uniproxy.AirCreatePnrRspBody) (*CreatePnrResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("AirCreateReservationRsp is empty")
	}

	if nil == rsp.UniversalRecord {
		return nil, errors.New("universal record is empty")
	}

	result := &CreatePnrResult{}

	//解析预订结果
	ParseUniversalRecord(result, rsp.UniversalRecord)

	return result, nil
}

//getAirReservationResult 转换AirReservation结果
func getAirReservationResult(result *airrs.AirReservation) *AirReservationResult {
	if nil == result {
		return nil
	}

	reservation := &AirReservationResult{}
	reservation.LocatorCode = string(result.LocatorCode)
	reservation.CreateDate = result.CreateDate
	reservation.ModifiedDate = result.ModifiedDate
	reservation.BookingTravelerKey = parseBookingTravelerRefs(result.BookingTravelerRef)

	//解析供应商信息
	if result.SupplierLocator != nil && len(result.SupplierLocator) > 0 {
		reservation.Suppliers = make([]*SupplierResult, 0)
		for _, sl := range result.SupplierLocator {
			supplier := &SupplierResult{
				Code:               string(sl.SupplierCode),
				LocatorCode:        sl.SupplierLocatorCode,
				ReservationInfoRef: string(sl.ProviderReservationInfoRef),
				CreateDate:         sl.CreateDateTime,
			}

			reservation.Suppliers = append(reservation.Suppliers, supplier)
		}
	}

	//解析航段信息
	if result.AirSegment != nil && len(result.AirSegment) > 0 {
		reservation.Segments = make([]*Segment, 0)
		for idx, airSegment := range result.AirSegment {
			segment := GetSegmentResult(airSegment, idx, false, "", nil)
			if segment != nil {
				reservation.Segments = append(reservation.Segments, segment)
			}
		}
	}

	return reservation
}

//pasreAirFault 解析机票错误
func pasreAirFault(fault *uniproxy.AirCreatePnrRspFault) string {
	if nil == fault {
		return "unkown error"
	}

	if fault.Error != nil {
		if fault.Error.AirSegmentError != nil && len(fault.Error.AirSegmentError) > 0 {
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

	//return fault.Error.Description
	return fmt.Sprintf("%s(%s)", fault.Code, fault.Text)
}
