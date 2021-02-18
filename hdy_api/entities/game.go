package entities

import (
	"time"
)

//TableName 游戏表名称
func (e *GameEntity) TableName() string {
	return "game"
}

//GameEntity 游戏表
type GameEntity struct {
	ID        int64     `xorm:"pk autoincr index"`           //自增ID
	Name      string    `xorm:"notnull nvarchar(20) unique"` //游戏名称
	CreatedAt time.Time `xorm:"notnull created"`             //创建时间
}

//TableName 游戏大区表名称
func (e *GameAreaEntity) TableName() string {
	return "game_area"
}

//GameAreaEntity 游戏大区
type GameAreaEntity struct {
	ID        int64     `xorm:"pk autoincr index"`    //自增ID
	Name      string    `xorm:"notnull nvarchar(20)"` //游戏大区名称
	GameID    int64     `xorm:"notnull index"`        //游戏ID
	CreatedAt time.Time `xorm:"notnull created"`      //创建时间
}

//TableName 游戏服务器表名称
func (e *GameServerEntity) TableName() string {
	return "game_server"
}

//GameServerEntity 游戏服务器
type GameServerEntity struct {
	ID         int64     `xorm:"pk autoincr index"`    //自增ID
	Name       string    `xorm:"notnull nvarchar(50)"` //游戏服名称
	GameAreaID int64     `xorm:"notnull index"`        //游戏大区ID
	CreatedAt  time.Time `xorm:"notnull created"`      //创建时间
}

//TableName 游戏段位表名称
func (e *GameDanEntity) TableName() string {
	return "game_dan"
}

//GameDanEntity 游戏段位
type GameDanEntity struct {
	ID        int64     `xorm:"pk autoincr index"`    //自增ID
	Name      string    `xorm:"notnull nvarchar(20)"` //段位名称
	GameID    int64     `xorm:"notnull index"`        //游戏ID
	CreatedAt time.Time `xorm:"notnull created"`      //创建时间
}

//TableName 游戏英雄表名称
func (e *GameHeroEntity) TableName() string {
	return "game_hero"
}

//GameHeroEntity 游戏英雄
type GameHeroEntity struct {
	ID        int64     `xorm:"pk autoincr index"`    //自增ID
	Name      string    `xorm:"notnull nvarchar(50)"` //英雄名称
	GameID    int64     `xorm:"notnull index"`        //游戏ID
	CreatedAt time.Time `xorm:"notnull created"`      //创建时间
}

//TableName 游戏角色表名称
func (e *GameRoleEntity) TableName() string {
	return "game_role"
}

//GameRoleEntity 游戏角色
type GameRoleEntity struct {
	ID           int64     `xorm:"pk autoincr index"`         //自增ID
	GUID         string    `xorm:"notnull varchar(32) index"` //惟一标识
	Name         string    `xorm:"notnull nvarchar(50)"`      //角色名称
	UserID       int64     `xorm:"notnull index"`             //用户ID
	GameServerID int64     `xorm:"notnull index"`             //游戏服ID
	IsGoodAt     bool      `xorm:"notnull default(0) index"`  //是否擅长
	CreatedAt    time.Time `xorm:"notnull created"`           //创建时间
}
