package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func main() {
	connectToAWS()

	result, err := svc.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}