package grpc

import (
	"context"

	pkggrpc "github.com/Nulandmori/micorservices-pattern/pkg/grpc"
	"github.com/Nulandmori/micorservices-pattern/services/item/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, logger logr.Logger) error {

	svc := &server{
		logger: logger.WithName("server"),
	}

	return pkggrpc.NewServer(port, logger, func(s *grpc.Server) {
		proto.RegisterItemServiceServer(s, svc)
	}).Start(ctx)
}
