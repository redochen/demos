package enum

//The different levels at which an optional service may be applied
type EnumOptionalServiceApplicabilityType int

const (
	OsatItinerary EnumOptionalServiceApplicabilityType = iota
	OsatPassenger
	OsatSegment
	OsatPassengerSegment
	OsatPassengerOD //PassengerOD stands for passenger origin destination.
	OsatOther
)

//获取EnumOptionalServiceApplicabilityType的字符串值
func (this EnumOptionalServiceApplicabilityType) String() string {
	switch this {
	case OsatItinerary:
		return "Itinerary"
	case OsatPassenger:
		return "Passenger"
	case OsatSegment:
		return "Segment"
	case OsatPassengerSegment:
		return "PassengerSegment"
	case OsatPassengerOD:
		return "PassengerOD"
	case OsatOther:
		return "Other"
	}
	return ""
}

//获取EnumOptionalServiceApplicabilityType的整数值
func (this EnumOptionalServiceApplicabilityType) Value() int {
	if this >= OsatItinerary && this <= OsatOther {
		return int(this)
	} else {
		return -1
	}
}
