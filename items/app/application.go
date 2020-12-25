package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore-golang/items/clients/elasticsearch"
)

var (
	router = gin.Default()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	router.Run(":8000")
}
