package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lequocduyquang/user-profile-service/db"
	"github.com/lequocduyquang/user-profile-service/domain/users"
)

var (
	router = gin.Default()
)

// StartApp Bootstraping user profile service
func StartApp() {
	db.ConnectToDB()
	db.Client.AutoMigrate(&users.User{})
	mapUrls()
	log.Fatal(router.Run(os.Getenv("PORT")))
}
