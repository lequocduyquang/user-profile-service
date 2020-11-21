package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApp Bootstraping user profile service
func StartApp() {
	mapUrls()
	router.Run(":5000")
}
