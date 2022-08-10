package IOFile

type UploadFile struct {
	Type       string `mapstructure:"type"`
	DomainName string `mapstructure:"domain_name"`
	*S3        `mapstructure:"s3"`
	*Eos       `mapstructure:"eos"`
	*Localhost `mapstructure:"localhost"`
}
type S3 struct {
	accessKeyId     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"Secret_access_key"`
	Region          string `mapstructure:"region"`
	BucketName      string `mapstructure:"bucket_name"`
}
type Eos struct {
	accessKeyId     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"Secret_access_key"`
	Url             string `mapstructure:"url"`
	BucketName      string `mapstructure:"bucket_name"`
}
type Localhost struct {
	PublicResourcePrefix  string `mapstructure:"public_resource_prefix"`
	PrivateResourcePrefix string `mapstructure:"private_resource_prefix"`
}
