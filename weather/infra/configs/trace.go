package configs

import "go.opentelemetry.io/otel/trace"

var Tracer trace.Tracer

func NewTracer(tracer trace.Tracer) {
	Tracer = tracer
}

func GetTracer() trace.Tracer {
	return Tracer
}
