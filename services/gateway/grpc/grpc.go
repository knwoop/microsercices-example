package grpc

import (
	"fmt"

	userpb "github.com/knwoop/microsercices-example/gen/proto/user/v1"
	"google.golang.org/grpc"
)

type GRPCServicesClient struct {
	User userpb.UserServiceServer
}

func NewGRPCServicesClient() (*GRPCServicesClient, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}
	userConn, err := grpc.DialContext(ctx, "localhost:50052", opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial to user grpc server: %w", err)
	}

	if err := userpb.RegisterUserServiceServer(ctx, mux, greetingpb.NewGreeterClient(greetingConn)); err != nil {
		return nil, fmt.Errorf("failed to create a catalog grpc client: %w", err)
	}

	return &GRPCServicesClient{}, nil
}
