package fileUploadUtils

import (
	"baize/app/setting"
	"baize/app/utils/dateUtils"
	"baize/app/utils/pathUtils"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Upload(baseDir string, file *multipart.FileHeader) string {
	pathFileName := baseDir + ExtractFilename(file)
	savePath := setting.Conf.Profile + pathFileName
	dir := filepath.Dir(savePath)
	pathUtils.CreateMutiDir(dir)
	src, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer src.Close()
	out, err := os.Create(savePath)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		panic(err)
	}
	return pathFileName
}

func ExtractFilename(file *multipart.FileHeader) string {

	return dateUtils.DatePath() + "/" + uuid.NewV4().String() + "." + getExtension(file)

}

func getExtension(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		return prefixGainExtension(file.Header["Content-Type"][0])
	}
	return ext
}

func prefixGainExtension(prefix string) string {
	switch prefix {
	case "image/png":
		return "png"
	case "image/jpg":
		return "jpg"
	case "image/jpeg":
		return "jpeg"
	case "image/bmp":
		return "bmp"
	case "image/gif":
		return "gif"
	default:
		return ""
	}
}
