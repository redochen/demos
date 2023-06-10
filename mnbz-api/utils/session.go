package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redochen/tools/cache"
	CcLog "github.com/redochen/tools/log"
)

//NewSession 获取新的会话ID
func NewSession(userId int32) string {
	if userId <= 0 {
		return ""
	}

	sid := NewGUID()

	duration := time.Hour * 12

	//12小时过期
	cache.SetString(userId, sid, duration)
	cache.SetInt(sid, int(userId), duration)

	return sid
}

//GetUserID 获取用户ID
func GetUserID(ctx *gin.Context, deleteOnGet ...bool) (int32, error) {
	//获取会话ID
	sid := getSessionParameter(ctx)
	if sid == "" {
		return 0, NewRequiredError("session")
	}

	//根据会话ID获取用户ID
	id := cache.GetInt(sid, deleteOnGet...)
	if id <= 0 {
		return 0, NewInvalidError("session")
	}

	CcLog.Debugf("sid=%s,userId=%d", sid, id)

	return int32(id), nil
}

//VerifySession 验证会话ID
func VerifySession(sid string, userId ...int32) error {
	if sid == "" {
		return NewInvalidError("session")
	}

	id := cache.GetInt(sid)
	if id <= 0 {
		return NewInvalidError("session")
	}

	if len(userId) > 0 && id != int(userId[0]) {
		return NewInvalidError("session")
	}

	return nil
}
