package main

import (
  "07.go-web-reouter/controller"
  "net/http"
)

func main() {
  controller.RegisterRoutes()
  http.ListenAndServe("localhost:8080", nil)
}
