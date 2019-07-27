package enum

//IdentificationType to define how the customer will identify himself when collecting the ticket
type EnumFulfillmentIDType int

const (
	FidtBahnCard EnumFulfillmentIDType = iota
	FidtCreditCard
	FidtEuroChequeCard
	FidtCollectionReference
)

//获取EnumFulfillmentIDType的字符串值
func (this EnumFulfillmentIDType) String() string {
	switch this {
	case FidtBahnCard:
		return "Bahn Card"
	case FidtCreditCard:
		return "Credit Card"
	case FidtEuroChequeCard:
		return "Euro Cheque Card"
	case FidtCollectionReference:
		return "Collection Reference"
	}
	return ""
}

//获取EnumFulfillmentIDType的整数值
func (this EnumFulfillmentIDType) Value() int {
	if this >= FidtBahnCard && this <= FidtCollectionReference {
		return int(this)
	} else {
		return -1
	}
}
