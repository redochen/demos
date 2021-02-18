package services

import (
	"github.com/gin-gonic/gin"
	"github.com/redochen/demos/hdy_api/biz"
	"github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/demos/hdy_api/utils"
	CcFunc "github.com/redochen/tools/function"
	CcStr "github.com/redochen/tools/string"
)

//GamesAsync 获取游戏列表接口
func GamesAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(c chan<- interface{}) {
		c <- getGames()
	}(ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GamesAsync")
}

//getGames 获取游戏列表功能
func getGames() interface{} {
	defer CcFunc.CheckPanic()

	games, err := biz.GetGames()
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResult(games)
}

//GameAreasAsync 获取游戏大区列表接口
func GameAreasAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	gameID := CcStr.ParseInt(ctx.Query("gameID"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- getGameAreas(id)
	}(gameID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameAreasAsync")
}

//getGameAreas 获取游戏大区列表功能
func getGameAreas(gameID int) interface{} {
	defer CcFunc.CheckPanic()

	areas, err := biz.GetGameAreas(gameID)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResult(areas)
}

//GameServersAsync 获取游戏服务器列表接口
func GameServersAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	areaID := CcStr.ParseInt(ctx.Query("areaId"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- gameGameServers(id)
	}(areaID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameServersAsync")
}

//gameGameServers 获取游戏服务器列表功能
func gameGameServers(areaID int) interface{} {
	defer CcFunc.CheckPanic()

	servers, err := biz.GetGameServers(areaID)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResult(servers)
}

//GameDansAsync 获取游戏段位列表接口
func GameDansAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	gameID := CcStr.ParseInt(ctx.Query("gameID"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- getGameDans(id)
	}(gameID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameDansAsync")
}

//getGameDans 获取游戏段位列表功能
func getGameDans(gameID int) interface{} {
	defer CcFunc.CheckPanic()

	dans, err := biz.GetGameDans(gameID)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResult(dans)
}

//GameHeroesAsync 获取游戏英雄列表接口
func GameHeroesAsync(ctx *gin.Context) {
	defer CcFunc.CheckPanic()

	gameID := CcStr.ParseInt(ctx.Query("gameID"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- getGameHeroes(id)
	}(gameID, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameHeroesAsync")
}

//getGameHeroes 获取游戏英雄列表功能
func getGameHeroes(gameID int) interface{} {
	defer CcFunc.CheckPanic()

	heroes, err := biz.GetGameHeroes(gameID)
	if err != nil {
		return models.NewResult(status.CustomError, err.Error())
	}

	return models.NewListResult(heroes)
}
