package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore-golang/oauth/clients/cassandra"
	"github.com/tejasa97/bookstore-golang/oauth/domain/access_token"
	"github.com/tejasa97/bookstore-golang/oauth/http"
	"github.com/tejasa97/bookstore-golang/oauth/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Ping database or panic
	_, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
