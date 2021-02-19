package config

import (
	"fmt"

	"github.com/redochen/tools/config"
)

const (
	//DefaultPort 监听端口
	DefaultPort string = "80"

	//DefaultTimeoutSeconds 接口超时秒数
	DefaultTimeoutSeconds int = 30

	//DefaultOutputToFile 是否输出到文件
	DefaultOutputToFile int = 0

	//server配置节点
	sectionSvr    = "server"
	optionSvrPort = "port"
	optionTimeout = "timeoutSeconds"
	optionOutput  = "outputToFile"

	//endpoints配置节点
	sectionEndpoints = "endpoints"
	optionUserID     = "userId"
	optionPassword   = "password"
	optionBranchCode = "branchCode"
	optionEndPoint   = "endPoint"

	//log配置节点
	sectionLog = "log"
)

var (
	//Port 监听端口
	Port = "80"

	//TimeoutSeconds 接口超时秒数
	TimeoutSeconds = 30

	//OutputToFile 是否输出到文件
	OutputToFile = false
)

//初始化
func init() {
	if nil == config.Conf || !config.Conf.IsValid() {
		return
	}

	fmt.Println("loading travelport config...")

	Port = config.Conf.StringEx(sectionSvr, optionSvrPort, DefaultPort)
	TimeoutSeconds = config.Conf.IntEx(sectionSvr, optionTimeout, DefaultTimeoutSeconds)
	outputToFile := config.Conf.IntEx(sectionSvr, optionOutput, DefaultOutputToFile)
	OutputToFile = (outputToFile != 0)

	InitEndpoints()
	InitPccs()

	fmt.Println("loading travelport config done.")
}
