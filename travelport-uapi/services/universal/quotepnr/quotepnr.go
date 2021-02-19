package quotepnr

import (
	"time"

	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/universal/quotepnr"
	. "github.com/redochen/demos/travelport-uapi/services/air/airprice"
	. "github.com/redochen/demos/travelport-uapi/services/universal/retrievepnr"
	. "github.com/redochen/demos/travelport-uapi/util"
	CcFunc "github.com/redochen/tools/function"
	CcJson "github.com/redochen/tools/json"
	"github.com/redochen/tools/log"
)

//QuotePnrAsync 异步询价PNR接口
func QuotePnrAsync(param *QuotePnrParam) *QuotePnrResult {
	defer CcFunc.CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *QuotePnrResult, 1)

	//异步执行
	go func(p *QuotePnrParam, c chan<- *QuotePnrResult) {
		c <- QuotePnr(p)
	}(param, ch)

	var result *QuotePnrResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		log.Error(param.LogContext, " quote pnr timed out!!!")
		break
	}

	elapsed := time.Since(start)
	log.Infof("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//QuotePnr 询价PNR接口
func QuotePnr(param *QuotePnrParam) *QuotePnrResult {
	defer CcFunc.CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	//获取RetrievePnr参数
	retrievePnrParam := getRetrievePnrParam(param)
	if nil == retrievePnrParam {
		log.Error(param.LogContext, "translateQuotePnrParamToRetrievePnrParam failed")
		return SetErrorCode(ErrParseParameterError)
	}

	//提取PNR信息
	retrievePnrResult := RetrievePnr(retrievePnrParam)
	if nil == retrievePnrResult {
		log.Error(param.LogContext, "RetrievePnr failed")
		return SetErrorCode(ErrRetrievePnrError)
	}

	result := &QuotePnrResult{}
	result.FillinByRetrievePnrResult(retrievePnrResult)

	if result.Status == 0 {
		airPriceParam := getAirPriceParam(retrievePnrResult, param)
		if nil == airPriceParam {
			result.SetErrorCode(ErrPnrAlreadyCancelled)
		} else {
			//获取行程价格
			airPriceResult := AirPrice(airPriceParam)
			if nil == airPriceResult {
				log.Error(param.LogContext, "AirPrice failed")
				result.SetErrorCode(ErrQuotePriceError)
			} else {
				result.FillinByAirPriceResult(airPriceResult)
			}
		}
	}

	if cfg.OutputToFile {
		resultJSON, _ := CcJson.Serialize(result)
		if resultJSON != "" {
			DumpFile(param.LogContext+"_Result.json", resultJSON, false)
		}
	}

	return result
}
