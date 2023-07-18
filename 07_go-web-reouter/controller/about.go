package controller

import (
  "fmt"
  "net/http"
)

func registerAboutRouter() {
  http.HandleFunc("/about", handleAbout())
}

func handleAbout() func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, r.URL.Path)
  }
}
