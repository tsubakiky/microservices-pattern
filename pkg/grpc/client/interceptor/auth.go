package interceptor

import (
	"context"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const metadataAuthorizationKey = "authorization"

func AuthServiceUnnaryClientInterceptor(audience string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		tokenSource, err := idtoken.NewTokenSource(ctx, audience)
		if err != nil {
			return err
		}

		token, err := tokenSource.Token()
		if err != nil {
			return err
		}

		ctx = metadata.AppendToOutgoingContext(ctx, metadataAuthorizationKey, "Bearer "+token.AccessToken)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
