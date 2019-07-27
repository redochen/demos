package access

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/redochen/demos/hdy_api/config"
	. "github.com/redochen/demos/hdy_api/entities"
	. "github.com/redochen/tools/function"
	"log"
	"time"
)

var engine *xorm.Engine

func init() {
	go initDb()
}

func initDb() {
	db, err := xorm.NewEngine(config.DataDriver, config.DataSource)
	if nil == db {
		log.Fatalf("[accsess.init] Can NOT connect to mysql", err)
	}

	db.Sync2(new(GameEntity))
	db.Sync2(new(GameAreaEntity))
	db.Sync2(new(GameServerEntity))
	db.Sync2(new(GameDanEntity))
	db.Sync2(new(GameHeroEntity))
	db.Sync2(new(GameRoleEntity))

	db.Sync2(new(UserEntity))
	db.Sync2(new(RelationEntity))
	db.Sync2(new(InvitationEntity))
	db.Sync2(new(EvaluationEntity))

	engine = db

	keepAlive()
}

func keepAlive() {
	t := time.NewTicker(time.Minute * 1)

	go func() {
		for _ = range t.C {
			defer CheckPanic()

			if engine != nil {
				engine.Ping()
			}
		}
	}()
}
