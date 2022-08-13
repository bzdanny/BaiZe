package IOFile

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type s3IOFile struct {
	s3Config   *s3.Client
	bucket     string
	domainName string
}

func (s *s3IOFile) PublicUploadFile(file *fileParams) (string, error) {
	obj := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(file.keyName),
		Body:        file.data,
		ContentType: aws.String(file.contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	}
	_, err := s.s3Config.PutObject(context.TODO(), obj)
	if err != nil {
		return "", err
	}
	return s.domainName + "/" + file.keyName, nil
}

func (s *s3IOFile) privateUploadFile(file *fileParams) (string, error) {
	obj := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(file.keyName),
		Body:        file.data,
		ContentType: aws.String(file.contentType),
		ACL:         types.ObjectCannedACLPrivate,
	}
	_, err := s.s3Config.PutObject(context.TODO(), obj)
	if err != nil {
		return "", err
	}
	return file.keyName, nil
}
