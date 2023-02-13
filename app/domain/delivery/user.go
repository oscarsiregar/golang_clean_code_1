package delivery

import "github.com/gin-gonic/gin"

type UserDeliveryInterface interface {
	CreateUser(c *gin.Context)
}
