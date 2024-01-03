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

	// _, err2 := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{BlogId: "adfadsf"})
	// if err2 != nil {
	// 	log.Fatalf("Error happened while reading: %v ", err)
	// }

	blogId := createBlogRes.Blog.GetId()
	readBlogRes, _ := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{BlogId: blogId})

	fmt.Printf("Blog was read: %v", readBlogRes)

	newBlog := &pb.Blog{
		Id:       blogId,
		AuthorId: "Change author",
		Title:    "First",
		Content:  "Content",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &pb.UpdateBlogRequest{Blog: newBlog})

	if updateErr != nil {
		fmt.Printf("Error happened whie update: %v", err)
	}

	fmt.Printf("blog was update %v", updateRes)

	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{BlogId: blogId})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", updateErr)
	}

	fmt.Printf("blog was delete %v", deleteRes)

}
