package retrievepnr

import (
	"encoding/xml"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/universal/retrievepnr"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestRetrievePnrRQ 测试提取PNR RQ
func TestRetrievePnrRQ() string {
	param := &RetrievePnrParam{
		UrlCode: "9WCPPZ",
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/RetrievePnrParam.json", json)

	result := RetrievePnr(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestRetrievePnrRS 测试提取PNR RS
func TestRetrievePnrRS() *RetrievePnrResult {
	val, err := LoadFile("samples/RetrievePnrRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope uniproxy.RetrievePnrRspEnvelope
	xml.Unmarshal([]byte(val), &envelope)

	result, err := getResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}
