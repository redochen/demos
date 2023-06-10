package biz

import (
	"time"

	"github.com/redochen/demos/mnbz-api/access"
	"github.com/redochen/demos/mnbz-api/models"
	"github.com/redochen/demos/mnbz-api/utils"

	CcStr "github.com/redochen/tools/string"
)

// Register 注册
func Register(model *models.RegisterModel) (int32, error) {
	if nil == model {
		return 0, utils.NewInvalidError("register")
	}

	if model.Password == "" {
		return 0, utils.NewRequiredError("password")
	}

	//检查账号是否已经注册
	user, err := access.CheckRegisterAccount(model.Email, model.Mobile)
	if err != nil {
		return 0, err
	} else if user != nil && user.ID > 0 {
		return 0, utils.NewExistedError("account")
	}

	//检查OpenID是否已记录
	if model.OpenID != "" {
		user, err = access.GetUserByOpenID(model.OpenID, false)
		if err != nil {
			return 0, err
		} else if user != nil && user.ID > 0 {
			return 0, utils.NewExistedError("account")
		}
	}

	company := model.GetCompany()

	//先保存公司，这里需要判断是否已存在
	companyId, err := access.AddCompany(company, false)
	if err != nil {
		return 0, err
	}

	user = model.GetUser()
	user.CompanyID = companyId
	user.LoginAt = utils.FormatDateTime(time.Now())

	_, err = access.AddUser(user)
	if err != nil {
		return 0, err
	}

	found, err := access.GetUserByGUID(user.GUID, false)
	if err != nil {
		return 0, err
	} else if nil == found || found.ID <= 0 {
		return 0, utils.NewFailedError("user", "save")
	}

	return found.ID, nil
}

// Login 登录
func Login(account, password, openID string) (string, error) {
	var user *models.User
	var err error

	if openID != "" { //匿名登录
		user, err = access.GetUserByOpenID(openID, true)
		if err != nil {
			return "", err
		} else if nil == user || user.ID <= 0 {
			return "", utils.NewInvalidError("open ID")
		}

		//记录新用户
		// if nil == user || user.ID <= 0 {
		// 	user = &models.UserEntity{
		// 		WechatOpenID: openID,
		// 	}

		// 	user.GUID, _ = CcStr.NewGUID()

		// 	_, err = access.AddUser(user)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// }
	} else { //账号密码登录
		user, err = access.GetUserByAccount(account, password)
		if err != nil {
			return "", err
		} else if nil == user || user.ID <= 0 {
			return "", utils.NewNotExistedError("account")
		}
	}

	//更新登录时间
	access.UpdateUserLoginTime(user.ID)

	return utils.NewSession(user.ID), nil
}

// UpdateUser 更新用户信息
func UpdateUser(model *models.User, userID int32) error {
	if nil == model {
		return utils.NewInvalidError("user")
	}

	var err error
	var user *models.User

	if model.GUID != "" {
		user, err = access.GetUserByGUID(model.GUID, true)
	} else if model.OpenID != "" {
		user, err = access.GetUserByOpenID(model.OpenID, true)
	}

	if err != nil {
		return err
	} else if nil == user || user.ID <= 0 {
		return utils.NewNotExistedError("user")
	}

	//只允许修改自己的信息
	if userID != user.ID {
		return utils.NewNotAllowedError("user", "modify")
	}

	//user.Account = CcStr.FirstValid(model.Account, user.Account) //账号不允许修改
	//user.Password = CcStr.FirstValid(model.Password, user.Password) //密码走修改密码接口
	user.NickName = CcStr.FirstValid(model.NickName, user.NickName)
	user.Mobile = CcStr.FirstValid(model.Mobile, user.Mobile)
	user.Email = CcStr.FirstValid(model.Email, user.Email)
	user.Wechat = CcStr.FirstValid(model.Wechat, user.Wechat)
	user.OpenID = CcStr.FirstValid(model.OpenID, user.OpenID)
	user.QQ = CcStr.FirstValid(model.QQ, user.QQ)

	_, err = access.UpdateUser(user)
	return err
}
