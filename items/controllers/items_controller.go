package controllers

import (
	"net/http"

	"github.com/tejasa97/bookstore-golang/items/domain/queries"

	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore-golang/items/domain/items"
	"github.com/tejasa97/bookstore-golang/items/services"
	"github.com/tejasa97/utils-go/rest_errors"

	"github.com/tejasa97/oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Search(*gin.Context)
	Delete(*gin.Context)
}

type itemsController struct{}

func (i *itemsController) Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	var itemRequest items.Item
	if err := c.ShouldBindJSON(&itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// Assign `seller` to the caller of the request
	itemRequest.Seller = oauth.GetClientId(c.Request)

	item, err := services.ItemsService.Create(itemRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (i *itemsController) Get(c *gin.Context) {
	itemID := c.Param("item_id")

	item, err := services.ItemsService.Get(itemID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, item)

}

func (i *itemsController) Search(c *gin.Context) {

	var eq queries.EsQuery
	if err := c.ShouldBindJSON(&eq); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	items, err := services.ItemsService.Search(eq)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (i *itemsController) Delete(c *gin.Context) {
	itemID := c.Param("item_id")

	err := services.ItemsService.Delete(itemID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.Status(http.StatusNoContent)
}
