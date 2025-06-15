package logger

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

var globalLogger *slog.Logger

func InitSlog(serviceName, appMode string, level slog.Level) {
	// if output == nil {
	// 	output = os.Stdout
	// }
	output := os.Stdout

	prettyHandler := NewPrettyHandler(output, PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: level,
		},
		ServiceName: serviceName,
	})
	baseLogger := slog.New(prettyHandler)

	globalLogger = baseLogger
	slog.SetDefault(globalLogger)
}

func Get() *slog.Logger {
	if globalLogger == nil {
		return slog.Default()
	}

	return globalLogger
}

func Ctx(ctx context.Context, baseLogger *slog.Logger) *slog.Logger {
	l := slog.Default()

	if baseLogger != nil {
		l = baseLogger
	}

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		return l.With(
			slog.String("trace_id", span.SpanContext().TraceID().String()),
			slog.String("span_id", span.SpanContext().SpanID().String()),
		)
	}

	return l
}

func Info(ctx context.Context, msg string, args ...any) {
	l := Ctx(ctx, globalLogger)
	l.Info(msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	l := Ctx(ctx, globalLogger)
	l.Warn(msg, args...)
}

func Error(ctx context.Context, msg string, err error, args ...any) {
	allArgs := append([]any{slog.Any("error", err.Error())}, args...)
	l := Ctx(ctx, globalLogger)
	l.Error(msg, allArgs...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	l := Ctx(ctx, globalLogger)
	l.Debug(msg, args...)
}
