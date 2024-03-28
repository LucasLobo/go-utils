package logger_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"go-utils/logger"
)

func TestStuff(t *testing.T) {

	var programLevel = new(slog.LevelVar) // Info by default
	programLevel.Set(slog.LevelDebug)

	handler := logger.NewCtxHandler(os.Stderr, &slog.HandlerOptions{
		Level: programLevel,
	})

	l := slog.New(handler)
	slog.SetDefault(l)

	ctx := context.Background()
	ctx = logger.WithValue(ctx, "custom_key_1", "custom_val_1")

	slog.DebugContext(ctx, "Hello, World!")
}
