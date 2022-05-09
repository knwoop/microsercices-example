package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/knwoop/microsercices-example/gen/proto/greeting/v1"
)

var _ pb.GreeterServiceServer = (*server)(nil)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.SayHelloResponse{Message: "Hello " + in.GetName()}, nil
}

func New() (*grpc.Server, error) {
	s := grpc.NewServer()
	pb.RegisterGreeterServiceServer(s, &server{})

	reflection.Register(s)

	return s, nil
}
