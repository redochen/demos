package hoteldetails

import (
	"fmt"

	"github.com/redochen/demos/travelport-uapi/models"
	CcTime "github.com/redochen/tools/time"
)

//HotelDetails接口参数
type HotelDetailsParam struct {
	models.BaseParam
	Chain           string          `xml:"chain" json:"chain"`                                         //酒店连锁代码
	Code            string          `xml:"code json:"code"`                                            //酒店代码
	Name            string          `xml:"name,omitempty" json:"name,omitempty"`                       //酒店名称
	CheckinDate     string          `xml:"checkin" json:"checkin"`                                     //入住日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	CheckoutDate    string          `xml:"checkout" json:"checkout"`                                   //退房日期，格式为“yyyy-MM-dd”或“yyyyMMdd”
	Adults          int             `xml:"adults,omitempty" json:"adults,omitempty"`                   //成人数量，默认为1
	Rooms           int             `xml:"rooms,omitempty" json:"rooms,omitempty"`                     //房间数量，默认为1
	Currency        string          `xml:"currency" json:"currency"`                                   //货币代码
	HostToken       *HostTokenParam `xml:"hostToken,omitempty" json:"hostToken,omitempty"`             //主机口令
	Aggregator      string          `xml:"aggregator,omitempty" json:"aggregator,omitempty"`           //集成者代码，即酒店平台代码
	Provider        string          `xml:"provider,omitempty" json:"provider,omitempty"`               //1G,1P,1P,1J,TRM等
	OnlyDescription int             `xml:"onlyDescription,omitempty" json:"onlyDescription,omitempty"` //仅返回描述：0-Complete，接口变为HotelDescription，其他-None，接口变为HotelRateAndRule
}

//主机口令。TRM使用，由Hotel Search接口返回，有效期为15分钟
type HostTokenParam struct {
	Host  string `xml:"host" json:"host"`
	Key   string `xml:"key" json:"key"`
	Token string `xml:"token" json:"token"`
}

//预检查
func (this *HotelDetailsParam) PreCheck() {
	this.ServiceName = "HotelDetails"

	if this.Chain == "" {
		this.Chain = "00"
	}

	if this.Adults <= 0 {
		this.Adults = 1
	}

	if this.Currency == "" {
		this.Currency = "CNY"
	}

	if this.LogContext == "" {
		this.LogContext = this.GetLogContext()
	}

	this.BaseParam.PreCheck()
}

//获取日志上下文
func (this *HotelDetailsParam) GetLogContext() string {
	var logContext string

	logContext += fmt.Sprintf("[%s-%s]", this.Chain, this.Code)
	//logContext += fmt.Sprintf("[%s]", this.Name)
	logContext += fmt.Sprintf("[%s-%s]",
		CcTime.RemoveDateSeparator(this.CheckinDate),
		CcTime.RemoveDateSeparator(this.CheckoutDate))

	if this.OnlyDescription != 0 {
		logContext += "[None]"
	} else {
		logContext += "[Complete]"
	}

	return logContext
}
