package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/Nulandmori/micorservices-pattern/pkg/env"
	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	authority "github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	catalog "github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/proto"
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

	authorityServiceAddr := env.MustGetEnv("AUTHORITY_SERVICE_ADDR")
	catalogServiceAddr := env.MustGetEnv("CATALOG_SERVICE_ADDR")

	if strings.Contains(authorityServiceAddr, defaultTLSPort) && strings.Contains(catalogServiceAddr, defaultTLSPort) {
		creds := credentials.NewTLS(&tls.Config{})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	aconn, err := grpc.DialContext(ctx, authorityServiceAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial authority grpc server: %w", err)
	}

	cconn, err := grpc.DialContext(ctx, catalogServiceAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial catalog grpc server: %w", err)
	}

	svc := &server{
		authorityClient: authority.NewAuthorityServiceClient(aconn),
		catalogClient:   catalog.NewCatalogServiceClient(cconn),
	}
	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterGatewayServiceServer(s, svc)
	}).Start(ctx)
}
