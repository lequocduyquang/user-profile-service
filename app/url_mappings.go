package app

import "github.com/lequocduyquang/user-profile-service/controllers"

func mapUrls() {
	router.GET("/ping", controllers.PingController.Ping)
}
