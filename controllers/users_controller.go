package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lequocduyquang/user-profile-service/domain/users"
	"github.com/lequocduyquang/user-profile-service/services"
	"github.com/lequocduyquang/user-profile-service/utils"
)

var (
	// UserController exported
	UserController UserControllerInterface = &userController{}
)

// UserControllerInterface interface
type UserControllerInterface interface {
	GetAll(c *gin.Context)
	Register(c *gin.Context)
}

type userController struct{}

func (u *userController) GetAll(c *gin.Context) {
	result, err := services.UserService.GetAll()
	if err != nil {
		restErr := utils.NewNotFoundError("not found data")
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *userController) Register(c *gin.Context) {
	var user users.User
	fmt.Printf("Value of user %v", &user)
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, savedErr := services.UserService.Create(user)
	if savedErr != nil {
		c.JSON(http.StatusBadRequest, savedErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
