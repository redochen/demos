package models

import (
	"github.com/redochen/demos/mnbz-api/utils"
	"time"
)

// TableName 表名
func (e *Config) TableName() string {
	return "config"
}

// Config 配置
type Config struct {
	BaseModel   `xorm:"-"`
	ID          int32     `xorm:"pk autoincr index 'id'" json:"id"`                    //ID
	GUID        string    `xorm:"notnull varchar(32) index 'guid'" json:"guid"`        //惟一标识
	UserID      int32     `xorm:"notnull 'userId'" json:"userId"`                      //用户Id
	GdsType     string    `xorm:" varchar(10) notnull index 'gdsType'" json:"gdsType"` //GDS类型：1G (1G/1A/1S/1E四选1)
	PCC         string    `xorm:" varchar(10) notnull index 'pcc'" json:"pcc"`         //PCC
	Account     string    `xorm:"varchar(50) 'account'" json:"account"`                //账号
	Password    string    `xorm:"varchar(50) 'password'" json:"password"`              //密码
	Status      string    `xorm:"varchar(10) 'status'" json:"status"`                  //状态：normal/expired/invalid
	CreatedTime time.Time `xorm:"notnull created 'createdAt'" json:"-"`                //创建时间
	UpdatedTime time.Time `xorm:"updated 'updatedAt'" json:"-"`                        //修改时间
	DeletedTime time.Time `xorm:"deleted 'deletedAt'" json:"-"`                        //删除时间
}

// HandleResponse 处理响应结果
func (m *Config) HandleResponse() {
	if m != nil {
		m.CreatedAt = utils.FormatDateTime(m.CreatedTime)
		m.UpdatedAt = utils.FormatDateTime(m.UpdatedTime)
	}
}
