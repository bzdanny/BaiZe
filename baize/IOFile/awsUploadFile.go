package IOFile

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/bzdanny/BaiZe/baize/setting"
	"io"
)

type s3IOFile struct {
	s3Config   *s3.Client
	bucket     string
	domainName string
}

func (s *s3IOFile) PublicUploadFile(key string, data io.Reader) (string, error) {
	obj := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   data,
		ACL:    types.ObjectCannedACLPublicRead,
	}
	_, err := s.s3Config.PutObject(context.TODO(), obj)
	if err != nil {
		return "", err
	}
	return s.domainName + key, nil
}

func (s *s3IOFile) privateUploadFile(key string, data io.Reader) (string, error) {
	obj := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   data,
		ACL:    types.ObjectCannedACLPrivate,
	}
	_, err := s.s3Config.PutObject(context.TODO(), obj)
	if err != nil {
		return "", err
	}
	return setting.Conf.UploadFile.DomainName + key, nil
}
