package grpc

import (
	"context"
	"fmt"
	"net"

	"time"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/Nulandmori/micorservices-pattern/pkg/grpc/server/interceptor"
	"github.com/go-logr/logr"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

var defaultNOPAuthFunc = func(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, logger logr.Logger, register func(server *grpc.Server)) *Server {
	interceptors := []grpc.UnaryServerInterceptor{
		interceptor.NewRequestLogger(logger.WithName("request")),
		grpc_auth.UnaryServerInterceptor(defaultNOPAuthFunc),
		otelgrpc.UnaryServerInterceptor(),
	}

	opts := []grpc.ServerOption{
		middleware.WithUnaryServerChain(interceptors...),
	}

	server := grpc.NewServer(opts...)

	register(server)

	reflection.Register(server)
	channelz.RegisterChannelzServiceToServer(server)

	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on %d: %v", s.port, err)
	}

	errCh := make(chan error, 1)

	go func() {
		if err := s.server.Serve(listener); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("server has stopped with error: %v", err)
		}
		return nil
	case <-ctx.Done():
		s.server.GracefulStop()
		return <-errCh
	}
}

const projectID = "gaudiy-integration-test"

func initTraceProvider(logger logr.Logger) {
	// When running on GCP, authentication is handled automatically
	// using default credentials. This environment variable check
	// is to help debug projects running locally. It's possible for this
	// warning to be printed while the exporter works normally. See
	// https://developers.google.com/identity/protocols/application-default-credentials
	// for more details.

	// projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if len(projectID) == 0 {
		logger.V(1).Info("GOOGLE_CLOUD_PROJECT not set")
	}
	for i := 1; i <= 3; i++ {
		exporter, err := texporter.New(texporter.WithProjectID(projectID))
		if err != nil {
			logger.Info("failed to initialize exporter: %v", err)
		} else {
			// Create trace provider with the exporter.
			// The AlwaysSample sampling policy is used here for demonstration
			// purposes and should not be used in production environments.
			tp := sdktrace.NewTracerProvider(
				sdktrace.WithSampler(sdktrace.AlwaysSample()),
				sdktrace.WithBatcher(exporter),
			)
			if err == nil {
				logger.Info("initialized trace provider")
				otel.SetTracerProvider(tp)
				otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
				return
			} else {
				d := time.Second * 10 * time.Duration(i)
				logger.Info("sleeping %v to retry initializing trace provider", d)
				time.Sleep(d)
			}
		}
	}
	logger.V(1).Info("failed to initialize trace provider")
}
