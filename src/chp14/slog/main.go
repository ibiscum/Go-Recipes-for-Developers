package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

func LoggingWithLevels() {

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	logger.Debug("This is a debug message")
	logger.Info("This is an info message with an integer argument", "arg", 42)
	logger.Warn("This is a warning message with a string argument", "arg", "foo")

	// Checking if logging is enabled for a specific level
	if logger.Enabled(context.Background(), slog.LevelError) {
		logger.Error("This is an error message", slog.String("arg", "foo"))
	}
}

func ChangingLogLevel() {

	level := &slog.LevelVar{}
	level.Set(slog.LevelDebug)
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))

	logger.Debug("This is a debug message")
	logger.Info("This is an info message with an integer argument", "arg", 42)
	logger.Warn("This is a warning message with a string argument", "arg", "foo")

	// Set the log level to info
	level.Set(slog.LevelInfo)
	logger.Debug("This is a debug message")
	logger.Info("This is an info message with an integer argument", "arg", 42)
	logger.Warn("This is a warning message with a string argument", "arg", "foo")

}

type ContextIDHandler struct {
	slog.Handler
}

func (h ContextIDHandler) Handle(ctx context.Context, r slog.Record) error {
	// If the context has a string id, retrieve it and add it to the record
	if id, ok := ctx.Value("id").(string); ok {
		r.Add(slog.String("id", id))
	}
	return h.Handler.Handle(ctx, r)
}

func JSONLogging() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	},
	))

	logger.Info("This is a debug message in JSON",
		slog.Time("now", time.Now()),
		slog.String("stringValue", "bar"),
		slog.Duration("durationValue", time.Second))
	logger.Info("This is an info message in JSON", slog.Time("now", time.Now()))

}

func AddingContextInfo() {
	logger := slog.New(&ContextIDHandler{
		Handler: slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
		),
	})

	// Create a new context with an id
	ctx := context.WithValue(context.Background(), "id", "123")

	logger.Info("This is an info message without a context id")
	logger.InfoContext(ctx, "This is an info message without a context id")
}

func Grouping() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	},
	))

	logger.Debug("This is a debug message with no group", slog.String("key1", "value1"))
	logger.Debug("This is a debug message with group as argument", slog.Group("group1", slog.String("key1", "value1")))
	l1 := logger.WithGroup("group2")
	l1.Debug("This is a debug message in group2", slog.String("key2", "value2"))

}

func WithAttributes() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	},
	))

	logger.Debug("This is a debug message with no additional attributes", slog.String("key", "value"))

	l1 := logger.With(slog.String("handler", "a"), slog.String("reqId", "reqId"))
	l1.Debug("This is a debug message with id", slog.String("key", "value"))

}

func main() {
	LoggingWithLevels()
	ChangingLogLevel()
	JSONLogging()
	AddingContextInfo()
	Grouping()
	WithAttributes()
}
