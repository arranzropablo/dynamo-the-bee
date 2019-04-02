package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB

func connectToAWS() {

	awsSession := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("dynamodb.us-east-1.amazonaws.com"),
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", ""),
	}))

	db = dynamo.New(awsSession)
}