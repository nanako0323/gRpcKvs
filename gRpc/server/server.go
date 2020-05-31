package main

import (
	pb "gogRpcKvs/gRpc"
	"gogRpcKvs/gRpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")

	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	dogService := &service.MyDogService{}

	pb.RegisterDogServer(server, dogService)
	server.Serve(listenPort)
}
