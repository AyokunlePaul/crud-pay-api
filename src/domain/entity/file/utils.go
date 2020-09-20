package file

import (
	"github.com/AyokunlePaul/crud-pay-api/src/pkg/response"
	"mime/multipart"
	"net/http"
)

func New() *CrudPayFile {
	return new(CrudPayFile)
}

func NewList(request *http.Request, folder string, headers []*multipart.FileHeader) ([]CrudPayFile, *response.BaseResponse) {
	files := make([]CrudPayFile, len(headers))
	for index, header := range headers {
		currentFile := CrudPayFile{
			Folder:  folder,
			Header:  header,
			Request: request,
		}
		files[index] = currentFile
	}
	return files, nil
}
