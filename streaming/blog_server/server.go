package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/TTKirito/go/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// fix server golang

type server struct {
	pb.UnimplementedBlogServiceServer
}

var collection *mongo.Collection

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	fmt.Println("Create blog request::::")
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot convert to OID"),
		)
	}

	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *pb.ReadBlogRequest) (*pb.ReadBlogResponse, error) {
	blogId := req.GetBlogId()
	fmt.Println("Read blog request::::")
	oid, err := primitive.ObjectIDFromHex(blogId)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID "))

	}

	data := &blogItem{}
	filter := bson.D{{Key: "_id", Value: oid}}
	fmt.Println(oid, filter)
	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find blog with specified ID: %v", err))
	}

	return &pb.ReadBlogResponse{
		Blog: &pb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil
}

func (*server) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogResponse, error) {
	fmt.Println("Update blog request::::")
	blog := req.GetBlog()

	oid, err := primitive.ObjectIDFromHex(blog.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID "))

	}

	data := &blogItem{}
	filter := bson.D{{Key: "_id", Value: oid}}
	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find blog with specified ID: %v", err))
	}

	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	updateRes, updateErr := collection.ReplaceOne(context.Background(), filter, data)

	if updateErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr))

	}
	return &pb.UpdateBlogResponse{
		Blog: &pb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil
}

func main() {
	fmt.Println("Blog Service started")

	// connect mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://thuanton98:thuan123@cluster0.igcyw.mongodb.net/"))

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("myDB").Collection("blog")

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
