package main

import (
  "log"
  "net/http"
)

func test020201() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("url path:%v\n", r.URL.Path)
    http.ServeFile(w, r, "wwwroot"+r.URL.Path)
  })

  http.ListenAndServe(":8080", nil)
}

func test020202() {
  http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
}
func main() {

}
