package main

import (
	"fmt"
	"gogRpcKvs/kvs/models"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	result := delete("Mako", "Human")
	fmt.Println(result)
}

func delete(name string, kind string) bool {

	// open dynamodb session
	svc := utils.OpenDynamoDb()

	//create delete param
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(name),
			},
			"Kind": {
				S: aws.String(kind),
			},
		},
		TableName: aws.String(models.TableName),
	}

	//execute
	_, err := svc.DeleteItem(input)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
