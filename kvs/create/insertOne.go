package create

import (
	"fmt"
	"gogRpcKvs/kvs/models"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// InsertOne ... insert one data to dynamoDb.
// If succeed, return true
func InsertOne(name string, kind string) bool {
	svc := utils.OpenDynamoDb()

	data := models.Dog{Name: name, Kind: kind}

	av, isErr := mapToAttributeValue(data)

	if isErr {
		return false
	}

	isSuccess := putItem(av, svc)

	return isSuccess
}

func exitWithError(err error, msg string) {
	fmt.Println(msg)
	fmt.Println(err.Error())
}

func mapToAttributeValue(data models.Dog) (map[string]*dynamodb.AttributeValue, bool) {
	av, err := dynamodbattribute.MarshalMap(data)

	if err != nil {
		exitWithError(err, "Got error marshalling new item:")
		return nil, true
	}

	return av, false
}

func putItem(av map[string]*dynamodb.AttributeValue, svc *dynamodb.DynamoDB) bool {
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(models.TableName),
	}

	_, puterr := svc.PutItem(input)

	if puterr != nil {
		exitWithError(puterr, "Got error calling PutItem:")
		return false
	}

	return true
}
