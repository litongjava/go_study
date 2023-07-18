package main

import (
  "html/template"
  "net/http"
  "time"
)

func main() {
  test01()
}

func test01() {
  var server http.Server = http.Server{Addr: "localhost:8080"}

  http.HandleFunc("/process", process01)
  server.ListenAndServe()

}

func process01(w http.ResponseWriter, r *http.Request) {
  var funcMap template.FuncMap = template.FuncMap{"fdate": formatDate}
  var name string = "template-demo.html"
  var t *template.Template = template.New(name).Funcs(funcMap)
  t.ParseFiles(name)
  t.Execute(w, time.Now())
}

func formatDate(t time.Time) string {
  var layout = "2006-01-02"
  return t.Format(layout)
}
