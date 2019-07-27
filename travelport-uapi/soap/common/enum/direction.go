package enum

type EnumDirection int

const (
	South EnumDirection = iota
	North
	East
	West
	SouthEast
	SouthWest
	NorthEast
	NorthWest
)

//获取EnumDirection的字符串值
func (this EnumDirection) String() string {
	switch this {
	case South:
		return "S"
	case North:
		return "N"
	case East:
		return "E"
	case West:
		return "W"
	case SouthEast:
		return "SE"
	case SouthWest:
		return "NW"
	case NorthEast:
		return "SE"
	case NorthWest:
		return "NW"
	}
	return ""
}

//获取EnumDirection的整数值
func (this EnumDirection) Value() int {
	if this >= South && this <= NorthWest {
		return int(this)
	} else {
		return -1
	}
}
