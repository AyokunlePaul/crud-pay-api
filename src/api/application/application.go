package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/api/middleware"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

var crudPayRouter *gin.Engine

func init() {
	crudPayRouter = gin.New()
	zapLogger := logger.GetLogger()
	crudPayRouter.Use(middleware.RequestLoggerMiddleware())
	crudPayRouter.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	crudPayRouter.Use(ginzap.RecoveryWithZap(zapLogger, true))
}

func StartApplication() {
	initializeRepositories()
	setUpRepositoriesAndManagers()
	mapRoutes()
	logger.Error("application start error", crudPayRouter.Run(":8080"))
}
