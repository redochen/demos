package enum

//List the types of private fares, Agency private fare, Airline private Fare and Unknown.
//Also, this enumaration list includes PrivateFare to indetify private fares for GDSs where we can not identify specific private fares.
type EnumPrivateFare int

const (
	PfUnknownType EnumPrivateFare = iota
	PfPrivateFare
	PfAgencyPrivateFare
	PfAirlinePrivateFare
)

//获取EnumPrivateFare的字符串值
func (this EnumPrivateFare) String() string {
	switch this {
	case PfUnknownType:
		return "UnknownType"
	case PfPrivateFare:
		return "PrivateFare"
	case PfAgencyPrivateFare:
		return "AgencyPrivateFare"
	case PfAirlinePrivateFare:
		return "AirlinePrivateFare"
	}
	return ""
}

//获取EnumPrivateFare的整数值
func (this EnumPrivateFare) Value() int {
	if this >= PfUnknownType && this <= PfAirlinePrivateFare {
		return int(this)
	} else {
		return -1
	}
}
