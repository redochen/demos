package ping

import (
	"encoding/xml"
	sysproxy "github.com/redochen/demos/travelport-uapi/soap/system"
	. "github.com/redochen/demos/travelport-uapi/util"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestPing 测试Ping
func TestPing() (string, error) {
	val, err := LoadFile("PingRsq.xml")
	if err != nil {
		return "", err
	}

	var envelope sysproxy.PingRspEnvelope
	xml.Unmarshal([]byte(val), &envelope)

	json, _ := CcJson.Serialize(envelope)
	ccfile.DumpFile("json.txt", json)

	return json, nil
}
