package handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	productUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/product"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/string_utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Product interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetVendorProduct(*gin.Context)
	Update(*gin.Context)
	Search(*gin.Context)
}

type productHandler struct {
	useCase productUseCase.UseCase
}

func ForProduct(useCase productUseCase.UseCase) Product {
	return &productHandler{
		useCase: useCase,
	}
}

func (handler *productHandler) Create(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]

	newProduct := product.New()
	_ = context.BindJSON(&newProduct)

	createError := handler.useCase.CreateProduct(token, newProduct)
	if createError != nil {
		context.JSON(createError.Status, createError)
		return
	}

	context.JSON(http.StatusCreated, response.NewCreatedResponse("product created", newProduct))
}

func (handler *productHandler) Get(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	productId := context.Param("product_id")

	if string_utilities.IsEmpty(productId) {
		result, productsError := handler.useCase.GetAllMyProducts(token)
		if productsError != nil {
			context.JSON(productsError.Status, productsError)
			return
		}
		context.JSON(http.StatusOK, response.NewOkResponse("products fetched", result))
		return
	}

	result, productError := handler.useCase.GetProductWithId(token, productId)
	if productError != nil {
		context.JSON(productError.Status, productError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("product fetched", result))
}

func (handler *productHandler) GetVendorProduct(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	ownerId := context.Param("owner_id")

	products, getProductsError := handler.useCase.GetAllProductsCreatedByUserWithId(token, ownerId)
	if getProductsError != nil {
		context.JSON(getProductsError.Status, getProductsError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("products fetched", products))
}

func (handler *productHandler) Update(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	productId := context.Param("product_id")

	newProduct := new(product.Product)
	_ = context.BindJSON(&newProduct)

	updatedProduct, updateError := handler.useCase.UpdateProduct(token, productId, newProduct)
	if updateError != nil {
		context.JSON(updateError.Status, updateError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("product successfully updated", updatedProduct))
}

func (handler *productHandler) Search(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	query := context.Query("name")
	if string_utilities.IsEmpty(query) {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("query cannot be empty"))
		return
	}
	result, searchError := handler.useCase.SearchProduct(token, query)
	if searchError != nil {
		context.JSON(searchError.Status, searchError)
		return
	}

	context.JSON(http.StatusOK, response.NewOkResponse("products fetched", result))
}
