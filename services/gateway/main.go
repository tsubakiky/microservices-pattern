package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/Nulandmori/micorservices-pattern/pkg/logger"
	"github.com/Nulandmori/micorservices-pattern/pkg/run"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/http"
)

func main() {
	run.Run(server)
}

func server(ctx context.Context) int {
	httpPort := 8080
	grpcPort := 9090

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		return 1
	}
	glogger := l.WithName("gateway")

	if len(os.Getenv("PORT")) > 0 {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			fmt.Printf("cannot convert %q to number!\n", os.Getenv("PORT"))
		}
		httpPort = p
	}

	grpcErrCh := make(chan error, 1)
	go func() {
		grpcErrCh <- grpc.RunServer(ctx, grpcPort, glogger.WithName("grpc"))
	}()

	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- http.RunServer(ctx, httpPort, grpcPort)
	}()

	select {
	case err := <-grpcErrCh:
		glogger.Error(err, "failed to serve grpc server")
		return 1
	case err := <-httpErrCh:
		glogger.Error(err, "failed to serve http server")
		return 1
	case <-ctx.Done():
		glogger.Info("shutting down...")
		return 0
	}
}
