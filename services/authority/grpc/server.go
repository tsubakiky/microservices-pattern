package grpc

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/json"
	"encoding/pem"
	"fmt"

	grpccontext "github.com/Nulandmori/micorservices-pattern/pkg/grpc/context"
	"github.com/Nulandmori/micorservices-pattern/services/authority/proto"
	customer "github.com/Nulandmori/micorservices-pattern/services/customer/proto"
	"github.com/go-logr/logr"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	rsaPrivateKey *rsa.PrivateKey
	_             proto.AuthorityServiceServer = (*server)(nil)
)

//go:embed private-key.pem
var privateKeyFile []byte

const (
	issuer = "authority"
	kid    = "aa7c6287-c45d-4966-84b4-a1633e4e3a64"
)

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
	res, err := s.customerClient.GetCustomerByName(ctx, &customer.GetCustomerByNameRequest{Name: req.Name})
	if err != nil {
		s.log(ctx).Info(fmt.Sprintf("failed to authenticate the customer: %s", err))
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	token, err := createAccessToken(res.GetCustomer().Id)
	if err != nil {
		s.log(ctx).Error(err, "failed to create the access token")
		return nil, status.Error(codes.Internal, "failed to create access token")
	}
	return &proto.SigninResponse{
		AccessToken: string(token),
	}, nil
}

func createAccessToken(sub string) ([]byte, error) {
	token := jwt.New()

	if err := token.Set(jwt.IssuerKey, issuer); err != nil {
		return nil, fmt.Errorf("failed to set the issuer key to the token: %w", err)
	}

	if err := token.Set(jwt.SubjectKey, sub); err != nil {
		return nil, fmt.Errorf("failed to set the subject key to the token: %w", err)
	}
	headers := jws.NewHeaders()
	if err := headers.Set(jws.KeyIDKey, kid); err != nil {
		return nil, fmt.Errorf("failed to create jws headers: %w", err)
	}

	if err := headers.Set(jws.AlgorithmKey, jwa.RS256); err != nil {
		return nil, fmt.Errorf("failed to set the alg key to the token: %w", err)
	}

	if err := headers.Set(jws.TypeKey, "JWT"); err != nil {
		return nil, fmt.Errorf("failed to set the typ key to the token: %w", err)
	}

	b, err := json.Marshal(token)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the token: %w", err)
	}

	signedToken, err := jws.Sign(b, jwa.RS256, rsaPrivateKey, jws.WithHeaders(headers))
	if err != nil {
		return nil, fmt.Errorf("failed to sign the token: %w", err)
	}
	return signedToken, nil
}

func (s *server) log(ctx context.Context) logr.Logger {
	reqid := grpccontext.GetRequestID(ctx)

	return s.logger.WithValues("request_id", reqid)
}

func init() {
	block, _ := pem.Decode(privateKeyFile)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %s", err))
	}
	rsaPrivateKey = key
}
