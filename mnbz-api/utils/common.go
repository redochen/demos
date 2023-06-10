package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/redochen/demos/mnbz-api/status"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/redochen/tools/cache"
	"github.com/redochen/tools/log"
	CcStr "github.com/redochen/tools/string"
)

const (
	defaultTimeoutSeconds = 30
	defaultPageIndex      = 1
	defaultPageSize       = 10

	AllStatus             = 10000
	DefaultDateTimeFormat = "yyyy-MM-dd HH:mm:ss"

	paramTimeoutSeconds = "timeoutSeconds"
	paramPageIndex      = "pageIndex"
	paramPageSize       = "pageSize"
	paramCaptchaID      = "captchaId"
	paramCaptchaCode    = "captchaCode"
	paramCaptchaLength  = "length"
	paramID             = "id"
	paramGUID           = "guid"
	paramSessionID      = "sid"
)

// ParseDateTime 解析日期时间
func ParseDateTime(str string, format ...string) time.Time {
	fmt := GetDateTimeFormat(format...)
	return CcStr.ParseTime(str, fmt, false)
}

// FormatDateTime 输出日期时间
func FormatDateTime(t time.Time, format ...string) string {
	fmt := GetDateTimeFormat(format...)
	return CcStr.FormatTime(t, fmt)
}

// GetDateTimeFormat 获取日期时间格式
func GetDateTimeFormat(format ...string) string {
	fmt := DefaultDateTimeFormat

	if len(format) > 0 {
		fmt = CcStr.FirstValid(format...)
	}

	return fmt
}

// GetOffsetAndLimit 获取分页查询的offset和limit
func GetOffsetAndLimit(pageIndex, pageSize int) (offset, limit int) {
	if pageSize <= 0 {
		pageSize = defaultPageSize
	}

	if pageIndex <= 0 {
		pageIndex = defaultPageIndex
	}

	limit = pageSize
	offset = (pageIndex - 1) * pageSize

	return
}

// GetRequestParameter 获取请求参数
func GetRequestParameter(ctx *gin.Context, name string) string {
	if nil == ctx || name == "" {
		return ""
	}

	return CcStr.FirstValid(ctx.Query(name), ctx.PostForm(name), ctx.GetHeader(name))
}

// GetTimeoutParameter 获取超时秒数参数
func GetTimeoutParameter(ctx *gin.Context) (seconds int) {
	seconds = CcStr.ParseInt(GetRequestParameter(ctx, paramTimeoutSeconds))
	if seconds <= 0 {
		seconds = defaultTimeoutSeconds
	}

	return
}

// GetPageParameters 获取分页参数
func GetPageParameters(ctx *gin.Context) (pageIndex int, pageSize int) {
	pageIndex = CcStr.ParseInt(GetRequestParameter(ctx, paramPageIndex))
	pageSize = CcStr.ParseInt(GetRequestParameter(ctx, paramPageSize))

	return
}

// GetCaptchaParameters 获取验证码参数
func GetCaptchaParameters(ctx *gin.Context) (captchaID string, captchaCode string) {
	captchaID = GetRequestParameter(ctx, paramCaptchaID)
	captchaCode = GetRequestParameter(ctx, paramCaptchaCode)

	return
}

// GetCaptchaLengthParameter 获取验证码长度参数
func GetCaptchaLengthParameter(ctx *gin.Context) (length int) {
	length = CcStr.ParseInt(GetRequestParameter(ctx, paramCaptchaLength))
	if length <= 0 {
		length = 6
	}

	return
}

// GetIDParameter 获取Id参数
func GetIDParameter(ctx *gin.Context) (Id int) {
	Id = CcStr.ParseInt(GetRequestParameter(ctx, paramID))
	return
}

// GetGUIDParameter 获取Guid参数
func GetGUIDParameter(ctx *gin.Context) (Guid string) {
	Guid = GetRequestParameter(ctx, paramGUID)
	return
}

// GetSessionParameter 获取Sid参数
func getSessionParameter(ctx *gin.Context) (Sid string) {
	Sid = GetRequestParameter(ctx, paramSessionID)
	return
}

// WaitAndResponse 等待响应结果
func WaitAndResponse(ctx *gin.Context, ch <-chan interface{}, name string) {
	var result interface{}

	start := time.Now()
	seconds := time.Duration(GetTimeoutParameter(ctx))

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(seconds * time.Second):
		//result = NewResult(status.Timeout)
		log.Errorf("[%s] %s", name, status.GetErrMessage(status.Timeout))
		break
	}

	elapsed := time.Since(start)
	log.Debugf("[%s] spent %f seconds", name, elapsed.Seconds())

	ctx.JSON(http.StatusOK, result)
}

// VerifyCaptcha 检验验证码
func VerifyCaptcha(captchaID, captchaCode string, mustHasID, mustHasCode bool) error {
	if captchaCode == "" {
		if mustHasCode {
			return errors.New("invalid captchaCode")
		}

		return nil
	}

	if captchaID != "" {
		if ok := captcha.VerifyString(captchaID, captchaCode); ok {
			return nil
		}

		return errors.New("wrong captchaCode")
	} else if mustHasID {
		return errors.New("invalid captchaId")
	}

	//读取缓存
	if value := cache.GetString(captchaCode); value != "" {
		return nil
	}

	return errors.New("wrong captchaCode")
}

// NewGUID 获取新的GUID
func NewGUID() string {
	guid, _ := CcStr.NewGUID()
	if guid != "" {
		guid = CcStr.ReplaceAll(guid, "", "-")
	}

	return guid
}

// NewRequiredError 获取字段必填错误实例
func NewRequiredError(field string) error {
	return fmt.Errorf("%s is required", field)
}

// NewInvalidError 获取无效的参数错误实例
func NewInvalidError(parameter string) error {
	return fmt.Errorf("invalid %s parameter", parameter)
}

// NewExistedError 获取数据已存在错误实例
func NewExistedError(name string) error {
	return fmt.Errorf("%s already exists", name)
}

// NewNotExistedError 获取数据不存在错误实例
func NewNotExistedError(name string) error {
	return fmt.Errorf("%s not exists", name)
}

// NewFailedError 获取操作失败错误实例
func NewFailedError(name, operation string) error {
	return fmt.Errorf("failed to %s %s", operation, name)
}

// NewNotAllowedError 获取非法操作错误实例
func NewNotAllowedError(name, operation string) error {
	return fmt.Errorf("not allowed to %s %s", operation, name)
}
