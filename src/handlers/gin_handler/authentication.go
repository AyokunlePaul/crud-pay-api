package gin_handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user/user_service"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/models"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthenticationHandler interface {
	Login(*gin.Context)
	CreateAccount(*gin.Context)
	ResetPassword(*gin.Context)
	UpdateUser(*gin.Context)
	RefreshToken(*gin.Context)
}

type authenticationHandler struct {
	service user_service.Service
}

func NewAuthenticationHandler(service user_service.Service) AuthenticationHandler {
	return &authenticationHandler{
		service: service,
	}
}

func (handler *authenticationHandler) Login(context *gin.Context) {
	var userPayload models.UserPayload
	_ = context.BindJSON(&userPayload.Payload)

	if validationError := userPayload.CanLogin(); validationError != nil {
		context.JSON(validationError.Status, validationError)
		return
	}

	result, loginError := handler.service.Get(userPayload.ToDomainUser())
	if loginError != nil {
		context.JSON(loginError.Status, loginError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("login successful", result))
}

func (handler *authenticationHandler) CreateAccount(context *gin.Context) {
	var userPayload models.UserPayload
	_ = context.BindJSON(&userPayload.Payload)

	if validationError := userPayload.CanBeCreated(); validationError != nil {
		context.JSON(validationError.Status, validationError)
		return
	}

	result, loginError := handler.service.CreateUser(userPayload.ToDomainUser())
	if loginError != nil {
		context.JSON(loginError.Status, loginError)
		return
	}
	context.JSON(http.StatusCreated, response.NewCreatedResponse("user successfully created", result))
}

func (handler *authenticationHandler) ResetPassword(context *gin.Context) {
	var payload map[string]string
	_ = context.BindJSON(&payload)
	userEmail, ok := payload["email"]
	if !ok {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("email is missing"))
		return
	}
	if passwordResetError := handler.service.ResetPassword(userEmail); passwordResetError != nil {
		context.JSON(passwordResetError.Status, passwordResetError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("verification code sent", nil))
}

func (handler *authenticationHandler) UpdateUser(context *gin.Context) {
	userToken := strings.Split(context.GetHeader("Authorization"), " ")[1]

	var userPayload models.UserPayload
	_ = context.BindJSON(&userPayload.Payload)

	result, updateUserError := handler.service.Update(userPayload.ToDomainUser(), userToken)

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
	result, updateUserError := handler.service.RefreshToken(userToken)
	if updateUserError != nil {
		context.JSON(updateUserError.Status, updateUserError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("token successfully refreshed", result))
}
