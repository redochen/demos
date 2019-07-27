package enum

//The type of action the agent performed.
type EnumAgentActionType int

const (
	AatCreated EnumAgentActionType = iota
	AatModified
	AatTicketed
)

//获取EnumAgentActionType的字符串值
func (this EnumAgentActionType) String() string {
	switch this {
	case AatCreated:
		return "Created"
	case AatModified:
		return "Modified"
	case AatTicketed:
		return "Ticketed"

	}
	return ""
}

//获取EnumAgentActionType的整数值
func (this EnumAgentActionType) Value() int {
	if this >= AatCreated && this <= AatTicketed {
		return int(this)
	} else {
		return -1
	}
}
