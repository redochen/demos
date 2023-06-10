package models

import (
	"github.com/redochen/demos/mnbz-api/utils"
	"time"
)

//TableName 表名
func (e *Need) TableName() string {
	return "need"
}

//Need 需求发布
type Need struct {
	BaseModel      `xorm:"-"`
	ID             int32     `xorm:"pk autoincr index 'id'" json:"id"`                    //ID
	GUID           string    `xorm:"notnull varchar(32) index 'guid'" json:"guid"`        //惟一标识
	UserID         int32     `xorm:"notnull 'userId'" json:"userId"`                      //用户Id
	TripType       string    `xorm:"varchar(5) notnull index 'tripType'" json:"tripType"` //单程OW、往返RT
	FromCity       string    `xorm:"varchar(50) notnull 'fromCity'" json:"fromCity"`      //出发城市
	ToCity         string    `xorm:"varchar(50) notnull 'toCity'" json:"toCity"`          //抵达城市
	DepartDate     string    `xorm:"varchar(50) notnull 'depDate'" json:"depDate"`        //出发日期
	BackDate       string    `xorm:"varchar(50) 'arrDate'" json:"arrDate,omitempty"`      //返程日期
	Airlines       string    `xorm:"varchar(100) 'arilines'" json:"arilines,omitempty"`   //航空公司
	PassengerCount int32     `xorm:"notnull default 1 'count'" json:"count"`              //乘机人数
	Remarks        string    `xorm:"nvarchar(500) 'remarks'" json:"remarks,omitempty"`    //备注
	Status         string    `xorm:"varchar(10) 'status'" json:"status"`                  //状态：publish/private
	CreatedTime    time.Time `xorm:"notnull created 'createdAt'" json:"-"`                //创建时间
	UpdatedTime    time.Time `xorm:"updated 'updatedAt'" json:"-"`                        //修改时间
	DeletedTime    time.Time `xorm:"deleted 'deletedAt'" json:"-"`                        //删除时间
}

//HandleResponse 处理响应结果
func (m *Need) HandleResponse() {
	if m != nil {
		m.CreatedAt = utils.FormatDateTime(m.CreatedTime)
		m.UpdatedAt = utils.FormatDateTime(m.UpdatedTime)
	}
}
