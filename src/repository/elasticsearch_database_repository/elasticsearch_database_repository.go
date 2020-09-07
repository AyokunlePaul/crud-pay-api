package elasticsearch_database_repository

import (
	"github.com/AyokunlePaul/crud-pay-api/src/clients/elasticsearch_client"
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product"
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product/product_search"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities"
	"reflect"
)

type productSearchRepository struct{}

func New() product_search.Repository {
	return &productSearchRepository{}
}

func (repository *productSearchRepository) Search(query string) ([]product.Product, *response.BaseResponse) {
	var products []product.Product
	var productType product.Product

	searchResult, searchError := elasticsearch_client.Client.Search(query)
	if searchError != nil {
		return nil, utilities.HandleElasticSearchError(searchError)
	}

	for _, currentProduct := range searchResult.Each(reflect.TypeOf(productType)) {
		products = append(products, currentProduct.(product.Product))
	}

	return products, nil
}
