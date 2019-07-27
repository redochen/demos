package models

import (
	"github.com/redochen/demos/hdy_api/status"
	. "github.com/redochen/tools/object"
)

//模型基类
type BaseModel struct {
	Guid      string `json:"guid,omitempty"`       //惟一标识
	CreatedAt string `json:"createTime,omitempty"` //创建时间，格式为:yyyy-MM-dd HH:mm:ss
	UpdatedAt string `json:"updateTime,omitempty"` //修改时间，格式为:yyyy-MM-dd HH:mm:ss
}

//响应基类
type BaseResult struct {
	ErrorCode    int    `json:"errCode"` //错误代码，0为成功
	ErrorMessage string `json:"errMsg"`  //错误消息
}

//设置错误
func (this *BaseResult) SetError(errCode status.ErrorCode, errMsg ...string) {
	this.ErrorCode = int(errCode)
	this.ErrorMessage = status.GetErrMessage(errCode, errMsg...)
}

//获取结果
func NewResult(errCode status.ErrorCode, errMsg ...string) *BaseResult {
	result := &BaseResult{}
	result.SetError(errCode, errMsg...)

	return result
}

//列表数据
type ListResult struct {
	BaseResult
	TotalCount int64       `json:"totalCount,omitempty"` //总共项数
	PageCount  int64       `json:"pageCount,omitempty"`  //总共页数
	Items      interface{} `json:"list,omitempty"`       //列表
}

//获取列表结果
func NewListResult(list interface{}) *ListResult {
	pageCount := 0
	totalCount := CcObject.GetLengthOfCollection(list)
	if totalCount > 0 {
		pageCount = 1
	}

	return NewListResultEx(list, int64(totalCount), int64(pageCount))
}

//获取列表结果
func NewListResultEx(list interface{}, totalCount, pageCount int64) *ListResult {
	result := &ListResult{
		TotalCount: totalCount,
		PageCount:  pageCount,
		Items:      list,
	}

	result.SetError(status.Success)
	return result
}
