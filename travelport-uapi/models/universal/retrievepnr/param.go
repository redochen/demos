package retrievepnr

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/models"
)

//RetrievePnrParam 提取PNR参数
type RetrievePnrParam struct {
	models.BaseParam
	UrlCode string `xml:"urlCode" json:"urlCode"` //UR编号
}

//PreCheck 预检查
func (retrievePnrParam *RetrievePnrParam) PreCheck() {
	retrievePnrParam.ServiceName = "RetrievePnr"

	//if retrievePnrParam.LogContext == "" {
	retrievePnrParam.LogContext = fmt.Sprintf("[%s]", retrievePnrParam.UrlCode)
	//}

	retrievePnrParam.BaseParam.PreCheck()
}
