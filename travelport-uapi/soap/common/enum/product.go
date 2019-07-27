package enum

//Available product types
type EnumProduct int

const (
	ProductAir EnumProduct = iota
	ProductVehicle
	ProductHotel
	ProductRail
	ProductCruise
	ProductOther
)

//获取EnumProduct的字符串值
func (this EnumProduct) String() string {
	switch this {
	case ProductAir:
		return "Air"
	case ProductVehicle:
		return "Vehicle"
	case ProductHotel:
		return "Hotel"
	case ProductRail:
		return "Rail"
	case ProductCruise:
		return "Cruise"
	case ProductOther:
		return "Other"
	}
	return ""
}

//获取EnumProduct的整数值
func (this EnumProduct) Value() int {
	if this >= ProductAir && this <= ProductOther {
		return int(this)
	} else {
		return -1
	}
}
