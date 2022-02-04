package tracing

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/log"
)

type JaegerSettings struct {
	Enabled                     bool
	Host                        string        `cfg:"host" default:"localhost"`
	Port                        string        `cfg:"port" default:"6831"`
	AttemptReconnecting         bool          `cfg:"attempt_reconnecting" dfault:"true"`
	AttemptReconnectingInterval time.Duration `cfg:"attempt_reconnecting_interval" default:"1m"`
	MaxPacketSize               int           `cfg:"max_packet_size" default:"65000"`
}

type jaegerTracer struct {
	cfg.AppId
	enabled bool
	tracer  trace.Tracer
}

func NewJaegerTracer(config cfg.Config, _ log.Logger) (Tracer, error) {
	appId := cfg.AppId{}
	appId.PadFromConfig(config)

	settings := &TracerSettings{}
	config.UnmarshalKey("tracing", settings)

	jaegerSettings := &JaegerSettings{}
	config.UnmarshalKey("tracing.jaeger", jaegerSettings)

	options := []jaeger.AgentEndpointOption{
		jaeger.WithAgentHost(jaegerSettings.Host),
		jaeger.WithAgentPort(jaegerSettings.Port),
		jaeger.WithMaxPacketSize(jaegerSettings.MaxPacketSize),
	}

	if !jaegerSettings.AttemptReconnecting {
		options = append(options, jaeger.WithDisableAttemptReconnecting())
	}

	if jaegerSettings.AttemptReconnectingInterval <= 0 {
		options = append(options, jaeger.WithAttemptReconnectingInterval(jaegerSettings.AttemptReconnectingInterval))
	}

	jaegerExporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(options...),
	)
	if err != nil {
		return nil, err
	}

	tp := newTraceProvider(appId, jaegerExporter)

	otel.SetTracerProvider(tp)

	tracer := tp.Tracer(namingStrategy(appId))

	return NewJaegerTracerWithInterfaces(appId, tracer, settings.Enabled)
}

func NewJaegerTracerWithInterfaces(appId cfg.AppId, tracer trace.Tracer, enabled bool) (*jaegerTracer, error) {
	return &jaegerTracer{
		AppId:   appId,
		enabled: enabled,
		tracer:  tracer,
	}, nil
}

func (t *jaegerTracer) StartSubSpan(ctx context.Context, name string) (context.Context, Span) {
	if !t.enabled {
		return ctx, disabledJaegerSpan()
	}

	ctxWithSpan, span := t.tracer.Start(ctx, name)
	span.AddEvent(name)

	return ctxWithSpan, &jaegerSpan{span: span}
}

func (t *jaegerTracer) StartSpan(name string) (context.Context, Span) {
	ctx := context.Background()
	if !t.enabled {
		return ctx, disabledJaegerSpan()
	}

	ctx, span := t.tracer.Start(ctx, name)

	return ctx, &jaegerSpan{span: span}
}

func (t *jaegerTracer) StartSpanFromContext(ctx context.Context, name string) (context.Context, Span) {
	if !t.enabled {
		return ctx, disabledJaegerSpan()
	}

	parentSpan := trace.SpanFromContext(ctx)
	parentSpan.AddEvent(name)

	return ctx, &jaegerSpan{span: parentSpan}
}

func (t *jaegerTracer) HttpHandler(h http.Handler) http.Handler {
	if !t.enabled {
		return h
	}

	name := fmt.Sprintf("%v-%v-%v-%v", t.Project, t.Environment, t.Family, t.Application)
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(ctx)

		ctx, _ = t.tracer.Start(ctx, name)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})

	return handlerFunc
}

func newTraceProvider(appId cfg.AppId, exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	r := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(namingStrategy(appId)),
	)

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // TODO: allow to configure the sampler
	)
}

func namingStrategy(appId cfg.AppId) string {
	return fmt.Sprintf("%s.%s.%s.%s", appId.Project, appId.Environment, appId.Family, appId.Application)
}

func disabledJaegerSpan() *jaegerSpan {
	return &jaegerSpan{span: trace.SpanFromContext(nil)}
}
