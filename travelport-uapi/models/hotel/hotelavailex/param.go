package hotelavailex

import (
	hotav "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
)

//HotelAvailEx接口参数
type HotelAvailExParam struct {
	hotav.HotelAvailParam
}

//预检查
func (this *HotelAvailExParam) PreCheck() {
	this.HotelAvailParam.PreCheck()
	this.ServiceName = "HotelAvailEx"
}
