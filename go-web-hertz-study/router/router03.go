package main

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
  h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
  // This handler will match: "/hertz/version", but will not match : "/hertz/" or "/hertz"
  h.GET("/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
    version := c.Param("version")
    c.String(consts.StatusOK, "Hello %s", version)
  })
  h.Spin()
}
