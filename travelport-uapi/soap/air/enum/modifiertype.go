package enum

//Type of Brand modifiers. e.g FareFamilyDisplay or BasicDetailOnly.
type EnumModifierType int

const (
	MtFareFamilyDisplay EnumModifierType = iota
	MtBasicDetailOnly
)

//获取EnumModifierType的字符串值
func (this EnumModifierType) String() string {
	switch this {
	case MtFareFamilyDisplay:
		return "FareFamilyDisplay"
	case MtBasicDetailOnly:
		return "BasicDetailOnly"
	}
	return ""
}

//获取EnumModifierType的整数值
func (this EnumModifierType) Value() int {
	if this >= MtFareFamilyDisplay && this <= MtBasicDetailOnly {
		return int(this)
	} else {
		return -1
	}
}
