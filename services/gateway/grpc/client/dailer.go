package client

import (
	"context"

	"google.golang.org/grpc"
)

func DialContext(ctx context.Context, serverAddr string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	conn, err := grpc.DialContext(ctx, serverAddr, opts)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
