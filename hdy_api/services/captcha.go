package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/redochen/demos/hdy_api/config"
	. "github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/demos/hdy_api/utils"
	"github.com/redochen/tools/cache"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/http"
	rnd "github.com/redochen/tools/random"
	. "github.com/redochen/tools/string"
	"net/http"
	"path"
	"time"
)

const (
	smsSendOk     = "success"
	cacheDuration = 10 * time.Minute
)

//获取图片、音频验证码接口
func GetCaptcha(ctx *gin.Context) {
	defer CheckPanic()

	length := utils.GetCaptchaLengthParameter(ctx)
	code := captcha.NewLen(length)

	ctx.String(http.StatusOK, code)
}

//获取短信验证码接口
func GetSmsCaptchaAsync(ctx *gin.Context) {
	defer CheckPanic()

	cellphone := ctx.Query("cellphone")
	length := utils.GetCaptchaLengthParameter(ctx)
	captchaId, captchaCode := utils.GetCaptchaParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(l int, mobile, id, code string, c chan<- interface{}) {
		c <- getSmsCaptcha(l, mobile, id, code)
	}(length, cellphone, captchaId, captchaCode, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetSmsCaptchaAsync")
}

//获取短信验证码功能
func getSmsCaptcha(length int, cellphone, captchaId, captchaCode string) interface{} {
	defer CheckPanic()

	//检测验证码
	err := utils.VerifyCaptcha(captchaId, captchaCode, true, true)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	smsCaptchaCode := rnd.GetRandomNumber(length)

	//保存到缓存
	cache.SetString(smsCaptchaCode, cellphone, cacheDuration)

	err = sendSmsCaptcha(cellphone, smsCaptchaCode)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	} else {
		return NewResult(status.Success)
	}
}

//验证图片、音频验证码接口
func VerifyCaptcha(ctx *gin.Context) {
	defer CheckPanic()

	captchaId, captchaCode := utils.GetCaptchaParameters(ctx)
	err := utils.VerifyCaptcha(captchaId, captchaCode, true, true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResult(status.CustomError, err.Error()))
	} else {
		ctx.JSON(http.StatusOK, NewResult(status.Success))
	}
}

//加载图片、音频验证码资源接口
func LoadCaptcha(ctx *gin.Context) {
	defer CheckPanic()

	file := ctx.Param("file")
	width := CcStr.ParseInt(ctx.Query("width"))
	height := CcStr.ParseInt(ctx.Query("height"))
	lang := ctx.Query("lang")
	reload := ctx.Query("reload")
	download := ctx.Query("download")

	if width <= 0 {
		width = captcha.StdWidth
	}

	if height <= 0 {
		height = captcha.StdHeight
	}

	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		ctx.String(http.StatusNotFound, "")
		return
	}

	if reload != "" {
		captcha.Reload(id)
	}

	var contentType string
	var content bytes.Buffer

	switch ext {
	case ".png":
		contentType = "image/png"
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		contentType = "audio/x-wav"
		captcha.WriteAudio(&content, id, lang)
	default:
		ctx.String(http.StatusNotFound, "")
		return
	}

	if download != "" {
		contentType = "application/octet-stream"
	}

	ctx.Data(http.StatusOK, contentType, content.Bytes())
}

//发送短信验证码功能
func sendSmsCaptcha(cellphone, code string) error {
	defer CheckPanic()

	if config.SmsUrl == "" {
		return errors.New("invalid sms url")
	}

	if config.SmsToken == "" {
		return errors.New("invalid sms token")
	}

	url := fmt.Sprintf("%s?token=%s&cellphone=%s&captcha=%s",
		config.SmsUrl, config.SmsToken, cellphone, code)

	val, err := CcHttp.Get(url)
	if err != nil {
		return err
	}

	if val != smsSendOk {
		return errors.New(val)
	}

	return nil
}
