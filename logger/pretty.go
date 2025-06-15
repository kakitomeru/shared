package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"strings"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts    *slog.HandlerOptions
	ServiceName string
}

type PrettyHandler struct {
	slog.Handler
	l           *log.Logger
	serviceName string
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	default:
		level = r.Level.String()
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b := ""
	if len(fields) > 0 {
		r, err := json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}

		b = string(r)
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(color.WhiteString(strings.ToUpper(h.serviceName)), timeStr, level, msg, color.WhiteString(b))

	return nil
}

func NewPrettyHandler(w io.Writer, opts PrettyHandlerOptions) *PrettyHandler {
	return &PrettyHandler{
		Handler:     slog.NewJSONHandler(w, opts.SlogOpts),
		l:           log.New(w, "", 0),
		serviceName: opts.ServiceName,
	}
}
