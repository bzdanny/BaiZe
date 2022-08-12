package IOFile

import (
	"bytes"
	"io/ioutil"
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
	err = ioutil.WriteFile(l.publicPath+file.keyName, b, 0664)
	if err != nil {
		return "", err
	}
	return l.domainName + file.keyName, nil
}

func (l *localHostIOFile) privateUploadFile(file *fileParams) (string, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(file.data)
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	err = ioutil.WriteFile(l.privatePath+file.keyName, b, 0664)
	if err != nil {
		return "", err
	}
	return file.keyName, nil
}
