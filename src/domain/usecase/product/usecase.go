package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/search"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type productUseCase struct {
	productManager product.Manager
	tokenManager   token.Manager
	searchManager  search.Manager
}

func NewUseCase(manager product.Manager, tokenManager token.Manager, searchManager search.Manager) UseCase {
	return &productUseCase{
		productManager: manager,
		tokenManager:   tokenManager,
		searchManager:  searchManager,
	}
}

func (useCase *productUseCase) CreateProduct(token string, product *product.Product) *response.BaseResponse {
	ownerId, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return ownerIdError
	}
	product.OwnerId, _ = entity.StringToCrudPayId(ownerId)
	return useCase.productManager.Create(product)
}

func (useCase *productUseCase) UpdateProduct(token string, product *product.Product) *response.BaseResponse {
	_, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return ownerIdError
	}
	return useCase.productManager.Update(product)
}

func (useCase *productUseCase) SearchProduct(token string, query string) ([]product.Product, *response.BaseResponse) {
	_, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return nil, ownerIdError
	}
	searchParameter := search.Param{
		Index: "crudpay.products",
		Query: query,
		Name:  "product_name",
	}
	results, searchError := useCase.searchManager.Search(searchParameter)
	if searchError != nil {
		return nil, searchError
	}

	return results.([]product.Product), nil
}

func (useCase *productUseCase) GetProductWithId(token string, productId string) (*product.Product, *response.BaseResponse) {
	_, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return nil, ownerIdError
	}
	id, _ := entity.StringToCrudPayId(productId)
	return useCase.productManager.Get(id)
}

func (useCase *productUseCase) GetAllCreatedProducts(token string) ([]product.Product, *response.BaseResponse) {
	ownerId, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return nil, ownerIdError
	}
	id, _ := entity.StringToCrudPayId(ownerId)
	return useCase.productManager.GetProducts(id)
}