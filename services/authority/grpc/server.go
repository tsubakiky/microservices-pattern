package grpc

import (
	"context"
	"fmt"

	grpccontext "github.com/Nulandmori/micorservices-pattern/pkg/grpc/context"
	"github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	customer "github.com/Nulandmori/micorservices-pattern/services/customer/proto"
	"github.com/go-logr/logr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ proto.AuthorityServiceServer = (*server)(nil)

type server struct {
	proto.UnimplementedAuthorityServiceServer

	customerClient customer.CustomerServiceClient
	logger         logr.Logger
}

func (s *server) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	fmt.Println("Start Signup!")
	c, err := s.customerClient.CreateCustomer(ctx, &customer.CreateCustomerRequest{Name: req.Name})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			return nil, status.Error(codes.AlreadyExists, "customer already exists by given name")
		}
		s.log(ctx).Error(err, "failed to call customer service")
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &proto.SignupResponse{
		Customer: &customer.Customer{
			Id:   c.GetCustomer().Id,
			Name: c.GetCustomer().Name,
		},
	}, nil
}

func (s *server) Signin(ctx context.Context, req *proto.SigninRequest) (*proto.SigninResponse, error) {
	_, err := s.customerClient.GetCustomerByName(ctx, &customer.GetCustomerByNameRequest{Name: req.Name})
	if err != nil {
		s.log(ctx).Info(fmt.Sprintf("failed to authenticate the customer: %s", err))
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	// TODO: create accessToken
	return &proto.SigninResponse{
		AccessToken: "test",
	}, nil
}

func (s *server) log(ctx context.Context) logr.Logger {
	reqid := grpccontext.GetRequestID(ctx)

	return s.logger.WithValues("request_id", reqid)
}
