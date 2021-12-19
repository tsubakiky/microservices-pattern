package grpc

import (
	"context"

	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/customer/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {

	svc := &server{}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterCustomerServiceServer(s, svc)
	}).Start(ctx)
}
