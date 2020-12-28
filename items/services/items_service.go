package services

import (
	"github.com/tejasa97/bookstore-golang/items/domain/items"
	"github.com/tejasa97/bookstore-golang/items/domain/queries"
	"github.com/tejasa97/utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, *rest_errors.RestErr)
	Delete(string) *rest_errors.RestErr
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {

	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {

	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, *rest_errors.RestErr) {

	itemsDao := items.Item{}
	items, err := itemsDao.Search(query)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *itemsService) Delete(id string) *rest_errors.RestErr {

	itemsDao := items.Item{ID: id}
	err := itemsDao.Delete()
	if err != nil {
		return err
	}

	return nil
}
