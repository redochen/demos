package config

import (
	"fmt"
	"log"

	"github.com/redochen/tools/config"
)

const (
	defaultServerPort = 1111
	defaultMysqlPort  = 3306

	dbDriverMysql    = "mysql"
	dbDriverPostgres = "postgres"
	dbDriverMsSql    = "mssql"
	dbDriverSQLite   = "sqlite3"

	//server配置节
	sectionSvr     = "server"
	optionSvrPort  = "port"
	optionSvrToken = "token"

	//database配置节
	sectionDb        = "database"
	optionDbDriver   = "driver"
	optionDbHost     = "host"
	optionDbPort     = "port"
	optionDbDbName   = "database"
	optionDbCharset  = "charset"
	optionDbUserName = "username"
	optionDbPassword = "password"

	//sms配置节
	sectionSms     = "sms"
	optionSmsURL   = "url"
	optionSmsToken = "token"
)

var (
	ServerPort  int
	ServerToken string

	DataDriver string
	DataSource string

	SmsUrl   string
	SmsToken string
)

func init() {
	if nil == config.Conf || !config.Conf.IsValid() {
		return
	}

	var err error

	ServerPort = config.Conf.IntEx(sectionSvr, optionSvrPort, defaultServerPort)
	ServerToken, _ = config.Conf.String(sectionSvr, optionSvrToken)

	if !config.Conf.HasSection(sectionDb) {
		log.Fatal("[config.init] Can NOT read database section\n")
	}

	DataDriver, err = config.Conf.String(sectionDb, optionDbDriver)
	if err != nil || DataDriver == "" {
		log.Fatalf("[config.init] Can NOT read database driver option:\n%v", err)
	}

	dbHost, _ := config.Conf.String(sectionDb, optionDbHost)
	// if err != nil || dbHost == "" {
	// 	log.Fatalf("[config.init] Can NOT read database host option:\n%v", err)
	// }

	dbPort, _ := config.Conf.Int(sectionDb, optionDbPort)
	if dbPort == 0 {
		dbPort = defaultMysqlPort
	}

	dbName, err := config.Conf.String(sectionDb, optionDbDbName)
	if err != nil || dbName == "" {
		log.Fatalf("[config.init] Can NOT read database name option:\n%v", err)
	}

	dbCharset, _ := config.Conf.String(sectionDb, optionDbCharset)

	dbUser, _ := config.Conf.String(sectionDb, optionDbUserName)
	// if err != nil || dbUser == "" {
	// 	log.Fatalf("[config.init] Can NOT read database username option:\n%v", err)
	// }

	dbPass, _ := config.Conf.String(sectionDb, optionDbPassword)
	// if err != nil || dbPass == "" {
	// 	log.Fatalf("[config.init] Can NOT read database password option:\n%v", err)
	// }

	switch DataDriver {
	case dbDriverMysql:
		DataSource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		if dbCharset != "" {
			DataSource = DataSource + "?charset=" + dbCharset
		}
	case dbDriverSQLite:
		DataSource = dbName
		break
	}

	if config.Conf.HasSection(sectionSms) {
		SmsUrl, _ = config.Conf.String(sectionSms, optionSmsURL)
		SmsToken, _ = config.Conf.String(sectionSms, optionSmsToken)
	}
}
