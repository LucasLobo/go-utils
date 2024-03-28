package logger_test

import (
	"context"
	"log/slog"
	"os"

	"go-utils/logger"
)

func ExampleCtxHandler() {

	// defining the level allows us to change the log level at runtime
	var programLevel = new(slog.LevelVar) // Info by default
	programLevel.Set(slog.LevelDebug)

	handler := logger.NewCtxHandler(os.Stdout, &slog.HandlerOptions{
		Level:       programLevel,
		ReplaceAttr: removeTime, // Remove time from log output otherwise tests will fail
	})

	l := slog.New(handler)
	slog.SetDefault(l)

	ctx := context.Background()
	ctx = logger.WithValue(ctx, "custom_key", "custom_val")

	slog.DebugContext(ctx, "Hello, World!")

	// Output:
	// {"level":"DEBUG","msg":"Hello, World!","custom_key":"custom_val"}
}

// removeTime is a function that removes the time attribute from the log output. This is only required because we're
// comparing the full output and the time will always be different.
func removeTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}
