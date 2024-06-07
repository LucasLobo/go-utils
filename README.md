# go-utils

go-utils is a library that contains a collection of my personal Golang utilities

# Features

## Logger

The `logger` package provides an extension to the `slog.JSONHandler` that adds context information to the logs.

### Usage

```go

// create the logger using the custom CtxHandler using [logger.NewCtxHandler]
l := slog.New(logger.NewCtxHandler(os.Stdout, &slog.HandlerOptions{}))

// add a custom key-value pair to the ctx using [logger.WithValue]
ctx := logger.WithValue(context.Background(), "custom_key", "custom_val")

// log a message using the custom ctx, the message will contain the custom key-value pair
l.InfoContext(ctx, "Hello, World!")

```

## ...
More to come
