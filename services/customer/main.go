package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"github.com/Nulandmori/micorservices-pattern/pkg/logger"
	"github.com/Nulandmori/micorservices-pattern/services/customer/grpc"
	"golang.org/x/sys/unix"
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	port := 8080

	ctx, stop := signal.NotifyContext(ctx, unix.SIGTERM, unix.SIGINT)
	defer stop()

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			// Unhandleable, something went wrong...
			panic(fmt.Sprintf("failed to write log:`%s` original error is:`%s`", ferr, err))
		}
		return 1
	}
	clogger := l.WithName("customer")

	if len(os.Getenv("PORT")) > 0 {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			fmt.Printf("cannot convert %q to number!\n", os.Getenv("PORT"))
		}
		port = p
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.RunServer(ctx, port, clogger.WithName("grpc"))
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
