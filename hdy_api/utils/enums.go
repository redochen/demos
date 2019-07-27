package utils

type RelationStatus int

const (
	StatusAll = 100000
)

const (
	RelationStatusAll      = StatusAll //所有状态
	RelationStatusDefault  = 0         //待验证
	RelationStatusBlacked  = -1        //已拉黑
	RelationStatusAccepted = 1         //已验证
)

const (
	InvitationStatusAll      = StatusAll //所有邀请
	InvitationStatusDefault  = 0         //待接受
	InvitationStatusRefused  = -1        //已拒绝
	InvitationStatusAccepted = 1         //已接受
)
