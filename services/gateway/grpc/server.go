package grpc

import (
	"context"

	catalog "github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/proto"
)

var (
	_ proto.GatewayServiceServer = (*server)(nil)
)

type server struct {
	proto.UnimplementedGatewayServiceServer
	catalogClient catalog.CatalogServiceClient
}

func (s *server) CreateItem(ctx context.Context, req *catalog.CreateItemRequest) (*catalog.CreateItemResponse, error) {
	return s.catalogClient.CreateItem(ctx, req)
}

func (s *server) GetItem(ctx context.Context, req *catalog.GetItemRequest) (*catalog.GetItemResponse, error) {
	return s.catalogClient.GetItem(ctx, req)
}
