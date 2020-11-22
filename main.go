package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lequocduyquang/user-profile-service/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}
	app.StartApp()
}
