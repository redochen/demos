package config

import (
	"fmt"
	. "github.com/redochen/tools/config"
)

const (
	DefaultPort           string = "80" //监听端口
	DefaultTimeoutSeconds int    = 30   //接口超时秒数
	DefaultOutputToFile   int    = 0    //是否输出到文件

	//server配置节点
	sectionSvr    = "server"
	optionSvrPort = "port"
	optionTimeout = "timeoutSeconds"
	optionOutput  = "outputToFile"

	//endpoints配置节点
	sectionEndpoints = "endpoints"
	optionUserId     = "userId"
	optionPassword   = "password"
	optionBranchCode = "branchCode"
	optionEndPoint   = "endPoint"

	//log配置节点
	sectionLog = "log"
)

var (
	Port           = "80"  //监听端口
	TimeoutSeconds = 30    //接口超时秒数
	OutputToFile   = false //是否输出到文件
)

//初始化
func init() {
	if nil == Conf || !Conf.IsValid() {
		return
	}

	fmt.Println("loading travelport config...")

	Port = Conf.StringEx(sectionSvr, optionSvrPort, DefaultPort)
	TimeoutSeconds = Conf.IntEx(sectionSvr, optionTimeout, DefaultTimeoutSeconds)
	outputToFile := Conf.IntEx(sectionSvr, optionOutput, DefaultOutputToFile)
	OutputToFile = (outputToFile != 0)

	InitEndpoints()
	InitPccs()

	fmt.Println("loading travelport config done.")
}
