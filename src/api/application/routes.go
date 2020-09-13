package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/handler"
	"github.com/AyokunlePaul/crud-pay-api/src/api/middleware"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/search"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/authentication"
	productUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/product"
	"github.com/AyokunlePaul/crud-pay-api/src/infra/database"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/password_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/token_service"
)

func initializeDatabases() {
	database.Init()
	user.Init()
	purchase.Init()
	product.Init()
	search.Init()
}

var (
	authenticationHandler handler.Authentication
	productHandler        handler.Product
	tokenService          token_service.Service
)

func setUpRepositoriesAndManagers() {
	tokenService = token_service.New()
	errorService := error_service.New()

	userManager := user.NewManager(user.NewDatabaseRepository(errorService), password_service.New())
	tokenManager := token.NewManager(token.NewDatabaseRepository(errorService), tokenService)
	productManager := product.NewManager(product.NewDatabaseRepository(errorService))
	searchManager := search.NewManager(search.NewDatabaseRepository(errorService))
	_ = purchase.NewManager(purchase.NewDatabaseRepository(errorService))

	authenticationHandler = handler.ForAuthentication(authentication.NewUseCase(tokenManager, userManager))
	productHandler = handler.ForProduct(productUseCase.NewUseCase(productManager, tokenManager, searchManager, userManager))
}

func mapRoutes() {
	v1Group := crudPayRouter.Group("/v1")
	{
		authenticationGroup := v1Group.Group("/user")
		{
			authenticationGroup.POST("/login", authenticationHandler.Login)
			authenticationGroup.POST("/create", authenticationHandler.CreateAccount)
			authenticationGroup.PUT("/update", middleware.AuthorizationMiddleWare(tokenService), authenticationHandler.UpdateUser)
			authenticationGroup.POST("/reset_password", middleware.AuthorizationMiddleWare(tokenService), authenticationHandler.ResetPassword)
			authenticationGroup.POST("/refresh_token", authenticationHandler.RefreshToken)
		}
		productGroup := v1Group.Group("/product")
		{
			productGroup.POST("/", middleware.AuthorizationMiddleWare(tokenService), productHandler.Create)
			productGroup.GET("/:product_id", middleware.AuthorizationMiddleWare(tokenService), productHandler.Get)
		}
		searchGroup := v1Group.Group("/search")
		{
			searchGroup.GET("/product", middleware.AuthorizationMiddleWare(tokenService), productHandler.Search)
		}
	}
}
