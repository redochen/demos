package hotelavail

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/hotel"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//HotelAvail接口结果
type HotelAvailResult struct {
	models.BaseResult
	Hotels    []*HotelInfoResult `xml:"hotels" json:"hotels"` //酒店列表
	HostToken *HostTokenResult   `xml:"hostToken" json:"hostToken"`
}

//主机口令。TRM使用，由Hotel Search接口返回，有效期为15分钟
type HostTokenResult struct {
	Host  string `xml:"host" json:"host"`   //主机
	Key   string `xml:"key" json:"key"`     //键
	Token string `xml:"token" json:"token"` //口令
}

//酒店信息
type HotelInfoResult struct {
	Vendors      []*VendorInfoResult  `xml:"vendors" json:"vendors"`       //供应商列表
	Properties   []*HotelProperty     `xml:"properties" json:"properties"` //属性列表
	Discounts    []*CorporateDiscount `xml:"discounts" json:"discounts"`   //折扣列表
	Rates        []*RateInfoResult    `xml:"rates" json:"rates"`           //价格列表
	Media        *MediaInfoResult     `xml:"media" json:"media"`           //媒介信息
	SourceLink   bool                 `xml:"sourceLink" json:"sourceLink"` //指示数据是来自供应商还是数据库
	Description  string               `xml:"description" json:"description"`
	ProviderCode string               `xml:"providerCode" json:"providerCode"`
}

//公司折扣？
type CorporateDiscount struct {
	Value                string `xml:"value" json:"value"`
	IsNegotiatedRateCode bool   `xml:"isNegotiate" json:"isNegotiate"`
}

//媒介信息
type MediaInfoResult struct {
	Caption  string `xml:"caption" json:"caption"`   //标题
	Height   uint   `xml:"height" json:"height"`     //高度
	Width    uint   `xml:"width" json:"width"`       //宽度
	Type     string `xml:"type" json:"type"`         //类型
	Url      string `xml:"url" json:"url"`           //图片链接
	Icon     string `xml:"icon" json:"icon"`         //图标链接
	SizeCode string `xml:"sizeCode" json:"sizeCode"` //尺寸代码
}

//供应商信息
type VendorInfoResult struct {
	Key              string `xml:"key" json:"key"` //键值
	ProviderCode     string `xml:"providerCode" json:"providerCode"`
	VendorCode       string `xml:"vendorCode" json:"vendorCode"`
	VendorLocationID string `xml:"vendorLocationID" json:"vendorLocationID"`
}

//价格信息
type RateInfoResult struct {
	MinAmount                        float32 `xml:"minAmount" json:"minAmount"`                                               //原始最低价格
	MinAmountCurrency                string  `xml:"minAmountCurrency" json:"minAmountCurrency"`                               //原始最低价格货币代码
	ApproximateMinAmount             float32 `xml:"approximateMinAmount" json:"approximateMinAmount"`                         //转换最低价格
	ApproximateMinAmountCurrency     string  `xml:"approximateMinAmountCurrency" json:"approximateMinAmountCurrency"`         //转换最低价格货币代码
	MinAmountRateChanged             bool    `xml:"minAmountRateChanged" json:"minAmountRateChanged"`                         //最低价格是否会变价
	MaxAmount                        float32 `xml:"maxAmount" json:"maxAmount"`                                               //原始最高价格
	MaxAmountCurrency                string  `xml:"maxAmountCurrency" json:"maxAmountCurrency"`                               //原始最高价格货币代码
	ApproximateMaxAmount             float32 `xml:"approximateMaxAmount" json:"approximateMaxAmount"`                         //原始最高价格
	ApproximateMaxAmountCurrency     string  `xml:"approximateMaxAmountCurrency" json:"approximateMaxAmountCurrency"`         //原始最高价格货币代码
	MaxAmountRateChanged             bool    `xml:"maxAmountRateChanged" json:"maxAmountRateChanged"`                         //最高价格是否会变价
	MinStayAmount                    float32 `xml:"minStayAmount" json:"minStayAmount"`                                       //原始入住期间最低价格
	MinStayAmountCurrency            string  `xml:"minStayAmountCurrency" json:"minStayAmountCurrency"`                       //原始入住期间最低价格货币代码
	ApproximateMinStayAmount         float32 `xml:"approximateMinStayAmount" json:"approximateMinStayAmount"`                 //转换入住期间最低价格
	ApproximateMinStayAmountCurrency string  `xml:"approximateMinStayAmountCurrency" json:"approximateMinStayAmountCurrency"` //转换入住期间最低价格货币代码
	Commission                       string  `xml:"commission" json:"commission"`
	RateSupplier                     string  `xml:"rateSupplier" json:"rateSupplier"`
	RateSupplierLogo                 string  `xml:"rateSupplierLogo" json:"rateSupplierLogo"`
	PaymentType                      string  `xml:"paymentType" json:"paymentType"`
}

//设置错误代码
func SetErrorCode(code ErrorCode) *HotelAvailResult {
	result := &HotelAvailResult{}
	result.SetErrorCode(code)
	return result
}

//设置错误代码
func (this *HotelAvailResult) SetErrorCode(code ErrorCode) *HotelAvailResult {
	this.BaseResult.SetErrorCode(code)
	return this
}
