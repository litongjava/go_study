package main

import (
  "github.com/cloudwego/hertz/pkg/app/server"
  "go-web-hertz-web-demo01/controller"
)

func main() {
  ports := server.WithHostPorts("0.0.0.0:9001")
  h := server.Default(ports)
  controller.RegisterHadlder(h)
  h.Spin()
}
