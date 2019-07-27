package cancelpnr

import (
	"encoding/xml"
	"fmt"
	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/universal/cancelpnr"
	"github.com/redochen/demos/travelport-uapi/soap"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/json"
	. "github.com/redochen/tools/log"
	"time"
)

//CancelPnrAsync 异步取消PNR接口
func CancelPnrAsync(param *CancelPnrParam) *CancelPnrResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *CancelPnrResult, 1)

	//异步执行
	go func(p *CancelPnrParam, c chan<- *CancelPnrResult) {
		c <- CancelPnr(p)
	}(param, ch)

	var result *CancelPnrResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		Logger.Error("CancelPnrAsync", param.LogContext+" cancel pnr timed out!!!")
		break
	}

	elapsed := time.Since(start)
	Logger.Info("CancelPnrAsync", fmt.Sprintf("%s spent %f seconds.", param.LogContext, elapsed.Seconds()))

	return result
}

//CancelPnr 取消PNR接口
func CancelPnr(param *CancelPnrParam) *CancelPnrResult {
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

	paramJSON, _ := CcJson.Serialize(param)
	if paramJSON != "" {
		DumpFile(param.LogContext+"_Param.json", paramJSON, true)
	}

	//转换查询参数
	reqEnvelope, err := getReqEnvolpe(param, pcc.BranchCode)
	if err != nil {
		Logger.Error("CancelPnr", param.LogContext+err.Error())
		return SetErrorCode(ErrInvalidParameter)
	}

	//序列化请求XML文本
	reqXML, err := xml.Marshal(reqEnvelope)
	if err != nil {
		Logger.Error("CancelPnr", param.LogContext+err.Error())
		return SetErrorCode(ErrParseParameterError)
	}

	if cfg.OutputToFile {
		DumpFile(param.LogContext+"_RQ.xml", string(reqXML), true)
	}

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.UniversalServiceName, reqXML)
	if err != nil {
		Logger.Error("CancelPnr", param.LogContext+err.Error())
		return SetErrorCode(ErrProcessError)
	}

	DumpFile(param.LogContext+"_RS.xml", string(rspXML), true)

	//反序列化响应XML文件
	var rspEnvelope uniproxy.CancelPnrRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	//转换查询结果
	result, err := getResult(rspEnvelope.Body)
	if err != nil {
		Logger.Error("CancelPnr", param.LogContext+err.Error())
		return SetErrorCode(ErrCancelPnrError)
	}

	resultJSON, _ := CcJson.Serialize(result)
	if resultJSON != "" {
		DumpFile(param.LogContext+"_Result.json", resultJSON, true)
	}

	return result
}
