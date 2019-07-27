package hotelmedia

import (
	"errors"
	hotavmodel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	. "github.com/redochen/demos/travelport-uapi/models/hotel/hotelmedia"
	hotavsvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelavail"
	"github.com/redochen/demos/travelport-uapi/soap"
	com "github.com/redochen/demos/travelport-uapi/soap/common"
	comrs "github.com/redochen/demos/travelport-uapi/soap/common/response"
	hot "github.com/redochen/demos/travelport-uapi/soap/hotel"
	hotproxy "github.com/redochen/demos/travelport-uapi/soap/hotel/proxy"
	hotrq "github.com/redochen/demos/travelport-uapi/soap/hotel/request"
	hotrs "github.com/redochen/demos/travelport-uapi/soap/hotel/response"
)

//getReqEnvolpe 获取请求参数
func getReqEnvolpe(param *HotelMediaParam, branchCode string) (*soap.ReqEnvelope, error) {
	if nil == param {
		return nil, errors.New("param cannot be nil")
	}

	if nil == param.HotelProperties || len(param.HotelProperties) == 0 {
		return nil, errors.New("should contain at least one hotel property")
	}

	body := hotproxy.NewHotelMediaReqBody(branchCode)
	if nil == body {
		return nil, errors.New("failed to create HotelMediaReqBody")
	}

	req := body.Request

	if param.HotelProperties != nil && len(param.HotelProperties) > 0 {
		req.HotelProperty = make([]*hotrq.HotelProperty, 0)
		for _, hp := range param.HotelProperties {
			property := translateHotelPropertyParam(hp)
			if property != nil {
				req.HotelProperty = append(req.HotelProperty, property)
			}
		}
	}

	return soap.NewReqEnvelope(body), nil
}

//translateHotelPropertyParam 转换HotelProperty参数
func translateHotelPropertyParam(property *HotelPropertyParam) *hotrq.HotelProperty {
	if nil == property {
		return nil
	}

	result := &hotrq.HotelProperty{
		HotelChain: com.TypeHotelChainCode(property.Chain),
		HotelCode:  com.TypeHotelCode(property.Code),
	}

	//设置名称
	if property.Name != "" {
		result.Name = property.Name
	}

	//设置地点
	if property.Location != "" {
		result.HotelLocation = hot.TypeHotelLocationCode(property.Location)
	}

	//设置交通
	if property.Transportation > 0 {
		result.HotelTransportation = com.TypeOTACode(property.Transportation)
	}

	//设置地址
	if property.Address != nil && len(property.Address) > 0 {
		result.PropertyAddress = &hotrq.UnstructuredAddress{
			Address: make([]string, 0),
		}

		for _, addr := range property.Address {
			result.PropertyAddress.Address = append(result.PropertyAddress.Address, addr)
		}
	}

	//设置其他
	if property.ParticipationLevel != "" {
		result.ParticipationLevel = com.StringLength1(property.ParticipationLevel)
	}

	if property.ReserveRequirement != "" {
		result.ReserveRequirement = com.TypeReserveRequirement(property.ReserveRequirement)
	}

	if property.VendorLocationKey != "" {
		result.VendorLocationKey = property.VendorLocationKey
	}

	return result
}

//getResult 转换结果
func getResult(body *hotproxy.HotelMediaRspBody) (*HotelMediaResult, error) {
	if nil == body || nil == body.Response {
		return nil, errors.New("response is empty")
	}

	rsp := body.Response
	if nil == rsp {
		return nil, errors.New("HotelMediaLinksRsp is empty")
	}

	result := &HotelMediaResult{}

	if rsp.HotelPropertyWithMediaItems != nil && len(rsp.HotelPropertyWithMediaItems) > 0 {
		result.HotelProperties = make([]*HotelPropertyExResult, 0)

		for _, hpwmi := range rsp.HotelPropertyWithMediaItems {
			property := translateHotelPropertyWithMediaItemsToHotelProperty(hpwmi)
			if property != nil {
				result.HotelProperties = append(result.HotelProperties, property)
			}
		}
	}

	return result, nil
}

//translateHotelPropertyWithMediaItemsToHotelProperty 转换结果模式
func translateHotelPropertyWithMediaItemsToHotelProperty(hotel *hotrs.HotelPropertyWithMediaItems) *HotelPropertyExResult {
	if nil == hotel || nil == hotel.HotelProperty {
		return nil
	}

	result := &HotelPropertyExResult{
		Property: hotavsvc.TranslateHotelPropertyResult(hotel.HotelProperty),
	}

	//解析媒介信息
	if result != nil && hotel.MediaItem != nil && len(hotel.MediaItem) > 0 {
		result.Medias = make([]*hotavmodel.MediaInfoResult, 0)
		for _, mi := range hotel.MediaItem {
			media := translateMediaItem(mi)
			if media != nil {
				result.Medias = append(result.Medias, media)
			}
		}
	}

	return result
}

//translateMediaItem ...
func translateMediaItem(item *comrs.MediaItem) *hotavmodel.MediaInfoResult {
	if nil == item {
		return nil
	}

	media := &hotavmodel.MediaInfoResult{
		Caption:  item.Caption,
		Height:   item.Height,
		Width:    item.Width,
		Type:     item.Type,
		Url:      item.Url,
		Icon:     item.Icon,
		SizeCode: item.SizeCode,
	}

	return media
}
