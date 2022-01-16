package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Nulandmori/micorservices-pattern/services/catalog/proto"
	customer "github.com/Nulandmori/micorservices-pattern/services/customer/proto"
	item "github.com/Nulandmori/micorservices-pattern/services/item/proto"
)

var _ proto.CatalogServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedCatalogServiceServer
	itemClient     item.ItemServiceClient
	customerClient customer.CustomerServiceClient
}

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	res, err := s.itemClient.CreateItem(ctx, &item.CreateItemRequest{
		CustomerId: "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
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

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	ires, err := s.itemClient.GetItem(ctx, &item.GetItemRequest{Id: req.Id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "not found")
		}
	}

	i := ires.GetItem()
	if i == nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	cres, err := s.customerClient.GetCustomer(ctx, &customer.GetCustomerRequest{Id: i.CustomerId})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}
	c := cres.GetCustomer()
	if c == nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.GetItemResponse{
		Item: &proto.Item{
			Id:           i.Id,
			CustomerId:   i.CustomerId,
			CustomerName: c.Name,
			Title:        i.Title,
			Price:        int64(i.Price),
		},
	}, nil
}
