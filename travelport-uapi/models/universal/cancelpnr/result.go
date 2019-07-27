package cancelpnr

import (
	"github.com/redochen/demos/travelport-uapi/models"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//CancelPnrResult 取消PNR结果
type CancelPnrResult struct {
	models.BaseResult
	ResStatus []*ProviderReservationStatus `xml:"resStatus" json:"resStatus"` //状态列表
}

//ProviderReservationStatus 供应商预订状态
type ProviderReservationStatus struct {
	LocatorCode  string `xml:"locatorCode" json:"locatorCode"`   //预订编号
	Cancelled    bool   `xml:"cancelled" json:"cancelled"`       //是否已取消
	ProviderCode string `xml:"providerCode" json:"providerCode"` //供应商代码：1G
	Message      string `xml:"message" json:"message"`           //相关信息
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *CancelPnrResult {
	result := &CancelPnrResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (cancelPnrResult *CancelPnrResult) SetErrorCode(code ErrorCode) *CancelPnrResult {
	cancelPnrResult.BaseResult.SetErrorCode(code)
	return cancelPnrResult
}
