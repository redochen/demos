package hotelrules

import (
	"fmt"

	"github.com/redochen/demos/travelport-uapi/models"
	CcTime "github.com/redochen/tools/time"
)

//HotelRules接口参数
type HotelRulesParam struct {
	models.BaseParam
	UrlCode      string  `xml:"urlCode,omitempty" json:"urlCode,omitempty"`           //UR编号，当该参数有值时，将忽略其他所有参数
	Chain        string  `xml:"chain,omitempty" json:"chain,omitempty"`               //酒店连锁代码
	Code         string  `xml:"code,omitempty json:"code,omitempty"`                  //酒店代码
	Name         string  `xml:"name,omitempty" json:"name,omitempty"`                 //酒店名称
	BaseAmount   float32 `xml:"baseAmount,omitempty" json:"baseAmount,omitempty"`     //价格
	Currency     string  `xml:"currency,omitempty" json:"currency,omitempty"`         //货币代码
	RatePlanType string  `xml:"ratePlanType,omitempty" json:"ratePlanType,omitempty"` //运价编码
	CheckinDate  string  `xml:"checkin,omitempty" json:"checkin,omitempty"`           //入住日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	CheckoutDate string  `xml:"checkout,omitempty" json:"checkout,omitempty"`         //退房日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	Rooms        int     `xml:"rooms,omitempty" json:"rooms,omitempty"`               //房间数量
	Adults       int     `xml:"adults,omitempty" json:"adults,omitempty"`             //成人数量
	Children     int     `xml:"children,omitempty" json:"children,omitempty"`         //儿童数量
	Provider     string  `xml:"provider,omitempty" json:"provider,omitempty"`         //1G,1P,1P,1J,TRM等
	RequestType  int     `xml:"requestType,omitempty" json:"requestType,omitempty"`   //请求类型：1-仅详细，2-仅规则，0及其他-所有
}

//预检查
func (this *HotelRulesParam) PreCheck() {
	this.ServiceName = "HotelRules"

	if this.LogContext == "" {
		this.LogContext = this.GetLogContext()
	}

	this.BaseParam.PreCheck()
}

//获取日志上下文
func (this *HotelRulesParam) GetLogContext() string {
	var logContext string

	if len(this.UrlCode) > 0 {
		logContext += fmt.Sprintf("[%s]", this.UrlCode)
	} else {
		if len(this.Currency) <= 0 {
			this.Currency = "CNY"
		}

		logContext += fmt.Sprintf("[%s-%s]", this.Chain, this.Code)
		//logContext += fmt.Sprintf("[%s]", this.Name)
		logContext += fmt.Sprintf("[%s-%s]",
			CcTime.RemoveDateSeparator(this.CheckinDate),
			CcTime.RemoveDateSeparator(this.CheckoutDate))
		logContext += fmt.Sprintf("[%s-%s%.2f]", this.RatePlanType, this.Currency, this.BaseAmount)

		switch this.RequestType {
		case 1:
			logContext += "[Details]"
			break
		case 2:
			logContext += "[Rules]"
			break
		default:
			logContext += "[All]"
			break
		}
	}

	return logContext
}
