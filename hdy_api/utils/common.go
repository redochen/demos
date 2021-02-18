package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	. "github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
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
	paramUserGUID       = "userGuid"
)

//ParseDateTime 解析日期时间
func ParseDateTime(str string, format ...string) time.Time {
	fmt := GetDateTimeFormat(format...)
	return CcStr.ParseTime(str, fmt, false)
}

//FormatDateTime 输出日期时间
func FormatDateTime(t time.Time, format ...string) string {
	fmt := GetDateTimeFormat(format...)
	return CcStr.FormatTime(t, fmt)
}

//GetDateTimeFormat 获取日期时间格式
func GetDateTimeFormat(format ...string) string {
	fmt := DefaultDateTimeFormat

	if format != nil && len(format) > 0 {
		fmt = CcStr.FirstValid(format...)
	}

	return fmt
}

//GetOffsetAndLimit 获取分页查询的offset和limit
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

//GetRequestParameter 获取请求参数
func GetRequestParameter(ctx *gin.Context, name string) string {
	if nil == ctx || name == "" {
		return ""
	}

	return CcStr.FirstValid(ctx.Query(name), ctx.PostForm(name))
}

//GetTimeoutParameter 获取超时秒数参数
func GetTimeoutParameter(ctx *gin.Context) (seconds int) {
	seconds = CcStr.ParseInt(GetRequestParameter(ctx, paramTimeoutSeconds))
	if seconds <= 0 {
		seconds = defaultTimeoutSeconds
	}

	return
}

//GetPageParameters 获取分页参数
func GetPageParameters(ctx *gin.Context) (pageIndex int, pageSize int) {
	pageIndex = CcStr.ParseInt(GetRequestParameter(ctx, paramPageIndex))
	pageSize = CcStr.ParseInt(GetRequestParameter(ctx, paramPageSize))

	return
}

//GetCaptchaParameters 获取验证码参数
func GetCaptchaParameters(ctx *gin.Context) (captchaID string, captchaCode string) {
	captchaID = GetRequestParameter(ctx, paramCaptchaID)
	captchaCode = GetRequestParameter(ctx, paramCaptchaCode)

	return
}

//GetCaptchaLengthParameter 获取验证码长度参数
func GetCaptchaLengthParameter(ctx *gin.Context) (length int) {
	length = CcStr.ParseInt(GetRequestParameter(ctx, paramCaptchaLength))
	if length <= 0 {
		length = 6
	}

	return
}

//GetUserGUIDParameter 获取用户Guid参数
func GetUserGUIDParameter(ctx *gin.Context) (userGUID string) {
	userGUID = GetRequestParameter(ctx, paramUserGUID)

	return
}

//WaitAndResponse 等待响应结果
func WaitAndResponse(ctx *gin.Context, ch <-chan interface{}, name string) {
	var result interface{}

	start := time.Now()
	seconds := time.Duration(GetTimeoutParameter(ctx))

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(seconds * time.Second):
		result = NewResult(status.Timeout)
		log.Errorf("[%s] %s", name, status.GetErrMessage(status.Timeout))
		break
	}

	elapsed := time.Since(start)
	log.Debugf("[%s] spent %f seconds", name, elapsed.Seconds())

	ctx.JSON(http.StatusOK, result)
}

//VerifyCaptcha 检验验证码
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
