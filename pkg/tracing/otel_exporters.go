package tracing

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

type OtelExporterFactory func(ctx context.Context, config cfg.Config, logger log.Logger) (*otlptrace.Exporter, error)

func AddTraceExporter(name string, exporter OtelExporterFactory) {
	TraceExporters[name] = exporter
}

var TraceExporters = map[string]OtelExporterFactory{
	"otel_http": NewOtelHttpTracer,
}

func NewOtelHttpTracer(ctx context.Context, config cfg.Config, logger log.Logger) (*otlptrace.Exporter, error) {
	client := otlptracehttp.NewClient()
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP trace exporter: %w", err)
	}

	return exporter, nil
}
