package access

import (
	"errors"
	"math"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/utils"
	CcStr "github.com/redochen/tools/string"
)

//AddUser 添加用户
func AddUser(user *entities.UserEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == user {
		return 0, errors.New("invalid parameter")
	}

	return engine.InsertOne(user)
}

//UpdateUser 更新用户信息
func UpdateUser(user *entities.UserEntity) (int64, error) {
	if nil == engine {
		return 0, errors.New("engine not initialized")
	}

	if nil == user {
		return 0, errors.New("invalid parameter")
	}

	if user.ID <= 0 {
		return 0, errors.New("invalid user ID")
	}

	return engine.Id(user.ID).Update(user)
}

//GetUsers 获取用户列表
func GetUsers(pageIndex, pageSize int) (users []*entities.UserEntity, totalCount, pageCount int64, err error) {
	if nil == engine {
		err = errors.New("engine not initialized")
		return
	}

	user := new(entities.UserEntity)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(user)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Desc("login_at").Limit(limit, offset).Rows(user)
	if err != nil {
		return
	}

	defer rows.Close()

	users = make([]*entities.UserEntity, 0)

	for rows.Next() {
		err = rows.Scan(user)
		if err != nil {
			return
		}

		users = append(users, user)

		//注意：这里应重新分配内存
		user = new(entities.UserEntity)
	}

	return
}

//UpdateUserLoginTime 更新用户登录时间
func UpdateUserLoginTime(userID int64) error {
	if nil == engine {
		return errors.New("engine not initialized")
	}

	if userID <= 0 {
		return errors.New("invalid user ID")
	}

	_, err := engine.Id(userID).Cols("login_at").Update(&entities.UserEntity{LoginAt: time.Now()})
	return err
}

//GetUserByAccount 根据账号获取用户
func GetUserByAccount(account string, checkPass bool, password ...string) (*entities.UserEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if account == "" {
		return nil, errors.New("invalid account")
	}

	var pass string

	if checkPass {
		if pass = CcStr.FirstValid(password...); pass == "" {
			return nil, errors.New("invalid password")
		}
	}

	var user entities.UserEntity
	var err error

	if checkPass {
		_, err = engine.Where("account = ?", account).And("password = ?", pass).Limit(1).Get(&user)
	} else {
		_, err = engine.Where("account = ?", account).Limit(1).Get(&user)
	}

	return &user, err
}

//GetUserByOpenID 根据微信OpenID获取用户
func GetUserByOpenID(openID string) (*entities.UserEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if openID == "" {
		return nil, errors.New("invalid OpenID")
	}

	var user entities.UserEntity
	_, err := engine.Where("wechat_openid = ?", openID).Limit(1).Get(&user)

	return &user, err
}

//GetUserByGUID 根据GUID获取用户
func GetUserByGUID(guid string) (*entities.UserEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if guid == "" {
		return nil, errors.New("invalid user Guid")
	}

	var user entities.UserEntity
	_, err := engine.Where("guid = ?", guid).Limit(1).Get(&user)

	return &user, err
}

//GetUserByGUIDOrOpenID 根据GUID或微信OpenID获取用户
func GetUserByGUIDOrOpenID(guidOrOpenID string) (*entities.UserEntity, error) {
	if nil == engine {
		return nil, errors.New("engine not initialized")
	}

	if guidOrOpenID == "" {
		return nil, errors.New("invalid parameter")
	}

	var user entities.UserEntity
	_, err := engine.Where("guid = ?", guidOrOpenID).
		Or("wechat_openid = ?", guidOrOpenID).Limit(1).Get(&user)

	return &user, err
}
