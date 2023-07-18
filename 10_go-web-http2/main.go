package main

import (
  "10.go-web-http2/controller"
  "net/http"
)

func main() {
  controller.RegisterRoutes()
  //拦截css目录到public下面
  http.Handle("/css/", http.FileServer(http.Dir("public")))
  certFile := "cert.pem"
  keyFile := "key.pem"
  http.ListenAndServeTLS("localhost:8080", certFile, keyFile, nil)
}
