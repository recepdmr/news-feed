package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/recepdmr/news-feed/services/post/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedPostServiceServer
}

func (s *server) GetPosts(ctx context.Context, in *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	return &pb.GetPostsResponse{
		Data: []*pb.Post{
			{
				Id:        "recep",
				Content:   "test content",
				CreatedAt: time.Now().UTC().String(),
				DeletedAt: "",
			},
		},
	}, nil
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Panic("PORT is waiting")
	}

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
