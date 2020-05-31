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

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Dog"),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String("Ume"),
			},
			"Kind": {
				S: aws.String("Mix"),
			},
		},
	}

	result, err := svc.GetItem(input)

	if err != nil {
		fmt.Println("GetItem Error", err)
		return
	}

	fmt.Println(result)

	dog := &Dog{}

	//musicにmapし、エラーが発生した場合
	if err := dynamodbattribute.UnmarshalMap(result.Item, dog); err != nil {
		fmt.Println("Unmarshal Error", err)
		return
	}

	j, _ := json.Marshal(dog)
	fmt.Println(string(j))
}
