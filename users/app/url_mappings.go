package app

import (
	"github.com/tejasa97/bookstore-golang/users/controllers/ping"
	"github.com/tejasa97/bookstore-golang/users/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/internal/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.POST("/users/login", users.LoginUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)

}
