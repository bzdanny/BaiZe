package stringUtils

import (
	"github.com/gogf/gf/v2/util/gconv"
	"mime/multipart"
	"path"
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

func GetTenantRandomName(userId int64, fileName string) string {
	t := time.Now()
	nameKey := "/" + gconv.String(userId) + "/" + t.Format("06/01/02") + "/" + GetUUID() + path.Ext(fileName)
	return nameKey
}
func GetTenantOriginalName(userId int64, fileName string) string {
	if fileName == "" {
		return GetTenantRandomName(userId, fileName)
	}
	t := time.Now()
	nameKey := "/" + gconv.String(userId) + "/" + t.Format("06/01/02") + "/" + GetUUID() + "/" + fileName
	return nameKey
}

func getExtension(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		return fileExtension[file.Header["Content-Type"][0]]
	}
	return ext
}
