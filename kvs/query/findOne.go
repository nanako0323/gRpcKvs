package query

import (
	"fmt"
	"gogRpcKvs/kvs/models"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//FindOne ... Get Single Item from dynamoDb
func FindOne(name string) (string, string) {

	svc := utils.OpenDynamoDb()

	ume := models.Dog{Name: "Ume", Kind: "Mix"}
	selctUme := selectItem(ume)

	result, err := selctUme(svc)

	if err != nil {
		fmt.Println("GetItem Error", err)
		return "", ""
	}

	item := formatToDog(result)
	return item.Name, item.Kind
}

func selectItem(param models.Dog) func(svc *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error) {

	return func(svc *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error) {

		input := &dynamodb.GetItemInput{
			TableName: aws.String("Dog"),
			Key: map[string]*dynamodb.AttributeValue{
				"Name": {
					S: aws.String(param.Name),
				},
				"Kind": {
					S: aws.String(param.Kind),
				},
			},
		}

		return svc.GetItem(input)
	}
}

func formatToDog(result *dynamodb.GetItemOutput) models.Dog {

	one := &models.Dog{}

	if err := dynamodbattribute.UnmarshalMap(result.Item, one); err != nil {
		fmt.Println("Unmarshal Error", err)
		return models.Dog{Name: "", Kind: ""}
	}

	return *one

}
