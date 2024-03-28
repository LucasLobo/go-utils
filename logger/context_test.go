package logger_test

import (
	"context"
	"log/slog"
	"testing"

	"go-utils/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithValue(t *testing.T) {
	t.Run("context map contains value set in context", func(t *testing.T) {
		ctx := logger.WithValue(context.Background(), "a", "1")
		assert.NotNil(t, ctx)

		actual, ok := logger.GetCtxMap(ctx)
		assert.True(t, ok)

		expectedMap := logger.LoggerCtxMap{
			"a": slog.String("a", "1"),
		}
		assert.EqualValues(t, expectedMap, actual)
	})
	t.Run("context map value is replaced in child context", func(t *testing.T) {
		parentCtx := logger.WithValue(context.Background(), "a", "parent")
		assert.NotNil(t, parentCtx)

		childCtx := logger.WithValue(parentCtx, "a", "child")
		assert.NotNil(t, childCtx)

		actual, ok := logger.GetCtxMap(childCtx)
		assert.True(t, ok)

		expectedMap := logger.LoggerCtxMap{
			"a": slog.String("a", "child"),
		}
		assert.EqualValues(t, expectedMap, actual)
	})
	t.Run("child context does not leak into parent context", func(t *testing.T) {
		// Assert
		parentCtx := logger.WithValue(context.Background(), "a", "parent")
		parentMap, ok := logger.GetCtxMap(parentCtx)
		require.True(t, ok)
		_ = logger.WithValue(parentCtx, "a", "child")

		expectedParent := logger.LoggerCtxMap{
			"a": slog.String("a", "parent"),
		}
		// parent maps should not change after adding values to a child context!
		assert.Equal(t, logger.LoggerCtxMap(parentMap), expectedParent)
		// not even after getting ctx map again!
		parentMap, ok = logger.GetCtxMap(parentCtx)
		require.True(t, ok)
		assert.Equal(t, logger.LoggerCtxMap(parentMap), expectedParent)
	})
}

func TestWithLoggerValues(t *testing.T) {
	t.Run("context map contains values set in context", func(t *testing.T) {
		ctx := logger.WithValues(context.Background(), map[string]string{"a": "1", "b": "2"})
		assert.NotNil(t, ctx)

		actual, ok := logger.GetCtxMap(ctx)
		assert.True(t, ok)

		expectedMap := logger.LoggerCtxMap{
			"a": slog.String("a", "1"),
			"b": slog.String("b", "2"),
		}
		assert.EqualValues(t, expectedMap, actual)
	})
	t.Run("context map values are replaced in child context", func(t *testing.T) {
		parentCtx := logger.WithValues(context.Background(), map[string]string{"a": "parent", "b": "parent"})
		assert.NotNil(t, parentCtx)

		childCtx := logger.WithValues(parentCtx, map[string]string{"a": "child", "b": "child"})
		assert.NotNil(t, childCtx)

		actual, ok := logger.GetCtxMap(childCtx)
		assert.True(t, ok)

		expectedMap := logger.LoggerCtxMap{
			"a": slog.String("a", "child"),
			"b": slog.String("b", "child"),
		}
		assert.EqualValues(t, expectedMap, actual)
	})
	t.Run("child context does not leak into parent context", func(t *testing.T) {
		// Assert
		parentCtx := logger.WithValues(context.Background(), map[string]string{"a": "parent", "b": "parent"})
		parentMap, ok := logger.GetCtxMap(parentCtx)
		require.True(t, ok)
		_ = logger.WithValues(parentCtx, map[string]string{"a": "child", "b": "child"})

		expectedParent := logger.LoggerCtxMap{
			"a": slog.String("a", "parent"),
			"b": slog.String("b", "parent"),
		}
		// parent maps should not change after adding values to a child context!
		assert.Equal(t, logger.LoggerCtxMap(parentMap), expectedParent)
		// not even after getting ctx map again!
		parentMap, ok = logger.GetCtxMap(parentCtx)
		require.True(t, ok)
		assert.Equal(t, logger.LoggerCtxMap(parentMap), expectedParent)
	})
}
