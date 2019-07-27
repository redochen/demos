package createpnr

import (
	. "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	unirs "github.com/redochen/demos/travelport-uapi/soap/universal/response"
	"strings"
)

//ParseUniversalRecord 解析预订结果
func ParseUniversalRecord(result *CreatePnrResult, universalRecord *unirs.UniversalRecord) {
	//解析PNR编码
	result.ParsePnrCode(universalRecord)

	//解析旅客信息
	if universalRecord.BookingTraveler != nil && len(universalRecord.BookingTraveler) > 0 {
		result.Travelers = make([]*Traveler, 0)
		for _, bt := range universalRecord.BookingTraveler {
			traveler := getTraveler(bt)
			if traveler != nil {
				result.Travelers = append(result.Travelers, traveler)
			}
		}
	}

	//解析供应商信息
	if universalRecord.ProviderReservationInfo != nil && len(universalRecord.ProviderReservationInfo) > 0 {
		result.ProviderInfos = make([]*ProviderInfo, 0)
		for _, pri := range universalRecord.ProviderReservationInfo {
			provider := getProviderInfo(pri)
			if provider != nil {
				result.ProviderInfos = append(result.ProviderInfos, provider)
			}
		}
	}

	//解析OSI
	if universalRecord.OSI != nil && len(universalRecord.OSI) > 0 {
		result.OtherServiceInfos = make([]*OtherServiceInfo, 0)
		for _, osi := range universalRecord.OSI {
			otherSvcInfo := getOtherServiceInfo(osi)
			if otherSvcInfo != nil {
				result.OtherServiceInfos = append(result.OtherServiceInfos, otherSvcInfo)
			}
		}
	}

	//解析代理商信息
	result.AgencyInfos = getAgencyInfos(universalRecord.AgencyInfo)

	//解析支付信息
	if universalRecord.FormOfPayment != nil && len(universalRecord.FormOfPayment) > 0 {
		result.Payments = make([]*FormOfPayment, 0)
		for _, p := range universalRecord.FormOfPayment {
			payment := getFormOfPaymentResult(p)
			if payment != nil {
				result.Payments = append(result.Payments, payment)
			}
		}
	}

	//解析机票预订结果
	if universalRecord.AirReservation != nil && len(universalRecord.AirReservation) > 0 {
		result.AirReservations = make([]*AirReservationResult, 0)
		for _, ar := range universalRecord.AirReservation {
			reservation := getAirReservationResult(ar)
			if reservation != nil {
				result.AirReservations = append(result.AirReservations, reservation)
			}
		}
	}

	//解析酒店预订结果
	if universalRecord.HotelReservation != nil && len(universalRecord.HotelReservation) > 0 {
		result.HotelReservations = make([]*HotelReservationResult, 0)
		for _, hr := range universalRecord.HotelReservation {
			reservation := getHotelReservationResult(hr)
			if reservation != nil {
				result.HotelReservations = append(result.HotelReservations, reservation)
			}
		}
	}
}

