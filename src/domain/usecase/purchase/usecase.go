package purchase

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/string_utilities"
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
		purchase.Id, productToBeBought.Amount, purchase.ShippingFee,
		purchase.NumberOfInstallments, purchase.Duration, purchase.Type,
	)

	purchase.Timeline = timelines
	purchase.OwnerId = productToBeBought.OwnerId
	purchase.Amount = productToBeBought.Amount
	purchase.TimelineAmount = timelines[0].Amount
	purchase.TotalAmount = purchase.TimelineAmount + purchase.ShippingFee
	purchase.CreatedBy, _ = entity.StringToCrudPayId(userId)

	_ = useCase.userManager.IncrementTotalPurchase(purchase.CreatedBy)

	return useCase.purchaseManager.Create(purchase)
}

func (useCase *useCase) UpdatePurchase(token string, update purchase.Update) (*purchase.Purchase, *response.BaseResponse) {
	userId, tokenError := useCase.tokenManager.Get(token)
	if tokenError != nil {
		return nil, tokenError
	}

	if string_utilities.IsEmpty(update.Reference) {
		return nil, response.NewBadRequestError("invalid payment reference")
	}

	currentPurchase := new(purchase.Purchase)
	currentPurchase.Id, _ = entity.StringToCrudPayId(update.PurchaseId)

	if getPurchaseError := useCase.purchaseManager.Get(currentPurchase); getPurchaseError != nil {
		return nil, getPurchaseError
	}
	currentPurchase.Reference = update.Reference
	if currentPurchase.Successful {
		return currentPurchase, nil
	}

	if currentPurchase.CreatedBy.Hex() != userId {
		return nil, response.NewBadRequestError("user not authorized to update purchase")
	}

	currentTime := time.Now()

	//Validate and update payment timeline
	for index, currentTimeline := range currentPurchase.Timeline {
		//Break after the first unpaid timeline
		if !currentTimeline.Paid {
			if update.Amount != (currentTimeline.Amount + currentTimeline.ShippingFee) {
				return nil, response.NewBadRequestError("payment amount doesn't match expected amount")
			}
			currentPurchase.Timeline[index].Paid = true
			currentPurchase.Timeline[index].ActualPaymentDate = &currentTime
			break
		}
	}
	//A purchase is successful when all the payment timeline are paid
	currentPurchase.Successful = len(funk.Filter(currentPurchase.Timeline, func(currentTimeline timeline.Timeline) bool {
		return !currentTimeline.Paid
	}).([]timeline.Timeline)) == 0

	//Validate payment against paystack here

	//Update local database
	if updatePurchaseError := useCase.purchaseManager.Update(currentPurchase); updatePurchaseError != nil {
		return nil, updatePurchaseError
	}

	return currentPurchase, nil
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
