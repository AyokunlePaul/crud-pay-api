package search

import (
	"context"
	"fmt"
	crudPayError "github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
)

var (
	containerName        = os.Getenv("ELASTIC_SEARCH_CONTAINER_NAME")
	client               *elastic.Client
	elasticsearchUri     = fmt.Sprintf("http://%s:9200", containerName)
	elasticsearchContext = context.Background()
)

func Init() {
	var clientError error
	client, clientError = elastic.NewClient(
		elastic.SetURL(elasticsearchUri),
		elastic.SetSniff(false),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
			"X-Caller-Id": []string{"..."},
		}),
	)
	if clientError != nil {
		panic(clientError)
	}
}

type elasticSearchRepository struct {
	errorService crudPayError.Service
}

func NewDatabaseRepository(errorService crudPayError.Service) Repository {
	return &elasticSearchRepository{
		errorService: errorService,
	}
}

func (repository *elasticSearchRepository) Search(parameter Param) (interface{}, *response.BaseResponse) {
	queryTerm := elastic.NewMatchQuery(parameter.Name, parameter.Query)
	searchResults, searchError := client.Search().
		Index(parameter.Index).
		Query(queryTerm).
		Pretty(true).
		Do(elasticsearchContext)
	if searchError != nil {
		return nil, repository.errorService.HandleElasticSearchError(searchError)
	}
	return searchResults, nil
}
