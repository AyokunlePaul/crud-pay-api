package middlewares

import (
	"github.com/AyokunlePaul/crud-pay-api/src/utils/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthorizationMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenPayload := context.GetHeader("Authorization")
		if tokenPayload == "" {
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
		if tokenValidityError := utilities.CheckTokenValidity(userToken, true); tokenValidityError != nil {
			context.JSON(http.StatusUnauthorized, response.NewUnAuthorizedError())
			context.Abort()
			return
		}
		context.Next()
	}
}
