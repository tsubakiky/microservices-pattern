package grpc

import (
	"context"
	"fmt"

	"github.com/Nulandmori/micorservices-pattern/services/customer/proto"
)

var _ proto.CustomerServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedCustomerServiceServer
}

func (s *server) GetCustomer(ctx context.Context, req *proto.GetCustomerRequest) (*proto.GetCustomerResponse, error) {
	return &proto.GetCustomerResponse{
		Customer: &proto.Customer{
			Id:   "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
			Name: "goldie",
		},
	}, nil
}

func (s *server) CreateCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	fmt.Println("Create Customer!")
	return &proto.CreateCustomerResponse{
		Customer: &proto.Customer{
			Id:   "7c0cde05-4df0-47f4-94c4-978dd9f56e5c",
			Name: "goldie",
		},
	}, nil
}
