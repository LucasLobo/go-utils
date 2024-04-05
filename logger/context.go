package logger

import (
	"context"
	"log/slog"
	"maps"
)

// loggerCtxKey is a context key, private to this package to avoid leaking the data out of the package.
type loggerCtxKey struct{}

type loggerCtxMap map[string]slog.Attr

// WithValue returns a new context with the key-value pair added for logging purposes.
// All values added to the context using this function will be included in the logs when logger uses [CtxHandler]
// and the context is passed to the logger.
func WithValue(ctx context.Context, key, val string) context.Context {
	// we clone the map to avoid modifying the original map
	ctxMap := cloneCtxMap(ctx)
	ctxMap[key] = slog.String(key, val)
	// create a new context with the new map
	return context.WithValue(ctx, loggerCtxKey{}, ctxMap)
}

// WithValues returns a new context with all the key-value pairs added for logging purposes.
// This should be used instead of WithLoggerValue when you want to add multiple values, as less cloning operations
// will be performed.
// Read WithValue for more information.
func WithValues(ctx context.Context, values map[string]string) context.Context {
	ctxMap := cloneCtxMap(ctx)
	for key, val := range values {
		ctxMap[key] = slog.String(key, val)
	}
	ctx = context.WithValue(ctx, loggerCtxKey{}, ctxMap)
	return ctx
}

func cloneCtxMap(ctx context.Context) loggerCtxMap {
	if ctxMap, ok := mapFromCtx(ctx); ok {
		return maps.Clone(ctxMap)
	}
	return loggerCtxMap{}
}

// mapFromCtx is a simple alias function to ctxMap, ok := ctx.Value(loggerCtxKey{}).(loggerCtxMap)
func mapFromCtx(ctx context.Context) (loggerCtxMap, bool) {
	ctxMap, ok := ctx.Value(loggerCtxKey{}).(loggerCtxMap)
	return ctxMap, ok
}
