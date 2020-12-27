package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tejasa97/bookstore-golang/items/clients/elasticsearch"
	"github.com/tejasa97/utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	i.ID = result.Id
	return nil
}

func (i *Item) Get() *rest_errors.RestErr {
	itemID := i.ID
	result, err := elasticsearch.Client.Get(indexItems, i.ID)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.ID))
		}

		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get item with id %s", i.ID), errors.New("database error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error trying to parse database response", errors.New("database error"))
	}

	i.ID = itemID
	return nil
}
