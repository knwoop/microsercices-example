package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	greetingpb "github.com/knwoop/microsercices-example/services/greeting/proto"
)

func New(ctx context.Context, port int) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: false,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	greetingConn, err := grpc.DialContext(ctx, "localhost:50051", opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial to geeting grpc server: %w", err)
	}
	if err := greetingpb.RegisterGreeterHandlerClient(ctx, mux, greetingpb.NewGreeterClient(greetingConn)); err != nil {
		return nil, fmt.Errorf("failed to create a catalog grpc client: %w", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return server, nil
}
