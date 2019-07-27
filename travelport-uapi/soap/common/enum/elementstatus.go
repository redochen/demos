package enum

//Values to specify the state of the element. "A" refers to "Add" , "M" refers to "Modified" and "C" refers to error conditions when value provided in "Key" attribute is not retained in response
type EnumElementStatus int

const (
	EsAdd EnumElementStatus = iota
	EsModified
	EsErrorConditions
)

//获取EnumElementStatus的字符串值
func (this EnumElementStatus) String() string {
	switch this {
	case EsAdd:
		return "A"
	case EsModified:
		return "M"
	case EsErrorConditions:
		return "C"
	}
	return ""
}

//获取EnumElementStatus的整数值
func (this EnumElementStatus) Value() int {
	if this >= EsAdd && this <= EsErrorConditions {
		return int(this)
	} else {
		return -1
	}
}
