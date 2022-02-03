package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func main() {

	filename := flag.String("f", "", "The file to upload")
	flag.Parse()

	if *filename == "" {
		fmt.Println("You must supply the file")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load the configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(*filename)

	if err != nil {
		fmt.Println("Unable to open the file" + *filename)
		return
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String("phanimullapudi"),
		Key:    filename,
		Body:   file,
	}

	_, err = PutFile(context.TODO(), client, input)

	if err != nil {
		fmt.Println("Got error uploading file:" + *filename)
		fmt.Println(err)
		return
	}

}
