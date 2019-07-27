package quotepnr

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/models"
)

//QuotePnrParam 询价PNR参数
type QuotePnrParam struct {
	models.BaseParam
	UrlCode  string `xml:"urlCode" json:"urlCode"`   //UR编号
	Carrier  string `xml:"carrier" json:"carrier"`   //开票航司
	Currency string `xml:"currency" json:"currency"` //货币代码
}

//PreCheck 预检查
func (quotePnrParam *QuotePnrParam) PreCheck() {
	quotePnrParam.ServiceName = "QuotePnr"

	//if quotePnrParam.LogContext == "" {
	quotePnrParam.LogContext = fmt.Sprintf("[%s][%s]", quotePnrParam.UrlCode, quotePnrParam.Carrier)
	if len(quotePnrParam.Currency) > 0 {
		quotePnrParam.LogContext += fmt.Sprintf("[%s]", quotePnrParam.Currency)
	}
	//}

	quotePnrParam.BaseParam.PreCheck()
}
