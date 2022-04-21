package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	httppkg "net/http"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/knwoop/microsercices-example/services/gateway/http"
)

var (
	port = flag.Int("port", 4000, "The server port")
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	httpServer, err := http.New(ctx, *port)
	if err != nil {
		fmt.Printf("faield to set up http server: %s", err)
		return 1
	}

	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- httpServer.ListenAndServe()
	}()

	select {
	case err := <-httpErrCh:
		fmt.Println("failed to serve http server", err)
		return 1
	case <-ctx.Done():
		fmt.Println("shutting down...")
		if err := httpServer.Shutdown(ctx); err != nil {
			fmt.Printf("failed to shutdown http server: %v", err)
			return 1
		}

		if err := <-httpErrCh; err != nil && !errors.Is(err, httppkg.ErrServerClosed) {
			fmt.Printf("failed to close http server: %v", err)
			return 1
		}

		return 0
	}
}
