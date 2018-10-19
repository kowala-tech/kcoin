package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kowala-tech/kcoin/client/version"
	"log"
	"path/filepath"
	"strings"
)

const (
	Bucket    = "releases.kowala.tech"
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

	fmt.Printf("file %s updated in %s", IndexName, Bucket)
}

func listS3Objects(svc *s3.S3, bucket string) bytes.Buffer {
	var buffer bytes.Buffer

	listObjectsInput := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	err := svc.ListObjectsPages(listObjectsInput, func(p *s3.ListObjectsOutput, _ bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			if _, err := version.FilenameParser(*obj.Key); err != nil {
				// skip file, it's not parsable not a binary release file
				continue
			}
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
