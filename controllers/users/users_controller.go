package users

import (
	"net/http"
	"strconv"

	"github.com/tejasa97/bookstore_users-api/services"
	"github.com/tejasa97/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore_users-api/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("invalid user id format")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		err := errors.NewBadRequest("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusAccepted, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
