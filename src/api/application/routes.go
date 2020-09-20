package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/handler"
	"github.com/AyokunlePaul/crud-pay-api/src/api/middleware"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/search"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/authentication"
	productUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/product"
	purchaseUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/infra/database"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/password_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/token_service"
	"github.com/gin-gonic/gin"
)

func initializeDatabases() {
	database.Init()
	user.Init()
	purchase.Init()
	timeline.Init()
	product.Init()
	search.Init()
}

var (
	authenticationHandler   handler.Authentication
	productHandler          handler.Product
	purchaseHandler         handler.Purchase
	authorizationMiddleware gin.HandlerFunc
)

func setUpRepositoriesAndManagers() {
	tokenService := token_service.New()
	errorService := error_service.New()

	authorizationMiddleware = middleware.AuthorizationMiddleWare(tokenService)

	userManager := user.NewManager(user.NewDatabaseRepository(errorService))
	tokenManager := token.NewManager(token.NewDatabaseRepository(errorService), tokenService)
	productManager := product.NewManager(product.NewDatabaseRepository(errorService))
	searchManager := search.NewManager(search.NewDatabaseRepository(errorService))
	timelineManager := timeline.NewManager(timeline.NewDatabaseRepository(errorService))
	purchaseManager := purchase.NewManager(purchase.NewDatabaseRepository(errorService))

	authenticationHandler = handler.ForAuthentication(authentication.NewUseCase(tokenManager, userManager, password_service.New()))
	productHandler = handler.ForProduct(productUseCase.New(productManager, tokenManager, searchManager, userManager))
	purchaseHandler = handler.ForPurchase(purchaseUseCase.New(tokenManager, userManager, timelineManager, purchaseManager, productManager))
}

func mapRoutes() {
	v1Group := crudPayRouter.Group("/v1")
	{
		authenticationGroup := v1Group.Group("/user")
		{
			authenticationGroup.POST("/login", authenticationHandler.Login)
			authenticationGroup.POST("/create", authenticationHandler.CreateAccount)
			authenticationGroup.PUT("/update", authorizationMiddleware, authenticationHandler.UpdateUser)
			authenticationGroup.POST("/reset_password", authorizationMiddleware, authenticationHandler.ResetPassword)
			authenticationGroup.POST("/refresh_token", authenticationHandler.RefreshToken)
		}
		productGroup := v1Group.Group("/product")
		{
			productGroup.POST("/create", authorizationMiddleware, productHandler.Create)
			productGroup.GET("/:product_id", authorizationMiddleware, productHandler.Get)
		}
		searchGroup := v1Group.Group("/search")
		{
			searchGroup.GET("/product", authorizationMiddleware, productHandler.Search)
		}
		purchaseGroup := v1Group.Group("/purchase")
		{
			purchaseGroup.POST("/create", authorizationMiddleware, purchaseHandler.Create)
			purchaseGroup.GET("/all", authorizationMiddleware, purchaseHandler.List)
			purchaseGroup.GET("/product/:product_id", authorizationMiddleware, purchaseHandler.Get)
		}
	}
}
