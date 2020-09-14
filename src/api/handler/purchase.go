package handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	purchaseUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Purchase interface {
	Create(*gin.Context)
	Get(*gin.Context)
	List(*gin.Context)
}

type purchaseHandler struct {
	useCase purchaseUseCase.UseCase
}

func ForPurchase(useCase purchaseUseCase.UseCase) Purchase {
	return &purchaseHandler{
		useCase: useCase,
	}
}

func (handler *purchaseHandler) Create(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]

	newPurchase := purchase.New()
	_ = context.BindJSON(&newPurchase)

	if createPurchaseError := handler.useCase.CreatePurchase(token, newPurchase); createPurchaseError != nil {
		context.JSON(createPurchaseError.Status, createPurchaseError)
		return
	} else {
		context.JSON(http.StatusCreated, response.NewCreatedResponse("purchase created", newPurchase))
		return
	}
}

func (handler *purchaseHandler) Get(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	productId := context.Param("product_id")
	if currentPurchase, getPurchaseError := handler.useCase.GetPurchase(token, productId); getPurchaseError != nil {
		context.JSON(getPurchaseError.Status, getPurchaseError)
		return
	} else {
		context.JSON(http.StatusOK, response.NewOkResponse("purchase fetched", currentPurchase))
		return
	}
}

func (handler *purchaseHandler) List(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]
	if purchases, getPurchasesError := handler.useCase.GetAllPurchasesMadeByUser(token); getPurchasesError != nil {
		context.JSON(getPurchasesError.Status, getPurchasesError)
		return
	} else {
		context.JSON(http.StatusOK, response.NewOkResponse("all purchases fetched", purchases))
		return
	}
}
