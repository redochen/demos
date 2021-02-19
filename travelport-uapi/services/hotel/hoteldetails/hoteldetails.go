package hoteldetails

import (
	"encoding/xml"
	"time"

	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hoteldetails"
	"github.com/redochen/demos/travelport-uapi/soap"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	CcJson "github.com/redochen/tools/json"
	"github.com/redochen/tools/log"
)

//HotelDetailsAsync 异步获取酒店详情接口
func HotelDetailsAsync(param *HotelDetailsParam) *HotelDetailsResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *HotelDetailsResult, 1)

	//异步执行
	go func(p *HotelDetailsParam, c chan<- *HotelDetailsResult) {
		c <- HotelDetails(p)
	}(param, ch)

	var result *HotelDetailsResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		log.Error(param.LogContext, " hotel details search timed out!!!")
		break
	}

	elapsed := time.Since(start)
	log.Infof("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//HotelDetails 获取酒店详情接口
func HotelDetails(param *HotelDetailsParam) *HotelDetailsResult {
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

	DumpFile(param.LogContext+"_RQ.xml", string(reqXML), false)

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.HotelServiceName, reqXML)
	if err != nil {
		log.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrProcessError)
	}

	DumpFile(param.LogContext+"_RS.xml", string(rspXML), false)

	//反序列化响应XML文件
	var rspEnvelope hotproxy.HotelDetailsRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError)
	}

	//转换查询结果
	result, err := getResult(rspEnvelope.Body)
	if err != nil {
		log.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrNoAvailDataReturned)
	}

	if cfg.OutputToFile {
		resultJSON, _ := CcJson.Serialize(result)
		if resultJSON != "" {
			DumpFile(param.LogContext+"_Result.json", resultJSON, false)
		}
	}

	return result
}
