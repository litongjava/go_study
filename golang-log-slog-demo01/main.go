// github.com/bigwhite/experiments/tree/master/slog-examples/demo2/main.go
package main

import (
  "io"
  "log"
  "log/slog"
  "net"
  "os"
)

func main() {
  opts := slog.HandlerOptions{
    //AddSource: true,
  }
  conn, err := net.Dial("tcp", "127.0.0.1:5000")
  if err != nil {
    log.Fatal(err)
  }
  multiWriter := io.MultiWriter(conn, os.Stdout)
  handler := slog.NewTextHandler(multiWriter, &opts)
  logger := slog.New(handler)
  slog.SetDefault(logger)
  slog.Info("hello", "name", "Al")
  slog.Error("oops", net.ErrClosed, "status", 500)
}
