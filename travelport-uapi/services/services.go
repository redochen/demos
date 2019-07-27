package services

import (
	"github.com/gin-gonic/gin"
	AirAvailModel "github.com/redochen/demos/travelport-uapi/models/air/airavail"
	AirPriceModel "github.com/redochen/demos/travelport-uapi/models/air/airprice"
	LowFareModel "github.com/redochen/demos/travelport-uapi/models/air/lowfare"
	HotelAvailModel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavail"
	HotelAvailExModel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelavailex"
	HotelDetailsModel "github.com/redochen/demos/travelport-uapi/models/hotel/hoteldetails"
	HotelMediaModel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelmedia"
	HotelRulesModel "github.com/redochen/demos/travelport-uapi/models/hotel/hotelrules"
	CancelPnrModel "github.com/redochen/demos/travelport-uapi/models/universal/cancelpnr"
	CreatePnrModel "github.com/redochen/demos/travelport-uapi/models/universal/createpnr"
	QuotePnrModel "github.com/redochen/demos/travelport-uapi/models/universal/quotepnr"
	RetrievePnrModel "github.com/redochen/demos/travelport-uapi/models/universal/retrievepnr"
	AirAvailSvc "github.com/redochen/demos/travelport-uapi/services/air/airavail"
	AirPriceSvc "github.com/redochen/demos/travelport-uapi/services/air/airprice"
	LowFareSvc "github.com/redochen/demos/travelport-uapi/services/air/lowfare"
	HotelAvailSvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelavail"
	HotelAvailExSvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelavailex"
	HotelDetailsSvc "github.com/redochen/demos/travelport-uapi/services/hotel/hoteldetails"
	HotelMediaSvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelmedia"
	HotelRulesSvc "github.com/redochen/demos/travelport-uapi/services/hotel/hotelrules"
	PingSvc "github.com/redochen/demos/travelport-uapi/services/system/ping"
	CancelPnrSvc "github.com/redochen/demos/travelport-uapi/services/universal/cancelpnr"
	CreatePnrSvc "github.com/redochen/demos/travelport-uapi/services/universal/createpnr"
	QuotePnrSvc "github.com/redochen/demos/travelport-uapi/services/universal/quotepnr"
	RetrievePnrSvc "github.com/redochen/demos/travelport-uapi/services/universal/retrievepnr"
	"net/http"
)

//PING接口
func Ping(c *gin.Context) {
	rsp, err := PingSvc.Ping("Hello World!")
	if err != nil {
		c.String(http.StatusOK, err.Error())
	} else {
		c.String(http.StatusOK, rsp)
	}
}

//LowFare接口
func LowFare(c *gin.Context) {
	var param LowFareModel.LowFareParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, LowFareSvc.TestLowFareRQ())
	//c.JSON(http.StatusOK, LowFareSvc.TestLowFareRS())
	c.JSON(http.StatusOK, LowFareSvc.LowFareAsync(&param))
}

//AirAvail接口
func AirAvail(c *gin.Context) {
	var param AirAvailModel.AirAvailParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, AirAvailSvc.TestAirAvailRQ())
	//c.JSON(http.StatusOK, AirAvailSvc.TestAirAvailRS())
	c.JSON(http.StatusOK, AirAvailSvc.AirAvailAsync(&param))
}

//AirPrice接口
func AirPrice(c *gin.Context) {
	var param AirPriceModel.AirPriceParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, AirPriceSvc.TestAirPriceRQ())
	//c.JSON(http.StatusOK, AirPriceSvc.TestAirPriceRS())
	c.JSON(http.StatusOK, AirPriceSvc.AirPriceAsync(&param))
}

//HotelAvail接口
func HotelAvail(c *gin.Context) {
	var param HotelAvailModel.HotelAvailParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, HotelAvailSvc.TestHotelAvailRQ())
	//c.JSON(http.StatusOK, HotelAvailSvc.TestHotelAvailRS())
	c.JSON(http.StatusOK, HotelAvailSvc.HotelAvailAsync(&param))
}

//HotelAvailEx接口
func HotelAvailEx(c *gin.Context) {
	var param HotelAvailExModel.HotelAvailExParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, HotelAvailExSvc.TestHotelAvailExRQ())
	c.JSON(http.StatusOK, HotelAvailExSvc.HotelAvailExAsync(&param))
}

//HotelMedia接口
func HotelMedia(c *gin.Context) {
	var param HotelMediaModel.HotelMediaParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, HotelMediaSvc.TestHotelMediaRQ())
	//c.JSON(http.StatusOK, HotelMediaSvc.TestHotelMediaRS())
	c.JSON(http.StatusOK, HotelMediaSvc.HotelMediaAsync(&param))
}

//HotelDetails接口
func HotelDetails(c *gin.Context) {
	var param HotelDetailsModel.HotelDetailsParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, HotelDetailsSvc.TestHotelRateAndRuleSearchRQ())
	//c.JSON(http.StatusOK, HotelDetailsSvc.TestHotelDescriptionRQ())
	//c.JSON(http.StatusOK, HotelDetailsSvc.TestHotelDetailsRS())
	c.JSON(http.StatusOK, HotelDetailsSvc.HotelDetailsAsync(&param))
}

//HotelRules接口
func HotelRules(c *gin.Context) {
	var param HotelRulesModel.HotelRulesParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, HotelRulesSvc.TestHotelRulesRQ())
	//c.JSON(http.StatusOK, HotelRulesSvc.TestHotelRulesRS())
	c.JSON(http.StatusOK, HotelRulesSvc.HotelRulesAsync(&param))
}

//CreatePnr接口
func CreatePnr(c *gin.Context) {
	var param CreatePnrModel.CreatePnrParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, CreatePnrSvc.TestCreateAirPnrRQ())
	//c.JSON(http.StatusOK, CreatePnrSvc.TestCreateHotelPnrRQ())
	//c.JSON(http.StatusOK, CreatePnrSvc.TestCreateAirPnrRS())
	//c.JSON(http.StatusOK, CreatePnrSvc.TestCreateHotelPnrRS())
	c.JSON(http.StatusOK, CreatePnrSvc.CreatePnrAsync(&param))
}

//RetrievePnr接口
func RetrievePnr(c *gin.Context) {
	var param RetrievePnrModel.RetrievePnrParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, RetrievePnrSvc.TestRetrievePnrRQ())
	//c.JSON(http.StatusOK, RetrievePnrSvc.TestRetrievePnrRS())
	c.JSON(http.StatusOK, RetrievePnrSvc.RetrievePnrAsync(&param))
}

//QuotePnr接口
func QuotePnr(c *gin.Context) {
	var param QuotePnrModel.QuotePnrParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, QuotePnrSvc.TestQuotePnrRQ())
	c.JSON(http.StatusOK, QuotePnrSvc.QuotePnrAsync(&param))
}

//CancelPnr接口
func CancelPnr(c *gin.Context) {
	var param CancelPnrModel.CancelPnrParam
	c.Bind(&param)
	//c.JSON(http.StatusOK, CancelPnrSvc.TestCancelPnrRQ())
	//c.JSON(http.StatusOK, CancelPnrSvc.TestCancelPnrRS())
	c.JSON(http.StatusOK, CancelPnrSvc.CancelPnrAsync(&param))
}
