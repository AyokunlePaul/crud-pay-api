package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/thoas/go-funk"
	"time"
)

type useCase struct {
	purchaseManager purchase.Manager
	tokenManager    token.Manager
	userManager     user.Manager
	productManager  product.Manager
	timelineManager timeline.Manager
}

func New(
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
	productToBeBought, productError := useCase.productManager.Get(purchase.ProductId)
	if productError != nil {
		return productError
	}
	if validationError := productToBeBought.CanBePurchased(userId, purchase); validationError != nil {
		return validationError
	}

	timelines := timeline.NewTimeline(
		purchase.Id, productToBeBought.Amount, purchase.DeliveryArea.ShippingFee,
		purchase.NumberOfInstallments, purchase.Duration, purchase.Type,
	)

	purchase.Timeline = timelines
	purchase.OwnerId = productToBeBought.OwnerId
	purchase.Amount = productToBeBought.Amount
	purchase.TimelineAmount = timelines[0].Amount
	purchase.ShippingFee = timelines[0].ShippingFee
	purchase.TotalAmount = purchase.TimelineAmount + purchase.ShippingFee
	purchase.CreatedBy, _ = entity.StringToCrudPayId(userId)

	return useCase.purchaseManager.Create(purchase)
}

func (useCase *useCase) UpdatePurchase(token, purchaseId, reference string, amount float64) *response.BaseResponse {
	_, tokenError := useCase.tokenManager.Get(token)
	if tokenError != nil {
		return tokenError
	}

	currentPurchase := new(purchase.Purchase)
	currentPurchase.Id, _ = entity.StringToCrudPayId(purchaseId)

	if getPurchaseError := useCase.purchaseManager.Get(currentPurchase); getPurchaseError != nil {
		return getPurchaseError
	}
	currentPurchase.Reference = reference

	currentTime := time.Now()

	//Validate and update payment timeline
	for _, currentTimeline := range currentPurchase.Timeline {
		//Break after the first unpaid timeline
		if !currentTimeline.Paid {
			if amount != (currentTimeline.Amount + currentTimeline.ShippingFee) {
				return response.NewBadRequestError("payment amount doesn't match expected amount")
			}
			currentTimeline.Paid = true
			currentTimeline.ActualPaymentDate = &currentTime
			break
		}
	}
	//A purchase is successful when all the payment timeline are paid
	currentPurchase.Successful = len(funk.Filter(currentPurchase.Timeline, func(currentTimeline timeline.Timeline) bool {
		return !currentTimeline.Paid
	}).([]timeline.Timeline)) == 0

	//Validate payment against paystack here

	if updatePurchaseError := useCase.purchaseManager.Update(currentPurchase); updatePurchaseError != nil {
		return updatePurchaseError
	}

	return nil
}

func (useCase *useCase) GetAllPurchasesMadeByUser(token string) ([]purchase.Purchase, *response.BaseResponse) {
	id, tokenError := useCase.tokenManager.Get(token)
	if tokenError != nil {
		return nil, tokenError
	}
	userId, _ := entity.StringToCrudPayId(id)
	return useCase.purchaseManager.List(userId)
}

func (useCase *useCase) GetPurchase(token, purchaseId string) (*purchase.Purchase, *response.BaseResponse) {
	_, tokenError := useCase.tokenManager.Get(token)
	if tokenError != nil {
		return nil, tokenError
	}
	currentPurchase := new(purchase.Purchase)
	currentPurchase.Id, _ = entity.StringToCrudPayId(purchaseId)

	if getPurchaseError := useCase.purchaseManager.Get(currentPurchase); getPurchaseError != nil {
		return nil, getPurchaseError
	} else {
		return currentPurchase, nil
	}
}
