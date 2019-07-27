package lowfare

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/air/lowfare"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestLowFareRQ 测试运价查询RQ
func TestLowFareRQ() string {
	param := &LowFareParam{
		Routes:     make([]*Route, 0),
		Modifiers:  &SearchModifiers{},
		CabinClass: "Y",
		AdultCount: 1,
		Currency:   "CNY",
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60
	param.Modifiers.NumOfStops = -1
	param.PlatingCarrier = "MU"

	dep := &Route{
		Origin:        "HKG",
		Destination:   "SHA",
		DepartureDate: "2018-01-01",
	}

	param.Routes = append(param.Routes, dep)

	ret := &Route{
		Origin:        "HKG",
		Destination:   "SHA",
		DepartureDate: "2018-01-10",
	}
	param.Routes = append(param.Routes, ret)

	param.Modifiers.IncludeAirlines = make([]string, 0)
	param.Modifiers.IncludeAirlines = append(param.Modifiers.IncludeAirlines, "CX")
	param.Modifiers.IncludeAirlines = append(param.Modifiers.IncludeAirlines, "KA")

	/*
		param.Modifiers.ExcludeAirlines = make([]string, 0)
		param.Modifiers.ExcludeAirlines = append(param.Modifiers.ExcludeAirlines, "CX")
		param.Modifiers.ExcludeAirlines = append(param.Modifiers.ExcludeAirlines, "KA")
	*/

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/LowFareParam.json", json)

	result := LowFare(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestLowFareRS 测试运价查询RS
func TestLowFareRS() *LowFareResult {
	val, err := LoadFile("samples/LowFareRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope airproxy.LowFareRspEnvelope
	xml.Unmarshal([]byte(val), &envelope)

	result, err := getResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}
