package biz

import (
	"errors"
	"time"

	"github.com/redochen/demos/hdy_api/access"
	"github.com/redochen/demos/hdy_api/entities"
	"github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/utils"
	CcStr "github.com/redochen/tools/string"
)

//Register 注册
func Register(model *models.UserModel) (string, error) {
	if nil == model {
		return "", errors.New("invalid parameters")
	}

	//检查账号是否已经注册
	user, err := access.GetUserByAccount(model.Account, false)
	if err != nil {
		return "", err
	} else if user != nil && user.ID > 0 {
		return "", errors.New("account already exists")
	}

	//检查OpenID是否已记录
	if model.OpenID != "" {
		user, err = access.GetUserByOpenID(model.OpenID)
		if err != nil {
			return "", err
		}
	}

	if user != nil && user.ID > 0 {
		user.Account = model.Account
		user.Password = model.Password
		user.NickName = CcStr.FirstValid(model.NickName, user.NickName)
		user.Avator = CcStr.FirstValid(model.Avator, user.Avator)
		user.Cellphone = CcStr.FirstValid(model.Cellphone, user.Cellphone)
		user.Email = CcStr.FirstValid(model.Email, user.Email)
		user.Wechat = CcStr.FirstValid(model.Wechat, user.Wechat)
		user.WechatOpenID = CcStr.FirstValid(model.OpenID, user.WechatOpenID)
		user.QQ = CcStr.FirstValid(model.QQ, user.QQ)
		user.Signature = CcStr.FirstValid(model.Signature, user.Signature)

		_, err = access.UpdateUser(user)
		if err != nil {
			return "", err
		}

		return user.GUID, nil
	}

	entity := getEntityOfUserModel(model)
	entity.GUID, _ = CcStr.NewGUID()
	entity.LoginAt = time.Now()

	_, err = access.AddUser(entity)
	if err != nil {
		return "", err
	}

	return entity.GUID, nil
}

//Login 登录
func Login(account, password, openID string) (*models.UserModel, error) {
	var user *entities.UserEntity
	var err error

	if account != "" && password != "" { //账号密码登录
		user, err = access.GetUserByAccount(account, true, password)
		if err != nil {
			return nil, err
		}

		if nil == user || user.ID <= 0 {
			return nil, errors.New("account not exists")
		}
	} else { //匿名登录
		if openID == "" {
			return nil, errors.New("invalid openID")
		}

		user, err = access.GetUserByOpenID(openID)
		if err != nil {
			return nil, err
		}

		//记录新用户
		if nil == user || user.ID <= 0 {
			user = &entities.UserEntity{
				WechatOpenID: openID,
			}

			user.GUID, _ = CcStr.NewGUID()

			_, err = access.AddUser(user)
			if err != nil {
				return nil, err
			}
		}
	}

	//更新登录时间
	access.UpdateUserLoginTime(user.ID)

	return getModelOfUserEntity(user, true), nil
}

//UpdateUser 更新用户信息
func UpdateUser(model *models.UserModel) error {
	if nil == model {
		return errors.New("invalid parameters")
	}

	var user *entities.UserEntity
	var err error

	if model.GUID != "" {
		user, err = access.GetUserByGUID(model.GUID)
	} else if model.OpenID != "" {
		user, err = access.GetUserByOpenID(model.OpenID)
	}

	if err != nil {
		return err
	}

	if nil == user || user.ID <= 0 {
		return errors.New("user not exists")
	}

	user.Account = CcStr.FirstValid(model.Account, user.Account)
	user.Password = CcStr.FirstValid(model.Password, user.Password)
	user.NickName = CcStr.FirstValid(model.NickName, user.NickName)
	user.Avator = CcStr.FirstValid(model.Avator, user.Avator)
	user.Cellphone = CcStr.FirstValid(model.Cellphone, user.Cellphone)
	user.Email = CcStr.FirstValid(model.Email, user.Email)
	user.Wechat = CcStr.FirstValid(model.Wechat, user.Wechat)
	user.WechatOpenID = CcStr.FirstValid(model.OpenID, user.WechatOpenID)
	user.QQ = CcStr.FirstValid(model.QQ, user.QQ)
	user.Signature = CcStr.FirstValid(model.Signature, user.Signature)

	_, err = access.UpdateUser(user)
	return err
}

//GetUser 获取用户详情
func GetUser(guid string) (*models.UserModel, error) {
	entity, err := access.GetUserByGUID(guid)
	if err != nil {
		return nil, err
	}

	return getModelOfUserEntity(entity, true), nil
}

//GetUsers 获取用户列表
func GetUsers(pageIndex, pageSize int) (users []*models.UserModel, totalCount, pageCount int64, err error) {
	entities, totalCount, pageCount, err := access.GetUsers(pageIndex, pageSize)
	if err != nil {
		return
	}

	users = make([]*models.UserModel, 0)

	if entities != nil && len(entities) > 0 {
		for _, entity := range entities {
			user := getModelOfUserEntity(entity, false)
			users = append(users, user)
		}
	}

	return
}

func getEntityOfUserModel(model *models.UserModel) *entities.UserEntity {
	if nil == model {
		return nil
	}

	entity := &entities.UserEntity{
		GUID:         model.GUID,
		Account:      model.Account,
		Password:     model.Password,
		NickName:     model.NickName,
		Avator:       model.Avator,
		Cellphone:    model.Cellphone,
		Email:        model.Email,
		Wechat:       model.Wechat,
		WechatOpenID: model.OpenID,
		QQ:           model.QQ,
		Signature:    model.Signature,
		Points:       model.Points,
		LoginAt:      utils.ParseDateTime(model.LastLogin),
		CreatedAt:    utils.ParseDateTime(model.CreatedAt),
		UpdatedAt:    utils.ParseDateTime(model.UpdatedAt),
	}

	return entity
}

func getModelOfUserEntity(entity *entities.UserEntity, details bool) *models.UserModel {
	if nil == entity {
		return nil
	}

	model := &models.UserModel{
		Account:   entity.Account,
		Password:  "", //entity.Password,
		NickName:  entity.NickName,
		Avator:    entity.Avator,
		OpenID:    entity.WechatOpenID,
		Signature: entity.Signature,
		Points:    entity.Points,
		LastLogin: utils.FormatDateTime(entity.LoginAt),
	}

	if details {
		model.Cellphone = entity.Cellphone
		model.Email = entity.Email
		model.Wechat = entity.Wechat
		model.QQ = entity.QQ
	}

	model.GUID = entity.GUID
	model.CreatedAt = utils.FormatDateTime(entity.CreatedAt)
	model.UpdatedAt = utils.FormatDateTime(entity.UpdatedAt)

	return model
}
