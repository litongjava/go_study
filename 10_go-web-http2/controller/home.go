package controller

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
)

func registerHomeRoutes() {
  http.HandleFunc("/home", handleHome)
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
  if pusher, ok := writer.(http.Pusher); ok {
    fmt.Println("push app.css")
    pusher.Push("/css/app.css", &http.PushOptions{
      Header: http.Header{"Content-Type": []string{"text/css"}},
    })
  } else {
    fmt.Println("not suppored push")
  }
  t, err := template.ParseFiles("public/home.html")
  if err != nil {
    //open tmpl.html: The system cannot find the file specified.
    log.Println(err)
  } else {
    t.Execute(writer, "Hello World!")
  }
}
