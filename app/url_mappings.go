package app

import "github.com/lequocduyquang/user-profile-service/controllers"

func mapUrls() {
	router.GET("/ping", controllers.PingController.Ping)

	router.GET("/users", controllers.UserController.GetAll)
	router.POST("/users/register", controllers.UserController.Register)
}
