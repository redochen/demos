package entities

import (
	"time"
)

//游戏表名称
func (self *GameEntity) TableName() string {
	return "game"
}

//游戏表
type GameEntity struct {
	Id        int64     `xorm:"pk autoincr index"`           //自增ID
	Name      string    `xorm:"notnull nvarchar(20) unique"` //游戏名称
	CreatedAt time.Time `xorm:"notnull created"`             //创建时间
}

//游戏大区表名称
func (self *GameAreaEntity) TableName() string {
	return "game_area"
}

//游戏大区
type GameAreaEntity struct {
	Id        int64     `xorm:"pk autoincr index"`    //自增ID
	Name      string    `xorm:"notnull nvarchar(20)"` //游戏大区名称
	GameId    int64     `xorm:"notnull index"`        //游戏ID
	CreatedAt time.Time `xorm:"notnull created"`      //创建时间
}

//游戏服务器表名称
func (self *GameServerEntity) TableName() string {
	return "game_server"
}

//游戏服务器
type GameServerEntity struct {
	Id         int64     `xorm:"pk autoincr index"`    //自增ID
	Name       string    `xorm:"notnull nvarchar(50)"` //游戏服名称
	GameAreaId int64     `xorm:"notnull index"`        //游戏大区ID
	CreatedAt  time.Time `xorm:"notnull created"`      //创建时间
}

//游戏段位表名称
func (self *GameDanEntity) TableName() string {
	return "game_dan"
}

//游戏段位
type GameDanEntity struct {
	Id        int64     `xorm:"pk autoincr index"`    //自增ID
	Name      string    `xorm:"notnull nvarchar(20)"` //段位名称
	GameId    int64     `xorm:"notnull index"`        //游戏ID
	CreatedAt time.Time `xorm:"notnull created"`      //创建时间
}

//游戏英雄表名称
func (self *GameHeroEntity) TableName() string {
	return "game_hero"
}

//游戏英雄
type GameHeroEntity struct {
	Id        int64     `xorm:"pk autoincr index"`    //自增ID
	Name      string    `xorm:"notnull nvarchar(50)"` //英雄名称
	GameId    int64     `xorm:"notnull index"`        //游戏ID
	CreatedAt time.Time `xorm:"notnull created"`      //创建时间
}

//游戏角色表名称
func (self *GameRoleEntity) TableName() string {
	return "game_role"
}

//游戏角色
type GameRoleEntity struct {
	Id           int64     `xorm:"pk autoincr index"`         //自增ID
	Guid         string    `xorm:"notnull varchar(32) index"` //惟一标识
	Name         string    `xorm:"notnull nvarchar(50)"`      //角色名称
	UserId       int64     `xorm:"notnull index"`             //用户ID
	GameServerId int64     `xorm:"notnull index"`             //游戏服ID
	IsGoodAt     bool      `xorm:"notnull default(0) index"`  //是否擅长
	CreatedAt    time.Time `xorm:"notnull created"`           //创建时间
}
