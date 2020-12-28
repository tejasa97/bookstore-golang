package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/tejasa97/bookstore-golang/items/clients/elasticsearch"
	"github.com/tejasa97/bookstore-golang/items/domain/queries"
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

func (i *Item) Search(query queries.EsQuery) ([]Item, *rest_errors.RestErr) {
	finalQuery := query.Build()
	result, err := elasticsearch.Client.Search(indexItems, finalQuery)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()

		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}

		item.ID = hit.Id
		items[index] = item
	}

	return items, nil
}

func (i *Item) Delete() *rest_errors.RestErr {
	_, err := elasticsearch.Client.Delete(indexItems, i.ID)
	if err != nil {
		return rest_errors.NewNotFoundError(fmt.Sprintf("error when trying to find item with id %s for deletion", i.ID))
	}

	return nil
}

func (i *Item) Update(id string) *rest_errors.RestErr {
	_, err := elasticsearch.Client.Update(indexItems, id, i)
	if err != nil {
		return rest_errors.NewNotFoundError(fmt.Sprintf("error when trying to find item with id %s for updation", id))
	}

	return nil
}
