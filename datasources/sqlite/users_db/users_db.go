package users_db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Client *gorm.DB
)

func init() {

	const DBName = "tejas.db"

	var err error
	Client, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database %s \n", DBName))
	}
	fmt.Printf("Connection successfully opened to database %s \n", DBName)
}