//getBookingTraveler 转换旅客信息参数
func getBookingTraveler(param *Traveler, category, airline string) *comrq.BookingTraveler {
	if nil == param {
		return nil
	}

	if strings.EqualFold(param.PassengerType, "CHD") {
		param.PassengerType = "CNN"
	}

	traveler := &comrq.BookingTraveler{
		Key:          com.TypeRef(param.Key),
		DOB:          param.BirthDate,
		Gender:       com.TypeGender(param.Gender),
		TravelerType: com.TypePTC(param.PassengerType),
		BookingTravelerName: &comrq.BookingTravelerName{
			First: param.GivenName,
			Last:  com.TypeTravelerLastName(param.Surname),
		},
		Nationality: com.TypeCountry(param.Nationality),
	}

	//设置称呼
	if len(param.ReferenceName) > 0 {
		traveler.BookingTravelerName.Prefix = param.ReferenceName
	}

	//处理儿童年龄
	if strings.EqualFold(param.PassengerType, "CNN") {
		//更改儿童类型为C0N
		traveler.TravelerType = com.TypePTC(param.GetChildYear(false))

		rmk := &comrq.NameRemark{
			Category:   category,
			RemarkData: param.GetChildYear(true),
		}

		traveler.NameRemark = make([]*comrq.NameRemark, 0)
		traveler.NameRemark = append(traveler.NameRemark, rmk)
	}

	//设置电话信息
	traveler.PhoneNumber = make([]*comrq.PhoneNumber, 0)

	if param.Phones != nil && len(param.Phones) > 0 {
		for _, pn := range param.Phones {
			phone := getPhoneNumber(pn)
			if phone != nil {
				traveler.PhoneNumber = append(traveler.PhoneNumber, phone)
			}
		}
	} else {
		phone := &comrq.PhoneNumber{
			Location:    "SHA",
			CountryCode: "86",
			AreaCode:    "21",
			Number:      "62626289",
		}
		traveler.PhoneNumber = append(traveler.PhoneNumber, phone)
	}

	//设置邮件信息
	if param.Emails != nil && len(param.Emails) > 0 {
		traveler.Email = make([]*comrq.Email, 0)
		for _, m := range param.Emails {
			email := getEmailParam(m)
			if email != nil {
				traveler.Email = append(traveler.Email, email)
			}
		}
	}

	//设置地址信息
	if param.Addresses != nil && len(param.Addresses) > 0 {
		traveler.Address = make([]*comrq.TypeStructuredAddress, 0)
		for _, addr := range param.Addresses {
			address := getAddressParam(addr)
			if address != nil {
				traveler.Address = append(traveler.Address, address)
			}
		}
	}

	//设置常旅客信息
	if len(param.FFPCardNo) > 0 {
		traveler.LoyaltyCard = make([]*comrq.LoyaltyCard, 0)
		loyaltyCard := &comrq.LoyaltyCard{
			SupplierType: "Air",
			SupplierCode: com.TypeCarrier(param.FFPCarrier),
			CardNumber:   com.TypeCardNumber(param.FFPCardNo),
		}

		if len(airline) > 0 {
			loyaltyCard.ProviderReservationSpecificInfo = make([]*comrq.ProviderReservationSpecificInfo, 0)

			specificInfo := &comrq.ProviderReservationSpecificInfo{
				OperatedBy: make([]*comrq.OperatedBy, 0),
			}
			specificInfo.OperatedBy = append(specificInfo.OperatedBy, &comrq.OperatedBy{
				Value: airline,
			})

			loyaltyCard.ProviderReservationSpecificInfo = append(loyaltyCard.ProviderReservationSpecificInfo, specificInfo)
		}

		traveler.LoyaltyCard = append(traveler.LoyaltyCard, loyaltyCard)
	}

	return traveler
}

//getTraveler 转换旅客信息结果
func getTraveler(result *comrs.BookingTraveler) *Traveler {
	if nil == result {
		return nil
	}

	traveler := &Traveler{
		PassengerType: string(result.TravelerType),
		BirthDate:     result.DOB,
		Gender:        string(result.Gender),
		Nationality:   string(result.Nationality),
		Key:           string(result.Key),
		ElStat:        result.ElStat,
	}

	//解析姓名和称呼
	if result.BookingTravelerName != nil {
		traveler.Surname = string(result.BookingTravelerName.Last)
		traveler.GivenName = result.BookingTravelerName.First
		traveler.ReferenceName = result.BookingTravelerName.Prefix
	}

	//解析电话信息
	if result.PhoneNumber != nil && len(result.PhoneNumber) > 0 {
		traveler.Phones = make([]*Phone, 0)
		for _, pn := range result.PhoneNumber {
			phone := getPhone(pn)
			if phone != nil {
				traveler.Phones = append(traveler.Phones, phone)
			}
		}
	}

	//解析邮件信息
	if result.Email != nil && len(result.Email) > 0 {
		traveler.Emails = make([]*Email, 0)
		for _, m := range result.Email {
			email := getEmailResult(m)
			if email != nil {
				traveler.Emails = append(traveler.Emails, email)
			}
		}
	}

	//解析地址信息
	if result.Address != nil && len(result.Address) > 0 {
		traveler.Addresses = make([]*Address, 0)
		for _, addr := range result.Address {
			address := getAddressResult(addr)
			if address != nil {
				traveler.Addresses = append(traveler.Addresses, address)
			}
		}
	}

	return traveler
}

