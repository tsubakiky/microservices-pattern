package grpc

import (
	"context"
	"fmt"

	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"

	"github.com/Nulandmori/micorservices-pattern/services/item/pkg/db"
	"github.com/Nulandmori/micorservices-pattern/services/item/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	db, err := db.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create db client: %v", err)
	}
	defer db.Close()

	svc := &server{
		logger: logger.WithName("server"),
		db:     db,
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterItemServiceServer(s, svc)
	}).Start(ctx)
}
