package systemdlog

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
)

type SystemdHandler struct {
	slog.Handler
	inner *log.Logger
}

func (h *SystemdHandler) Handle(ctx context.Context, r slog.Record) error {
	level := priorityFromSlog(r.Level).String()

	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	msg := level + r.Message

	if r.NumAttrs() > 0 {
		b, err := json.Marshal(fields)
		if err != nil {
			return err
		}

		h.inner.Println(msg, string(b))
	} else {
		h.inner.Println(msg)
	}

	return nil
}

func NewSystemdHandler(out io.Writer, opts slog.HandlerOptions) *SystemdHandler {
	h := &SystemdHandler{
		Handler: slog.NewJSONHandler(out, &opts),
		inner:   log.New(out, "", 0),
	}

	return h
}
