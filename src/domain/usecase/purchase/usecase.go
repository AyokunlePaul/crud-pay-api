package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type useCase struct {
	purchaseManager purchase.Manager
	tokenManager    token.Manager
	userManager     user.Manager
	productManager  product.Manager
	timelineManager timeline.Manager
}

func NewUseCase(
	tokenManager token.Manager, userManager user.Manager, timelineManager timeline.Manager,
	purchaseManager purchase.Manager, productManager product.Manager,
) UseCase {
	return &useCase{
		tokenManager:    tokenManager,
		userManager:     userManager,
		purchaseManager: purchaseManager,
		productManager:  productManager,
		timelineManager: timelineManager,
	}
}

func (useCase *useCase) CreatePurchase(token string, purchase *purchase.Purchase) *response.BaseResponse {
	userId, tokenError := useCase.tokenManager.Get(token)
	if tokenError != nil {
		return tokenError
	}
	if !purchase.HasValidPaymentFrequency() {
		return response.NewBadRequestError("invalid payment frequency")
	}
	productToBeBought, productError := useCase.productManager.Get(purchase.ProductId)
	if productError != nil {
		return productError
	}
	if purchase.Type == timeline.TypeInstallment && !productToBeBought.AllowInstallment {
		return response.NewBadRequestError("product does not allow installment payment")
	}
	if userId == productToBeBought.OwnerId.Hex() {
		return response.NewBadRequestError("you can't buy your own product")
	}
	if purchase.NumberOfInstallments > productToBeBought.MaxInstallments {
		return response.NewBadRequestError("specify a lower payment installment number")
	}
	timelines := timeline.NewTimeline(purchase.Id, productToBeBought.Amount, purchase.NumberOfInstallments, purchase.Duration, purchase.Frequency)
	if timelineCreationError := useCase.timelineManager.CreateList(timelines); timelineCreationError != nil {
		return timelineCreationError
	}
	purchase.CreatedBy, _ = entity.StringToCrudPayId(userId)

	return useCase.purchaseManager.Create(purchase)
}

func (useCase *useCase) UpdatePurchase(token string, purchase *purchase.Purchase) *response.BaseResponse {
	panic("implement me")
}

func (useCase *useCase) GetAllPurchasesMadeByUser(token string) ([]purchase.Purchase, *response.BaseResponse) {
	panic("implement me")
}

func (useCase *useCase) GetPurchase(purchaseId, token string) (*purchase.Purchase, *response.BaseResponse) {
	panic("implement me")
}
