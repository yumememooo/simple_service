package apis

import (
	"net/http"
	"simple_service/internal/service"

	"github.com/gin-gonic/gin"
)

// @Tags pet
// @Summary Get Pet by animal_kind
// @Success 200 {array} model.Pet
// @Param animal_kind query string false "search animal_kind:{貓/狗}"
// @Router /api/v1/pet [get]
func FindPet(c *gin.Context) {
	animal_kind := c.Query("animal_kind")
	if res, err := service.FindPet(animal_kind); err != nil {
		BadRequest(c, err)
	} else {
		// c.Header(internal.CustomHeaderForTotal, totalCount)
		c.JSON(http.StatusOK, res)
	}
}
