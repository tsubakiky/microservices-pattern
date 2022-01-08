package grpc

import (
	"context"
	"fmt"

	"github.com/Nulandmori/micorservices-pattern/pkg/env"
	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	item "github.com/Nulandmori/micorservices-pattern/services/item/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}
	conn, err := grpc.DialContext(ctx, env.GetEnv("ITEM_SERVICE_ADDR", "item_app:8080"), opts...)
	if err != nil {
		return fmt.Errorf("failed to dial item grpc server: %w", err)
	}

	itemClient := item.NewItemServiceClient(conn)

	svc := &server{
		itemClient: itemClient,
	}
	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterCatalogServiceServer(s, svc)
	}).Start(ctx)
}
