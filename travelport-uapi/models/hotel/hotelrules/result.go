package hotelrules

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//HotelRules接口结果
type HotelRulesResult struct {
	models.BaseResult
	Rates      []*HotelRateDetail `xml:"rates" json:"rates"`           //价格列表
	Rules      []*HotelRuleItem   `xml:"rules" json:"rules"`           //规则列表
	SourceLink bool               `xml:"sourceLink" json:"sourceLink"` //数据来自供应商还是数据库
}

//设置错误代码
func SetErrorCode(code ErrorCode) *HotelRulesResult {
	result := &HotelRulesResult{}
	result.SetErrorCode(code)
	return result
}

//设置错误代码
func (this *HotelRulesResult) SetErrorCode(code ErrorCode) *HotelRulesResult {
	this.BaseResult.SetErrorCode(code)
	return this
}
