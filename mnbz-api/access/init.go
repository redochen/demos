package access

import (
	"errors"
	"log"
	"time"

	"github.com/redochen/demos/mnbz-api/config"
	"github.com/redochen/demos/mnbz-api/models"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-xorm/xorm"
	CcFunc "github.com/redochen/tools/function"
)

var engine *xorm.Engine

func init() {
	go initDb()
}

func initDb() {
	db, err := xorm.NewEngine(config.DataDriver, config.DataSource)
	if nil == db {
		log.Fatalf("[access.init] failed to connect to database:\n%v", err)
	}

	//db.ShowSQL(true)
	//db.Logger().SetLevel(xorm.LOG_DEBUG)
	//db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	//db.SetConnMaxLifetime()

	syncTable(db, new(models.Company))
	syncTable(db, new(models.User))
	syncTable(db, new(models.Config))
	syncTable(db, new(models.Need))
	syncTable(db, new(models.Resource))
	syncTable(db, new(models.Task))

	engine = db

	//keepAlive()
}

//checkEngine 检查数据库引擎
func checkEngine() error {
	if nil == engine {
		return errors.New("engine not initialized")
	}
	return nil
}

//syncTable 检查并同步表
func syncTable(db *xorm.Engine, bean interface{}) {
	exist, err := db.IsTableExist(bean)
	if err != nil {
		log.Fatalf("[access.syncTable] failed to check table:\n%v", err)
		return
	}

	if !exist {
		err = db.CreateTables(bean)
		if err != nil {
			log.Fatalf("[access.syncTable] failed to create table:\n%v", err)
			return
		}
	}

	db.Sync2(bean)
}

func keepAlive() {
	defer CcFunc.CheckPanic()
	t := time.NewTicker(time.Minute * 1)

	go func() {
		for range t.C {
			if engine != nil {
				engine.Ping()
			}
		}
	}()
}
