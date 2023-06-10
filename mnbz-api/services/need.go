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

// @Tags 需求
// @Summary 添加需求
// @Description 添加需求接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Need true "需求参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/need/add [post]
func AddNeedAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	need := &models.Need{}

	if err = ctx.BindJSON(need); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	need.UserID = userID

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Need, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.AddNeed(m)
		handleResult(err1, c)
	}(need, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "AddNeedAsync")
}

// @Tags 需求
// @Summary 更新需求
// @Description 更新需求接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Need true "需求参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/need/update [post]
func UpdateNeedAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	need := &models.Need{}

	if err = ctx.BindJSON(need); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Need, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.UpdateNeed(m, uid)
		handleResult(err1, c)
	}(need, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateNeedAsync")
}

// @Tags 需求
// @Summary 删除需求
// @Description 删除需求接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Need true "需求置参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/need/delete [post]
func DeleteNeedAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	need := &models.Need{}

	if err = ctx.BindJSON(need); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Need, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.DeleteNeed(m, uid)
		handleResult(err1, c)
	}(need, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "DeleteNeedAsync")
}

// @Tags 需求
// @Summary 需求详情
// @Description 获取需求详情接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param id query int false "需求ID"
// @Param guid query string false "需求GUID"
// @Success 0 {object} models.Need
// @Failure 1 {object} models.BaseResult
// @Router /api/need/details [get]
func GetNeedAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	needId := utils.GetIDParameter(ctx)
	needGuid := utils.GetGUIDParameter(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, guid string, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		model := &models.Need{
			ID:   int32(id),
			GUID: guid,
		}
		need, err := access.GetNeed(model, uid)
		handleDataResult(need, err, c)
	}(needId, needGuid, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetNeedAsync")
}

// @Tags 需求
// @Summary 需求列表
// @Description 获取需求列表接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} []models.Need
// @Failure 1 {object} models.BaseResult
// @Router /api/needs [get]
func GetNeedsAsync(ctx *gin.Context) {
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
		items, totalCount, pageCount, err := access.GetNeeds(pageIndex, pageSize, uid)
		handleListResult(items, totalCount, pageCount, err, c)
	}(pageIndex, pageSize, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetNeedsAsync")
}
