package IOFile

import (
	"bytes"
	"io"
	"io/ioutil"
)

type localHostIOFile struct {
	publicPath  string
	privatePath string
	domainName  string
}

func (l *localHostIOFile) PublicUploadFile(key string, data io.Reader) (string, error) {

	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(data)
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	err = ioutil.WriteFile(l.publicPath+key, b, 0664)
	if err != nil {
		return "", err
	}
	return l.domainName + key, nil
}

func (l *localHostIOFile) privateUploadFile(key string, data io.Reader) (string, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(data)
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	err = ioutil.WriteFile(l.privatePath+key, b, 0664)
	if err != nil {
		return "", err
	}
	return key, nil
}
