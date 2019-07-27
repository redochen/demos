package enum

//Deposit, Guarantee, or Prepayment required to hold/book the room. Applicable only for HotelSupershopper, Hotel Details and Hotel rules.
type EnumGuaranteeType int

const (
	GtDeposit EnumGuaranteeType = iota
	GtGuarantee
	GtPrepayment
)

//获取EnumGuaranteeType的字符串值
func (this EnumGuaranteeType) String() string {
	switch this {
	case GtDeposit:
		return "Deposit"
	case GtGuarantee:
		return "Guarantee"
	case GtPrepayment:
		return "Prepayment"
	}
	return ""
}

//获取EnumGuaranteeType的整数值
func (this EnumGuaranteeType) Value() int {
	if this >= GtDeposit && this <= GtPrepayment {
		return int(this)
	} else {
		return -1
	}
}
