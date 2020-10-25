package admin

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/search"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"time"
)

type useCase struct {
	tokenManager    token.Manager
	purchaseManager purchase.Manager
	userManager     user.Manager
	searchManager   search.Manager
}

func New(tokenManager token.Manager, purchaseManager purchase.Manager, userManager user.Manager, searchManager search.Manager) UseCase {
	return &useCase{
		tokenManager:    tokenManager,
		purchaseManager: purchaseManager,
		userManager:     userManager,
		searchManager:   searchManager,
	}
}

func (useCase *useCase) CreateNew(user *user.User) *response.BaseResponse {
	userToken := token.NewCrudPayToken()

	if emailValidationError := user.IsValidEmail(); emailValidationError != nil {
		return emailValidationError
	}
	if tokenError := useCase.tokenManager.CreateToken(userToken, user.Id.Hex()); tokenError != nil {
		return tokenError
	}
	user.Token = userToken.AccessToken
	user.RefreshToken = userToken.RefreshToken
	user.IsAdmin = true

	if adminCreationError := useCase.userManager.Create(user); adminCreationError != nil {
		return adminCreationError
	}

	return nil
}

func (useCase *useCase) GetAll() ([]user.User, *response.BaseResponse) {
	return useCase.userManager.ListAdmin()
}

func (useCase *useCase) GetDailyStat() (map[string]interface{}, *response.BaseResponse) {
	from := time.Now().AddDate(0, 0, -1)
	to := time.Now()

	newSignUpCount, countError := useCase.userManager.List(from, to)
	if countError != nil {
		return nil, countError
	}

	totalPaymentsCount, paymentCountError := useCase.purchaseManager.ListData(from, to)
	if paymentCountError != nil {
		return nil, paymentCountError
	}

	return map[string]interface{}{
		"new_signups":   newSignUpCount,
		"payments_made": totalPaymentsCount,
	}, nil
}

func (useCase *useCase) Delete(user *user.User) *response.BaseResponse {
	return useCase.userManager.Delete(user.Id)
}

func (useCase *useCase) Search(query string) ([]user.User, *response.BaseResponse) {
	searchParameter := search.Param{
		Index: "crudpay.users",
		Query: query,
		Names: []string{"first_name", "last_name"},
	}
	users, userSearchError := useCase.searchManager.SearchUser(searchParameter)
	if userSearchError != nil {
		return nil, userSearchError
	}
	return users, nil
}
