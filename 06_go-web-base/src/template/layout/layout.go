package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
)

func init() {
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
func main() {
  test01()
}

func test01() {
  http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
    files, err := template.ParseFiles("layout.html", "home.html")
    if err != nil {
      fmt.Println(err)
    }
    //因为在html中使用define定义了模板名称,所以在调用时可以去掉后缀
    files.ExecuteTemplate(w, "layout", "Hello World")
  })
  http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    files, err := template.ParseFiles("layout.html", "about.html")
    if err != nil {
      fmt.Println(err)
    }
    //因为在html中使用define定义了模板名称,所以在调用时可以去掉后缀
    files.ExecuteTemplate(w, "layout", "Hello World")
  })
  http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
    files, err := template.ParseFiles("layout.html")
    if err != nil {
      fmt.Println(err)
    }
    //因为在html中使用define定义了模板名称,所以在调用时可以去掉后缀
    err = files.ExecuteTemplate(w, "layout", "")
    log.Println(err)
  })
  http.ListenAndServe("localhost:8080", nil)
}
