package enum

//Used in Low Fare Search to better promote unique results
type EnumDiversity int

const (
	DtBlend EnumDiversity = iota
	DtAirports
	DtCarrier
	DtOrigin
	DtDestination
	DtDateCombination
	DtFirstODDate
	DtSecondODDate
	DtFirstOD
	DtSecondOD
)

//获取EnumDiversity的字符串值
func (this EnumDiversity) String() string {
	switch this {
	case DtBlend:
		return "Blend"
	case DtAirports:
		return "Airports"
	case DtCarrier:
		return "Carrier"
	case DtOrigin:
		return "Origin"
	case DtDestination:
		return "Destination"
	case DtDateCombination:
		return "DateCombination"
	case DtFirstODDate:
		return "FirstODDate"
	case DtSecondODDate:
		return "SecondODDate"
	case DtFirstOD:
		return "FirstOD"
	case DtSecondOD:
		return "SecondOD"
	}
	return ""
}

//获取EnumDiversity的整数值
func (this EnumDiversity) Value() int {
	if this >= DtBlend && this <= DtSecondOD {
		return int(this)
	} else {
		return -1
	}
}
