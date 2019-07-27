package config

import (
	"fmt"
	. "github.com/redochen/tools/config"
	"log"
)

const (
	defaultServerPort = 1111
	defaultMysqlPort  = 3306

	dbDriverMysql = "mysql"
	//dbDriverPostgres = "postgres"
	//dbDriverMsSql    = "mssql"
	//dbDriverSQLite   = "sqlite3"

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
	optionSmsUrl   = "url"
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
	if nil == Conf || !Conf.IsValid() {
		return
	}

	var err error

	ServerPort = Conf.IntEx(sectionSvr, optionSvrPort, defaultServerPort)
	ServerToken, _ = Conf.String(sectionSvr, optionSvrToken)

	if !Conf.HasSection(sectionDb) {
		log.Fatal("[config.init] Can NOT read database section\n")
	}

	DataDriver, err = Conf.String(sectionDb, optionDbDriver)
	if err != nil || DataDriver == "" {
		log.Fatalf("[config.init] Can NOT read database driver option\n", err)
	}

	dbHost, err := Conf.String(sectionDb, optionDbHost)
	if err != nil || dbHost == "" {
		log.Fatalf("[config.init] Can NOT read database host option\n", err)
	}

	dbPort, _ := Conf.Int(sectionDb, optionDbPort)
	if dbPort == 0 {
		dbPort = defaultMysqlPort
	}

	dbName, err := Conf.String(sectionDb, optionDbDbName)
	if err != nil || dbName == "" {
		log.Fatalf("[config.init] Can NOT read database name option\n", err)
	}

	dbCharset, _ := Conf.String(sectionDb, optionDbCharset)

	dbUser, err := Conf.String(sectionDb, optionDbUserName)
	if err != nil || dbUser == "" {
		log.Fatalf("[config.init] Can NOT read database username option\n", err)
	}

	dbPass, err := Conf.String(sectionDb, optionDbPassword)
	if err != nil || dbPass == "" {
		log.Fatalf("[config.init] Can NOT read database password option\n", err)
	}

	switch DataDriver {
	case dbDriverMysql:
		DataSource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		if dbCharset != "" {
			DataSource = DataSource + "?charset=" + dbCharset
		}
	}

	if Conf.HasSection(sectionSms) {
		SmsUrl, _ = Conf.String(sectionSms, optionSmsUrl)
		SmsToken, _ = Conf.String(sectionSms, optionSmsToken)
	}
}
