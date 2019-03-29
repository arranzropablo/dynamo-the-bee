package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func connectToAWS() {
	fmt.Println("Introduce your Access key ID:")
	accessId := readString()
	fmt.Println("Introduce your Secret access key:")
	secret := readString()

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(accessId, secret, ""),
	})

}