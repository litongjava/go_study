// github.com/bigwhite/experiments/tree/master/slog-examples/demo2/main.go
package main

import (
	"log/slog"
	"net"
	"os"
)

func main() {
	opts := slog.HandlerOptions{
		AddSource: true,
	}
	handler := slog.NewTextHandler(os.Stderr, &opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
}
