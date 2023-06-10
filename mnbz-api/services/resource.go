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

// @Tags 资源
// @Summary 添加资源
// @Description 添加资源接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Resource true "资源参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/resource/add [post]
func AddResourceAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	resource := &models.Resource{}

	if err = ctx.BindJSON(resource); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	resource.UserID = userID

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Resource, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.AddResource(m)
		handleResult(err1, c)
	}(resource, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "AddResourceAsync")
}

// @Tags 资源
// @Summary 更新资源
// @Description 更新资源接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Resource true "资源参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/resource/update [post]
func UpdateResourceAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	resource := &models.Resource{}

	if err = ctx.BindJSON(resource); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Resource, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.UpdateResource(m, uid)
		handleResult(err1, c)
	}(resource, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateResourceAsync")
}

// @Tags 资源
// @Summary 删除资源
// @Description 删除资源接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Resource true "资源置参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/resource/delete [post]
func DeleteResourceAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	resource := &models.Resource{}

	if err = ctx.BindJSON(resource); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Resource, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.DeleteResource(m, uid)
		handleResult(err1, c)
	}(resource, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "DeleteResourceAsync")
}

// @Tags 资源
// @Summary 资源详情
// @Description 获取资源详情接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param id query int false "资源ID"
// @Param guid query string false "资源GUID"
// @Success 0 {object} models.Resource
// @Failure 1 {object} models.BaseResult
// @Router /api/resource/details [get]
func GetResourceAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	resourceId := utils.GetIDParameter(ctx)
	resourceGuid := utils.GetGUIDParameter(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, guid string, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		model := &models.Resource{
			ID:   int32(id),
			GUID: guid,
		}
		resource, err := access.GetResource(model, uid)
		handleDataResult(resource, err, c)
	}(resourceId, resourceGuid, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetResourceAsync")
}

// @Tags 资源
// @Summary 资源列表
// @Description 获取资源列表接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} []models.Resource
// @Failure 1 {object} models.BaseResult
// @Router /api/resources [get]
func GetResourcesAsync(ctx *gin.Context) {
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
		items, totalCount, pageCount, err := access.GetResources(pageIndex, pageSize, uid)
		handleListResult(items, totalCount, pageCount, err, c)
	}(pageIndex, pageSize, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetResourcesAsync")
}
