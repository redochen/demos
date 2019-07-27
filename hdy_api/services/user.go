package services

import (
	"github.com/gin-gonic/gin"
	"github.com/redochen/demos/hdy_api/biz"
	. "github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/demos/hdy_api/utils"
	. "github.com/redochen/tools/function"
	"net/http"
)

//注册接口
func RegisterAsync(ctx *gin.Context) {
	defer CheckPanic()

	var err error
	user := &UserModel{}

	if err = ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, NewResult(status.CustomError, err.Error()))
		return
	}

	captchaId, captchaCode := utils.GetCaptchaParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(u *UserModel, id, code string, c chan<- interface{}) {
		c <- register(u, id, code)
	}(user, captchaId, captchaCode, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "RegisterAsync")
}

//注册功能
func register(user *UserModel, captchaId, captchaCode string) interface{} {
	defer CheckPanic()

	if nil == user {
		return NewResult(status.CustomError, "invalid parameter")
	}

	mustHasCode := false

	//有手机号码时必须要有验证码
	if user.Cellphone != "" {
		mustHasCode = true
	}

	//检测验证码
	err := utils.VerifyCaptcha(captchaId, captchaCode, false, mustHasCode)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	guid, err := biz.Register(user)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewRegisterResult(guid)
}

//登录接口
func LoginAsync(ctx *gin.Context) {
	defer CheckPanic()

	openId := ctx.Query("openid")
	account := ctx.Query("account")
	password := ctx.Query("password")

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(acc, pass, id string, c chan<- interface{}) {
		c <- login(acc, pass, id)
	}(account, password, openId, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "LoginAsync")
}

//注册功能
func login(account, password, openid string) interface{} {
	defer CheckPanic()

	user, err := biz.Login(account, password, openid)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	} else {
		return NewUserResult(user)
	}
}

//更新用户接口
func UpdateUserAsync(ctx *gin.Context) {
	defer CheckPanic()

	var err error
	user := &UserModel{}

	if err = ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, NewResult(status.CustomError, err.Error()))
		return
	}

	captchaId, captchaCode := utils.GetCaptchaParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(u *UserModel, id, code string, c chan<- interface{}) {
		c <- register(u, id, code)
	}(user, captchaId, captchaCode, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateUserAsync")

}

//更新用户功能
func updateUser(user *UserModel, captchaId, captchaCode string) interface{} {
	defer CheckPanic()

	//检测验证码
	err := utils.VerifyCaptcha(captchaId, captchaCode, false, false)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	err = biz.UpdateUser(user)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewResult(status.Success)
}

//获取用户详情接口
func GetUserAsync(ctx *gin.Context) {
	defer CheckPanic()

	userGuid := utils.GetUserGuidParameter(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(guid string, c chan<- interface{}) {
		c <- getUser(guid)
	}(userGuid, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetUserAsync")
}

//获取用户详情功能
func getUser(guid string) interface{} {
	defer CheckPanic()

	user, err := biz.GetUser(guid)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	} else {
		return NewUserResult(user)
	}
}

//获取用户列表接口
func GetUsersAsync(ctx *gin.Context) {
	defer CheckPanic()

	pageIndex, pageSize := utils.GetPageParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(index, size int, c chan<- interface{}) {
		c <- getUsers(index, size)
	}(pageIndex, pageSize, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetUsersAsync")
}

//获取用户列表功能
func getUsers(pageIndex, pageSize int) interface{} {
	defer CheckPanic()

	users, totalCount, pageCount, err := biz.GetUsers(pageIndex, pageSize)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	} else {
		return NewListResultEx(users, totalCount, pageCount)
	}
}
