package api

import (
	"context"
	"fmt"

	pb "github.com/knwoop/microsercices-example/gen/proto/gateway/v1"
	greetingapi "github.com/knwoop/microsercices-example/gen/proto/greeting/v1"
)

var _ pb.GreeterSer3viceServer = (*server)(nil)

type GreetingAPI struct {
	pb.UnimplementedGreeterServiceServer

	greetingAPI greetingapi.GreeterServiceServer
}

func (api *GreetingAPI) GetUser(ctx context.Context, req *pb.SayHelloRequest) (*pb.GreeterServiceSayHelloResponse, error) {
	apiReq := &greetingapi.SayHelloRequest{
		Name: req.Name,
	}

	res, err := api.greetingAPI.SayHello(ctx, apiReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call greeting api SayHello: %w", err)
	}

	return &pb.GreeterServiceSayHelloResponse{
		Message: res.Message,
	}, nil
}
