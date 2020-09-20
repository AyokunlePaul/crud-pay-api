package file

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/file"
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/token"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type fileUploadUseCase struct {
	tokenManager token.Manager
	fileManager  file.Manager
}

func New(tokenManager token.Manager, fileManager file.Manager) UseCase {
	return &fileUploadUseCase{
		tokenManager: tokenManager,
		fileManager:  fileManager,
	}
}

func (useCase *fileUploadUseCase) UploadFile(token string, payFile *file.CrudPayFile) *response.BaseResponse {
	if userId, userIdError := useCase.tokenManager.Get(token); userIdError != nil {
		return userIdError
	} else {
		payFile.Folder = userId //Use the user id to create folder
	}
	return useCase.fileManager.Create(payFile)
}

func (useCase *fileUploadUseCase) UploadFiles(token string, crudPayFiles []file.CrudPayFile) *response.BaseResponse {
	if userId, userIdError := useCase.tokenManager.Get(token); userIdError != nil {
		return userIdError
	} else {
		//Use the user id to create folder
		for _, currentFile := range crudPayFiles {
			currentFile.Folder = userId
		}
	}
	return useCase.fileManager.CreateList(crudPayFiles)
}
