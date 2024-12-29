package tracer

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("simple-http-server"),
	)

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resources),
	)

	otel.SetTracerProvider(traceProvider)

	return func() {

		err := traceProvider.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}

	}
}
