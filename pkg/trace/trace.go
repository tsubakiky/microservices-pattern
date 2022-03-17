package trace

import (
	"context"
	"fmt"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/Nulandmori/micorservices-pattern/pkg/env"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTraceProvider(ctx context.Context) (func(), error) {
	projectID := env.MustGetEnv("GOOGLE_CLOUD_PROJECT")
	if len(projectID) == 0 {
		return nil, fmt.Errorf("GOOGLE_CLOUD_PROJECT not set")
	}
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		return nil, fmt.Errorf("texporter.NewExporter(): %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter))

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return func() {
		_ = tp.ForceFlush(ctx) // flushes any pending spans
	}, nil
}
