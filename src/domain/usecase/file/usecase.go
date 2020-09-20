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
	userId, userIdError := useCase.tokenManager.Get(token)
	if userIdError != nil {
		return userIdError
	}
	return useCase.fileManager.Create(userId, payFile)
}

func (useCase *fileUploadUseCase) UploadFiles(token string, crudPayFiles []*file.CrudPayFile) *response.BaseResponse {
	userId, userIdError := useCase.tokenManager.Get(token)
	if userIdError != nil {
		return userIdError
	}
	return useCase.fileManager.CreateList(userId, crudPayFiles)
}
