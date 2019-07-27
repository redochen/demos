package enum

//Available product types
type EnumPolicy int

const (
	PolicyAir EnumPolicy = iota
	PolicyVehicle
	PolicyHotel
	PolicyRail
	PolicyTicketing
)

//获取EnumPolicy的字符串值
func (this EnumPolicy) String() string {
	switch this {
	case PolicyAir:
		return "Air"
	case PolicyVehicle:
		return "Vehicle"
	case PolicyHotel:
		return "Hotel"
	case PolicyRail:
		return "Rail"
	case PolicyTicketing:
		return "Ticketing"
	}
	return ""
}

//获取EnumPolicy的整数值
func (this EnumPolicy) Value() int {
	if this >= PolicyAir && this <= PolicyTicketing {
		return int(this)
	} else {
		return -1
	}
}
