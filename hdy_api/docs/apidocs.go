package docs

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	. "github.com/redochen/demos/hdy_api/models"
	"github.com/redochen/demos/hdy_api/status"
	"github.com/redochen/demos/hdy_api/utils"
	. "github.com/redochen/tools/json"
	"html/template"
	"net/http"
	"time"
)

//接口文档列表
func ApiDocs(ctx *gin.Context) {
	var list string

	list += "<p><ul>"
	list += "<strong>游戏</strong>"
	list += "</br>"
	list += "<li><a href='/docs/games'>获取游戏列表</a></li>"
	list += "<li><a href='/docs/game/areas'>获取大区列表</a></li>"
	list += "<li><a href='/docs/game/servers'>获取服务器列表</a></li>"
	list += "<li><a href='/docs/game/dans'>获取段位列表</a></li>"
	list += "<li><a href='/docs/game/heroes'>获取英雄列表</a></li>"
	list += "</ul></p>"

	list += "<p><ul>"
	list += "<strong>用户</strong>"
	list += "</br>"
	list += "<li><a href='/docs/user/register'>用户注册</a></li>"
	list += "<li><a href='/docs/user/login'>用户登录</a></li>"
	list += "<li><a href='/docs/user/update'>更新用户信息</a></li>"
	list += "<li><a href='/docs/user/details'>获取用户详情</a></li>"
	list += "<li><a href='/docs/users'>获取用户列表</a></li>"
	list += "</ul></p>"

	list += "<p><ul>"
	list += "<strong>其他</strong>"
	list += "</br>"
	list += "<li><a href='/docs/captcha'>获取图片、音频验证码</a></li>"
	list += "<li><a href='/docs/captcha/sms'>获取短信验证码</a></li>"
	list += "<li><a href='/docs/captcha/verify'>验证图片、音频验证码</a></li>"
	list += "<li><a href='/docs/captcha/resource'>图片、音频验证码链接URL地址说明文档</a></li>"
	list += "</ul></p>"

	ctx.HTML(http.StatusOK, "api_doc_list.tmpl", gin.H{
		"list": template.HTML(list),
	})
}

//获取图片、音频验证码接口文档
func ApidocCaptcha(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取图片、音频验证码接口",
		"url": template.HTML(`/api/captcha?length=6
			<br/><br/>length：验证码的长度，默认长度为8，可选参数。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": template.HTML(`图片、音频验证码的资源ID，详情请参考 <a href="/docs/captcha/sample" target="_blank">图片、音频验证码示例</a>。
		<br/>
		<br/><span style="color:red;">注意：验证码10分钟内有效，不可重复使用</span>`),
	})
}

//获取短信验证码接口文档
func ApiDocSmsCaptcha(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取短信验证码接口",
		"url": template.HTML(`/api/captcha/sms?length=6&cellphone=13xxxxxxxxx&captchaId=xxxxxx&captchaCode=xxxxxx&timeoutSeconds=30
			<br/><br/>length：验证码的长度，默认长度为8，可选参数。
			<br/>cellphone：接收验证码的手机号，必填参数。
			<br/>captchaId：图片、音频验证码的资源ID，可通过 <a href="/docs/captcha" target="_blank">图片、音频验证码接口</a> 获取，必填参数。
			<br/>captchaCode：验证码，可通过查看验证码图片或收听验证码音频获取，必填参数。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": template.HTML(CcJson.GetString(NewResult(status.Success)) + `<br/><br/><span style=color:red;">注意：验证码10分钟内有效，不可重复使用</span>`),
	})
}

//验证图片、音频验证码接口文档
func ApiDocVerifyCaptcha(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "验证图片、音频验证码接口",
		"url": template.HTML(`/api/captcha/verify?captchaId=xxxxxx&captchaCode=xxxxxx
			<br/><br/>captchaId：图片、音频验证码的资源ID，可通过 <a href="/docs/captcha" target="_blank">图片、音频验证码接口</a> 获取，必填参数。
			<br/>captchaCode：验证码，可通过查看验证码图片或收听验证码音频获取，必填参数。`),
		"method": "POST",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewResult(status.Success)),
	})
}

//图片、音频验证码链接URL地址说明文档
func ApiDocCaptchaResource(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "图片、音频验证码链接URL地址",
		"url": template.HTML(`/captcha/captchaId.png 或 /captcha/captchaId.wav
			<br/><br/>captchaId：图片、音频验证码的资源ID，可通过 <a href="/docs/captcha" target="_blank">图片、音频验证码接口</a> 获取，必填参数。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": template.HTML(`图片或音频资源，详情请参考 <a href="/docs/captcha/sample" target="_blank">图片、音频验证码示例</a>。
		<br/>
		<br/><span style="color:red;">注意：验证码10分钟内有效，不可重复使用</span>`),
	})
}

//验证码示例文档
func ApiDocCaptchaSample(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "captcha_sample.tmpl", gin.H{
		"CaptchaId": captcha.New(),
	})
}

//获取游戏列表接口文档
func ApiDocGames(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取游戏列表接口",
		"url": template.HTML(`/api/games?timeoutSeconds=30
			<br/><br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewListResult(getGames())),
	})
}

