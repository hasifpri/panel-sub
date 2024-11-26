package infrastructure

import (
	"context"
	"fmt"
	"log"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"
	"strings"

	"time"

	"go.opentelemetry.io/otel"
	//"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"

	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	"go.opentelemetry.io/otel/exporters/jaeger"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var TRACEPROVIDER *sdktrace.TracerProvider

func InitializeTracer() {
	ctx := context.Background()
	stdoutExporter, err := stdout.New(stdout.WithPrettyPrint())
	jaegerExporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(infrastructureconfiguration.JaegerEndpoint)))
	otlpTraceHttpExporterForAspecto, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("collector.aspecto.io"),
		otlptracehttp.WithHeaders(map[string]string{"Authorization": infrastructureconfiguration.AspectoKey}))

	if err != nil {
		log.Fatal(err)
	}

	var tp *sdktrace.TracerProvider
	if strings.Contains(infrastructureconfiguration.TracingTool, "STDOUT") {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(stdoutExporter),
			sdktrace.WithResource(
				resource.NewWithAttributes(
					semconv.SchemaURL,
					semconv.ServiceNameKey.String(infrastructureconfiguration.ServiceName),
					attribute.String("environment", infrastructureconfiguration.Environment),
				)),
		)
	}

	if strings.Contains(infrastructureconfiguration.TracingTool, "JAEGER") {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(jaegerExporter),
			sdktrace.WithResource(
				resource.NewWithAttributes(
					semconv.SchemaURL,
					semconv.ServiceNameKey.String(infrastructureconfiguration.ServiceName),
					attribute.String("environment", infrastructureconfiguration.Environment),
				)),
		)
	}

	if strings.Contains(infrastructureconfiguration.TracingTool, "STDOUT") && strings.Contains(infrastructureconfiguration.TracingTool, "JAEGER") {
		fmt.Println("infrastructureconfiguration.TracingTool CCC: ", infrastructureconfiguration.TracingTool)
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(jaegerExporter),
			sdktrace.WithBatcher(stdoutExporter),
			sdktrace.WithResource(
				resource.NewWithAttributes(
					semconv.SchemaURL,
					semconv.ServiceNameKey.String(infrastructureconfiguration.ServiceName),
					attribute.String("environment", infrastructureconfiguration.Environment),
				)),
		)
	}

	if strings.Contains(infrastructureconfiguration.TracingTool, "STDOUT") && strings.Contains(infrastructureconfiguration.TracingTool, "JAEGER") && strings.Contains(infrastructureconfiguration.TracingTool, "ASPECTO") {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(jaegerExporter),
			sdktrace.WithBatcher(otlpTraceHttpExporterForAspecto),
			sdktrace.WithBatcher(stdoutExporter),
			sdktrace.WithResource(
				resource.NewWithAttributes(
					semconv.SchemaURL,
					semconv.ServiceNameKey.String(infrastructureconfiguration.ServiceName),
					attribute.String("environment", infrastructureconfiguration.Environment),
				)),
		)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	TRACEPROVIDER = tp
}

func CloseTracer(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	// Cleanly shutdown and flush telemetry when the application exits.
	// Do not make the application hang when it is shutdown.
	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	if err := TRACEPROVIDER.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
