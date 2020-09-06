package elasticsearch_client

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
)

var client *elastic.Client
const elasticsearchContainerName = "ELASTIC_SEARCH_CONTAINER_NAME"

func init() {
	var elasticsearchClientError error
	elasticsearchContainer := os.Getenv(elasticsearchContainerName)

	url := fmt.Sprintf("http://%s:9200", elasticsearchContainer)

	client, elasticsearchClientError = elastic.NewClient(
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
}

func Get() *elastic.Client {
	return client
}
