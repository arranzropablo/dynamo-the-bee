package main

import (
	"gopkg.in/underarmour/dynago.v2"
)

var client *dynago.Client

func connectToAWS() {

	client = dynago.NewClient(dynago.NewAwsExecutor(dynago.ExecutorConfig{
		Region:    "us-east-1",
		AccessKey: "",
		SecretKey: "",
		SessionToken: "",
	}))

}