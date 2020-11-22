package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
)

var (
	// Client instance be exported
	Client *gorm.DB
)

// ConnectToDB init db
func ConnectToDB() {
	var err error
	Client, err = gorm.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Print("\nCannot connect to database")
		panic(err)
	} else {
		fmt.Print("\nConnected to the database successfully")
	}
}
