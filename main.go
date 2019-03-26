package main

import (
	"bytes"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nosajio/markdown-to-json/mdtojson"
	"os"
)

// ProcessorEvent would represent the incoming lambda event, but we only require
// a signal for the time being, hence the empty struct
type ProcessorEvent struct {
}

// HandleEvent is run when the lambda fn is triggered. It will download and
// process an entire repo and then save the resulting JSON to the specified S3
// bucket
func HandleEvent(e ProcessorEvent) error {
	repoURL := os.Getenv("REPO")
	tmpDIR := os.Getenv("DIR")
	s3Region := os.Getenv("S3_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")
	s3Key := os.Getenv("S3_KEY")

	// Initiate the S3 session
	s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if err != nil {
		return err
	}

	// Download and process posts
	json, err := mdtojson.ProcessRepo(repoURL, tmpDIR)
	if err != nil {
		return err
	}
	contentLen := int64(len(json))

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s3Bucket),
		Key:           aws.String(s3Key),
		Body:          bytes.NewReader([]byte(json)),
		ContentLength: aws.Int64(contentLen),
		ContentType:   aws.String("application/json"),
	})

	return nil
}

func main() {
	lambda.Start(HandleEvent)
}
