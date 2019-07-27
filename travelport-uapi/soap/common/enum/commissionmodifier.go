package enum

//Optional commission modifier.
type EnumCommissionModifier int

const (
	CmFarePercent                      EnumCommissionModifier = iota //Commission percentage applied to the fare
	CmFareAmount                                                     //Commission amount applied to the fare
	CmCommissionAmount                                               //Specific commission amount to be applied
	CmLessStandardCommission                                         //Indicates commission percentage applied to the fare less the standard commission
	CmStandardPlusSupplementaryPercent                               //Indicates commission percentage includes standard and supplementary commission
	CmSupplementaryPercent                                           //Supplementary commission percent which is applied to the fare
	CmSupplementaryAmount                                            //Supplementary commission amount which is applied to the fare
)

//获取EnumCommissionModifier的字符串值
func (this EnumCommissionModifier) String() string {
	switch this {
	case CmFarePercent:
		return "FarePercent"
	case CmFareAmount:
		return "FareAmount"
	case CmCommissionAmount:
		return "CommissionAmount"
	case CmLessStandardCommission:
		return "LessStandardCommission"
	case CmStandardPlusSupplementaryPercent:
		return "StandardPlusSupplementaryPercent"
	case CmSupplementaryPercent:
		return "SupplementaryPercent"
	case CmSupplementaryAmount:
		return "SupplementaryAmount"
	}
	return ""
}

//获取EnumCommissionModifier的整数值
func (this EnumCommissionModifier) Value() int {
	if this >= CmFarePercent && this <= CmSupplementaryAmount {
		return int(this)
	} else {
		return -1
	}
}
