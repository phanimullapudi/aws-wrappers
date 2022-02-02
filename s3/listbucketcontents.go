package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Load the Shared AWS configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2"))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("gg-wordpress"),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("First page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s   size=%d   lastmodified=%v", aws.ToString(object.Key), object.Size, object.LastModified)
	}
}
