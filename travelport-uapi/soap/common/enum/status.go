package enum

type EnumStatus int

const (
	Issued       EnumStatus = iota //The service fee has been issued.
	ReadyToIssue                   //The service fee is ready to be issued.
	IssueLater                     //The service fee can be issued later.
)

//获取EnumStatus的字符串值
func (this EnumStatus) String() string {
	switch this {
	case Issued:
		return "Issued"
	case ReadyToIssue:
		return "ReadyToIssue"
	case IssueLater:
		return "IssueLater"
	}
	return ""
}

//获取EnumStatus的整数值
func (this EnumStatus) Value() int {
	if this >= Issued && this <= IssueLater {
		return int(this)
	} else {
		return -1
	}
}
