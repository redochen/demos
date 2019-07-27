package enum

//Values to specify the state of the element. "A" refers to "Add" , "M" refers to "Modified" and "C" refers to error conditions when value provided in "Key" attribute is not retained in response
type EnumFormOfRefund int

const (
	ForMCO EnumFormOfRefund = iota
	ForFormOfPayment
)

//获取EnumFormOfRefund的字符串值
func (this EnumFormOfRefund) String() string {
	switch this {
	case ForMCO:
		return "MCO"
	case ForFormOfPayment:
		return "FormOfPayment"
	}
	return ""
}

//获取EnumFormOfRefund的整数值
func (this EnumFormOfRefund) Value() int {
	if this >= ForMCO && this <= ForFormOfPayment {
		return int(this)
	} else {
		return -1
	}
}
