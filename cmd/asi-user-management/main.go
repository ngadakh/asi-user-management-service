package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ngadakh/asi-user-management-service/generated/go/asi-user-management-service"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("CreateUser called with request: %v", in)

	return &pb.CreateUserResponse{
		Id:    "123",
		Name:  "Navnath Gadakh",
		Email: "navnathgadakh@gmail.com",
	}, nil
}

func main() {
	fmt.Println("Hello World!!")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
