package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func test02040801() {
  server := http.Server{
    Addr: "localhost:8080",
  }

  http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
    //r.ParseForm()
    //fmt.Fprintln(w, r.Form)
    //fmt.Fprintln(w, r.PostForm)
    r.ParseMultipartForm(1024 * 1024)
    fmt.Fprintln(w, r.MultipartForm)
  })

  server.ListenAndServe()
}

func test02041101() {
  server := http.Server{
    Addr: "localhost:8080",
  }

  http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
    //r.ParseForm()
    //fmt.Fprintln(w, r.Form)
    //fmt.Fprintln(w, r.PostForm)
    r.ParseMultipartForm(1024)
    fmt.Fprintln(w, r.MultipartForm)
  })
  server.ListenAndServe()
}
func upload(w http.ResponseWriter, r *http.Request) {
  //r.ParseMultipartForm(1024)

  //读取第一个文件
  //fileHeader := r.MultipartForm.File["file"][0]
  //打开文件
  //file, err := fileHeader.Open()
  file, _, err := r.FormFile("file")
  if err == nil {
    data, err := ioutil.ReadAll(file)
    if err == nil {
      fmt.Fprintln(w, string(data))
    }
  }
}
func test020412() {
  server := http.Server{
    Addr: "localhost:8080",
  }
  http.HandleFunc("/process", upload)
  server.ListenAndServe()
}
func test020416() {

}
func main() {
  test020412()
}
