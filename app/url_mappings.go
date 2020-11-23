package app

import "github.com/lequocduyquang/user-profile-service/controllers"

func mapUrls() {
	router.GET("/ping", controllers.PingController.Ping)

	router.GET("/users", controllers.UserController.GetAll)
	router.GET("/users/:id", controllers.UserController.GetByID)
	router.POST("/users/register", controllers.UserController.Register)
	router.DELETE("/users/:id", controllers.UserController.DeleteByID)
}
