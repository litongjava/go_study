package controller

import (
  "fmt"
  "net/http"
)

func registerContactRouter() {
  http.HandleFunc("/contact", handleAbout())
}

func handleContact() func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, r.URL.Path)
  }
}
