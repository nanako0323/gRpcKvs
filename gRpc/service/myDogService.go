package service

import (
	"context"
	"errors"
	pb "gogRpcKvs/gRpc"
)

// MyDogService ...struct
type MyDogService struct {
}

// GetMyDog ... get message , return response
func (s *MyDogService) GetMyDog(ctx context.Context, message *pb.GetMyDogMessage) (*pb.GetMyDogResponse, error) {

	switch message.TargetDog {
	case "ume":
		return &pb.GetMyDogResponse{
			Name: "ume",
			Kind: "mix",
		}, nil

	case "kojiro":
		return &pb.GetMyDogResponse{
			Name: "kojiro",
			Kind: "Labrador Retriever",
		}, nil
	}

	return nil, errors.New("Not Found YourDog")
}
