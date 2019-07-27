package cancelpnr

import (
	"errors"
	"fmt"
	. "github.com/redochen/demos/travelport-uapi/models/universal/cancelpnr"
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	uniproxy "github.com/redochen/demos/travelport-uapi/soap/universal/proxy"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *CancelPnrParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if len(param.UrlCode) == 0 {
		return nil, errors.New("UniversalRecord locator code cannot be nil")
	}

	if len(param.Version) == 0 {
		return nil, errors.New("version cannot be nil")
	}

	body := uniproxy.NewCancelPnrReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create CancelPnrReqBody")
	}

	req := body.Request
	req.UniversalRecordLocatorCode = com.TypeLocatorCode(param.UrlCode)
	req.Version = com.TypeURVersion(param.Version)

	return soap.NewReqEnvelope(body), nil
}

//getResult 解析结果
func getResult(body *uniproxy.CancelPnrRspBody) (*CancelPnrResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("UniversalRecordCancelRsp is nil")
	}

	result := &CancelPnrResult{}

	if rsp.ProviderReservationStatus != nil && len(rsp.ProviderReservationStatus) > 0 {
		result.ResStatus = make([]*ProviderReservationStatus, 0)
		for _, prs := range rsp.ProviderReservationStatus {
			p := &ProviderReservationStatus{
				LocatorCode:  string(prs.LocatorCode),
				Cancelled:    prs.Cancelled,
				ProviderCode: string(prs.ProviderCode),
			}

			if prs.CancelInfo != nil {
				p.Message = fmt.Sprintf("[%s]%s(%d)",
					prs.CancelInfo.Type,
					prs.CancelInfo.Value,
					prs.CancelInfo.Code)
			}

			result.ResStatus = append(result.ResStatus, p)
		}
	}

	return result, nil
}
