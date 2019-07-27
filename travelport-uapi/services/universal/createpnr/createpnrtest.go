package createpnr

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestCreateAirPnrRQ 测试创建机票PNR RQ
func TestCreateAirPnrRQ() string {
	param := &CreatePnrParam{
		Travelers: make([]*Traveler, 0),
		Air:       getAirParam(),
		//Payments:  make([]*FormOfPayment, 0),
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60
	param.Travelers = append(param.Travelers, getAdultTraveler())
	//param.Payments = append(param.Payments, getFormOfPayment())

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/CreateAirPnrParam.json", json)

	return testCreatePnrRQ(param)
}

//TestCreateHotelPnrRQ 测试创建酒店PNR RQ
func TestCreateHotelPnrRQ() string {
	param := &CreatePnrParam{
		Travelers: make([]*Traveler, 0),
		Hotel:     getHotelParam(),
		Payments:  make([]*FormOfPayment, 0),
	}

	param.Travelers = append(param.Travelers, getAdultTraveler())
	param.Payments = append(param.Payments, getDefaultFormOfPayment())

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/CreateHotelPnrParam.json", json)

	return testCreatePnrRQ(param)
}

//testCreatePnrRQ
func testCreatePnrRQ(param *CreatePnrParam) string {
	if nil == param {
		return ""
	}

	result := CreatePnr(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestCreateAirPnrRS 测试创建机票PNR RS
func TestCreateAirPnrRS() *CreatePnrResult {
	val, err := LoadFile("samples/CreateAirPnrRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope uniproxy.AirCreatePnrRspEnvelope
	err = xml.Unmarshal([]byte(val), &envelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	if envelope.Body != nil && envelope.Body.Fault != nil {
		return SetErrorMessage(pasreAirFault(envelope.Body.Fault))
	}

	result, err := getAirResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}

//TestCreateHotelPnrRS 测试创建酒店PNR RS
func TestCreateHotelPnrRS() *CreatePnrResult {
	val, err := LoadFile("samples/CreateHotelPnrRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope uniproxy.HotelCreatePnrRspEnvelope
	err = xml.Unmarshal([]byte(val), &envelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	if envelope.Body != nil && envelope.Body.Fault != nil {
		return SetErrorMessage(parseHotelFault(envelope.Body.Fault))
	}

	result, err := getHotelResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}

//getAdultTraveler 获取成人旅客
func getAdultTraveler() *Traveler {
	adult := &Traveler{
		Key:           "1",
		Surname:       "Li",
		GivenName:     "Si",
		PassengerType: "ADT",
		ReferenceName: "Mr",
		Gender:        "M",
		BirthDate:     "1980-08-08",
		Phones:        make([]*Phone, 0),
		Emails:        make([]*Email, 0),
		Addresses:     make([]*Address, 0),
	}

	phone := &Phone{
		CountryCode: "86",
		Location:    "SHA",
		AreaCode:    "21",
		Number:      "123456",
		Type:        "Home",
	}
	adult.Phones = append(adult.Phones, phone)

	email := &Email{
		Id:   "lisi@email.com",
		Type: "Home",
	}
	adult.Emails = append(adult.Emails, email)

	address := &Address{
		Country:    "CN",
		State:      "SHA",
		City:       "Shanghai",
		Address:    "Xuihui Dist.",
		PostalCode: "200000",
		Street:     make([]string, 0),
	}
	address.Street = append(address.Street, "No.8 Linyun Rd")

	adult.Addresses = append(adult.Addresses, address)

	/*
		child := &Traveler{
			Surname:       "Zhang",
			GivenName:     "Sisi",
			PassengerType: "CNN",
			//ReferenceName: "Miss*P-C5",
			ReferenceName: "Miss",
			Gender:        "F",
			BirthDate:     "2010-06-08",
		}
		param.Travelers = append(param.Travelers, child)
	*/

	return adult
}

//获取支付信息
func getDefaultFormOfPayment() *FormOfPayment {
	param := &FormOfPayment{
		Key:  "1",
		Type: "Credit",
		CreditCard: &CreditCard{
			Type:       "MC",
			CVV:        "xxxxxx",
			BankName:   "China Bank",
			Number:     "123456789012",
			ExpiryDate: "2018-12",
			BillingAddress: &Address{
				Country:    "CN",
				State:      "SHA",
				City:       "Shanghai",
				Address:    "Xuihui Dist.",
				PostalCode: "200000",
				Street:     make([]string, 0),
			},
		},
	}

	param.CreditCard.BillingAddress.Street = append(
		param.CreditCard.BillingAddress.Street,
		"No.8 Linyun Rd")

	return param
}

//获取机票参数
func getAirParam() *AirParam {
	param := &AirParam{
		Segments:         make([]*Segment, 0),
		TicketDate:       "2017-11-01T00:00:00",
		TicketingCarrier: "CX",
	}

	s1 := &Segment{
		SegmentKey:                "1",
		Group:                     0,
		Carrier:                   "CX",
		FlightNumber:              711,
		DepartureAirport:          "HKG",
		ArrivalAirport:            "SIN",
		DepartureTime:             "2017-12-01T16:10:00.000+08:00",
		ArrivalTime:               "2017-12-01T20:05:00.000+08:00",
		LinkAvailability:          true,
		AvailabilityDisplayType:   "General",
		AvailabilitySource:        "S",
		OptionalServicesIndicator: false,
		ParticipantLevel:          "Secure Sell",
		PolledAvailabilityOption:  "Polled avail used",
		ProviderCode:              "1G",
		BookingCode:               "I",
	}
	param.Segments = append(param.Segments, s1)

	s2 := &Segment{
		SegmentKey:                "2",
		Group:                     0,
		Carrier:                   "CX",
		FlightNumber:              734,
		DepartureAirport:          "SIN",
		ArrivalAirport:            "HKG",
		DepartureTime:             "2017-12-10T16:30:00.000+08:00",
		ArrivalTime:               "2017-12-10T20:25:00.000+08:00",
		LinkAvailability:          true,
		AvailabilityDisplayType:   "General",
		AvailabilitySource:        "S",
		OptionalServicesIndicator: false,
		ParticipantLevel:          "Secure Sell",
		PolledAvailabilityOption:  "Polled avail used",
		ProviderCode:              "1G",
		BookingCode:               "I",
	}
	param.Segments = append(param.Segments, s2)

	return param
}

//获取酒店参数
func getHotelParam() *HotelParam {
	param := &HotelParam{
		Property: &HotelProperty{
			Chain:     "PH",
			Code:      "95553",
			Name:      "REGAL AIRPORT HOTEL LIFESTYLE",
			Addresses: make([]string, 0),
		},
		Rates:        make([]*HotelRate, 0),
		Rooms:        1,
		Adults:       1,
		CheckinDate:  "2015-12-01",
		CheckoutDate: "2015-12-07",
		HostToken:    nil,
	}

	param.Property.Addresses = append(param.Property.Addresses, "9 Cheong Tat Road")
	param.Property.Addresses = append(param.Property.Addresses, "Hong Kong HK 000000")

	rate := &HotelRate{
		RatePlanType: "B2TPAP",
	}
	param.Rates = append(param.Rates, rate)

	return param
}
