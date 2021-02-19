package cancelpnr

import (
	"encoding/xml"
	"fmt"

	. "github.com/redochen/demos/travelport-uapi/models/universal/cancelpnr"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	CcFile "github.com/redochen/tools/file"
	CcJson "github.com/redochen/tools/json"
)

//TestCancelPnrRQ 测试取消PNR RQ
func TestCancelPnrRQ() string {
	param := &CancelPnrParam{
		UrlCode: "9WCPPZ",
		Version: "0",
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60

	json, _ := CcJson.Serialize(param)
	CcFile.DumpFile("samples/CancelPnrParam.json", json)

	result := CancelPnr(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}

//TestCancelPnrRS 测试取消PNR RS
func TestCancelPnrRS() *CancelPnrResult {
	val, err := LoadFile("samples/CancelPnrRS.xml")
	if err != nil {
		return SetErrorCode(ErrInvalidParameter)
	}

	var envelope uniproxy.CancelPnrRspEnvelope
	xml.Unmarshal([]byte(val), &envelope)

	result, err := getResult(envelope.Body)
	if err != nil {
		return SetErrorCode(ErrFormatResultError)
	}

	return result
}