//getPhoneNumber 转换电话参数
func getPhoneNumber(param *Phone) *comrq.PhoneNumber {
	if nil == param {
		return nil
	}

	phone := &comrq.PhoneNumber{
		Location:    param.Location,
		CountryCode: param.CountryCode,
		AreaCode:    param.AreaCode,
		Number:      param.Number,
		Type:        param.Type,
	}

	return phone
}

//getPhone 转换电话结果
func getPhone(result *comrs.PhoneNumber) *Phone {
	if nil == result {
		return nil
	}

	phone := &Phone{
		Location:    result.Location,
		CountryCode: result.CountryCode,
		AreaCode:    result.AreaCode,
		Number:      result.Number,
		Type:        result.Type,
	}

	return phone
}

//getEmailParam 转换邮件参数
func getEmailParam(param *Email) *comrq.Email {
	if nil == param {
		return nil
	}

	email := &comrq.Email{
		EmailID: param.Id,
		Type:    com.TypeRef(param.Type),
	}

	return email
}

//getEmailResult 转换邮件结果
func getEmailResult(result *comrs.Email) *Email {
	if nil == result {
		return nil
	}

	email := &Email{
		Id:   result.EmailID,
		Type: string(result.Type),
	}

	return email
}

//getAddressParam 转换地址参数
func getAddressParam(param *Address) *comrq.TypeStructuredAddress {
	if nil == param {
		return nil
	}
	address := &comrq.TypeStructuredAddress{
		AddressName: param.Address,
		Country:     param.Country,
		State: &comrq.State{
			Value: param.State,
		},
		City:       param.City,
		Street:     param.Street,
		PostalCode: param.PostalCode,
	}

	return address
}

//getAddressResult 转换地址结果
func getAddressResult(result *comrs.TypeStructuredAddress) *Address {
	if nil == result {
		return nil
	}

	address := &Address{
		Address:    result.AddressName,
		Country:    result.Country,
		City:       result.City,
		Street:     result.Street,
		PostalCode: result.PostalCode,
		Key:        string(result.Key),
		ElStat:     result.ElStat,
	}

	if result.State != nil {
		address.State = result.State.Value
	}

	return address
}

//getOSI 转换其他服务信息参数
func getOSI(param *OtherServiceInfo) *comrq.OSI {
	if nil == param {
		return nil
	}

	osi := &comrq.OSI{
		Key:                        com.TypeRef(param.Key),
		Carrier:                    com.TypeCarrier(param.Carrier),
		Code:                       param.Code,
		Text:                       param.Text,
		ProviderReservationInfoRef: com.TypeRef(param.ProviderInfoRef),
		ProviderCode:               com.TypeProviderCode(param.ProviderCode),
	}

	return osi
}

//getOtherServiceInfo 转换其他服务信息结果
func getOtherServiceInfo(result *comrs.OSI) *OtherServiceInfo {
	if nil == result {
		return nil
	}

	osi := &OtherServiceInfo{
		Carrier:         string(result.Carrier),
		Code:            result.Code,
		Text:            result.Text,
		ProviderInfoRef: string(result.ProviderReservationInfoRef),
		ProviderCode:    string(result.ProviderCode),
		Key:             string(result.Key),
		ElStat:          result.ElStat,
	}

	return osi
}

//getFormOfPaymentParam 转换支付信息参数
func getFormOfPaymentParam(param *FormOfPayment) *comrq.FormOfPayment {
	if nil == param {
		return nil
	}

	payment := &comrq.FormOfPayment{
		Key:        com.TypeRef(param.Key),
		ProfileID:  param.ProfileID,
		ProfileKey: com.TypeRef(param.ProfileKey),
	}

	//转换供应商信息引用Key
	if param.ProviderInfoKeys != nil && len(param.ProviderInfoKeys) > 0 {
		payment.ProviderReservationInfoRef = make([]*comrq.ProviderReservationInfoRef, 0)
		for _, pik := range param.ProviderInfoKeys {
			payment.ProviderReservationInfoRef = append(payment.ProviderReservationInfoRef,
				&comrq.ProviderReservationInfoRef{
					Key: com.TypeRef(pik),
				})
		}
	}

	//转换信用卡信息
	if param.CreditCard != nil {
		payment.Type = "Credit"
		payment.CreditCard = getCreditCardParam(param.CreditCard)
	}

	return payment
}

