package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore-golang/oauth/http"
	"github.com/tejasa97/bookstore-golang/oauth/repository/db"
	"github.com/tejasa97/bookstore-golang/oauth/repository/rest"
	"github.com/tejasa97/bookstore-golang/oauth/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(rest.NewRepository(), db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8001")
}
