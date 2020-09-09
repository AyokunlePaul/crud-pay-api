package authentication

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type authenticationUseCase struct {
	tokenManager token.Manager
	userManager  user.Manager
}

func NewUseCase(tokenManager token.Manager, userManager user.Manager) UseCase {
	return &authenticationUseCase{
		tokenManager: tokenManager,
		userManager:  userManager,
	}
}

func (authentication *authenticationUseCase) Create(user *user.User) *response.BaseResponse {
	userToken := token.NewCrudPayToken()

	if tokenError := authentication.tokenManager.CreateToken(userToken, user.Id.Hex()); tokenError != nil {
		return tokenError
	}
	if validationError := user.CanBeCreated(); validationError != nil {
		return validationError
	}
	user.Token = userToken.AccessToken
	user.RefreshToken = userToken.RefreshToken

	return authentication.userManager.Create(user)
}

func (authentication *authenticationUseCase) LogIn(user *user.User) *response.BaseResponse {
	if validationError := user.CanLogin(); validationError != nil {
		return validationError
	}
	if getUserError := authentication.userManager.Get(user); getUserError != nil {
		return getUserError
	}
	return nil
}

func (authentication *authenticationUseCase) Update(token string, newUser user.User) (*user.User, *response.BaseResponse) {
	userId, userIdError := authentication.tokenManager.Get(token)
	if userIdError != nil {
		return nil, userIdError
	}
	oldUser := new(user.User)
	oldUser.Id, _ = entity.StringToCrudPayId(userId)

	if getUserError := authentication.userManager.Get(oldUser); getUserError != nil {
		return nil, getUserError
	}

	if validationError := oldUser.CanBeUpdatedWith(newUser); validationError != nil {
		return nil, validationError
	}

	if userUpdateError := authentication.userManager.Update(oldUser); userUpdateError != nil {
		return nil, userUpdateError
	}
	return oldUser, nil
}

func (authentication *authenticationUseCase) ForgotPassword(token string, email string) *response.BaseResponse {
	panic("implement me")
}

func (authentication *authenticationUseCase) RefreshToken(refreshToken string) (*user.User, *response.BaseResponse) {
	panic("implement me")
}
