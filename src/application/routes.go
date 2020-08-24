package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/gin_handler"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/authentication"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/token_repository"
)

func mapRoutes() {
	tokenRepository := token_repository.NewTokenRepository()
	userDatabaseRepository := authentication.NewUserDatabaseRepository(tokenRepository)
	userService := user.NewUserService(userDatabaseRepository)
	userEndpointHandler := gin_handler.NewAuthenticationHandler(userService)

	v1Router := router.Group("/v1")
	{
		authenticationGroup := v1Router.Group("/user")
		{
			authenticationGroup.POST("/login", userEndpointHandler.Login)
			authenticationGroup.POST("/create", userEndpointHandler.CreateAccount)
			authenticationGroup.PUT("/update", userEndpointHandler.UpdateUser)
			authenticationGroup.POST("/reset_password", userEndpointHandler.ResetPassword)
		}
	}
}
