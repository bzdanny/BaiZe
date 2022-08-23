package stringUtils

import (
	"github.com/gogf/gf/v2/util/gconv"
	"mime/multipart"
	"path/filepath"
	"time"
)

var fileExtension = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
	"image/bmp":  "bmp",
	"image/gif":  "gif",
}

func GetTenantRandomName(userId int64, extensionName string) string {
	t := time.Now()
	nameKey := gconv.String(userId) + "/" + t.Format("06/01/02") + "/" + GetUUID() + "." + extensionName
	return nameKey
}

func GetExtension(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		return fileExtension[file.Header["Content-Type"][0]]
	}
	return ext
}
