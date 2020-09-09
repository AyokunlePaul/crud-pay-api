package handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/presenter/models/user_payload"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/authentication"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Authentication interface {
	Login(*gin.Context)
	CreateAccount(*gin.Context)
	ResetPassword(*gin.Context)
	UpdateUser(*gin.Context)
	RefreshToken(*gin.Context)
}

type authenticationHandler struct {
	useCase authentication.UseCase
}

func ForAuthentication(useCase authentication.UseCase) Authentication {
	return &authenticationHandler{
		useCase: useCase,
	}
}

func (handler *authenticationHandler) Login(context *gin.Context) {
	var userPayload user_payload.UserPayload
	_ = context.BindJSON(&userPayload.Payload)

	user := userPayload.FromEmailAndPasswordToDomain()
	loginError := handler.useCase.LogIn(user)

	if loginError != nil {
		context.JSON(loginError.Status, loginError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("login successful", user))
}

func (handler *authenticationHandler) CreateAccount(context *gin.Context) {
	var userPayload user_payload.UserPayload
	_ = context.BindJSON(&userPayload.Payload)

	user := userPayload.ToDomain()
	createUserError := handler.useCase.Create(user)

	if createUserError != nil {
		context.JSON(createUserError.Status, createUserError)
		return
	}
	context.JSON(http.StatusCreated, response.NewCreatedResponse("user successfully created", user))
}

func (handler *authenticationHandler) ResetPassword(context *gin.Context) {
	userToken := strings.Split(context.GetHeader("Authorization"), " ")[1]
	var payload map[string]string
	_ = context.BindJSON(&payload)
	userEmail, ok := payload["email"]
	if !ok {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("email is missing"))
		return
	}
	if passwordResetError := handler.useCase.ForgotPassword(userToken, userEmail); passwordResetError != nil {
		context.JSON(passwordResetError.Status, passwordResetError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("verification code sent", nil))
}

func (handler *authenticationHandler) UpdateUser(context *gin.Context) {
	userToken := strings.Split(context.GetHeader("Authorization"), " ")[1]

	var userPayload user_payload.UserPayload
	_ = context.BindJSON(&userPayload.Payload)
	user := userPayload.ToDomain()

	result, updateUserError := handler.useCase.Update(userToken, *user)

	if updateUserError != nil {
		context.JSON(updateUserError.Status, updateUserError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("user updated successfully", result))
}

func (handler *authenticationHandler) RefreshToken(context *gin.Context) {
	tokenMap := make(map[string]string, 1)
	_ = context.BindJSON(&tokenMap)
	var userToken string
	var ok bool

	if len(tokenMap) < 1 {
		context.JSON(http.StatusUnauthorized, response.NewUnAuthorizedError())
		return
	} else {
		if userToken, ok = tokenMap["refresh_token"]; !ok {
			context.JSON(http.StatusUnauthorized, response.NewUnAuthorizedError())
			return
		}
	}
	result, updateUserError := handler.useCase.RefreshToken(userToken)
	if updateUserError != nil {
		context.JSON(updateUserError.Status, updateUserError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("token successfully refreshed", result))
}