//获取游戏大区列表接口文档
func ApiDocGameAreas(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取游戏大区列表接口",
		"url": template.HTML(`/api/game/areas?gamdId=10001&timeoutSeconds=30
			<br/><br/>gamdId：游戏ID，可通过 <a href="/docs/games" target="_blank">游戏列表接口</a> 获取，必填参数。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewListResult(getGameAreas())),
	})
}

//获取游戏服务器列表接口文档
func ApiDocGameServers(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取游戏服务器列表接口",
		"url": template.HTML(`/api/game/servers?areaId=10001&timeoutSeconds=30
			<br/><br/>areaId：游戏大区ID，可通过 <a href="/docs/game/areas" target="_blank">游戏大区列表接口</a> 获取，必填参数。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewListResult(getGameServers())),
	})
}

//获取游戏段位列表接口文档
func ApiDocGameDans(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取游戏段位列表接口",
		"url": template.HTML(`/api/game/dans?gamdId=10001&timeoutSeconds=30
			<br/><br/>gamdId：游戏ID，可通过 <a href="/docs/games" target="_blank">游戏列表接口</a> 获取，必填参数。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "无参见URL说明",
		"result": CcJson.GetString(NewListResult(getGameDans())),
	})
}

//获取游戏英雄列表接口文档
func ApiDocGameHeroes(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取游戏英雄列表接口",
		"url": template.HTML(`/api/game/heroes?gamdId=10001&timeoutSeconds=30
			<br/><br/>gamdId：游戏ID，可通过 <a href="/docs/games" target="_blank">游戏列表接口</a> 获取，必填参数。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewListResult(getGameHeroes())),
	})
}

//注册接口文档
func ApiDocRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "用户注册接口",
		"url": template.HTML(`/api/user/register?captchaCode=xxxxxx&timeoutSeconds=30
			<br/><br/>captchaCode：短信验证码，可通过 <a href="/docs/captcha/sms" target="_blank">短信验证码接口</a> 获取，可选参数。<span style='color:red;'>如果传递有效的手机号，则必须传入该参数。</span>
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "POST",
		"param":  CcJson.GetString(getUser(false, true)),
		"result": CcJson.GetString(NewRegisterResult("xxxxxx")),
	})
}

//登录接口文档
func ApiDocLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "用户登录接口",
		"url": template.HTML(`/api/user/login?openid=xxxxxx&timeoutSeconds=30 或者 /api/user/login?account=xxxxxx&password=xxxxxx&timeoutSeconds=30
			<br/><br/>openid：微信OpenID，可选参数。<span style='color:red;'>如果没有传递account和password参数，则必须传递openid参数。</span>
			<br/>account：账号，可选参数。<span style='color:red;'>如果没有传递openid参数，则必须传递account和password参数。</span>
			<br/>password：密码，可选参数。<span style='color:red;'>如果没有传递openid参数，则必须传递account和password参数。</span>
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewUserResult(getUser(true, true))),
	})
}

