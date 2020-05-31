package query

import (
	"fmt"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dog struct {
	Name string
	Kind string
}

//FindOne ... Get Single Item from dynamoDb
func FindOne(name string) (string, string) {

	svc := utils.OpenDynamoDb()

	ume := dog{Name: "Ume", Kind: "Mix"}
	selctUme := selectItem(ume)

	result, err := selctUme(svc)

	if err != nil {
		fmt.Println("GetItem Error", err)
		return "", ""
	}

	item := formatToDog(result)
	return item.Name, item.Kind
}

func selectItem(param dog) func(svc *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error) {

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

func formatToDog(result *dynamodb.GetItemOutput) dog {

	one := &dog{}

	if err := dynamodbattribute.UnmarshalMap(result.Item, one); err != nil {
		fmt.Println("Unmarshal Error", err)
		return dog{Name: "", Kind: ""}
	}

	return *one

}
