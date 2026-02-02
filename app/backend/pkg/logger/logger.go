package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"

	"rmn-backend/pkg/constants"
)

type ZeroLogger struct {
	logger zerolog.Logger
	caller int // caller skip frame - if caller is nil, logger not show caller
}

type Config struct {
	Service          string
	Level            Level
	DisableTimestamp bool
	DisableCaller    bool

	CallerSkipFrame int
}

func NewLogger(config Config) (*ZeroLogger, error) {
	zerolog.SetGlobalLevel(config.Level.ToZerologLevel())

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		With()

	if !config.DisableTimestamp {
		logger = logger.Timestamp()
	}
	if !config.DisableCaller {
		logger = logger.Caller()
	}

	caller := config.CallerSkipFrame
	if caller <= 0 {
		caller = 1
	}

	return &ZeroLogger{
		logger: logger.Logger(),
		caller: caller,
	}, nil
}

// Debug logs a message at debug level
func (z *ZeroLogger) Debug(ctx context.Context, msg string, fields ...interface{}) {
	event := z.logger.Debug()
	event = event.CallerSkipFrame(z.caller)
	event = addContextFields(event, ctx)
	event = addFields(event, fields)
	event.Msg(msg)
}

// Info logs a message at info level
func (z *ZeroLogger) Info(ctx context.Context, msg string, fields ...interface{}) {
	event := z.logger.Info()
	event = event.CallerSkipFrame(z.caller)
	event = addContextFields(event, ctx)
	event = addFields(event, fields)
	event.Msg(msg)
}

// Warn logs a message at warn level
func (z *ZeroLogger) Warn(ctx context.Context, msg string, fields ...interface{}) {
	event := z.logger.Warn()
	event = event.CallerSkipFrame(z.caller)
	event = addContextFields(event, ctx)
	event = addFields(event, fields)
	event.Msg(msg)
}

// Error logs a message at error level
func (z *ZeroLogger) Error(ctx context.Context, msg string, fields ...interface{}) {
	event := z.logger.Error()
	event = event.CallerSkipFrame(z.caller)
	event = addContextFields(event, ctx)
	event = addFields(event, fields)
	event.Msg(msg)
}

// Fatal logs a message at fatal level and exits the program
func (z *ZeroLogger) Fatal(ctx context.Context, msg string, fields ...interface{}) {
	event := z.logger.Fatal()
	event = event.CallerSkipFrame(z.caller)
	event = addContextFields(event, ctx)
	event = addFields(event, fields)
	event.Msg(msg)
}

// With creates a new logger with additional fields
func (z *ZeroLogger) With(fields ...interface{}) ILogger {
	ctx := z.logger.With()
	for i := 0; i < len(fields)-1; i += 2 {
		if key, ok := fields[i].(string); ok {
			ctx = ctx.Interface(key, fields[i+1])
		}
	}
	return &ZeroLogger{
		logger: ctx.Logger(),
	}
}

// addContextFields extracts values from context and adds them to the event
func addContextFields(event *zerolog.Event, ctx context.Context) *zerolog.Event {
	if ctx == nil {
		return event
	}

	// Extract trace ID from context
	if traceID := ctx.Value(constants.ContextKeyRequestID); traceID != nil {
		event = event.Interface(constants.ContextKeyRequestID, traceID)
	}

	return event
}

// addFields adds fields to a zerolog event
// Assumes fields come in pairs: key1, value1, key2, value2, ...
func addFields(event *zerolog.Event, fields []interface{}) *zerolog.Event {
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			if key, ok := fields[i].(string); ok {
				event = event.Interface(key, fields[i+1])
			}
		}
	}
	return event
}
