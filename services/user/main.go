package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/knwoop/microsercices-example/services/user/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen on port %d: %s\n", *port, err)
	}

	grpcServer, err := grpc.New()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpcServer.Serve(listener)
	}()

	select {
	case err := <-errCh:
		fmt.Println(err.Error())
		return 1
	case <-ctx.Done():
		fmt.Println("shutting down...")
		return 0
	}
}
