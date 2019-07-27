package models

//关系
type RelationModel struct {
	BaseModel
	UserGuid   string `json:"myGuid"`     //我的用户Guid
	FriendGuid string `json:"friendGuid"` //对方用户Guid
	Status     int    `json:"status"`     //-1:已拉黑;0:待验证;1:已验证
	Remarks    string `json:"remarks"`    //备注
}
