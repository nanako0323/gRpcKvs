package service

import (
	"context"
	"errors"
	pb "gogRpcKvs/gRpc"
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
