package app

import (
	"github.com/tejasa97/bookstore-golang/items/controllers"
)

func mapUrls() {
	// Ping
	router.GET("/ping", controllers.PingController.Ping)

	// TODO: Items
	router.POST("/items", controllers.ItemsController.Create)
	router.GET("/items/:item_id", controllers.ItemsController.Get)
	router.POST("/items/search", controllers.ItemsController.Search)

}
