package main

import (
	"context"
	"fmt"
	"io/ioutil"
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

func uploadfiles(foldername string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load the configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	files, err := ioutil.ReadDir("/Users/phanimullapudi/Documents/test/" + foldername)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			fmt.Println("Need code for recurssive")
		} else {

			file, err := os.Open("/Users/phanimullapudi/Documents/test/" + foldername + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			input := &s3.PutObjectInput{
				Bucket: aws.String("phanimullapudi"),
				Key:    aws.String(file.Name()),
				Body:   file,
			}

			_, err = PutFile(context.TODO(), client, input)

			if err != nil {
				fmt.Println("Got error uploading file:" + file.Name())
				fmt.Println(err)
				return
			}
		}
	}

}

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Failed to load the configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	files, err := ioutil.ReadDir("/Users/phanimullapudi/Documents/test")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			go uploadfiles(file.Name())
		} else {
			file, err := os.Open("/Users/phanimullapudi/Documents/test/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			input := &s3.PutObjectInput{
				Bucket: aws.String("phanimullapudi"),
				Key:    aws.String(file.Name()),
				Body:   file,
			}

			_, err = PutFile(context.TODO(), client, input)

			if err != nil {
				fmt.Println("Got error uploading file:" + file.Name())
				fmt.Println(err)
				return
			}

		}

	}

	var input string
	fmt.Scanln(&input)
}
