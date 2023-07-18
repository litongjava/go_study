package main

import (
  "fmt"
  "net/http"
)

func main() {
  server := http.Server{
    Addr: "localhost:8080",
  }
  http.HandleFunc("/", indexProcess)
  server.ListenAndServe()
}

func indexProcess(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hello,world", r.URL.Path)
}
