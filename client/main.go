package main

import (
	"context"
	"fmt"
	constants "github.com/heavybr/go_grpc_example"
	"github.com/heavybr/go_grpc_example/proto/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial(fmt.Sprintf("localhost:%d", constants.PORT), grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error on gRPC server dial: %v", err)
	}

	defer connection.Close()

	doSampleRequest(connection, err)
}

func doSampleRequest(connection *grpc.ClientConn, err error) {
	client := pb.NewHelloServiceClient(connection)

	request := &pb.HelloRequest{
		Name: "Matheus Cumpian",
	}

	res, err := client.Hello(context.Background(), request)

	if err != nil {
		log.Printf("error during request to hello server: %v", err)
	}

	log.Println(res)
}
