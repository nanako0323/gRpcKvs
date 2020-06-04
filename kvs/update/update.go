package update

import (
	"fmt"
	"gogRpcKvs/kvs/models"
	"gogRpcKvs/kvs/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//DogRemark ... update remark from key
func DogRemark(name string, kind string, remark string) bool {
	//begin session
	svc := utils.OpenDynamoDb()

	key := models.Dog{Name: name, Kind: kind}

	//set param
	input := setUpdateParam(key, remark)

	//update
	_, err := svc.UpdateItem(input)

	//error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func setUpdateParam(dog models.Dog, remark string) *dynamodb.UpdateItemInput {

	return &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(remark),
			},
		},
		TableName: aws.String(models.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(dog.Name),
			},
			"Kind": {
				S: aws.String(dog.Kind),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Remark = :r"),
	}

}
