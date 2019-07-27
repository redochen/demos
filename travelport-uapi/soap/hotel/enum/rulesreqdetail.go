package enum

//Request details for Rules, Details, or All. Default is All. Applicable for 1p/1j.
type EnumRulesRequestDetails int

const (
	RdRules EnumRulesRequestDetails = iota
	RdDetails
	RdAll
)

//获取EnumRulesRequestDetails的字符串值
func (this EnumRulesRequestDetails) String() string {
	switch this {
	case RdRules:
		return "Rules"
	case RdDetails:
		return "Details"
	case RdAll:
		return "All"
	}
	return ""
}

//获取EnumRulesRequestDetails的整数值
func (this EnumRulesRequestDetails) Value() int {
	if this >= RdRules && this <= RdAll {
		return int(this)
	} else {
		return -1
	}
}
