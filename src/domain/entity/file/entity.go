package file

import (
	"mime/multipart"
	"net/http"
)

type Type string

type CrudPayFile struct {
	Folder      string
	Header      *multipart.FileHeader
	Request     *http.Request
	UploadedUrl string
}