//getFormOfPaymentResult 转换支付信息结果
func getFormOfPaymentResult(result *comrs.FormOfPayment) *FormOfPayment {
	if nil == result {
		return nil
	}

	payment := &FormOfPayment{
		Type:       result.Type,
		ProfileID:  result.ProfileID,
		ProfileKey: string(result.ProfileKey),
		Reusable:   result.Reusable,
		Key:        string(result.Key),
		ElStat:     result.ElStat,
	}

	//转换供应商信息引用Key
	if result.ProviderReservationInfoRef != nil && len(result.ProviderReservationInfoRef) > 0 {
		payment.ProviderInfoKeys = make([]string, 0)
		for _, pik := range result.ProviderReservationInfoRef {
			payment.ProviderInfoKeys = append(payment.ProviderInfoKeys, string(pik.Key))
		}
	}

	//转换信用卡信息
	if result.CreditCard != nil {
		payment.CreditCard = getCreditCardResult(result.CreditCard)
	}

	return payment
}

//getCreditCardParam 转换信用卡参数
func getCreditCardParam(param *CreditCard) *comrq.CreditCard {
	if nil == param {
		return nil
	}

	cc := &comrq.CreditCard{
		Key: com.TypeRef(param.Key),
	}

	cc.BankCountryCode = com.TypeCountry(param.CountryCode)
	cc.BankStateCode = com.TypeState(param.StateCode)
	cc.BankName = param.BankName
	cc.Number = com.TypeCreditCardNumber(param.Number)
	cc.Type = com.TypeCardMerchantType(param.Type)
	cc.ExpDate = param.ExpiryDate
	cc.CVV = param.CVV
	cc.ApprovalCode = param.ApprovalCode
	cc.BillingAddress = getAddressParam(param.BillingAddress)
	cc.Key = com.TypeRef(param.Key)
	cc.ProfileID = param.ProfileID

	return cc
}

//getCreditCardResult 转换信用卡结果
func getCreditCardResult(result *comrs.CreditCard) *CreditCard {
	if nil == result {
		return nil
	}

	cc := &CreditCard{
		CountryCode:    string(result.BankCountryCode),
		StateCode:      string(result.BankStateCode),
		BankName:       result.BankName,
		Number:         string(result.Number),
		Type:           string(result.Type),
		ExpiryDate:     result.ExpDate,
		CVV:            result.CVV,
		ApprovalCode:   result.ApprovalCode,
		BillingAddress: getAddressResult(result.BillingAddress),
		Key:            string(result.Key),
		ProfileID:      result.ProfileID,
	}

	return cc
}

//getProviderInfo 转换供应商预订信息结果
func getProviderInfo(result *unirs.ProviderReservationInfo) *ProviderInfo {
	if nil == result {
		return nil
	}

	provider := &ProviderInfo{
		ProviderCode:   string(result.ProviderCode),
		LocatorCode:    string(result.LocatorCode),
		CreateDate:     result.CreateDate,
		ModifiedDate:   result.ModifiedDate,
		HostCreateDate: result.HostCreateDate,
		OwningPCC:      string(result.OwningPCC),
		Key:            string(result.Key),
		ElStat:         result.ElStat,
	}

	return provider
}

//getAgencyInfos 转换代理商信息结果
func getAgencyInfos(result *comrs.AgencyInfo) []*AgencyInfo {
	if nil == result || nil == result.AgentAction || len(result.AgentAction) == 0 {
		return nil
	}

	agencyInfos := make([]*AgencyInfo, 0)
	for _, aa := range result.AgentAction {
		agencyInfo := &AgencyInfo{
			ActionType: aa.ActionType,
			AgentCode:  aa.AgentCode,
			BranchCode: string(aa.BranchCode),
			AgencyCode: aa.AgencyCode,
			EventTime:  aa.EventTime,
		}
		agencyInfos = append(agencyInfos, agencyInfo)
	}

	return agencyInfos
}

//parseBookingTravelerRefs 解析预订旅客引用信息
func parseBookingTravelerRefs(result []*comrs.BookingTravelerRef) []string {
	if nil == result || len(result) == 0 {
		return nil
	}

	array := make([]string, 0)
	for _, btr := range result {
		array = append(array, string(btr.Key))
	}

	return array
}
