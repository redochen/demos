package airprice

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/air/airprice"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestAirPriceRQ 测试获取价格RQ
func TestAirPriceRQ() string {
	param := &AirPriceParam{
		Segments:   make([]*Segment, 0),
		Passengers: make([]*PassengerParam, 0),
		Carrier:    "CX",
		Currency:   "CNY",
		//CabinClass:         "Economy",
		CheckInventory:     true,
		SpecifyBookingCode: true,
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60

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

	adult := &PassengerParam{
		PassengerType: "ADT",
		ReferenceKey:  "1",
	}
	param.Passengers = append(param.Passengers, adult)

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/AirPriceParam.json", json)

	result := AirPrice(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestAirPriceRS 测试获取价格RS
func TestAirPriceRS() *AirPriceResult {
	val, err := LoadFile("samples/AirPriceRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope airproxy.AirPriceRspEnvelope
	xml.Unmarshal([]byte(val), &envelope)

	result, err := getResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}
