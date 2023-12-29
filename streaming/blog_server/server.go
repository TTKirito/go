package main

import (
	"fmt"
	"log"
	"net"

	"github.com/TTKirito/go/pb"
	"google.golang.org/grpc"
)

// fix server golang

type server struct {
	pb.UnimplementedBlogServiceServer
}

func main() {
	fmt.Println("Blog Service started")

	// connect mongodb

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen :%v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	pb.RegisterBlogServiceServer(s, &server{})

	fmt.Println("Starting Server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
