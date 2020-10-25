package authentication

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/password_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type useCase struct {
	tokenManager    token.Manager
	userManager     user.Manager
	passwordService password_service.Service
}

func NewUseCase(tokenManager token.Manager, userManager user.Manager, service password_service.Service) UseCase {
	return &useCase{
		tokenManager:    tokenManager,
		userManager:     userManager,
		passwordService: service,
	}
}

func (useCase *useCase) Create(user *user.User) *response.BaseResponse {
	userToken := token.NewCrudPayToken()

	if validationError := user.CanBeCreated(); validationError != nil {
		return validationError
	}
	if tokenError := useCase.tokenManager.CreateToken(userToken, user.Id.Hex()); tokenError != nil {
		return tokenError
	}
	user.Token = userToken.AccessToken
	user.RefreshToken = userToken.RefreshToken
	hashedPassword, passwordHashError := useCase.passwordService.Generate(user.Password)
	if passwordHashError != nil {
		return response.NewInternalServerError(response.ErrorCreatingUser)
	}
	user.Password = hashedPassword

	return useCase.userManager.Create(user)
}

func (useCase *useCase) LogIn(user *user.User) *response.BaseResponse {
	userPassword := user.Password
	if validationError := user.CanLogin(); validationError != nil {
		return validationError
	}
	if getUserError := useCase.userManager.Get(user); getUserError != nil {
		return getUserError
	}
	if passwordComparisonError := useCase.passwordService.Compare(user.Password, userPassword); passwordComparisonError != nil {
		return response.NewBadRequestError(response.AuthenticationError)
	}
	return nil
}

func (useCase *useCase) Update(token string, newUser user.User) (*user.User, *response.BaseResponse) {
	userId, userIdError := useCase.tokenManager.Get(token)
	if userIdError != nil {
		return nil, userIdError
	}
	oldUser := new(user.User)
	oldUser.Id, _ = entity.StringToCrudPayId(userId)

	if getUserError := useCase.userManager.Get(oldUser); getUserError != nil {
		return nil, getUserError
	}

	if validationError := oldUser.CanBeUpdatedWith(newUser); validationError != nil {
		return nil, validationError
	}

	if userUpdateError := useCase.userManager.Update(oldUser); userUpdateError != nil {
		return nil, userUpdateError
	}
	return oldUser, nil
}

func (useCase *useCase) ForgotPassword(token string, email string) *response.BaseResponse {
	panic("implement me")
}

func (useCase *useCase) RefreshToken(refreshToken string) (*user.User, *response.BaseResponse) {
	userToken := token.NewCrudPayToken()
	refreshTokenError := useCase.tokenManager.RefreshToken(userToken, refreshToken, "")
	if refreshTokenError != nil {
		return nil, refreshTokenError
	}

	userId, _ := useCase.tokenManager.Get(userToken.AccessToken)
	newUser := user.User{
		Token:        userToken.AccessToken,
		RefreshToken: userToken.RefreshToken,
	}
	newUser.Id, _ = entity.StringToCrudPayId(userId)

	return useCase.Update(userToken.AccessToken, newUser)
}
