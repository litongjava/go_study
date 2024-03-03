package main

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
  h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

  v1 := h.Group("/v1")
  v1.GET("/get", func(ctx context.Context, c *app.RequestContext) {
    c.String(consts.StatusOK, "get")
  })

  v1.POST("/post", func(ctx context.Context, c *app.RequestContext) {
    c.String(consts.StatusOK, "post")
  })
  v2 := h.Group("/v2")
  v2.PUT("/put", func(ctx context.Context, c *app.RequestContext) {
    c.String(consts.StatusOK, "put")
  })
  v2.DELETE("/delete", func(ctx context.Context, c *app.RequestContext) {
    c.String(consts.StatusOK, "delete")
  })
  h.Spin()
}
