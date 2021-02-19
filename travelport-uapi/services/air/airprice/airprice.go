package airprice

import (
	"encoding/xml"
	"time"

	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/air/airprice"
	"github.com/redochen/demos/travelport-uapi/soap"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	CcJson "github.com/redochen/tools/json"
	"github.com/redochen/tools/log"
)

//AirPriceAsync 异步获取价格接口
func AirPriceAsync(param *AirPriceParam) *AirPriceResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *AirPriceResult, 1)

	//异步执行
	go func(p *AirPriceParam, c chan<- *AirPriceResult) {
		c <- AirPrice(p)
	}(param, ch)

	var result *AirPriceResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		log.Error(param.LogContext, " air price timed out!!!")
		break
	}

	elapsed := time.Since(start)
	log.Infof("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//AirPrice 获取价格接口
func AirPrice(param *AirPriceParam) *AirPriceResult {
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
		log.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrInvalidParameter)
	}

	//序列化请求XML文本
	reqXML, err := xml.Marshal(reqEnvelope)
	if err != nil {
		log.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrParseParameterError)
	}

	DumpFile(param.LogContext+"_RQ.xml", string(reqXML), true)

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.AirServiceName, reqXML)
	if err != nil {
		log.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrProcessError)
	}

	DumpFile(param.LogContext+"_RS.xml", string(rspXML), true)

	//反序列化响应XML文件
	var rspEnvelope airproxy.AirPriceRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	//解析错误信息
	if rspEnvelope.Body != nil && rspEnvelope.Body.Fault != nil {
		return SetErrorMessage(pasreFault(rspEnvelope.Body.Fault))
	}

	result, err := getResult(rspEnvelope.Body)
	if err != nil {
		log.Error(param.LogContext, err.Error())
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
