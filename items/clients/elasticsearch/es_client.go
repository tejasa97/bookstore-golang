package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/tejasa97/utils-go/logger"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("https://99a6772afe4c438690f2c1d6a8f30f41.asia-south1.gcp.elastic-cloud.com:9243"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth("elastic", "jJrQNm2AW11VN6ua7Locwazr"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, item interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(item).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}

func (c *esClient) Get(index string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get document with id %s in index %s", id, index), err)
		return nil, err
	}

	return result, nil
}
