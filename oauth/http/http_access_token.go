package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	atDomain "github.com/tejasa97/bookstore-golang/oauth/domain/access_token"
	"github.com/tejasa97/bookstore-golang/oauth/services/access_token"
	"github.com/tejasa97/bookstore-golang/oauth/utils/errors"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at_request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&at_request); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(at_request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}
