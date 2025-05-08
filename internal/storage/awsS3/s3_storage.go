package awsS3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	client *s3.S3
}

func (s *S3Storage) InitSession() {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"), ""),
		Region: aws.String("us-east-2"),
	})
	if err != nil {
		panic(err)
	}
	s.client = s3.New(sess)
}

func (s *S3Storage) UploadFile(fileData []byte, fileName, contentType string) {

}
