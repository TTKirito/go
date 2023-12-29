package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TTKirito/go/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	c := pb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	blog := &pb.Blog{
		AuthorId: "stephane",
		Title:    "test",
		Content:  "test",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{Blog: blog})

	if err != nil {
		log.Fatalf("unexpected error : %v", err)
	}

	fmt.Printf("Blog has been created: %v", createBlogRes)
}
