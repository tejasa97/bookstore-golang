package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()

	logger.Info("starting the app...")
	router.Run(":8000")
}
