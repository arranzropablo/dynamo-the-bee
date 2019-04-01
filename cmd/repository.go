package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

var svc *dynamodb.DynamoDB

func connectToAWS() {
	fmt.Println("Introduce your Access key ID:")
	accessId := readString()
	fmt.Println("Introduce your Secret access key:")
	secret := readString()
	fmt.Println("Introduce your token:")
	token := readString()

	awsSession := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("dynamodb.us-east-1.amazonaws.com"),
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(accessId, secret, token),
	}))

	svc = dynamodb.New(awsSession)
}

func ListTables() (result *dynamodb.ListTablesOutput) {
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func CreateTable(name string, fields map[string]string, hashfield string, rangefield string, readCapacity int64, writeCapacity int64) (result *dynamodb.CreateTableOutput) {
	for _, t := range fields {
		switch t {
		case "N", "S", "BOOL", "B", "SS", "NS", "BS":
			continue
		default:
			log.Fatalf("Type %s is not accepted\n", t)
		}
	}

	var attributes []*dynamodb.AttributeDefinition

	for f, t := range fields {
		attributes = append(attributes, &dynamodb.AttributeDefinition{
			AttributeName: aws.String(f),
			AttributeType: aws.String(t),
		})
	}

	if _, ok := fields[hashfield]; !ok {
		log.Fatalf("Hash field can not be %s\n", hashfield)
	}

	keys := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String(hashfield),
			KeyType:       aws.String("HASH"),
		},
	}

	if _, ok := fields[rangefield]; !ok && len(rangefield) > 0 {
		log.Fatalf("Range field can not be %s\n", rangefield)
	} else if ok {
		keys = append(keys, &dynamodb.KeySchemaElement{
			AttributeName: aws.String(rangefield),
			KeyType:       aws.String("RANGE"),
		})
	}

	//default values
	if readCapacity == 0  {
		readCapacity = 10
	}

	if writeCapacity == 0  {
		writeCapacity = 10
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: attributes,
		KeySchema: keys,
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readCapacity),
			WriteCapacityUnits: aws.Int64(writeCapacity),
		},
		TableName: aws.String(name),
	}
	result, err := svc.CreateTable(input)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func GetItemsFromTable(tableName string) {
	//como puedo hacer un get de un objeto abstracto...?
	//https://github.com/guregu/dynamo
}