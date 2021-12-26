package grpc

import (
	"context"
	"fmt"

	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	customer "github.com/Nulandmori/micorservices-pattern/services/customer/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true)),
	}
	conn, err := grpc.DialContext(ctx, "customer_app:8080", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial grpc server: %w", err)
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
