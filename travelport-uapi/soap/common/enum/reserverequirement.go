package enum

//Type of payment required to reserve travel i.e. Hotel Reservation requirement
type EnumReserveRequiremen int

const (
	RrDeposit EnumReserveRequiremen = iota
	RrGuarantee
	RrPrepayment
	RrOther
)

//获取EnumReserveRequiremen的字符串值
func (this EnumReserveRequiremen) String() string {
	switch this {
	case RrDeposit:
		return "Deposit"
	case RrGuarantee:
		return "Guarantee"
	case RrPrepayment:
		return "Prepayment"
	case RrOther:
		return "Other"
	}
	return ""
}

//获取EnumReserveRequiremen的整数值
func (this EnumReserveRequiremen) Value() int {
	if this >= RrDeposit && this <= RrOther {
		return int(this)
	} else {
		return -1
	}
}
