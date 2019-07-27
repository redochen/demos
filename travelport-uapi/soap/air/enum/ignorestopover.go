package enum

type EnumIgnoreStopOver int

const (
	IsoNoStopOver EnumIgnoreStopOver = iota
	IsoStopOver
	IsoIgnoreSegment
)

//获取EnumIgnoreStopOver的字符串值
func (this EnumIgnoreStopOver) String() string {
	switch this {
	case IsoNoStopOver:
		return "NoStopOver"
	case IsoStopOver:
		return "StopOver"
	case IsoIgnoreSegment:
		return "IgnoreSegment"
	}
	return ""
}

//获取EnumIgnoreStopOver的整数值
func (this EnumIgnoreStopOver) Value() int {
	if this >= IsoNoStopOver && this <= IsoIgnoreSegment {
		return int(this)
	} else {
		return -1
	}
}
