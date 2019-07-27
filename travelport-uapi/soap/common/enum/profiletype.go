package enum

//A type for unique party identifiers of any party role.
type EnumProfileType int

const (
	PtAgencyGroup EnumProfileType = iota
	PtAgency
	PtBranchGroup
	PtBranch
	PtAgent
	PtAccount
	PtTravelerGroup
	PtTraveler
)

//获取EnumProfileType的字符串值
func (this EnumProfileType) String() string {
	switch this {
	case PtAgencyGroup:
		return "AgencyGroup"
	case PtAgency:
		return "Agency"
	case PtBranchGroup:
		return "BranchGroup"
	case PtBranch:
		return "Branch"
	case PtAgent:
		return "Agent"
	case PtAccount:
		return "Account"
	case PtTravelerGroup:
		return "TravelerGroup"
	case PtTraveler:
		return "Traveler"
	}
	return ""
}

//获取EnumProfileType的整数值
func (this EnumProfileType) Value() int {
	if this >= PtAgencyGroup && this <= PtTraveler {
		return int(this)
	} else {
		return -1
	}
}
