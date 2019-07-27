package models

//游戏
type GameModel struct {
	Id   int64  `json:"id"`   //自增ID
	Name string `json:"name"` //游戏名称
}

//游戏大区
type AreaModel struct {
	Id     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //游戏大区名称
	GameId int64  `json:"gameId"` //游戏ID
}

//游戏服务器
type ServerModel struct {
	Id     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //游戏服名称
	AreaId int64  `json:"areaId"` //游戏大区ID
}

//游戏段位
type DanModel struct {
	Id     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //段位名称
	GameId int64  `json:"gameId"` //游戏ID
}

//游戏英雄
type HeroModel struct {
	Id     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //英雄名称
	GameId int64  `json:"gameId"` //游戏ID
}
