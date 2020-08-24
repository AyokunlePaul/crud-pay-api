package application

import (
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication() {
	mapRoutes()
	logger.Error("application start error", router.Run(":8080"))
}
