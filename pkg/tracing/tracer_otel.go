package tracing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	instrumentationName    = "https://github.com/justtrackio/gosoline"
	instrumentationVersion = "v0.8.0"
)

type OtelSettings struct {
	Exporter string `cfg:"exporter"`
}

type otelTracer struct {
	cfg.AppId
	tracer trace.Tracer
}

func NewOtelTracer(ctx context.Context, config cfg.Config, logger log.Logger) (Tracer, error) {
	appId := cfg.AppId{}
	appId.PadFromConfig(config)

	settings := &OtelSettings{}
	config.UnmarshalKey("tracing.otel", settings)

	otelExporterFactory := TraceExporters[settings.Exporter]

	exporter, err := otelExporterFactory(ctx, config, logger)
	if err != nil {
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(fmt.Sprintf("%v-%v-%v-%v", appId.Project, appId.Environment, appId.Family, appId.Application)),
		)),
		// sdktrace.WithSampler() // TODO: use existing samping config to populate this sampler
		// sdktrace.WithRawSpanLimits() // TODO: check if we were supporting similar attribute for xray i.e StreamingMaxSubsegmentCount
	)

	otel.SetTracerProvider(tracerProvider)

	tracer := otel.GetTracerProvider().Tracer(
		instrumentationName,
		trace.WithInstrumentationVersion(instrumentationVersion),
		trace.WithSchemaURL(semconv.SchemaURL),
	)

	return NewOtelTracerWithInterfaces(appId, tracer), nil
}

func NewOtelTracerWithInterfaces(appId cfg.AppId, tracer trace.Tracer) *otelTracer {
	return &otelTracer{
		AppId:  appId,
		tracer: tracer,
	}
}

func (t *otelTracer) StartSubSpan(ctx context.Context, name string) (context.Context, Span) {
	ctx, span := t.tracer.Start(ctx, name)

	return newOtelSpan(ctx, span)
}

func (t *otelTracer) StartSpan(name string) (context.Context, Span) {
	return t.newRootSpan(context.Background(), name)
}

func (t *otelTracer) StartSpanFromContext(ctx context.Context, name string) (context.Context, Span) {
	if parentSpan := GetSpanFromContext(ctx); parentSpan != nil {
		parentTrace := parentSpan.GetTrace()
		var tFlags trace.TraceFlags
		if parentTrace.GetSampled() {
			tFlags = trace.FlagsSampled
		}

		tID, _ := trace.TraceIDFromHex(parentTrace.GetTraceId())
		sID, _ := trace.SpanIDFromHex(parentSpan.GetId())

		spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
			TraceID:    tID,
			SpanID:     sID,
			TraceFlags: tFlags,
		})
		ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)

		return t.newRootSpan(ctx, name)
	}

	if ctxTrace := GetTraceFromContext(ctx); ctxTrace != nil {
		var tFlags trace.TraceFlags
		if ctxTrace.GetSampled() {
			tFlags = trace.FlagsSampled
		}

		tID, _ := trace.TraceIDFromHex(ctxTrace.GetTraceId())
		sID, _ := trace.SpanIDFromHex(ctxTrace.GetId())

		spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
			TraceID:    tID,
			SpanID:     sID,
			TraceFlags: tFlags,
		})
		ctx = trace.ContextWithRemoteSpanContext(ctx, spanCtx)

		return t.newRootSpan(ctx, name)
	}

	return t.newRootSpan(ctx, name)
}

func (t *otelTracer) HttpHandler(h http.Handler) http.Handler {
	name := fmt.Sprintf("%v-%v-%v-%v", t.Project, t.Environment, t.Family, t.Application)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, span := t.tracer.Start(r.Context(), name)

		ctx, _ = newOtelSpan(ctx, span)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})

	return otelhttp.NewHandler(handlerFunc, name)
}

func (t *otelTracer) newRootSpan(ctx context.Context, name string) (context.Context, *otelSpan) {
	ctx, span := t.tracer.Start(ctx, name)

	return newOtelSpan(ctx, span)
}
