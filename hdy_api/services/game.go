package services

import (
	"github.com/gin-gonic/gin"
	"github.com/redochen/demos/hdy_api/biz"
	. "github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/demos/hdy_api/utils"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/string"
)

//获取游戏列表接口
func GamesAsync(ctx *gin.Context) {
	defer CheckPanic()

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(c chan<- interface{}) {
		c <- getGames()
	}(ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GamesAsync")
}

//获取游戏列表功能
func getGames() interface{} {
	defer CheckPanic()

	games, err := biz.GetGames()
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewListResult(games)
}

//获取游戏大区列表接口
func GameAreasAsync(ctx *gin.Context) {
	defer CheckPanic()

	gameId := CcStr.ParseInt(ctx.Query("gameId"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- getGameAreas(id)
	}(gameId, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameAreasAsync")
}

//获取游戏大区列表功能
func getGameAreas(gameId int) interface{} {
	defer CheckPanic()

	areas, err := biz.GetGameAreas(gameId)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewListResult(areas)
}

//获取游戏服务器列表接口
func GameServersAsync(ctx *gin.Context) {
	defer CheckPanic()

	areaId := CcStr.ParseInt(ctx.Query("areaId"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- gameGameServers(id)
	}(areaId, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameServersAsync")
}

//获取游戏服务器列表功能
func gameGameServers(areaId int) interface{} {
	defer CheckPanic()

	servers, err := biz.GetGameServers(areaId)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewListResult(servers)
}

//获取游戏段位列表接口
func GameDansAsync(ctx *gin.Context) {
	defer CheckPanic()

	gameId := CcStr.ParseInt(ctx.Query("gameId"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- getGameDans(id)
	}(gameId, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameDansAsync")
}

//获取游戏段位列表功能
func getGameDans(gameId int) interface{} {
	defer CheckPanic()

	dans, err := biz.GetGameDans(gameId)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewListResult(dans)
}

//获取游戏英雄列表接口
func GameHeroesAsync(ctx *gin.Context) {
	defer CheckPanic()

	gameId := CcStr.ParseInt(ctx.Query("gameId"))

	//创建一个chan用于接收异步处理结果
	ch := make(chan interface{}, 1)

	//异步执行
	go func(id int, c chan<- interface{}) {
		c <- getGameHeroes(id)
	}(gameId, ch)

	//等待异步处理结果并返回响应
	utils.WaitAndResponse(ctx, ch, "GameHeroesAsync")
}

//获取游戏英雄列表功能
func getGameHeroes(gameId int) interface{} {
	defer CheckPanic()

	heroes, err := biz.GetGameHeroes(gameId)
	if err != nil {
		return NewResult(status.CustomError, err.Error())
	}

	return NewListResult(heroes)
}
