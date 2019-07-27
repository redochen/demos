package enum

//Identifies the type of action (if any) to take on this air reservation. Only TTL, TAU, TAX and TAW can be set by the user.
type EnumActionType int

const (
	TAW EnumActionType = iota
	TTL
	TLCXL
	ACTIVE
	CXL
	TAU //Equivalent to TAX in Worldspan
	TRH
)

//获取EnumActionType的字符串值
func (this EnumActionType) String() string {
	switch this {
	case TAW:
		return "TAW"
	case TTL:
		return "TTL"
	case TLCXL:
		return "TLCXL"
	case ACTIVE:
		return "ACTIVE"
	case CXL:
		return "CXL"
	case TAU:
		return "TAU"
	case TRH:
		return "TRH"

	}
	return ""
}

//获取EnumActionType的整数值
func (this EnumActionType) Value() int {
	if this >= TAW && this <= TRH {
		return int(this)
	} else {
		return -1
	}
}
