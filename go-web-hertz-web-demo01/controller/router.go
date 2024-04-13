package controller

import "github.com/cloudwego/hertz/pkg/app/server"

func RegisterHadlder(h *server.Hertz) {
  h.GET("/PingHandler", PingHandler)
  h.POST("/PersonBind", PersonBind)
}
