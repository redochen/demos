package enum

//The modification types supported
type EnumModificationType int

const (
	MtAddSegment      EnumModificationType = iota //Add a segment to the itinerary
	MtRemoveSegment                               //Delete a segment from the itinerary
	MtReplaceSegment                              //Replace a segment in the itinerary with a new segment
	MtAddPassenger                                //Add a passenger to the itinerary
	MtRemovePassenger                             //Remove a passenger from the itinerary
	MtOptionsOnly                                 //Modification where only options are added / removed from the itinerary
	MtOther                                       //Other modification types
)

//获取EnumModificationType的字符串值
func (this EnumModificationType) String() string {
	switch this {
	case MtAddSegment:
		return "AddSegment"
	case MtRemoveSegment:
		return "RemoveSegment"
	case MtReplaceSegment:
		return "ReplaceSegment"
	case MtAddPassenger:
		return "AddPassenger"
	case MtRemovePassenger:
		return "RemovePassenger"
	case MtOptionsOnly:
		return "OptionsOnly"
	case MtOther:
		return "Other"
	}
	return ""
}

//获取EnumModificationType的整数值
func (this EnumModificationType) Value() int {
	if this >= MtAddSegment && this <= MtOther {
		return int(this)
	} else {
		return -1
	}
}
