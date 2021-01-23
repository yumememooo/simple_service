package apis

import (
	"fmt"
	"net/http"
	"simple_service/internal/log"
	"strings"

	"github.com/gin-gonic/gin"
)

// General HTTP Response
type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ConflictRequest(c *gin.Context, err error) {
	c.JSON(http.StatusConflict, &HttpResponse{
		Code:    http.StatusConflict,
		Message: err.Error(),
	})
	log.SugarLogger.Errorf("%v", err.Error())
}
func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, &HttpResponse{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	})
	log.SugarLogger.Errorf("%v", err.Error())
}

//handle backend-service api response error message and return status code
func HandleHttpResponse(c *gin.Context, err error) {
	log.SugarLogger.Error(err.Error())
	if strings.Contains(err.Error(), "Timeout exceeded") || strings.Contains(err.Error(), "connected party did not properly respond after a period of time") {
		c.JSON(http.StatusRequestTimeout, &HttpResponse{
			Code:    http.StatusRequestTimeout,
			Message: err.Error(),
		})
		return
	}
	if strings.Contains(err.Error(), " No connection could be made because the target machine actively refused it") {
		c.JSON(http.StatusBadGateway, &HttpResponse{
			Code:    http.StatusBadGateway,
			Message: err.Error(),
		})
		return
	}
	if strings.Contains(err.Error(), "connection refuse") {
		c.JSON(http.StatusBadGateway, &HttpResponse{
			Code:    http.StatusBadGateway,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest, &HttpResponse{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	})
}

func NotFound(c *gin.Context, subject string) {
	c.JSON(http.StatusNotFound, &HttpResponse{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", subject),
	})
}

func ExpectationFailed(c *gin.Context, err error) {
	c.JSON(http.StatusExpectationFailed, &HttpResponse{
		Code:    http.StatusExpectationFailed,
		Message: err.Error(),
	})
	log.SugarLogger.Errorf("%v", err.Error())
}

func LockedFailed(c *gin.Context, err error) {
	c.JSON(http.StatusLocked, &HttpResponse{
		Code:    http.StatusLocked,
		Message: err.Error(),
	})
	log.SugarLogger.Error(err.Error())
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, &HttpResponse{
		Code:    http.StatusOK,
		Message: "success",
	})
}

func SuccessMessage(c *gin.Context, m string) {
	c.JSON(http.StatusOK, &HttpResponse{
		Code:    http.StatusOK,
		Message: m,
	})
}
