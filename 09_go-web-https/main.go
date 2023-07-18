package main

import (
  "fmt"
  "net/http"
)

func main() {
  listenHttps()
}

func listenHttps() {
  http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintln(writer, "Hello World")
  })
  //certFile := "cert.pem"
  //keyFile := "key.pem"
  //http.ListenAndServeTLS("localhost:8080", certFile, keyFile, nil)
  http.ListenAndServe("localhost:8080", nil)
}
