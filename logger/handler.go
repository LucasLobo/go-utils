package logger

import (
	"context"
	"io"
	"log/slog"
)

// CtxHandler is a [slog.Handler] that adds context attributes to the log record.
// Attributes must be added using the [WithValue] or [WithValues] function.
type CtxHandler struct {
	internal slog.Handler
}

// NewCtxHandler creates a [CtxHandler] based on [slog.JSONHandler] that writes to w,
// using the given options.
// If opts is nil, the default options are used.
func NewCtxHandler(w io.Writer, opts *slog.HandlerOptions) *CtxHandler {
	return &CtxHandler{
		internal: slog.NewJSONHandler(w, opts),
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *CtxHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.internal.Enabled(ctx, level)
}

// WithAttrs returns a new [CtxHandler] whose attributes consists
// of h's attributes followed by attrs.
func (h *CtxHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CtxHandler{internal: h.internal.WithAttrs(attrs)}
}

// WithGroup returns a new [CtxHandler] whose group is set to name.
func (h *CtxHandler) WithGroup(name string) slog.Handler {
	return &CtxHandler{internal: h.internal.WithGroup(name)}
}

// Handle extends the default [slog.JSONHandler] Handle method by adding context attributes to the log record.
func (h *CtxHandler) Handle(ctx context.Context, r slog.Record) error {
	ctxMap, ok := getCtxMap(ctx)
	if ok {
		r = r.Clone()
		for _, val := range ctxMap {
			r.AddAttrs(val)
		}
	}
	return h.internal.Handle(ctx, r)
}
