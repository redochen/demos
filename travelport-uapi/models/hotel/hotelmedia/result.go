package hotelmedia

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	hotav "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//HotelMedia接口结果
type HotelMediaResult struct {
	models.BaseResult
	HotelProperties []*HotelPropertyExResult `xml:"hotelProperties" json:"hotelProperties"`
}

//酒店属性
type HotelPropertyExResult struct {
	Property *HotelProperty           `xml:"property" json:"property"`
	Medias   []*hotav.MediaInfoResult `xml:"medias" json:"medias"` //媒介信息
}

//设置错误代码
func SetErrorCode(code ErrorCode) *HotelMediaResult {
	result := &HotelMediaResult{}
	result.SetErrorCode(code)
	return result
}

//设置错误代码
func (this *HotelMediaResult) SetErrorCode(code ErrorCode) *HotelMediaResult {
	this.BaseResult.SetErrorCode(code)
	return this
}
