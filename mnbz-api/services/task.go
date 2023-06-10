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

// @Tags 任务
// @Summary 添加任务
// @Description 添加任务接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Task true "任务参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/task/add [post]
func AddTaskAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	task := &models.Task{}

	if err = ctx.BindJSON(task); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	task.UserID = userID

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Task, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.AddTask(m)
		handleResult(err1, c)
	}(task, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "AddTaskAsync")
}

// @Tags 任务
// @Summary 更新任务
// @Description 更新任务接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Task true "任务参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/task/update [post]
func UpdateTaskAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	task := &models.Task{}

	if err = ctx.BindJSON(task); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Task, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.UpdateTask(m, uid)
		handleResult(err1, c)
	}(task, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "UpdateTaskAsync")
}

// @Tags 任务
// @Summary 删除任务
// @Description 删除任务接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param data body models.Task true "任务置参数"
// @Success 0 {object} models.BaseResult
// @Failure 1 {object} models.BaseResult
// @Router /api/task/delete [post]
func DeleteTaskAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	task := &models.Task{}

	if err = ctx.BindJSON(task); err != nil {
		ctx.JSON(http.StatusBadRequest,
			models.NewResult(status.CustomError, err.Error()))
		return
	}

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(m *models.Task, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		_, err1 := access.DeleteTask(m, uid)
		handleResult(err1, c)
	}(task, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "DeleteTaskAsync")
}

// @Tags 任务
// @Summary 任务详情
// @Description 获取任务详情接口
// @Produce json
// @Param sid header string true "会话ID"
// @Param id query int false "任务ID"
// @Param guid query string false "任务GUID"
// @Success 0 {object} models.Task
// @Failure 1 {object} models.BaseResult
// @Router /api/task/details [get]
func GetTaskAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//获取当前登录的用户ID
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResult(status.CustomError, err.Error()))
		return
	}

	taskId := utils.GetIDParameter(ctx)
	taskGuid := utils.GetGUIDParameter(ctx)

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, guid string, uid int32, c chan<- interface{}) {
		defer CcFunc.CheckPanic()
		model := &models.Task{
			ID:   int32(id),
			GUID: guid,
		}
		task, err := access.GetTask(model, uid)
		handleDataResult(task, err, c)
	}(taskId, taskGuid, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetTaskAsync")
}

// @Tags 任务
// @Summary 任务列表
// @Description 获取任务列表接口
// @Produce json
// @Param sid header string true "会话ID"
// @Success 0 {object} []models.Task
// @Failure 1 {object} models.BaseResult
// @Router /api/tasks [get]
func GetTasksAsync(ctx *gin.Context) {
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
		items, totalCount, pageCount, err := access.GetTasks(pageIndex, pageSize, uid)
		handleListResult(items, totalCount, pageCount, err, c)
	}(pageIndex, pageSize, userID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GetTasksAsync")
}
