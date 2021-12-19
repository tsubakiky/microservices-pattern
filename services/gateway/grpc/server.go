package grpc

import (
	"context"

	authority "github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/proto"
)

var (
	_ proto.GatewayServiceServer = (*server)(nil)
)

type server struct {
	proto.UnimplementedGatewayServiceServer

	authorityClient authority.AuthorityServiceClient
}

func (s *server) Signup(ctx context.Context, req *authority.SignupRequest) (*authority.SignupResponse, error) {
	return s.authorityClient.Signup(ctx, req)
}
