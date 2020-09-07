package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/authentication/domain/user/user_service"
	"github.com/AyokunlePaul/crud-pay-api/src/clients/elasticsearch_client"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/gin_handler/authentication_handler"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/gin_handler/middlewares"
	"github.com/AyokunlePaul/crud-pay-api/src/handlers/gin_handler/product_handler"
	"github.com/AyokunlePaul/crud-pay-api/src/product/domain/product/product_service"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/elasticsearch_database_repository"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/product_database_repository"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/token_repository"
	"github.com/AyokunlePaul/crud-pay-api/src/repository/user_database_repository"
)

func initializeDatabases() {
	elasticsearch_client.Init()
}

func mapRoutes() {
	tokenRepository := token_repository.New()
	searchRepository := elasticsearch_database_repository.New()

	userDatabaseRepository := user_database_repository.New(tokenRepository)
	userService := user_service.New(userDatabaseRepository)
	userEndpointHandler := authentication_handler.New(userService)

	productDatabaseRepository := product_database_repository.New(tokenRepository, searchRepository)
	productService := product_service.New(productDatabaseRepository)
	productHandler := product_handler.New(productService)

	v1Group := crudPayRouter.Group("/v1")
	{
		authenticationGroup := v1Group.Group("/user")
		{
			authenticationGroup.POST("/login", userEndpointHandler.Login)
			authenticationGroup.POST("/create", userEndpointHandler.CreateAccount)
			authenticationGroup.PUT("/update", middlewares.AuthorizationMiddleWare(), userEndpointHandler.UpdateUser)
			authenticationGroup.POST("/reset_password", middlewares.AuthorizationMiddleWare(), userEndpointHandler.ResetPassword)
			authenticationGroup.POST("/refresh_token", userEndpointHandler.RefreshToken)
		}
		productGroup := v1Group.Group("/product")
		{
			productGroup.POST("/", middlewares.AuthorizationMiddleWare(), productHandler.Create)
			productGroup.GET("/:product_id", middlewares.AuthorizationMiddleWare(), productHandler.Get)
		}
		searchGroup := v1Group.Group("/search")
		{
			searchGroup.GET("/product", middlewares.AuthorizationMiddleWare(), productHandler.Search)
		}
	}
}
