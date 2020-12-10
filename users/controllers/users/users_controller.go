package users

import (
	"net/http"
	"strconv"

	"github.com/tejasa97/bookstore-golang/users/services"
	"github.com/tejasa97/bookstore-golang/users/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/tejasa97/bookstore-golang/users/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("invalid user id format")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		err := errors.NewBadRequest("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func LoginUser(c *gin.Context) {
	var user_req users.LoginRequest
	if err := c.ShouldBindJSON(&user_req); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := services.UsersService.LoginUser(user_req)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	var user users.User

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("invalid user id format")
		c.JSON(err.Status, err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// To update only certain fields
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.UpdateUser(isPartial, userID, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequest("invalid user id format")
		c.JSON(err.Status, err)
		return
	}

	if deleteErr := services.UsersService.DeleteUser(userID); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func SearchUser(c *gin.Context) {

	status := c.Query("status")

	users, err := services.UsersService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
