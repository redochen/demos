package models

import (
	"github.com/redochen/demos/hdy_api/status"
)

//用户
type UserModel struct {
	BaseModel
	Account   string `json:"account"`             //账号
	Password  string `json:"pass,omitempty"`      //密码
	NickName  string `json:"nick,omitempty"`      //昵称
	Avator    string `json:"avator,omitempty"`    //头像URL
	Cellphone string `json:"mobile,omitempty"`    //手机号码
	Email     string `json:"email,omitempty"`     //邮箱
	Wechat    string `json:"wechat,omitempty"`    //微信
	OpenID    string `json:"openId"`              //微信OpenID
	QQ        string `json:"qq,omitempty"`        //QQ
	Signature string `json:"sign,omitempty"`      //签名
	Points    int    `json:"points"`              //积分
	LastLogin string `json:"lastLogin,omitempty"` //最后登录时间，格式为:yyyy-MM-dd HH:mm:ss
}

//注册结果
type RegisterResult struct {
	BaseResult
	Guid string `json:"guid,omitempty"` //惟一标识
}

//获取注册结果
func NewRegisterResult(guid string) *RegisterResult {
	result := &RegisterResult{
		Guid: guid,
	}

	result.SetError(status.Success)
	return result
}

//用户结果
type UserResult struct {
	BaseResult
	User *UserModel `json:"user, omitempty"` //用户信息
}

//获取登录结果
func NewUserResult(user *UserModel) *UserResult {
	result := &UserResult{
		User: user,
	}

	result.SetError(status.Success)
	return result
}
