package lowfare

import (
	"encoding/xml"
	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/air/lowfare"
	"github.com/redochen/demos/travelport-uapi/soap"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/json"
	. "github.com/redochen/tools/log"
	"time"
)

//LowFareAsync 异步运价查询接口
func LowFareAsync(param *LowFareParam) *LowFareResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *LowFareResult, 1)

	//异步执行
	go func(p *LowFareParam, c chan<- *LowFareResult) {
		c <- LowFare(p)
	}(param, ch)

	var result *LowFareResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		Logger.Error(param.LogContext, " low search timed out!!!")
		break
	}

	elapsed := time.Since(start)
	Logger.Infof("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//LowFare 运价查询接口
func LowFare(param *LowFareParam) *LowFareResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	//获取PCC配置
	pcc, err := cfg.PCCs.Get(param.GdsAccount)
	if err != nil {
		return SetErrorCode(ErrInvalidPcc)
	}

	if cfg.OutputToFile {
		paramJSON, _ := CcJson.Serialize(param)
		if paramJSON != "" {
			DumpFile(param.LogContext+"_Param.json", paramJSON, false)
		}
	}

	//转换查询参数
	reqEnvelope, err := getReqEnvolpe(param, pcc.BranchCode)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrInvalidParameter)
	}

	//序列化请求XML文本
	reqXML, err := xml.Marshal(reqEnvelope)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrParseParameterError)
	}

	DumpFile(param.LogContext+"_RQ.xml", string(reqXML), false)

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.AirServiceName, reqXML)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrProcessError)
	}

	DumpFile(param.LogContext+"_RS.xml", string(rspXML), false)

	//反序列化响应XML文件
	var rspEnvelope airproxy.LowFareRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	//转换查询结果
	result, err := getResult(rspEnvelope.Body)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrNoFareDataReturned)
	}

	if cfg.OutputToFile {
		resultJSON, _ := CcJson.Serialize(result)
		if resultJSON != "" {
			DumpFile(param.LogContext+"_Result.json", resultJSON, false)
		}
	}

	return result
}
