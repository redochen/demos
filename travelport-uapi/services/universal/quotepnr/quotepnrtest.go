package quotepnr

import (
	"fmt"

	. "github.com/redochen/demos/travelport-uapi/models/universal/quotepnr"
	CcFile "github.com/redochen/tools/file"
	CcJson "github.com/redochen/tools/json"
)

//TestQuotePnrRQ 测试询价PNR RQ
func TestQuotePnrRQ() string {
	param := &QuotePnrParam{
		UrlCode:  "9WCPPZ",
		Carrier:  "CX",
		Currency: "CNY",
	}

	param.GdsAccount = "9UP"
	param.TimeoutSeconds = 60

	json, _ := CcJson.Serialize(param)
	CcFile.DumpFile("samples/QuotePnrParam.json", json)

	result := QuotePnr(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}
