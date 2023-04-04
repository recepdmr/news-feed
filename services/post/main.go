package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/recepdmr/news-feed/services/post/domain"
	"github.com/recepdmr/news-feed/services/post/mongo"
	pb "github.com/recepdmr/news-feed/services/post/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

type server struct {
	pb.UnimplementedPostServiceServer
}

func (s *server) Get(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	repository := mongo.NewMongoDBPostRepository()

	posts := repository.GetAll()

	responsePosts := make([]*pb.Post, len(posts))

	for i, p := range posts {
		responsePosts[i] = &pb.Post{
			Id:        p.ID.Hex(),
			Content:   p.Content,
			CreatedAt: p.CreatedAt.String(),
		}
	}

	return &pb.GetPostsResponse{
		Data: responsePosts,
	}, nil
}

func (s *server) GetById(ctx context.Context, in *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	repository := mongo.NewMongoDBPostRepository()
	post, err := repository.GetById(in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetByIdResponse{
		Data: &pb.Post{
			Id:        post.ID.Hex(),
			Content:   post.Content,
			CreatedAt: post.CreatedAt.String(),
		},
	}, nil
}

func (s *server) Insert(ctx context.Context, in *pb.InsertPostRequest) (*pb.InsertPostResponse, error) {
	repository := mongo.NewMongoDBPostRepository()

	return &pb.InsertPostResponse{}, repository.Insert(domain.Post{
		Content: in.GetContent(),
	})
}

func (s *server) Delete(ctx context.Context, in *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	repo := mongo.NewMongoDBPostRepository()

	postId, err := primitive.ObjectIDFromHex(in.GetId())
	if err != nil {
		return nil, err
	}

	repo.Delete(domain.Post{
		ID: postId,
	})
	return &pb.DeletePostResponse{}, nil
}

func init() {
	port = os.Getenv("PORT")

	if port == "" {
		panic("PORT is waiting")
	}

	if os.Getenv("MONGODB_URI") == "" {
		panic("MONGODB_URI is waiting")
	}
}

func main() {
	log.Printf("Post service is available at %v", port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterPostServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
