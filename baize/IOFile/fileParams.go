package IOFile

import (
	"io"
	"mime/multipart"
	"path/filepath"
)

type fileParams struct {
	keyName     string
	contentType string
	data        io.Reader
}

var fileExtension = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
	"image/bmp":  "bmp",
	"image/gif":  "gif",
}

var contentType = map[string]string{
	".json": "application/json",
	".html": "text/html",
	".js":   "application/javascript",
	".css":  "text/css",
	".gif":  "image/gif",
	".png":  "image/png",
	".gz":   "application/x-gzip",
	".svg":  "image/svg+xml",
	".pdf":  "application/pdf",
	".jpeg": "image/jpeg",
}

func NewFileParamsRandomName(keyName string, file multipart.File) *fileParams {

	f := new(fileParams)
	f.keyName = keyName
	f.data = file
	f.contentType = contentType[filepath.Ext(keyName)]
	return f
}
