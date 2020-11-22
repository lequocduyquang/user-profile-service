package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lequocduyquang/user-profile-service/services"
)

var (
	// UserController exported
	UserController UserControllerInterface = &userController{}
)

// UserControllerInterface interface
type UserControllerInterface interface {
	GetAll(c *gin.Context)
}

type userController struct{}

func (u *userController) GetAll(c *gin.Context) {
	result, err := services.UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}
