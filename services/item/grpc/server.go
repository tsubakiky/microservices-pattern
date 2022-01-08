package grpc

import (
	"context"

	"github.com/Nulandmori/micorservices-pattern/services/item/proto"
)

var _ proto.ItemServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedItemServiceServer
}

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	return &proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         "bda92da6-3270-4255-a756-dbe7d0aa333e",
			CustomerId: req.CustomerId,
			Title:      "Keyboard",
			Price:      30000,
		},
	}, nil
}
