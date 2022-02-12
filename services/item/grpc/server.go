package grpc

import (
	"context"

	"github.com/Nulandmori/micorservices-pattern/services/item/ent"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Nulandmori/micorservices-pattern/services/item/proto"
	"github.com/go-logr/logr"
)

var (
	_ proto.ItemServiceServer = (*server)(nil)
)

type server struct {
	proto.UnimplementedItemServiceServer
	logger logr.Logger
	db     *ent.Client
}

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	item, err := s.db.Item.Create().SetCustomerID(req.CustomerId).SetTitle(req.Title).SetPrice(req.Price).Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         item.ID.String(),
			CustomerId: item.CustomerID,
			Title:      item.Title,
			Price:      item.Price,
		},
	}, nil
}

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	item, err := s.db.Item.Get(ctx, uuid.MustParse(req.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.GetItemResponse{
		Item: &proto.Item{
			Id:         req.Id,
			CustomerId: item.CustomerID,
			Title:      item.Title,
			Price:      item.Price,
		},
	}, nil
}
