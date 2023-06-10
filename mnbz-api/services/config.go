package services

import (
	"github.com/redochen/demos/mnbz-api/access"
	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/status"
	"github.com/redochen/demos/mnbz-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	CcFunc "github.com/redochen/tools/function"
)

// @Tags 配置
// @Summary 添加配置
// @Description 添加配置接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Config true "配置参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/config/add [post]
func AddConfigAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	config := &models.Config{}

	if err = ctx.BindJSON(config); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	config.UserID = userID

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Config, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.AddConfig(m)
		handleResult(err1, c)
	}(config, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "AddConfigAsync")
}

// @Tags 配置
// @Summary 更新配置
// @Description 更新配置接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Config true "配置参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/config/update [post]
func UpdateConfigAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	config := &models.Config{}

	if err = ctx.BindJSON(config); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Config, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.UpdateConfig(m, uid)
		handleResult(err1, c)
	}(config, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateConfigAsync")
}

// @Tags 配置
// @Summary 删除配置
// @Description 删除配置接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Config true "配置参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/config/delete [post]
func DeleteConfigAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	config := &models.Config{}

	if err = ctx.BindJSON(config); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Config, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.DeleteConfig(m, uid)
		handleResult(err1, c)
	}(config, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "DeleteConfigAsync")
}

// @Tags 配置
// @Summary 配置详情
// @Description 获取配置详情接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param id query int false "配置ID"
// @Param guid query string false "配置GUID"
// @Success 0 {object} models.Config
// @Failure 1 {object} models.BaseResult
// @Router /api/config/details [get]
func GetConfigAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	configId := utils.GetIDParameter(ctx)
	configGuid := utils.GetGUIDParameter(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, guid string, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		model := &models.Config{
			ID:   int32(id),
			GUID: guid,
		}
		config, err := access.GetConfig(model, uid)
		handleDataResult(config, err, c)
	}(configId, configGuid, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetConfigAsync")
}

// @Tags 配置
// @Summary 配置列表
// @Description 获取配置列表接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} []models.Config
// @Failure 1 {object} models.BaseResult
// @Router /api/configs [get]
func GetConfigsAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	pageIndex, pageSize := utils.GetPageParameters(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(index, size int, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		items, totalCount, pageCount, err := access.GetConfigs(pageIndex, pageSize, uid)
		handleListResult(items, totalCount, pageCount, err, c)
	}(pageIndex, pageSize, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetConfigsAsync")
}
