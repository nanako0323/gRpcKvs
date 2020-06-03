package main

import (
	"fmt"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	result := Update()
	fmt.Println(result)
}

// Update ... get key param and update param
// return bool
func Update() bool {
	//begin session
	svc := utils.OpenDynamoDb()

	tableName := "Dog"
	dogName := "Kojiro"
	dogKind := "Labrador Retriever"

	//set param
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String("He is dead."),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(dogName),
			},
			"Kind": {
				S: aws.String(dogKind),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Remark = :r"),
	}

	//update
	_, err := svc.UpdateItem(input)

	//error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
