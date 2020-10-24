package handler

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/file"
	fileUseCase "github.com/AyokunlePaul/crud-pay-api/src/domain/usecase/file"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/logger"
	"github.com/AyokunlePaul/crud-pay-api/src/utils/string_utilities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type File interface {
	Create(*gin.Context)
}

type fileHandler struct {
	useCase fileUseCase.UseCase
}

func ForFileUpload(useCase fileUseCase.UseCase) File {
	return &fileHandler{
		useCase: useCase,
	}
}

func (handler *fileHandler) Create(context *gin.Context) {
	token := strings.Split(context.GetHeader("Authorization"), " ")[1]

	form, formError := context.MultipartForm()
	if formError != nil {
		logger.Error("parse form error", formError)
		context.JSON(http.StatusInternalServerError, response.NewInternalServerError("an error occurred"))
		return
	}
	uploadedFiles := form.File["files"]

	folder := context.PostForm("folder")
	if string_utilities.IsEmpty(folder) {
		context.JSON(http.StatusBadRequest, response.NewBadRequestError("invalid folder type"))
		return
	}
	newFiles, fileCreationError := file.NewList(context.Request, folder, uploadedFiles)
	if fileCreationError != nil {
		context.JSON(fileCreationError.Status, fileCreationError)
		return
	}

	if fileUploadError := handler.useCase.UploadFiles(token, newFiles); fileUploadError != nil {
		context.JSON(fileUploadError.Status, fileUploadError)
		return
	}
	uploadResponse := make([]string, len(uploadedFiles))
	for index, currentFile := range newFiles {
		uploadResponse[index] = currentFile.UploadedUrl
	}
	context.JSON(http.StatusCreated, response.NewOkResponse("file uploaded successfully", uploadResponse))
}
