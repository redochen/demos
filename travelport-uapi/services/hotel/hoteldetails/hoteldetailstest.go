package hoteldetails

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hoteldetails"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestHotelRateAndRuleSearchRQ 测试酒店价格列表查询RQ
func TestHotelRateAndRuleSearchRQ() string {
	param := &HotelDetailsParam{
		Chain:        "PH",
		Code:         "95553",
		Name:         "REGAL AIRPORT HOTEL LIFESTYLE",
		Adults:       1,
		CheckinDate:  "2016-01-10",
		CheckoutDate: "2016-01-12",
		Provider:     "1G",
	}

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/HotelRateAndRuleSearchParam.txt", json)

	return testHotelDetailsRQ(param)
}

//TestHotelDescriptionRQ 测试酒店描述查询RQ
func TestHotelDescriptionRQ() string {
	param := &HotelDetailsParam{
		Chain:           "PH",
		Code:            "95553",
		Name:            "REGAL AIRPORT HOTEL LIFESTYLE",
		Adults:          1,
		CheckinDate:     "2016-01-10",
		CheckoutDate:    "2016-01-12",
		OnlyDescription: 1,
		Provider:        "1G",
	}

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/HotelDescriptionParam.json", json)

	return testHotelDetailsRQ(param)
}

//testHotelDetailsRQ 测试酒店详情RQ
func testHotelDetailsRQ(param *HotelDetailsParam) string {
	if nil == param {
		return "invalid parameters"
	}

	result := HotelDetails(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestHotelDetailsRS 测试酒店详情RS
func TestHotelDetailsRS() *HotelDetailsResult {
	val, err := LoadFile("samples/HotelDetailsRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope hotproxy.HotelDetailsRspEnvelope
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
