package hotelmedia

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelmedia"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestHotelMediaRQ 测试酒店媒介RQ
func TestHotelMediaRQ() string {
	param := &HotelMediaParam{
		HotelProperties: make([]*HotelPropertyParam, 0),
	}

	property1 := &HotelPropertyParam{
		Chain: "MC",
		Code:  "60411",
	}

	param.HotelProperties = append(param.HotelProperties, property1)

	property2 := &HotelPropertyParam{
		Chain: "PN",
		Code:  "81538",
	}

	param.HotelProperties = append(param.HotelProperties, property2)

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/HotelMediaParam.json", json)

	result := HotelMedia(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestHotelMediaRS 测试酒店媒介RS
func TestHotelMediaRS() *HotelMediaResult {
	val, err := LoadFile("samples/HotelMediaRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope hotproxy.HotelMediaRspEnvelope
	err = xml.Unmarshal([]byte(val), &envelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	result, err := getResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}
