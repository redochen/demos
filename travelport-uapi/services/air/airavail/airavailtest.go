package airavail

import (
	"encoding/xml"
	"fmt"

	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/models/air/airavail"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	CcJson "github.com/redochen/tools/json"
)

//TestAirAvailRQ 测试AV查询RQ
func TestAirAvailRQ() string {
	param := &AirAvailParam{
		Routes:    make([]*Route, 0),
		Modifiers: &SearchModifiers{},
		//CabinClass: "Y",
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60
	param.Modifiers.NumOfStops = 1

	dep := &Route{
		Origin:        "SHA",
		Destination:   "SIN",
		DepartureDate: "2017-12-01",
	}

	param.Routes = append(param.Routes, dep)

	ret := &Route{
		Origin:        "SIN",
		Destination:   "SHA",
		DepartureDate: "2017-12-10",
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
	ccfile.DumpFile("samples/AirAvailParam.json", json)

	result := AirAvail(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestAirAvailRS 测试AV查询RS
func TestAirAvailRS() *AirAvailResult {
	val, err := LoadFile("samples/AirAvailRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope airproxy.AirAvailRspEnvelope
	xml.Unmarshal([]byte(val), &envelope)

	result, _, err := getResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}
