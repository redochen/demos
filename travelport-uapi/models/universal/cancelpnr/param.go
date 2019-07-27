package cancelpnr

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/models"
)

//CancelPnrParam 取消PNR参数
type CancelPnrParam struct {
	models.BaseParam
	UrlCode string `xml:"urlCode" json:"urlCode"` //UR编号
	Version string `xml:"version" json:"version"` //UR版本号
}

//PreCheck 预检查
func (cancelPnrParam *CancelPnrParam) PreCheck() {
	cancelPnrParam.ServiceName = "CancelPnr"

	//if cancelPnrParam.LogContext == "" {
	cancelPnrParam.LogContext = fmt.Sprintf("[%s][%s]", cancelPnrParam.UrlCode, cancelPnrParam.Version)
	//}

	cancelPnrParam.BaseParam.PreCheck()
}
