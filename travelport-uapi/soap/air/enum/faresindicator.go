package enum

//Defines the type of fares to return (Only public fares, Only private fares, Only agency private fares, Only airline private fares or all fares)
type EnumFaresIndicator int

const (
	FiPublicFaresOnly EnumFaresIndicator = iota
	FiPrivateFaresOnly
	FiAgencyPrivateFaresOnly
	FiAirlinePrivateFaresOnly
	FiPublicAndPrivateFares
	FiNetFaresOnly
	FiAllFares //Applicable for 1G/1V air shop only
)

//获取EnumFaresIndicator的字符串值
func (this EnumFaresIndicator) String() string {
	switch this {
	case FiPublicFaresOnly:
		return "PublicFaresOnly"
	case FiPrivateFaresOnly:
		return "PrivateFaresOnly"
	case FiAgencyPrivateFaresOnly:
		return "AgencyPrivateFaresOnly"
	case FiAirlinePrivateFaresOnly:
		return "AirlinePrivateFaresOnly"
	case FiPublicAndPrivateFares:
		return "PublicAndPrivateFares"
	case FiNetFaresOnly:
		return "NetFaresOnly"
	case FiAllFares:
		return "AllFares"
	}
	return ""
}

//获取EnumFaresIndicator的整数值
func (this EnumFaresIndicator) Value() int {
	if this >= FiPublicFaresOnly && this <= FiAllFares {
		return int(this)
	} else {
		return -1
	}
}
