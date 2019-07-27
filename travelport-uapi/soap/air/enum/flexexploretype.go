package enum

//Type of flex explore to be performed
type EnumFlexExploreType int

const (
	FetAnyWhere EnumFlexExploreType = iota
	FetArea
	FetZone
	FetCountry
	FetState
	FetDistanceInMiles
	FetDistanceInKilometers
	FetDestination
	FetGroup
)

//获取EnumFlexExploreType的字符串值
func (this EnumFlexExploreType) String() string {
	switch this {
	case FetAnyWhere:
		return "AnyWhere"
	case FetArea:
		return "Area"
	case FetZone:
		return "Zone"
	case FetCountry:
		return "Country"
	case FetState:
		return "State"
	case FetDistanceInMiles:
		return "DistanceInMiles"
	case FetDistanceInKilometers:
		return "DistanceInKilometers"
	case FetDestination:
		return "Destination"
	case FetGroup:
		return "Group"
	}
	return ""
}

//获取EnumFlexExploreType的整数值
func (this EnumFlexExploreType) Value() int {
	if this >= FetAnyWhere && this <= FetGroup {
		return int(this)
	} else {
		return -1
	}
}
