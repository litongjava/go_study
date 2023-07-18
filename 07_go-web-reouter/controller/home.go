package controller

import (
  "fmt"
  "net/http"
)

func registerHomeRouters() {
  http.HandleFunc("/home", handleHome())
}

func handleHome() func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, r.URL.Path)
  }
}
