package main

import (
  "github.com/kataras/iris/v12"
)

func main() {
  app := iris.New()
  app.Get("/ping", func(c iris.Context) {
    c.JSON(iris.Map{
      "message": "pong",
    })
  })
  app.Listen(":8080")
}
