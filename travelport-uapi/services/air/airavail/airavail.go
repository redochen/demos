package airavail

import (
	"encoding/xml"
	"fmt"
	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/air/airavail"
	"github.com/redochen/demos/travelport-uapi/soap"
	airproxy "github.com/redochen/demos/travelport-uapi/soap/air/proxy"
	comrq "github.com/redochen/demos/travelport-uapi/soap/common/request"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	. "github.com/redochen/tools/json"
	. "github.com/redochen/tools/log"
	"time"
)

//AirAvailAsync 异步AV查询接口
func AirAvailAsync(param *AirAvailParam) *AirAvailResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *AirAvailResult, 1)

	//异步执行
	go func(p *AirAvailParam, c chan<- *AirAvailResult) {
		c <- AirAvail(p)
	}(param, ch)

	var result *AirAvailResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		Logger.Error(param.LogContext, " air avail search timed out!!!")
		break
	}

	elapsed := time.Since(start)
	Logger.Debugf("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//AirAvail AV查询接口
func AirAvail(param *AirAvailParam) *AirAvailResult {
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

	var result *AirAvailResult
	var nextRspRefs []*comrs.NextResultReference
	pageIndex := 0

	for firstTime := true; firstTime || (nextRspRefs != nil && len(nextRspRefs) > 0); {
		var av *AirAvailResult
		av, nextRspRefs = getAvail(param, nextRspRefs, pageIndex, pcc)

		if firstTime {
			result = av //第一次查询
		} else if av != nil && av.Flights != nil {
			if nil == result {
				result = &AirAvailResult{}
			}

			if nil == result.Flights {
				result.Flights = make([]*Flight, 0)
			}

			for _, flight := range av.Flights {
				if nil == flight {
					continue
				}

				result.Flights = append(result.Flights, flight)
			}
		}

		//未请求翻页，或者已经翻够页
		if param.MaxPageDown <= 0 || pageIndex >= param.MaxPageDown {
			break
		}

		//没有后续页
		if nil == nextRspRefs || len(nextRspRefs) <= 0 {
			break
		}

		firstTime = false
		pageIndex++
	}

	if cfg.OutputToFile {
		resultJSON, _ := CcJson.Serialize(result)
		if resultJSON != "" {
			DumpFile(param.LogContext+"_Result.json", resultJSON, false)
		}
	}

	return result
}

//getAvail 获取AV
func getAvail(param *AirAvailParam, nextRspResultReferences []*comrs.NextResultReference, pageIndex int, pcc *cfg.PCC) (*AirAvailResult, []*comrs.NextResultReference) {
	nextReqResultReferences := translateNextResultReferences(nextRspResultReferences, param.ProviderCode)

	reqEnvelope, err := getReqEnvolpe(param, nextReqResultReferences, pcc.BranchCode)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrInvalidParameter), nil
	}

	//序列化请求XML文本
	reqXML, err := xml.Marshal(reqEnvelope)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrParseParameterError), nil
	}

	var rqDumpFile string
	var rsDumpFile string

	if pageIndex > 0 {
		rqDumpFile = fmt.Sprintf("%s_P%d_RQ.xml", param.LogContext, pageIndex)
		rsDumpFile = fmt.Sprintf("%s_P%d_RS.xml", param.LogContext, pageIndex)
	} else {
		rqDumpFile = fmt.Sprintf("%s_RQ.xml", param.LogContext)
		rsDumpFile = fmt.Sprintf("%s_RS.xml", param.LogContext)
	}

	DumpFile(rqDumpFile, string(reqXML), false)

	//调用Galileo接口
	rspXML, err := PostRequest(pcc, soap.AirServiceName, reqXML)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrProcessError), nil
	}

	DumpFile(rsDumpFile, string(rspXML), false)

	//反序列化响应XML文件
	var rspEnvelope airproxy.AirAvailRspEnvelope
	err = xml.Unmarshal([]byte(rspXML), &rspEnvelope)
	if err != nil {
		return SetErrorCode(ErrParseResultError), nil
	}

	//转换查询结果
	result, nextRefs, err := getResult(rspEnvelope.Body)
	if err != nil {
		Logger.Error(param.LogContext, err.Error())
		return SetErrorCode(ErrNoAvailDataReturned), nextRefs
	}

	return result, nextRefs
}

//translateNextResultReferences 转换参数
func translateNextResultReferences(nextRspResultReferences []*comrs.NextResultReference, providerCode string) []*comrq.NextResultReference {
	if nil == nextRspResultReferences || len(nextRspResultReferences) <= 0 {
		return nil
	}

	nextReqResultReferences := make([]*comrq.NextResultReference, 0)

	for _, nextRspResultReference := range nextRspResultReferences {
		if nil == nextRspResultReference {
			continue
		}

		nextReqResultReference := &comrq.NextResultReference{
			Value:        nextRspResultReference.Value,
			ProviderCode: nextRspResultReference.ProviderCode,
		}

		if len(nextReqResultReference.ProviderCode) <= 0 {
			nextReqResultReference.ProviderCode = providerCode
		}

		nextReqResultReferences = append(nextReqResultReferences, nextReqResultReference)
	}

	return nextReqResultReferences
}
