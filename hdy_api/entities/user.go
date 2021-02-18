package entities

import (
	"time"
)

//TableName 用户表名称
func (e *UserEntity) TableName() string {
	return "user"
}

//UserEntity 用户
type UserEntity struct {
	ID           int64     `xorm:"pk autoincr index"`                 //自增ID
	GUID         string    `xorm:"notnull varchar(32) index"`         //惟一标识
	Account      string    `xorm:"varchar(50) index"`                 //账号
	Password     string    `xorm:"varchar(50) index"`                 //密码
	NickName     string    `xorm:"nvarchar(50) index"`                //昵称
	Avator       string    `xorm:"varchar(200)"`                      //头像URL
	Cellphone    string    `xorm:"varchar(20) index"`                 //手机号码
	Email        string    `xorm:"varchar(50)"`                       //邮箱
	Wechat       string    `xorm:"varchar(50) index"`                 //微信
	WechatOpenID string    `xorm:"varchar(50) index 'wechat_openid'"` //微信OpenID
	QQ           string    `xorm:"varchar(20) index 'qq'"`            //QQ
	Signature    string    `xorm:"nvarchar(100)"`                     //签名
	Points       int       `xorm:"notnull default(0)"`                //积分
	LoginAt      time.Time //最后登录时间
	CreatedAt    time.Time `xorm:"notnull created"` //创建时间
	UpdatedAt    time.Time `xorm:"updated"`         //修改时间
}

//TableName 关系表名称
func (e *RelationEntity) TableName() string {
	return "relation"
}

//RelationEntity 关系
type RelationEntity struct {
	ID           int64     `xorm:"pk autoincr index"`         //自增ID
	GUID         string    `xorm:"notnull varchar(32) index"` //惟一标识
	UserID       int64     `xorm:"notnull index"`             //用户ID
	FriendUserID int64     `xorm:"notnull index"`             //对方用户ID
	Status       int       `xorm:"notnull default(0) index"`  //-1:已拉黑;0:待验证;1:已验证
	Remarks      string    `xorm:"nvarchar(100)"`             //备注
	CreatedAt    time.Time `xorm:"notnull created"`           //创建时间
	UpdatedAt    time.Time `xorm:"updated"`                   //修改时间
}

//TableName 邀请表名称
func (e *InvitationEntity) TableName() string {
	return "invitation"
}

//InvitationEntity 邀请
type InvitationEntity struct {
	ID           int64     `xorm:"pk autoincr index"`         //自增ID
	GUID         string    `xorm:"notnull varchar(32) index"` //惟一标识
	UserID       int64     `xorm:"notnull index"`             //用户ID
	FriendUserID int64     `xorm:"notnull index"`             //对方用户ID
	Message      string    `xorm:"nvarchar(200)"`             //消息
	Status       int       `xorm:"notnull default(0) index"`  //-1:已拒绝;0-待接受;1-已接受
	PlayTime     time.Time `xorm:"notnull"`                   //开黑时间
	CreatedAt    time.Time `xorm:"notnull created"`           //创建时间
}

//TableName 评价表名称
func (e *EvaluationEntity) TableName() string {
	return "evaluation"
}

//EvaluationEntity 评价
type EvaluationEntity struct {
	ID           int64     `xorm:"pk autoincr index"`         //自增ID
	GUID         string    `xorm:"notnull varchar(32) index"` //惟一标识
	InvitationID int64     `xorm:"notnull index"`             //相关的邀请ID
	Content      string    `xorm:"notnull nvarchar(500)"`     //评价内容
	ReplyTo      int       `xorm:"notnull default(0)"`        //回复评价的原始评价ID
	IsDeleted    bool      `xorm:"notnull default(0) index"`  //是否已删除
	CreatedAt    time.Time `xorm:"notnull created"`           //创建时间
	UpdatedAt    time.Time `xorm:"updated"`                   //修改时间
	DeletedAt    time.Time `xorm:"deleted"`                   //删除时间
}
