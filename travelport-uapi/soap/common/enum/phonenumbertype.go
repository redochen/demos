package enum

type EnumPhoneNumberType int

const (
	PntAgency EnumPhoneNumberType = iota
	PntBusiness
	PntMobile
	PntHome
	PntFax
	PntHotel
	PntOther
	PntNone
	PntEmail
	PntReservations
)

//获取EnumPhoneNumberType的字符串值
func (this EnumPhoneNumberType) String() string {
	switch this {
	case PntAgency:
		return "Agency"
	case PntBusiness:
		return "Business"
	case PntMobile:
		return "Mobile"
	case PntHome:
		return "Home"
	case PntFax:
		return "Fax"
	case PntHotel:
		return "Hotel"
	case PntOther:
		return "Other"
	case PntNone:
		return "None"
	case PntEmail:
		return "Email"
	case PntReservations:
		return "Reservations"
	}
	return ""
}

//获取EnumPhoneNumberType的整数值
func (this EnumPhoneNumberType) Value() int {
	if this >= PntAgency && this <= PntReservations {
		return int(this)
	} else {
		return -1
	}
}
