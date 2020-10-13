package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DBName = "tejas.db"
)

var (
	Client *sql.DB
)

func init() {

	var err error
	Client, err = sql.Open("sqlite3", "./tejas.db")
	fmt.Println("in DB", Client)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Printf("%s DB successfully configured", DBName)
}
