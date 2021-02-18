package models

//GameModel 游戏
type GameModel struct {
	ID   int64  `json:"id"`   //自增ID
	Name string `json:"name"` //游戏名称
}

//AreaModel 游戏大区
type AreaModel struct {
	ID     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //游戏大区名称
	GameID int64  `json:"gameId"` //游戏ID
}

//ServerModel 游戏服务器
type ServerModel struct {
	ID     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //游戏服名称
	AreaID int64  `json:"areaId"` //游戏大区ID
}

//DanModel 游戏段位
type DanModel struct {
	ID     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //段位名称
	GameID int64  `json:"gameId"` //游戏ID
}

//HeroModel 游戏英雄
type HeroModel struct {
	ID     int64  `json:"id"`     //自增ID
	Name   string `json:"name"`   //英雄名称
	GameID int64  `json:"gameId"` //游戏ID
}
