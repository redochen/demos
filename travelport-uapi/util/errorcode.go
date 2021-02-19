package util

//ErrorCode 错误代码
type ErrorCode int

const (
	ErrSuccess ErrorCode = iota //成功
	ErrOther                    //其他错误

	ErrUnauthenticatedIPAddress //未认证的IP地址
	ErrInvalidParameter         //无效的请求参数
	ErrInvalidPcc               //无效的PCC账号
	ErrNotSupported             //不支持该操作
	ErrParseParameterError      //解析参数错误
	ErrParseResultError         //解析结果错误
	ErrFormatResultError        //格式化参数错误
	ErrProcessError             //处理错误
	ErrTimeout                  //接口超时
	ErrNoDataReturned           //无数据返回
	ErrGetRequestEnvelope       //获取请求SOAP失败

	//机票相关错误
	ErrNoFareDataReturned  //未获取到有效的运价数据
	ErrNoAvailDataReturned //未获取到有效的AV数据

	//酒店相关错误
	ErrHotelAvailError //查询酒店失败
	ErrHotelMediaError //获取酒店媒介失败

	//PNR相关
	ErrCreatePnrError      //创建PNR失败
	ErrQuotePriceError     //获取价格失败
	ErrCancelPnrError      //取消PNR失败
	ErrRetrievePnrError    //提取PNR失败
	ErrPnrAlreadyCancelled //PNR已被取消

	ErrUnkown //未知错误
)

//String 获取错误代码的字符串值
func (self ErrorCode) String() string {
	switch self {
	case ErrSuccess:
		return "success"
	case ErrOther:
		return "other error"
	case ErrUnauthenticatedIPAddress:
		return "unauthenticated ip address"
	case ErrInvalidParameter:
		return "invalid parameter"
	case ErrInvalidPcc:
		return "invalid gds account"
	case ErrNotSupported:
		return "not supported"
	case ErrParseParameterError:
		return "failed to parse parameter"
	case ErrParseResultError:
		return "failed to parse result"
	case ErrFormatResultError:
		return "failed to format result"
	case ErrProcessError:
		return "failed to process"
	case ErrTimeout:
		return "Timed out"
	case ErrNoDataReturned:
		return "no data returned"
	case ErrGetRequestEnvelope:
		return "failed to get request envelope"
	case ErrNoFareDataReturned:
		return "no fare data returned"
	case ErrNoAvailDataReturned:
		return "no avail data returned"
	case ErrQuotePriceError:
		return "failed to quote price"
	case ErrCreatePnrError:
		return "failed to create pnr"
	case ErrCancelPnrError:
		return "failed to cancel pnr"
	case ErrRetrievePnrError:
		return "failed to retrieve pnr"
	case ErrPnrAlreadyCancelled:
		return "pnr has already been cancelled"
	case ErrHotelAvailError:
		return "failed to get availiable hotel"
	case ErrHotelMediaError:
		return "failed to get hotel media info"
	}
	return "unkown error"
}

//Value 获取错误代码的整数值
func (self ErrorCode) Value() int {
	if self >= ErrSuccess && self < ErrUnkown {
		return int(self)
	}

	return int(ErrUnkown)
}
