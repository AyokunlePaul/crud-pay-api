package search

import (
	"encoding/json"
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
)

type manager struct {
	repository Repository
}

func NewManager(repository Repository) Manager {
	return &manager{
		repository: repository,
	}
}

func (manager *manager) search(parameter Param) (interface{}, *response.BaseResponse) {
	searchResults, searchError := manager.repository.search(parameter)
	if searchError != nil {
		return nil, searchError
	}
	return searchResults, nil
}

func (manager *manager) SearchProduct(parameter Param) ([]product.Product, *response.BaseResponse) {
	searchResults, searchError := manager.repository.search(parameter)
	if searchError != nil {
		return nil, searchError
	}

	var products []product.Product
	results := searchResults.(*entity.SearchResult)
	if results.TotalHits() > 0 {
		for _, hit := range results.Hits.Hits {
			var currentProduct product.Product
			marshalError := json.Unmarshal(hit.Source, &currentProduct)
			if marshalError != nil {
				errorMessage := fmt.Sprintf("product marshal error: ID: %s", hit.Id)
				logger.Error(errorMessage, marshalError)
				continue
			}
			currentProduct.Id, _ = entity.StringToCrudPayId(hit.Id)
			products = append(products, currentProduct)
		}
	}
	return products, nil
}

func (manager *manager) SearchUser(parameter Param) ([]user.User, *response.BaseResponse) {
	searchResults, searchError := manager.repository.search(parameter)
	if searchError != nil {
		return nil, searchError
	}

	var users []user.User
	results := searchResults.(*entity.SearchResult)
	if results.TotalHits() > 0 {
		for _, hit := range results.Hits.Hits {
			var currentUser user.User
			marshalError := json.Unmarshal(hit.Source, &currentUser)
			if marshalError != nil {
				errorMessage := fmt.Sprintf("user marshal error: ID: %s", hit.Id)
				logger.Error(errorMessage, marshalError)
				continue
			}
			currentUser.Id, _ = entity.StringToCrudPayId(hit.Id)
			users = append(users, currentUser)
		}
	}
	return users, nil
}
