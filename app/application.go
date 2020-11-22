package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lequocduyquang/user-profile-service/db"
	"github.com/lequocduyquang/user-profile-service/domain/users"
)

var (
	router = gin.Default()
)

// StartApp Bootstraping user profile service
func StartApp() {
	db.Client.AutoMigrate(&users.User{})
	mapUrls()
	router.Run(":5000")
}
