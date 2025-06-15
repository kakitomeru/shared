package logger

import (
	"context"
	"log/slog"
	"os"
	"time"

	"go.opentelemetry.io/otel/trace"
)

var globalLogger *slog.Logger

func InitSlog(serviceName, appMode string, level slog.Level) {
	// if output == nil {
	// 	output = os.Stdout
	// }
	output := os.Stdout

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
				a.Value = slog.StringValue(a.Value.Time().Format(time.RFC3339Nano))
			}

			return a
		},
	}

	handler := slog.NewTextHandler(output, opts)
	baseLogger := slog.New(handler).With(slog.String("service", serviceName))

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
	Ctx(ctx, globalLogger).Info(msg, args...)
}
func Warn(ctx context.Context, msg string, args ...any) {
	Ctx(ctx, globalLogger).Warn(msg, args...)
}

func Error(ctx context.Context, msg string, err error, args ...any) {
	allArgs := append([]any{slog.Any("error", err)}, args...)
	Ctx(ctx, globalLogger).Error(msg, allArgs...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	Ctx(ctx, globalLogger).Debug(msg, args...)
}
