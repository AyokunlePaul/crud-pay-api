package product

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/search"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type productUseCase struct {
	productManager product.Manager
	tokenManager   token.Manager
	searchManager  search.Manager
	userManager    user.Manager
}

func New(manager product.Manager, tokenManager token.Manager, searchManager search.Manager, userManager user.Manager) UseCase {
	return &productUseCase{
		productManager: manager,
		tokenManager:   tokenManager,
		searchManager:  searchManager,
		userManager:    userManager,
	}
}

func (useCase *productUseCase) CreateProduct(token string, product *product.Product) *response.BaseResponse {
	ownerId, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return ownerIdError
	}

	if validationError := product.CanBeCreated(); validationError != nil {
		return validationError
	}

	owner := user.New()
	owner.Id, _ = entity.StringToCrudPayId(ownerId)
	if getUserError := useCase.userManager.Get(owner); getUserError != nil {
		return getUserError
	}
	if !owner.IsVendor {
		return response.NewBadRequestError("only vendors can create a product")
	}

	product.OwnerId, _ = entity.StringToCrudPayId(ownerId)
	return useCase.productManager.Create(product)
}

func (useCase *productUseCase) UpdateProduct(token, productId string, product *product.Product) (*product.Product, *response.BaseResponse) {
	userId, userIdError := useCase.tokenManager.Get(token)
	if userIdError != nil {
		return nil, userIdError
	}

	product.Id, _ = entity.StringToCrudPayId(productId)
	oldProduct, getProductError := useCase.productManager.Get(product.Id)
	if oldProduct.OwnerId.Hex() != userId {
		return nil, response.NewBadRequestError("not your product")
	}

	if getProductError != nil {
		return nil, getProductError
	}
	if productUpdateError := oldProduct.CanBeUpdatedWith(product); productUpdateError != nil {
		return nil, productUpdateError
	}

	return oldProduct, useCase.productManager.Update(oldProduct)
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

func (useCase *productUseCase) GetAllProductsCreatedByUserWithId(token string) ([]product.Product, *response.BaseResponse) {
	ownerId, ownerIdError := useCase.tokenManager.Get(token)
	if ownerIdError != nil {
		return nil, ownerIdError
	}
	id, _ := entity.StringToCrudPayId(ownerId)
	return useCase.productManager.List(id)
}
