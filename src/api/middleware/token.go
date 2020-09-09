package middleware

import (
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/token_service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthorizationMiddleWare(service token_service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenPayload := context.GetHeader("Authorization")
		if strings.TrimSpace(tokenPayload) == "" {
			context.JSON(http.StatusUnauthorized, response.NewUnAuthorizedError())
			context.Abort()
			return
		}
		bearerToken := strings.Split(tokenPayload, " ")
		if len(bearerToken) != 2 {
			context.JSON(http.StatusUnauthorized, response.NewUnAuthorizedError())
			context.Abort()
			return
		}
		userToken := bearerToken[1]
		if tokenValidityError := service.CheckTokenValidity(userToken, true); tokenValidityError != nil {
			context.JSON(http.StatusUnauthorized, response.NewUnAuthorizedError())
			context.Abort()
			return
		}
		context.Next()
	}
}
