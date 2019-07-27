package enum

//The passenger residency type.Residence Type can be Employee, National or Resident
type EnumResidency int

const (
	ResEmployee EnumResidency = iota
	ResNational
	ResResident
)

//获取EnumResidency的字符串值
func (this EnumResidency) String() string {
	switch this {
	case ResEmployee:
		return "Employee"
	case ResNational:
		return "National"
	case ResResident:
		return "Resident"
	}
	return ""
}

//获取EnumResidency的整数值
func (this EnumResidency) Value() int {
	if this >= ResEmployee && this <= ResResident {
		return int(this)
	} else {
		return -1
	}
}
