package retrievepnr

import (
	"errors"
	. "github.com/redochen/demos/travelport-uapi/models/universal/retrievepnr"
	. "github.com/redochen/demos/travelport-uapi/services/universal/createpnr"
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *RetrievePnrParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if len(param.UrlCode) == 0 {
		return nil, errors.New("UniversalRecord locator code cannot be nil")
	}

	body := uniproxy.NewRetrievePnrReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create RetrievePnrReqBody")
	}

	req := body.Request
	req.UniversalRecordLocatorCode = com.TypeLocatorCode(param.UrlCode)

	return soap.NewReqEnvelope(body), nil
}

//getResult 解析结果
func getResult(body *uniproxy.RetrievePnrRspBody) (*RetrievePnrResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("RetrievePnrRspBody is nil")
	}

	if nil == rsp.UniversalRecord {
		return nil, errors.New("universal record is empty")
	}

	result := &RetrievePnrResult{}

	//解析预订结果
	ParseUniversalRecord(&result.CreatePnrResult, rsp.UniversalRecord)

	return result, nil
}
