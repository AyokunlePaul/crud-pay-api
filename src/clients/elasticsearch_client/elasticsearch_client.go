package elasticsearch_client

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
)

var (
	Client    elasticsearchClientInterface = &elasticsearchClient{}
	container string                       = os.Getenv(elasticsearchContainerName)
)

const elasticsearchContainerName = "ELASTIC_SEARCH_CONTAINER_NAME"

type elasticsearchClient struct {
	client *elastic.Client
}

type elasticsearchClientInterface interface {
	Search(string) (*elastic.SearchResult, error)
	setClient(*elastic.Client)
}

func Init() {
	url := fmt.Sprintf("http://%s:9200", container)

	client, elasticsearchClientError := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)

	if elasticsearchClientError != nil {
		panic(elasticsearchClientError)
	}
	Client.setClient(client)
}

func (client *elasticsearchClient) Search(query string) (*elastic.SearchResult, error) {
	elasticsearchContext := context.Background()
	queryTerm := elastic.NewMatchQuery("product_name", query)
	return client.client.Search().
		Index("crudpay.products").
		Query(queryTerm).
		Pretty(true).
		Do(elasticsearchContext)
}

func (client *elasticsearchClient) setClient(elasticClient *elastic.Client) {
	client.client = elasticClient
}