//更新用户接口文档
func ApiDocUpdateUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "更新用户信息接口",
		"url": template.HTML(`/api/user/update?captchaCode=xxxxxx&timeoutSeconds=30
			<br/><br/>captchaCode：短信验证码，可通过 <a href="/docs/captcha/sms" target="_blank">短信验证码接口</a> 获取，可选参数。<span style='color:red;'>如果手机号有变更，则必须传入该参数。</span>
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "POST",
		"param":  CcJson.GetString(getUser(false, true)),
		"result": CcJson.GetString(NewResult(status.Success)),
	})
}

//获取用户详情接口文档
func ApiDocGetUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取用户详情接口",
		"url": template.HTML(`/api/user/details?userGuid=xxxxxx&timeoutSeconds=30
			<br/><br/>userGuid：用户GUID，必填参数。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewUserResult(getUser(true, true))),
	})
}

//获取用户列表接口文档
func ApiDocGetUsers(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api_doc.tmpl", gin.H{
		"title": "获取用户列表接口",
		"url": template.HTML(`/api/users?pageIndex=1&pageSize=10&timeoutSeconds=30
			<br/><br/>pageIndex：当前页索引号，可选参数。
			<br/>pageSize：每页显示的用户数，可选参数，默认为10个。
			<br/>timeoutSeconds：接口超时秒数，可选参数，默认为30秒超时。`),
		"method": "GET",
		"param":  "参见URL说明",
		"result": CcJson.GetString(NewListResult(getUsers())),
	})
}

func getGames() []*GameModel {
	games := make([]*GameModel, 0)

	games = append(games, &GameModel{Id: 10001, Name: "游戏1"})
	games = append(games, &GameModel{Id: 10002, Name: "游戏2"})

	return games
}

func getGameAreas() []*AreaModel {
	areas := make([]*AreaModel, 0)

	areas = append(areas, &AreaModel{Id: 20001, GameId: 10001, Name: "游戏大区1"})
	areas = append(areas, &AreaModel{Id: 20002, GameId: 10001, Name: "游戏大区2"})

	return areas
}

func getGameServers() []*ServerModel {
	servers := make([]*ServerModel, 0)

	servers = append(servers, &ServerModel{Id: 30001, AreaId: 20001, Name: "游戏服务器1"})
	servers = append(servers, &ServerModel{Id: 30002, AreaId: 20001, Name: "游戏服务器2"})

	return servers
}

func getGameDans() []*DanModel {
	dans := make([]*DanModel, 0)

	dans = append(dans, &DanModel{Id: 20001, GameId: 10001, Name: "段位1"})
	dans = append(dans, &DanModel{Id: 20002, GameId: 10001, Name: "段位2"})

	return dans
}

func getGameHeroes() []*HeroModel {
	heroes := make([]*HeroModel, 0)

	heroes = append(heroes, &HeroModel{Id: 20001, GameId: 10001, Name: "英雄1"})
	heroes = append(heroes, &HeroModel{Id: 20002, GameId: 10001, Name: "英雄2"})

	return heroes
}

func getUsers() []*UserModel {
	users := make([]*UserModel, 0)

	users = append(users, getUser(true, false))
	users = append(users, getUser(true, false))

	return users
}

func getUser(isResult bool, details bool) *UserModel {
	user := &UserModel{
		Account:   "xxxxxx",
		NickName:  "xxxxxx",
		Avator:    "http://www.url/image",
		OpenID:    "xxxxxx",
		Signature: "xxxxxx",
	}

	if details {
		user.Cellphone = "xxxxxx"
		user.Email = "xxxxxx@mail.com"
		user.Wechat = "xxxxxx"
		user.QQ = "xxxxxx"
	}

	if isResult {
		user.Guid = "xxxxxx"
		user.CreatedAt = utils.FormatDateTime(time.Now())
		user.UpdatedAt = utils.FormatDateTime(time.Now())
		user.LastLogin = utils.FormatDateTime(time.Now())
	} else {
		user.Password = "xxxxxx"
	}

	return user
}
