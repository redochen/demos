package hotelrules

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelrules"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestHotelRulesRQ 测试酒店媒介RQ
func TestHotelRulesRQ() string {
	param := &HotelRulesParam{
		//UrlCode: "",
		Chain:        "PH",
		Code:         "95553",
		Name:         "REGAL AIRPORT HOTEL LIFESTYLE",
		BaseAmount:   2400.00,
		Currency:     "HKD",
		RatePlanType: "B2TPAP",
		Adults:       1,
		CheckinDate:  "2016-01-10",
		CheckoutDate: "2016-01-12",
		Provider:     "1G",
		RequestType:  2,
	}

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/HotelRulesParam.json", json)

	result := HotelRules(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestHotelRulesRS 测试酒店媒介RS
func TestHotelRulesRS() *HotelRulesResult {
	val, err := LoadFile("samples/HotelRulesRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope hotproxy.HotelRulesRspEnvelope
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
