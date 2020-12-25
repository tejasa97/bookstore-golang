package services

import (
	"net/http"

	"github.com/tejasa97/bookstore-golang/items/domain/items"
	"github.com/tejasa97/utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(name string) (*items.Item, *rest_errors.RestErr) {
	return nil, &rest_errors.RestErr{
		Status:  http.StatusNotImplemented,
		Message: "Not implemented yet",
		Error:   "not_implemented",
	}
}
