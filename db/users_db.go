package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	"github.com/joho/godotenv"
)

var (
	// Client instance be exported
	Client *gorm.DB
	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_HOST")
)

// Init db
func init() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost = "suleiman.db.elephantsql.com"
	dbPort = "5432"
	dbName = "colmcodb"
	dbUser = "colmcodb"
	dbPass = "ctC_fXIsfCq7p8tYvWWtoFzmhZbW0xxb"
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbPort, dbUser, dbName, dbPass)
	Client, err = gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", dbName)
		panic(err)
	} else {
		fmt.Printf("Connected to the database %s successfully\n", dbName)
	}
	// Client.Debug().AutoMigrate(&users.User{})
}
