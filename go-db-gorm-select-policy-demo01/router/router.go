package router

import (
  "github.com/cloudwego/hertz/pkg/app/server"
  "go-db-gorm-select-policy-demo01/handler"
)

func Register(hertz *server.Hertz) {
  hertz.GET("/ping", handler.Ping)
  hertz.GET("/time", handler.Time)
  hertz.GET("/feedback/list", handler.FeedbackList)
  hertz.GET("/ResponseTimesMap", handler.GetResponseTimesMap)
}
