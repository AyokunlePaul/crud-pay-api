package authentication

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/password_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type authenticationUseCase struct {
	tokenManager    token.Manager
	userManager     user.Manager
	passwordService password_service.Service
}

func NewUseCase(tokenManager token.Manager, userManager user.Manager, service password_service.Service) UseCase {
	return &authenticationUseCase{
		tokenManager:    tokenManager,
		userManager:     userManager,
		passwordService: service,
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
	hashedPassword, passwordHashError := authentication.passwordService.Generate(user.Password)
	if passwordHashError != nil {
		return response.NewInternalServerError(response.ErrorCreatingUser)
	}
	user.Password = hashedPassword

	return authentication.userManager.Create(user)
}

func (authentication *authenticationUseCase) LogIn(user *user.User) *response.BaseResponse {
	userPassword := user.Password
	if validationError := user.CanLogin(); validationError != nil {
		return validationError
	}
	if getUserError := authentication.userManager.Get(user); getUserError != nil {
		return getUserError
	}
	if passwordComparisonError := authentication.passwordService.Compare(user.Password, userPassword); passwordComparisonError != nil {
		return response.NewBadRequestError(response.AuthenticationError)
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
