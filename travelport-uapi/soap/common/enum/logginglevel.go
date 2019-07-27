package enum

type EnumLoggingLevel int

const (
	Trace EnumLoggingLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

//获取EnumLoggingLevel的字符串值
func (this EnumLoggingLevel) String() string {
	switch this {
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return ""
}

//获取EnumLoggingLevel的整数值
func (this EnumLoggingLevel) Value() int {
	if this >= Trace && this <= Fatal {
		return int(this)
	} else {
		return -1
	}
}
