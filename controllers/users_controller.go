package controllers

import (
	"net/http"
	"strconv"

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
	GetByID(c *gin.Context)
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
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := utils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, savedErr := services.UserService.Create(user)
	if savedErr != nil {
		restErr := utils.NewBadRequestError(savedErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (u *userController) GetByID(c *gin.Context) {
	var (
		err    error
		userID int64
	)
	if userID, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		restErr := utils.NewBadRequestError("User id is invalid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, foundErr := services.UserService.GetByID(userID)
	if foundErr != nil {
		restErr := utils.NewBadRequestError(foundErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
