package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/redochen/demos/mnbz-api/config"
	svc "github.com/redochen/demos/mnbz-api/services"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redochen/tools/log"

	_ "github.com/redochen/demos/mnbz-api/docs" //其中docs是你生成docs的路径
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 码农搬砖接口
// @description 码农搬砖的接口系统
// @host localhost:1111
func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(keepAlive())
	r.Use(ginLogger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.BestCompression))
	// 允许使用跨域请求	 全局中间件
	r.Use(Cors())

	r.LoadHTMLGlob("templates/*")

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})

	//接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.GET("/docs", docs.APIDocs)
	// r.GET("/docs/captcha", docs.APIDocCaptcha)
	// r.GET("/docs/captcha/sms", docs.APIDocSmsCaptcha)
	// r.GET("/docs/captcha/verify", docs.APIDocVerifyCaptcha)
	// r.GET("/docs/captcha/resource", docs.APIDocCaptchaResource)
	// r.GET("/docs/captcha/sample", docs.APIDocCaptchaSample)

	//用户接口
	r.POST("/api/user/register", svc.RegisterAsync)
	r.POST("/api/user/login", svc.LoginAsync)
	r.POST("/api/user/update", svc.UpdateUserAsync)
	r.GET("/api/user/details", svc.GetUserAsync)
	//r.GET("/api/users", svc.GetUsersAsync)
	r.POST("/api/user/logout", svc.LogoutAsync)

	//配置接口
	r.POST("/api/config/add", svc.AddConfigAsync)
	r.POST("/api/config/update", svc.UpdateConfigAsync)
	r.POST("/api/config/delete", svc.DeleteConfigAsync)
	r.GET("/api/config/details", svc.GetConfigAsync)
	r.GET("/api/configs", svc.GetConfigsAsync)

	//需求接口
	r.POST("/api/need/add", svc.AddNeedAsync)
	r.POST("/api/need/update", svc.UpdateNeedAsync)
	r.POST("/api/need/delete", svc.DeleteNeedAsync)
	r.GET("/api/need/details", svc.GetNeedAsync)
	r.GET("/api/needs", svc.GetNeedsAsync)

	//资源接口
	r.POST("/api/resource/add", svc.AddResourceAsync)
	r.POST("/api/resource/update", svc.UpdateResourceAsync)
	r.POST("/api/resource/delete", svc.DeleteResourceAsync)
	r.GET("/api/resource/details", svc.GetResourceAsync)
	r.GET("/api/resources", svc.GetResourcesAsync)

	//抢票接口
	r.POST("/api/task/add", svc.AddTaskAsync)
	r.POST("/api/task/update", svc.UpdateTaskAsync)
	r.POST("/api/task/delete", svc.DeleteTaskAsync)
	r.GET("/api/task/details", svc.GetTaskAsync)
	r.GET("/api/tasks", svc.GetTasksAsync)

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

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			// header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, sid,Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //	处理请求
	}
}

// 日志记录器
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
