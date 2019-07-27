package retrievepnr

import (
	cp "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	. "github.com/redochen/demos/travelport-uapi/util"
)

//RetrievePnrResult 提取PNR结果
type RetrievePnrResult struct {
	cp.CreatePnrResult
}

//SetErrorCode 设置错误代码
func SetErrorCode(code ErrorCode) *RetrievePnrResult {
	result := &RetrievePnrResult{}
	result.SetErrorCode(code)
	return result
}

//SetErrorCode 设置错误代码
func (retrievePnrResult *RetrievePnrResult) SetErrorCode(code ErrorCode) *RetrievePnrResult {
	retrievePnrResult.BaseResult.SetErrorCode(code)
	return retrievePnrResult
}
