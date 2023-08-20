package main

import (
	"context"
	"fmt"
	pb "hello-server/service"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	// open port
	listen, _ := net.Listen("tcp", ":9090")
	// create grpc service
	grpcServer := grpc.NewServer()
	// register service
	pb.RegisterSayHelloServer(grpcServer, &server{})
	// run service
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("fail to serve: %v", err)
		return
	}
}
