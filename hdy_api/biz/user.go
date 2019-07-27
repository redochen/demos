package biz

import (
	"errors"
	"github.com/redochen/demos/hdy_api/access"
	. "github.com/redochen/demos/hdy_api/entities"
	. "github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/utils"
	. "github.com/redochen/tools/string"
	"time"
)

//注册
func Register(model *UserModel) (string, error) {
	if nil == model {
		return "", errors.New("invalid parameters")
	}

	//检查账号是否已经注册
	user, err := access.GetUserByAccount(model.Account, false)
	if err != nil {
		return "", err
	} else if user != nil && user.Id > 0 {
		return "", errors.New("account already exists")
	}

	//检查OpenID是否已记录
	if model.OpenID != "" {
		user, err = access.GetUserByOpenId(model.OpenID)
		if err != nil {
			return "", err
		}
	}

	if user != nil && user.Id > 0 {
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

		return user.Guid, nil
	}

	entity := getEntityOfUserModel(model)
	entity.Guid, _ = CcStr.NewGuid()
	entity.LoginAt = time.Now()

	_, err = access.AddUser(entity)
	if err != nil {
		return "", err
	}

	return entity.Guid, nil
}

//登录
func Login(account, password, openId string) (*UserModel, error) {
	var user *UserEntity
	var err error

	if account != "" && password != "" { //账号密码登录
		user, err = access.GetUserByAccount(account, true, password)
		if err != nil {
			return nil, err
		}

		if nil == user || user.Id <= 0 {
			return nil, errors.New("account not exists")
		}
	} else { //匿名登录
		if openId == "" {
			return nil, errors.New("invalid openID")
		}

		user, err = access.GetUserByOpenId(openId)
		if err != nil {
			return nil, err
		}

		//记录新用户
		if nil == user || user.Id <= 0 {
			user = &UserEntity{
				WechatOpenID: openId,
			}

			user.Guid, _ = CcStr.NewGuid()

			_, err = access.AddUser(user)
			if err != nil {
				return nil, err
			}
		}
	}

	//更新登录时间
	access.UpdateUserLoginTime(user.Id)

	return getModelOfUserEntity(user, true), nil
}

//更新用户信息
func UpdateUser(model *UserModel) error {
	if nil == model {
		return errors.New("invalid parameters")
	}

	var user *UserEntity
	var err error

	if model.Guid != "" {
		user, err = access.GetUserByGuid(model.Guid)
	} else if model.OpenID != "" {
		user, err = access.GetUserByOpenId(model.OpenID)
	}

	if err != nil {
		return err
	}

	if nil == user || user.Id <= 0 {
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

//获取用户详情
func GetUser(guid string) (*UserModel, error) {
	entity, err := access.GetUserByGuid(guid)
	if err != nil {
		return nil, err
	}

	return getModelOfUserEntity(entity, true), nil
}

//获取用户列表
func GetUsers(pageIndex, pageSize int) (models []*UserModel, totalCount, pageCount int64, err error) {
	entities, totalCount, pageCount, err := access.GetUsers(pageIndex, pageSize)
	if err != nil {
		return
	}

	models = make([]*UserModel, 0)

	if entities != nil && len(entities) > 0 {
		for _, entity := range entities {
			model := getModelOfUserEntity(entity, false)
			models = append(models, model)
		}
	}

	return
}

func getEntityOfUserModel(model *UserModel) *UserEntity {
	if nil == model {
		return nil
	}

	entity := &UserEntity{
		Guid:         model.Guid,
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

func getModelOfUserEntity(entity *UserEntity, details bool) *UserModel {
	if nil == entity {
		return nil
	}

	model := &UserModel{
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

	model.Guid = entity.Guid
	model.CreatedAt = utils.FormatDateTime(entity.CreatedAt)
	model.UpdatedAt = utils.FormatDateTime(entity.UpdatedAt)

	return model
}
