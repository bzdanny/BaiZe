package IOFile

import (
	"bytes"
	"github.com/bzdanny/BaiZe/baize/constants"
	"github.com/bzdanny/BaiZe/baize/utils/pathUtils"
	"io/ioutil"
	"path/filepath"
)

type localHostIOFile struct {
	publicPath  string
	privatePath string
	domainName  string
}

func (l *localHostIOFile) PublicUploadFile(file *fileParams) (string, error) {

	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(file.data)
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	pathAndName := l.publicPath + file.keyName
	err = pathUtils.CreateMutiDir(filepath.Dir(pathAndName))
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(pathAndName, b, 0664)
	if err != nil {
		return "", err
	}
	return l.domainName + constants.ResourcePrefix + "/" + file.keyName, nil
}

func (l *localHostIOFile) privateUploadFile(file *fileParams) (string, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(file.data)
	if err != nil {
		return "", err
	}
	pathAndName := l.privatePath + file.keyName
	err = pathUtils.CreateMutiDir(filepath.Dir(pathAndName))
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	err = ioutil.WriteFile(pathAndName, b, 0664)
	if err != nil {
		return "", err
	}
	return file.keyName, nil
}
