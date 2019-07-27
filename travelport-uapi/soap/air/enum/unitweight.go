package enum

//Defines the ability to eticket an entity (Yes, No, Required, Ticketless)
type EnumUnitWeight int

const (
	Kilograms EnumUnitWeight = iota
	Pounds
)

//获取EnumUnitWeight的字符串值
func (w EnumUnitWeight) String() string {
	switch w {
	case Kilograms:
		return "Kilograms"
	case Pounds:
		return "Pounds"
	}
	return ""
}

//获取EnumUnitWeight的整数值
func (w EnumUnitWeight) Value() int {
	if w >= Kilograms && w <= Pounds {
		return int(w)
	} else {
		return -1
	}
}
