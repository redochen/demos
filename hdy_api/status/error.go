package status

import (
	"fmt"

	CcStr "github.com/redochen/tools/string"
)

//ErrorCode 错误代码
type ErrorCode int

const (
	Success ErrorCode = iota
	CustomError
	InvalidParameters
	Timeout
)

//GetErrMessage 获取错误信息
func GetErrMessage(errCode ErrorCode, errMsg ...string) string {
	if CustomError == errCode {
		msg := CcStr.FirstValid(errMsg...)
		if msg != "" {
			return msg
		}
	}

	switch errCode {
	case Success:
		return "success"
	case InvalidParameters:
		return "invalid paramters"
	case Timeout:
		return "request timed out"
	default:
		return fmt.Sprintf("unkown error code %d", errCode)
	}
}
