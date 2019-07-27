package enum

//Defines the ability to eticket an entity (Yes, No, Required, Ticketless)
type EnumEticketability int

const (
	EtYes EnumEticketability = iota
	EtNo
	EtRequired
	EtTicketless
)

//获取EnumEticketability的字符串值
func (this EnumEticketability) String() string {
	switch this {
	case EtYes:
		return "Yes"
	case EtNo:
		return "No"
	case EtRequired:
		return "Required"
	case EtTicketless:
		return "Ticketless"
	}
	return ""
}

//获取EnumEticketability的整数值
func (this EnumEticketability) Value() int {
	if this >= EtYes && this <= EtTicketless {
		return int(this)
	} else {
		return -1
	}
}
