package api

import (
	"context"
	"fmt"

	pb "github.com/knwoop/microsercices-example/gen/proto/gateway/v1"
	userapi "github.com/knwoop/microsercices-example/gen/proto/user/v1"
)

var _ pb.GreeterSer3viceServer = (*server)(nil)

type server struct {
	pb.UnimplementedGreeterServiceServer

	userAPI userapi.UserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	apiReq := &userapi.GetUserRequest{
		Id: req.Id,
	}

	res, err := s.userAPI.GetUser(ctx, apiReq)
	if err != nil {
		return nil, fmt.Errorf("failed userapi get user: %w", err)
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:   res.User.Id,
			Name: res.User.Name,
		},
	}, nil
}
