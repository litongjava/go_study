package main

import (
  "html/template"
  "log"
  "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("tmpl.html")
  if err != nil {
    //open tmpl.html: The system cannot find the file specified.
    log.Println(err)
  } else {
    t.Execute(w, "Hello World!")
  }

}

func test020601() {
  server := http.Server{
    Addr: "localhost:8080",
  }
  http.HandleFunc("/process", process)
  server.ListenAndServe()
}

func main() {
  test020601()
}
