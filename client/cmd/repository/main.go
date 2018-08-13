package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"bytes"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"path/filepath"
	"strings"
	"fmt"
)

const (
	Bucket    = "releases.kowala.io"
	Region    = "us-east-1"
	IndexName = "index.txt"
)

func main() {
	config := &aws.Config{Region: aws.String(Region)}
	sess := session.Must(session.NewSession(config))

	svc := s3.New(sess)
	uploader := s3manager.NewUploader(sess)

	filenames := listS3Objects(svc, Bucket)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String(filepath.Base(IndexName)),
		ACL:    aws.String("public-read"),
		Body:   strings.NewReader(filenames.String()),
	})
	if err != nil {
		log.Fatal("failed to list objects", err)
	}

	fmt.Printf("file %s updated n %s", IndexName, Bucket)
}

func listS3Objects(svc *s3.S3, bucket string) bytes.Buffer {
	var buffer bytes.Buffer

	listObjectsInput := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	err := svc.ListObjectsPages(listObjectsInput, func(p *s3.ListObjectsOutput, _ bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			_, err := buffer.WriteString(*obj.Key + "\n")
			if err != nil {
				log.Fatal("failed to list objects", err)
			}
		}
		return true
	})
	if err != nil {
		log.Fatal("failed to list objects", err)
	}
	return buffer
}
