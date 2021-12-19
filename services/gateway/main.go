package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Nulandmori/micorservices-pattern/pkg/logger"
	"github.com/Nulandmori/micorservices-pattern/pkg/run"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/http"
)

func main() {
	run.Run(server)
}

func server(ctx context.Context) int {
	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		return 1
	}
	glogger := l.WithName("gateway")

	grpcErrCh := make(chan error, 1)
	go func() {
		grpcErrCh <- grpc.RunServer(ctx, 4000, glogger.WithName("grpc"))
	}()

	httpErrCh := make(chan error, 1)
	go func() {
		httpErrCh <- http.RunServer(ctx, 4001, 4000)
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
