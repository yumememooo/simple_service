package http_server

import (
	"simple_service/internal/http_server/apis"

	"github.com/gin-gonic/gin"
)

func loadRoutes(e *gin.Engine) {
	root := e.Group("api/v1")

	myGroup := root.Group("")
	{
		myGroup.GET("version", apis.Version)
		myGroup.GET("pet", apis.FindPet)
		myGroup.POST("pet", apis.AddPet)

	}
}
