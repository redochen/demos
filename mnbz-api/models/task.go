package models

import (
	"github.com/redochen/demos/mnbz-api/utils"
	"time"
)

// TableName 表名
func (e *Task) TableName() string {
	return "task"
}

// Task 抢票
type Task struct {
	BaseModel      `xorm:"-"`
	ID             int32     `xorm:"pk autoincr index 'id'" json:"id"`             //ID
	GUID           string    `xorm:"notnull varchar(32) index 'guid'" json:"guid"` //惟一标识
	UserID         int32     `xorm:"notnull 'userId'" json:"userId"`               //用户Id
	MaxCount       int32     `xorm:"default 0 'maxCount'" json:"maxCount"`         //最大抢票次数，0表示不限制
	RunCount       int32     `xorm:"default 0 'runCount'" json:"runCount"`         //累计抢票次数
	PNR            string    `xorm:"varchar(6) index 'pnrCode'" json:"pnrCode"`    //PNR
	AvCmd          string    `xorm:"varchar(500) 'avCmd'" json:"avCmd"`            //查询指令
	BookCmd        string    `xorm:"nvarchar(500) 'bookCmd'" json:"bookCmd"`       //预订指令
	Interval       int32     `xorm:"default 10 'interval'" json:"interval"`        //查询间隔，单位秒
	PassengerNames string    `xorm:"varchar(1000) 'passNames'" json:"passNames"`   //旅客姓名
	PassengerDocs  string    `xorm:"varchar(1000) 'passDocs'" json:"passDocs"`     //旅客DOCS
	Status         string    `xorm:"varchar(10) index 'status'" json:"status"`     //状态：stopped|pending|running|done|cancelled|expired
	CreatedTime    time.Time `xorm:"notnull created 'createdAt'" json:"-"`         //创建时间
	UpdatedTime    time.Time `xorm:"updated 'updatedAt'" json:"-"`                 //修改时间
	DeletedTime    time.Time `xorm:"deleted 'deletedAt'" json:"-"`                 //删除时间
}

// HandleResponse 处理响应结果
func (m *Task) HandleResponse() {
	if m != nil {
		m.CreatedAt = utils.FormatDateTime(m.CreatedTime)
		m.UpdatedAt = utils.FormatDateTime(m.UpdatedTime)
	}
}
