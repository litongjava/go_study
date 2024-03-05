package main

import (
  "github.com/cloudwego/hertz/pkg/app/server"
  "go-db-gorm-select-policy-demo01/db"
  "go-db-gorm-select-policy-demo01/router"
)

func main() {
  print("start")
  hertz := server.Default()
  db.Init()
  router.Register(hertz)
  hertz.Spin()
  print("end")
}
