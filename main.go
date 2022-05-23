package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// S3の接続情報
	accessKey := os.Getenv("AWS_ACCSESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	region := aws.String(os.Getenv("DEFAULT_REGION"))

	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	session := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      region,
	}))
	svc := s3.New(session)
	c, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("バケット名"),
		Key:    aws.String("/diectory/filename"),
	})

	url, err := c.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return
	}

	fmt.Println(url)
}
