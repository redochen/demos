package models

import (
	"github.com/redochen/demos/mnbz-api/utils"
	"time"
)

// TableName 表名
func (e *Company) TableName() string {
	return "company"
}

// Company 公司名称
type Company struct {
	BaseModel    `xorm:"-"`
	ID           int32     `xorm:"pk autoincr index 'id'" json:"id"`                     //ID
	GUID         string    `xorm:"notnull varchar(32) index 'guid'" json:"guid"`         //惟一标识
	Name         string    `xorm:"nvarchar(50) notnull index 'name'" json:"companyName"` //公司名称
	Country      string    `xorm:"varchar(50) 'country'" json:"country"`                 //所在国家
	City         string    `xorm:"varchar(100) 'city'" json:"city"`                      //所在城市
	ContractUser string    `xorm:"varchar(50) 'contract'" json:"contractUser"`           //联系人
	CreatedTime  time.Time `xorm:"notnull created 'createdAt'" json:"-"`                 //创建时间
	UpdatedTime  time.Time `xorm:"updated 'updatedAt'" json:"-"`                         //修改时间
	DeletedTime  time.Time `xorm:"deleted 'deletedAt'" json:"-"`                         //删除时间
}

// HandleResponse 处理响应结果
func (m *Company) HandleResponse() {
	if m != nil {
		m.CreatedAt = utils.FormatDateTime(m.CreatedTime)
		m.UpdatedAt = utils.FormatDateTime(m.UpdatedTime)
	}
}
