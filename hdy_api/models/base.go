package models

import (
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/tools/object"
)

//BaseModel 模型基类
type BaseModel struct {
	GUID      string `json:"guid,omitempty"`       //惟一标识
	CreatedAt string `json:"createTime,omitempty"` //创建时间，格式为:yyyy-MM-dd HH:mm:ss
	UpdatedAt string `json:"updateTime,omitempty"` //修改时间，格式为:yyyy-MM-dd HH:mm:ss
}

//BaseResult 响应基类
type BaseResult struct {
	ErrorCode    int    `json:"errCode"` //错误代码，0为成功
	ErrorMessage string `json:"errMsg"`  //错误消息
}

//SetError 设置错误
func (r *BaseResult) SetError(errCode status.ErrorCode, errMsg ...string) {
	r.ErrorCode = int(errCode)
	r.ErrorMessage = status.GetErrMessage(errCode, errMsg...)
}

//NewResult 获取结果
func NewResult(errCode status.ErrorCode, errMsg ...string) *BaseResult {
	result := &BaseResult{}
	result.SetError(errCode, errMsg...)

	return result
}

//ListResult 列表数据
type ListResult struct {
	BaseResult
	TotalCount int64       `json:"totalCount,omitempty"` //总共项数
	PageCount  int64       `json:"pageCount,omitempty"`  //总共页数
	Items      interface{} `json:"list,omitempty"`       //列表
}

//NewListResult 获取列表结果
func NewListResult(list interface{}) *ListResult {
	pageCount := 0
	totalCount := object.GetLengthOfCollection(list)
	if totalCount > 0 {
		pageCount = 1
	}

	return NewListResultEx(list, int64(totalCount), int64(pageCount))
}

//NewListResultEx 获取列表结果
func NewListResultEx(list interface{}, totalCount, pageCount int64) *ListResult {
	result := &ListResult{
		TotalCount: totalCount,
		PageCount:  pageCount,
		Items:      list,
	}

	result.SetError(status.Success)
	return result
}
