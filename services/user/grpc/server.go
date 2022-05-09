package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/knwoop/microsercices-example/gen/proto/user/v1"
)

var _ pb.UserServiceServer = (*server)(nil)

type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser implements UserServer
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	if req.Id != "gig" {
		return nil, status.Error(codes.NotFound, "not found user")
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:   req.Id,
			Name: "showcase gig",
		},
	}, nil
}

func New() (*grpc.Server, error) {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	reflection.Register(s)

	return s, nil
}
