package IOFile

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/bzdanny/BaiZe/baize/baizeSet"
	"github.com/bzdanny/BaiZe/baize/constants"
	"github.com/bzdanny/BaiZe/baize/setting"
	"io"
)

const (
	awsS3     = "s3"
	yiDong    = "eos"
	localhost = "localhost"
)

var FileType = baizeSet.Set[string]{}

type IOFile interface {
	PublicUploadFile(key string, data io.Reader) (string, error)
	privateUploadFile(key string, data io.Reader) (string, error)
}

var ioFile IOFile

func GetConfig() IOFile {
	return ioFile
}

func Init() {
	FileType.Add(awsS3)
	FileType.Add(yiDong)

	switch setting.Conf.UploadFile.Type {
	case awsS3:
		config := aws.Config{
			Credentials: credentials.NewStaticCredentialsProvider(setting.Conf.UploadFile.S3.accessKeyId, setting.Conf.UploadFile.S3.SecretAccessKey, ""),
			Region:      setting.Conf.UploadFile.S3.Region,
		}
		s := new(s3IOFile)
		s.s3Config = s3.NewFromConfig(config)
		s.bucket = setting.Conf.UploadFile.S3.BucketName
		s.domainName = setting.Conf.UploadFile.DomainName
	case yiDong:
		config := aws.Config{
			Credentials: credentials.NewStaticCredentialsProvider(setting.Conf.UploadFile.Eos.accessKeyId, setting.Conf.UploadFile.Eos.SecretAccessKey, ""),
			EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID: "aws",
					URL:         "https://eos-wuxi-1.cmecloud.cn",
				}, nil
			}),
		}
		s := new(s3IOFile)
		s.s3Config = s3.NewFromConfig(config)
		s.bucket = setting.Conf.UploadFile.Eos.BucketName
		s.domainName = setting.Conf.UploadFile.DomainName
		ioFile = s
	default:
		l := new(localHostIOFile)
		l.domainName = setting.Conf.UploadFile.DomainName
		pubPath := setting.Conf.UploadFile.Localhost.PublicResourcePrefix

		if pubPath == "" {
			pubPath = constants.DefaultPublicPath
		}
		l.publicPath = pubPath
		priPath := setting.Conf.UploadFile.Localhost.PrivateResourcePrefix
		if priPath == "" {
			priPath = constants.DefaultPublicPath
		}
		l.privatePath = priPath
		ioFile = l
	}
	return
}
