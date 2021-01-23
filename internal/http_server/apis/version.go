package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags version
// @Summary Get version
// @Success 200 {object} object "success"
// @Router /api/v1/version [get]
func Version(c *gin.Context) {

	c.JSON(http.StatusOK, "1.1.0")

}
