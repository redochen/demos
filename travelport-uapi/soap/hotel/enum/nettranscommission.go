package enum

//Type to support net trans commission indicator. Supported for 1G and 1V provider.
type EnumNetTransCommission int

const (
	NtcA EnumNetTransCommission = iota //A stands for Excellent.
	NtcB                               //B stands for Good.
	NtcC                               //C stands for Poor.
	NtcP                               //P stands for Payment Bureau
	NtcX                               //X stands for Unknown. To support any other value than A,B,C and P.
)

//获取EnumNetTransCommission的字符串值
func (this EnumNetTransCommission) String() string {
	switch this {
	case NtcA:
		return "A"
	case NtcB:
		return "B"
	case NtcC:
		return "C"
	case NtcP:
		return "P"
	case NtcX:
		return "X"
	}
	return ""
}

//获取EnumNetTransCommission的整数值
func (this EnumNetTransCommission) Value() int {
	if this >= NtcA && this <= NtcX {
		return int(this)
	} else {
		return -1
	}
}
