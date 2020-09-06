package product_handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/models/product_payload"
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product/product_service"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities/string_utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Handler interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Search(*gin.Context)
}

type handler struct {
	service product_service.Service
}

func New(service product_service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (handler *handler) Create(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	var payload product_payload.ProductPayload
	_ = context.BindJSON(&payload.Payload)

	result, createError := handler.service.Create(payload.ToDomain(), token)
	if createError != nil {
		context.JSON(createError.Status, createError)
		return
	}

	context.JSON(http.StatusCreated, response.NewCreatedResponse("product created", result))
}

func (handler *handler) Get(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	productId := context.Param("product_id")

	if string_utilities.IsEmpty(productId) {
		result, productsError := handler.service.GetProducts(token)
		if productsError != nil {
			context.JSON(productsError.Status, productsError)
			return
		}
		context.JSON(http.StatusOK, response.NewOkResponse("products fetched", result))
		return
	}

	result, productError := handler.service.Get(productId, token)
	if productError != nil {
		context.JSON(productError.Status, productError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("product fetched", result))
}

func (handler *handler) Search(context *gin.Context) {
	panic("implement me")
}
