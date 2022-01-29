package grpc

import (
	"context"
	"fmt"

	"github.com/Nulandmori/micorservices-pattern/pkg/firebase"
	grpccontext "github.com/Nulandmori/micorservices-pattern/pkg/grpc/context"
	catalog "github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	"github.com/Nulandmori/micorservices-pattern/services/gateway/proto"
	"github.com/go-logr/logr"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	_ proto.GatewayServiceServer        = (*server)(nil)
	_ grpc_auth.ServiceAuthFuncOverride = (*server)(nil)

	publicRPCMethods = map[string]struct{}{
		"/gateway.GatewayService/GetItem": {},
	}
)

type server struct {
	proto.UnimplementedGatewayServiceServer
	catalogClient catalog.CatalogServiceClient
	firebaseAuth  firebase.FirebaseAuth
	logger        logr.Logger
}

func (s *server) CreateItem(ctx context.Context, req *catalog.CreateItemRequest) (*catalog.CreateItemResponse, error) {
	return s.catalogClient.CreateItem(ctx, req)
}

func (s *server) GetItem(ctx context.Context, req *catalog.GetItemRequest) (*catalog.GetItemResponse, error) {
	return s.catalogClient.GetItem(ctx, req)
}

func (s *server) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	_, ok := publicRPCMethods[fullMethodName]
	if ok {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		s.log(ctx).Info("failed to get token from authorization header")
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	_, err = s.firebaseAuth.VerifyIDToken(ctx, token)
	if err != nil {
		s.log(ctx).Info(fmt.Sprintf("failed to verify token: %s", err.Error()))
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	return ctx, nil
}

func (s *server) log(ctx context.Context) logr.Logger {
	reqid := grpccontext.GetRequestID(ctx)

	return s.logger.WithValues("request_id", reqid)
}
