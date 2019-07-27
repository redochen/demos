package hotelmedia

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/models"
)

//HotelMedia接口参数
type HotelMediaParam struct {
	models.BaseParam
	HotelProperties []*HotelPropertyParam `xml:"hotelProperties" json:"hotelProperties"`
}

//酒店物业参数
type HotelPropertyParam struct {
	Chain              string   `xml:"chain" json:"chain"`                                               //酒店连锁代码
	Code               string   `xml:"code json:"code"`                                                  //酒店代码
	Name               string   `xml:"name,omitempty" json:"name,omitempty"`                             //酒店名称
	Location           string   `xml:"location,omitempty" json:"location,omitempty"`                     //酒店地点
	Transportation     int      `xml:"transportation,omitempty" json:"transportation,omitempty"`         //交通设施
	ParticipationLevel string   `xml:"participationLevel,omitempty" json:"participationLevel,omitempty"` //参与级别？
	ReserveRequirement string   `xml:"reserveRequirement,omitempty" json:"reserveRequirement,omitempty"` //预订要求
	VendorLocationKey  string   `xml:"vendorLocationKey,omitempty" json:"vendorLocationKey,omitempty"`   //供应商地址键值
	Address            []string `xml:"address,omitempty" json:"address,omitempty"`                       //地址列表
}

//预检查
func (this *HotelMediaParam) PreCheck() {
	this.ServiceName = "HotelMedia"

	if this.LogContext == "" {
		if this.HotelProperties != nil &&
			len(this.HotelProperties) > 0 {
			for i, hp := range this.HotelProperties {
				this.LogContext += fmt.Sprintf("[%s-%s]", hp.Chain, hp.Code)
				if i > 0 {
					this.LogContext += "~"
					break
				}
			}
		}
	}

	this.BaseParam.PreCheck()
}
