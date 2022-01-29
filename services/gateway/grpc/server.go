package grpc

import (
	"context"

	authority "github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	catalog "github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/proto"
)

var (
	_ proto.GatewayServiceServer = (*server)(nil)
)

type server struct {
	proto.UnimplementedGatewayServiceServer
	catalogClient   catalog.CatalogServiceClient
	authorityClient authority.AuthorityServiceClient
}

func (s *server) Signup(ctx context.Context, req *authority.SignupRequest) (*authority.SignupResponse, error) {
	return s.authorityClient.Signup(ctx, req)
}

func (s *server) Signin(ctx context.Context, req *authority.SigninRequest) (*authority.SigninResponse, error) {
	return s.authorityClient.Signin(ctx, req)
}

func (s *server) CreateItem(ctx context.Context, req *catalog.CreateItemRequest) (*catalog.CreateItemResponse, error) {
	return s.catalogClient.CreateItem(ctx, req)
}

func (s *server) GetItem(ctx context.Context, req *catalog.GetItemRequest) (*catalog.GetItemResponse, error) {
	return s.catalogClient.GetItem(ctx, req)
}
