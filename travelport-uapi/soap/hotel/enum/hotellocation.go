package enum

type EnumHotelLocation int

const (
	HlCity EnumHotelLocation = iota
	HlAirport
)

//获取EnumHotelLocation的字符串值
func (this EnumHotelLocation) String() string {
	switch this {
	case HlCity:
		return "City"
	case HlAirport:
		return "Airport"
	}
	return ""
}

//获取EnumHotelLocation的整数值
func (this EnumHotelLocation) Value() int {
	if this >= HlCity && this <= HlAirport {
		return int(this)
	} else {
		return -1
	}
}
