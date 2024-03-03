package main

import (
  "context"
  "github.com/cloudwego/hertz/pkg/app"
  "github.com/cloudwego/hertz/pkg/app/server"
  "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
  h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
  // However, this one will match "/hertz/v1/" and "/hertz/v2/send"
  h.GET("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
    version := c.Param("version")
    action := c.Param("action")
    message := version + " is " + action
    c.String(consts.StatusOK, message)
  })
  h.Spin()
}
