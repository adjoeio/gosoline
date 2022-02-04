package tracing

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type jaegerSpan struct {
	enabled bool
	span    trace.Span
}

func (s *jaegerSpan) AddAnnotation(key string, value string) {
	s.span.SetAttributes(attribute.String(key, value))
}

func (s *jaegerSpan) AddError(err error) {
	s.span.RecordError(err)
}

func (s *jaegerSpan) AddMetadata(key string, value interface{}) {
	v, ok := value.(fmt.Stringer)
	if !ok {
		return
	}
	s.span.SetAttributes(attribute.Stringer(key, v))
}

func (s *jaegerSpan) Finish() {
	s.span.End()
}

func (s *jaegerSpan) GetId() string {
	return s.span.SpanContext().SpanID().String()
}

func (s *jaegerSpan) GetTrace() *Trace {
	return &Trace{
		TraceId:  s.span.SpanContext().TraceID().String(),
		Id:       s.GetId(),
		ParentId: "",
		Sampled:  s.span.SpanContext().IsSampled(),
	}
}
