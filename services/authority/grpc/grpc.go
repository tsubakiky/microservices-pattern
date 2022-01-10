package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/Nulandmori/micorservices-pattern/pkg/env"
	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	customer "github.com/Nulandmori/micorservices-pattern/services/customer/proto"
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

	customerServiceAddr := env.MustGetEnv("CUSTOMER_SERVICE_ADDR")

	if strings.Contains(customerServiceAddr, defaultTLSPort) {
		creds := credentials.NewTLS(&tls.Config{})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.DialContext(ctx, customerServiceAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial customer grpc server: %w", err)
	}

	customerClient := customer.NewCustomerServiceClient(conn)

	svc := &server{
		customerClient: customerClient,
		logger:         logger.WithName("server"),
	}
	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterAuthorityServiceServer(s, svc)
	}).Start(ctx)
}
