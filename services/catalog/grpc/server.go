package grpc

import (
	"bytes"
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/lestrrat-go/jwx/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	item "github.com/Nulandmori/micorservices-pattern/services/item/proto"
)

var _ proto.CatalogServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedCatalogServiceServer
	itemClient item.ItemServiceClient
}

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	tokenStr, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "token not found")
	}

	token, err := jwt.Parse(bytes.NewBufferString(tokenStr).Bytes())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed to parse access token")
	}

	res, err := s.itemClient.CreateItem(ctx, &item.CreateItemRequest{
		CustomerId: token.Subject(),
		Title:      req.Title,
		Price:      req.Price,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	item := res.GetItem()

	return &proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         item.Id,
			CustomerId: item.CustomerId,
			Title:      item.Title,
			Price:      item.Price,
		},
	}, nil
}
