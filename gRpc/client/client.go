package main

import (
	"context"
	"fmt"
	pb "gogRpcKvs/gRpc"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())

	if err != nil {
		log.Fatal("client connection error:", err)
	}

	defer conn.Close()

	client := pb.NewDogClient(conn)
	message := &pb.AddMyDogMessage{Name: "Mako"}
	res, err := client.AddMyDog(context.TODO(), message)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error:%#v \n", err)
}
