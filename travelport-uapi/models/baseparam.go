package models

import (
	"fmt"
	"github.com/redochen/demos/travelport-uapi/config"
	. "github.com/redochen/tools/random"
	"strings"
)

//BaseParam 接口参数基类
type BaseParam struct {
	GdsAccount     string `xml:"gdsAccount,omitempty" json:"gdsAccount,omitempty"`         //接口PCC号
	ProviderCode   string `xml:"providerCode,omitempty" json:"providerCode,omitempty"`     //接口代码：1G/1P/1V等
	ServiceName    string `xml:"-" json:"-"`                                               //服务名称
	TimeoutSeconds int    `xml:"timeoutSeconds,omitempty" json:"timeoutSeconds,omitempty"` //超时秒数
	LogContext     string `xml:"logContext,omitempty" json:"logContext,omitempty"`         //日志上下文
	RandomText     string `xml:"-" json:"-"`                                               //查询特征码
}

//PreCheck 预检查
func (baseParam *BaseParam) PreCheck() {
	if baseParam.ProviderCode == "" {
		baseParam.ProviderCode = "1G"
	}

	if baseParam.TimeoutSeconds <= 0 {
		baseParam.TimeoutSeconds = config.TimeoutSeconds
	}

	if baseParam.RandomText == "" {
		baseParam.RandomText = GetRandomText(8, true, true)
	}

	if baseParam.LogContext == "" || !strings.Contains(baseParam.LogContext, baseParam.RandomText) {
		baseParam.LogContext = fmt.Sprintf("[%s][%s]%s",
			baseParam.RandomText,
			baseParam.ServiceName,
			baseParam.LogContext)
	}
}
