package hotelavail

import (
	"encoding/xml"
	"fmt"

	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	CcFile "github.com/redochen/tools/file"
	CcJson "github.com/redochen/tools/json"
)

//TestHotelAvailRQ  测试酒店查询RQ
func TestHotelAvailRQ() string {
	param := &HotelAvailParam{
		Location:     "SHA",
		CheckinDate:  "2015-12-01",
		CheckoutDate: "2015-12-07",
	}

	json, _ := CcJson.Serialize(param)
	CcFile.DumpFile("samples/HotelAvailParam.json", json)

	result := HotelAvail(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestHotelAvailRS 测试酒店查询RS
func TestHotelAvailRS() *HotelAvailResult {
	val, err := LoadFile("samples/HotelAvailRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope hotproxy.HotelAvailRspEnvelope
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
