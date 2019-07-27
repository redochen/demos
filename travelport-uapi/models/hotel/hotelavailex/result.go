package hotelavailex

import (
	"github.com/redochen/demos/travelport-uapi/models"
	hotav "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	hotmi "github.com/redochen/demos/travelport-uapi/models/hotel/hotelmedia"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//HotelAvailEx接口结果
type HotelAvailExResult struct {
	models.BaseResult
	Hotels    []*HotelInfoExResult   `xml:"hotels" json:"hotels"` //酒店列表
	HostToken *hotav.HostTokenResult `xml:"hostToken" json:"hostToken"`
}

//酒店信息
type HotelInfoExResult struct {
	Vendors      []*hotav.VendorInfoResult      `xml:"vendors" json:"vendors"`       //供应商列表
	Properties   []*hotmi.HotelPropertyExResult `xml:"properties" json:"properties"` //属性列表
	Discounts    []*hotav.CorporateDiscount     `xml:"discounts" json:"discounts"`   //折扣列表
	Rates        []*hotav.RateInfoResult        `xml:"rates" json:"rates"`           //价格列表
	Media        *hotav.MediaInfoResult         `xml:"media" json:"media"`           //媒介信息
	SourceLink   bool                           `xml:"sourceLink" json:"sourceLink"` //指示数据是来自供应商还是数据库
	Description  string                         `xml:"description" json:"description"`
	ProviderCode string                         `xml:"providerCode" json:"providerCode"`
}

//根据HotelAvailResult填充结果
func (this *HotelAvailExResult) FillinByHotelAvailResult(result *hotav.HotelAvailResult) {
	if nil == result {
		return
	}

	if result.Status != 0 {
		this.Status = result.Status
		this.Message = result.Message
	}

	if result.HostToken != nil {
		this.HostToken = result.HostToken
	}

	if result.Hotels != nil &&
		len(result.Hotels) > 0 {
		this.Hotels = make([]*HotelInfoExResult, 0)

		for _, h := range result.Hotels {
			hotel := &HotelInfoExResult{
				Vendors:      h.Vendors,
				Discounts:    h.Discounts,
				Rates:        h.Rates,
				Media:        h.Media,
				SourceLink:   h.SourceLink,
				Description:  h.Description,
				ProviderCode: h.ProviderCode,
			}

			if h.Properties != nil &&
				len(h.Properties) > 0 {
				hotel.Properties = make([]*hotmi.HotelPropertyExResult, 0)
				for _, p := range h.Properties {
					property := &hotmi.HotelPropertyExResult{
						Property: p,
					}
					hotel.Properties = append(hotel.Properties, property)
				}
			}

			this.Hotels = append(this.Hotels, hotel)
		}
	}
}

//根据HotelMediaResult填充结果
func (this *HotelAvailExResult) FillinByHotelMediaResult(result *hotmi.HotelMediaResult) {
	if nil == result {
		return
	}

	if result.Status != 0 {
		this.Status = result.Status
		this.Message = result.Message
	}

	if result.HotelProperties != nil &&
		len(result.HotelProperties) > 0 {
		for _, p := range result.HotelProperties {
			this.setHotelPropertyMedias(p)
		}

	}
}

//设置酒店媒介信息
func (this *HotelAvailExResult) setHotelPropertyMedias(property *hotmi.HotelPropertyExResult) {
	if nil == property {
		return
	}

	p := this.getHotelPropertyExResult(property.Property.Chain, property.Property.Code)
	if p != nil {
		p.Medias = property.Medias
	}
}

//获取HotelPropertyExResult
func (this *HotelAvailExResult) getHotelPropertyExResult(chain, code string) *hotmi.HotelPropertyExResult {
	if nil == this.Hotels ||
		len(this.Hotels) <= 0 {
		return nil
	}

	for _, h := range this.Hotels {
		if nil == h.Properties ||
			len(h.Properties) <= 0 {
			continue
		}

		for _, p := range h.Properties {
			if nil == p ||
				nil == p.Property {
				continue
			}

			if p.Property.Chain == chain &&
				p.Property.Code == code {
				return p
			}
		}
	}

	return nil
}

//设置错误代码
func SetErrorCode(code ErrorCode) *HotelAvailExResult {
	result := &HotelAvailExResult{}
	result.SetErrorCode(code)
	return result
}

//设置错误代码
func (this *HotelAvailExResult) SetErrorCode(code ErrorCode) *HotelAvailExResult {
	this.BaseResult.SetErrorCode(code)
	return this
}
