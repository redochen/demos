package enum

//Kestrel Long Fare Rule Category Codes
type EnumFareRuleCategoryCode int

const (
	APP EnumFareRuleCategoryCode = iota //Rule App/Other Conditions
	WHO                                 //Eligibility
	DAY                                 //Day/Time
	SEA                                 //Seasonal
	FLT                                 //Flight App
	ADV                                 //Advance Res/Tkt
	MIN                                 //Minimum Stay
	MAX                                 //Maximum Stay
	STP                                 //Stopovers
	TRF                                 //Transfers/Routing
	CMB                                 //Combinability
	BLA                                 //Blackouts
	SUR                                 //Surcharges
	ACC                                 //Accompanied
	TVL                                 //Travel Restrictions
	TKT                                 //Sales Restrictions
	CHG                                 //Penalties
	HIP                                 //HIP and Mileage Exceptions
	END                                 //Ticket Endorsements
	CHD                                 //Children"s Discounts
	TUC                                 //Tour Conductor Disc
	AGT                                 //Agent Discounts
	DSC                                 //All Other Disc
	MIS                                 //Misc Fare Tags
	FBR                                 //Fare By Rule
	GRP                                 //Groups
	TUR                                 //Tours
	VAC                                 //Visit Another Country
	DEP                                 //Deposits
	VOL                                 //Voluntary Changes
	IVE                                 //Involuntary Exchanges
	VOR                                 //Voluntary Refunds
	IVR                                 //Involuntary Refunds
	NET                                 //Negotiated Fares
	OTH                                 //Other
)

//获取EnumFareRuleCategoryCode的字符串值
func (this EnumFareRuleCategoryCode) String() string {
	switch this {
	case APP:
		return "APP"
	case WHO:
		return "WHO"
	case DAY:
		return "DAY"
	case SEA:
		return "SEA"
	case FLT:
		return "FLT"
	case ADV:
		return "ADV"
	case MIN:
		return "MIN"
	case MAX:
		return "MAX"
	case STP:
		return "STP"
	case TRF:
		return "TRF"
	case CMB:
		return "CMB"
	case BLA:
		return "BLA"
	case SUR:
		return "SUR"
	case ACC:
		return "ACC"
	case TVL:
		return "TVL"
	case TKT:
		return "TKT"
	case CHG:
		return "CHG"
	case HIP:
		return "HIP"
	case END:
		return "END"
	case CHD:
		return "CHD"
	case TUC:
		return "TUC"
	case AGT:
		return "AGT"
	case DSC:
		return "DSC"
	case MIS:
		return "MIS"
	case FBR:
		return "FBR"
	case GRP:
		return "GRP"
	case TUR:
		return "TUR"
	case VAC:
		return "VAC"
	case DEP:
		return "DEP"
	case VOL:
		return "VOL"
	case IVE:
		return "IVE"
	case VOR:
		return "VOR"
	case IVR:
		return "IVR"
	case NET:
		return "NET"
	case OTH:
		return "OTH"
	}
	return ""
}

//获取EnumFareRuleCategoryCode的整数值
func (this EnumFareRuleCategoryCode) Value() int {
	if this >= APP && this <= OTH {
		return int(this)
	} else {
		return -1
	}
}
