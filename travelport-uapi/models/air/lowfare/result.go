package lowfare

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//LowFareResult LowFareSearch接口结果
type LowFareResult struct {
	models.BaseResult
	Fares []*FlightFare `xml:"fares" json:"fares"` //运价列表
}

//FlightFare 运价结果类
type FlightFare struct {
	Segments []*Segment     `xml:"segments" json:"segments"` //航段列表
	Prices   []*PricingInfo `xml:"prices" json:"prices"`     //价格列表
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *LowFareResult {
	result := &LowFareResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (lowFareResult *LowFareResult) SetErrorCode(code ErrorCode) *LowFareResult {
	lowFareResult.BaseResult.SetErrorCode(code)
	return lowFareResult
}
