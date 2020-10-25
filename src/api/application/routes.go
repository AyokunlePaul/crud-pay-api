package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/handler"
	"github.com/AyokunlePaul/crud-pay-api/src/api/middleware"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/file"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/product"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/search"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/timeline"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/user"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/admin"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/authentication"
	fileUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/file"
	productUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/product"
	purchaseUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/purchase"
	"github.com/AyokunlePaul/crud-pay-api/src/infra/database"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/password_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/token_service"
	"github.com/gin-gonic/gin"
)

func initializeRepositories() {
	database.Init()
	user.Init()
	purchase.Init()
	timeline.Init()
	product.Init()
	search.Init()
}

var (
	fileHandler             handler.File
	adminHandler            handler.Admin
	productHandler          handler.Product
	purchaseHandler         handler.Purchase
	authenticationHandler   handler.Authentication
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
	fileManager := file.NewManager(file.NewStorageRepository(errorService))

	fileHandler = handler.ForFileUpload(fileUseCase.New(tokenManager, fileManager))
	authenticationHandler = handler.ForAuthentication(authentication.NewUseCase(tokenManager, userManager, password_service.New()))
	productHandler = handler.ForProduct(productUseCase.New(productManager, tokenManager, searchManager, userManager))
	purchaseHandler = handler.ForPurchase(purchaseUseCase.New(tokenManager, userManager, timelineManager, purchaseManager, productManager))
	adminHandler = handler.ForAdmin(admin.New(tokenManager, purchaseManager, userManager, searchManager))
}

func mapRoutes() {
	v1Group := crudPayRouter.Group("/v1")
	{
		authenticationGroup := v1Group.Group("/user")
		{
			authenticationGroup.POST("/login", authenticationHandler.Login)
			authenticationGroup.POST("/create", authenticationHandler.CreateAccount)
			authenticationGroup.PUT("/", authorizationMiddleware, authenticationHandler.UpdateUser)
			authenticationGroup.POST("/reset_password", authorizationMiddleware, authenticationHandler.ResetPassword)
			authenticationGroup.POST("/refresh_token", authenticationHandler.RefreshToken)
		}
		productGroup := v1Group.Group("/product", authorizationMiddleware)
		{
			productGroup.POST("/", productHandler.Create)
			productGroup.GET("/", productHandler.Get)
			productGroup.GET("/vendor/:owner_id", productHandler.GetVendorProduct)
			productGroup.GET("/details/:product_id", productHandler.Get)
			productGroup.PUT("/:product_id", productHandler.Update)
		}
		searchGroup := v1Group.Group("/search", authorizationMiddleware)
		{
			searchGroup.GET("/product", productHandler.Search)
		}
		purchaseGroup := v1Group.Group("/purchase", authorizationMiddleware)
		{
			purchaseGroup.POST("/", purchaseHandler.Create)
			purchaseGroup.POST("/:purchase_id", purchaseHandler.Update)
			purchaseGroup.GET("/", purchaseHandler.List)
			purchaseGroup.GET("/product/:purchase_id", purchaseHandler.Get)
		}
		fileUploadGroup := v1Group.Group("/file", authorizationMiddleware)
		{
			fileUploadGroup.POST("/", fileHandler.Create)
		}
		adminGroup := v1Group.Group("/admin")
		{
			adminGroup.GET("/", adminHandler.GetDailyStat)
			adminGroup.GET("/search", adminHandler.Search)
		}
	}
}
