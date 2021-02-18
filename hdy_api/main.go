package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redochen/demos/hdy_api/config"
	"github.com/redochen/demos/hdy_api/docs"
	svc "github.com/redochen/demos/hdy_api/services"
	"github.com/redochen/tools/log"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(keepAlive())
	r.Use(ginLogger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.BestCompression))

	r.LoadHTMLGlob("templates/*")

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	//文档接口
	r.GET("/docs", docs.APIDocs)
	r.GET("/docs/captcha", docs.APIDocCaptcha)
	r.GET("/docs/captcha/sms", docs.APIDocSmsCaptcha)
	r.GET("/docs/captcha/verify", docs.APIDocVerifyCaptcha)
	r.GET("/docs/captcha/resource", docs.APIDocCaptchaResource)
	r.GET("/docs/captcha/sample", docs.APIDocCaptchaSample)
	r.GET("/docs/games", docs.APIDocGames)
	r.GET("/docs/game/areas", docs.APIDocGameAreas)
	r.GET("/docs/game/servers", docs.APIDocGameServers)
	r.GET("/docs/game/dans", docs.APIDocGameDans)
	r.GET("/docs/game/heroes", docs.APIDocGameHeroes)
	r.GET("/docs/user/register", docs.APIDocRegister)
	r.GET("/docs/user/login", docs.APIDocLogin)
	r.GET("/docs/user/update", docs.APIDocUpdateUser)
	r.GET("/docs/user/details", docs.APIDocGetUser)
	r.GET("/docs/users", docs.APIDocGetUsers)

	//游戏接口
	r.GET("/api/games", svc.GamesAsync)
	r.GET("/api/game/areas", svc.GameAreasAsync)
	r.GET("/api/game/servers", svc.GameServersAsync)
	r.GET("/api/game/dans", svc.GameDansAsync)
	r.GET("/api/game/heroes", svc.GameHeroesAsync)

	//用户接口
	r.POST("/api/user/register", svc.RegisterAsync)
	r.GET("/api/user/login", svc.LoginAsync)
	r.POST("/api/user/update", svc.UpdateUserAsync)
	r.GET("/api/user/details", svc.GetUserAsync)
	r.GET("/api/users", svc.GetUsersAsync)

	//验证码接口
	r.GET("/api/captcha", svc.GetCaptcha)
	r.GET("/api/captcha/sms", svc.GetSmsCaptchaAsync)
	r.POST("/api/captcha/verify", svc.VerifyCaptcha)
	r.GET("/captcha/:file", svc.LoadCaptcha)

	//启动服务，请在此之前添加新的路由映射
	r.Run(":" + strconv.Itoa(config.ServerPort))
}

// 保持活动连接
func keepAlive() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !strings.Contains(ctx.Request.Header.Get("Connection"), "Keep-Alive") {
			return
		}
		ctx.Writer.Header().Set("Connection", "Keep-Alive")
		ctx.Next()
	}
}

//日志记录器
func ginLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		// before request
		addr := ctx.Request.Header.Get("X-Real-IP")
		if addr == "" {
			addr = ctx.Request.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = ctx.Request.RemoteAddr
			}
		}

		msg := fmt.Sprintf("Started %s %s for %s at %v", ctx.Request.Method,
			ctx.Request.URL.Path, addr, start.Local())
		ctx.Next()

		// after request
		latency := time.Since(start)
		//log.Print(latency)

		// access the status we are sending
		log.Infof("%s; Completed %v %s in %v.",
			msg, ctx.Writer.Status(), http.StatusText(ctx.Writer.Status()), latency)
	}
}
