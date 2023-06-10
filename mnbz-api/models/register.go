package models

import "github.com/redochen/demos/mnbz-api/status"

//RegisterModel 用户注册模型
type RegisterModel struct {
	CompanyName  string `json:"companyName"`  //公司名称
	Country      string `json:"country"`      //所在国家
	City         string `json:"city"`         //所在城市
	Password     string `json:"password"`     //密码
	ContractUser string `json:"contractUser"` //联系人
	NickName     string `json:"nickName"`     //昵称
	Mobile       string `json:"mobile"`       //手机号码
	Email        string `json:"email"`        //邮箱
	Wechat       string `json:"wechat"`       //微信
	OpenID       string `json:"wechatOpenid"` //微信OpenID
	QQ           string `json:"-"`            //QQ
}

//GetCompany 获取注册公司信息
func (r *RegisterModel) GetCompany() *Company {
	if nil == r {
		return nil
	}

	return &Company{
		Name:         r.CompanyName,  //公司名称
		Country:      r.Country,      //所在国家
		City:         r.City,         //所在城市
		ContractUser: r.ContractUser, //联系人
	}
}

//GetUser 获取注册用户信息
func (r *RegisterModel) GetUser() *User {
	if nil == r {
		return nil
	}

	return &User{
		Password: r.Password, //密码
		NickName: r.NickName, //昵称
		Mobile:   r.Mobile,   //手机号码
		Email:    r.Email,    //邮箱
		Wechat:   r.Wechat,   //微信
		OpenID:   r.OpenID,   //微信OpenID
		QQ:       r.QQ,       //QQ
	}
}

//RegisterResult 注册结果
type RegisterResult struct {
	BaseResult
	UserID int32 `json:"userId,omitempty"` //用户ID
}

//NewRegisterResult 获取注册结果
func NewRegisterResult(userId int32) *RegisterResult {
	result := &RegisterResult{
		UserID: userId,
	}

	result.SetError(status.Success)
	return result
}
