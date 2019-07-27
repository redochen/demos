package airprice

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/models/air"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//AirPriceResult AirPrice接口结果
type AirPriceResult struct {
	models.BaseResult
	Segments []*Segment     `xml:"segments" json:"segments"` //航段列表
	Prices   []*PriceResult `xml:"prices" json:"prices"`     //价格列表
}

//PriceResult 价格信息
type PriceResult struct {
	AirPricingSolutions []*AirPricingSolution `xml:"airPricingSolutions" json:"airPricingSolutions"` //机票价格解决方案列表
}

//AirPricingSolution 机票价格解决方案类
type AirPricingSolution struct {
	Key               string         `xml:"key" json:"key"`                             //引用键
	TotalPrice        float32        `xml:"totalPrice" json:"totalPrice"`               //总价
	TotalCurrency     string         `xml:"totalCurrency" json:"totalCurrency"`         //总价货币代码
	AppTotalPrice     float32        `xml:"appTotalPrice" json:"appTotalPrice"`         //近似总价
	AppTotalCurrency  string         `xml:"appTotalCurrency" json:"appTotalCurrency"`   //近似总价货币代码
	BasePrice         float32        `xml:"basePrice" json:"basePrice"`                 //基础价
	BaseCurrency      string         `xml:"baseCurrency" json:"baseCurrency"`           //基础价货币代码
	AppBasePrice      float32        `xml:"appBasePrice" json:"appBasePrice"`           //近似基础价
	AppBaseCurrency   string         `xml:"appBaseCurrency" json:"appBaseCurrency"`     //近似基础价货币代码
	EquivBasePrice    float32        `xml:"equivBasePrice" json:"equivBasePrice"`       //等效基础价
	EquivBaseCurrency string         `xml:"equivBaseCurrency" json:"equivBaseCurrency"` //等效基础价货币代码
	Taxes             float32        `xml:"taxes" json:"taxes"`                         //税费
	TaxesCurrency     string         `xml:"taxesCurrency" json:"taxesCurrency"`         //税费货币代码
	AppTaxes          float32        `xml:"appTaxes" json:"appTaxes"`                   //近似税费
	AppTaxesCurrency  string         `xml:"appTaxesCurrency" json:"appTaxesCurrency"`   //近似税费货币代码
	QuoteDate         string         `xml:"quoteDate" json:"quoteDate"`                 //询价日期
	PricingInfos      []*PricingInfo `xml:"pricingInfos" json:"pricingInfos"`           //价格信息列表
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *AirPriceResult {
	result := &AirPriceResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (airPriceResult *AirPriceResult) SetErrorCode(code ErrorCode) *AirPriceResult {
	airPriceResult.BaseResult.SetErrorCode(code)
	return airPriceResult
}

//SetErrorMessage 设置错误消息
func SetErrorMessage(message string) *AirPriceResult {
	result := &AirPriceResult{}
	result.SetErrorMessage(message)
	return result
}

//SetErrorMessage 设置错误消息
func (airPriceResult *AirPriceResult) SetErrorMessage(message string) *AirPriceResult {
	airPriceResult.BaseResult.SetErrorMessage(message)
	return airPriceResult
}
