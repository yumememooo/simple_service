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
