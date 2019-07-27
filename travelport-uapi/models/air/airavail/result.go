package airavail

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//AirAvailResult AirAvail接口结果
type AirAvailResult struct {
	models.BaseResult
	Flights []*Flight `xml:"flights" json:"flights"` //航班列表
}

//Flight ...
type Flight struct {
	Segments []*Segment `xml:"segments" json:"segments"` //航段列表
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *AirAvailResult {
	result := &AirAvailResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (airAvailResult *AirAvailResult) SetErrorCode(code ErrorCode) *AirAvailResult {
	airAvailResult.BaseResult.SetErrorCode(code)
	return airAvailResult
}
