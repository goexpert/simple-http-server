package tracer

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func InitTracer() func() {
	ctx := context.Background()
	exporter := os.Getenv("exporter")

	exporterJaeger, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint("jaeger:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	exporterPretty, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("simple-http-server"),
	)

	var traceProvider *trace.TracerProvider

	if exporter == "jaeger" {
		traceProvider = trace.NewTracerProvider(
			trace.WithBatcher(exporterJaeger),
			trace.WithResource(resources),
		)

	} else {

		traceProvider = trace.NewTracerProvider(
			trace.WithBatcher(exporterPretty),
			trace.WithResource(resources),
		)
	}

	otel.SetTracerProvider(traceProvider)

	return func() {

		err := traceProvider.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}

	}
}
