package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user/user_service"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/gin_handler"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/gin_handler/middlewares"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/token_repository"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/user_database_repository"
)

func mapRoutes() {
	tokenRepository := token_repository.NewTokenRepository()
	userDatabaseRepository := user_database_repository.New(tokenRepository)
	userService := user_service.New(userDatabaseRepository)
	userEndpointHandler := gin_handler.NewAuthenticationHandler(userService)

	v1Router := crudPayRouter.Group("/v1")
	{
		authenticationGroup := v1Router.Group("/user")
		{
			authenticationGroup.POST("/login", userEndpointHandler.Login)
			authenticationGroup.POST("/create", userEndpointHandler.CreateAccount)
			authenticationGroup.PUT("/update", middlewares.AuthorizationMiddleWare(), userEndpointHandler.UpdateUser)
			authenticationGroup.POST("/reset_password", middlewares.AuthorizationMiddleWare(), userEndpointHandler.ResetPassword)
			authenticationGroup.POST("/refresh_token", userEndpointHandler.RefreshToken)
		}
	}
}
