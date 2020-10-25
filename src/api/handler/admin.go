package handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/admin"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Admin interface {
	GetDailyStat(*gin.Context)
	Search(*gin.Context)
}

type adminHandler struct {
	adminUseCase admin.UseCase
}

func ForAdmin(adminUseCase admin.UseCase) Admin {
	return &adminHandler{
		adminUseCase: adminUseCase,
	}
}

func (handler *adminHandler) GetDailyStat(context *gin.Context) {
	dailyStat, dailyStatError := handler.adminUseCase.GetDailyStat()
	if dailyStatError != nil {
		context.JSON(dailyStatError.Status, dailyStatError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("daily stat fetched", dailyStat))
}

func (handler *adminHandler) Search(context *gin.Context) {
	query := context.Query("query")
	users, searchError := handler.adminUseCase.Search(query)
	if searchError != nil {
		context.JSON(searchError.Status, searchError)
		return
	}
	context.JSON(http.StatusOK, response.NewOkResponse("users fetched", users))
}
