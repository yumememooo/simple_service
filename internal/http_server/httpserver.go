package http_server

import (
	"context"
	"fmt"
	"net/http"
	"simple_service/internal/config"
	"simple_service/internal/log"
	"time"

	_ "simple_service/docs" //important swag docs

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log/level"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartHttpServer(errChan chan error) {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(ginLogHandler(), gin.Recovery())
	engine.Use(Cors())
	timeout := time.Duration(config.Configuration.Service.Timeout) * time.Millisecond
	engine.Use(timeoutHandler(timeout))

	engine.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	loadRoutes(engine)

	endpoint := fmt.Sprintf(":%d", config.Configuration.Service.Port)
	go func() { errChan <- engine.Run(endpoint) }()

	log.SugarLogger.Infof("Listening on port: %d", config.Configuration.Service.Port)
}

func ginLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		log.SugarLogger.Infof("level:%s,clientIP:%s,msg:%s", level.DebugValue(),
			clientIP, fmt.Sprintf("Request in %s %s", method, path))

		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		log.SugarLogger.Infof("level:%s,clientIP:%s,latency:%s,statusCode:%d,msg:%s", level.DebugValue(),
			clientIP, fmt.Sprintf("%v", latency), statusCode,
			fmt.Sprintf("%s %s", method, path))
	}
}

func timeoutHandler(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //請求頭部
		if origin != "" {
			//接收客戶端傳送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//伺服器支援的所有跨域請求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允許跨域設定可以返回其他子段，可以自定義欄位
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允許瀏覽器（客戶端）可以解析的頭部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//設定快取時間
			c.Header("Access-Control-Max-Age", "172800")
			//允許客戶端傳遞校驗資訊比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允許型別校驗
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.SugarLogger.Errorf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
