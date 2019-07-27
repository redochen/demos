package enum

//ATA/IATA Standard commission level.
type EnumCommissionLevel int

const (
	ClRecalled EnumCommissionLevel = iota
	ClFare
	ClPenalty
)

//获取EnumCommissionLevel的字符串值
func (this EnumCommissionLevel) String() string {
	switch this {
	case ClRecalled:
		return "Recalled"
	case ClFare:
		return "Fare"
	case ClPenalty:
		return "Penalty"
	}
	return ""
}

//获取EnumCommissionLevel的整数值
func (this EnumCommissionLevel) Value() int {
	if this >= ClRecalled && this <= ClPenalty {
		return int(this)
	} else {
		return -1
	}
}
