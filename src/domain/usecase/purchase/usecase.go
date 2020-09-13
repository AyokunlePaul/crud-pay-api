package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	purchaseEntity "github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type useCase struct {
	purchaseManager purchaseEntity.Manager
	tokenManager    token.Manager
	userManager     user.Manager
	productManager  product.Manager
	timelineManager timeline.Manager
}

func NewUseCase(
	tokenManager token.Manager, userManager user.Manager, timelineManager timeline.Manager,
	purchaseManager purchaseEntity.Manager, productManager product.Manager,
) UseCase {
	return &useCase{
		tokenManager:    tokenManager,
		userManager:     userManager,
		purchaseManager: purchaseManager,
		productManager:  productManager,
		timelineManager: timelineManager,
	}
}

func (useCase *useCase) CreatePurchase(token string, purchase *purchaseEntity.Purchase) *response.BaseResponse {
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
	if userId == productToBeBought.OwnerId.Hex() {
		return response.NewBadRequestError("you can't buy your own product")
	}
	if purchase.Type == timeline.TypeInstallment && !productToBeBought.AllowInstallment {
		return response.NewBadRequestError("product does not allow installment payment")
	}
	if purchase.NumberOfInstallments > productToBeBought.MaxInstallments {
		return response.NewBadRequestError("specify a lower payment installment number")
	}

	return nil
}

func (useCase *useCase) UpdatePurchase(token string, purchase *purchaseEntity.Purchase) *response.BaseResponse {
	panic("implement me")
}

func (useCase *useCase) GetAllPurchasesMadeByUser(token string) ([]purchaseEntity.Purchase, *response.BaseResponse) {
	panic("implement me")
}

func (useCase *useCase) GetPurchase(purchaseId, token string) (*purchaseEntity.Purchase, *response.BaseResponse) {
	panic("implement me")
}
