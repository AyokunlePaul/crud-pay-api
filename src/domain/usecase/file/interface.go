package file

import (
	"github.com/AyokunlePaul/crud-pay-api/src/domain/entity/file"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
)

type UseCase interface {
	UploadFile(string, *file.CrudPayFile) *response.BaseResponse
	UploadFiles(string, []file.CrudPayFile) *response.BaseResponse
}
