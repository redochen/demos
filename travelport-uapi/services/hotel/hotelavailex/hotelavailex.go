package hotelavailex

import (
	"time"

	cfg "github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavailex"
	hotavsvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelavail"
	hotmisvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelmedia"
	. "github.com/redochen/demos/travelport-uapi/util"
	. "github.com/redochen/tools/function"
	CcJson "github.com/redochen/tools/json"
	"github.com/redochen/tools/log"
)

//HotelAvailExAsync 异步酒店查询接口
func HotelAvailExAsync(param *HotelAvailExParam) *HotelAvailExResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	ch := make(chan *HotelAvailExResult, 1)

	//异步执行
	go func(p *HotelAvailExParam, c chan<- *HotelAvailExResult) {
		c <- HotelAvailEx(p)
	}(param, ch)

	var result *HotelAvailExResult
	timeoutSeconds := time.Duration(param.TimeoutSeconds)
	start := time.Now()

	//带超时等待
	select {
	case result = <-ch:
		break
	case <-time.After(timeoutSeconds * time.Second):
		result = SetErrorCode(ErrTimeout)
		log.Error(param.LogContext, " hotel availex search timed out!!!")
		break
	}

	elapsed := time.Since(start)
	log.Infof("%s spent %f seconds.", param.LogContext, elapsed.Seconds())

	return result
}

//HotelAvailEx 酒店查询接口
func HotelAvailEx(param *HotelAvailExParam) *HotelAvailExResult {
	defer CheckPanic()

	if nil == param {
		return SetErrorCode(ErrInvalidParameter)
	}

	param.PreCheck()

	//调用HotelAvail接口
	hotelAvailResult := hotavsvc.HotelAvail(&param.HotelAvailParam)
	if nil == hotelAvailResult {
		log.Error(param.LogContext, "HotelAvail failed")
		return SetErrorCode(ErrHotelAvailError)
	}

	result := &HotelAvailExResult{}
	result.FillinByHotelAvailResult(hotelAvailResult)

	if result.Status == 0 {
		hotelMediaParam := getHotelMediaParamFromHotelAvailResult(hotelAvailResult, param)
		if nil == hotelMediaParam {
			log.Error(param.LogContext, "getHotelMediaParamFromHotelAvailResult failed")
			result.SetErrorCode(ErrNoAvailDataReturned)
		} else {
			//调用HotelMedia接口
			hotelMediaResult := hotmisvc.HotelMedia(hotelMediaParam)
			if nil == hotelMediaResult {
				log.Error(param.LogContext, "HotelMedia failed")
				result.SetErrorCode(ErrHotelMediaError)
			} else {
				result.FillinByHotelMediaResult(hotelMediaResult)
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
