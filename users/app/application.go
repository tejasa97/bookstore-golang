package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()
	router.Run(":8000")
}
