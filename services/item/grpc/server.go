package grpc

import (
	"context"

	grpccontext "github.com/Nulandmori/micorservices-pattern/pkg/grpc/context"
	"github.com/Nulandmori/micorservices-pattern/services/item/ent"

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
	return &proto.CreateItemResponse{
		Item: &proto.Item{
			Id:         "bda92da6-3270-4255-a756-dbe7d0aa333e",
			CustomerId: req.CustomerId,
			Title:      "Keyboard",
			Price:      30000,
		},
	}, nil
}

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	return &proto.GetItemResponse{
		Item: &proto.Item{
			Id:         req.Id,
			CustomerId: "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
			Title:      "Keyboard",
			Price:      30000,
		},
	}, nil
}

func (s *server) log(ctx context.Context) logr.Logger {
	reqid := grpccontext.GetRequestID(ctx)

	return s.logger.WithValues("request_id", reqid)
}
