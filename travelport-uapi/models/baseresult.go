package models

import (
	"github.com/redochen/demos/travelport-uapi/util"
	"strings"
)

//BaseResult 接口结果基类
type BaseResult struct {
	Status     int    `xml:"status" json:"status"`         //0, 成功；其他，失败。
	Message    string `xml:"message" json:"message"`       //提示信息，长度小于64。
	IsTimedOut bool   `xml:"isTimedOut" json:"isTimedOut"` //是否超时
}

//SetErrorCode 设置错误代码
func (baseResult *BaseResult) SetErrorCode(code util.ErrorCode) {
	baseResult.Status = code.Value()
	baseResult.Message = code.String()

	if code == util.ErrTimeout {
		baseResult.IsTimedOut = true
	}
}

//SetErrorMessage 设置错误消息
func (baseResult *BaseResult) SetErrorMessage(message string) {
	baseResult.Status = int(util.ErrOther)
	baseResult.Message = message

	if strings.Contains(message, "Timed out") {
		baseResult.IsTimedOut = true
	}
}
