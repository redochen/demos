package services

import (
	"net/http"

	"github.com/redochen/demos/mnbz-api/access"
	"github.com/redochen/demos/mnbz-api/biz"
	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/status"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/gin-gonic/gin"
	CcFunc "github.com/redochen/tools/function"
)

// @Tags 用户
// @Summary 用户注册
// @Description 用户注册接口
// @Produce json
// @Param captchaId query string true "验证码ID"
// @Param captchaCode query string true "验证码"
// @Param data body models.RegisterModel true "注册参数"
// @Success 0 {object} models.RegisterResult
// @Failure 1 {object} models.BaseResult
// @Router /api/user/register [post]
func RegisterAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	var err error
	model := &models.RegisterModel{}

	if err = ctx.BindJSON(model); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	captchaId, captchaCode := utils.GetCaptchaParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.RegisterModel, id, code string, c chan<- interface{}) {
		c <- register(m, id, code)
	}(model, captchaId, captchaCode, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "RegisterAsync")
}

// register 注册功能
func register(model *models.RegisterModel, captchaId, captchaCode string) interface{} {
	defer CcFunc.CheckPanic()

	if nil == model {
		return models.NewResult(status.CustomError, "invalid parameter")
	}

	mustHasCode := false

	//有手机号码时必须要有验证码
	if model.Mobile != "" {
		mustHasCode = true
	}

	//检测验证码
	err := utils.VerifyCaptcha(captchaId, captchaCode, false, mustHasCode)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	userId, err := biz.Register(model)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewRegisterResult(userId)
}

// @Tags 用户
// @Summary 用户登录
// @Description 用户登录接口
// @Produce json
// @Param account query string false "账号，可选参数"
// @Param password query string false "密码，可选参数"
// @Param openid query string false "微信开放id，可选参数"
// @Success 0 {object} string
// @Failure 1 {object} models.BaseResult
// @Router /api/user/login [post]
func LoginAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	openID := utils.GetRequestParameter(ctx, "openid")
	account := utils.GetRequestParameter(ctx, "account")
	password := utils.GetRequestParameter(ctx, "password")

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(acc, pass, id string, c chan<- interface{}) {
		c <- login(acc, pass, id)
	}(account, password, openID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "LoginAsync")
}

// login 登录功能
func login(account, password, openid string) interface{} {
	defer CcFunc.CheckPanic()

	sid, err := biz.Login(account, password, openid)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewDataResult(sid)
}

// @Tags 用户
// @Summary 更新用户
// @Description 更新用户信息接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.User true "更新参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/user/update [post]
func UpdateUserAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	var err error
	user := &models.User{}

	if err = ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(u *models.User, uid int32, c chan<- interface{}) {
		c <- updateUser(u, uid)
	}(user, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateUserAsync")
}

// updateUser 更新用户功能
func updateUser(user *models.User, userID int32) interface{} {
	defer CcFunc.CheckPanic()

	//检测验证码
	// err := utils.VerifyCaptcha(captchaId, captchaCode, false, false)
	// if err != nil {
	// 	return models.NewResult(status.CustomError, err.Error())
	// }

	err := biz.UpdateUser(user, userID)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewResult(status.Success)
}

// @Tags 用户
// @Summary 用户详情
// @Description 获取用户详情接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} models.User
// @Failure 1 {object} models.BaseResult
// @Router /api/user/details [get]
func GetUserAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(uid int32, c chan<- interface{}) {
		c <- getUser(uid)
	}(userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetUserAsync")
}

// getUser 获取用户详情功能
func getUser(userID int32) interface{} {
	defer CcFunc.CheckPanic()

	user, err := access.GetUserByID(userID, true)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewDataResult(user)
}

// @Tags 用户
// @Summary 用户列表
// @Description 获取用户列表接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} []models.User
// @Failure 1 {object} models.BaseResult
// @Router /api/users [get]
func GetUsersAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	_, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

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

// getUsers 获取用户列表功能
func getUsers(pageIndex, pageSize int) interface{} {
	defer CcFunc.CheckPanic()

	users, totalCount, pageCount, err := access.GetUsers(pageIndex, pageSize, true)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResultEx(users, totalCount, pageCount)
}

// @Tags 用户
// @Summary 用户登出
// @Description 用户登接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/user/logout [post]
func LogoutAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	_, _ = utils.GetUserID(ctx)

	ctx.JSON(http.StatusOK, models.NewResult(status.Success))
}
