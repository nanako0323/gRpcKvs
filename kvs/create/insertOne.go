package main

import (
	"fmt"
	"gogRpcKvs/kvs/utils"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type item struct {
	Name string
	Kind string
}

const tableName string = "Dog"

func main() {
	InsertOne()
}

// InsertOne ... insert one data to dynamoDb
func InsertOne() {
	svc := utils.OpenDynamoDb()

	data := item{Name: "Tao", Kind: "Cat"}

	av := mapToAttributeValue(data)

	putItem(av, svc)
}

func exitWithError(err error, msg string) {
	fmt.Println(msg)
	fmt.Println(err.Error())
	os.Exit(1)
}

func mapToAttributeValue(data item) map[string]*dynamodb.AttributeValue {
	av, err := dynamodbattribute.MarshalMap(data)

	if err != nil {
		exitWithError(err, "Got error marshalling new item:")
	}

	return av
}

func putItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB) {
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, puterr := svc.PutItem(input)

	if puterr != nil {
		exitWithError(puterr, "Got error calling PutItem:")
	}
}
