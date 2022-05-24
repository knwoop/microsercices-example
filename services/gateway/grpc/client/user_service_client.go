package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	userv1 "github.com/knwoop/microsercices-example/gen/proto/user/v1"
)

type UserServiceClient struct {
	client userv1.UserServiceClient
	conn   *grpc.ClientConn
}

func NewUserServiceClient(ctx context.Context, serverAddr string) (*UserServiceClient, error) {
	conn, err := DialContext(ctx, serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed DialContext: %w", err)
	}
	return &UserServiceClient{
		client: userv1.NewUserServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *UserServiceClient) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	res, err := c.client.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *UserServiceClient) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}
