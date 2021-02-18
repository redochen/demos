package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redochen/demos/hdy_api/biz"
	"github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/demos/hdy_api/utils"
	CcFunc "github.com/redochen/tools/function"
)

//RegisterAsync 注册接口
func RegisterAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	var err error
	user := &models.UserModel{}

	if err = ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResult(status.CustomError, err.Error()))
		return
	}

	captchaID, captchaCode := utils.GetCaptchaParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(u *models.UserModel, id, code string, c chan<- interface{}) {
		c <- register(u, id, code)
	}(user, captchaID, captchaCode, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "RegisterAsync")
}

//register 注册功能
func register(user *models.UserModel, captchaID, captchaCode string) interface{} {
	defer CcFunc.CheckPanic()

	if nil == user {
		return models.NewResult(status.CustomError, "invalid parameter")
	}

	mustHasCode := false

	//有手机号码时必须要有验证码
	if user.Cellphone != "" {
		mustHasCode = true
	}

	//检测验证码
	err := utils.VerifyCaptcha(captchaID, captchaCode, false, mustHasCode)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	guid, err := biz.Register(user)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewRegisterResult(guid)
}

//LoginAsync 登录接口
func LoginAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	openID := ctx.Query("openid")
	account := ctx.Query("account")
	password := ctx.Query("password")

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(acc, pass, id string, c chan<- interface{}) {
		c <- login(acc, pass, id)
	}(account, password, openID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "LoginAsync")
}

//login 登录功能
func login(account, password, openid string) interface{} {
	defer CcFunc.CheckPanic()

	user, err := biz.Login(account, password, openid)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewUserResult(user)
}

//UpdateUserAsync 更新用户接口
func UpdateUserAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	var err error
	user := &models.UserModel{}

	if err = ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResult(status.CustomError, err.Error()))
		return
	}

	captchaID, captchaCode := utils.GetCaptchaParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(u *models.UserModel, id, code string, c chan<- interface{}) {
		c <- register(u, id, code)
	}(user, captchaID, captchaCode, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateUserAsync")

}

//updateUser 更新用户功能
func updateUser(user *models.UserModel, captchaID, captchaCode string) interface{} {
	defer CcFunc.CheckPanic()

	//检测验证码
	err := utils.VerifyCaptcha(captchaID, captchaCode, false, false)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	err = biz.UpdateUser(user)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewResult(status.Success)
}

//GetUserAsync 获取用户详情接口
func GetUserAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	userGUID := utils.GetUserGUIDParameter(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(guid string, c chan<- interface{}) {
		c <- getUser(guid)
	}(userGUID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetUserAsync")
}

//getUser 获取用户详情功能
func getUser(guid string) interface{} {
	defer CcFunc.CheckPanic()

	user, err := biz.GetUser(guid)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewUserResult(user)
}

//GetUsersAsync 获取用户列表接口
func GetUsersAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

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

//getUsers 获取用户列表功能
func getUsers(pageIndex, pageSize int) interface{} {
	defer CcFunc.CheckPanic()

	users, totalCount, pageCount, err := biz.GetUsers(pageIndex, pageSize)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResultEx(users, totalCount, pageCount)
}
