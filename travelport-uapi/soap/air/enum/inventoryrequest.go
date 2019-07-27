package enum

//The valid inventory types are Seamless - A, DirectAccess - B, Basic - C
type EnumInventoryRequest int

const (
	IrSeamless EnumInventoryRequest = iota
	IrDirectAccess
	IrBasic
)

//获取EnumInventoryRequest的字符串值
func (this EnumInventoryRequest) String() string {
	switch this {
	case IrSeamless:
		return "Seamless"
	case IrDirectAccess:
		return "DirectAccess"
	case IrBasic:
		return "Basic"
	}
	return ""
}

//获取EnumInventoryRequest的整数值
func (this EnumInventoryRequest) Value() int {
	if this >= IrSeamless && this <= IrBasic {
		return int(this)
	} else {
		return -1
	}
}
