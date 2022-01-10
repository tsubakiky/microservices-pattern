package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/Nulandmori/micorservices-pattern/pkg/env"
	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	item "github.com/Nulandmori/micorservices-pattern/services/item/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultTLSPort = "443"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}

	itemServiceAddr := env.MustGetEnv("ITEM_SERVICE_ADDR")

	if strings.Contains(itemServiceAddr, defaultTLSPort) {
		creds := credentials.NewTLS(&tls.Config{})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.DialContext(ctx, itemServiceAddr, opts...)
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
