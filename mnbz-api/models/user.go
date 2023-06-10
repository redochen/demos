package models

import (
	"github.com/redochen/demos/mnbz-api/utils"
	"time"
)

// TableName 表名
func (e *User) TableName() string {
	return "user"
}

// User 用户
type User struct {
	BaseModel   `xorm:"-"`
	ID          int32     `xorm:"pk autoincr index 'id'" json:"id"`                 //ID
	GUID        string    `xorm:"notnull varchar(32) index 'guid'" json:"guid"`     //惟一标识
	CompanyID   int32     `xorm:"notnull 'companyId'" json:"companyId"`             //公司Id
	Password    string    `xorm:"varchar(20) 'password'" json:"password"`           //密码
	NickName    string    `xorm:"nvarchar(50) 'nickName'" json:"nickName"`          //昵称
	Mobile      string    `xorm:"notnull varchar(20) index 'mobile'" json:"mobile"` //手机号码
	Email       string    `xorm:"notnull varchar(50) index 'email'" json:"email"`   //邮箱
	Wechat      string    `xorm:"nvarchar(50) index 'wechat'" json:"wechat"`        //微信
	OpenID      string    `xorm:"varchar(50) index 'wechatOpenId'" json:"-"`        //微信OpenID
	QQ          string    `xorm:"varchar(20) index 'qq'" json:"-"`                  //QQ
	LoginAt     string    `xorm:"varchar(20) 'loginAt'" json:"lastLogin,omitempty"` //最后登录时间，格式为:yyyy-MM-dd HH:mm:ss
	CreatedTime time.Time `xorm:"notnull created 'createdAt'" json:"-"`             //创建时间
	UpdatedTime time.Time `xorm:"updated 'updatedAt'" json:"-"`                     //修改时间
	DeletedTime time.Time `xorm:"deleted 'deletedAt'" json:"-"`                     //删除时间
}

// HideUserDetails 隐藏用户详情
func (u *User) HideUserDetails() {
	if u != nil {
		u.Mobile = ""
		u.Email = ""
		u.Wechat = ""
		u.QQ = ""
	}
}

// HandleResponse 处理响应结果
func (m *User) HandleResponse() {
	if m != nil {
		m.Password = ""
		m.CreatedAt = utils.FormatDateTime(m.CreatedTime)
		m.UpdatedAt = utils.FormatDateTime(m.UpdatedTime)
	}
}
