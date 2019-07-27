package hoteldetails

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	hotav "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//HotelDetails接口结果
type HotelDetailsResult struct {
	models.BaseResult
	Property            *HotelProperty           `xml:"property" json:"property"`                       //酒店属性
	AlternateProperties []*HotelProperty         `xml:"alternateProperties" json:"alternateProperties"` //可选酒店列表
	Items               []*HotelDetailItem       `xml:"items" json:"items"`                             //详情列表
	Rates               []*HotelRateDetail       `xml:"rates" json:"rates"`                             //价格列表
	Medias              []*hotav.MediaInfoResult `xml:"medias" json:"medias"`                           //媒介信息
	Comments            []*Comments              `xml:"comments" json:"comments"`                       //客户评论
	SourceLink          bool                     `xml:"sourceLink" json:"sourceLink"`                   //数据来自供应商还是数据库
}

//设置错误代码
func SetErrorCode(code ErrorCode) *HotelDetailsResult {
	result := &HotelDetailsResult{}
	result.SetErrorCode(code)
	return result
}

//设置错误代码
func (this *HotelDetailsResult) SetErrorCode(code ErrorCode) *HotelDetailsResult {
	this.BaseResult.SetErrorCode(code)
	return this
}
