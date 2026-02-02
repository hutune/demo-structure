package logger

import (
	"context"

	"github.com/rs/zerolog"
)

type Level string

const (
	TraceLv    Level = "trace"
	DebugLv    Level = "debug"
	InfoLv     Level = "info"
	WarnLv     Level = "warn"
	ErrorLv    Level = "error"
	FatalLv    Level = "fatal"
	PanicLv    Level = "panic"
	NoLevelLv  Level = "no_level"
	DisabledLv Level = "disabled"
)

func (l Level) ToZerologLevel() zerolog.Level {
	switch l {
	case TraceLv:
		return zerolog.TraceLevel
	case DebugLv:
		return zerolog.DebugLevel
	case InfoLv:
		return zerolog.InfoLevel
	case WarnLv:
		return zerolog.WarnLevel
	case ErrorLv:
		return zerolog.ErrorLevel
	case FatalLv:
		return zerolog.FatalLevel
	case PanicLv:
		return zerolog.PanicLevel
	case NoLevelLv:
		return zerolog.NoLevel
	case DisabledLv:
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}

type ILogger interface {
	Debug(ctx context.Context, msg string, fields ...interface{})
	Info(ctx context.Context, msg string, fields ...interface{})
	Warn(ctx context.Context, msg string, fields ...interface{})
	Error(ctx context.Context, msg string, fields ...interface{})
	Fatal(ctx context.Context, msg string, fields ...interface{})
	With(fields ...interface{}) ILogger
}
