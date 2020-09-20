package file

import (
	"mime/multipart"
	"net/http"
)

type CrudPayFile struct {
	BucketName  string
	Folder      string
	File        multipart.File
	Request     *http.Request
	FileName    string
	UploadedUrl string
}
