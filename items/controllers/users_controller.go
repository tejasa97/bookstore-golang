package controllers

import (
	"github.com/gin-gonic/gin"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(*gin.Context)
}

type pingController struct{}

func (i *pingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}
