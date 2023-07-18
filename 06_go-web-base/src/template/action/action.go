package main

import (
  "fmt"
  "html/template"
  "math/rand"
  "net/http"
  "time"
)

func main() {
  test04()
}
func test04() {
  server := http.Server{Addr: "localhost:8080"}
  http.HandleFunc("/process", process04)
  server.ListenAndServe()
}
func process04(w http.ResponseWriter, r *http.Request) {
  files, err := template.ParseFiles("tmpl04-1.html", "tmpl04-2.html")
  if err != nil {
    fmt.Println(err)
    return
  }
  files.Execute(w, "hello world")
}

func test03() {
  server := http.Server{Addr: "localhost:8080"}
  http.HandleFunc("/process", process03)
  server.ListenAndServe()
}
func process03(w http.ResponseWriter, r *http.Request) {
  files, err := template.ParseFiles("tmpl03.html")
  if err != nil {
    fmt.Println(err)
    return
  }
  files.Execute(w, "hello")
}
func test02() {
  server := http.Server{Addr: "localhost:8080"}
  http.HandleFunc("/process", process02)
  server.ListenAndServe()
}
func process02(w http.ResponseWriter, r *http.Request) {
  files, err := template.ParseFiles("tmpl02.html")
  if err != nil {
    fmt.Println(err)
    return
  }
  //dayOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
  dayOfWeek := []string{}
  files.Execute(w, dayOfWeek)
}

func test01() {
  server := http.Server{Addr: "localhost:8080"}
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
  files, _ := template.ParseFiles("tmpl.html")
  rand.Seed(time.Now().Unix())
  files.Execute(w, rand.Intn(10) > 5)
}
