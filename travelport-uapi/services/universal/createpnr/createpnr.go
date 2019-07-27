package createpnr

import (
	"encoding/xml"
	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	"github.com/redochen/demos/travelport-uapi/soap"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/json"
	. "github.com/redochen/tools/log"
	"time"
)

//CreatePnrAsync 异步实时订座接口
func CreatePnrAsync(param *CreatePnrParam) *CreatePnrResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *CreatePnrResult, 1)

	//异步执行
	go func(p *CreatePnrParam, c chan<- *CreatePnrResult) {
		c <- CreatePnr(p)
	}(param, ch)

	var result *CreatePnrResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		Logger.Error(param.LogContext, " create pnr timed out!!!")
		break
	}

	elapsed := time.Since(start)
	Logger.Infof("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//CreatePnr 实时订座接口
func CreatePnr(param *CreatePnrParam) *CreatePnrResult {
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

	if param.Air != nil {
		return CreateAirPnr(pcc, param)
	} else if param.Hotel != nil {
		return CreateHotelPnr(pcc, param)
	}

	return SetErrorCode(ErrNotSupported)
}

//CreateAirPnr 机票实时订座
func CreateAirPnr(pcc *cfg.PCC, param *CreatePnrParam) *CreatePnrResult {
	//转换查询参数
	reqEnvelope, err := getAirReqEnvolpe(param, pcc.BranchCode)
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

	DumpFile(param.LogContext+"_RQ.xml", string(reqXML), true)

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.AirServiceName, reqXML)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrProcessError)
	}

	DumpFile(param.LogContext+"_RS.xml", string(rspXML), true)

	//反序列化响应XML文件
	var rspEnvelope uniproxy.AirCreatePnrRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	//转换查询结果
	if rspEnvelope.Body != nil && rspEnvelope.Body.Fault != nil {
		return SetErrorMessage(pasreAirFault(rspEnvelope.Body.Fault))
	}

	result, err := getAirResult(rspEnvelope.Body)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrCreatePnrError)
	}

	resultJSON, _ := CcJson.Serialize(result)
	if resultJSON != "" {
		DumpFile(param.LogContext+"_Result.json", resultJSON, true)
	}

	return result
}

//CreateHotelPnr 酒店实时订座
func CreateHotelPnr(pcc *cfg.PCC, param *CreatePnrParam) *CreatePnrResult {
	//转换查询参数
	reqEnvelope, err := getHotelReqEnvolpe(param, pcc.BranchCode)
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

	DumpFile(param.LogContext+"_RQ.xml", string(reqXML), true)

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.HotelServiceName, reqXML)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrProcessError)
	}

	DumpFile(param.LogContext+"_RS.xml", string(rspXML), true)

	//反序列化响应XML文件
	var rspEnvelope uniproxy.HotelCreatePnrRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	//转换查询结果
	if rspEnvelope.Body != nil && rspEnvelope.Body.Fault != nil {
		return SetErrorMessage(parseHotelFault(rspEnvelope.Body.Fault))
	}

	result, err := getHotelResult(rspEnvelope.Body)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrCreatePnrError)
	}

	resultJSON, _ := CcJson.Serialize(result)
	if resultJSON != "" {
		DumpFile(param.LogContext+"_Result.json", resultJSON, true)
	}

	return result
}
