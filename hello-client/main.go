package main

import (
	"context"
	"fmt"
	pb "hello-server/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	client := pb.NewSayHelloClient(conn)

	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: " Napleon"})

	fmt.Println(resp.GetResponseMsg())
}
