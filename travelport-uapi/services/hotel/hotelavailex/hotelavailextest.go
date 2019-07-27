package hotelavailex

import (
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavailex"
	ccfile "github.com/redochen/tools/file"
	. "github.com/redochen/tools/json"
)

//TestHotelAvailExRQ 测试酒店查询RQ
func TestHotelAvailExRQ() string {
	param := &HotelAvailExParam{}
	param.Location = "HKG"
	param.CheckinDate = "2016-01-10"
	param.CheckoutDate = "2016-01-12"

	json, _ := CcJson.Serialize(param)
	ccfile.DumpFile("samples/HotelAvailExParam.json", json)

	result := HotelAvailEx(param)
	val, err := CcJson.Serialize(result)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return val
}
