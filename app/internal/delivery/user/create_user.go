package delivery

import (
	"Latihan_1/app/domain/dto"
	"Latihan_1/app/domain/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uh userHandler) CreateUser(c *gin.Context) {
	var dto dto.User
	err := c.ShouldBind(&dto)
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(400, errorMessage)
		return
	}

	res, err := uh.userUsecase.CreateUser(entity.User{
		Name:  dto.Name,
		Age:   dto.Age,
		Email: dto.Email,
	})
	if err != nil {
		errorMessage := gin.H{
			"message": err.Error(),
		}
		c.JSON(400, errorMessage)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
