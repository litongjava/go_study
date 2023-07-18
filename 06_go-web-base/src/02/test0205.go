package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
  str := `<html>
    <head><title>Go Web</title></head>
    <body><h1>Hello World</h1></body>
    </html>`
  w.Write([]byte(str))
}
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(501)
  fmt.Fprintln(w, "No such service,try next door")
}
func headerExample(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Location", "http://ppnt.top")
  w.WriteHeader(302)
}

type Post struct {
  User    string
  Threads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  post := &Post{
    User:    "Ping E Lee",
    Threads: []string{"first", "second", "third"},
  }
  json, _ := json.Marshal(post)
  w.Write(json)
}
func test020501() {
  server := http.Server{
    Addr: "localhost:8080",
  }
  http.HandleFunc("/write", writeExample)
  http.HandleFunc("/writeHeader", writeHeaderExample)
  http.HandleFunc("/redirect", headerExample)
  http.HandleFunc("/json", jsonExample)
  server.ListenAndServe()
}
func main() {
  test020501()
}
