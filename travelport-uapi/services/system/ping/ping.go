package ping

import (
	"encoding/xml"
	"errors"

	"github.com/redochen/demos/travelport-uapi/soap"
	sysproxy "github.com/redochen/demos/travelport-uapi/soap/system"
	. "github.com/redochen/demos/travelport-uapi/util"
	CcFile "github.com/redochen/tools/file"
	CcFunc "github.com/redochen/tools/function"
	CcJson "github.com/redochen/tools/json"
)

//Ping PING测试
func Ping(payload string) (string, error) {
	defer CcFunc.CheckPanic()

	body := sysproxy.NewPingReqBody("", payload)
	if nil == body {
		return "", errors.New("failed to create PingReqBody")
	}

	reqEnvelope := soap.NewReqEnvelope(body)

	reqXML, err := xml.Marshal(reqEnvelope)
	if err != nil {
		return "", err
	}

	CcFile.DumpFile("PingReq.xml", string(reqXML))

	rspXML, err := PostRequest(nil, soap.SystemServiceName, reqXML)
	if err != nil {
		return "", err
	}

	CcFile.DumpFile("PingRsq.xml", rspXML)

	var rspEnvelope sysproxy.PingRspEnvelope
	xml.Unmarshal([]byte(rspXML), &rspEnvelope)

	json, _ := CcJson.Serialize(rspEnvelope)
	CcFile.DumpFile("PingJson.txt", json)

	return json, nil
}
