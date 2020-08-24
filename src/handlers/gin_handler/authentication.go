package gin_handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
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
}

type authenticationHandler struct {
	service user.Service
}

func NewAuthenticationHandler(service user.Service) AuthenticationHandler {
	return &authenticationHandler{
		service: service,
	}
}

func (handler *authenticationHandler) Login(context *gin.Context) {
	var userPayload user.User
	_ = context.BindJSON(&userPayload)
	result, loginError := handler.service.Get(userPayload)
	if loginError != nil {
		context.JSON(loginError.Status, loginError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("login successful", result))
}

func (handler *authenticationHandler) CreateAccount(context *gin.Context) {
	var userPayload user.User
	_ = context.BindJSON(&userPayload)
	result, loginError := handler.service.CreateUser(userPayload)
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
	bearerToken := context.GetHeader("Authorization")
	tokenArray := strings.Split(bearerToken, " ")
	var userToken string
	if len(tokenArray) != 2 {
		userToken = tokenArray[1]
	} else {
		context.JSON(http.StatusUnauthorized, response.NewBadRequestError("invalid token"))
		return
	}
	result, updateUserError := handler.service.Update(user.User{
		Token: userToken,
	})
	if updateUserError != nil {
		context.JSON(updateUserError.Status, updateUserError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("user updated successfully", result))
}
