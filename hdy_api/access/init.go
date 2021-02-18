package access

import (
	"log"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/redochen/demos/hdy_api/config"
	"github.com/redochen/demos/hdy_api/entities"
	CcFunc "github.com/redochen/tools/function"
)

var engine *xorm.Engine

func init() {
	go initDb()
}

func initDb() {
	db, err := xorm.NewEngine(config.DataDriver, config.DataSource)
	if nil == db {
		log.Fatalf("[accsess.init] Can NOT connect to mysql:\n%v", err)
	}

	db.Sync2(new(entities.GameEntity))
	db.Sync2(new(entities.GameAreaEntity))
	db.Sync2(new(entities.GameServerEntity))
	db.Sync2(new(entities.GameDanEntity))
	db.Sync2(new(entities.GameHeroEntity))
	db.Sync2(new(entities.GameRoleEntity))

	db.Sync2(new(entities.UserEntity))
	db.Sync2(new(entities.RelationEntity))
	db.Sync2(new(entities.InvitationEntity))
	db.Sync2(new(entities.EvaluationEntity))

	engine = db

	keepAlive()
}

func keepAlive() {
	t := time.NewTicker(time.Minute * 1)

	go func() {
		for range t.C {
			defer CcFunc.CheckPanic()

			if engine != nil {
				engine.Ping()
			}
		}
	}()
}
