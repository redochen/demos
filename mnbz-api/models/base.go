package models

import (
	"github.com/redochen/demos/mnbz-api/status"

	"github.com/redochen/tools/object"
)

// BaseModel 模型基类
type BaseModel struct {
	CreatedAt string `json:"createdAt,omitempty"` //创建时间，格式为:yyyy-MM-dd HH:mm:ss
	UpdatedAt string `json:"updatedAt,omitempty"` //修改时间，格式为:yyyy-MM-dd HH:mm:ss
}

// BaseResult 响应基类
type BaseResult struct {
	Code    int    `json:"code"`    //错误代码，0为成功
	Message string `json:"message"` //错误消息
}

// SetError 设置错误
func (r *BaseResult) SetError(code status.ErrorCode, errMsg ...string) {
	r.Code = int(code)
	r.Message = status.GetErrMessage(code, errMsg...)
}

// NewResult 获取结果
func NewResult(code status.ErrorCode, errMsg ...string) *BaseResult {
	result := &BaseResult{}
	result.SetError(code, errMsg...)

	return result
}

// DataResult 单体数据
type DataResult struct {
	BaseResult
	Data interface{} `json:"data,omitempty"` //数据
}

// NewDataResult 获取数据结果
func NewDataResult(data interface{}) *DataResult {
	result := &DataResult{
		Data: data,
	}

	result.SetError(status.Success)
	return result
}

// ListResult 列表数据
type ListResult struct {
	BaseResult
	TotalCount int64       `json:"totalCount"`     //总共项数
	PageCount  int64       `json:"pageCount"`      //总共页数
	Items      interface{} `json:"list,omitempty"` //列表
}

// NewListResult 获取列表结果
func NewListResult(items interface{}) *ListResult {
	pageCount := 0
	totalCount := object.GetLengthOfCollection(items)
	if totalCount > 0 {
		pageCount = 1
	}

	return NewListResultEx(items, int64(totalCount), int64(pageCount))
}

// NewListResultEx 获取列表结果
func NewListResultEx(items interface{}, totalCount, pageCount int64) *ListResult {
	result := &ListResult{
		TotalCount: totalCount,
		PageCount:  pageCount,
		Items:      items,
	}

	result.SetError(status.Success)
	return result
}
