package access

import (
	"errors"
	"math"
	"time"

	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	"github.com/go-xorm/xorm"
)

//AddUser 添加用户
func AddUser(user *models.User) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == user {
		return 0, utils.NewInvalidError("user")
	} else if user.GUID == "" {
		user.GUID = utils.NewGUID()
	}

	return engine.InsertOne(user)
}

//UpdateUser 更新用户信息
func UpdateUser(user *models.User) (int64, error) {
	err := checkEngine()
	if err != nil {
		return 0, err
	}

	if nil == user {
		return 0, utils.NewInvalidError("user")
	} else if user.ID <= 0 {
		return 0, utils.NewRequiredError("user ID")
	}

	return engine.Id(user.ID).Update(user)
}

//GetUsers 获取用户列表
func GetUsers(pageIndex, pageSize int, hideDetail bool) (users []*models.User, totalCount, pageCount int64, err error) {
	err = checkEngine()
	if err != nil {
		return
	}

	user := new(models.User)
	offset, limit := utils.GetOffsetAndLimit(pageIndex, pageSize)

	totalCount, err = engine.Count(user)
	if err != nil {
		return
	}

	if totalCount > 0 && limit > 0 {
		pageCount = int64(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var rows *xorm.Rows
	rows, err = engine.Desc("loginAt").Limit(limit, offset).Rows(user)
	if err != nil {
		return
	}

	defer rows.Close()

	users = make([]*models.User, 0)

	for rows.Next() {
		err = rows.Scan(user)
		if err != nil {
			return
		}

		if hideDetail {
			user.HideUserDetails()
		}

		user.HandleResponse()
		users = append(users, user)

		//注意：这里应重新分配内存
		user = new(models.User)
	}

	return
}

//UpdateUserLoginTime 更新用户登录时间
func UpdateUserLoginTime(userID int32) error {
	err := checkEngine()
	if err != nil {
		return err
	}

	if userID <= 0 {
		return utils.NewRequiredError("user ID")
	}

	_, err = engine.Id(userID).Cols("loginAt").Update(
		&models.User{LoginAt: utils.FormatDateTime(time.Now())})
	return err
}

//GetUserByID 根据id获取用户
func GetUserByID(id int32, errIfNotExists bool) (*models.User, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if id <= 0 {
		return nil, utils.NewRequiredError("user id")
	}

	user := new(models.User)
	_, err = engine.Id(id).Get(user)

	if err != nil {
		return nil, err
	} else if errIfNotExists && user.ID <= 0 {
		return nil, utils.NewNotExistedError("user")
	}

	user.HandleResponse()

	return user, err
}

//CheckRegisterAccount 检查注册时使用的邮箱或手机号
func CheckRegisterAccount(email, mobile string) (*models.User, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if email == "" && mobile == "" {
		return nil, utils.NewRequiredError("email or mobile")
	}

	user := new(models.User)
	if email == "" {
		_, err = engine.Where("mobile = ?", mobile).Limit(1).Get(user)
	} else if mobile == "" {
		_, err = engine.Where("email = ?", email).Limit(1).Get(user)
	} else {
		_, err = engine.Where("mobile = ?", mobile).
			And("email = ?", email).Limit(1).Get(user)
	}

	if err != nil {
		return nil, err
	}

	user.HandleResponse()

	return user, err
}

//GetUserByAccount 根据账号获取用户
func GetUserByAccount(account, password string) (*models.User, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if account == "" {
		return nil, utils.NewRequiredError("account")
	}

	if password == "" {
		return nil, utils.NewRequiredError("password")
	}

	user := new(models.User)
	_, err = engine.Where("mobile = ?", account).
		Or("email = ?", account).Limit(1).Get(user)

	if err != nil {
		return nil, err
	}

	if user.ID <= 0 {
		return nil, utils.NewNotExistedError("user")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	user.HandleResponse()

	return user, err
}

//GetUserByOpenID 根据微信OpenID获取用户
func GetUserByOpenID(openID string, errIfNotExists bool) (*models.User, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if openID == "" {
		return nil, utils.NewRequiredError("open ID")
	}

	user := new(models.User)
	_, err = engine.Where("wechatOpenId = ?", openID).Limit(1).Get(user)

	if err != nil {
		return nil, err
	} else if errIfNotExists && user.ID <= 0 {
		return nil, utils.NewNotExistedError("user")
	}

	user.HandleResponse()

	return user, err
}

//GetUserByGUID 根据GUID获取用户
func GetUserByGUID(guid string, errIfNotExists bool) (*models.User, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guid == "" {
		return nil, utils.NewRequiredError("user guid")
	}

	user := new(models.User)
	_, err = engine.Where("guid = ?", guid).Limit(1).Get(user)

	if err != nil {
		return nil, err
	} else if errIfNotExists && user.ID <= 0 {
		return nil, utils.NewNotExistedError("user")
	}

	user.HandleResponse()

	return user, err
}

//GetUserByGUIDOrOpenID 根据GUID或微信OpenID获取用户
func GetUserByGUIDOrOpenID(guidOrOpenID string, errIfNotExists bool) (*models.User, error) {
	err := checkEngine()
	if err != nil {
		return nil, err
	}

	if guidOrOpenID == "" {
		return nil, utils.NewRequiredError("user guid or open ID")
	}

	user := new(models.User)
	_, err = engine.Where("guid = ?", guidOrOpenID).
		Or("wechatOpenId = ?", guidOrOpenID).Limit(1).Get(user)

	if err != nil {
		return nil, err
	} else if errIfNotExists && user.ID <= 0 {
		return nil, utils.NewNotExistedError("user")
	}

	user.HandleResponse()

	return user, err
}
