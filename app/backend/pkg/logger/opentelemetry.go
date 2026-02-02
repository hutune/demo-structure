package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetryDecorator struct{}

func NewOpenTelemetryDecorator() *OpenTelemetryDecorator {
	return &OpenTelemetryDecorator{}
}

func (d *OpenTelemetryDecorator) Decorate(logger ILogger) ILogger {
	return &openTelemetryLogger{
		logger: logger,
	}
}

type openTelemetryLogger struct {
	logger ILogger
}

func (l *openTelemetryLogger) log(ctx context.Context, level string, msg string, fields ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		traceID := span.SpanContext().TraceID().String()
		spanID := span.SpanContext().SpanID().String()

		fields = append(fields, "trace_id", traceID, "span_id", spanID)
	}

	switch level {
	case "debug":
		l.logger.Debug(ctx, msg, fields...)
	case "info":
		l.logger.Info(ctx, msg, fields...)
	case "warn":
		l.logger.Warn(ctx, msg, fields...)
	case "error":
		l.logger.Error(ctx, msg, fields...)
	case "fatal":
		l.logger.Fatal(ctx, msg, fields...)
	}
}

func (l *openTelemetryLogger) Debug(ctx context.Context, msg string, fields ...interface{}) {
	l.log(ctx, "debug", msg, fields...)
}

func (l *openTelemetryLogger) Info(ctx context.Context, msg string, fields ...interface{}) {
	l.log(ctx, "info", msg, fields...)
}

func (l *openTelemetryLogger) Warn(ctx context.Context, msg string, fields ...interface{}) {
	l.log(ctx, "warn", msg, fields...)
}

func (l *openTelemetryLogger) Error(ctx context.Context, msg string, fields ...interface{}) {
	l.log(ctx, "error", msg, fields...)
}

func (l *openTelemetryLogger) Fatal(ctx context.Context, msg string, fields ...interface{}) {
	l.log(ctx, "fatal", msg, fields...)
}

func (l *openTelemetryLogger) With(fields ...interface{}) ILogger {
	return &openTelemetryLogger{
		logger: l.logger.With(fields...),
	}
}
