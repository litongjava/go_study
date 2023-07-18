package main

import (
  "fmt"
  "html/template"
  "net/http"
)

func main() {
  test01()
}

func test01() {
  var server = http.Server{Addr: "localhost:8080"}
  http.HandleFunc("/", process01)
  server.ListenAndServe()
}

func process01(w http.ResponseWriter, r *http.Request) {
  files, err := template.ParseFiles("template-demo.html")
  if err != nil {
    fmt.Println(err)
  }
  files.Execute(w, "Hello")
}
