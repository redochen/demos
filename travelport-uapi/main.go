package main

import (
	"fmt"
	. "github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/redochen/demos/travelport-uapi/config"
	svc "github.com/redochen/demos/travelport-uapi/services"
	. "github.com/redochen/tools/log"
	. "github.com/redochen/tools/string"
	"net/http"
	"runtime"
	"strings"
	"time"
)

//入口方法
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	//设置模式
	//gin.SetMode(gin.ReleaseMode)

	//初始化
	r := gin.New()
	r.Use(keepAlive())
	//r.Use(ginLogger())
	r.Use(Gzip(BestCompression))

	//路由映射
	r.GET("/Hello/:name", hello)
	r.GET("/Ping", svc.Ping)

	r.POST("/LowFare", svc.LowFare)
	r.POST("/AirAvail", svc.AirAvail)
	r.POST("/AirPrice", svc.AirPrice)

	r.POST("/HotelAvail", svc.HotelAvail)
	r.POST("/HotelMedia", svc.HotelMedia)
	r.POST("/HotelAvailEx", svc.HotelAvailEx)
	r.POST("/HotelDetails", svc.HotelDetails)
	r.POST("/HotelRules", svc.HotelRules)

	r.POST("/CreatePnr", svc.CreatePnr)
	r.POST("/RetrievePnr", svc.RetrievePnr)
	r.POST("/QuotePnr", svc.QuotePnr)
	r.POST("/CancelPnr", svc.CancelPnr)

	//启动服务
	r.Run(fmt.Sprintf(":%s", config.Port))
}

// 保持活动连接
func keepAlive() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.Header.Get("Connection"), "Keep-Alive") {
			return
		}
		c.Writer.Header().Set("Connection", "Keep-Alive")
		c.Next()
	}
}

//日志记录器
func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// before request
		addr := c.Request.Header.Get("X-Real-IP")
		if addr == "" {
			addr = c.Request.Header.Get("X-Forwarded-For")
			if addr == "" {
				addr = c.Request.RemoteAddr
			}
		}

		msg := fmt.Sprintf("Started %s %s for %s at %v", c.Request.Method,
			c.Request.URL.Path, addr, start.Local())
		c.Next()

		// after request
		latency := time.Since(start)
		//log.Print(latency)

		// access the status we are sending
		Logger.Infof("%s; Completed %v %s in %v.",
			msg, c.Writer.Status(), http.StatusText(c.Writer.Status()), latency)
	}
}

//测试接口
func hello(c *gin.Context) {
	name := c.Params.ByName("name")
	now := CcStr.FormatTime(time.Now(), "yyyy-MM-ddTHH:mm:ss.fffzzz")
	message := "Hello " + name + ", at " + now

	/*
		start := time.Now()

		//带超时等待
		select {
		case <-time.After(1 * time.Second):
			dur := time.Since(start)
			fmt.Printf("spent %f seconds.\n", dur.Seconds())
			break
		}
	*/

	//ts := time.Now().Local().Format("2006-01-02 15:04:05")
	//message += " at " + RemoveDateSeparator(ts)

	c.String(http.StatusOK, message)
}
