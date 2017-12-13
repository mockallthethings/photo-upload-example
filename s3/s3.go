package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader = s3manager.Uploader
type UploadInput = s3manager.UploadInput
type UploadOutput = s3manager.UploadOutput

type CreateBucketInput = s3.CreateBucketInput
type CreateBucketOutput = s3.CreateBucketOutput

type client struct {
	svc      *s3.S3
	uploader *s3manager.Uploader
}

func (c *client) CreateBucket(input *CreateBucketInput) (*CreateBucketOutput, error) {
	return c.svc.CreateBucket(input)
}

func (c *client) Upload(ui *UploadInput) (*UploadOutput, error) {
	return c.uploader.Upload(ui)
}

type Client interface {
	Upload(ui *UploadInput) (*UploadOutput, error)
	CreateBucket(input *CreateBucketInput) (*CreateBucketOutput, error)
}

const S3Endpoint = "http://localstack:4572"

func NewClient() (Client, error) {
	cfg := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
		S3ForcePathStyle: aws.Bool(true),
		// same default region as localstackj
		Region:   aws.String(endpoints.UsEast1RegionID),
		Endpoint: aws.String(S3Endpoint),
	}
	sess := session.Must(session.NewSession(cfg))
	svc := s3.New(sess)
	defaultBucket := "mockallthethings-example"
	_, err := svc.CreateBucket(&CreateBucketInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(defaultBucket),
	})
	if err != nil {
		return nil, err
	}
	return &client{
		svc:      svc,
		uploader: s3manager.NewUploader(sess),
	}, nil
}
