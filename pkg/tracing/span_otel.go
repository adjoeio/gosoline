package tracing

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type otelSpan struct {
	span trace.Span
}

func (o *otelSpan) AddAnnotation(key string, value string) {
	o.span.SetAttributes(attribute.String(key, value))
}

func (o *otelSpan) AddError(err error) {
	o.span.RecordError(err)
}

func (o *otelSpan) AddMetadata(key string, value interface{}) {
	// TODO: handle better the metadata. EventOption can be passed to the AddEvent. the value can be set in the attributes
	o.span.AddEvent(fmt.Sprintf("%s:%v", key, value))
}

func (o *otelSpan) Finish() {
	o.span.End()
}

func (o *otelSpan) GetId() string {
	return o.span.SpanContext().SpanID().String()
}

func (o *otelSpan) GetTrace() *Trace {
	var parentID string
	spanCtx := o.span.SpanContext()
	if parentSpanCtx := o.GetParentSpanContext(); parentSpanCtx != nil {
		parentID = parentSpanCtx.SpanID().String()
	}

	return &Trace{
		TraceId:  spanCtx.TraceID().String(),
		Id:       spanCtx.SpanID().String(),
		ParentId: parentID,
		Sampled:  spanCtx.IsSampled(),
	}
}

func (o *otelSpan) GetParentSpanContext() *trace.SpanContext {
	roSpan, ok := o.span.(sdktrace.ReadOnlySpan)
	if !ok {
		return nil
	}

	parent := roSpan.Parent()

	return &parent

}

func newOtelSpan(ctx context.Context, span trace.Span, app cfg.AppId) (context.Context, *otelSpan) {
	s := &otelSpan{
		span: span,
	}

	appFamily := fmt.Sprintf("%v-%v-%v", app.Project, app.Environment, app.Family)
	appId := fmt.Sprintf("%v-%v-%v-%v", app.Project, app.Environment, app.Family, app.Application)
	s.AddAnnotation("appFamily", appFamily)
	s.AddAnnotation("appId", appId)

	return ContextWithSpan(ctx, s), s
}
