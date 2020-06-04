package delete

import (
	"fmt"
	"gogRpcKvs/kvs/models"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Dog ... delete item from key
func Dog(name string, kind string) bool {

	// open dynamodb session
	svc := utils.OpenDynamoDb()

	key := models.Dog{Name: name, Kind: kind}

	//create delete param
	input := setDeleteParam(key)

	//execute
	_, err := svc.DeleteItem(input)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func setDeleteParam(dog models.Dog) *dynamodb.DeleteItemInput {
	return &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(dog.Name),
			},
			"Kind": {
				S: aws.String(dog.Kind),
			},
		},
		TableName: aws.String(models.TableName),
	}
}
