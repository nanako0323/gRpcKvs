package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Dog ... dynamoDb struct
type Dog struct {
	Name string
	Kind string
}

func main() {

	svc := openDynamoDb()

	ume := Dog{Name: "Ume", Kind: "Mix"}
	selctUme := selectItem(ume)

	result, err := selctUme(svc)

	if err != nil {
		fmt.Println("GetItem Error", err)
		return
	}

	item := formatToString(result)
	fmt.Println(item)
}

func openDynamoDb() *dynamodb.DynamoDB {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	return svc
}

func selectItem(param Dog) func(svc *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error) {

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

func formatToString(result *dynamodb.GetItemOutput) string {

	dog := &Dog{}

	if err := dynamodbattribute.UnmarshalMap(result.Item, dog); err != nil {
		fmt.Println("Unmarshal Error", err)
		return ""
	}

	j, _ := json.Marshal(dog)
	return string(j)

}
