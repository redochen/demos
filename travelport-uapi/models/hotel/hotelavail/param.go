package hotelavail

import (
	"fmt"

	"github.com/redochen/demos/travelport-uapi/models"
	CcTime "github.com/redochen/tools/time"
)

//HotelAvail接口参数
type HotelAvailParam struct {
	models.BaseParam
	HotelAvailExternParam
	Location     string `xml:"location" json:"location"`                     //地点
	LocationType int    `xml:"locationType" json:"locationType"`             //地点类型：1-城市；2-机场。
	Name         string `xml:"name" json:"name"`                             //酒店名称
	CheckinDate  string `xml:"checkin" json:"checkin"`                       //入住日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	CheckoutDate string `xml:"checkout" json:"checkout"`                     //退房日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	Rooms        int    `xml:"rooms,omitempty" json:"rooms,omitempty"`       //房间数量，默认为1
	Adults       int    `xml:"adults,omitempty" json:"adults,omitempty"`     //成人数量，默认为1
	Children     int    `xml:"children,omitempty" json:"children,omitempty"` //儿童数量，默认为0
}

//HotelAvail接口扩展参数
type HotelAvailExternParam struct {
	Reference     string   `xml:"reference" json:"reference"`                             //著名地点，比如 Eiffel Tower
	IncludeHotels []string `xml:"includeHotels,omitempty" json:"includeHotels,omitempty"` //包含的酒店列表
	ExcludeHotels []string `xml:"excludeHotels,omitempty" json:"excludeHotels,omitempty"` //排除的酒店列表
	RatingType    string   `xml:"ratingType" json:"ratingType"`                           //星级类型：AAA或NTM
	Ratings       []int    `xml:"ratings" json:"ratings"`                                 //星级列表（与星级区间二选一）
	MinRating     int      `xml:"minRating" json:"minRating"`                             //最小星级（与星级列表二选一）
	MaxRating     int      `xml:"maxRating" json:"maxRating"`                             //最大星级（与星级列表二选一）
	Categories    []int    `xml:"categories" json:"categories"`                           //价格类型
	Amenities     []int    `xml:"amenities" json:"amenities"`                             //房间特性
	Currency      string   `xml:"currency" json:"currency"`                               //货币代码
}

//预检查
func (this *HotelAvailParam) PreCheck() {
	this.ServiceName = "HotelAvail"

	if this.Adults <= 0 {
		this.Adults = 1
	}

	if this.Rooms <= 0 {
		this.Rooms = 1
	}

	if this.LogContext == "" {
		this.LogContext = this.GetLogContext()
	}

	this.BaseParam.PreCheck()
}

//获取日志上下文
func (this *HotelAvailParam) GetLogContext() string {
	var logContext string

	logContext += fmt.Sprintf("[%s-%d]", this.Location, this.LocationType)
	/*
		if len(this.Name) > 0 {
			logContext += fmt.Sprintf("[%s]", this.Name)
		}
	*/

	logContext += fmt.Sprintf("[%s-%s]",
		CcTime.RemoveDateSeparator(this.CheckinDate),
		CcTime.RemoveDateSeparator(this.CheckoutDate))

	logContext += fmt.Sprintf("[%d]", this.Rooms)

	if this.Children > 0 {
		logContext += fmt.Sprintf("[%d-%d]", this.Adults, this.Children)
	} else {
		logContext += fmt.Sprintf("[%d-0]", this.Adults)
	}

	if len(this.RatingType) > 0 {
		if this.Ratings != nil &&
			len(this.Ratings) > 0 {
			logContext += fmt.Sprintf("[%s-", this.RatingType)
			for _, r := range this.Ratings {
				logContext += fmt.Sprintf("%d/", r)
			}
			logContext += "]"
		} else {

			if this.MinRating <= 0 {
				this.MinRating = 1
			}

			if this.MaxRating > 5 {
				this.MaxRating = 5
			}

			logContext += fmt.Sprintf("[%s-%d^%d]", this.RatingType, this.MinRating, this.MaxRating)
		}
	}

	return logContext
}
