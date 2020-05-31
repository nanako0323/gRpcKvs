package service

import (
	"context"
	"errors"
	pb "gogRpcKvs/gRpc"
	"gogRpcKvs/kvs/create"
	"gogRpcKvs/kvs/query"
)

// MyDogService ...struct
type MyDogService struct {
}

// GetMyDog ... get message , return response
func (s *MyDogService) GetMyDog(ctx context.Context, message *pb.GetMyDogMessage) (*pb.GetMyDogResponse, error) {

	name, kind := query.FindOne(message.TargetDog)

	if name == "" || kind == "" {
		return nil, errors.New("Not Found YourDog")
	}

	return &pb.GetMyDogResponse{
		Name: name,
		Kind: kind,
	}, nil
}

// AddMyDog ... get message, and put item to dynamodb
func (s *MyDogService) AddMyDog(ctx context.Context, message *pb.AddMyDogMessage) (*pb.AddMyDogResponse, error) {

	success := create.InsertOne(message.Name, message.Kind)

	if !success {
		return nil, errors.New("Failed Put Item")
	}

	return &pb.AddMyDogResponse{
		Result: "success",
	}, nil
}
