package main

import (
	"context"
	"fmt"
	constants "github.com/heavybr/go_grpc_example"
	"github.com/heavybr/go_grpc_example/proto/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s server) Hello(_ context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("HelloRequest received")
	msg := fmt.Sprintf("hello %s", request.GetName())
	return &pb.HelloResponse{
		Msg: msg,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", constants.PORT))

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &server{})

	log.Printf("😀 gRPC server running on port %d", constants.PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}


}
