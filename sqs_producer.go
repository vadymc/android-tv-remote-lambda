package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	SQS_QUEUE = "tv_remote"
)

var (
	sqsClient *sqs.SQS
	queueUrl  *string
)

func startSqs() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)
	if err != nil {
		panic(err)
	}
	sqsClient = sqs.New(sess)
	initQueueUrl()
	log.Println("Connected to SQS")
}

func initQueueUrl() {
	result, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(SQS_QUEUE),
	})

	if err != nil {
		log.Fatalf("Error", err)
		return
	}
	queueUrl = result.QueueUrl
}

func sendToSqs(command string) {
	result, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(command),
		QueueUrl:    queueUrl,
	})

	if err != nil {
		log.Println("Error", err)
		return
	}

	log.Println("Success", *result.MessageId)
}
