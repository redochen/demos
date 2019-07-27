package enum

type EnumOtherGuaranteeInfo int

const (
	OgiIATANumber    EnumOtherGuaranteeInfo = iota //IATA/ARC Number
	OgiAgencyAddress                               //Agency Address
	OgiDepositTaken                                //Deposit Taken
	OgiOthers                                      //Others
)

//获取EnumOtherGuaranteeInfo的字符串值
func (this EnumOtherGuaranteeInfo) String() string {
	switch this {
	case OgiIATANumber:
		return "IATA/ARC Number"
	case OgiAgencyAddress:
		return "Agency Address"
	case OgiDepositTaken:
		return "Deposit Taken"
	case OgiOthers:
		return "Others"
	}
	return ""
}

//获取EnumOtherGuaranteeInfo的整数值
func (this EnumOtherGuaranteeInfo) Value() int {
	if this >= OgiIATANumber && this <= OgiOthers {
		return int(this)
	} else {
		return -1
	}
}
