package file

import (
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/error_service"
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
)

type repository struct {
	errorService error_service.Service
}

const (
	bucketName = "crud_pay_test_bucket"
	bucketUrl  = "https://storage.cloud.google.com"
)

func NewStorageRepository(errorService error_service.Service) Repository {
	return &repository{
		errorService: errorService,
	}
}

func (repository *repository) Create(userId string, file *CrudPayFile) *response.BaseResponse {
	clientContext := appengine.NewContext(file.Request)
	client, clientError := storage.NewClient(clientContext, option.WithCredentialsFile("keys.json"))
	if clientError != nil {
		return repository.errorService.HandleGoogleStorageError(clientError)
	}

	name := fmt.Sprintf("%s/%s/%s", userId, file.Folder, file.Header.Filename)
	storageWriter := client.Bucket(bucketName).Object(name).NewWriter(clientContext)

	headerFile, fileOpenError := file.Header.Open()
	if fileOpenError != nil {
		return repository.errorService.HandleGoogleStorageError(fileOpenError)
	}

	if _, fileCopyError := io.Copy(storageWriter, headerFile); fileCopyError != nil {
		return repository.errorService.HandleGoogleStorageError(fileCopyError)
	}

	if storageWriterCloseError := storageWriter.Close(); storageWriterCloseError != nil {
		return repository.errorService.HandleGoogleStorageError(storageWriterCloseError)
	}

	file.UploadedUrl = bucketUrl + "/" + bucketName + "/" + storageWriter.Attrs().Name
	return nil
}

func (repository *repository) CreateList(userId string, files []*CrudPayFile) *response.BaseResponse {
	for _, file := range files {
		repository.Create(userId, file)
	}
	return nil
}
