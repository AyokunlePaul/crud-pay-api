package file

import (
	"cloud.google.com/go/storage"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

type repository struct {
	errorService error_service.Service
}

func NewStorageRepository(errorService error_service.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) Create(file *CrudPayFile) *response.BaseResponse {
	clientContext := appengine.NewContext(file.Request)
	client, clientError := storage.NewClient(clientContext, option.WithCredentialsFile("test_keys.json"))
}

func (repository *repository) CreateList(files []CrudPayFile) *response.BaseResponse {
	panic("implement me")
}
