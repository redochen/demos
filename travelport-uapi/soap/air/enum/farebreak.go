package enum

//Types of fare break.
type EnumFareBreak int

const (
	FbMustBreak     EnumFareBreak = iota //Break Fare at the associated segment. Multiple Breaks or No Breaks may be allowed.
	FbMustOnlyBreak                      //Only Break Fare at the associated segment. Fare Break in the entire itinerary is allowed only at the concerned segment.
	FbMustNotBreak                       //No Fare Break allowed at the associated segment.
)

//获取EnumFareBreak的字符串值
func (this EnumFareBreak) String() string {
	switch this {
	case FbMustBreak:
		return "MustBreak"
	case FbMustOnlyBreak:
		return "MustOnlyBreak"
	case FbMustNotBreak:
		return "MustNotBreak"
	}
	return ""
}

//获取EnumFareBreak的整数值
func (this EnumFareBreak) Value() int {
	if this >= FbMustBreak && this <= FbMustNotBreak {
		return int(this)
	} else {
		return -1
	}
}
